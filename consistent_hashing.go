package consistentHashing

import (
	"hash/crc32"
	"sync"
)

// hashing function
type HashFunc func(key []byte) uint32

type option func(*hashMgr)

type hashMgr struct {
	hashMap     map[string]uint32
	keys        units
	calc        HashFunc
	mutex       *sync.Mutex
	replicasNum int
	sorted      bool
}

func NewMgr(opt ...option) *hashMgr {
	defaultHash := &hashMgr{
		hashMap:     make(map[string]uint32),
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
