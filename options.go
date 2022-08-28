package consistentHashing

// seting the ReplicaNum of key
func WithReplicaNum(num int) option {
	return func(hm *hashMgr) {
		hm.replicasNum = num
	}
}

// seting the hash function
func WithHashFunc(hf HashFunc) option {
	return func(hm *hashMgr) {
		hm.calc = hf
	}
}
