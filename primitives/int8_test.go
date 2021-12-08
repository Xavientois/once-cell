/*
This file was auto-generated. Do not modify it, lest your changes be overwritten.
*/

package oncecell

import (
	"testing"

	. "github.com/onsi/gomega"
)

var testInt8Value int8 = -8

func TestSetInt8Value(t *testing.T) {
	Default = NewGomegaWithT(t)
	var o OnceInt8

	err := o.SetValue(testInt8Value)
	Ω(err).ShouldNot(HaveOccurred())

	v, err := o.Value()
	Ω(err).ShouldNot(HaveOccurred())
	Ω(v).Should(Equal(testInt8Value))
}

func TestGetInt8Value(t *testing.T) {
	Default = NewGomegaWithT(t)
	var o OnceInt8

	err := o.SetValue(testInt8Value)

	v, err := o.Value()
	Ω(err).ShouldNot(HaveOccurred())
	Ω(v).Should(Equal(testInt8Value))
}

func TestMustSetInt8Value(t *testing.T) {
	Default = NewGomegaWithT(t)
	var o OnceInt8

	Ω(func() { o.MustSetValue(testInt8Value) }).ShouldNot(Panic())
}

func TestMustGetInt8Value(t *testing.T) {
	Default = NewGomegaWithT(t)
	var o OnceInt8

	err := o.SetValue(testInt8Value)
	Ω(err).ShouldNot(HaveOccurred())

	Ω(func() { o.MustValue() }).ShouldNot(Panic())
	Ω(o.MustValue()).Should(Equal(testInt8Value))
}

func TestMultiSetInt8Value(t *testing.T) {
	Default = NewGomegaWithT(t)
	var o OnceInt8

	err := o.SetValue(testInt8Value)
	Ω(err).ShouldNot(HaveOccurred())

	err = o.SetValue(testInt8Value)
	Ω(err).Should(HaveOccurred())
}

func TestMultiMustSetInt8Value(t *testing.T) {
	Default = NewGomegaWithT(t)
	var o OnceInt8

	Ω(func() { o.MustSetValue(testInt8Value) }).ShouldNot(Panic())
	Ω(func() { o.MustSetValue(testInt8Value) }).Should(Panic())
}

func TestSyncMultiSetInt8Value(t *testing.T) {
	Default = NewGomegaWithT(t)
	var o OnceInt8

	errorChannel := make(chan error)

	set := func() {
		err := o.SetValue(testInt8Value)
		if err != nil {
			errorChannel <- err
		}
	}

	go set()
	go set()

	Eventually(errorChannel).Should(Receive())
}
