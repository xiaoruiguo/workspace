package apis

import (
	v1 "github.com/lijiaocn/workspace/studys/study-k8s-develop/code-generator/example-2/xx/test/pkg/apis/test/v1"
)

func init() {
	// Register the types with the Scheme so the components can map objects to GroupVersionKinds and back
	AddToSchemes = append(AddToSchemes, v1.SchemeBuilder.AddToScheme)
}
