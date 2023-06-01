package events

// Autogenerated file. DO NOT MODIFY DIRECTLY!
/*
 *  Copyright (c) 2022 Avesha, Inc. All rights reserved.
 *
 *  SPDX-License-Identifier: Apache-2.0
 *
 *  Licensed under the Apache License, Version 2.0 (the "License");
 *  you may not use this file except in compliance with the License.
 *  You may obtain a copy of the License at
 *
 *  http://www.apache.org/licenses/LICENSE-2.0
 *
 *  Unless required by applicable law or agreed to in writing, software
 *  distributed under the License is distributed on an "AS IS" BASIS,
 *  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *  See the License for the specific language governing permissions and
 *  limitations under the License.
 */

import "github.com/kubeslice/kubeslice-monitoring/pkg/events"

var EventsMap = map[events.EventName]*events.EventSchema{
	"ProjectDeleted": {
		Name:                "ProjectDeleted",
		Reason:              "ProjectDeleted",
		Action:              "DeleteProject",
		Type:                events.EventTypeWarning,
		ReportingController: "controller",
		Message:             "Project got deleted.",
	},
	"ProjectDeletionFailed": {
		Name:                "ProjectDeletionFailed",
		Reason:              "ProjectDeletionFailed",
		Action:              "DeleteProject",
		Type:                events.EventTypeWarning,
		ReportingController: "controller",
		Message:             "Project deletion failed.",
	},
	"ClusterInstallationInProgress": {
		Name:                "ClusterInstallationInProgress",
		Reason:              "ClusterInstallationInProgress",
		Action:              "AutoInstallation",
		Type:                events.EventTypeNormal,
		ReportingController: "controller",
		Message:             "Worker Operator installation in progress.",
	},
	"ClusterInstallationFailed": {
		Name:                "ClusterInstallationFailed",
		Reason:              "ClusterInstallationFailed",
		Action:              "AutoInstallation",
		Type:                events.EventTypeWarning,
		ReportingController: "controller",
		Message:             "Worker Operator installation failed.",
	},
	"ClusterInstallationPending": {
		Name:                "ClusterInstallationPending",
		Reason:              "ClusterInstallationPending",
		Action:              "AutoInstallation",
		Type:                events.EventTypeWarning,
		ReportingController: "controller",
		Message:             "Worker Operator verification failed. Pods might be in error state.",
	},
	"ClusterDeleted": {
		Name:                "ClusterDeleted",
		Reason:              "ClusterDeleted",
		Action:              "DeleteCluster",
		Type:                events.EventTypeWarning,
		ReportingController: "controller",
		Message:             "Cluster got deleted.",
	},
	"ClusterDeletionFailed": {
		Name:                "ClusterDeletionFailed",
		Reason:              "ClusterDeletionFailed",
		Action:              "DeleteCluster",
		Type:                events.EventTypeWarning,
		ReportingController: "controller",
		Message:             "Cluster deletion failed.",
	},
	"ClusterDeregistrationInProgress": {
		Name:                "ClusterDeregistrationInProgress",
		Reason:              "ClusterDeregistrationInProgress",
		Action:              "DeregisterCluster",
		Type:                events.EventTypeNormal,
		ReportingController: "controller",
		Message:             "Cluster deregistration in progress.",
	},
	"ClusterDeregistered": {
		Name:                "ClusterDeregistered",
		Reason:              "ClusterDeregistered",
		Action:              "DeregisterCluster",
		Type:                events.EventTypeNormal,
		ReportingController: "controller",
		Message:             "Cluster deregistered.",
	},
	"ClusterDeregisterTimeout": {
		Name:                "ClusterDeregisterTimeout",
		Reason:              "ClusterDeregisterTimeout",
		Action:              "DeregisterCluster",
		Type:                events.EventTypeWarning,
		ReportingController: "controller",
		Message:             "Cluster deregister timeout.",
	},
	"ClusterDeregisterFailed": {
		Name:                "ClusterDeregisterFailed",
		Reason:              "ClusterDeregisterFailed",
		Action:              "DeregisterCluster",
		Type:                events.EventTypeWarning,
		ReportingController: "controller",
		Message:             "Cluster deregister failed.",
	},
	"SliceConfigDeleted": {
		Name:                "SliceConfigDeleted",
		Reason:              "SliceConfigDeleted",
		Action:              "DeleteSliceConfig",
		Type:                events.EventTypeWarning,
		ReportingController: "controller",
		Message:             "Slice config got deleted.",
	},
	"SliceConfigDeletionFailed": {
		Name:                "SliceConfigDeletionFailed",
		Reason:              "SliceConfigDeletionFailed",
		Action:              "DeleteSliceConfig",
		Type:                events.EventTypeWarning,
		ReportingController: "controller",
		Message:             "Slice config deletion failed.",
	},
	"ServiceExportConfigDeleted": {
		Name:                "ServiceExportConfigDeleted",
		Reason:              "ServiceExportConfigDeleted",
		Action:              "DeleteServiceExportConfig",
		Type:                events.EventTypeWarning,
		ReportingController: "controller",
		Message:             "Service export config got deleted.",
	},
	"ServiceExportConfigDeletionFailed": {
		Name:                "ServiceExportConfigDeletionFailed",
		Reason:              "ServiceExportConfigDeletionFailed",
		Action:              "DeleteServiceExportConfig",
		Type:                events.EventTypeWarning,
		ReportingController: "controller",
		Message:             "Service export config deletion failed.",
	},
	"DefaultSliceQoSConfigCreated": {
		Name:                "DefaultSliceQoSConfigCreated",
		Reason:              "DefaultSliceQoSConfigCreated",
		Action:              "CreateSliceQoSConfig",
		Type:                events.EventTypeNormal,
		ReportingController: "controller",
		Message:             "Default Slice QoS config created.",
	},
	"SliceQoSConfigDeleted": {
		Name:                "SliceQoSConfigDeleted",
		Reason:              "SliceQoSConfigDeleted",
		Action:              "DeleteSliceQoSConfig",
		Type:                events.EventTypeWarning,
		ReportingController: "controller",
		Message:             "Slice QoS config got deleted.",
	},
	"SliceQoSConfigDeletionFailed": {
		Name:                "SliceQoSConfigDeletionFailed",
		Reason:              "SliceQoSConfigDeletionFailed",
		Action:              "DeleteSliceQoSConfig",
		Type:                events.EventTypeWarning,
		ReportingController: "controller",
		Message:             "Slice QoS config deletion failed.",
	},
	"SecretDeleted": {
		Name:                "SecretDeleted",
		Reason:              "SecretDeleted",
		Action:              "DeleteSecret",
		Type:                events.EventTypeWarning,
		ReportingController: "controller",
		Message:             "Secret got deleted.",
	},
	"SecretDeletionFailed": {
		Name:                "SecretDeletionFailed",
		Reason:              "SecretDeletionFailed",
		Action:              "DeleteSecret",
		Type:                events.EventTypeWarning,
		ReportingController: "controller",
		Message:             "Secret deletion failed.",
	},
	"NamespaceCreated": {
		Name:                "NamespaceCreated",
		Reason:              "NamespaceCreated",
		Action:              "CreateNamespace",
		Type:                events.EventTypeNormal,
		ReportingController: "controller",
		Message:             "Namespace got created.",
	},
	"NamespaceCreationFailed": {
		Name:                "NamespaceCreationFailed",
		Reason:              "NamespaceCreationFailed",
		Action:              "CreateNamespace",
		Type:                events.EventTypeWarning,
		ReportingController: "controller",
		Message:             "Namespace creation failed.",
	},
	"NamespaceDeleted": {
		Name:                "NamespaceDeleted",
		Reason:              "NamespaceDeleted",
		Action:              "DeleteNamespace",
		Type:                events.EventTypeWarning,
		ReportingController: "controller",
		Message:             "Namespace got deleted.",
	},
	"NamespaceDeletionFailed": {
		Name:                "NamespaceDeletionFailed",
		Reason:              "NamespaceDeletionFailed",
		Action:              "DeleteNamespace",
		Type:                events.EventTypeWarning,
		ReportingController: "controller",
		Message:             "Namespace deletion failed.",
	},
	"WorkerClusterRoleCreated": {
		Name:                "WorkerClusterRoleCreated",
		Reason:              "WorkerClusterRoleCreated",
		Action:              "CreateWorkerClusterRole",
		Type:                events.EventTypeNormal,
		ReportingController: "controller",
		Message:             "Worker cluster role got created.",
	},
	"WorkerClusterRoleCreationFailed": {
		Name:                "WorkerClusterRoleCreationFailed",
		Reason:              "WorkerClusterRoleCreationFailed",
		Action:              "CreateWorkerClusterRole",
		Type:                events.EventTypeWarning,
		ReportingController: "controller",
		Message:             "Worker cluster role creation failed.",
	},
	"WorkerClusterRoleUpdated": {
		Name:                "WorkerClusterRoleUpdated",
		Reason:              "WorkerClusterRoleUpdated",
		Action:              "UpdateWorkerClusterRole",
		Type:                events.EventTypeNormal,
		ReportingController: "controller",
		Message:             "Worker cluster role got updated.",
	},
	"WorkerClusterRoleUpdateFailed": {
		Name:                "WorkerClusterRoleUpdateFailed",
		Reason:              "WorkerClusterRoleUpdateFailed",
		Action:              "UpdateWorkerClusterRole",
		Type:                events.EventTypeWarning,
		ReportingController: "controller",
		Message:             "Worker cluster role update failed.",
	},
	"ReadOnlyRoleCreated": {
		Name:                "ReadOnlyRoleCreated",
		Reason:              "ReadOnlyRoleCreated",
		Action:              "CreateReadOnlyRole",
		Type:                events.EventTypeNormal,
		ReportingController: "controller",
		Message:             "Read only role got created.",
	},
	"ReadOnlyRoleCreationFailed": {
		Name:                "ReadOnlyRoleCreationFailed",
		Reason:              "ReadOnlyRoleCreationFailed",
		Action:              "CreateReadOnlyRole",
		Type:                events.EventTypeWarning,
		ReportingController: "controller",
		Message:             "Read only role creation failed.",
	},
	"ReadOnlyRoleUpdated": {
		Name:                "ReadOnlyRoleUpdated",
		Reason:              "ReadOnlyRoleUpdated",
		Action:              "UpdateReadOnlyRole",
		Type:                events.EventTypeNormal,
		ReportingController: "controller",
		Message:             "Read only role got updated.",
	},
	"ReadOnlyRoleUpdateFailed": {
		Name:                "ReadOnlyRoleUpdateFailed",
		Reason:              "ReadOnlyRoleUpdateFailed",
		Action:              "UpdateReadOnlyRole",
		Type:                events.EventTypeWarning,
		ReportingController: "controller",
		Message:             "Read only role update failed.",
	},
	"ReadWriteRoleCreated": {
		Name:                "ReadWriteRoleCreated",
		Reason:              "ReadWriteRoleCreated",
		Action:              "CreateReadWriteRole",
		Type:                events.EventTypeNormal,
		ReportingController: "controller",
		Message:             "Read write role got created.",
	},
	"ReadWriteRoleCreationFailed": {
		Name:                "ReadWriteRoleCreationFailed",
		Reason:              "ReadWriteRoleCreationFailed",
		Action:              "CreateReadWriteRole",
		Type:                events.EventTypeWarning,
		ReportingController: "controller",
		Message:             "Read write role creation failed.",
	},
	"ReadWriteRoleUpdated": {
		Name:                "ReadWriteRoleUpdated",
		Reason:              "ReadWriteRoleUpdated",
		Action:              "UpdateReadWriteRole",
		Type:                events.EventTypeNormal,
		ReportingController: "controller",
		Message:             "Read write role got updated.",
	},
	"ReadWriteRoleUpdateFailed": {
		Name:                "ReadWriteRoleUpdateFailed",
		Reason:              "ReadWriteRoleUpdateFailed",
		Action:              "UpdateReadWriteRole",
		Type:                events.EventTypeWarning,
		ReportingController: "controller",
		Message:             "Read write role update failed.",
	},
	"ServiceAccountCreated": {
		Name:                "ServiceAccountCreated",
		Reason:              "ServiceAccountCreated",
		Action:              "CreateServiceAccount",
		Type:                events.EventTypeNormal,
		ReportingController: "controller",
		Message:             "Service account got created.",
	},
	"ServiceAccountCreationFailed": {
		Name:                "ServiceAccountCreationFailed",
		Reason:              "ServiceAccountCreationFailed",
		Action:              "CreateServiceAccount",
		Type:                events.EventTypeWarning,
		ReportingController: "controller",
		Message:             "Service account creation failed.",
	},
	"ServiceAccountSecretCreated": {
		Name:                "ServiceAccountSecretCreated",
		Reason:              "ServiceAccountSecretCreated",
		Action:              "CreateServiceAccountSecret",
		Type:                events.EventTypeNormal,
		ReportingController: "controller",
		Message:             "Service account secret got created.",
	},
	"ServiceAccountSecretCreationFailed": {
		Name:                "ServiceAccountSecretCreationFailed",
		Reason:              "ServiceAccountSecretCreationFailed",
		Action:              "CreateServiceAccountSecret",
		Type:                events.EventTypeWarning,
		ReportingController: "controller",
		Message:             "Service account secret creation failed.",
	},
	"DefaultRoleBindingCreated": {
		Name:                "DefaultRoleBindingCreated",
		Reason:              "DefaultRoleBindingCreated",
		Action:              "CreateDefaultRoleBinding",
		Type:                events.EventTypeNormal,
		ReportingController: "controller",
		Message:             "Default role binding got created.",
	},
	"DefaultRoleBindingCreationFailed": {
		Name:                "DefaultRoleBindingCreationFailed",
		Reason:              "DefaultRoleBindingCreationFailed",
		Action:              "CreateDefaultRoleBinding",
		Type:                events.EventTypeWarning,
		ReportingController: "controller",
		Message:             "Default role binding creation failed.",
	},
	"DefaultRoleBindingUpdated": {
		Name:                "DefaultRoleBindingUpdated",
		Reason:              "DefaultRoleBindingUpdated",
		Action:              "UpdateDefaultRoleBinding",
		Type:                events.EventTypeNormal,
		ReportingController: "controller",
		Message:             "Default role binding got updated.",
	},
	"DefaultRoleBindingUpdateFailed": {
		Name:                "DefaultRoleBindingUpdateFailed",
		Reason:              "DefaultRoleBindingUpdateFailed",
		Action:              "UpdateDefaultRoleBinding",
		Type:                events.EventTypeWarning,
		ReportingController: "controller",
		Message:             "Default role binding update failed.",
	},
	"DefaultRoleBindingDeleted": {
		Name:                "DefaultRoleBindingDeleted",
		Reason:              "DefaultRoleBindingDeleted",
		Action:              "DeleteDefaultRoleBinding",
		Type:                events.EventTypeWarning,
		ReportingController: "controller",
		Message:             "Default role binding got deleted.",
	},
	"DefaultRoleBindingDeletionFailed": {
		Name:                "DefaultRoleBindingDeletionFailed",
		Reason:              "DefaultRoleBindingDeletionFailed",
		Action:              "DeleteDefaultRoleBinding",
		Type:                events.EventTypeWarning,
		ReportingController: "controller",
		Message:             "Default role binding deletion failed.",
	},
	"InactiveRoleBindingDeleted": {
		Name:                "InactiveRoleBindingDeleted",
		Reason:              "InactiveRoleBindingDeleted",
		Action:              "DeleteInactiveRoleBinding",
		Type:                events.EventTypeWarning,
		ReportingController: "controller",
		Message:             "Inactive role binding got deleted.",
	},
	"InactiveRoleBindingDeletionFailed": {
		Name:                "InactiveRoleBindingDeletionFailed",
		Reason:              "InactiveRoleBindingDeletionFailed",
		Action:              "DeleteInactiveRoleBinding",
		Type:                events.EventTypeWarning,
		ReportingController: "controller",
		Message:             "Inactive role binding deletion failed.",
	},
	"InactiveServiceAccountDeleted": {
		Name:                "InactiveServiceAccountDeleted",
		Reason:              "InactiveServiceAccountDeleted",
		Action:              "DeleteInactiveServiceAccount",
		Type:                events.EventTypeWarning,
		ReportingController: "controller",
		Message:             "Inactive service account got deleted.",
	},
	"InactiveServiceAccountDeletionFailed": {
		Name:                "InactiveServiceAccountDeletionFailed",
		Reason:              "InactiveServiceAccountDeletionFailed",
		Action:              "DeleteInactiveServiceAccount",
		Type:                events.EventTypeWarning,
		ReportingController: "controller",
		Message:             "Inactive service account deletion failed.",
	},
	"ServiceAccountDeleted": {
		Name:                "ServiceAccountDeleted",
		Reason:              "ServiceAccountDeleted",
		Action:              "DeleteServiceAccount",
		Type:                events.EventTypeWarning,
		ReportingController: "controller",
		Message:             "Service account got deleted.",
	},
	"ServiceAccountDeletionFailed": {
		Name:                "ServiceAccountDeletionFailed",
		Reason:              "ServiceAccountDeletionFailed",
		Action:              "DeleteServiceAccount",
		Type:                events.EventTypeWarning,
		ReportingController: "controller",
		Message:             "Service account deletion failed.",
	},
	"WorkerServiceImportDeletedForcefully": {
		Name:                "WorkerServiceImportDeletedForcefully",
		Reason:              "WorkerServiceImportDeletedForcefully",
		Action:              "DeleteWorkerServiceImport",
		Type:                events.EventTypeWarning,
		ReportingController: "controller",
		Message:             "Worker service import got deleted forcefully.",
	},
	"WorkerServiceImportRecreationFailed": {
		Name:                "WorkerServiceImportRecreationFailed",
		Reason:              "WorkerServiceImportRecreationFailed",
		Action:              "CreateWorkerServiceImport",
		Type:                events.EventTypeWarning,
		ReportingController: "controller",
		Message:             "Worker service import recreation failed after forceful deletion.",
	},
	"WorkerServiceImportRecreated": {
		Name:                "WorkerServiceImportRecreated",
		Reason:              "WorkerServiceImportRecreated",
		Action:              "CreateWorkerServiceImport",
		Type:                events.EventTypeNormal,
		ReportingController: "controller",
		Message:             "Worker service import got recreated after forceful deletion.",
	},
	"WorkerServiceImportCreationFailed": {
		Name:                "WorkerServiceImportCreationFailed",
		Reason:              "WorkerServiceImportCreationFailed",
		Action:              "CreateWorkerServiceImport",
		Type:                events.EventTypeWarning,
		ReportingController: "controller",
		Message:             "Worker service import creation failed.",
	},
	"WorkerServiceImportCreated": {
		Name:                "WorkerServiceImportCreated",
		Reason:              "WorkerServiceImportCreated",
		Action:              "CreateWorkerServiceImport",
		Type:                events.EventTypeNormal,
		ReportingController: "controller",
		Message:             "Worker service import got created.",
	},
	"WorkerServiceImportUpdateFailed": {
		Name:                "WorkerServiceImportUpdateFailed",
		Reason:              "WorkerServiceImportUpdateFailed",
		Action:              "UpdateWorkerServiceImport",
		Type:                events.EventTypeWarning,
		ReportingController: "controller",
		Message:             "Worker service import update failed.",
	},
	"WorkerServiceImportUpdated": {
		Name:                "WorkerServiceImportUpdated",
		Reason:              "WorkerServiceImportUpdated",
		Action:              "UpdateWorkerServiceImport",
		Type:                events.EventTypeNormal,
		ReportingController: "controller",
		Message:             "Worker service import got updated.",
	},
	"WorkerServiceImportDeleted": {
		Name:                "WorkerServiceImportDeleted",
		Reason:              "WorkerServiceImportDeleted",
		Action:              "DeleteWorkerServiceImport",
		Type:                events.EventTypeWarning,
		ReportingController: "controller",
		Message:             "Worker service import got deleted.",
	},
	"WorkerServiceImportDeletionFailed": {
		Name:                "WorkerServiceImportDeletionFailed",
		Reason:              "WorkerServiceImportDeletionFailed",
		Action:              "DeleteWorkerServiceImport",
		Type:                events.EventTypeWarning,
		ReportingController: "controller",
		Message:             "Worker service import deletion failed.",
	},
	"WorkerSliceConfigDeletedForcefully": {
		Name:                "WorkerSliceConfigDeletedForcefully",
		Reason:              "WorkerSliceConfigDeletedForcefully",
		Action:              "DeleteWorkerSliceConfig",
		Type:                events.EventTypeWarning,
		ReportingController: "controller",
		Message:             "Worker slice config got deleted forcefully.",
	},
	"WorkerSliceConfigRecreationFailed": {
		Name:                "WorkerSliceConfigRecreationFailed",
		Reason:              "WorkerSliceConfigRecreationFailed",
		Action:              "CreateWorkerSliceConfig",
		Type:                events.EventTypeWarning,
		ReportingController: "controller",
		Message:             "Worker slice config recreation failed after forceful deletion.",
	},
	"WorkerSliceConfigRecreated": {
		Name:                "WorkerSliceConfigRecreated",
		Reason:              "WorkerSliceConfigRecreated",
		Action:              "CreateWorkerSliceConfig",
		Type:                events.EventTypeNormal,
		ReportingController: "controller",
		Message:             "Worker slice config got recreated after forceful deletion.",
	},
	"WorkerSliceConfigCreationFailed": {
		Name:                "WorkerSliceConfigCreationFailed",
		Reason:              "WorkerSliceConfigCreationFailed",
		Action:              "CreateWorkerSliceConfig",
		Type:                events.EventTypeWarning,
		ReportingController: "controller",
		Message:             "Worker slice config creation failed.",
	},
	"WorkerSliceConfigCreated": {
		Name:                "WorkerSliceConfigCreated",
		Reason:              "WorkerSliceConfigCreated",
		Action:              "CreateWorkerSliceConfig",
		Type:                events.EventTypeNormal,
		ReportingController: "controller",
		Message:             "Worker slice config got created.",
	},
	"WorkerSliceConfigUpdateFailed": {
		Name:                "WorkerSliceConfigUpdateFailed",
		Reason:              "WorkerSliceConfigUpdateFailed",
		Action:              "UpdateWorkerSliceConfig",
		Type:                events.EventTypeWarning,
		ReportingController: "controller",
		Message:             "Worker slice config update failed.",
	},
	"WorkerSliceConfigUpdated": {
		Name:                "WorkerSliceConfigUpdated",
		Reason:              "WorkerSliceConfigUpdated",
		Action:              "UpdateWorkerSliceConfig",
		Type:                events.EventTypeNormal,
		ReportingController: "controller",
		Message:             "Worker slice config got updated.",
	},
	"WorkerSliceConfigDeleted": {
		Name:                "WorkerSliceConfigDeleted",
		Reason:              "WorkerSliceConfigDeleted",
		Action:              "DeleteWorkerSliceConfig",
		Type:                events.EventTypeWarning,
		ReportingController: "controller",
		Message:             "Worker slice config got deleted.",
	},
	"WorkerSliceConfigDeletionFailed": {
		Name:                "WorkerSliceConfigDeletionFailed",
		Reason:              "WorkerSliceConfigDeletionFailed",
		Action:              "DeleteWorkerSliceConfig",
		Type:                events.EventTypeWarning,
		ReportingController: "controller",
		Message:             "Worker slice config deletion failed.",
	},
	"WorkerSliceGatewayDeletedForcefully": {
		Name:                "WorkerSliceGatewayDeletedForcefully",
		Reason:              "WorkerSliceGatewayDeletedForcefully",
		Action:              "DeleteWorkerSliceGateway",
		Type:                events.EventTypeWarning,
		ReportingController: "controller",
		Message:             "Worker slice gateway got deleted forcefully.",
	},
	"WorkerSliceGatewayRecreationFailed": {
		Name:                "WorkerSliceGatewayRecreationFailed",
		Reason:              "WorkerSliceGatewayRecreationFailed",
		Action:              "CreateWorkerSliceGateway",
		Type:                events.EventTypeWarning,
		ReportingController: "controller",
		Message:             "Worker slice gateway recreation failed after forceful deletion.",
	},
	"WorkerSliceGatewayRecreated": {
		Name:                "WorkerSliceGatewayRecreated",
		Reason:              "WorkerSliceGatewayRecreated",
		Action:              "CreateWorkerSliceGateway",
		Type:                events.EventTypeNormal,
		ReportingController: "controller",
		Message:             "Worker slice gateway got recreated after forceful deletion.",
	},
	"WorkerSliceGatewayDeletionFailed": {
		Name:                "WorkerSliceGatewayDeletionFailed",
		Reason:              "WorkerSliceGatewayDeletionFailed",
		Action:              "DeleteWorkerSliceGateway",
		Type:                events.EventTypeWarning,
		ReportingController: "controller",
		Message:             "Worker slice gateway deletion failed.",
	},
	"WorkerSliceGatewayDeleted": {
		Name:                "WorkerSliceGatewayDeleted",
		Reason:              "WorkerSliceGatewayDeleted",
		Action:              "DeleteWorkerSliceGateway",
		Type:                events.EventTypeWarning,
		ReportingController: "controller",
		Message:             "Worker slice gateway got deleted.",
	},
	"WorkerSliceGatewayCreationFailed": {
		Name:                "WorkerSliceGatewayCreationFailed",
		Reason:              "WorkerSliceGatewayCreationFailed",
		Action:              "CreateWorkerSliceGateway",
		Type:                events.EventTypeWarning,
		ReportingController: "controller",
		Message:             "Worker slice gateway creation failed.",
	},
	"WorkerSliceGatewayCreated": {
		Name:                "WorkerSliceGatewayCreated",
		Reason:              "WorkerSliceGatewayCreated",
		Action:              "CreateWorkerSliceGateway",
		Type:                events.EventTypeNormal,
		ReportingController: "controller",
		Message:             "Worker slice gateway got created.",
	},
	"SliceGatewayJobCreationFailed": {
		Name:                "SliceGatewayJobCreationFailed",
		Reason:              "SliceGatewayJobCreationFailed",
		Action:              "CreateSliceGatewayJob",
		Type:                events.EventTypeWarning,
		ReportingController: "controller",
		Message:             "Slice gateway job creation failed.",
	},
	"SliceGatewayJobCreated": {
		Name:                "SliceGatewayJobCreated",
		Reason:              "SliceGatewayJobCreated",
		Action:              "CreateSliceGatewayJob",
		Type:                events.EventTypeNormal,
		ReportingController: "controller",
		Message:             "Slice gateway job got created.",
	},
}

var (
	EventProjectDeleted                       events.EventName = "ProjectDeleted"
	EventProjectDeletionFailed                events.EventName = "ProjectDeletionFailed"
	EventClusterInstallationInProgress        events.EventName = "ClusterInstallationInProgress"
	EventClusterInstallationFailed            events.EventName = "ClusterInstallationFailed"
	EventClusterInstallationPending           events.EventName = "ClusterInstallationPending"
	EventClusterDeleted                       events.EventName = "ClusterDeleted"
	EventClusterDeletionFailed                events.EventName = "ClusterDeletionFailed"
	EventClusterDeregistrationInProgress      events.EventName = "ClusterDeregistrationInProgress"
	EventClusterDeregistered                  events.EventName = "ClusterDeregistered"
	EventClusterDeregisterTimeout             events.EventName = "ClusterDeregisterTimeout"
	EventClusterDeregisterFailed              events.EventName = "ClusterDeregisterFailed"
	EventSliceConfigDeleted                   events.EventName = "SliceConfigDeleted"
	EventSliceConfigDeletionFailed            events.EventName = "SliceConfigDeletionFailed"
	EventServiceExportConfigDeleted           events.EventName = "ServiceExportConfigDeleted"
	EventServiceExportConfigDeletionFailed    events.EventName = "ServiceExportConfigDeletionFailed"
	EventDefaultSliceQoSConfigCreated         events.EventName = "DefaultSliceQoSConfigCreated"
	EventSliceQoSConfigDeleted                events.EventName = "SliceQoSConfigDeleted"
	EventSliceQoSConfigDeletionFailed         events.EventName = "SliceQoSConfigDeletionFailed"
	EventSecretDeleted                        events.EventName = "SecretDeleted"
	EventSecretDeletionFailed                 events.EventName = "SecretDeletionFailed"
	EventNamespaceCreated                     events.EventName = "NamespaceCreated"
	EventNamespaceCreationFailed              events.EventName = "NamespaceCreationFailed"
	EventNamespaceDeleted                     events.EventName = "NamespaceDeleted"
	EventNamespaceDeletionFailed              events.EventName = "NamespaceDeletionFailed"
	EventWorkerClusterRoleCreated             events.EventName = "WorkerClusterRoleCreated"
	EventWorkerClusterRoleCreationFailed      events.EventName = "WorkerClusterRoleCreationFailed"
	EventWorkerClusterRoleUpdated             events.EventName = "WorkerClusterRoleUpdated"
	EventWorkerClusterRoleUpdateFailed        events.EventName = "WorkerClusterRoleUpdateFailed"
	EventReadOnlyRoleCreated                  events.EventName = "ReadOnlyRoleCreated"
	EventReadOnlyRoleCreationFailed           events.EventName = "ReadOnlyRoleCreationFailed"
	EventReadOnlyRoleUpdated                  events.EventName = "ReadOnlyRoleUpdated"
	EventReadOnlyRoleUpdateFailed             events.EventName = "ReadOnlyRoleUpdateFailed"
	EventReadWriteRoleCreated                 events.EventName = "ReadWriteRoleCreated"
	EventReadWriteRoleCreationFailed          events.EventName = "ReadWriteRoleCreationFailed"
	EventReadWriteRoleUpdated                 events.EventName = "ReadWriteRoleUpdated"
	EventReadWriteRoleUpdateFailed            events.EventName = "ReadWriteRoleUpdateFailed"
	EventServiceAccountCreated                events.EventName = "ServiceAccountCreated"
	EventServiceAccountCreationFailed         events.EventName = "ServiceAccountCreationFailed"
	EventServiceAccountSecretCreated          events.EventName = "ServiceAccountSecretCreated"
	EventServiceAccountSecretCreationFailed   events.EventName = "ServiceAccountSecretCreationFailed"
	EventDefaultRoleBindingCreated            events.EventName = "DefaultRoleBindingCreated"
	EventDefaultRoleBindingCreationFailed     events.EventName = "DefaultRoleBindingCreationFailed"
	EventDefaultRoleBindingUpdated            events.EventName = "DefaultRoleBindingUpdated"
	EventDefaultRoleBindingUpdateFailed       events.EventName = "DefaultRoleBindingUpdateFailed"
	EventDefaultRoleBindingDeleted            events.EventName = "DefaultRoleBindingDeleted"
	EventDefaultRoleBindingDeletionFailed     events.EventName = "DefaultRoleBindingDeletionFailed"
	EventInactiveRoleBindingDeleted           events.EventName = "InactiveRoleBindingDeleted"
	EventInactiveRoleBindingDeletionFailed    events.EventName = "InactiveRoleBindingDeletionFailed"
	EventInactiveServiceAccountDeleted        events.EventName = "InactiveServiceAccountDeleted"
	EventInactiveServiceAccountDeletionFailed events.EventName = "InactiveServiceAccountDeletionFailed"
	EventServiceAccountDeleted                events.EventName = "ServiceAccountDeleted"
	EventServiceAccountDeletionFailed         events.EventName = "ServiceAccountDeletionFailed"
	EventWorkerServiceImportDeletedForcefully events.EventName = "WorkerServiceImportDeletedForcefully"
	EventWorkerServiceImportRecreationFailed  events.EventName = "WorkerServiceImportRecreationFailed"
	EventWorkerServiceImportRecreated         events.EventName = "WorkerServiceImportRecreated"
	EventWorkerServiceImportCreationFailed    events.EventName = "WorkerServiceImportCreationFailed"
	EventWorkerServiceImportCreated           events.EventName = "WorkerServiceImportCreated"
	EventWorkerServiceImportUpdateFailed      events.EventName = "WorkerServiceImportUpdateFailed"
	EventWorkerServiceImportUpdated           events.EventName = "WorkerServiceImportUpdated"
	EventWorkerServiceImportDeleted           events.EventName = "WorkerServiceImportDeleted"
	EventWorkerServiceImportDeletionFailed    events.EventName = "WorkerServiceImportDeletionFailed"
	EventWorkerSliceConfigDeletedForcefully   events.EventName = "WorkerSliceConfigDeletedForcefully"
	EventWorkerSliceConfigRecreationFailed    events.EventName = "WorkerSliceConfigRecreationFailed"
	EventWorkerSliceConfigRecreated           events.EventName = "WorkerSliceConfigRecreated"
	EventWorkerSliceConfigCreationFailed      events.EventName = "WorkerSliceConfigCreationFailed"
	EventWorkerSliceConfigCreated             events.EventName = "WorkerSliceConfigCreated"
	EventWorkerSliceConfigUpdateFailed        events.EventName = "WorkerSliceConfigUpdateFailed"
	EventWorkerSliceConfigUpdated             events.EventName = "WorkerSliceConfigUpdated"
	EventWorkerSliceConfigDeleted             events.EventName = "WorkerSliceConfigDeleted"
	EventWorkerSliceConfigDeletionFailed      events.EventName = "WorkerSliceConfigDeletionFailed"
	EventWorkerSliceGatewayDeletedForcefully  events.EventName = "WorkerSliceGatewayDeletedForcefully"
	EventWorkerSliceGatewayRecreationFailed   events.EventName = "WorkerSliceGatewayRecreationFailed"
	EventWorkerSliceGatewayRecreated          events.EventName = "WorkerSliceGatewayRecreated"
	EventWorkerSliceGatewayDeletionFailed     events.EventName = "WorkerSliceGatewayDeletionFailed"
	EventWorkerSliceGatewayDeleted            events.EventName = "WorkerSliceGatewayDeleted"
	EventWorkerSliceGatewayCreationFailed     events.EventName = "WorkerSliceGatewayCreationFailed"
	EventWorkerSliceGatewayCreated            events.EventName = "WorkerSliceGatewayCreated"
	EventSliceGatewayJobCreationFailed        events.EventName = "SliceGatewayJobCreationFailed"
	EventSliceGatewayJobCreated               events.EventName = "SliceGatewayJobCreated"
)
