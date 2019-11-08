package target_funcs_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestTargetFuncs(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "TargetFuncs Suite")
}

var _ = BeforeSuite(func() {
	println("BeforeSuite")
})

var _ = AfterSuite(func() {
	println("AfterSuite")
})
