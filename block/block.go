package block

import (
	"crypto/sha256"
	"encoding/hex"
	"time"
)

/**
Index 是这个块在整个链中的位置
Timestamp 显而易见就是块生成时的时间戳
Hash 是这个块通过 SHA256 算法生成的散列值
PrevHash 代表前一个块的 SHA256 散列值
BPM 每分钟心跳数，也就是心率。示例数据
**/
type Block struct {
	Index int
	Timestamp string
	BPM int
	Hash string
	PrevHash string
}

var Blockchain []Block

func CalculateHash(block Block) string {
	record := string(block.Index) + block.Timestamp + string(block.BPM) + block.PrevHash
	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}

func GenerateBlock(oldBlock Block, BPM int) (Block, error) {
	var newBlock Block

	t := time.Now()
	newBlock.Index = oldBlock.Index + 1
	newBlock.Timestamp = t.String()
	newBlock.BPM = BPM
	newBlock.PrevHash = oldBlock.Hash
	newBlock.Hash = CalculateHash(newBlock)

	return newBlock, nil
}

func IsBlockValid(newBlock, oldBlock Block) bool {
	if oldBlock.Index+1 != newBlock.Index {
		return false
	}
	if oldBlock.Hash != newBlock.PrevHash {
		return false
	}
	if CalculateHash(newBlock) != newBlock.Hash {
		return false
	}
	return true
}

func ReplaceChain(newBlocks []Block) {
	if len(newBlocks) > len(Blockchain) {
		Blockchain = newBlocks
	}
}