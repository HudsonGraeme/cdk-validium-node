// Code generated by mockery v2.10.0. DO NOT EDIT.

package tree

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// storeMock is an autogenerated mock type for the Store type
type storeMock struct {
	mock.Mock
}

// BeginDBTransaction provides a mock function with given fields: ctx, txBundleID
func (_m *storeMock) BeginDBTransaction(ctx context.Context, txBundleID string) error {
	ret := _m.Called(ctx, txBundleID)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, txBundleID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Commit provides a mock function with given fields: ctx, txBundleID
func (_m *storeMock) Commit(ctx context.Context, txBundleID string) error {
	ret := _m.Called(ctx, txBundleID)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, txBundleID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Get provides a mock function with given fields: ctx, key, txBundleID
func (_m *storeMock) Get(ctx context.Context, key []byte, txBundleID string) ([]byte, error) {
	ret := _m.Called(ctx, key, txBundleID)

	var r0 []byte
	if rf, ok := ret.Get(0).(func(context.Context, []byte, string) []byte); ok {
		r0 = rf(ctx, key, txBundleID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, []byte, string) error); ok {
		r1 = rf(ctx, key, txBundleID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Rollback provides a mock function with given fields: ctx, txBundleID
func (_m *storeMock) Rollback(ctx context.Context, txBundleID string) error {
	ret := _m.Called(ctx, txBundleID)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, txBundleID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Set provides a mock function with given fields: ctx, key, value, txBundleID
func (_m *storeMock) Set(ctx context.Context, key []byte, value []byte, txBundleID string) error {
	ret := _m.Called(ctx, key, value, txBundleID)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, []byte, []byte, string) error); ok {
		r0 = rf(ctx, key, value, txBundleID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SupportsDBTransactions provides a mock function with given fields:
func (_m *storeMock) SupportsDBTransactions() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}