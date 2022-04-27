// Code generated by swagger-doc. DO NOT EDIT.

package v1alpha1

func (VirtualMachineSnapshot) SwaggerDoc() map[string]string {
	return map[string]string{
		"":       "VirtualMachineSnapshot defines the operation of snapshotting a VM\n+genclient\n+k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object",
		"status": "+optional",
	}
}

func (VirtualMachineSnapshotSpec) SwaggerDoc() map[string]string {
	return map[string]string{
		"":                "VirtualMachineSnapshotSpec is the spec for a VirtualMachineSnapshot resource",
		"deletionPolicy":  "+optional",
		"failureDeadline": "This time represents the number of seconds we permit the vm snapshot\nto take. In case we pass this deadline we mark this snapshot\nas failed.\nDefaults to DefaultFailureDeadline - 5min\n+optional",
	}
}

func (VirtualMachineSnapshotStatus) SwaggerDoc() map[string]string {
	return map[string]string{
		"":                                  "VirtualMachineSnapshotStatus is the status for a VirtualMachineSnapshot resource",
		"sourceUID":                         "+optional",
		"virtualMachineSnapshotContentName": "+optional",
		"creationTime":                      "+optional\n+nullable",
		"phase":                             "+optional",
		"readyToUse":                        "+optional",
		"error":                             "+optional",
		"conditions":                        "+optional",
		"indications":                       "+optional\n+listType=set",
	}
}

func (Error) SwaggerDoc() map[string]string {
	return map[string]string{
		"":        "Error is the last error encountered during the snapshot/restore",
		"time":    "+optional",
		"message": "+optional",
	}
}

func (Condition) SwaggerDoc() map[string]string {
	return map[string]string{
		"":                   "Condition defines conditions",
		"lastProbeTime":      "+optional\n+nullable",
		"lastTransitionTime": "+optional\n+nullable",
		"reason":             "+optional",
		"message":            "+optional",
	}
}

func (VirtualMachineSnapshotList) SwaggerDoc() map[string]string {
	return map[string]string{
		"": "VirtualMachineSnapshotList is a list of VirtualMachineSnapshot resources\n+k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object",
	}
}

func (VirtualMachineSnapshotContent) SwaggerDoc() map[string]string {
	return map[string]string{
		"":       "VirtualMachineSnapshotContent contains the snapshot data\n+genclient\n+k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object",
		"status": "+optional",
	}
}

func (VirtualMachineSnapshotContentSpec) SwaggerDoc() map[string]string {
	return map[string]string{
		"":              "VirtualMachineSnapshotContentSpec is the spec for a VirtualMachineSnapshotContent resource",
		"volumeBackups": "+optional",
	}
}

func (SourceSpec) SwaggerDoc() map[string]string {
	return map[string]string{
		"":               "SourceSpec contains the appropriate spec for the resource being snapshotted",
		"virtualMachine": "+optional",
	}
}

func (PersistentVolumeClaim) SwaggerDoc() map[string]string {
	return map[string]string{
		"spec": "Spec defines the desired characteristics of a volume requested by a pod author.\nMore info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#persistentvolumeclaims\n+optional",
	}
}

func (VolumeBackup) SwaggerDoc() map[string]string {
	return map[string]string{
		"":                   "VolumeBackup contains the data neeed to restore a PVC",
		"volumeSnapshotName": "+optional",
	}
}

func (VirtualMachineSnapshotContentStatus) SwaggerDoc() map[string]string {
	return map[string]string{
		"":                     "VirtualMachineSnapshotContentStatus is the status for a VirtualMachineSnapshotStatus resource",
		"creationTime":         "+optional\n+nullable",
		"readyToUse":           "+optional",
		"error":                "+optional",
		"volumeSnapshotStatus": "+optional",
	}
}

func (VirtualMachineSnapshotContentList) SwaggerDoc() map[string]string {
	return map[string]string{
		"": "VirtualMachineSnapshotContentList is a list of VirtualMachineSnapshot resources\n+k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object",
	}
}

func (VolumeSnapshotStatus) SwaggerDoc() map[string]string {
	return map[string]string{
		"":             "VolumeSnapshotStatus is the status of a VolumeSnapshot",
		"creationTime": "+optional\n+nullable",
		"readyToUse":   "+optional",
		"error":        "+optional",
	}
}

func (VirtualMachineRestore) SwaggerDoc() map[string]string {
	return map[string]string{
		"":       "VirtualMachineRestore defines the operation of restoring a VM\n+genclient\n+k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object",
		"status": "+optional",
	}
}

func (VirtualMachineRestoreSpec) SwaggerDoc() map[string]string {
	return map[string]string{
		"":        "VirtualMachineRestoreSpec is the spec for a VirtualMachineRestoreresource",
		"target":  "initially only VirtualMachine type supported",
		"patches": "If the target for the restore does not exist, it will be created. Patches holds JSON patches that would be\napplied to the target manifest before it's created. Patches should fit the target's Kind.\n\nExample for a patch: {\"op\": \"replace\", \"path\": \"/metadata/name\", \"value\": \"new-vm-name\"}\n\n+optional\n+listType=atomic",
	}
}

func (VirtualMachineRestoreStatus) SwaggerDoc() map[string]string {
	return map[string]string{
		"":                   "VirtualMachineRestoreStatus is the spec for a VirtualMachineRestoreresource",
		"restores":           "+optional",
		"restoreTime":        "+optional",
		"deletedDataVolumes": "+optional",
		"complete":           "+optional",
		"conditions":         "+optional",
	}
}

func (VolumeRestore) SwaggerDoc() map[string]string {
	return map[string]string{
		"":               "VolumeRestore contains the data neeed to restore a PVC",
		"dataVolumeName": "+optional",
	}
}

func (VirtualMachineRestoreList) SwaggerDoc() map[string]string {
	return map[string]string{
		"": "VirtualMachineRestoreList is a list of VirtualMachineRestore resources\n+k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object",
	}
}
