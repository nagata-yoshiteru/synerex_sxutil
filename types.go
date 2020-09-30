package sxutil

import (
	"sync"

	api "github.com/synerex/synerex_api"
)

type RequestVars struct {
	CrMu            *sync.Mutex
	CrLoop          *bool
	SxServerAddress string
	CurrentNid      uint64
	StorageID       uint64
	ProSpAdmin      map[string][]*api.Supply
	ObjNames        []string
	AreaName        string
	Data            *[]byte //= make(chan []byte)
}
