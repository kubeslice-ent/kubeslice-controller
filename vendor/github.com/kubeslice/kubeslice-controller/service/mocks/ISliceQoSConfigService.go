// Code generated by mockery v2.28.1. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
	reconcile "sigs.k8s.io/controller-runtime/pkg/reconcile"
)

// ISliceQoSConfigService is an autogenerated mock type for the ISliceQoSConfigService type
type ISliceQoSConfigService struct {
	mock.Mock
}

// DeleteSliceQoSConfig provides a mock function with given fields: ctx, namespace
func (_m *ISliceQoSConfigService) DeleteSliceQoSConfig(ctx context.Context, namespace string) (reconcile.Result, error) {
	ret := _m.Called(ctx, namespace)

	var r0 reconcile.Result
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (reconcile.Result, error)); ok {
		return rf(ctx, namespace)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) reconcile.Result); ok {
		r0 = rf(ctx, namespace)
	} else {
		r0 = ret.Get(0).(reconcile.Result)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, namespace)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ReconcileSliceQoSConfig provides a mock function with given fields: ctx, req
func (_m *ISliceQoSConfigService) ReconcileSliceQoSConfig(ctx context.Context, req reconcile.Request) (reconcile.Result, error) {
	ret := _m.Called(ctx, req)

	var r0 reconcile.Result
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, reconcile.Request) (reconcile.Result, error)); ok {
		return rf(ctx, req)
	}
	if rf, ok := ret.Get(0).(func(context.Context, reconcile.Request) reconcile.Result); ok {
		r0 = rf(ctx, req)
	} else {
		r0 = ret.Get(0).(reconcile.Result)
	}

	if rf, ok := ret.Get(1).(func(context.Context, reconcile.Request) error); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewISliceQoSConfigService interface {
	mock.TestingT
	Cleanup(func())
}

// NewISliceQoSConfigService creates a new instance of ISliceQoSConfigService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewISliceQoSConfigService(t mockConstructorTestingTNewISliceQoSConfigService) *ISliceQoSConfigService {
	mock := &ISliceQoSConfigService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}