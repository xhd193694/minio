/*
 * MinIO Cloud Storage, (C) 2019 MinIO, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package iampolicy

import (
	"github.com/minio/minio/pkg/bucket/policy/condition"
)

// AdminAction - admin policy action.
type AdminAction string

const (
	// HealAdminAction - allows heal command
	HealAdminAction = "admin:Heal"

	// Service Actions

	// StorageInfoAdminAction - allow listing server info
	StorageInfoAdminAction = "admin:StorageInfo"
	// AccountingUsageInfoAdminAction - allow listing accounting usage info
	AccountingUsageInfoAdminAction = "admin:AccountingUsageInfo"
	// DataUsageInfoAdminAction - allow listing data usage info
	DataUsageInfoAdminAction = "admin:DataUsageInfo"
	// PerfInfoAdminAction - allow listing performance info
	PerfInfoAdminAction = "admin:PerfInfo"
	// TopLocksAdminAction - allow listing top locks
	TopLocksAdminAction = "admin:TopLocksInfo"
	// ProfilingAdminAction - allow profiling
	ProfilingAdminAction = "admin:Profiling"
	// TraceAdminAction - allow listing server trace
	TraceAdminAction = "admin:ServerTrace"
	// ConsoleLogAdminAction - allow listing console logs on terminal
	ConsoleLogAdminAction = "admin:ConsoleLog"
	// KMSKeyStatusAdminAction - allow getting KMS key status
	KMSKeyStatusAdminAction = "admin:KMSKeyStatus"
	// ServerHardwareInfoAdminAction - allow listing server hardware info
	ServerHardwareInfoAdminAction = "admin:HardwareInfo"
	// ServerInfoAdminAction - allow listing server info
	ServerInfoAdminAction = "admin:ServerInfo"

	// ServerUpdateAdminAction - allow MinIO binary update
	ServerUpdateAdminAction = "admin:ServerUpdate"

	//Config Actions

	// ConfigUpdateAdminAction - allow MinIO config management
	ConfigUpdateAdminAction = "admin:ConfigUpdate"

	// User Actions

	// CreateUserAdminAction - allow creating MinIO user
	CreateUserAdminAction = "admin:CreateUser"
	// DeleteUserAdminAction - allow deleting MinIO user
	DeleteUserAdminAction = "admin:DeleteUser"
	// ListUsersAdminAction - allow list users permission
	ListUsersAdminAction = "admin:ListUsers"
	// EnableUserAdminAction - allow enable user permission
	EnableUserAdminAction = "admin:EnableUser"
	// DisableUserAdminAction - allow disable user permission
	DisableUserAdminAction = "admin:DisableUser"
	// GetUserAdminAction - allows GET permission on user info
	GetUserAdminAction = "admin:GetUser"

	// Group Actions

	// AddUserToGroupAdminAction - allow adding user to group permission
	AddUserToGroupAdminAction = "admin:AddUserToGroup"
	// RemoveUserFromGroupAdminAction - allow removing user to group permission
	RemoveUserFromGroupAdminAction = "admin:RemoveUserFromGroup"
	// GetGroupAdminAction - allow getting group info
	GetGroupAdminAction = "admin:GetGroup"
	// ListGroupsAdminAction - allow list groups permission
	ListGroupsAdminAction = "admin:ListGroups"
	// EnableGroupAdminAction - allow enable group permission
	EnableGroupAdminAction = "admin:EnableGroup"
	// DisableGroupAdminAction - allow disable group permission
	DisableGroupAdminAction = "admin:DisableGroup"

	// Policy Actions

	// CreatePolicyAdminAction - allow create policy permission
	CreatePolicyAdminAction = "admin:CreatePolicy"
	// DeletePolicyAdminAction - allow delete policy permission
	DeletePolicyAdminAction = "admin:DeletePolicy"
	// GetPolicyAdminAction - allow get policy permission
	GetPolicyAdminAction = "admin:GetPolicy"
	// AttachPolicyAdminAction - allows attaching a policy to a user/group
	AttachPolicyAdminAction = "admin:AttachUserOrGroupPolicy"
	// ListUserPoliciesAdminAction - allows listing user policies
	ListUserPoliciesAdminAction = "admin:ListUserPolicies"
	// AllAdminActions - provides all admin permissions
	AllAdminActions = "admin:*"
)

// List of all supported admin actions.
var supportedAdminActions = map[AdminAction]struct{}{
	AllAdminActions:                {},
	HealAdminAction:                {},
	ServerInfoAdminAction:          {},
	StorageInfoAdminAction:         {},
	DataUsageInfoAdminAction:       {},
	PerfInfoAdminAction:            {},
	TopLocksAdminAction:            {},
	ProfilingAdminAction:           {},
	TraceAdminAction:               {},
	ConsoleLogAdminAction:          {},
	KMSKeyStatusAdminAction:        {},
	ServerHardwareInfoAdminAction:  {},
	ServerUpdateAdminAction:        {},
	ConfigUpdateAdminAction:        {},
	CreateUserAdminAction:          {},
	DeleteUserAdminAction:          {},
	ListUsersAdminAction:           {},
	EnableUserAdminAction:          {},
	DisableUserAdminAction:         {},
	GetUserAdminAction:             {},
	AddUserToGroupAdminAction:      {},
	RemoveUserFromGroupAdminAction: {},
	ListGroupsAdminAction:          {},
	EnableGroupAdminAction:         {},
	DisableGroupAdminAction:        {},
	CreatePolicyAdminAction:        {},
	DeletePolicyAdminAction:        {},
	GetPolicyAdminAction:           {},
	AttachPolicyAdminAction:        {},
	ListUserPoliciesAdminAction:    {},
}

func parseAdminAction(s string) (AdminAction, error) {
	action := AdminAction(s)
	if action.IsValid() {
		return action, nil
	}

	return action, Errorf("unsupported action '%v'", s)
}

// IsValid - checks if action is valid or not.
func (action AdminAction) IsValid() bool {
	_, ok := supportedAdminActions[action]
	return ok
}

// adminActionConditionKeyMap - holds mapping of supported condition key for an action.
var adminActionConditionKeyMap = map[Action]condition.KeySet{
	AllAdminActions:                condition.NewKeySet(condition.AllSupportedAdminKeys...),
	HealAdminAction:                condition.NewKeySet(condition.AllSupportedAdminKeys...),
	StorageInfoAdminAction:         condition.NewKeySet(condition.AllSupportedAdminKeys...),
	ServerInfoAdminAction:          condition.NewKeySet(condition.AllSupportedAdminKeys...),
	DataUsageInfoAdminAction:       condition.NewKeySet(condition.AllSupportedAdminKeys...),
	PerfInfoAdminAction:            condition.NewKeySet(condition.AllSupportedAdminKeys...),
	TopLocksAdminAction:            condition.NewKeySet(condition.AllSupportedAdminKeys...),
	ProfilingAdminAction:           condition.NewKeySet(condition.AllSupportedAdminKeys...),
	TraceAdminAction:               condition.NewKeySet(condition.AllSupportedAdminKeys...),
	ConsoleLogAdminAction:          condition.NewKeySet(condition.AllSupportedAdminKeys...),
	KMSKeyStatusAdminAction:        condition.NewKeySet(condition.AllSupportedAdminKeys...),
	ServerHardwareInfoAdminAction:  condition.NewKeySet(condition.AllSupportedAdminKeys...),
	ServerUpdateAdminAction:        condition.NewKeySet(condition.AllSupportedAdminKeys...),
	ConfigUpdateAdminAction:        condition.NewKeySet(condition.AllSupportedAdminKeys...),
	CreateUserAdminAction:          condition.NewKeySet(condition.AllSupportedAdminKeys...),
	DeleteUserAdminAction:          condition.NewKeySet(condition.AllSupportedAdminKeys...),
	ListUsersAdminAction:           condition.NewKeySet(condition.AllSupportedAdminKeys...),
	EnableUserAdminAction:          condition.NewKeySet(condition.AllSupportedAdminKeys...),
	DisableUserAdminAction:         condition.NewKeySet(condition.AllSupportedAdminKeys...),
	GetUserAdminAction:             condition.NewKeySet(condition.AllSupportedAdminKeys...),
	AddUserToGroupAdminAction:      condition.NewKeySet(condition.AllSupportedAdminKeys...),
	RemoveUserFromGroupAdminAction: condition.NewKeySet(condition.AllSupportedAdminKeys...),
	ListGroupsAdminAction:          condition.NewKeySet(condition.AllSupportedAdminKeys...),
	EnableGroupAdminAction:         condition.NewKeySet(condition.AllSupportedAdminKeys...),
	DisableGroupAdminAction:        condition.NewKeySet(condition.AllSupportedAdminKeys...),
	CreatePolicyAdminAction:        condition.NewKeySet(condition.AllSupportedAdminKeys...),
	DeletePolicyAdminAction:        condition.NewKeySet(condition.AllSupportedAdminKeys...),
	GetPolicyAdminAction:           condition.NewKeySet(condition.AllSupportedAdminKeys...),
	AttachPolicyAdminAction:        condition.NewKeySet(condition.AllSupportedAdminKeys...),
	ListUserPoliciesAdminAction:    condition.NewKeySet(condition.AllSupportedAdminKeys...),
}
