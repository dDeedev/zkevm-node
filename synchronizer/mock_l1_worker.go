// Code generated by mockery v2.22.1. DO NOT EDIT.

package synchronizer

import (
	context "context"
	sync "sync"

	mock "github.com/stretchr/testify/mock"

	time "time"
)

// workerMock is an autogenerated mock type for the worker type
type workerMock struct {
	mock.Mock
}

// String provides a mock function with given fields:
func (_m *workerMock) String() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// asyncRequestRollupInfoByBlockRange provides a mock function with given fields: ctx, ch, wg, blockRange, sleepBefore
func (_m *workerMock) asyncRequestRollupInfoByBlockRange(ctx contextWithCancel, ch chan responseRollupInfoByBlockRange, wg *sync.WaitGroup, blockRange blockRange, sleepBefore time.Duration) error {
	ret := _m.Called(ctx, ch, wg, blockRange, sleepBefore)

	var r0 error
	if rf, ok := ret.Get(0).(func(contextWithCancel, chan responseRollupInfoByBlockRange, *sync.WaitGroup, blockRange, time.Duration) error); ok {
		r0 = rf(ctx, ch, wg, blockRange, sleepBefore)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// isIdle provides a mock function with given fields:
func (_m *workerMock) isIdle() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// requestLastBlock provides a mock function with given fields: ctx
func (_m *workerMock) requestLastBlock(ctx context.Context) responseL1LastBlock {
	ret := _m.Called(ctx)

	var r0 responseL1LastBlock
	if rf, ok := ret.Get(0).(func(context.Context) responseL1LastBlock); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(responseL1LastBlock)
	}

	return r0
}

type mockConstructorTestingTnewWorkerMock interface {
	mock.TestingT
	Cleanup(func())
}

// newWorkerMock creates a new instance of workerMock. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func newWorkerMock(t mockConstructorTestingTnewWorkerMock) *workerMock {
	mock := &workerMock{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
