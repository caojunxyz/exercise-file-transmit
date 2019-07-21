package fileblock

import (
	"crypto/sha1"
	"fmt"
	"github.com/caojunxyz/superdeliver/deliver"
)

const (
	DefaultBlockSize = 2 * (1 << 20) // Byte
	// DefaultBlockSize = 2 * (1 << 10) // Byte
	// DefaultBlockSize = 4 // Byte
)

func newBlock(id int64, buf []byte) *deliver.Block {
	size := int64(len(buf))
	block := &deliver.Block{Id: id, Size: size}
	block.Raw = make([]byte, size, size)
	copy(block.Raw, buf[:size])
	block.Sha1 = fmt.Sprintf("%x", sha1.Sum(block.Raw))
	return block
}

func CheckBlock(block *deliver.Block) error {
	if block.Size != int64(len(block.Raw)) {
		return fmt.Errorf("size not match (%d - %d)", block.Size, len(block.Raw))
	}

	sum := fmt.Sprintf("%x", sha1.Sum(block.Raw))
	if block.Sha1 != sum {
		return fmt.Errorf("sha1 not match (%s - %s)", block.Sha1, sum)
	}
	return nil
}
