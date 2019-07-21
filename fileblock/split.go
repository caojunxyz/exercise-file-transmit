package fileblock

import (
	"crypto/sha1"
	"fmt"
	"github.com/caojunxyz/superdeliver/deliver"
	"io"
	"log"
	"os"
)

type SplitController struct {
	Artifact   *deliver.Artifact
	CurBlockId int64
	Done       chan struct{}
	Blocks     chan *deliver.Block
	file       *os.File
}

func NewSplitController(filename string, n int) *SplitController {
	return &SplitController{
		Artifact: &deliver.Artifact{Filename: filename},
		CurBlockId: 0, // id starts from 0
		Done:     make(chan struct{}),
		Blocks:   make(chan *deliver.Block, n),
	}
}

func (c *SplitController) Init() error {
	artifact := c.Artifact
	file, err := os.Open(artifact.Filename)
	if err != nil {
		log.Println(err)
		return err
	}
	c.file = file

	if artifact.Sha1 == "" {
		h := sha1.New()
		n, err := io.Copy(h, c.file)
		if err != nil {
			log.Println(err)
			return err
		}
		artifact.Size = n
		artifact.Sha1 = fmt.Sprintf("%x", h.Sum(nil))
	}

	if artifact.BlockSize == 0 {
		artifact.BlockSize = DefaultBlockSize
		if artifact.BlockSize > artifact.Size {
			artifact.BlockSize = artifact.Size
		}
	}

	if artifact.BlockNum == 0 {
		artifact.BlockNum = artifact.Size / artifact.BlockSize
		if artifact.Size%artifact.BlockSize > 0 {
			artifact.BlockNum += 1
		}
	}
	return nil
}

func (c *SplitController) Start() error {
	var n int
	var err error

	go func() {
		defer func() {
			if err := recover(); err != nil {
				log.Println("recover:", err)
			}
		}()

		defer func() {
			close(c.Done)
		}()

		log.Println("start splitting...")
		artifact := c.Artifact
		buf := make([]byte, artifact.BlockSize, artifact.BlockSize)
		offset := artifact.BlockSize * c.CurBlockId
		c.file.Seek(offset, io.SeekStart)
		for {
			n, err = c.file.Read(buf)
			if err != nil && err != io.EOF {
				log.Println(err)
				return
			}

			if n > 0 {
				c.Blocks <- newBlock(c.CurBlockId, buf[:n])
				c.CurBlockId += 1
			}

			if c.CurBlockId == artifact.BlockNum || err == io.EOF {
				return
			}
		}
	}()

	return nil
}
