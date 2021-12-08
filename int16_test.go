package oncecell

import (
	"testing"

	. "github.com/onsi/gomega"
)

var testInt16Value int16 = -16

func TestSetInt16Value(t *testing.T) {
	Default = NewGomegaWithT(t)
	var o OnceInt16

	err := o.SetValue(testInt16Value)
	Ω(err).ShouldNot(HaveOccurred())

	v, err := o.Value()
	Ω(err).ShouldNot(HaveOccurred())
	Ω(v).Should(Equal(testInt16Value))
}

func TestGetInt16Value(t *testing.T) {
	Default = NewGomegaWithT(t)
	var o OnceInt16

	err := o.SetValue(testInt16Value)

	v, err := o.Value()
	Ω(err).ShouldNot(HaveOccurred())
	Ω(v).Should(Equal(testInt16Value))
}

func TestMustSetInt16Value(t *testing.T) {
	Default = NewGomegaWithT(t)
	var o OnceInt16

	Ω(func() { o.MustSetValue(testInt16Value) }).ShouldNot(Panic())
}

func TestMustGetInt16Value(t *testing.T) {
	Default = NewGomegaWithT(t)
	var o OnceInt16

	err := o.SetValue(testInt16Value)
	Ω(err).ShouldNot(HaveOccurred())

	Ω(func() { o.MustValue() }).ShouldNot(Panic())
	Ω(o.MustValue()).Should(Equal(testInt16Value))
}

func TestMultiSetInt16Value(t *testing.T) {
	Default = NewGomegaWithT(t)
	var o OnceInt16

	err := o.SetValue(testInt16Value)
	Ω(err).ShouldNot(HaveOccurred())

	err = o.SetValue(testInt16Value)
	Ω(err).Should(HaveOccurred())
}

func TestMultiMustSetInt16Value(t *testing.T) {
	Default = NewGomegaWithT(t)
	var o OnceInt16

	Ω(func() { o.MustSetValue(testInt16Value) }).ShouldNot(Panic())
	Ω(func() { o.MustSetValue(testInt16Value) }).Should(Panic())
}

func TestSyncMultiSetInt16Value(t *testing.T) {
	Default = NewGomegaWithT(t)
	var o OnceInt16

	errorChannel := make(chan error)

	set := func() {
		err := o.SetValue(testInt16Value)
		if err != nil {
			errorChannel <- err
		}
	}

	go set()
	go set()

	Eventually(errorChannel).Should(Receive())
}
