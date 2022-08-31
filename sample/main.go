package main

import (
	"fmt"

	"github.com/yintamaa/consistentHashing"
)

func main() {
	mp := make(map[string]int)
	hashMgr := consistentHashing.NewMgr(consistentHashing.WithReplicaNum(1000))
	hashMgr.Add("1", "2", "3", "4")
	for i := 0; i < 100; i++ {
		id := hashMgr.Get(fmt.Sprintf("%d", i))
		mp[id]++
	}
	for k, v := range mp {
		fmt.Println(k, v)
	}
	fmt.Println("over 1")
	hashMgr.Remove("4")
	for k := range mp {
		delete(mp, k)
	}
	for i := 0; i < 100; i++ {
		id := hashMgr.Get(fmt.Sprintf("%d", i))
		mp[id]++
	}
	for k, v := range mp {
		fmt.Println(k, v)
	}
	fmt.Println("over 2")
	fmt.Println(hashMgr.Get("100"))
	fmt.Println(hashMgr.Get("100"))
	fmt.Println(hashMgr.Get("100"))
}
