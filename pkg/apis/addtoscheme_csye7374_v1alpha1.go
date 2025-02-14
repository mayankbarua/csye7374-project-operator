package apis

import (
	"github.com/csye7374-Advance-Cloud/csye7374-project-operator/pkg/apis/csye7374/v1alpha1"
)

func init() {
	// Register the types with the Scheme so the components can map objects to GroupVersionKinds and back
	AddToSchemes = append(AddToSchemes, v1alpha1.SchemeBuilder.AddToScheme)
}
