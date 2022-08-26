package consistentHashing

import (
	"sort"
	"strconv"
)

func (hm *hashMgr) Add(keys ...string) {
	for _, key := range keys {
		hm.mutex.Lock()
		if _, ok := hm.hashMap[key]; ok { // key has existed
			continue
		}
		for j := 0; j < hm.replicasNum; j++ {
			hashVal := hm.calc([]byte(key + strconv.Itoa(j)))
			hm.hashMap[key] = hashVal
			hm.keys = append(hm.keys, hashVal)
		}
		hm.mutex.Unlock()
	}
}

func (hm *hashMgr) Get(key string) {
	hashVal := hm.calc([]byte(key))
	hm.mutex.Lock()
	if !hm.sorted {
		sort.Sort(hm.keys)
	}
}

func (hm *hashMgr) Remove() {

}
