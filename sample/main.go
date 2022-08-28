package main

import (
	"fmt"

	"github.com/yintamaa/consistentHashing"
)

func main() {
	mp := make(map[string]int)
	hashMgr := consistentHashing.NewMgr(consistentHashing.WithReplicaNum(100))
	hashMgr.Add("1", "2", "3", "4", "5", "6")
	for i := 0; i < 100; i++ {
		id := hashMgr.Get(fmt.Sprintf("%d", i))
		mp[id]++
	}
	for k, v := range mp {
		fmt.Println(k, v)
	}
}
