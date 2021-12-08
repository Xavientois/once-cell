/*
This file was auto-generated. Do not modify it, lest your changes be overwritten.
*/

package oncecell

import (
	"testing"

	. "github.com/onsi/gomega"
)

var testComplex64Value complex64 = complex(6,4)

func TestSetComplex64Value(t *testing.T) {
	Default = NewGomegaWithT(t)
	var o OnceComplex64

	err := o.SetValue(testComplex64Value)
	Ω(err).ShouldNot(HaveOccurred())

	v, err := o.Value()
	Ω(err).ShouldNot(HaveOccurred())
	Ω(v).Should(Equal(testComplex64Value))
}

func TestGetComplex64Value(t *testing.T) {
	Default = NewGomegaWithT(t)
	var o OnceComplex64

	err := o.SetValue(testComplex64Value)

	v, err := o.Value()
	Ω(err).ShouldNot(HaveOccurred())
	Ω(v).Should(Equal(testComplex64Value))
}

func TestMustSetComplex64Value(t *testing.T) {
	Default = NewGomegaWithT(t)
	var o OnceComplex64

	Ω(func() { o.MustSetValue(testComplex64Value) }).ShouldNot(Panic())
}

func TestMustGetComplex64Value(t *testing.T) {
	Default = NewGomegaWithT(t)
	var o OnceComplex64

	err := o.SetValue(testComplex64Value)
	Ω(err).ShouldNot(HaveOccurred())

	Ω(func() { o.MustValue() }).ShouldNot(Panic())
	Ω(o.MustValue()).Should(Equal(testComplex64Value))
}

func TestMultiSetComplex64Value(t *testing.T) {
	Default = NewGomegaWithT(t)
	var o OnceComplex64

	err := o.SetValue(testComplex64Value)
	Ω(err).ShouldNot(HaveOccurred())

	err = o.SetValue(testComplex64Value)
	Ω(err).Should(HaveOccurred())
}

func TestMultiMustSetComplex64Value(t *testing.T) {
	Default = NewGomegaWithT(t)
	var o OnceComplex64

	Ω(func() { o.MustSetValue(testComplex64Value) }).ShouldNot(Panic())
	Ω(func() { o.MustSetValue(testComplex64Value) }).Should(Panic())
}

func TestSyncMultiSetComplex64Value(t *testing.T) {
	Default = NewGomegaWithT(t)
	var o OnceComplex64

	errorChannel := make(chan error)

	set := func() {
		err := o.SetValue(testComplex64Value)
		if err != nil {
			errorChannel <- err
		}
	}

	go set()
	go set()

	Eventually(errorChannel).Should(Receive())
}
