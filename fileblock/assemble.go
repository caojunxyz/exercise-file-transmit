package fileblock

import (
	"crypto/sha1"
	"fmt"
	"github.com/caojunxyz/superdeliver/deliver"
	"github.com/caojunxyz/superdeliver/utils"
	"io"
	"log"
	"os"
	"sync"
	"time"
)

type AssembleController struct {
	sync.RWMutex
	Artifact         *deliver.Artifact
	AssembledFlags   map[int64]bool
	IsComplete       bool
	chanWaitAssemble chan *deliver.Block // blocks waiting to be assembled
	chanAssembled    chan *deliver.BlockAssembled // finish assembling, wait for broadcast
	file             *os.File
}

func NewAssembleController(artifact *deliver.Artifact, ch chan *deliver.BlockAssembled) *AssembleController {
	c := &AssembleController{
		Artifact:         artifact,
		AssembledFlags:   make(map[int64]bool),
		chanWaitAssemble: make(chan *deliver.Block, 50),
		chanAssembled:    ch,
	}
	if err := c.createDummyFile(); err != nil {
		log.Panic(err)
	}
	go c.assemble()
	return c
}

func (c *AssembleController) createDummyFile() error {
	t0 := time.Now()

	artifact := c.Artifact
	bn := artifact.Size / artifact.BlockSize // body block num
	tn := artifact.BlockNum - bn             // tail block num

	file, err := os.Create(artifact.Filename)
	if err != nil {
		log.Println(err)
		return err
	}

	buf := make([]byte, artifact.BlockSize, artifact.BlockSize)
	for i := int64(0); i < bn; i++ {
		_, err := file.WriteAt(buf, i*artifact.BlockSize)
		if err != nil {
			log.Println(err)
			return err
		}
	}

	if tn > 0 {
		n := artifact.Size - artifact.BlockSize*bn
		_, err := file.WriteAt(buf[:n], (artifact.BlockNum-1)*artifact.BlockSize)
		if err != nil {
			log.Println(err)
			return err
		}
	}

	if err := file.Sync(); err != nil {
		log.Panic(err)
	}

	c.file = file
	log.Printf("create dummy file cost: %v\n", time.Now().Sub(t0))
	return nil
}

func (c *AssembleController) IsBlockAssembled(blockId int64) bool {
	c.RLock()
	defer c.RUnlock()
	return c.AssembledFlags[blockId]
}

func (c *AssembleController) ReadBlock(id int64) (*deliver.Block, error) {
	if !c.IsBlockAssembled(id){
		return nil, fmt.Errorf("block #%d not found!", id)
	}

	c.Lock()
	defer c.Unlock()
	_, err := c.file.Seek(c.Artifact.BlockSize*id, io.SeekStart)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	buf := make([]byte, c.Artifact.BlockSize, c.Artifact.BlockSize)
	n, err := c.file.Read(buf)
	if err != nil && err != io.EOF {
		log.Println(err)
		return nil, err
	}
	return newBlock(id, buf[:n]), nil
}

func (c *AssembleController) CheckComplete() error {
	if !c.IsComplete && int64(len(c.AssembledFlags)) == c.Artifact.BlockNum {
		c.Lock()
		defer c.Unlock()
		// all blocks have been assembled
		c.file.Seek(0, io.SeekStart)
		h := sha1.New()
		n, err := io.Copy(h, c.file)
		if err != nil {
			log.Println(n, err)
			return err
		}

		if n != c.Artifact.Size {
			return fmt.Errorf("size not equal(%v, %v)", n, c.Artifact.Size)
		}

		sum := fmt.Sprintf("%x", h.Sum(nil))
		if sum != c.Artifact.Sha1 {
			return fmt.Errorf("sha1 not equal(%v, %v)", sum, c.Artifact.Sha1)
		}
		c.IsComplete = true
		log.Println("complete!")
	}
	return nil
}

func (c *AssembleController) SendToAssemble(block *deliver.Block) {
	if !c.IsBlockAssembled(block.Id) {
		c.chanWaitAssemble <- block
	}
}

func (c *AssembleController) assemble() {
	defer func() {
		if err := recover(); err != nil {
			log.Println(err)
		}
	}()

	log.Println("assemble...")
	for {
		select {
		case block := <-c.chanWaitAssemble:
			if err := CheckBlock(block); err != nil {
				log.Panic(err)
			}
			c.Lock()
			if _, err := c.file.Seek(0, io.SeekStart); err != nil {
				log.Panic(err)
			}
			_, err := c.file.WriteAt(block.Raw, c.Artifact.BlockSize*block.Id)
			if err != nil {
				log.Panic(err)
			}
			if err := c.file.Sync(); err != nil {
				log.Panic(err)
			}
			c.AssembledFlags[block.Id] = true
			c.Unlock()

			if c.chanAssembled != nil {
				c.chanAssembled <- &deliver.BlockAssembled{Id: block.Id, PulledFrom: block.PulledFrom}
			}
			log.Printf("Block#%d (%d/%d) (%s)\n", block.Id, len(c.AssembledFlags), c.Artifact.BlockNum, utils.BlockFormatString(block))
			if err := c.CheckComplete(); err != nil {
				log.Panic(err)
			}
		}
	}
}

func (c *AssembleController) Progress() int64 {
	c.RLock()
	defer c.RUnlock()
	return int64(len(c.AssembledFlags))
}