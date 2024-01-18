// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

/*
Copyright 2022 Upbound Inc.
Copyright 2023 Jakob Schlagenhaufer, Jan Dittrich
*/

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type VolumeAttachV3InitParameters struct {

	// Specify whether to attach the volume as Read-Only
	// (ro) or Read-Write (rw). Only values of ro and rw are accepted.
	// If left unspecified, the Block Storage API will apply a default of rw.
	AttachMode *string `json:"attachMode,omitempty" tf:"attach_mode,omitempty"`

	// The device to tell the Block Storage service this
	// volume will be attached as. This is purely for informational purposes.
	// You can specify auto or a device such as /dev/vdc.
	Device *string `json:"device,omitempty" tf:"device,omitempty"`

	// The host to attach the volume to.
	HostName *string `json:"hostName,omitempty" tf:"host_name,omitempty"`

	// The IP address of the host_name above.
	IPAddress *string `json:"ipAddress,omitempty" tf:"ip_address,omitempty"`

	// The iSCSI initiator string to make the connection.
	Initiator *string `json:"initiator,omitempty" tf:"initiator,omitempty"`

	// Whether to connect to this volume via multipath.
	Multipath *bool `json:"multipath,omitempty" tf:"multipath,omitempty"`

	// The iSCSI initiator OS type.
	OsType *string `json:"osType,omitempty" tf:"os_type,omitempty"`

	// The iSCSI initiator platform.
	Platform *string `json:"platform,omitempty" tf:"platform,omitempty"`

	// The region in which to obtain the V3 Block Storage
	// client. A Block Storage client is needed to create a volume attachment.
	// If omitted, the region argument of the provider is used. Changing this
	// creates a new volume attachment.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// The ID of the Volume to attach to an Instance.
	VolumeID *string `json:"volumeId,omitempty" tf:"volume_id,omitempty"`

	// A wwnn name. Used for Fibre Channel connections.
	Wwnn *string `json:"wwnn,omitempty" tf:"wwnn,omitempty"`

	// An array of wwpn strings. Used for Fibre Channel
	// connections.
	Wwpn []*string `json:"wwpn,omitempty" tf:"wwpn,omitempty"`
}

type VolumeAttachV3Observation struct {

	// Specify whether to attach the volume as Read-Only
	// (ro) or Read-Write (rw). Only values of ro and rw are accepted.
	// If left unspecified, the Block Storage API will apply a default of rw.
	AttachMode *string `json:"attachMode,omitempty" tf:"attach_mode,omitempty"`

	// The device to tell the Block Storage service this
	// volume will be attached as. This is purely for informational purposes.
	// You can specify auto or a device such as /dev/vdc.
	Device *string `json:"device,omitempty" tf:"device,omitempty"`

	// The storage driver that the volume is based on.
	DriverVolumeType *string `json:"driverVolumeType,omitempty" tf:"driver_volume_type,omitempty"`

	// The host to attach the volume to.
	HostName *string `json:"hostName,omitempty" tf:"host_name,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The IP address of the host_name above.
	IPAddress *string `json:"ipAddress,omitempty" tf:"ip_address,omitempty"`

	// The iSCSI initiator string to make the connection.
	Initiator *string `json:"initiator,omitempty" tf:"initiator,omitempty"`

	// A mount point base name for shared storage.
	MountPointBase *string `json:"mountPointBase,omitempty" tf:"mount_point_base,omitempty"`

	// Whether to connect to this volume via multipath.
	Multipath *bool `json:"multipath,omitempty" tf:"multipath,omitempty"`

	// The iSCSI initiator OS type.
	OsType *string `json:"osType,omitempty" tf:"os_type,omitempty"`

	// The iSCSI initiator platform.
	Platform *string `json:"platform,omitempty" tf:"platform,omitempty"`

	// The region in which to obtain the V3 Block Storage
	// client. A Block Storage client is needed to create a volume attachment.
	// If omitted, the region argument of the provider is used. Changing this
	// creates a new volume attachment.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// The ID of the Volume to attach to an Instance.
	VolumeID *string `json:"volumeId,omitempty" tf:"volume_id,omitempty"`

	// A wwnn name. Used for Fibre Channel connections.
	Wwnn *string `json:"wwnn,omitempty" tf:"wwnn,omitempty"`

	// An array of wwpn strings. Used for Fibre Channel
	// connections.
	Wwpn []*string `json:"wwpn,omitempty" tf:"wwpn,omitempty"`
}

type VolumeAttachV3Parameters struct {

	// Specify whether to attach the volume as Read-Only
	// (ro) or Read-Write (rw). Only values of ro and rw are accepted.
	// If left unspecified, the Block Storage API will apply a default of rw.
	// +kubebuilder:validation:Optional
	AttachMode *string `json:"attachMode,omitempty" tf:"attach_mode,omitempty"`

	// The device to tell the Block Storage service this
	// volume will be attached as. This is purely for informational purposes.
	// You can specify auto or a device such as /dev/vdc.
	// +kubebuilder:validation:Optional
	Device *string `json:"device,omitempty" tf:"device,omitempty"`

	// The host to attach the volume to.
	// +kubebuilder:validation:Optional
	HostName *string `json:"hostName,omitempty" tf:"host_name,omitempty"`

	// The IP address of the host_name above.
	// +kubebuilder:validation:Optional
	IPAddress *string `json:"ipAddress,omitempty" tf:"ip_address,omitempty"`

	// The iSCSI initiator string to make the connection.
	// +kubebuilder:validation:Optional
	Initiator *string `json:"initiator,omitempty" tf:"initiator,omitempty"`

	// Whether to connect to this volume via multipath.
	// +kubebuilder:validation:Optional
	Multipath *bool `json:"multipath,omitempty" tf:"multipath,omitempty"`

	// The iSCSI initiator OS type.
	// +kubebuilder:validation:Optional
	OsType *string `json:"osType,omitempty" tf:"os_type,omitempty"`

	// The iSCSI initiator platform.
	// +kubebuilder:validation:Optional
	Platform *string `json:"platform,omitempty" tf:"platform,omitempty"`

	// The region in which to obtain the V3 Block Storage
	// client. A Block Storage client is needed to create a volume attachment.
	// If omitted, the region argument of the provider is used. Changing this
	// creates a new volume attachment.
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// The ID of the Volume to attach to an Instance.
	// +kubebuilder:validation:Optional
	VolumeID *string `json:"volumeId,omitempty" tf:"volume_id,omitempty"`

	// A wwnn name. Used for Fibre Channel connections.
	// +kubebuilder:validation:Optional
	Wwnn *string `json:"wwnn,omitempty" tf:"wwnn,omitempty"`

	// An array of wwpn strings. Used for Fibre Channel
	// connections.
	// +kubebuilder:validation:Optional
	Wwpn []*string `json:"wwpn,omitempty" tf:"wwpn,omitempty"`
}

// VolumeAttachV3Spec defines the desired state of VolumeAttachV3
type VolumeAttachV3Spec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     VolumeAttachV3Parameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider VolumeAttachV3InitParameters `json:"initProvider,omitempty"`
}

// VolumeAttachV3Status defines the observed state of VolumeAttachV3.
type VolumeAttachV3Status struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        VolumeAttachV3Observation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// VolumeAttachV3 is the Schema for the VolumeAttachV3s API. Creates an attachment connection to a Block Storage volume
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,openstack}
type VolumeAttachV3 struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.hostName) || (has(self.initProvider) && has(self.initProvider.hostName))",message="spec.forProvider.hostName is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.volumeId) || (has(self.initProvider) && has(self.initProvider.volumeId))",message="spec.forProvider.volumeId is a required parameter"
	Spec   VolumeAttachV3Spec   `json:"spec"`
	Status VolumeAttachV3Status `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VolumeAttachV3List contains a list of VolumeAttachV3s
type VolumeAttachV3List struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VolumeAttachV3 `json:"items"`
}

// Repository type metadata.
var (
	VolumeAttachV3_Kind             = "VolumeAttachV3"
	VolumeAttachV3_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: VolumeAttachV3_Kind}.String()
	VolumeAttachV3_KindAPIVersion   = VolumeAttachV3_Kind + "." + CRDGroupVersion.String()
	VolumeAttachV3_GroupVersionKind = CRDGroupVersion.WithKind(VolumeAttachV3_Kind)
)

func init() {
	SchemeBuilder.Register(&VolumeAttachV3{}, &VolumeAttachV3List{})
}