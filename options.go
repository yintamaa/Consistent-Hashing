package consistentHashing

func WithReplicaNum(num int) option {
	return func(hm *hashMgr) {
		hm.replicasNum = num
	}
}

func WithHashFunc(hf HashFunc) option {
	return func(hm *hashMgr) {
		hm.calc = hf
	}
}
