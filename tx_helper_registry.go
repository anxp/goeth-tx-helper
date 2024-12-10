package goeth_tx_helper

import (
	"fmt"
	"sync"
)

var txHelperRegistrySingleInstance *TxHelperRegistry
var initLock sync.Mutex

type TxHelperRegistry struct {
	registry   map[string]*EIP1559TransactionHelper
	accessLock sync.Mutex
}

func createTxHelperRegistry() *TxHelperRegistry {
	initLock.Lock()
	defer initLock.Unlock()

	if txHelperRegistrySingleInstance == nil {
		registry := make(map[string]*EIP1559TransactionHelper)

		txHelperRegistrySingleInstance = &TxHelperRegistry{
			registry: registry,
		}

		fmt.Printf("TxHelperRegistry Instance created. This action should be done only once (Singleton).\n")
	}

	return txHelperRegistrySingleInstance
}

func (thr *TxHelperRegistry) getTxHelper(rpcUrl string) *EIP1559TransactionHelper {
	thr.accessLock.Lock()
	defer thr.accessLock.Unlock()

	txHelper, ok := thr.registry[rpcUrl]

	if !ok {
		return nil
	}

	return txHelper
}

func (thr *TxHelperRegistry) addTxHelperToRegistry(txHelper *EIP1559TransactionHelper) {
	thr.accessLock.Lock()
	defer thr.accessLock.Unlock()

	if _, ok := thr.registry[txHelper.GetRpcUrl()]; ok {
		return
	}

	thr.registry[txHelper.GetRpcUrl()] = txHelper
}
