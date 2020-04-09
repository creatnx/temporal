// The MIT License
//
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.
//
// Copyright (c) 2020 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

// Code generated by MockGen. DO NOT EDIT.
// Source: timerQueueProcessor.go

// Package history is a generated GoMock package.
package history

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	persistence "github.com/temporalio/temporal/common/persistence"
)

// MocktimerQueueProcessor is a mock of timerQueueProcessor interface.
type MocktimerQueueProcessor struct {
	ctrl     *gomock.Controller
	recorder *MocktimerQueueProcessorMockRecorder
}

// MocktimerQueueProcessorMockRecorder is the mock recorder for MocktimerQueueProcessor.
type MocktimerQueueProcessorMockRecorder struct {
	mock *MocktimerQueueProcessor
}

// NewMocktimerQueueProcessor creates a new mock instance.
func NewMocktimerQueueProcessor(ctrl *gomock.Controller) *MocktimerQueueProcessor {
	mock := &MocktimerQueueProcessor{ctrl: ctrl}
	mock.recorder = &MocktimerQueueProcessorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MocktimerQueueProcessor) EXPECT() *MocktimerQueueProcessorMockRecorder {
	return m.recorder
}

// Start mocks base method.
func (m *MocktimerQueueProcessor) Start() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Start")
}

// Start indicates an expected call of Start.
func (mr *MocktimerQueueProcessorMockRecorder) Start() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MocktimerQueueProcessor)(nil).Start))
}

// Stop mocks base method.
func (m *MocktimerQueueProcessor) Stop() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Stop")
}

// Stop indicates an expected call of Stop.
func (mr *MocktimerQueueProcessorMockRecorder) Stop() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*MocktimerQueueProcessor)(nil).Stop))
}

// FailoverNamespace mocks base method.
func (m *MocktimerQueueProcessor) FailoverNamespace(namespaceIDs map[string]struct{}) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "FailoverNamespace", namespaceIDs)
}

// FailoverNamespace indicates an expected call of FailoverNamespace.
func (mr *MocktimerQueueProcessorMockRecorder) FailoverNamespace(namespaceIDs interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FailoverNamespace", reflect.TypeOf((*MocktimerQueueProcessor)(nil).FailoverNamespace), namespaceIDs)
}

// NotifyNewTimers mocks base method.
func (m *MocktimerQueueProcessor) NotifyNewTimers(clusterName string, timerTask []persistence.Task) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "NotifyNewTimers", clusterName, timerTask)
}

// NotifyNewTimers indicates an expected call of NotifyNewTimers.
func (mr *MocktimerQueueProcessorMockRecorder) NotifyNewTimers(clusterName, timerTask interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NotifyNewTimers", reflect.TypeOf((*MocktimerQueueProcessor)(nil).NotifyNewTimers), clusterName, timerTask)
}

// LockTaskProcessing mocks base method.
func (m *MocktimerQueueProcessor) LockTaskProcessing() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "LockTaskProcessing")
}

// LockTaskProcessing indicates an expected call of LockTaskProcessing.
func (mr *MocktimerQueueProcessorMockRecorder) LockTaskProcessing() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LockTaskProcessing", reflect.TypeOf((*MocktimerQueueProcessor)(nil).LockTaskProcessing))
}

// UnlockTaskProcessing mocks base method.
func (m *MocktimerQueueProcessor) UnlockTaskProcessing() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "UnlockTaskProcessing")
}

// UnlockTaskProcessing indicates an expected call of UnlockTaskProcessing.
func (mr *MocktimerQueueProcessorMockRecorder) UnlockTaskProcessing() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnlockTaskProcessing", reflect.TypeOf((*MocktimerQueueProcessor)(nil).UnlockTaskProcessing))
}
