// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/erda-project/erda/internal/tools/orchestrator/scheduler/impl/servicegroup (interfaces: ServiceGroup)

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"

	apistructs "github.com/erda-project/erda/apistructs"
)

// MockServiceGroup is a mock of ServiceGroup interface.
type MockServiceGroup struct {
	ctrl     *gomock.Controller
	recorder *MockServiceGroupMockRecorder
}

// MockServiceGroupMockRecorder is the mock recorder for MockServiceGroup.
type MockServiceGroupMockRecorder struct {
	mock *MockServiceGroup
}

// NewMockServiceGroup creates a new mock instance.
func NewMockServiceGroup(ctrl *gomock.Controller) *MockServiceGroup {
	mock := &MockServiceGroup{ctrl: ctrl}
	mock.recorder = &MockServiceGroupMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockServiceGroup) EXPECT() *MockServiceGroupMockRecorder {
	return m.recorder
}

// Cancel mocks base method.
func (m *MockServiceGroup) Cancel(arg0, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Cancel", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Cancel indicates an expected call of Cancel.
func (mr *MockServiceGroupMockRecorder) Cancel(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Cancel", reflect.TypeOf((*MockServiceGroup)(nil).Cancel), arg0, arg1)
}

// ConfigUpdate mocks base method.
func (m *MockServiceGroup) ConfigUpdate(arg0 apistructs.ServiceGroup) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConfigUpdate", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// ConfigUpdate indicates an expected call of ConfigUpdate.
func (mr *MockServiceGroupMockRecorder) ConfigUpdate(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConfigUpdate", reflect.TypeOf((*MockServiceGroup)(nil).ConfigUpdate), arg0)
}

// Create mocks base method.
func (m *MockServiceGroup) Create(arg0 apistructs.ServiceGroupCreateV2Request) (apistructs.ServiceGroup, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0)
	ret0, _ := ret[0].(apistructs.ServiceGroup)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockServiceGroupMockRecorder) Create(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockServiceGroup)(nil).Create), arg0)
}

// Delete mocks base method.
func (m *MockServiceGroup) Delete(arg0, arg1, arg2 string, arg3 map[string]string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockServiceGroupMockRecorder) Delete(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockServiceGroup)(nil).Delete), arg0, arg1, arg2, arg3)
}

// Info mocks base method.
func (m *MockServiceGroup) Info(arg0 context.Context, arg1, arg2 string) (apistructs.ServiceGroup, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Info", arg0, arg1, arg2)
	ret0, _ := ret[0].(apistructs.ServiceGroup)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Info indicates an expected call of Info.
func (mr *MockServiceGroupMockRecorder) Info(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Info", reflect.TypeOf((*MockServiceGroup)(nil).Info), arg0, arg1, arg2)
}

// InspectServiceGroupWithTimeout mocks base method.
func (m *MockServiceGroup) InspectServiceGroupWithTimeout(arg0, arg1 string) (*apistructs.ServiceGroup, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InspectServiceGroupWithTimeout", arg0, arg1)
	ret0, _ := ret[0].(*apistructs.ServiceGroup)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InspectServiceGroupWithTimeout indicates an expected call of InspectServiceGroupWithTimeout.
func (mr *MockServiceGroupMockRecorder) InspectServiceGroupWithTimeout(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InspectServiceGroupWithTimeout", reflect.TypeOf((*MockServiceGroup)(nil).InspectServiceGroupWithTimeout), arg0, arg1)
}

// KillPod mocks base method.
func (m *MockServiceGroup) KillPod(arg0 context.Context, arg1, arg2, arg3 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "KillPod", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// KillPod indicates an expected call of KillPod.
func (mr *MockServiceGroupMockRecorder) KillPod(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "KillPod", reflect.TypeOf((*MockServiceGroup)(nil).KillPod), arg0, arg1, arg2, arg3)
}

// Precheck mocks base method.
func (m *MockServiceGroup) Precheck(arg0 apistructs.ServiceGroupPrecheckRequest) (apistructs.ServiceGroupPrecheckData, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Precheck", arg0)
	ret0, _ := ret[0].(apistructs.ServiceGroupPrecheckData)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Precheck indicates an expected call of Precheck.
func (mr *MockServiceGroupMockRecorder) Precheck(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Precheck", reflect.TypeOf((*MockServiceGroup)(nil).Precheck), arg0)
}

// Restart mocks base method.
func (m *MockServiceGroup) Restart(arg0, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Restart", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Restart indicates an expected call of Restart.
func (mr *MockServiceGroupMockRecorder) Restart(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Restart", reflect.TypeOf((*MockServiceGroup)(nil).Restart), arg0, arg1)
}

// Scale mocks base method.
func (m *MockServiceGroup) Scale(arg0 *apistructs.ServiceGroup) (interface{}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Scale", arg0)
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Scale indicates an expected call of Scale.
func (mr *MockServiceGroupMockRecorder) Scale(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Scale", reflect.TypeOf((*MockServiceGroup)(nil).Scale), arg0)
}

// Update mocks base method.
func (m *MockServiceGroup) Update(arg0 apistructs.ServiceGroupUpdateV2Request) (apistructs.ServiceGroup, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0)
	ret0, _ := ret[0].(apistructs.ServiceGroup)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockServiceGroupMockRecorder) Update(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockServiceGroup)(nil).Update), arg0)
}
