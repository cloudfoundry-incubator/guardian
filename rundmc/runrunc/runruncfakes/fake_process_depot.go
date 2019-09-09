// Code generated by counterfeiter. DO NOT EDIT.
package runruncfakes

import (
	"sync"
	"time"

	"code.cloudfoundry.org/guardian/rundmc/runrunc"
	"code.cloudfoundry.org/lager"
)

type FakeProcessDepot struct {
	CreatedTimeStub        func(lager.Logger, string) (time.Time, error)
	createdTimeMutex       sync.RWMutex
	createdTimeArgsForCall []struct {
		arg1 lager.Logger
		arg2 string
	}
	createdTimeReturns struct {
		result1 time.Time
		result2 error
	}
	createdTimeReturnsOnCall map[int]struct {
		result1 time.Time
		result2 error
	}
	ListProcessDirsStub        func(lager.Logger, string) ([]string, error)
	listProcessDirsMutex       sync.RWMutex
	listProcessDirsArgsForCall []struct {
		arg1 lager.Logger
		arg2 string
	}
	listProcessDirsReturns struct {
		result1 []string
		result2 error
	}
	listProcessDirsReturnsOnCall map[int]struct {
		result1 []string
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeProcessDepot) CreatedTime(arg1 lager.Logger, arg2 string) (time.Time, error) {
	fake.createdTimeMutex.Lock()
	ret, specificReturn := fake.createdTimeReturnsOnCall[len(fake.createdTimeArgsForCall)]
	fake.createdTimeArgsForCall = append(fake.createdTimeArgsForCall, struct {
		arg1 lager.Logger
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("CreatedTime", []interface{}{arg1, arg2})
	fake.createdTimeMutex.Unlock()
	if fake.CreatedTimeStub != nil {
		return fake.CreatedTimeStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.createdTimeReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeProcessDepot) CreatedTimeCallCount() int {
	fake.createdTimeMutex.RLock()
	defer fake.createdTimeMutex.RUnlock()
	return len(fake.createdTimeArgsForCall)
}

func (fake *FakeProcessDepot) CreatedTimeCalls(stub func(lager.Logger, string) (time.Time, error)) {
	fake.createdTimeMutex.Lock()
	defer fake.createdTimeMutex.Unlock()
	fake.CreatedTimeStub = stub
}

func (fake *FakeProcessDepot) CreatedTimeArgsForCall(i int) (lager.Logger, string) {
	fake.createdTimeMutex.RLock()
	defer fake.createdTimeMutex.RUnlock()
	argsForCall := fake.createdTimeArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeProcessDepot) CreatedTimeReturns(result1 time.Time, result2 error) {
	fake.createdTimeMutex.Lock()
	defer fake.createdTimeMutex.Unlock()
	fake.CreatedTimeStub = nil
	fake.createdTimeReturns = struct {
		result1 time.Time
		result2 error
	}{result1, result2}
}

func (fake *FakeProcessDepot) CreatedTimeReturnsOnCall(i int, result1 time.Time, result2 error) {
	fake.createdTimeMutex.Lock()
	defer fake.createdTimeMutex.Unlock()
	fake.CreatedTimeStub = nil
	if fake.createdTimeReturnsOnCall == nil {
		fake.createdTimeReturnsOnCall = make(map[int]struct {
			result1 time.Time
			result2 error
		})
	}
	fake.createdTimeReturnsOnCall[i] = struct {
		result1 time.Time
		result2 error
	}{result1, result2}
}

func (fake *FakeProcessDepot) ListProcessDirs(arg1 lager.Logger, arg2 string) ([]string, error) {
	fake.listProcessDirsMutex.Lock()
	ret, specificReturn := fake.listProcessDirsReturnsOnCall[len(fake.listProcessDirsArgsForCall)]
	fake.listProcessDirsArgsForCall = append(fake.listProcessDirsArgsForCall, struct {
		arg1 lager.Logger
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("ListProcessDirs", []interface{}{arg1, arg2})
	fake.listProcessDirsMutex.Unlock()
	if fake.ListProcessDirsStub != nil {
		return fake.ListProcessDirsStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.listProcessDirsReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeProcessDepot) ListProcessDirsCallCount() int {
	fake.listProcessDirsMutex.RLock()
	defer fake.listProcessDirsMutex.RUnlock()
	return len(fake.listProcessDirsArgsForCall)
}

func (fake *FakeProcessDepot) ListProcessDirsCalls(stub func(lager.Logger, string) ([]string, error)) {
	fake.listProcessDirsMutex.Lock()
	defer fake.listProcessDirsMutex.Unlock()
	fake.ListProcessDirsStub = stub
}

func (fake *FakeProcessDepot) ListProcessDirsArgsForCall(i int) (lager.Logger, string) {
	fake.listProcessDirsMutex.RLock()
	defer fake.listProcessDirsMutex.RUnlock()
	argsForCall := fake.listProcessDirsArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeProcessDepot) ListProcessDirsReturns(result1 []string, result2 error) {
	fake.listProcessDirsMutex.Lock()
	defer fake.listProcessDirsMutex.Unlock()
	fake.ListProcessDirsStub = nil
	fake.listProcessDirsReturns = struct {
		result1 []string
		result2 error
	}{result1, result2}
}

func (fake *FakeProcessDepot) ListProcessDirsReturnsOnCall(i int, result1 []string, result2 error) {
	fake.listProcessDirsMutex.Lock()
	defer fake.listProcessDirsMutex.Unlock()
	fake.ListProcessDirsStub = nil
	if fake.listProcessDirsReturnsOnCall == nil {
		fake.listProcessDirsReturnsOnCall = make(map[int]struct {
			result1 []string
			result2 error
		})
	}
	fake.listProcessDirsReturnsOnCall[i] = struct {
		result1 []string
		result2 error
	}{result1, result2}
}

func (fake *FakeProcessDepot) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.createdTimeMutex.RLock()
	defer fake.createdTimeMutex.RUnlock()
	fake.listProcessDirsMutex.RLock()
	defer fake.listProcessDirsMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeProcessDepot) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ runrunc.ProcessDepot = new(FakeProcessDepot)