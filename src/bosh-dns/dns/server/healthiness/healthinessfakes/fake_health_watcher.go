// Code generated by counterfeiter. DO NOT EDIT.
package healthinessfakes

import (
	"bosh-dns/dns/server/healthiness"
	"sync"
)

type FakeHealthWatcher struct {
	IsHealthyStub        func(ip string) bool
	isHealthyMutex       sync.RWMutex
	isHealthyArgsForCall []struct {
		ip string
	}
	isHealthyReturns struct {
		result1 bool
	}
	isHealthyReturnsOnCall map[int]struct {
		result1 bool
	}
	UntrackStub        func(ip string)
	untrackMutex       sync.RWMutex
	untrackArgsForCall []struct {
		ip string
	}
	RunStub        func(signal <-chan struct{})
	runMutex       sync.RWMutex
	runArgsForCall []struct {
		signal <-chan struct{}
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeHealthWatcher) IsHealthy(ip string) bool {
	fake.isHealthyMutex.Lock()
	ret, specificReturn := fake.isHealthyReturnsOnCall[len(fake.isHealthyArgsForCall)]
	fake.isHealthyArgsForCall = append(fake.isHealthyArgsForCall, struct {
		ip string
	}{ip})
	fake.recordInvocation("IsHealthy", []interface{}{ip})
	fake.isHealthyMutex.Unlock()
	if fake.IsHealthyStub != nil {
		return fake.IsHealthyStub(ip)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.isHealthyReturns.result1
}

func (fake *FakeHealthWatcher) IsHealthyCallCount() int {
	fake.isHealthyMutex.RLock()
	defer fake.isHealthyMutex.RUnlock()
	return len(fake.isHealthyArgsForCall)
}

func (fake *FakeHealthWatcher) IsHealthyArgsForCall(i int) string {
	fake.isHealthyMutex.RLock()
	defer fake.isHealthyMutex.RUnlock()
	return fake.isHealthyArgsForCall[i].ip
}

func (fake *FakeHealthWatcher) IsHealthyReturns(result1 bool) {
	fake.IsHealthyStub = nil
	fake.isHealthyReturns = struct {
		result1 bool
	}{result1}
}

func (fake *FakeHealthWatcher) IsHealthyReturnsOnCall(i int, result1 bool) {
	fake.IsHealthyStub = nil
	if fake.isHealthyReturnsOnCall == nil {
		fake.isHealthyReturnsOnCall = make(map[int]struct {
			result1 bool
		})
	}
	fake.isHealthyReturnsOnCall[i] = struct {
		result1 bool
	}{result1}
}

func (fake *FakeHealthWatcher) Untrack(ip string) {
	fake.untrackMutex.Lock()
	fake.untrackArgsForCall = append(fake.untrackArgsForCall, struct {
		ip string
	}{ip})
	fake.recordInvocation("Untrack", []interface{}{ip})
	fake.untrackMutex.Unlock()
	if fake.UntrackStub != nil {
		fake.UntrackStub(ip)
	}
}

func (fake *FakeHealthWatcher) UntrackCallCount() int {
	fake.untrackMutex.RLock()
	defer fake.untrackMutex.RUnlock()
	return len(fake.untrackArgsForCall)
}

func (fake *FakeHealthWatcher) UntrackArgsForCall(i int) string {
	fake.untrackMutex.RLock()
	defer fake.untrackMutex.RUnlock()
	return fake.untrackArgsForCall[i].ip
}

func (fake *FakeHealthWatcher) Run(signal <-chan struct{}) {
	fake.runMutex.Lock()
	fake.runArgsForCall = append(fake.runArgsForCall, struct {
		signal <-chan struct{}
	}{signal})
	fake.recordInvocation("Run", []interface{}{signal})
	fake.runMutex.Unlock()
	if fake.RunStub != nil {
		fake.RunStub(signal)
	}
}

func (fake *FakeHealthWatcher) RunCallCount() int {
	fake.runMutex.RLock()
	defer fake.runMutex.RUnlock()
	return len(fake.runArgsForCall)
}

func (fake *FakeHealthWatcher) RunArgsForCall(i int) <-chan struct{} {
	fake.runMutex.RLock()
	defer fake.runMutex.RUnlock()
	return fake.runArgsForCall[i].signal
}

func (fake *FakeHealthWatcher) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.isHealthyMutex.RLock()
	defer fake.isHealthyMutex.RUnlock()
	fake.untrackMutex.RLock()
	defer fake.untrackMutex.RUnlock()
	fake.runMutex.RLock()
	defer fake.runMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeHealthWatcher) recordInvocation(key string, args []interface{}) {
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

var _ healthiness.HealthWatcher = new(FakeHealthWatcher)
