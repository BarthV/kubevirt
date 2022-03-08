// Code generated by swagger-doc. DO NOT EDIT.

package v1alpha1

func (VirtualMachineClone) SwaggerDoc() map[string]string {
	return map[string]string{
		"":       "VirtualMachineClone TODO EDIT ihol3\n\n+k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object\n+k8s:openapi-gen=true\n+genclient\n+genclient:nonNamespaced",
		"status": "+optional",
	}
}

func (CloneIncludeExclude) SwaggerDoc() map[string]string {
	return map[string]string{
		"cloneAllByDefault": "+optional",
		"exclude":           "+optional",
		"include":           "+optional",
	}
}

func (VirtualMachineCloneSpec) SwaggerDoc() map[string]string {
	return map[string]string{
		"target":      "+optional",
		"annotations": "+optional",
	}
}

func (VirtualMachineCloneStatus) SwaggerDoc() map[string]string {
	return map[string]string{}
}

func (VirtualMachineCloneList) SwaggerDoc() map[string]string {
	return map[string]string{
		"":      "VirtualMachineCloneList is a list of MigrationPolicy\n\n+k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object",
		"items": "+listType=atomic",
	}
}
