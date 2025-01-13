// SPDX-FileCopyrightText: 2025 Upbound Inc. <https://upbound.io>
//
// SPDX-License-Identifier: Apache-2.0
// Code generated by angryjet. DO NOT EDIT.

package v1beta1

import xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"

// GetCondition of this ProviderConfig.
func (p *ProviderConfig) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return p.Status.GetCondition(ct)
}

// GetUsers of this ProviderConfig.
func (p *ProviderConfig) GetUsers() int64 {
	return p.Status.Users
}

// SetConditions of this ProviderConfig.
func (p *ProviderConfig) SetConditions(c ...xpv1.Condition) {
	p.Status.SetConditions(c...)
}

// SetUsers of this ProviderConfig.
func (p *ProviderConfig) SetUsers(i int64) {
	p.Status.Users = i
}
