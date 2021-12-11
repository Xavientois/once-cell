/*
This file was auto-generated. Do not modify it, lest your changes be overwritten.
*/

package oncecell

import (
	"testing"

	. "github.com/onsi/gomega"
)

var testUint16Value uint16 = 16

func TestSetUint16Value(t *testing.T) {
	Default = NewGomegaWithT(t)
	var o OnceUint16

	err := o.SetValue(testUint16Value)
	Ω(err).ShouldNot(HaveOccurred())

	v, err := o.Value()
	Ω(err).ShouldNot(HaveOccurred())
	Ω(v).Should(Equal(testUint16Value))
}

func TestGetUint16Value(t *testing.T) {
	Default = NewGomegaWithT(t)
	var o OnceUint16

	o.SetValue(testUint16Value)

	v, err := o.Value()
	Ω(err).ShouldNot(HaveOccurred())
	Ω(v).Should(Equal(testUint16Value))
}

func TestMustSetUint16Value(t *testing.T) {
	Default = NewGomegaWithT(t)
	var o OnceUint16

	Ω(func() { o.MustSetValue(testUint16Value) }).ShouldNot(Panic())
}

func TestMustGetUint16Value(t *testing.T) {
	Default = NewGomegaWithT(t)
	var o OnceUint16

	err := o.SetValue(testUint16Value)
	Ω(err).ShouldNot(HaveOccurred())

	Ω(func() { o.MustValue() }).ShouldNot(Panic())
	Ω(o.MustValue()).Should(Equal(testUint16Value))
}

func TestMultiSetUint16Value(t *testing.T) {
	Default = NewGomegaWithT(t)
	var o OnceUint16

	err := o.SetValue(testUint16Value)
	Ω(err).ShouldNot(HaveOccurred())

	err = o.SetValue(testUint16Value)
	Ω(err).Should(HaveOccurred())
}

func TestMultiMustSetUint16Value(t *testing.T) {
	Default = NewGomegaWithT(t)
	var o OnceUint16

	Ω(func() { o.MustSetValue(testUint16Value) }).ShouldNot(Panic())
	Ω(func() { o.MustSetValue(testUint16Value) }).Should(Panic())
}

func TestSyncMultiSetUint16Value(t *testing.T) {
	Default = NewGomegaWithT(t)
	var o OnceUint16

	errorChannel := make(chan error)

	set := func() {
		err := o.SetValue(testUint16Value)
		if err != nil {
			errorChannel <- err
		}
	}

	go set()
	go set()

	Eventually(errorChannel).Should(Receive())
}
