// This file was generated by counterfeiter
package fakes

import (
	"sync"

	"github.com/cloudfoundry-incubator/guardian/rundmc"
)

type FakeExitStore struct {
	StoreStub        func(handle string, exit <-chan struct{})
	storeMutex       sync.RWMutex
	storeArgsForCall []struct {
		handle string
		exit   <-chan struct{}
	}
	UnstoreStub        func(handle string)
	unstoreMutex       sync.RWMutex
	unstoreArgsForCall []struct {
		handle string
	}
	WaitStub        func(handle string)
	waitMutex       sync.RWMutex
	waitArgsForCall []struct {
		handle string
	}
}

func (fake *FakeExitStore) Store(handle string, exit <-chan struct{}) {
	fake.storeMutex.Lock()
	fake.storeArgsForCall = append(fake.storeArgsForCall, struct {
		handle string
		exit   <-chan struct{}
	}{handle, exit})
	fake.storeMutex.Unlock()
	if fake.StoreStub != nil {
		fake.StoreStub(handle, exit)
	}
}

func (fake *FakeExitStore) StoreCallCount() int {
	fake.storeMutex.RLock()
	defer fake.storeMutex.RUnlock()
	return len(fake.storeArgsForCall)
}

func (fake *FakeExitStore) StoreArgsForCall(i int) (string, <-chan struct{}) {
	fake.storeMutex.RLock()
	defer fake.storeMutex.RUnlock()
	return fake.storeArgsForCall[i].handle, fake.storeArgsForCall[i].exit
}

func (fake *FakeExitStore) Unstore(handle string) {
	fake.unstoreMutex.Lock()
	fake.unstoreArgsForCall = append(fake.unstoreArgsForCall, struct {
		handle string
	}{handle})
	fake.unstoreMutex.Unlock()
	if fake.UnstoreStub != nil {
		fake.UnstoreStub(handle)
	}
}

func (fake *FakeExitStore) UnstoreCallCount() int {
	fake.unstoreMutex.RLock()
	defer fake.unstoreMutex.RUnlock()
	return len(fake.unstoreArgsForCall)
}

func (fake *FakeExitStore) UnstoreArgsForCall(i int) string {
	fake.unstoreMutex.RLock()
	defer fake.unstoreMutex.RUnlock()
	return fake.unstoreArgsForCall[i].handle
}

func (fake *FakeExitStore) Wait(handle string) {
	fake.waitMutex.Lock()
	fake.waitArgsForCall = append(fake.waitArgsForCall, struct {
		handle string
	}{handle})
	fake.waitMutex.Unlock()
	if fake.WaitStub != nil {
		fake.WaitStub(handle)
	}
}

func (fake *FakeExitStore) WaitCallCount() int {
	fake.waitMutex.RLock()
	defer fake.waitMutex.RUnlock()
	return len(fake.waitArgsForCall)
}

func (fake *FakeExitStore) WaitArgsForCall(i int) string {
	fake.waitMutex.RLock()
	defer fake.waitMutex.RUnlock()
	return fake.waitArgsForCall[i].handle
}

var _ rundmc.ExitStore = new(FakeExitStore)