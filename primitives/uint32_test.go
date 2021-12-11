/*
This file was auto-generated. Do not modify it, lest your changes be overwritten.
*/

package oncecell

import (
	"testing"

	. "github.com/onsi/gomega"
)

var testUint32Value uint32 = 32

func TestSetUint32Value(t *testing.T) {
	Default = NewGomegaWithT(t)
	var o OnceUint32

	err := o.SetValue(testUint32Value)
	Ω(err).ShouldNot(HaveOccurred())

	v, err := o.Value()
	Ω(err).ShouldNot(HaveOccurred())
	Ω(v).Should(Equal(testUint32Value))
}

func TestGetUint32Value(t *testing.T) {
	Default = NewGomegaWithT(t)
	var o OnceUint32

	o.SetValue(testUint32Value)

	v, err := o.Value()
	Ω(err).ShouldNot(HaveOccurred())
	Ω(v).Should(Equal(testUint32Value))
}

func TestMustSetUint32Value(t *testing.T) {
	Default = NewGomegaWithT(t)
	var o OnceUint32

	Ω(func() { o.MustSetValue(testUint32Value) }).ShouldNot(Panic())
}

func TestMustGetUint32Value(t *testing.T) {
	Default = NewGomegaWithT(t)
	var o OnceUint32

	err := o.SetValue(testUint32Value)
	Ω(err).ShouldNot(HaveOccurred())

	Ω(func() { o.MustValue() }).ShouldNot(Panic())
	Ω(o.MustValue()).Should(Equal(testUint32Value))
}

func TestMultiSetUint32Value(t *testing.T) {
	Default = NewGomegaWithT(t)
	var o OnceUint32

	err := o.SetValue(testUint32Value)
	Ω(err).ShouldNot(HaveOccurred())

	err = o.SetValue(testUint32Value)
	Ω(err).Should(HaveOccurred())
}

func TestMultiMustSetUint32Value(t *testing.T) {
	Default = NewGomegaWithT(t)
	var o OnceUint32

	Ω(func() { o.MustSetValue(testUint32Value) }).ShouldNot(Panic())
	Ω(func() { o.MustSetValue(testUint32Value) }).Should(Panic())
}

func TestSyncMultiSetUint32Value(t *testing.T) {
	Default = NewGomegaWithT(t)
	var o OnceUint32

	errorChannel := make(chan error)

	set := func() {
		err := o.SetValue(testUint32Value)
		if err != nil {
			errorChannel <- err
		}
	}

	go set()
	go set()

	Eventually(errorChannel).Should(Receive())
}
