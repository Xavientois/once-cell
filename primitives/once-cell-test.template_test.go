/*
This file was auto-generated. Do not modify it, lest your changes be overwritten.
*/

package oncecell

import (
	"testing"

	. "github.com/onsi/gomega"
)

var testMyTypeValue myType = myType{}

func TestSetMyTypeValue(t *testing.T) {
	Default = NewGomegaWithT(t)
	var o OnceMyType

	err := o.SetValue(testMyTypeValue)
	Ω(err).ShouldNot(HaveOccurred())

	v, err := o.Value()
	Ω(err).ShouldNot(HaveOccurred())
	Ω(v).Should(Equal(testMyTypeValue))
}

func TestGetMyTypeValue(t *testing.T) {
	Default = NewGomegaWithT(t)
	var o OnceMyType

	err := o.SetValue(testMyTypeValue)

	v, err := o.Value()
	Ω(err).ShouldNot(HaveOccurred())
	Ω(v).Should(Equal(testMyTypeValue))
}

func TestMustSetMyTypeValue(t *testing.T) {
	Default = NewGomegaWithT(t)
	var o OnceMyType

	Ω(func() { o.MustSetValue(testMyTypeValue) }).ShouldNot(Panic())
}

func TestMustGetMyTypeValue(t *testing.T) {
	Default = NewGomegaWithT(t)
	var o OnceMyType

	err := o.SetValue(testMyTypeValue)
	Ω(err).ShouldNot(HaveOccurred())

	Ω(func() { o.MustValue() }).ShouldNot(Panic())
	Ω(o.MustValue()).Should(Equal(testMyTypeValue))
}

func TestMultiSetMyTypeValue(t *testing.T) {
	Default = NewGomegaWithT(t)
	var o OnceMyType

	err := o.SetValue(testMyTypeValue)
	Ω(err).ShouldNot(HaveOccurred())

	err = o.SetValue(testMyTypeValue)
	Ω(err).Should(HaveOccurred())
}

func TestMultiMustSetMyTypeValue(t *testing.T) {
	Default = NewGomegaWithT(t)
	var o OnceMyType

	Ω(func() { o.MustSetValue(testMyTypeValue) }).ShouldNot(Panic())
	Ω(func() { o.MustSetValue(testMyTypeValue) }).Should(Panic())
}

func TestSyncMultiSetMyTypeValue(t *testing.T) {
	Default = NewGomegaWithT(t)
	var o OnceMyType

	errorChannel := make(chan error)

	set := func() {
		err := o.SetValue(testMyTypeValue)
		if err != nil {
			errorChannel <- err
		}
	}

	go set()
	go set()

	Eventually(errorChannel).Should(Receive())
}
