package consistentHashing

import (
	"hash/crc32"
	"sync"
)

// hashing function
type HashFunc func(key []byte) uint32

type option func(*HashMgr)

type HashMgr struct {
	hashMap     map[uint32]string // hashVal -> keys
	keys        uints             // hashval slice
	calc        HashFunc
	mutex       *sync.Mutex
	replicasNum int
	sorted      bool
}

// create a consistent hashing instance
func NewMgr(opt ...option) *HashMgr {
	defaultHash := &HashMgr{
		hashMap:     make(map[uint32]string),
		keys:        make([]uint32, 0),
		calc:        crc32.ChecksumIEEE,
		mutex:       &sync.Mutex{},
		replicasNum: 1,
		sorted:      false,
	}
	for _, v := range opt {
		v(defaultHash)
	}
	return defaultHash
}
