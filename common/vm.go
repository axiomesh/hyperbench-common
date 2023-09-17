package common

import (
	"context"
)

// TxContext is the context of transaction (`vm.Run`)
type TxContext struct {
	context.Context
	TxIndex
}

// TxIndex is the unique index of a transaction (`vm.Run`)
type TxIndex struct {
	TxIdx   int64 // TxIdx is the index of sent transaction
	MissIdx int64 // MissIdx is the index of missed transaction
}

// VMContext is the context of vm
type VMContext struct {
	WorkerIdx int64 `mapstructure:"worker"`   // WorkerIdx is the index of worker
	VMIdx     int64 `mapstructure:"vm"`       // VMIdx is the index of vm
	Engine    int64 `mapstructure:"engine"`   // Engine is the total engine(vm) num
	Accounts  int64 `mapstructure:"accounts"` // Accounts is the count of acutal test accounts
}
