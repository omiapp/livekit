// Code generated by counterfeiter. DO NOT EDIT.
package servicefakes

import (
	"context"
	"sync"

	version "github.com/hashicorp/go-version"
	"github.com/livekit/livekit-server/pkg/service"
	"github.com/livekit/protocol/livekit"
)

type FakeEgressStore struct {
	GetEgressVersionStub        func(context.Context) (*version.Version, error)
	getEgressVersionMutex       sync.RWMutex
	getEgressVersionArgsForCall []struct {
		arg1 context.Context
	}
	getEgressVersionReturns struct {
		result1 *version.Version
		result2 error
	}
	getEgressVersionReturnsOnCall map[int]struct {
		result1 *version.Version
		result2 error
	}
	ListEgressStub        func(context.Context, livekit.RoomName) ([]*livekit.EgressInfo, error)
	listEgressMutex       sync.RWMutex
	listEgressArgsForCall []struct {
		arg1 context.Context
		arg2 livekit.RoomName
	}
	listEgressReturns struct {
		result1 []*livekit.EgressInfo
		result2 error
	}
	listEgressReturnsOnCall map[int]struct {
		result1 []*livekit.EgressInfo
		result2 error
	}
	LoadEgressStub        func(context.Context, string) (*livekit.EgressInfo, error)
	loadEgressMutex       sync.RWMutex
	loadEgressArgsForCall []struct {
		arg1 context.Context
		arg2 string
	}
	loadEgressReturns struct {
		result1 *livekit.EgressInfo
		result2 error
	}
	loadEgressReturnsOnCall map[int]struct {
		result1 *livekit.EgressInfo
		result2 error
	}
	StoreEgressStub        func(context.Context, *livekit.EgressInfo) error
	storeEgressMutex       sync.RWMutex
	storeEgressArgsForCall []struct {
		arg1 context.Context
		arg2 *livekit.EgressInfo
	}
	storeEgressReturns struct {
		result1 error
	}
	storeEgressReturnsOnCall map[int]struct {
		result1 error
	}
	UpdateEgressStub        func(context.Context, *livekit.EgressInfo) error
	updateEgressMutex       sync.RWMutex
	updateEgressArgsForCall []struct {
		arg1 context.Context
		arg2 *livekit.EgressInfo
	}
	updateEgressReturns struct {
		result1 error
	}
	updateEgressReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeEgressStore) GetEgressVersion(arg1 context.Context) (*version.Version, error) {
	fake.getEgressVersionMutex.Lock()
	ret, specificReturn := fake.getEgressVersionReturnsOnCall[len(fake.getEgressVersionArgsForCall)]
	fake.getEgressVersionArgsForCall = append(fake.getEgressVersionArgsForCall, struct {
		arg1 context.Context
	}{arg1})
	stub := fake.GetEgressVersionStub
	fakeReturns := fake.getEgressVersionReturns
	fake.recordInvocation("GetEgressVersion", []interface{}{arg1})
	fake.getEgressVersionMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeEgressStore) GetEgressVersionCallCount() int {
	fake.getEgressVersionMutex.RLock()
	defer fake.getEgressVersionMutex.RUnlock()
	return len(fake.getEgressVersionArgsForCall)
}

func (fake *FakeEgressStore) GetEgressVersionCalls(stub func(context.Context) (*version.Version, error)) {
	fake.getEgressVersionMutex.Lock()
	defer fake.getEgressVersionMutex.Unlock()
	fake.GetEgressVersionStub = stub
}

func (fake *FakeEgressStore) GetEgressVersionArgsForCall(i int) context.Context {
	fake.getEgressVersionMutex.RLock()
	defer fake.getEgressVersionMutex.RUnlock()
	argsForCall := fake.getEgressVersionArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeEgressStore) GetEgressVersionReturns(result1 *version.Version, result2 error) {
	fake.getEgressVersionMutex.Lock()
	defer fake.getEgressVersionMutex.Unlock()
	fake.GetEgressVersionStub = nil
	fake.getEgressVersionReturns = struct {
		result1 *version.Version
		result2 error
	}{result1, result2}
}

func (fake *FakeEgressStore) GetEgressVersionReturnsOnCall(i int, result1 *version.Version, result2 error) {
	fake.getEgressVersionMutex.Lock()
	defer fake.getEgressVersionMutex.Unlock()
	fake.GetEgressVersionStub = nil
	if fake.getEgressVersionReturnsOnCall == nil {
		fake.getEgressVersionReturnsOnCall = make(map[int]struct {
			result1 *version.Version
			result2 error
		})
	}
	fake.getEgressVersionReturnsOnCall[i] = struct {
		result1 *version.Version
		result2 error
	}{result1, result2}
}

func (fake *FakeEgressStore) ListEgress(arg1 context.Context, arg2 livekit.RoomName) ([]*livekit.EgressInfo, error) {
	fake.listEgressMutex.Lock()
	ret, specificReturn := fake.listEgressReturnsOnCall[len(fake.listEgressArgsForCall)]
	fake.listEgressArgsForCall = append(fake.listEgressArgsForCall, struct {
		arg1 context.Context
		arg2 livekit.RoomName
	}{arg1, arg2})
	stub := fake.ListEgressStub
	fakeReturns := fake.listEgressReturns
	fake.recordInvocation("ListEgress", []interface{}{arg1, arg2})
	fake.listEgressMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeEgressStore) ListEgressCallCount() int {
	fake.listEgressMutex.RLock()
	defer fake.listEgressMutex.RUnlock()
	return len(fake.listEgressArgsForCall)
}

func (fake *FakeEgressStore) ListEgressCalls(stub func(context.Context, livekit.RoomName) ([]*livekit.EgressInfo, error)) {
	fake.listEgressMutex.Lock()
	defer fake.listEgressMutex.Unlock()
	fake.ListEgressStub = stub
}

func (fake *FakeEgressStore) ListEgressArgsForCall(i int) (context.Context, livekit.RoomName) {
	fake.listEgressMutex.RLock()
	defer fake.listEgressMutex.RUnlock()
	argsForCall := fake.listEgressArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeEgressStore) ListEgressReturns(result1 []*livekit.EgressInfo, result2 error) {
	fake.listEgressMutex.Lock()
	defer fake.listEgressMutex.Unlock()
	fake.ListEgressStub = nil
	fake.listEgressReturns = struct {
		result1 []*livekit.EgressInfo
		result2 error
	}{result1, result2}
}

func (fake *FakeEgressStore) ListEgressReturnsOnCall(i int, result1 []*livekit.EgressInfo, result2 error) {
	fake.listEgressMutex.Lock()
	defer fake.listEgressMutex.Unlock()
	fake.ListEgressStub = nil
	if fake.listEgressReturnsOnCall == nil {
		fake.listEgressReturnsOnCall = make(map[int]struct {
			result1 []*livekit.EgressInfo
			result2 error
		})
	}
	fake.listEgressReturnsOnCall[i] = struct {
		result1 []*livekit.EgressInfo
		result2 error
	}{result1, result2}
}

func (fake *FakeEgressStore) LoadEgress(arg1 context.Context, arg2 string) (*livekit.EgressInfo, error) {
	fake.loadEgressMutex.Lock()
	ret, specificReturn := fake.loadEgressReturnsOnCall[len(fake.loadEgressArgsForCall)]
	fake.loadEgressArgsForCall = append(fake.loadEgressArgsForCall, struct {
		arg1 context.Context
		arg2 string
	}{arg1, arg2})
	stub := fake.LoadEgressStub
	fakeReturns := fake.loadEgressReturns
	fake.recordInvocation("LoadEgress", []interface{}{arg1, arg2})
	fake.loadEgressMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeEgressStore) LoadEgressCallCount() int {
	fake.loadEgressMutex.RLock()
	defer fake.loadEgressMutex.RUnlock()
	return len(fake.loadEgressArgsForCall)
}

func (fake *FakeEgressStore) LoadEgressCalls(stub func(context.Context, string) (*livekit.EgressInfo, error)) {
	fake.loadEgressMutex.Lock()
	defer fake.loadEgressMutex.Unlock()
	fake.LoadEgressStub = stub
}

func (fake *FakeEgressStore) LoadEgressArgsForCall(i int) (context.Context, string) {
	fake.loadEgressMutex.RLock()
	defer fake.loadEgressMutex.RUnlock()
	argsForCall := fake.loadEgressArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeEgressStore) LoadEgressReturns(result1 *livekit.EgressInfo, result2 error) {
	fake.loadEgressMutex.Lock()
	defer fake.loadEgressMutex.Unlock()
	fake.LoadEgressStub = nil
	fake.loadEgressReturns = struct {
		result1 *livekit.EgressInfo
		result2 error
	}{result1, result2}
}

func (fake *FakeEgressStore) LoadEgressReturnsOnCall(i int, result1 *livekit.EgressInfo, result2 error) {
	fake.loadEgressMutex.Lock()
	defer fake.loadEgressMutex.Unlock()
	fake.LoadEgressStub = nil
	if fake.loadEgressReturnsOnCall == nil {
		fake.loadEgressReturnsOnCall = make(map[int]struct {
			result1 *livekit.EgressInfo
			result2 error
		})
	}
	fake.loadEgressReturnsOnCall[i] = struct {
		result1 *livekit.EgressInfo
		result2 error
	}{result1, result2}
}

func (fake *FakeEgressStore) StoreEgress(arg1 context.Context, arg2 *livekit.EgressInfo) error {
	fake.storeEgressMutex.Lock()
	ret, specificReturn := fake.storeEgressReturnsOnCall[len(fake.storeEgressArgsForCall)]
	fake.storeEgressArgsForCall = append(fake.storeEgressArgsForCall, struct {
		arg1 context.Context
		arg2 *livekit.EgressInfo
	}{arg1, arg2})
	stub := fake.StoreEgressStub
	fakeReturns := fake.storeEgressReturns
	fake.recordInvocation("StoreEgress", []interface{}{arg1, arg2})
	fake.storeEgressMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeEgressStore) StoreEgressCallCount() int {
	fake.storeEgressMutex.RLock()
	defer fake.storeEgressMutex.RUnlock()
	return len(fake.storeEgressArgsForCall)
}

func (fake *FakeEgressStore) StoreEgressCalls(stub func(context.Context, *livekit.EgressInfo) error) {
	fake.storeEgressMutex.Lock()
	defer fake.storeEgressMutex.Unlock()
	fake.StoreEgressStub = stub
}

func (fake *FakeEgressStore) StoreEgressArgsForCall(i int) (context.Context, *livekit.EgressInfo) {
	fake.storeEgressMutex.RLock()
	defer fake.storeEgressMutex.RUnlock()
	argsForCall := fake.storeEgressArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeEgressStore) StoreEgressReturns(result1 error) {
	fake.storeEgressMutex.Lock()
	defer fake.storeEgressMutex.Unlock()
	fake.StoreEgressStub = nil
	fake.storeEgressReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeEgressStore) StoreEgressReturnsOnCall(i int, result1 error) {
	fake.storeEgressMutex.Lock()
	defer fake.storeEgressMutex.Unlock()
	fake.StoreEgressStub = nil
	if fake.storeEgressReturnsOnCall == nil {
		fake.storeEgressReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.storeEgressReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeEgressStore) UpdateEgress(arg1 context.Context, arg2 *livekit.EgressInfo) error {
	fake.updateEgressMutex.Lock()
	ret, specificReturn := fake.updateEgressReturnsOnCall[len(fake.updateEgressArgsForCall)]
	fake.updateEgressArgsForCall = append(fake.updateEgressArgsForCall, struct {
		arg1 context.Context
		arg2 *livekit.EgressInfo
	}{arg1, arg2})
	stub := fake.UpdateEgressStub
	fakeReturns := fake.updateEgressReturns
	fake.recordInvocation("UpdateEgress", []interface{}{arg1, arg2})
	fake.updateEgressMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeEgressStore) UpdateEgressCallCount() int {
	fake.updateEgressMutex.RLock()
	defer fake.updateEgressMutex.RUnlock()
	return len(fake.updateEgressArgsForCall)
}

func (fake *FakeEgressStore) UpdateEgressCalls(stub func(context.Context, *livekit.EgressInfo) error) {
	fake.updateEgressMutex.Lock()
	defer fake.updateEgressMutex.Unlock()
	fake.UpdateEgressStub = stub
}

func (fake *FakeEgressStore) UpdateEgressArgsForCall(i int) (context.Context, *livekit.EgressInfo) {
	fake.updateEgressMutex.RLock()
	defer fake.updateEgressMutex.RUnlock()
	argsForCall := fake.updateEgressArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeEgressStore) UpdateEgressReturns(result1 error) {
	fake.updateEgressMutex.Lock()
	defer fake.updateEgressMutex.Unlock()
	fake.UpdateEgressStub = nil
	fake.updateEgressReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeEgressStore) UpdateEgressReturnsOnCall(i int, result1 error) {
	fake.updateEgressMutex.Lock()
	defer fake.updateEgressMutex.Unlock()
	fake.UpdateEgressStub = nil
	if fake.updateEgressReturnsOnCall == nil {
		fake.updateEgressReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.updateEgressReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeEgressStore) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getEgressVersionMutex.RLock()
	defer fake.getEgressVersionMutex.RUnlock()
	fake.listEgressMutex.RLock()
	defer fake.listEgressMutex.RUnlock()
	fake.loadEgressMutex.RLock()
	defer fake.loadEgressMutex.RUnlock()
	fake.storeEgressMutex.RLock()
	defer fake.storeEgressMutex.RUnlock()
	fake.updateEgressMutex.RLock()
	defer fake.updateEgressMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeEgressStore) recordInvocation(key string, args []interface{}) {
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

var _ service.EgressStore = new(FakeEgressStore)
