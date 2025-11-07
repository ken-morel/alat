package main

import "C"

// export get_node_port
func get_node_port(handle C.int) C.int {
	if instance := getInstance(handle); instance != nil {
		return C.int(instance.node.GetPort())
	} else {
		return -1
	}
}
