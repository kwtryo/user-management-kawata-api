// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package handler

import (
	"context"
	"sync"
)

// Ensure, that HealthServiceMock does implement HealthService.
// If this is not the case, regenerate this file with moq.
var _ HealthService = &HealthServiceMock{}

// HealthServiceMock is a mock implementation of HealthService.
//
//	func TestSomethingThatUsesHealthService(t *testing.T) {
//
//		// make and configure a mocked HealthService
//		mockedHealthService := &HealthServiceMock{
//			HealthCheckFunc: func(ctx context.Context) error {
//				panic("mock out the HealthCheck method")
//			},
//		}
//
//		// use mockedHealthService in code that requires HealthService
//		// and then make assertions.
//
//	}
type HealthServiceMock struct {
	// HealthCheckFunc mocks the HealthCheck method.
	HealthCheckFunc func(ctx context.Context) error

	// calls tracks calls to the methods.
	calls struct {
		// HealthCheck holds details about calls to the HealthCheck method.
		HealthCheck []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
		}
	}
	lockHealthCheck sync.RWMutex
}

// HealthCheck calls HealthCheckFunc.
func (mock *HealthServiceMock) HealthCheck(ctx context.Context) error {
	if mock.HealthCheckFunc == nil {
		panic("HealthServiceMock.HealthCheckFunc: method is nil but HealthService.HealthCheck was just called")
	}
	callInfo := struct {
		Ctx context.Context
	}{
		Ctx: ctx,
	}
	mock.lockHealthCheck.Lock()
	mock.calls.HealthCheck = append(mock.calls.HealthCheck, callInfo)
	mock.lockHealthCheck.Unlock()
	return mock.HealthCheckFunc(ctx)
}

// HealthCheckCalls gets all the calls that were made to HealthCheck.
// Check the length with:
//
//	len(mockedHealthService.HealthCheckCalls())
func (mock *HealthServiceMock) HealthCheckCalls() []struct {
	Ctx context.Context
} {
	var calls []struct {
		Ctx context.Context
	}
	mock.lockHealthCheck.RLock()
	calls = mock.calls.HealthCheck
	mock.lockHealthCheck.RUnlock()
	return calls
}