// This file was generated by counterfeiter
package applicationfakes

import (
	"sync"

	"github.com/cloudfoundry/cli/cf/commands/application"
	"github.com/cloudfoundry/cli/cf/models"
)

type FakeApplicationStagingWatcher struct {
	ApplicationWatchStagingStub        func(app models.Application, orgName string, spaceName string, startCommand func(app models.Application) (models.Application, error)) (updatedApp models.Application, err error)
	applicationWatchStagingMutex       sync.RWMutex
	applicationWatchStagingArgsForCall []struct {
		app          models.Application
		orgName      string
		spaceName    string
		startCommand func(app models.Application) (models.Application, error)
	}
	applicationWatchStagingReturns struct {
		result1 models.Application
		result2 error
	}
}

func (fake *FakeApplicationStagingWatcher) ApplicationWatchStaging(app models.Application, orgName string, spaceName string, startCommand func(app models.Application) (models.Application, error)) (updatedApp models.Application, err error) {
	fake.applicationWatchStagingMutex.Lock()
	fake.applicationWatchStagingArgsForCall = append(fake.applicationWatchStagingArgsForCall, struct {
		app          models.Application
		orgName      string
		spaceName    string
		startCommand func(app models.Application) (models.Application, error)
	}{app, orgName, spaceName, startCommand})
	fake.applicationWatchStagingMutex.Unlock()
	if fake.ApplicationWatchStagingStub != nil {
		return fake.ApplicationWatchStagingStub(app, orgName, spaceName, startCommand)
	} else {
		return fake.applicationWatchStagingReturns.result1, fake.applicationWatchStagingReturns.result2
	}
}

func (fake *FakeApplicationStagingWatcher) ApplicationWatchStagingCallCount() int {
	fake.applicationWatchStagingMutex.RLock()
	defer fake.applicationWatchStagingMutex.RUnlock()
	return len(fake.applicationWatchStagingArgsForCall)
}

func (fake *FakeApplicationStagingWatcher) ApplicationWatchStagingArgsForCall(i int) (models.Application, string, string, func(app models.Application) (models.Application, error)) {
	fake.applicationWatchStagingMutex.RLock()
	defer fake.applicationWatchStagingMutex.RUnlock()
	return fake.applicationWatchStagingArgsForCall[i].app, fake.applicationWatchStagingArgsForCall[i].orgName, fake.applicationWatchStagingArgsForCall[i].spaceName, fake.applicationWatchStagingArgsForCall[i].startCommand
}

func (fake *FakeApplicationStagingWatcher) ApplicationWatchStagingReturns(result1 models.Application, result2 error) {
	fake.ApplicationWatchStagingStub = nil
	fake.applicationWatchStagingReturns = struct {
		result1 models.Application
		result2 error
	}{result1, result2}
}

var _ application.ApplicationStagingWatcher = new(FakeApplicationStagingWatcher)