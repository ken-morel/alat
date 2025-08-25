// Package node: the core's core
package node

import (
	"alat/pkg/core/discovery"
	"alat/pkg/core/pair"
	"alat/pkg/core/storage"
)

type Node struct {
	Storage     storage.NodeStorage
	PairManager *pair.PairManager
	discovery   *discovery.Manager
}
