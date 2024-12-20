// Code generated by MockGen. DO NOT EDIT.
// Source: pagerduty.go
//
// Generated by this command:
//
//	mockgen --build_flags=--mod=readonly -source pagerduty.go -destination ./mock/pagerdutymock.go -package pdmock
//

// Package pdmock is a generated GoMock package.
package pdmock

import (
	reflect "reflect"

	pagerduty "github.com/openshift/configuration-anomaly-detection/pkg/pagerduty"
	gomock "go.uber.org/mock/gomock"
)

// MockClient is a mock of Client interface.
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *MockClientMockRecorder
	isgomock struct{}
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

// AddNote mocks base method.
func (m *MockClient) AddNote(notes string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddNote", notes)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddNote indicates an expected call of AddNote.
func (mr *MockClientMockRecorder) AddNote(notes any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddNote", reflect.TypeOf((*MockClient)(nil).AddNote), notes)
}

// CreateNewAlert mocks base method.
func (m *MockClient) CreateNewAlert(newAlert pagerduty.NewAlert, serviceID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateNewAlert", newAlert, serviceID)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateNewAlert indicates an expected call of CreateNewAlert.
func (mr *MockClientMockRecorder) CreateNewAlert(newAlert, serviceID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateNewAlert", reflect.TypeOf((*MockClient)(nil).CreateNewAlert), newAlert, serviceID)
}

// EscalateAlert mocks base method.
func (m *MockClient) EscalateAlert() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EscalateAlert")
	ret0, _ := ret[0].(error)
	return ret0
}

// EscalateAlert indicates an expected call of EscalateAlert.
func (mr *MockClientMockRecorder) EscalateAlert() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EscalateAlert", reflect.TypeOf((*MockClient)(nil).EscalateAlert))
}

// EscalateAlertWithNote mocks base method.
func (m *MockClient) EscalateAlertWithNote(notes string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EscalateAlertWithNote", notes)
	ret0, _ := ret[0].(error)
	return ret0
}

// EscalateAlertWithNote indicates an expected call of EscalateAlertWithNote.
func (mr *MockClientMockRecorder) EscalateAlertWithNote(notes any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EscalateAlertWithNote", reflect.TypeOf((*MockClient)(nil).EscalateAlertWithNote), notes)
}

// GetServiceID mocks base method.
func (m *MockClient) GetServiceID() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetServiceID")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetServiceID indicates an expected call of GetServiceID.
func (mr *MockClientMockRecorder) GetServiceID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetServiceID", reflect.TypeOf((*MockClient)(nil).GetServiceID))
}

// SilenceAlert mocks base method.
func (m *MockClient) SilenceAlert() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SilenceAlert")
	ret0, _ := ret[0].(error)
	return ret0
}

// SilenceAlert indicates an expected call of SilenceAlert.
func (mr *MockClientMockRecorder) SilenceAlert() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SilenceAlert", reflect.TypeOf((*MockClient)(nil).SilenceAlert))
}

// SilenceAlertWithNote mocks base method.
func (m *MockClient) SilenceAlertWithNote(notes string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SilenceAlertWithNote", notes)
	ret0, _ := ret[0].(error)
	return ret0
}

// SilenceAlertWithNote indicates an expected call of SilenceAlertWithNote.
func (mr *MockClientMockRecorder) SilenceAlertWithNote(notes any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SilenceAlertWithNote", reflect.TypeOf((*MockClient)(nil).SilenceAlertWithNote), notes)
}
