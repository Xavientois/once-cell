/*
This file was auto-generated. Do not modify it, lest your changes be overwritten.
*/

package oncecell

import (
	"testing"

	. "github.com/onsi/gomega"
)

var testInt64Value int64 = -64

func TestSetInt64Value(t *testing.T) {
	Default = NewGomegaWithT(t)
	var o OnceInt64

	err := o.SetValue(testInt64Value)
	Ω(err).ShouldNot(HaveOccurred())

	v, err := o.Value()
	Ω(err).ShouldNot(HaveOccurred())
	Ω(v).Should(Equal(testInt64Value))
}

func TestGetInt64Value(t *testing.T) {
	Default = NewGomegaWithT(t)
	var o OnceInt64

	o.SetValue(testInt64Value)

	v, err := o.Value()
	Ω(err).ShouldNot(HaveOccurred())
	Ω(v).Should(Equal(testInt64Value))
}

func TestMustSetInt64Value(t *testing.T) {
	Default = NewGomegaWithT(t)
	var o OnceInt64

	Ω(func() { o.MustSetValue(testInt64Value) }).ShouldNot(Panic())
}

func TestMustGetInt64Value(t *testing.T) {
	Default = NewGomegaWithT(t)
	var o OnceInt64

	err := o.SetValue(testInt64Value)
	Ω(err).ShouldNot(HaveOccurred())

	Ω(func() { o.MustValue() }).ShouldNot(Panic())
	Ω(o.MustValue()).Should(Equal(testInt64Value))
}

func TestMultiSetInt64Value(t *testing.T) {
	Default = NewGomegaWithT(t)
	var o OnceInt64

	err := o.SetValue(testInt64Value)
	Ω(err).ShouldNot(HaveOccurred())

	err = o.SetValue(testInt64Value)
	Ω(err).Should(HaveOccurred())
}

func TestMultiMustSetInt64Value(t *testing.T) {
	Default = NewGomegaWithT(t)
	var o OnceInt64

	Ω(func() { o.MustSetValue(testInt64Value) }).ShouldNot(Panic())
	Ω(func() { o.MustSetValue(testInt64Value) }).Should(Panic())
}

func TestSyncMultiSetInt64Value(t *testing.T) {
	Default = NewGomegaWithT(t)
	var o OnceInt64

	errorChannel := make(chan error)

	set := func() {
		err := o.SetValue(testInt64Value)
		if err != nil {
			errorChannel <- err
		}
	}

	go set()
	go set()

	Eventually(errorChannel).Should(Receive())
}
