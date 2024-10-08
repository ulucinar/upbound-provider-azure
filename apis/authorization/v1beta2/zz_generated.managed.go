// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0
// Code generated by angryjet. DO NOT EDIT.

package v1beta2

import xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"

// GetCondition of this ResourceGroupPolicyAssignment.
func (mg *ResourceGroupPolicyAssignment) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this ResourceGroupPolicyAssignment.
func (mg *ResourceGroupPolicyAssignment) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetManagementPolicies of this ResourceGroupPolicyAssignment.
func (mg *ResourceGroupPolicyAssignment) GetManagementPolicies() xpv1.ManagementPolicies {
	return mg.Spec.ManagementPolicies
}

// GetProviderConfigReference of this ResourceGroupPolicyAssignment.
func (mg *ResourceGroupPolicyAssignment) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

// GetPublishConnectionDetailsTo of this ResourceGroupPolicyAssignment.
func (mg *ResourceGroupPolicyAssignment) GetPublishConnectionDetailsTo() *xpv1.PublishConnectionDetailsTo {
	return mg.Spec.PublishConnectionDetailsTo
}

// GetWriteConnectionSecretToReference of this ResourceGroupPolicyAssignment.
func (mg *ResourceGroupPolicyAssignment) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this ResourceGroupPolicyAssignment.
func (mg *ResourceGroupPolicyAssignment) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this ResourceGroupPolicyAssignment.
func (mg *ResourceGroupPolicyAssignment) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetManagementPolicies of this ResourceGroupPolicyAssignment.
func (mg *ResourceGroupPolicyAssignment) SetManagementPolicies(r xpv1.ManagementPolicies) {
	mg.Spec.ManagementPolicies = r
}

// SetProviderConfigReference of this ResourceGroupPolicyAssignment.
func (mg *ResourceGroupPolicyAssignment) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

// SetPublishConnectionDetailsTo of this ResourceGroupPolicyAssignment.
func (mg *ResourceGroupPolicyAssignment) SetPublishConnectionDetailsTo(r *xpv1.PublishConnectionDetailsTo) {
	mg.Spec.PublishConnectionDetailsTo = r
}

// SetWriteConnectionSecretToReference of this ResourceGroupPolicyAssignment.
func (mg *ResourceGroupPolicyAssignment) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this ResourcePolicyAssignment.
func (mg *ResourcePolicyAssignment) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this ResourcePolicyAssignment.
func (mg *ResourcePolicyAssignment) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetManagementPolicies of this ResourcePolicyAssignment.
func (mg *ResourcePolicyAssignment) GetManagementPolicies() xpv1.ManagementPolicies {
	return mg.Spec.ManagementPolicies
}

// GetProviderConfigReference of this ResourcePolicyAssignment.
func (mg *ResourcePolicyAssignment) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

// GetPublishConnectionDetailsTo of this ResourcePolicyAssignment.
func (mg *ResourcePolicyAssignment) GetPublishConnectionDetailsTo() *xpv1.PublishConnectionDetailsTo {
	return mg.Spec.PublishConnectionDetailsTo
}

// GetWriteConnectionSecretToReference of this ResourcePolicyAssignment.
func (mg *ResourcePolicyAssignment) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this ResourcePolicyAssignment.
func (mg *ResourcePolicyAssignment) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this ResourcePolicyAssignment.
func (mg *ResourcePolicyAssignment) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetManagementPolicies of this ResourcePolicyAssignment.
func (mg *ResourcePolicyAssignment) SetManagementPolicies(r xpv1.ManagementPolicies) {
	mg.Spec.ManagementPolicies = r
}

// SetProviderConfigReference of this ResourcePolicyAssignment.
func (mg *ResourcePolicyAssignment) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

// SetPublishConnectionDetailsTo of this ResourcePolicyAssignment.
func (mg *ResourcePolicyAssignment) SetPublishConnectionDetailsTo(r *xpv1.PublishConnectionDetailsTo) {
	mg.Spec.PublishConnectionDetailsTo = r
}

// SetWriteConnectionSecretToReference of this ResourcePolicyAssignment.
func (mg *ResourcePolicyAssignment) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this SubscriptionPolicyAssignment.
func (mg *SubscriptionPolicyAssignment) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this SubscriptionPolicyAssignment.
func (mg *SubscriptionPolicyAssignment) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetManagementPolicies of this SubscriptionPolicyAssignment.
func (mg *SubscriptionPolicyAssignment) GetManagementPolicies() xpv1.ManagementPolicies {
	return mg.Spec.ManagementPolicies
}

// GetProviderConfigReference of this SubscriptionPolicyAssignment.
func (mg *SubscriptionPolicyAssignment) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

// GetPublishConnectionDetailsTo of this SubscriptionPolicyAssignment.
func (mg *SubscriptionPolicyAssignment) GetPublishConnectionDetailsTo() *xpv1.PublishConnectionDetailsTo {
	return mg.Spec.PublishConnectionDetailsTo
}

// GetWriteConnectionSecretToReference of this SubscriptionPolicyAssignment.
func (mg *SubscriptionPolicyAssignment) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this SubscriptionPolicyAssignment.
func (mg *SubscriptionPolicyAssignment) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this SubscriptionPolicyAssignment.
func (mg *SubscriptionPolicyAssignment) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetManagementPolicies of this SubscriptionPolicyAssignment.
func (mg *SubscriptionPolicyAssignment) SetManagementPolicies(r xpv1.ManagementPolicies) {
	mg.Spec.ManagementPolicies = r
}

// SetProviderConfigReference of this SubscriptionPolicyAssignment.
func (mg *SubscriptionPolicyAssignment) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

// SetPublishConnectionDetailsTo of this SubscriptionPolicyAssignment.
func (mg *SubscriptionPolicyAssignment) SetPublishConnectionDetailsTo(r *xpv1.PublishConnectionDetailsTo) {
	mg.Spec.PublishConnectionDetailsTo = r
}

// SetWriteConnectionSecretToReference of this SubscriptionPolicyAssignment.
func (mg *SubscriptionPolicyAssignment) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}
