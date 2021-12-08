package oncecell

import (
	"testing"

	. "github.com/onsi/gomega"
)

var testBoolValue bool = true

func TestSetBoolValue(t *testing.T) {
	Default = NewGomegaWithT(t)
	var o OnceBool

	err := o.SetValue(testBoolValue)
	Ω(err).ShouldNot(HaveOccurred())

	v, err := o.Value()
	Ω(err).ShouldNot(HaveOccurred())
	Ω(v).Should(Equal(testBoolValue))
}

func TestGetBoolValue(t *testing.T) {
	Default = NewGomegaWithT(t)
	var o OnceBool

	err := o.SetValue(testBoolValue)

	v, err := o.Value()
	Ω(err).ShouldNot(HaveOccurred())
	Ω(v).Should(Equal(testBoolValue))
}

func TestMustSetBoolValue(t *testing.T) {
	Default = NewGomegaWithT(t)
	var o OnceBool

	Ω(func() { o.MustSetValue(testBoolValue) }).ShouldNot(Panic())
}

func TestMustGetBoolValue(t *testing.T) {
	Default = NewGomegaWithT(t)
	var o OnceBool

	err := o.SetValue(testBoolValue)
	Ω(err).ShouldNot(HaveOccurred())

	Ω(func() { o.MustValue() }).ShouldNot(Panic())
	Ω(o.MustValue()).Should(Equal(testBoolValue))
}

func TestMultiSetBoolValue(t *testing.T) {
	Default = NewGomegaWithT(t)
	var o OnceBool

	err := o.SetValue(testBoolValue)
	Ω(err).ShouldNot(HaveOccurred())

	err = o.SetValue(testBoolValue)
	Ω(err).Should(HaveOccurred())
}

func TestMultiMustSetBoolValue(t *testing.T) {
	Default = NewGomegaWithT(t)
	var o OnceBool

	Ω(func() { o.MustSetValue(testBoolValue) }).ShouldNot(Panic())
	Ω(func() { o.MustSetValue(testBoolValue) }).Should(Panic())
}

func TestSyncMultiSetBoolValue(t *testing.T) {
	Default = NewGomegaWithT(t)
	var o OnceBool

	errorChannel := make(chan error)

	set := func() {
		err := o.SetValue(testBoolValue)
		if err != nil {
			errorChannel <- err
		}
	}

	go set()
	go set()

	Eventually(errorChannel).Should(Receive())
}
