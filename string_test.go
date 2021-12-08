package oncecell

import (
	"testing"

	. "github.com/onsi/gomega"
)

var testStringValue string = "test_value"

func TestSetStringValue(t *testing.T) {
	Default = NewGomegaWithT(t)
	var o OnceString

	err := o.SetValue(testStringValue)
	Ω(err).ShouldNot(HaveOccurred())

	v, err := o.Value()
	Ω(err).ShouldNot(HaveOccurred())
	Ω(v).Should(Equal(testStringValue))
}

func TestGetStringValue(t *testing.T) {
	Default = NewGomegaWithT(t)
	var o OnceString

	err := o.SetValue(testStringValue)

	v, err := o.Value()
	Ω(err).ShouldNot(HaveOccurred())
	Ω(v).Should(Equal(testStringValue))
}

func TestMustSetStringValue(t *testing.T) {
	Default = NewGomegaWithT(t)
	var o OnceString

	Ω(func() { o.MustSetValue(testStringValue) }).ShouldNot(Panic())
}

func TestMustGetStringValue(t *testing.T) {
	Default = NewGomegaWithT(t)
	var o OnceString

	err := o.SetValue(testStringValue)
	Ω(err).ShouldNot(HaveOccurred())

	Ω(func() { o.MustValue() }).ShouldNot(Panic())
	Ω(o.MustValue()).Should(Equal(testStringValue))
}

func TestMultiSetStringValue(t *testing.T) {
	Default = NewGomegaWithT(t)
	var o OnceString

	err := o.SetValue(testStringValue)
	Ω(err).ShouldNot(HaveOccurred())

	err = o.SetValue(testStringValue)
	Ω(err).Should(HaveOccurred())
}

func TestMultiMustSetStringValue(t *testing.T) {
	Default = NewGomegaWithT(t)
	var o OnceString

	Ω(func() { o.MustSetValue(testStringValue) }).ShouldNot(Panic())
	Ω(func() { o.MustSetValue(testStringValue) }).Should(Panic())
}

func TestSyncMultiSetStringValue(t *testing.T) {
	Default = NewGomegaWithT(t)
	var o OnceString

	errorChannel := make(chan error)

	set := func() {
		err := o.SetValue(testStringValue)
		if err != nil {
			errorChannel <- err
		}
	}

	go set()
	go set()

	Eventually(errorChannel).Should(Receive())
}
