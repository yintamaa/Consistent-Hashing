package consistentHashing

import (
	"fmt"
	"sort"
	"strconv"
)

func (hm *HashMgr) Add(keys ...string) {
	hm.mutex.Lock()
	for _, key := range keys {
		for j := 0; j < hm.replicasNum; j++ {
			tempVal := hm.calc([]byte(key + strconv.Itoa(j)))
			hashVal := hm.calc([]byte(fmt.Sprintf("%d", tempVal)))
			if _, ok := hm.hashMap[hashVal]; ok { // key has existed
				break
			}
			hm.hashMap[hashVal] = key
			hm.keys = append(hm.keys, hashVal)
		}
	}
	hm.sorted = false
	hm.mutex.Unlock()
}

func (hm *HashMgr) Get(key string) (string, bool) {
	hashVal := hm.calc([]byte(key))
	hm.mutex.Lock()
	defer hm.mutex.Unlock()
	if !hm.sorted {
		sort.Sort(hm.keys)
		hm.sorted = true
	}
	sliceLen := len(hm.keys)
	if sliceLen == 0 {
		return "", false
	}
	idx := sort.Search(sliceLen, func(i int) bool {
		return hm.keys[i] >= hashVal
	}) % sliceLen
	return hm.hashMap[hm.keys[idx]], true
}

func (hm *HashMgr) Remove(keys ...string) bool {
	hm.mutex.Lock()
	defer hm.mutex.Unlock()
	if !hm.sorted {
		sort.Sort(hm.keys)
		hm.sorted = true
	}
	for _, key := range keys {
		for j := 0; j < hm.replicasNum; j++ {
			tempVal := hm.calc([]byte(key + strconv.Itoa(j)))
			hashVal := hm.calc([]byte(fmt.Sprintf("%d", tempVal)))
			sliceLen := len(hm.keys)
			if sliceLen == 0 {
				return false
			}
			idx := sort.Search(sliceLen, func(i int) bool {
				return hm.keys[i] >= hashVal
			}) % sliceLen
			if hm.keys[idx] != hashVal { // 不合法
				return false
			}
			hm.keys = append(hm.keys[0:idx], hm.keys[idx+1:len(hm.keys)]...)
			delete(hm.hashMap, hashVal)
		}
		hm.sorted = false
	}
	return true
}
