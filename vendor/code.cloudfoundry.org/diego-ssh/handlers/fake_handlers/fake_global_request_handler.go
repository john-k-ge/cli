// This file was generated by counterfeiter
package fake_handlers

import (
	"sync"

	"code.cloudfoundry.org/diego-ssh/handlers"
	"code.cloudfoundry.org/lager"
	"golang.org/x/crypto/ssh"
)

type FakeGlobalRequestHandler struct {
	HandleRequestStub        func(logger lager.Logger, request *ssh.Request)
	handleRequestMutex       sync.RWMutex
	handleRequestArgsForCall []struct {
		logger  lager.Logger
		request *ssh.Request
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeGlobalRequestHandler) HandleRequest(logger lager.Logger, request *ssh.Request) {
	fake.handleRequestMutex.Lock()
	fake.handleRequestArgsForCall = append(fake.handleRequestArgsForCall, struct {
		logger  lager.Logger
		request *ssh.Request
	}{logger, request})
	fake.recordInvocation("HandleRequest", []interface{}{logger, request})
	fake.handleRequestMutex.Unlock()
	if fake.HandleRequestStub != nil {
		fake.HandleRequestStub(logger, request)
	}
}

func (fake *FakeGlobalRequestHandler) HandleRequestCallCount() int {
	fake.handleRequestMutex.RLock()
	defer fake.handleRequestMutex.RUnlock()
	return len(fake.handleRequestArgsForCall)
}

func (fake *FakeGlobalRequestHandler) HandleRequestArgsForCall(i int) (lager.Logger, *ssh.Request) {
	fake.handleRequestMutex.RLock()
	defer fake.handleRequestMutex.RUnlock()
	return fake.handleRequestArgsForCall[i].logger, fake.handleRequestArgsForCall[i].request
}

func (fake *FakeGlobalRequestHandler) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.handleRequestMutex.RLock()
	defer fake.handleRequestMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeGlobalRequestHandler) recordInvocation(key string, args []interface{}) {
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

var _ handlers.GlobalRequestHandler = new(FakeGlobalRequestHandler)
