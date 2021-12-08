/*
This file was auto-generated. Do not modify it, lest your changes be overwritten.
*/

package oncecell

import (
	"testing"

	. "github.com/onsi/gomega"
)

var testComplex128Value complex128 = complex(1,28)

func TestSetComplex128Value(t *testing.T) {
	Default = NewGomegaWithT(t)
	var o OnceComplex128

	err := o.SetValue(testComplex128Value)
	Ω(err).ShouldNot(HaveOccurred())

	v, err := o.Value()
	Ω(err).ShouldNot(HaveOccurred())
	Ω(v).Should(Equal(testComplex128Value))
}

func TestGetComplex128Value(t *testing.T) {
	Default = NewGomegaWithT(t)
	var o OnceComplex128

	err := o.SetValue(testComplex128Value)

	v, err := o.Value()
	Ω(err).ShouldNot(HaveOccurred())
	Ω(v).Should(Equal(testComplex128Value))
}

func TestMustSetComplex128Value(t *testing.T) {
	Default = NewGomegaWithT(t)
	var o OnceComplex128

	Ω(func() { o.MustSetValue(testComplex128Value) }).ShouldNot(Panic())
}

func TestMustGetComplex128Value(t *testing.T) {
	Default = NewGomegaWithT(t)
	var o OnceComplex128

	err := o.SetValue(testComplex128Value)
	Ω(err).ShouldNot(HaveOccurred())

	Ω(func() { o.MustValue() }).ShouldNot(Panic())
	Ω(o.MustValue()).Should(Equal(testComplex128Value))
}

func TestMultiSetComplex128Value(t *testing.T) {
	Default = NewGomegaWithT(t)
	var o OnceComplex128

	err := o.SetValue(testComplex128Value)
	Ω(err).ShouldNot(HaveOccurred())

	err = o.SetValue(testComplex128Value)
	Ω(err).Should(HaveOccurred())
}

func TestMultiMustSetComplex128Value(t *testing.T) {
	Default = NewGomegaWithT(t)
	var o OnceComplex128

	Ω(func() { o.MustSetValue(testComplex128Value) }).ShouldNot(Panic())
	Ω(func() { o.MustSetValue(testComplex128Value) }).Should(Panic())
}

func TestSyncMultiSetComplex128Value(t *testing.T) {
	Default = NewGomegaWithT(t)
	var o OnceComplex128

	errorChannel := make(chan error)

	set := func() {
		err := o.SetValue(testComplex128Value)
		if err != nil {
			errorChannel <- err
		}
	}

	go set()
	go set()

	Eventually(errorChannel).Should(Receive())
}
