// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/aws/amazon-vpc-resource-controller-k8s/pkg/aws/ec2/api (interfaces: EC2APIHelper)

// Package mock_api is a generated GoMock package.
package mock_api

import (
	reflect "reflect"

	ec2 "github.com/aws/aws-sdk-go/service/ec2"
	gomock "github.com/golang/mock/gomock"
)

// MockEC2APIHelper is a mock of EC2APIHelper interface.
type MockEC2APIHelper struct {
	ctrl     *gomock.Controller
	recorder *MockEC2APIHelperMockRecorder
}

// MockEC2APIHelperMockRecorder is the mock recorder for MockEC2APIHelper.
type MockEC2APIHelperMockRecorder struct {
	mock *MockEC2APIHelper
}

// NewMockEC2APIHelper creates a new mock instance.
func NewMockEC2APIHelper(ctrl *gomock.Controller) *MockEC2APIHelper {
	mock := &MockEC2APIHelper{ctrl: ctrl}
	mock.recorder = &MockEC2APIHelperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEC2APIHelper) EXPECT() *MockEC2APIHelperMockRecorder {
	return m.recorder
}

// AssignIPv4AddressesAndWaitTillReady mocks base method.
func (m *MockEC2APIHelper) AssignIPv4AddressesAndWaitTillReady(arg0 string, arg1 int) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AssignIPv4AddressesAndWaitTillReady", arg0, arg1)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AssignIPv4PrefixesAndWaitTillReady mocks base method.
func (m *MockEC2APIHelper) AssignIPv4PrefixesAndWaitTillReady(arg0 string, arg1 int) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AssignIPv4PrefixesAndWaitTillReady", arg0, arg1)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AssignIPv4AddressesAndWaitTillReady indicates an expected call of AssignIPv4AddressesAndWaitTillReady.
func (mr *MockEC2APIHelperMockRecorder) AssignIPv4AddressesAndWaitTillReady(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AssignIPv4AddressesAndWaitTillReady", reflect.TypeOf((*MockEC2APIHelper)(nil).AssignIPv4AddressesAndWaitTillReady), arg0, arg1)
}

// AssociateBranchToTrunk mocks base method.
func (m *MockEC2APIHelper) AssociateBranchToTrunk(arg0, arg1 *string, arg2 int) (*ec2.AssociateTrunkInterfaceOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AssociateBranchToTrunk", arg0, arg1, arg2)
	ret0, _ := ret[0].(*ec2.AssociateTrunkInterfaceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AssociateBranchToTrunk indicates an expected call of AssociateBranchToTrunk.
func (mr *MockEC2APIHelperMockRecorder) AssociateBranchToTrunk(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AssociateBranchToTrunk", reflect.TypeOf((*MockEC2APIHelper)(nil).AssociateBranchToTrunk), arg0, arg1, arg2)
}

// AttachNetworkInterfaceToInstance mocks base method.
func (m *MockEC2APIHelper) AttachNetworkInterfaceToInstance(arg0, arg1 *string, arg2 *int64) (*string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AttachNetworkInterfaceToInstance", arg0, arg1, arg2)
	ret0, _ := ret[0].(*string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AttachNetworkInterfaceToInstance indicates an expected call of AttachNetworkInterfaceToInstance.
func (mr *MockEC2APIHelperMockRecorder) AttachNetworkInterfaceToInstance(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AttachNetworkInterfaceToInstance", reflect.TypeOf((*MockEC2APIHelper)(nil).AttachNetworkInterfaceToInstance), arg0, arg1, arg2)
}

// CreateAndAttachNetworkInterface mocks base method.
func (m *MockEC2APIHelper) CreateAndAttachNetworkInterface(arg0, arg1 *string, arg2 []string, arg3 []*ec2.Tag, arg4 *int64, arg5, arg6 *string, arg7 int) (*ec2.NetworkInterface, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateAndAttachNetworkInterface", arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7)
	ret0, _ := ret[0].(*ec2.NetworkInterface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateAndAttachNetworkInterface indicates an expected call of CreateAndAttachNetworkInterface.
func (mr *MockEC2APIHelperMockRecorder) CreateAndAttachNetworkInterface(arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAndAttachNetworkInterface", reflect.TypeOf((*MockEC2APIHelper)(nil).CreateAndAttachNetworkInterface), arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7)
}

// CreateNetworkInterface mocks base method.
func (m *MockEC2APIHelper) CreateNetworkInterface(arg0, arg1 *string, arg2 []string, arg3 []*ec2.Tag, arg4 int, arg5 *string) (*ec2.NetworkInterface, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateNetworkInterface", arg0, arg1, arg2, arg3, arg4, arg5)
	ret0, _ := ret[0].(*ec2.NetworkInterface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateNetworkInterface indicates an expected call of CreateNetworkInterface.
func (mr *MockEC2APIHelperMockRecorder) CreateNetworkInterface(arg0, arg1, arg2, arg3, arg4, arg5 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateNetworkInterface", reflect.TypeOf((*MockEC2APIHelper)(nil).CreateNetworkInterface), arg0, arg1, arg2, arg3, arg4, arg5)
}

// DeleteNetworkInterface mocks base method.
func (m *MockEC2APIHelper) DeleteNetworkInterface(arg0 *string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteNetworkInterface", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteNetworkInterface indicates an expected call of DeleteNetworkInterface.
func (mr *MockEC2APIHelperMockRecorder) DeleteNetworkInterface(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteNetworkInterface", reflect.TypeOf((*MockEC2APIHelper)(nil).DeleteNetworkInterface), arg0)
}

// DescribeNetworkInterfaces mocks base method.
func (m *MockEC2APIHelper) DescribeNetworkInterfaces(arg0 []*string) ([]*ec2.NetworkInterface, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeNetworkInterfaces", arg0)
	ret0, _ := ret[0].([]*ec2.NetworkInterface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeNetworkInterfaces indicates an expected call of DescribeNetworkInterfaces.
func (mr *MockEC2APIHelperMockRecorder) DescribeNetworkInterfaces(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeNetworkInterfaces", reflect.TypeOf((*MockEC2APIHelper)(nil).DescribeNetworkInterfaces), arg0)
}

// DescribeTrunkInterfaceAssociation mocks base method.
func (m *MockEC2APIHelper) DescribeTrunkInterfaceAssociation(arg0 *string) ([]*ec2.TrunkInterfaceAssociation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeTrunkInterfaceAssociation", arg0)
	ret0, _ := ret[0].([]*ec2.TrunkInterfaceAssociation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeTrunkInterfaceAssociation indicates an expected call of DescribeTrunkInterfaceAssociation.
func (mr *MockEC2APIHelperMockRecorder) DescribeTrunkInterfaceAssociation(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeTrunkInterfaceAssociation", reflect.TypeOf((*MockEC2APIHelper)(nil).DescribeTrunkInterfaceAssociation), arg0)
}

// DetachAndDeleteNetworkInterface mocks base method.
func (m *MockEC2APIHelper) DetachAndDeleteNetworkInterface(arg0, arg1 *string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DetachAndDeleteNetworkInterface", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DetachAndDeleteNetworkInterface indicates an expected call of DetachAndDeleteNetworkInterface.
func (mr *MockEC2APIHelperMockRecorder) DetachAndDeleteNetworkInterface(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DetachAndDeleteNetworkInterface", reflect.TypeOf((*MockEC2APIHelper)(nil).DetachAndDeleteNetworkInterface), arg0, arg1)
}

// DetachNetworkInterfaceFromInstance mocks base method.
func (m *MockEC2APIHelper) DetachNetworkInterfaceFromInstance(arg0 *string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DetachNetworkInterfaceFromInstance", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DetachNetworkInterfaceFromInstance indicates an expected call of DetachNetworkInterfaceFromInstance.
func (mr *MockEC2APIHelperMockRecorder) DetachNetworkInterfaceFromInstance(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DetachNetworkInterfaceFromInstance", reflect.TypeOf((*MockEC2APIHelper)(nil).DetachNetworkInterfaceFromInstance), arg0)
}

// GetBranchNetworkInterface mocks base method.
func (m *MockEC2APIHelper) GetBranchNetworkInterface(arg0 *string) ([]*ec2.NetworkInterface, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBranchNetworkInterface", arg0)
	ret0, _ := ret[0].([]*ec2.NetworkInterface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBranchNetworkInterface indicates an expected call of GetBranchNetworkInterface.
func (mr *MockEC2APIHelperMockRecorder) GetBranchNetworkInterface(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBranchNetworkInterface", reflect.TypeOf((*MockEC2APIHelper)(nil).GetBranchNetworkInterface), arg0)
}

// GetInstanceDetails mocks base method.
func (m *MockEC2APIHelper) GetInstanceDetails(arg0 *string) (*ec2.Instance, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetInstanceDetails", arg0)
	ret0, _ := ret[0].(*ec2.Instance)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetInstanceDetails indicates an expected call of GetInstanceDetails.
func (mr *MockEC2APIHelperMockRecorder) GetInstanceDetails(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetInstanceDetails", reflect.TypeOf((*MockEC2APIHelper)(nil).GetInstanceDetails), arg0)
}

// GetInstanceNetworkInterface mocks base method.
func (m *MockEC2APIHelper) GetInstanceNetworkInterface(arg0 *string) ([]*ec2.InstanceNetworkInterface, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetInstanceNetworkInterface", arg0)
	ret0, _ := ret[0].([]*ec2.InstanceNetworkInterface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetInstanceNetworkInterface indicates an expected call of GetInstanceNetworkInterface.
func (mr *MockEC2APIHelperMockRecorder) GetInstanceNetworkInterface(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetInstanceNetworkInterface", reflect.TypeOf((*MockEC2APIHelper)(nil).GetInstanceNetworkInterface), arg0)
}

// GetSubnet mocks base method.
func (m *MockEC2APIHelper) GetSubnet(arg0 *string) (*ec2.Subnet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSubnet", arg0)
	ret0, _ := ret[0].(*ec2.Subnet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSubnet indicates an expected call of GetSubnet.
func (mr *MockEC2APIHelperMockRecorder) GetSubnet(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSubnet", reflect.TypeOf((*MockEC2APIHelper)(nil).GetSubnet), arg0)
}

// SetDeleteOnTermination mocks base method.
func (m *MockEC2APIHelper) SetDeleteOnTermination(arg0, arg1 *string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetDeleteOnTermination", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetDeleteOnTermination indicates an expected call of SetDeleteOnTermination.
func (mr *MockEC2APIHelperMockRecorder) SetDeleteOnTermination(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetDeleteOnTermination", reflect.TypeOf((*MockEC2APIHelper)(nil).SetDeleteOnTermination), arg0, arg1)
}

// UnassignPrivateIpAddresses mocks base method.
func (m *MockEC2APIHelper) UnassignPrivateIpAddresses(arg0 string, arg1 []string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UnassignPrivateIpAddresses", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UnassignPrivateIpPrefixes mocks base method.
func (m *MockEC2APIHelper) UnassignPrivateIpPrefixes(arg0 string, arg1 []string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UnassignPrivateIpPrefixes", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UnassignPrivateIpAddresses indicates an expected call of UnassignPrivateIpAddresses.
func (mr *MockEC2APIHelperMockRecorder) UnassignPrivateIpAddresses(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnassignPrivateIpAddresses", reflect.TypeOf((*MockEC2APIHelper)(nil).UnassignPrivateIpAddresses), arg0, arg1)
}

// WaitForNetworkInterfaceStatusChange mocks base method.
func (m *MockEC2APIHelper) WaitForNetworkInterfaceStatusChange(arg0 *string, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WaitForNetworkInterfaceStatusChange", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// WaitForNetworkInterfaceStatusChange indicates an expected call of WaitForNetworkInterfaceStatusChange.
func (mr *MockEC2APIHelperMockRecorder) WaitForNetworkInterfaceStatusChange(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WaitForNetworkInterfaceStatusChange", reflect.TypeOf((*MockEC2APIHelper)(nil).WaitForNetworkInterfaceStatusChange), arg0, arg1)
}
