// Code generated by mockery v2.28.1. DO NOT EDIT.

package mocks

import (
	context "context"

	controllerv1alpha1 "github.com/kubeslice/kubeslice-controller/apis/controller/v1alpha1"
	mock "github.com/stretchr/testify/mock"

	reconcile "sigs.k8s.io/controller-runtime/pkg/reconcile"

	v1alpha1 "github.com/kubeslice/kubeslice-controller/apis/worker/v1alpha1"
)

// IWorkerSliceGatewayService is an autogenerated mock type for the IWorkerSliceGatewayService type
type IWorkerSliceGatewayService struct {
	mock.Mock
}

// CreateMinimumWorkerSliceGateways provides a mock function with given fields: ctx, sliceName, clusterNames, namespace, label, clusterMap, sliceSubnet, clusterCidr
func (_m *IWorkerSliceGatewayService) CreateMinimumWorkerSliceGateways(ctx context.Context, sliceName string, clusterNames []string, namespace string, label map[string]string, clusterMap map[string]int, sliceSubnet string, clusterCidr string) (reconcile.Result, error) {
	ret := _m.Called(ctx, sliceName, clusterNames, namespace, label, clusterMap, sliceSubnet, clusterCidr)

	var r0 reconcile.Result
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, []string, string, map[string]string, map[string]int, string, string) (reconcile.Result, error)); ok {
		return rf(ctx, sliceName, clusterNames, namespace, label, clusterMap, sliceSubnet, clusterCidr)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, []string, string, map[string]string, map[string]int, string, string) reconcile.Result); ok {
		r0 = rf(ctx, sliceName, clusterNames, namespace, label, clusterMap, sliceSubnet, clusterCidr)
	} else {
		r0 = ret.Get(0).(reconcile.Result)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, []string, string, map[string]string, map[string]int, string, string) error); ok {
		r1 = rf(ctx, sliceName, clusterNames, namespace, label, clusterMap, sliceSubnet, clusterCidr)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteWorkerSliceGatewaysByLabel provides a mock function with given fields: ctx, label, namespace
func (_m *IWorkerSliceGatewayService) DeleteWorkerSliceGatewaysByLabel(ctx context.Context, label map[string]string, namespace string) error {
	ret := _m.Called(ctx, label, namespace)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, map[string]string, string) error); ok {
		r0 = rf(ctx, label, namespace)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ListWorkerSliceGateways provides a mock function with given fields: ctx, ownerLabel, namespace
func (_m *IWorkerSliceGatewayService) ListWorkerSliceGateways(ctx context.Context, ownerLabel map[string]string, namespace string) ([]v1alpha1.WorkerSliceGateway, error) {
	ret := _m.Called(ctx, ownerLabel, namespace)

	var r0 []v1alpha1.WorkerSliceGateway
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, map[string]string, string) ([]v1alpha1.WorkerSliceGateway, error)); ok {
		return rf(ctx, ownerLabel, namespace)
	}
	if rf, ok := ret.Get(0).(func(context.Context, map[string]string, string) []v1alpha1.WorkerSliceGateway); ok {
		r0 = rf(ctx, ownerLabel, namespace)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]v1alpha1.WorkerSliceGateway)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, map[string]string, string) error); ok {
		r1 = rf(ctx, ownerLabel, namespace)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NodeIpReconciliationOfWorkerSliceGateways provides a mock function with given fields: ctx, cluster, namespace
func (_m *IWorkerSliceGatewayService) NodeIpReconciliationOfWorkerSliceGateways(ctx context.Context, cluster *controllerv1alpha1.Cluster, namespace string) error {
	ret := _m.Called(ctx, cluster, namespace)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *controllerv1alpha1.Cluster, string) error); ok {
		r0 = rf(ctx, cluster, namespace)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ReconcileWorkerSliceGateways provides a mock function with given fields: ctx, req
func (_m *IWorkerSliceGatewayService) ReconcileWorkerSliceGateways(ctx context.Context, req reconcile.Request) (reconcile.Result, error) {
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

type mockConstructorTestingTNewIWorkerSliceGatewayService interface {
	mock.TestingT
	Cleanup(func())
}

// NewIWorkerSliceGatewayService creates a new instance of IWorkerSliceGatewayService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewIWorkerSliceGatewayService(t mockConstructorTestingTNewIWorkerSliceGatewayService) *IWorkerSliceGatewayService {
	mock := &IWorkerSliceGatewayService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
