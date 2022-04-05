/*
Copyright 2022 The Crossplane Authors.

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
// Code generated by angryjet. DO NOT EDIT.

package v1alpha1

import xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"

// GetCondition of this ActiveDirectoryAdministrator.
func (mg *ActiveDirectoryAdministrator) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this ActiveDirectoryAdministrator.
func (mg *ActiveDirectoryAdministrator) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this ActiveDirectoryAdministrator.
func (mg *ActiveDirectoryAdministrator) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this ActiveDirectoryAdministrator.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *ActiveDirectoryAdministrator) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetWriteConnectionSecretToReference of this ActiveDirectoryAdministrator.
func (mg *ActiveDirectoryAdministrator) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this ActiveDirectoryAdministrator.
func (mg *ActiveDirectoryAdministrator) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this ActiveDirectoryAdministrator.
func (mg *ActiveDirectoryAdministrator) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this ActiveDirectoryAdministrator.
func (mg *ActiveDirectoryAdministrator) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this ActiveDirectoryAdministrator.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *ActiveDirectoryAdministrator) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetWriteConnectionSecretToReference of this ActiveDirectoryAdministrator.
func (mg *ActiveDirectoryAdministrator) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this Database.
func (mg *Database) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this Database.
func (mg *Database) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this Database.
func (mg *Database) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this Database.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *Database) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetWriteConnectionSecretToReference of this Database.
func (mg *Database) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this Database.
func (mg *Database) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this Database.
func (mg *Database) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this Database.
func (mg *Database) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this Database.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *Database) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetWriteConnectionSecretToReference of this Database.
func (mg *Database) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this ElasticPool.
func (mg *ElasticPool) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this ElasticPool.
func (mg *ElasticPool) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this ElasticPool.
func (mg *ElasticPool) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this ElasticPool.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *ElasticPool) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetWriteConnectionSecretToReference of this ElasticPool.
func (mg *ElasticPool) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this ElasticPool.
func (mg *ElasticPool) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this ElasticPool.
func (mg *ElasticPool) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this ElasticPool.
func (mg *ElasticPool) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this ElasticPool.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *ElasticPool) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetWriteConnectionSecretToReference of this ElasticPool.
func (mg *ElasticPool) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this FirewallRule.
func (mg *FirewallRule) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this FirewallRule.
func (mg *FirewallRule) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this FirewallRule.
func (mg *FirewallRule) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this FirewallRule.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *FirewallRule) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetWriteConnectionSecretToReference of this FirewallRule.
func (mg *FirewallRule) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this FirewallRule.
func (mg *FirewallRule) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this FirewallRule.
func (mg *FirewallRule) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this FirewallRule.
func (mg *FirewallRule) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this FirewallRule.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *FirewallRule) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetWriteConnectionSecretToReference of this FirewallRule.
func (mg *FirewallRule) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this MSSQLDatabase.
func (mg *MSSQLDatabase) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this MSSQLDatabase.
func (mg *MSSQLDatabase) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this MSSQLDatabase.
func (mg *MSSQLDatabase) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this MSSQLDatabase.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *MSSQLDatabase) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetWriteConnectionSecretToReference of this MSSQLDatabase.
func (mg *MSSQLDatabase) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this MSSQLDatabase.
func (mg *MSSQLDatabase) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this MSSQLDatabase.
func (mg *MSSQLDatabase) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this MSSQLDatabase.
func (mg *MSSQLDatabase) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this MSSQLDatabase.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *MSSQLDatabase) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetWriteConnectionSecretToReference of this MSSQLDatabase.
func (mg *MSSQLDatabase) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this MSSQLDatabaseExtendedAuditingPolicy.
func (mg *MSSQLDatabaseExtendedAuditingPolicy) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this MSSQLDatabaseExtendedAuditingPolicy.
func (mg *MSSQLDatabaseExtendedAuditingPolicy) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this MSSQLDatabaseExtendedAuditingPolicy.
func (mg *MSSQLDatabaseExtendedAuditingPolicy) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this MSSQLDatabaseExtendedAuditingPolicy.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *MSSQLDatabaseExtendedAuditingPolicy) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetWriteConnectionSecretToReference of this MSSQLDatabaseExtendedAuditingPolicy.
func (mg *MSSQLDatabaseExtendedAuditingPolicy) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this MSSQLDatabaseExtendedAuditingPolicy.
func (mg *MSSQLDatabaseExtendedAuditingPolicy) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this MSSQLDatabaseExtendedAuditingPolicy.
func (mg *MSSQLDatabaseExtendedAuditingPolicy) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this MSSQLDatabaseExtendedAuditingPolicy.
func (mg *MSSQLDatabaseExtendedAuditingPolicy) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this MSSQLDatabaseExtendedAuditingPolicy.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *MSSQLDatabaseExtendedAuditingPolicy) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetWriteConnectionSecretToReference of this MSSQLDatabaseExtendedAuditingPolicy.
func (mg *MSSQLDatabaseExtendedAuditingPolicy) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this MSSQLDatabaseVulnerabilityAssessmentRuleBaseline.
func (mg *MSSQLDatabaseVulnerabilityAssessmentRuleBaseline) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this MSSQLDatabaseVulnerabilityAssessmentRuleBaseline.
func (mg *MSSQLDatabaseVulnerabilityAssessmentRuleBaseline) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this MSSQLDatabaseVulnerabilityAssessmentRuleBaseline.
func (mg *MSSQLDatabaseVulnerabilityAssessmentRuleBaseline) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this MSSQLDatabaseVulnerabilityAssessmentRuleBaseline.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *MSSQLDatabaseVulnerabilityAssessmentRuleBaseline) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetWriteConnectionSecretToReference of this MSSQLDatabaseVulnerabilityAssessmentRuleBaseline.
func (mg *MSSQLDatabaseVulnerabilityAssessmentRuleBaseline) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this MSSQLDatabaseVulnerabilityAssessmentRuleBaseline.
func (mg *MSSQLDatabaseVulnerabilityAssessmentRuleBaseline) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this MSSQLDatabaseVulnerabilityAssessmentRuleBaseline.
func (mg *MSSQLDatabaseVulnerabilityAssessmentRuleBaseline) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this MSSQLDatabaseVulnerabilityAssessmentRuleBaseline.
func (mg *MSSQLDatabaseVulnerabilityAssessmentRuleBaseline) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this MSSQLDatabaseVulnerabilityAssessmentRuleBaseline.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *MSSQLDatabaseVulnerabilityAssessmentRuleBaseline) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetWriteConnectionSecretToReference of this MSSQLDatabaseVulnerabilityAssessmentRuleBaseline.
func (mg *MSSQLDatabaseVulnerabilityAssessmentRuleBaseline) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this MSSQLElasticPool.
func (mg *MSSQLElasticPool) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this MSSQLElasticPool.
func (mg *MSSQLElasticPool) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this MSSQLElasticPool.
func (mg *MSSQLElasticPool) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this MSSQLElasticPool.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *MSSQLElasticPool) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetWriteConnectionSecretToReference of this MSSQLElasticPool.
func (mg *MSSQLElasticPool) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this MSSQLElasticPool.
func (mg *MSSQLElasticPool) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this MSSQLElasticPool.
func (mg *MSSQLElasticPool) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this MSSQLElasticPool.
func (mg *MSSQLElasticPool) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this MSSQLElasticPool.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *MSSQLElasticPool) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetWriteConnectionSecretToReference of this MSSQLElasticPool.
func (mg *MSSQLElasticPool) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this MSSQLFailoverGroup.
func (mg *MSSQLFailoverGroup) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this MSSQLFailoverGroup.
func (mg *MSSQLFailoverGroup) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this MSSQLFailoverGroup.
func (mg *MSSQLFailoverGroup) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this MSSQLFailoverGroup.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *MSSQLFailoverGroup) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetWriteConnectionSecretToReference of this MSSQLFailoverGroup.
func (mg *MSSQLFailoverGroup) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this MSSQLFailoverGroup.
func (mg *MSSQLFailoverGroup) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this MSSQLFailoverGroup.
func (mg *MSSQLFailoverGroup) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this MSSQLFailoverGroup.
func (mg *MSSQLFailoverGroup) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this MSSQLFailoverGroup.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *MSSQLFailoverGroup) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetWriteConnectionSecretToReference of this MSSQLFailoverGroup.
func (mg *MSSQLFailoverGroup) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this MSSQLFirewallRule.
func (mg *MSSQLFirewallRule) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this MSSQLFirewallRule.
func (mg *MSSQLFirewallRule) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this MSSQLFirewallRule.
func (mg *MSSQLFirewallRule) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this MSSQLFirewallRule.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *MSSQLFirewallRule) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetWriteConnectionSecretToReference of this MSSQLFirewallRule.
func (mg *MSSQLFirewallRule) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this MSSQLFirewallRule.
func (mg *MSSQLFirewallRule) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this MSSQLFirewallRule.
func (mg *MSSQLFirewallRule) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this MSSQLFirewallRule.
func (mg *MSSQLFirewallRule) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this MSSQLFirewallRule.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *MSSQLFirewallRule) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetWriteConnectionSecretToReference of this MSSQLFirewallRule.
func (mg *MSSQLFirewallRule) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this MSSQLJobAgent.
func (mg *MSSQLJobAgent) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this MSSQLJobAgent.
func (mg *MSSQLJobAgent) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this MSSQLJobAgent.
func (mg *MSSQLJobAgent) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this MSSQLJobAgent.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *MSSQLJobAgent) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetWriteConnectionSecretToReference of this MSSQLJobAgent.
func (mg *MSSQLJobAgent) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this MSSQLJobAgent.
func (mg *MSSQLJobAgent) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this MSSQLJobAgent.
func (mg *MSSQLJobAgent) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this MSSQLJobAgent.
func (mg *MSSQLJobAgent) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this MSSQLJobAgent.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *MSSQLJobAgent) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetWriteConnectionSecretToReference of this MSSQLJobAgent.
func (mg *MSSQLJobAgent) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this MSSQLJobCredential.
func (mg *MSSQLJobCredential) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this MSSQLJobCredential.
func (mg *MSSQLJobCredential) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this MSSQLJobCredential.
func (mg *MSSQLJobCredential) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this MSSQLJobCredential.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *MSSQLJobCredential) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetWriteConnectionSecretToReference of this MSSQLJobCredential.
func (mg *MSSQLJobCredential) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this MSSQLJobCredential.
func (mg *MSSQLJobCredential) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this MSSQLJobCredential.
func (mg *MSSQLJobCredential) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this MSSQLJobCredential.
func (mg *MSSQLJobCredential) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this MSSQLJobCredential.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *MSSQLJobCredential) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetWriteConnectionSecretToReference of this MSSQLJobCredential.
func (mg *MSSQLJobCredential) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this MSSQLServerSecurityAlertPolicy.
func (mg *MSSQLServerSecurityAlertPolicy) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this MSSQLServerSecurityAlertPolicy.
func (mg *MSSQLServerSecurityAlertPolicy) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this MSSQLServerSecurityAlertPolicy.
func (mg *MSSQLServerSecurityAlertPolicy) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this MSSQLServerSecurityAlertPolicy.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *MSSQLServerSecurityAlertPolicy) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetWriteConnectionSecretToReference of this MSSQLServerSecurityAlertPolicy.
func (mg *MSSQLServerSecurityAlertPolicy) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this MSSQLServerSecurityAlertPolicy.
func (mg *MSSQLServerSecurityAlertPolicy) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this MSSQLServerSecurityAlertPolicy.
func (mg *MSSQLServerSecurityAlertPolicy) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this MSSQLServerSecurityAlertPolicy.
func (mg *MSSQLServerSecurityAlertPolicy) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this MSSQLServerSecurityAlertPolicy.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *MSSQLServerSecurityAlertPolicy) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetWriteConnectionSecretToReference of this MSSQLServerSecurityAlertPolicy.
func (mg *MSSQLServerSecurityAlertPolicy) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this MSSQLServerVulnerabilityAssessment.
func (mg *MSSQLServerVulnerabilityAssessment) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this MSSQLServerVulnerabilityAssessment.
func (mg *MSSQLServerVulnerabilityAssessment) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this MSSQLServerVulnerabilityAssessment.
func (mg *MSSQLServerVulnerabilityAssessment) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this MSSQLServerVulnerabilityAssessment.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *MSSQLServerVulnerabilityAssessment) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetWriteConnectionSecretToReference of this MSSQLServerVulnerabilityAssessment.
func (mg *MSSQLServerVulnerabilityAssessment) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this MSSQLServerVulnerabilityAssessment.
func (mg *MSSQLServerVulnerabilityAssessment) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this MSSQLServerVulnerabilityAssessment.
func (mg *MSSQLServerVulnerabilityAssessment) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this MSSQLServerVulnerabilityAssessment.
func (mg *MSSQLServerVulnerabilityAssessment) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this MSSQLServerVulnerabilityAssessment.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *MSSQLServerVulnerabilityAssessment) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetWriteConnectionSecretToReference of this MSSQLServerVulnerabilityAssessment.
func (mg *MSSQLServerVulnerabilityAssessment) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this MSSQLVirtualNetworkRule.
func (mg *MSSQLVirtualNetworkRule) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this MSSQLVirtualNetworkRule.
func (mg *MSSQLVirtualNetworkRule) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this MSSQLVirtualNetworkRule.
func (mg *MSSQLVirtualNetworkRule) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this MSSQLVirtualNetworkRule.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *MSSQLVirtualNetworkRule) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetWriteConnectionSecretToReference of this MSSQLVirtualNetworkRule.
func (mg *MSSQLVirtualNetworkRule) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this MSSQLVirtualNetworkRule.
func (mg *MSSQLVirtualNetworkRule) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this MSSQLVirtualNetworkRule.
func (mg *MSSQLVirtualNetworkRule) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this MSSQLVirtualNetworkRule.
func (mg *MSSQLVirtualNetworkRule) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this MSSQLVirtualNetworkRule.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *MSSQLVirtualNetworkRule) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetWriteConnectionSecretToReference of this MSSQLVirtualNetworkRule.
func (mg *MSSQLVirtualNetworkRule) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this ManagedDatabase.
func (mg *ManagedDatabase) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this ManagedDatabase.
func (mg *ManagedDatabase) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this ManagedDatabase.
func (mg *ManagedDatabase) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this ManagedDatabase.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *ManagedDatabase) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetWriteConnectionSecretToReference of this ManagedDatabase.
func (mg *ManagedDatabase) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this ManagedDatabase.
func (mg *ManagedDatabase) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this ManagedDatabase.
func (mg *ManagedDatabase) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this ManagedDatabase.
func (mg *ManagedDatabase) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this ManagedDatabase.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *ManagedDatabase) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetWriteConnectionSecretToReference of this ManagedDatabase.
func (mg *ManagedDatabase) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this ManagedInstance.
func (mg *ManagedInstance) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this ManagedInstance.
func (mg *ManagedInstance) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this ManagedInstance.
func (mg *ManagedInstance) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this ManagedInstance.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *ManagedInstance) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetWriteConnectionSecretToReference of this ManagedInstance.
func (mg *ManagedInstance) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this ManagedInstance.
func (mg *ManagedInstance) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this ManagedInstance.
func (mg *ManagedInstance) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this ManagedInstance.
func (mg *ManagedInstance) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this ManagedInstance.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *ManagedInstance) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetWriteConnectionSecretToReference of this ManagedInstance.
func (mg *ManagedInstance) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}