package mid

import (
	"fmt"
	"github.com/bits-and-blooms/bitset"
	"hash/fnv"
)

// BloomFilterGlobal 全局布隆过滤器实例
var BloomFilterGlobal *BloomFilter

// BloomFilter 布隆过滤器结构体
type BloomFilter struct {
	bitSet    *bitset.BitSet
	numHashes int  // 对应多少个哈希函数
	size      uint // 位数组大小
}

// InitBloomFilter 初始化布隆过滤器（可以在程序启动时调用）
func InitBloomFilter(size uint, numHashes int) {
	BloomFilterGlobal = NewBloomFilter(size, numHashes)
}

// NewBloomFilter 创建布隆过滤器
func NewBloomFilter(size uint, numHashes int) *BloomFilter {
	return &BloomFilter{
		bitSet:    bitset.New(size),
		numHashes: numHashes,
		size:      size,
	}
}

// Add 使用多个哈希函数将元素插入布隆过滤器
func (bf *BloomFilter) Add(item string) {
	hashes := bf.getHashes(item)
	for _, h := range hashes {
		bf.bitSet.Set(h % bf.size) // 映射到位数组中
	}
}

// MightContain 检查元素是否可能存在于布隆过滤器中
func (bf *BloomFilter) MightContain(item string) bool {
	hashes := bf.getHashes(item)
	for _, h := range hashes {
		if !bf.bitSet.Test(h % bf.size) {
			return false
		}
	}
	return true
}

func (bf *BloomFilter) getHashes(item string) []uint {
	hashes := make([]uint, bf.numHashes)
	for i := 0; i < bf.numHashes; i++ {
		hashes[i] = bf.hash(item, i)
	}
	return hashes
}

func (bf *BloomFilter) hash(item string, seed int) uint {
	h := fnv.New32a()
	_, err := h.Write([]byte(fmt.Sprintf("%s%d", item, seed)))
	if err != nil {
		return 0
	}
	return uint(h.Sum32())
}
