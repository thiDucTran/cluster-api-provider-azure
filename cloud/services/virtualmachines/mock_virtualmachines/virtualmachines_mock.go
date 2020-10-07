/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by MockGen. DO NOT EDIT.
// Source: ../service.go

// Package mock_virtualmachines is a generated GoMock package.
package mock_virtualmachines

import (
	context "context"
	autorest "github.com/Azure/go-autorest/autorest"
	logr "github.com/go-logr/logr"
	gomock "github.com/golang/mock/gomock"
	v1 "k8s.io/api/core/v1"
	reflect "reflect"
	v1alpha3 "sigs.k8s.io/cluster-api-provider-azure/api/v1alpha3"
	azure "sigs.k8s.io/cluster-api-provider-azure/cloud"
)

// MockVMScope is a mock of VMScope interface.
type MockVMScope struct {
	ctrl     *gomock.Controller
	recorder *MockVMScopeMockRecorder
}

// MockVMScopeMockRecorder is the mock recorder for MockVMScope.
type MockVMScopeMockRecorder struct {
	mock *MockVMScope
}

// NewMockVMScope creates a new mock instance.
func NewMockVMScope(ctrl *gomock.Controller) *MockVMScope {
	mock := &MockVMScope{ctrl: ctrl}
	mock.recorder = &MockVMScopeMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockVMScope) EXPECT() *MockVMScopeMockRecorder {
	return m.recorder
}

// Info mocks base method.
func (m *MockVMScope) Info(msg string, keysAndValues ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{msg}
	for _, a := range keysAndValues {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Info", varargs...)
}

// Info indicates an expected call of Info.
func (mr *MockVMScopeMockRecorder) Info(msg interface{}, keysAndValues ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{msg}, keysAndValues...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Info", reflect.TypeOf((*MockVMScope)(nil).Info), varargs...)
}

// Enabled mocks base method.
func (m *MockVMScope) Enabled() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Enabled")
	ret0, _ := ret[0].(bool)
	return ret0
}

// Enabled indicates an expected call of Enabled.
func (mr *MockVMScopeMockRecorder) Enabled() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Enabled", reflect.TypeOf((*MockVMScope)(nil).Enabled))
}

// Error mocks base method.
func (m *MockVMScope) Error(err error, msg string, keysAndValues ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{err, msg}
	for _, a := range keysAndValues {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Error", varargs...)
}

// Error indicates an expected call of Error.
func (mr *MockVMScopeMockRecorder) Error(err, msg interface{}, keysAndValues ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{err, msg}, keysAndValues...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Error", reflect.TypeOf((*MockVMScope)(nil).Error), varargs...)
}

// V mocks base method.
func (m *MockVMScope) V(level int) logr.InfoLogger {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "V", level)
	ret0, _ := ret[0].(logr.InfoLogger)
	return ret0
}

// V indicates an expected call of V.
func (mr *MockVMScopeMockRecorder) V(level interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "V", reflect.TypeOf((*MockVMScope)(nil).V), level)
}

// WithValues mocks base method.
func (m *MockVMScope) WithValues(keysAndValues ...interface{}) logr.Logger {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range keysAndValues {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "WithValues", varargs...)
	ret0, _ := ret[0].(logr.Logger)
	return ret0
}

// WithValues indicates an expected call of WithValues.
func (mr *MockVMScopeMockRecorder) WithValues(keysAndValues ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WithValues", reflect.TypeOf((*MockVMScope)(nil).WithValues), keysAndValues...)
}

// WithName mocks base method.
func (m *MockVMScope) WithName(name string) logr.Logger {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WithName", name)
	ret0, _ := ret[0].(logr.Logger)
	return ret0
}

// WithName indicates an expected call of WithName.
func (mr *MockVMScopeMockRecorder) WithName(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WithName", reflect.TypeOf((*MockVMScope)(nil).WithName), name)
}

// SubscriptionID mocks base method.
func (m *MockVMScope) SubscriptionID() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SubscriptionID")
	ret0, _ := ret[0].(string)
	return ret0
}

// SubscriptionID indicates an expected call of SubscriptionID.
func (mr *MockVMScopeMockRecorder) SubscriptionID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubscriptionID", reflect.TypeOf((*MockVMScope)(nil).SubscriptionID))
}

// ClientID mocks base method.
func (m *MockVMScope) ClientID() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClientID")
	ret0, _ := ret[0].(string)
	return ret0
}

// ClientID indicates an expected call of ClientID.
func (mr *MockVMScopeMockRecorder) ClientID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClientID", reflect.TypeOf((*MockVMScope)(nil).ClientID))
}

// ClientSecret mocks base method.
func (m *MockVMScope) ClientSecret() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClientSecret")
	ret0, _ := ret[0].(string)
	return ret0
}

// ClientSecret indicates an expected call of ClientSecret.
func (mr *MockVMScopeMockRecorder) ClientSecret() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClientSecret", reflect.TypeOf((*MockVMScope)(nil).ClientSecret))
}

// CloudEnvironment mocks base method.
func (m *MockVMScope) CloudEnvironment() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CloudEnvironment")
	ret0, _ := ret[0].(string)
	return ret0
}

// CloudEnvironment indicates an expected call of CloudEnvironment.
func (mr *MockVMScopeMockRecorder) CloudEnvironment() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloudEnvironment", reflect.TypeOf((*MockVMScope)(nil).CloudEnvironment))
}

// TenantID mocks base method.
func (m *MockVMScope) TenantID() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TenantID")
	ret0, _ := ret[0].(string)
	return ret0
}

// TenantID indicates an expected call of TenantID.
func (mr *MockVMScopeMockRecorder) TenantID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TenantID", reflect.TypeOf((*MockVMScope)(nil).TenantID))
}

// BaseURI mocks base method.
func (m *MockVMScope) BaseURI() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BaseURI")
	ret0, _ := ret[0].(string)
	return ret0
}

// BaseURI indicates an expected call of BaseURI.
func (mr *MockVMScopeMockRecorder) BaseURI() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BaseURI", reflect.TypeOf((*MockVMScope)(nil).BaseURI))
}

// Authorizer mocks base method.
func (m *MockVMScope) Authorizer() autorest.Authorizer {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Authorizer")
	ret0, _ := ret[0].(autorest.Authorizer)
	return ret0
}

// Authorizer indicates an expected call of Authorizer.
func (mr *MockVMScopeMockRecorder) Authorizer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Authorizer", reflect.TypeOf((*MockVMScope)(nil).Authorizer))
}

// ResourceGroup mocks base method.
func (m *MockVMScope) ResourceGroup() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResourceGroup")
	ret0, _ := ret[0].(string)
	return ret0
}

// ResourceGroup indicates an expected call of ResourceGroup.
func (mr *MockVMScopeMockRecorder) ResourceGroup() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResourceGroup", reflect.TypeOf((*MockVMScope)(nil).ResourceGroup))
}

// ClusterName mocks base method.
func (m *MockVMScope) ClusterName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClusterName")
	ret0, _ := ret[0].(string)
	return ret0
}

// ClusterName indicates an expected call of ClusterName.
func (mr *MockVMScopeMockRecorder) ClusterName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClusterName", reflect.TypeOf((*MockVMScope)(nil).ClusterName))
}

// Location mocks base method.
func (m *MockVMScope) Location() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Location")
	ret0, _ := ret[0].(string)
	return ret0
}

// Location indicates an expected call of Location.
func (mr *MockVMScopeMockRecorder) Location() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Location", reflect.TypeOf((*MockVMScope)(nil).Location))
}

// AdditionalTags mocks base method.
func (m *MockVMScope) AdditionalTags() v1alpha3.Tags {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AdditionalTags")
	ret0, _ := ret[0].(v1alpha3.Tags)
	return ret0
}

// AdditionalTags indicates an expected call of AdditionalTags.
func (mr *MockVMScopeMockRecorder) AdditionalTags() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AdditionalTags", reflect.TypeOf((*MockVMScope)(nil).AdditionalTags))
}

// VMSpecs mocks base method.
func (m *MockVMScope) VMSpecs() []azure.VMSpec {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VMSpecs")
	ret0, _ := ret[0].([]azure.VMSpec)
	return ret0
}

// VMSpecs indicates an expected call of VMSpecs.
func (mr *MockVMScopeMockRecorder) VMSpecs() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VMSpecs", reflect.TypeOf((*MockVMScope)(nil).VMSpecs))
}

// GetBootstrapData mocks base method.
func (m *MockVMScope) GetBootstrapData(ctx context.Context) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBootstrapData", ctx)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBootstrapData indicates an expected call of GetBootstrapData.
func (mr *MockVMScopeMockRecorder) GetBootstrapData(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBootstrapData", reflect.TypeOf((*MockVMScope)(nil).GetBootstrapData), ctx)
}

// GetVMImage mocks base method.
func (m *MockVMScope) GetVMImage() (*v1alpha3.Image, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetVMImage")
	ret0, _ := ret[0].(*v1alpha3.Image)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetVMImage indicates an expected call of GetVMImage.
func (mr *MockVMScopeMockRecorder) GetVMImage() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetVMImage", reflect.TypeOf((*MockVMScope)(nil).GetVMImage))
}

// SetAnnotation mocks base method.
func (m *MockVMScope) SetAnnotation(arg0, arg1 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetAnnotation", arg0, arg1)
}

// SetAnnotation indicates an expected call of SetAnnotation.
func (mr *MockVMScopeMockRecorder) SetAnnotation(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetAnnotation", reflect.TypeOf((*MockVMScope)(nil).SetAnnotation), arg0, arg1)
}

// SetProviderID mocks base method.
func (m *MockVMScope) SetProviderID(arg0 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetProviderID", arg0)
}

// SetProviderID indicates an expected call of SetProviderID.
func (mr *MockVMScopeMockRecorder) SetProviderID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetProviderID", reflect.TypeOf((*MockVMScope)(nil).SetProviderID), arg0)
}

// SetAddresses mocks base method.
func (m *MockVMScope) SetAddresses(arg0 []v1.NodeAddress) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetAddresses", arg0)
}

// SetAddresses indicates an expected call of SetAddresses.
func (mr *MockVMScopeMockRecorder) SetAddresses(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetAddresses", reflect.TypeOf((*MockVMScope)(nil).SetAddresses), arg0)
}

// SetVMState mocks base method.
func (m *MockVMScope) SetVMState(arg0 v1alpha3.VMState) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetVMState", arg0)
}

// SetVMState indicates an expected call of SetVMState.
func (mr *MockVMScopeMockRecorder) SetVMState(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetVMState", reflect.TypeOf((*MockVMScope)(nil).SetVMState), arg0)
}
