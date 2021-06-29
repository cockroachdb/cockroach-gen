// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/cockroachdb/cockroach/pkg/kv/kvclient/kvcoord (interfaces: Transport)

// Package kvcoord is a generated GoMock package.
package kvcoord

import (
	context "context"
	reflect "reflect"

	roachpb "github.com/cockroachdb/cockroach/pkg/roachpb"
	gomock "github.com/golang/mock/gomock"
)

// MockTransport is a mock of Transport interface.
type MockTransport struct {
	ctrl     *gomock.Controller
	recorder *MockTransportMockRecorder
}

// MockTransportMockRecorder is the mock recorder for MockTransport.
type MockTransportMockRecorder struct {
	mock *MockTransport
}

// NewMockTransport creates a new mock instance.
func NewMockTransport(ctrl *gomock.Controller) *MockTransport {
	mock := &MockTransport{ctrl: ctrl}
	mock.recorder = &MockTransportMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTransport) EXPECT() *MockTransportMockRecorder {
	return m.recorder
}

// IsExhausted mocks base method.
func (m *MockTransport) IsExhausted() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsExhausted")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsExhausted indicates an expected call of IsExhausted.
func (mr *MockTransportMockRecorder) IsExhausted() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsExhausted", reflect.TypeOf((*MockTransport)(nil).IsExhausted))
}

// MoveToFront mocks base method.
func (m *MockTransport) MoveToFront(arg0 roachpb.ReplicaDescriptor) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "MoveToFront", arg0)
}

// MoveToFront indicates an expected call of MoveToFront.
func (mr *MockTransportMockRecorder) MoveToFront(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MoveToFront", reflect.TypeOf((*MockTransport)(nil).MoveToFront), arg0)
}

// NextInternalClient mocks base method.
func (m *MockTransport) NextInternalClient(arg0 context.Context) (context.Context, roachpb.InternalClient, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NextInternalClient", arg0)
	ret0, _ := ret[0].(context.Context)
	ret1, _ := ret[1].(roachpb.InternalClient)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// NextInternalClient indicates an expected call of NextInternalClient.
func (mr *MockTransportMockRecorder) NextInternalClient(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NextInternalClient", reflect.TypeOf((*MockTransport)(nil).NextInternalClient), arg0)
}

// NextReplica mocks base method.
func (m *MockTransport) NextReplica() roachpb.ReplicaDescriptor {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NextReplica")
	ret0, _ := ret[0].(roachpb.ReplicaDescriptor)
	return ret0
}

// NextReplica indicates an expected call of NextReplica.
func (mr *MockTransportMockRecorder) NextReplica() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NextReplica", reflect.TypeOf((*MockTransport)(nil).NextReplica))
}

// Release mocks base method.
func (m *MockTransport) Release() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Release")
}

// Release indicates an expected call of Release.
func (mr *MockTransportMockRecorder) Release() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Release", reflect.TypeOf((*MockTransport)(nil).Release))
}

// SendNext mocks base method.
func (m *MockTransport) SendNext(arg0 context.Context, arg1 roachpb.BatchRequest) (*roachpb.BatchResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendNext", arg0, arg1)
	ret0, _ := ret[0].(*roachpb.BatchResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SendNext indicates an expected call of SendNext.
func (mr *MockTransportMockRecorder) SendNext(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendNext", reflect.TypeOf((*MockTransport)(nil).SendNext), arg0, arg1)
}

// SkipReplica mocks base method.
func (m *MockTransport) SkipReplica() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SkipReplica")
}

// SkipReplica indicates an expected call of SkipReplica.
func (mr *MockTransportMockRecorder) SkipReplica() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SkipReplica", reflect.TypeOf((*MockTransport)(nil).SkipReplica))
}
