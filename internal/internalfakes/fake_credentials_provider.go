// Code generated by counterfeiter. DO NOT EDIT.
package internalfakes

import (
	"sync"

	"github.com/rabbitmq/messaging-topology-operator/internal"
)

type FakeCredentialsProvider struct {
	GetPasswordStub        func() string
	getPasswordMutex       sync.RWMutex
	getPasswordArgsForCall []struct {
	}
	getPasswordReturns struct {
		result1 string
	}
	getPasswordReturnsOnCall map[int]struct {
		result1 string
	}
	GetUserStub        func() string
	getUserMutex       sync.RWMutex
	getUserArgsForCall []struct {
	}
	getUserReturns struct {
		result1 string
	}
	getUserReturnsOnCall map[int]struct {
		result1 string
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeCredentialsProvider) GetPassword() string {
	fake.getPasswordMutex.Lock()
	ret, specificReturn := fake.getPasswordReturnsOnCall[len(fake.getPasswordArgsForCall)]
	fake.getPasswordArgsForCall = append(fake.getPasswordArgsForCall, struct {
	}{})
	stub := fake.GetPasswordStub
	fakeReturns := fake.getPasswordReturns
	fake.recordInvocation("GetPassword", []interface{}{})
	fake.getPasswordMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeCredentialsProvider) GetPasswordCallCount() int {
	fake.getPasswordMutex.RLock()
	defer fake.getPasswordMutex.RUnlock()
	return len(fake.getPasswordArgsForCall)
}

func (fake *FakeCredentialsProvider) GetPasswordCalls(stub func() string) {
	fake.getPasswordMutex.Lock()
	defer fake.getPasswordMutex.Unlock()
	fake.GetPasswordStub = stub
}

func (fake *FakeCredentialsProvider) GetPasswordReturns(result1 string) {
	fake.getPasswordMutex.Lock()
	defer fake.getPasswordMutex.Unlock()
	fake.GetPasswordStub = nil
	fake.getPasswordReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeCredentialsProvider) GetPasswordReturnsOnCall(i int, result1 string) {
	fake.getPasswordMutex.Lock()
	defer fake.getPasswordMutex.Unlock()
	fake.GetPasswordStub = nil
	if fake.getPasswordReturnsOnCall == nil {
		fake.getPasswordReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.getPasswordReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeCredentialsProvider) GetUser() string {
	fake.getUserMutex.Lock()
	ret, specificReturn := fake.getUserReturnsOnCall[len(fake.getUserArgsForCall)]
	fake.getUserArgsForCall = append(fake.getUserArgsForCall, struct {
	}{})
	stub := fake.GetUserStub
	fakeReturns := fake.getUserReturns
	fake.recordInvocation("GetUser", []interface{}{})
	fake.getUserMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeCredentialsProvider) GetUserCallCount() int {
	fake.getUserMutex.RLock()
	defer fake.getUserMutex.RUnlock()
	return len(fake.getUserArgsForCall)
}

func (fake *FakeCredentialsProvider) GetUserCalls(stub func() string) {
	fake.getUserMutex.Lock()
	defer fake.getUserMutex.Unlock()
	fake.GetUserStub = stub
}

func (fake *FakeCredentialsProvider) GetUserReturns(result1 string) {
	fake.getUserMutex.Lock()
	defer fake.getUserMutex.Unlock()
	fake.GetUserStub = nil
	fake.getUserReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeCredentialsProvider) GetUserReturnsOnCall(i int, result1 string) {
	fake.getUserMutex.Lock()
	defer fake.getUserMutex.Unlock()
	fake.GetUserStub = nil
	if fake.getUserReturnsOnCall == nil {
		fake.getUserReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.getUserReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeCredentialsProvider) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getPasswordMutex.RLock()
	defer fake.getPasswordMutex.RUnlock()
	fake.getUserMutex.RLock()
	defer fake.getUserMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeCredentialsProvider) recordInvocation(key string, args []interface{}) {
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

var _ internal.CredentialsProvider = new(FakeCredentialsProvider)