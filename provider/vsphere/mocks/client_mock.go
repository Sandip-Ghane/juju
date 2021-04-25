// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/provider/vsphere (interfaces: Client)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	vsphereclient "github.com/juju/juju/provider/vsphere/internal/vsphereclient"
	object "github.com/vmware/govmomi/object"
	mo "github.com/vmware/govmomi/vim25/mo"
	types "github.com/vmware/govmomi/vim25/types"
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

// Close mocks base method.
func (m *MockClient) Close(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockClientMockRecorder) Close(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockClient)(nil).Close), arg0)
}

// ComputeResources mocks base method.
func (m *MockClient) ComputeResources(arg0 context.Context) ([]vsphereclient.ComputeResource, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ComputeResources", arg0)
	ret0, _ := ret[0].([]vsphereclient.ComputeResource)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ComputeResources indicates an expected call of ComputeResources.
func (mr *MockClientMockRecorder) ComputeResources(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ComputeResources", reflect.TypeOf((*MockClient)(nil).ComputeResources), arg0)
}

// CreateTemplateVM mocks base method.
func (m *MockClient) CreateTemplateVM(arg0 context.Context, arg1 vsphereclient.ImportOVAParameters) (*object.VirtualMachine, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTemplateVM", arg0, arg1)
	ret0, _ := ret[0].(*object.VirtualMachine)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateTemplateVM indicates an expected call of CreateTemplateVM.
func (mr *MockClientMockRecorder) CreateTemplateVM(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTemplateVM", reflect.TypeOf((*MockClient)(nil).CreateTemplateVM), arg0, arg1)
}

// CreateVirtualMachine mocks base method.
func (m *MockClient) CreateVirtualMachine(arg0 context.Context, arg1 vsphereclient.CreateVirtualMachineParams) (*mo.VirtualMachine, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateVirtualMachine", arg0, arg1)
	ret0, _ := ret[0].(*mo.VirtualMachine)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateVirtualMachine indicates an expected call of CreateVirtualMachine.
func (mr *MockClientMockRecorder) CreateVirtualMachine(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateVirtualMachine", reflect.TypeOf((*MockClient)(nil).CreateVirtualMachine), arg0, arg1)
}

// Datastores mocks base method.
func (m *MockClient) Datastores(arg0 context.Context) ([]mo.Datastore, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Datastores", arg0)
	ret0, _ := ret[0].([]mo.Datastore)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Datastores indicates an expected call of Datastores.
func (mr *MockClientMockRecorder) Datastores(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Datastores", reflect.TypeOf((*MockClient)(nil).Datastores), arg0)
}

// DeleteDatastoreFile mocks base method.
func (m *MockClient) DeleteDatastoreFile(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteDatastoreFile", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteDatastoreFile indicates an expected call of DeleteDatastoreFile.
func (mr *MockClientMockRecorder) DeleteDatastoreFile(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteDatastoreFile", reflect.TypeOf((*MockClient)(nil).DeleteDatastoreFile), arg0, arg1)
}

// DestroyVMFolder mocks base method.
func (m *MockClient) DestroyVMFolder(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DestroyVMFolder", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DestroyVMFolder indicates an expected call of DestroyVMFolder.
func (mr *MockClientMockRecorder) DestroyVMFolder(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DestroyVMFolder", reflect.TypeOf((*MockClient)(nil).DestroyVMFolder), arg0, arg1)
}

// EnsureVMFolder mocks base method.
func (m *MockClient) EnsureVMFolder(arg0 context.Context, arg1, arg2 string) (*object.Folder, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnsureVMFolder", arg0, arg1, arg2)
	ret0, _ := ret[0].(*object.Folder)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EnsureVMFolder indicates an expected call of EnsureVMFolder.
func (mr *MockClientMockRecorder) EnsureVMFolder(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnsureVMFolder", reflect.TypeOf((*MockClient)(nil).EnsureVMFolder), arg0, arg1, arg2)
}

// FindFolder mocks base method.
func (m *MockClient) FindFolder(arg0 context.Context, arg1 string) (*object.Folder, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindFolder", arg0, arg1)
	ret0, _ := ret[0].(*object.Folder)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindFolder indicates an expected call of FindFolder.
func (mr *MockClientMockRecorder) FindFolder(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindFolder", reflect.TypeOf((*MockClient)(nil).FindFolder), arg0, arg1)
}

// FindVMTemplatesByName mocks base method.
func (m *MockClient) FindVMTemplatesByName(arg0 context.Context, arg1, arg2 string) ([]*object.VirtualMachine, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindVMTemplatesByName", arg0, arg1, arg2)
	ret0, _ := ret[0].([]*object.VirtualMachine)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindVMTemplatesByName indicates an expected call of FindVMTemplatesByName.
func (mr *MockClientMockRecorder) FindVMTemplatesByName(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindVMTemplatesByName", reflect.TypeOf((*MockClient)(nil).FindVMTemplatesByName), arg0, arg1, arg2)
}

// Folders mocks base method.
func (m *MockClient) Folders(arg0 context.Context) (*object.DatacenterFolders, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Folders", arg0)
	ret0, _ := ret[0].(*object.DatacenterFolders)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Folders indicates an expected call of Folders.
func (mr *MockClientMockRecorder) Folders(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Folders", reflect.TypeOf((*MockClient)(nil).Folders), arg0)
}

// GetTargetDatastore mocks base method.
func (m *MockClient) GetTargetDatastore(arg0 context.Context, arg1 *mo.ComputeResource, arg2 string) (*object.Datastore, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTargetDatastore", arg0, arg1, arg2)
	ret0, _ := ret[0].(*object.Datastore)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTargetDatastore indicates an expected call of GetTargetDatastore.
func (mr *MockClientMockRecorder) GetTargetDatastore(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTargetDatastore", reflect.TypeOf((*MockClient)(nil).GetTargetDatastore), arg0, arg1, arg2)
}

// ListVMTemplates mocks base method.
func (m *MockClient) ListVMTemplates(arg0 context.Context, arg1 string) ([]*object.VirtualMachine, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListVMTemplates", arg0, arg1)
	ret0, _ := ret[0].([]*object.VirtualMachine)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListVMTemplates indicates an expected call of ListVMTemplates.
func (mr *MockClientMockRecorder) ListVMTemplates(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListVMTemplates", reflect.TypeOf((*MockClient)(nil).ListVMTemplates), arg0, arg1)
}

// MoveVMFolderInto mocks base method.
func (m *MockClient) MoveVMFolderInto(arg0 context.Context, arg1, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MoveVMFolderInto", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// MoveVMFolderInto indicates an expected call of MoveVMFolderInto.
func (mr *MockClientMockRecorder) MoveVMFolderInto(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MoveVMFolderInto", reflect.TypeOf((*MockClient)(nil).MoveVMFolderInto), arg0, arg1, arg2)
}

// MoveVMsInto mocks base method.
func (m *MockClient) MoveVMsInto(arg0 context.Context, arg1 string, arg2 ...types.ManagedObjectReference) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "MoveVMsInto", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// MoveVMsInto indicates an expected call of MoveVMsInto.
func (mr *MockClientMockRecorder) MoveVMsInto(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MoveVMsInto", reflect.TypeOf((*MockClient)(nil).MoveVMsInto), varargs...)
}

// RemoveVirtualMachines mocks base method.
func (m *MockClient) RemoveVirtualMachines(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveVirtualMachines", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveVirtualMachines indicates an expected call of RemoveVirtualMachines.
func (mr *MockClientMockRecorder) RemoveVirtualMachines(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveVirtualMachines", reflect.TypeOf((*MockClient)(nil).RemoveVirtualMachines), arg0, arg1)
}

// ResourcePools mocks base method.
func (m *MockClient) ResourcePools(arg0 context.Context, arg1 string) ([]*object.ResourcePool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResourcePools", arg0, arg1)
	ret0, _ := ret[0].([]*object.ResourcePool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ResourcePools indicates an expected call of ResourcePools.
func (mr *MockClientMockRecorder) ResourcePools(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResourcePools", reflect.TypeOf((*MockClient)(nil).ResourcePools), arg0, arg1)
}

// UpdateVirtualMachineExtraConfig mocks base method.
func (m *MockClient) UpdateVirtualMachineExtraConfig(arg0 context.Context, arg1 *mo.VirtualMachine, arg2 map[string]string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateVirtualMachineExtraConfig", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateVirtualMachineExtraConfig indicates an expected call of UpdateVirtualMachineExtraConfig.
func (mr *MockClientMockRecorder) UpdateVirtualMachineExtraConfig(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateVirtualMachineExtraConfig", reflect.TypeOf((*MockClient)(nil).UpdateVirtualMachineExtraConfig), arg0, arg1, arg2)
}

// UserHasRootLevelPrivilege mocks base method.
func (m *MockClient) UserHasRootLevelPrivilege(arg0 context.Context, arg1 string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UserHasRootLevelPrivilege", arg0, arg1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UserHasRootLevelPrivilege indicates an expected call of UserHasRootLevelPrivilege.
func (mr *MockClientMockRecorder) UserHasRootLevelPrivilege(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UserHasRootLevelPrivilege", reflect.TypeOf((*MockClient)(nil).UserHasRootLevelPrivilege), arg0, arg1)
}

// VirtualMachineObjectToManagedObject mocks base method.
func (m *MockClient) VirtualMachineObjectToManagedObject(arg0 context.Context, arg1 *object.VirtualMachine) (mo.VirtualMachine, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VirtualMachineObjectToManagedObject", arg0, arg1)
	ret0, _ := ret[0].(mo.VirtualMachine)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// VirtualMachineObjectToManagedObject indicates an expected call of VirtualMachineObjectToManagedObject.
func (mr *MockClientMockRecorder) VirtualMachineObjectToManagedObject(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VirtualMachineObjectToManagedObject", reflect.TypeOf((*MockClient)(nil).VirtualMachineObjectToManagedObject), arg0, arg1)
}

// VirtualMachines mocks base method.
func (m *MockClient) VirtualMachines(arg0 context.Context, arg1 string) ([]*mo.VirtualMachine, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VirtualMachines", arg0, arg1)
	ret0, _ := ret[0].([]*mo.VirtualMachine)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// VirtualMachines indicates an expected call of VirtualMachines.
func (mr *MockClientMockRecorder) VirtualMachines(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VirtualMachines", reflect.TypeOf((*MockClient)(nil).VirtualMachines), arg0, arg1)
}
