/*
This file was auto-generated. Do not modify it, lest your changes be overwritten.
*/

package oncecell

import (
	"testing"

	. "github.com/onsi/gomega"
)

var testUintValue uint = 4

func TestSetUintValue(t *testing.T) {
	Default = NewGomegaWithT(t)
	var o OnceUint

	err := o.SetValue(testUintValue)
	Ω(err).ShouldNot(HaveOccurred())

	v, err := o.Value()
	Ω(err).ShouldNot(HaveOccurred())
	Ω(v).Should(Equal(testUintValue))
}

func TestGetUintValue(t *testing.T) {
	Default = NewGomegaWithT(t)
	var o OnceUint

	o.SetValue(testUintValue)

	v, err := o.Value()
	Ω(err).ShouldNot(HaveOccurred())
	Ω(v).Should(Equal(testUintValue))
}

func TestMustSetUintValue(t *testing.T) {
	Default = NewGomegaWithT(t)
	var o OnceUint

	Ω(func() { o.MustSetValue(testUintValue) }).ShouldNot(Panic())
}

func TestMustGetUintValue(t *testing.T) {
	Default = NewGomegaWithT(t)
	var o OnceUint

	err := o.SetValue(testUintValue)
	Ω(err).ShouldNot(HaveOccurred())

	Ω(func() { o.MustValue() }).ShouldNot(Panic())
	Ω(o.MustValue()).Should(Equal(testUintValue))
}

func TestMultiSetUintValue(t *testing.T) {
	Default = NewGomegaWithT(t)
	var o OnceUint

	err := o.SetValue(testUintValue)
	Ω(err).ShouldNot(HaveOccurred())

	err = o.SetValue(testUintValue)
	Ω(err).Should(HaveOccurred())
}

func TestMultiMustSetUintValue(t *testing.T) {
	Default = NewGomegaWithT(t)
	var o OnceUint

	Ω(func() { o.MustSetValue(testUintValue) }).ShouldNot(Panic())
	Ω(func() { o.MustSetValue(testUintValue) }).Should(Panic())
}

func TestSyncMultiSetUintValue(t *testing.T) {
	Default = NewGomegaWithT(t)
	var o OnceUint

	errorChannel := make(chan error)

	set := func() {
		err := o.SetValue(testUintValue)
		if err != nil {
			errorChannel <- err
		}
	}

	go set()
	go set()

	Eventually(errorChannel).Should(Receive())
}
