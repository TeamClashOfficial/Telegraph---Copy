package api

import (
	"sync"
	"telegraph/binance"
)

var (
	SyncUrlMap  *sync.Map // used for msgs routing between nodes
	ReshareFlag bool      // flag to check choose between keygen and reshare
)

func init() {
	SyncUrlMap = new(sync.Map)
}

func init() {
	ReshareFlag = false // resharing is disabled
}

func UpdateSyncUrlMap(m map[string]binance.SortedIP) {
	// SyncUrlMap = MapToSync(m)
}
