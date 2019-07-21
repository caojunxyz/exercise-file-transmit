package utils

import (
	"fmt"
	"github.com/caojunxyz/superdeliver/deliver"
)

func BlockFormatString(block *deliver.Block) string {
	ri := 10
	if ri > len(block.Raw) {
		ri = len(block.Raw)
	}
	return fmt.Sprintf("Block#%d (size: %d bytes sha1: %s raw: %s", block.Id, block.Size, block.Sha1, fmt.Sprintf("%s", block.Raw[:ri]))
}

func ProgressFormatString(progress *deliver.Progress, blockNum int64) string {
	result := ""
	for k, v := range progress.AssembledBlocks {
		s := fmt.Sprintf("node-%d: %d/%d\n", k, v, blockNum)
		result += s
	}
	return result
}