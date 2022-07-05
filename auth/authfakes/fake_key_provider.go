// Code generated by counterfeiter. DO NOT EDIT.
package authfakes

import (
	"sync"

	"github.com/tomxiong/protocol/auth"
)

type FakeKeyProvider struct {
	GetSecretStub        func(string) string
	getSecretMutex       sync.RWMutex
	getSecretArgsForCall []struct {
		arg1 string
	}
	getSecretReturns struct {
		result1 string
	}
	getSecretReturnsOnCall map[int]struct {
		result1 string
	}
	NumKeysStub        func() int
	numKeysMutex       sync.RWMutex
	numKeysArgsForCall []struct {
	}
	numKeysReturns struct {
		result1 int
	}
	numKeysReturnsOnCall map[int]struct {
		result1 int
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeKeyProvider) GetSecret(arg1 string) string {
	fake.getSecretMutex.Lock()
	ret, specificReturn := fake.getSecretReturnsOnCall[len(fake.getSecretArgsForCall)]
	fake.getSecretArgsForCall = append(fake.getSecretArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.GetSecretStub
	fakeReturns := fake.getSecretReturns
	fake.recordInvocation("GetSecret", []interface{}{arg1})
	fake.getSecretMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeKeyProvider) GetSecretCallCount() int {
	fake.getSecretMutex.RLock()
	defer fake.getSecretMutex.RUnlock()
	return len(fake.getSecretArgsForCall)
}

func (fake *FakeKeyProvider) GetSecretCalls(stub func(string) string) {
	fake.getSecretMutex.Lock()
	defer fake.getSecretMutex.Unlock()
	fake.GetSecretStub = stub
}

func (fake *FakeKeyProvider) GetSecretArgsForCall(i int) string {
	fake.getSecretMutex.RLock()
	defer fake.getSecretMutex.RUnlock()
	argsForCall := fake.getSecretArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeKeyProvider) GetSecretReturns(result1 string) {
	fake.getSecretMutex.Lock()
	defer fake.getSecretMutex.Unlock()
	fake.GetSecretStub = nil
	fake.getSecretReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeKeyProvider) GetSecretReturnsOnCall(i int, result1 string) {
	fake.getSecretMutex.Lock()
	defer fake.getSecretMutex.Unlock()
	fake.GetSecretStub = nil
	if fake.getSecretReturnsOnCall == nil {
		fake.getSecretReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.getSecretReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeKeyProvider) NumKeys() int {
	fake.numKeysMutex.Lock()
	ret, specificReturn := fake.numKeysReturnsOnCall[len(fake.numKeysArgsForCall)]
	fake.numKeysArgsForCall = append(fake.numKeysArgsForCall, struct {
	}{})
	stub := fake.NumKeysStub
	fakeReturns := fake.numKeysReturns
	fake.recordInvocation("NumKeys", []interface{}{})
	fake.numKeysMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeKeyProvider) NumKeysCallCount() int {
	fake.numKeysMutex.RLock()
	defer fake.numKeysMutex.RUnlock()
	return len(fake.numKeysArgsForCall)
}

func (fake *FakeKeyProvider) NumKeysCalls(stub func() int) {
	fake.numKeysMutex.Lock()
	defer fake.numKeysMutex.Unlock()
	fake.NumKeysStub = stub
}

func (fake *FakeKeyProvider) NumKeysReturns(result1 int) {
	fake.numKeysMutex.Lock()
	defer fake.numKeysMutex.Unlock()
	fake.NumKeysStub = nil
	fake.numKeysReturns = struct {
		result1 int
	}{result1}
}

func (fake *FakeKeyProvider) NumKeysReturnsOnCall(i int, result1 int) {
	fake.numKeysMutex.Lock()
	defer fake.numKeysMutex.Unlock()
	fake.NumKeysStub = nil
	if fake.numKeysReturnsOnCall == nil {
		fake.numKeysReturnsOnCall = make(map[int]struct {
			result1 int
		})
	}
	fake.numKeysReturnsOnCall[i] = struct {
		result1 int
	}{result1}
}

func (fake *FakeKeyProvider) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getSecretMutex.RLock()
	defer fake.getSecretMutex.RUnlock()
	fake.numKeysMutex.RLock()
	defer fake.numKeysMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeKeyProvider) recordInvocation(key string, args []interface{}) {
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

var _ auth.KeyProvider = new(FakeKeyProvider)
