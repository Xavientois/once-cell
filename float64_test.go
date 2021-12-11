/*
This file was auto-generated. Do not modify it, lest your changes be overwritten.
*/

package oncecell

import (
	"testing"

	. "github.com/onsi/gomega"
)

var testFloat64Value float64 = 0.64

func TestSetFloat64Value(t *testing.T) {
	Default = NewGomegaWithT(t)
	var o OnceFloat64

	err := o.SetValue(testFloat64Value)
	Ω(err).ShouldNot(HaveOccurred())

	v, err := o.Value()
	Ω(err).ShouldNot(HaveOccurred())
	Ω(v).Should(Equal(testFloat64Value))
}

func TestGetFloat64Value(t *testing.T) {
	Default = NewGomegaWithT(t)
	var o OnceFloat64

	o.SetValue(testFloat64Value)

	v, err := o.Value()
	Ω(err).ShouldNot(HaveOccurred())
	Ω(v).Should(Equal(testFloat64Value))
}

func TestMustSetFloat64Value(t *testing.T) {
	Default = NewGomegaWithT(t)
	var o OnceFloat64

	Ω(func() { o.MustSetValue(testFloat64Value) }).ShouldNot(Panic())
}

func TestMustGetFloat64Value(t *testing.T) {
	Default = NewGomegaWithT(t)
	var o OnceFloat64

	err := o.SetValue(testFloat64Value)
	Ω(err).ShouldNot(HaveOccurred())

	Ω(func() { o.MustValue() }).ShouldNot(Panic())
	Ω(o.MustValue()).Should(Equal(testFloat64Value))
}

func TestMultiSetFloat64Value(t *testing.T) {
	Default = NewGomegaWithT(t)
	var o OnceFloat64

	err := o.SetValue(testFloat64Value)
	Ω(err).ShouldNot(HaveOccurred())

	err = o.SetValue(testFloat64Value)
	Ω(err).Should(HaveOccurred())
}

func TestMultiMustSetFloat64Value(t *testing.T) {
	Default = NewGomegaWithT(t)
	var o OnceFloat64

	Ω(func() { o.MustSetValue(testFloat64Value) }).ShouldNot(Panic())
	Ω(func() { o.MustSetValue(testFloat64Value) }).Should(Panic())
}

func TestSyncMultiSetFloat64Value(t *testing.T) {
	Default = NewGomegaWithT(t)
	var o OnceFloat64

	errorChannel := make(chan error)

	set := func() {
		err := o.SetValue(testFloat64Value)
		if err != nil {
			errorChannel <- err
		}
	}

	go set()
	go set()

	Eventually(errorChannel).Should(Receive())
}
