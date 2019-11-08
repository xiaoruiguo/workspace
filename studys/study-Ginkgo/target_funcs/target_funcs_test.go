package target_funcs_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "studyginkgo/target_funcs"
)

var _ = Describe("TargetFuncs", func() {

	BeforeEach(func() {
		println("BeforeEach-2")
	})

	BeforeEach(func() {
		println("BeforeEach-1")
	})

	JustBeforeEach(func() {
		println("JustBeforeEach-1")
	})

	JustBeforeEach(func() {
		println("JustBeforeEach-2")
	})

	JustAfterEach(func() {
		println("JustAfterEach-1")
	})

	JustAfterEach(func() {
		println("JustAfterEach-2")
	})

	AfterEach(func() {
		println("AfterEach-1")
	})

	AfterEach(func() {
		println("AfterEach-2")
	})

	Describe("ReturnInt", func() {
		Context("default", func() {

			var (
				input  int
				result int
			)

			BeforeEach(func() {
				println("BeforeEach in Context")
				input = 1
				result = 1
			})

			AfterEach(func() {
				println("AfterEach in Context")
				input = 0
				result = 0
			})

			It("return value", func() {
				println("Exec Test Case")
				v := ReturnInt(input)
				Expect(v).To(Equal(result))
			})
		})
	})
})
