package main

import "C"
import (
	"sync"

	"alat/pkg/core/pair"
)

var (
	pairManagers      = make(map[int]*pair.PairManager)
	pairManagersMutex = &sync.Mutex{}
	nextPairManagerID = 1
)
