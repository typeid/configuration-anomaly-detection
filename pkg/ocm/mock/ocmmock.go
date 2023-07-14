// Code generated by MockGen. DO NOT EDIT.
// Source: ocm.go

// Package ocmmock is a generated GoMock package.
package ocmmock

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1 "github.com/openshift-online/ocm-sdk-go/clustersmgmt/v1"
	ocm "github.com/openshift/configuration-anomaly-detection/pkg/ocm"
)

// MockClient is a mock of Client interface.
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *MockClientMockRecorder
}

// MockClientMockRecorder is the mock recorder for MockClient.
type MockClientMockRecorder struct {
	mock *MockClient
}

// NewMockClient creates a new mock instance.
func NewMockClient(ctrl *gomock.Controller) *MockClient {
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &MockClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClient) EXPECT() *MockClientMockRecorder {
	return m.recorder
}

// DeleteLimitedSupportReasons mocks base method.
func (m *MockClient) DeleteLimitedSupportReasons(ls ocm.LimitedSupportReason, internalClusterID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteLimitedSupportReasons", ls, internalClusterID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteLimitedSupportReasons indicates an expected call of DeleteLimitedSupportReasons.
func (mr *MockClientMockRecorder) DeleteLimitedSupportReasons(ls, internalClusterID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteLimitedSupportReasons", reflect.TypeOf((*MockClient)(nil).DeleteLimitedSupportReasons), ls, internalClusterID)
}

// GetClusterMachinePools mocks base method.
func (m *MockClient) GetClusterMachinePools(internalClusterID string) ([]*v1.MachinePool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetClusterMachinePools", internalClusterID)
	ret0, _ := ret[0].([]*v1.MachinePool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetClusterMachinePools indicates an expected call of GetClusterMachinePools.
func (mr *MockClientMockRecorder) GetClusterMachinePools(internalClusterID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetClusterMachinePools", reflect.TypeOf((*MockClient)(nil).GetClusterMachinePools), internalClusterID)
}

// GetSupportRoleARN mocks base method.
func (m *MockClient) GetSupportRoleARN(internalClusterID string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSupportRoleARN", internalClusterID)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSupportRoleARN indicates an expected call of GetSupportRoleARN.
func (mr *MockClientMockRecorder) GetSupportRoleARN(internalClusterID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSupportRoleARN", reflect.TypeOf((*MockClient)(nil).GetSupportRoleARN), internalClusterID)
}

// IsInLimitedSupport mocks base method.
func (m *MockClient) IsInLimitedSupport(internalClusterID string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsInLimitedSupport", internalClusterID)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsInLimitedSupport indicates an expected call of IsInLimitedSupport.
func (mr *MockClientMockRecorder) IsInLimitedSupport(internalClusterID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsInLimitedSupport", reflect.TypeOf((*MockClient)(nil).IsInLimitedSupport), internalClusterID)
}

// LimitedSupportReasonExists mocks base method.
func (m *MockClient) LimitedSupportReasonExists(ls ocm.LimitedSupportReason, internalClusterID string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LimitedSupportReasonExists", ls, internalClusterID)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LimitedSupportReasonExists indicates an expected call of LimitedSupportReasonExists.
func (mr *MockClientMockRecorder) LimitedSupportReasonExists(ls, internalClusterID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LimitedSupportReasonExists", reflect.TypeOf((*MockClient)(nil).LimitedSupportReasonExists), ls, internalClusterID)
}

// PostLimitedSupportReason mocks base method.
func (m *MockClient) PostLimitedSupportReason(limitedSupportReason ocm.LimitedSupportReason, internalClusterID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PostLimitedSupportReason", limitedSupportReason, internalClusterID)
	ret0, _ := ret[0].(error)
	return ret0
}

// PostLimitedSupportReason indicates an expected call of PostLimitedSupportReason.
func (mr *MockClientMockRecorder) PostLimitedSupportReason(limitedSupportReason, internalClusterID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PostLimitedSupportReason", reflect.TypeOf((*MockClient)(nil).PostLimitedSupportReason), limitedSupportReason, internalClusterID)
}

// UnrelatedLimitedSupportExists mocks base method.
func (m *MockClient) UnrelatedLimitedSupportExists(ls ocm.LimitedSupportReason, internalClusterID string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UnrelatedLimitedSupportExists", ls, internalClusterID)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UnrelatedLimitedSupportExists indicates an expected call of UnrelatedLimitedSupportExists.
func (mr *MockClientMockRecorder) UnrelatedLimitedSupportExists(ls, internalClusterID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnrelatedLimitedSupportExists", reflect.TypeOf((*MockClient)(nil).UnrelatedLimitedSupportExists), ls, internalClusterID)
}
