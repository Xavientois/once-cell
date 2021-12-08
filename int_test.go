package oncecell

import (
	"testing"

	. "github.com/onsi/gomega"
)

var testIntValue int = -4

func TestSetIntValue(t *testing.T) {
	Default = NewGomegaWithT(t)
	var o OnceInt

	err := o.SetValue(testIntValue)
	Ω(err).ShouldNot(HaveOccurred())

	v, err := o.Value()
	Ω(err).ShouldNot(HaveOccurred())
	Ω(v).Should(Equal(testIntValue))
}

func TestGetIntValue(t *testing.T) {
	Default = NewGomegaWithT(t)
	var o OnceInt

	err := o.SetValue(testIntValue)

	v, err := o.Value()
	Ω(err).ShouldNot(HaveOccurred())
	Ω(v).Should(Equal(testIntValue))
}

func TestMustSetIntValue(t *testing.T) {
	Default = NewGomegaWithT(t)
	var o OnceInt

	Ω(func() { o.MustSetValue(testIntValue) }).ShouldNot(Panic())
}

func TestMustGetIntValue(t *testing.T) {
	Default = NewGomegaWithT(t)
	var o OnceInt

	err := o.SetValue(testIntValue)
	Ω(err).ShouldNot(HaveOccurred())

	Ω(func() { o.MustValue() }).ShouldNot(Panic())
	Ω(o.MustValue()).Should(Equal(testIntValue))
}

func TestMultiSetIntValue(t *testing.T) {
	Default = NewGomegaWithT(t)
	var o OnceInt

	err := o.SetValue(testIntValue)
	Ω(err).ShouldNot(HaveOccurred())

	err = o.SetValue(testIntValue)
	Ω(err).Should(HaveOccurred())
}

func TestMultiMustSetIntValue(t *testing.T) {
	Default = NewGomegaWithT(t)
	var o OnceInt

	Ω(func() { o.MustSetValue(testIntValue) }).ShouldNot(Panic())
	Ω(func() { o.MustSetValue(testIntValue) }).Should(Panic())
}

func TestSyncMultiSetIntValue(t *testing.T) {
	Default = NewGomegaWithT(t)
	var o OnceInt

	errorChannel := make(chan error)

	set := func() {
		err := o.SetValue(testIntValue)
		if err != nil {
			errorChannel <- err
		}
	}

	go set()
	go set()

	Eventually(errorChannel).Should(Receive())
}
