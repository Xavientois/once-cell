/*
This file was auto-generated. Do not modify it, lest your changes be overwritten.
*/

package oncecell

import (
	"testing"

	. "github.com/onsi/gomega"
)

var testInt32Value int32 = -32

func TestSetInt32Value(t *testing.T) {
	Default = NewGomegaWithT(t)
	var o OnceInt32

	err := o.SetValue(testInt32Value)
	Ω(err).ShouldNot(HaveOccurred())

	v, err := o.Value()
	Ω(err).ShouldNot(HaveOccurred())
	Ω(v).Should(Equal(testInt32Value))
}

func TestGetInt32Value(t *testing.T) {
	Default = NewGomegaWithT(t)
	var o OnceInt32

	err := o.SetValue(testInt32Value)

	v, err := o.Value()
	Ω(err).ShouldNot(HaveOccurred())
	Ω(v).Should(Equal(testInt32Value))
}

func TestMustSetInt32Value(t *testing.T) {
	Default = NewGomegaWithT(t)
	var o OnceInt32

	Ω(func() { o.MustSetValue(testInt32Value) }).ShouldNot(Panic())
}

func TestMustGetInt32Value(t *testing.T) {
	Default = NewGomegaWithT(t)
	var o OnceInt32

	err := o.SetValue(testInt32Value)
	Ω(err).ShouldNot(HaveOccurred())

	Ω(func() { o.MustValue() }).ShouldNot(Panic())
	Ω(o.MustValue()).Should(Equal(testInt32Value))
}

func TestMultiSetInt32Value(t *testing.T) {
	Default = NewGomegaWithT(t)
	var o OnceInt32

	err := o.SetValue(testInt32Value)
	Ω(err).ShouldNot(HaveOccurred())

	err = o.SetValue(testInt32Value)
	Ω(err).Should(HaveOccurred())
}

func TestMultiMustSetInt32Value(t *testing.T) {
	Default = NewGomegaWithT(t)
	var o OnceInt32

	Ω(func() { o.MustSetValue(testInt32Value) }).ShouldNot(Panic())
	Ω(func() { o.MustSetValue(testInt32Value) }).Should(Panic())
}

func TestSyncMultiSetInt32Value(t *testing.T) {
	Default = NewGomegaWithT(t)
	var o OnceInt32

	errorChannel := make(chan error)

	set := func() {
		err := o.SetValue(testInt32Value)
		if err != nil {
			errorChannel <- err
		}
	}

	go set()
	go set()

	Eventually(errorChannel).Should(Receive())
}
