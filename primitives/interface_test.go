package oncecell

import (
	"testing"

	. "github.com/onsi/gomega"
)

type myStruct struct {
	field int
}

var testValue myStruct = myStruct{field: 1}

func TestSetValue(t *testing.T) {
	Default = NewGomegaWithT(t)
	var o Once

	err := o.SetValue(testValue)
	Ω(err).ShouldNot(HaveOccurred())

	v, err := o.Value()
	Ω(err).ShouldNot(HaveOccurred())
	Ω(v).Should(Equal(testValue))
}

func TestGetValue(t *testing.T) {
	Default = NewGomegaWithT(t)
	var o Once

	o.SetValue(testValue)

	v, err := o.Value()
	Ω(err).ShouldNot(HaveOccurred())
	Ω(v).Should(Equal(testValue))
}

func TestMustSetValue(t *testing.T) {
	Default = NewGomegaWithT(t)
	var o Once

	Ω(func() { o.MustSetValue(testValue) }).ShouldNot(Panic())
}

func TestMustGetValue(t *testing.T) {
	Default = NewGomegaWithT(t)
	var o Once

	err := o.SetValue(testValue)
	Ω(err).ShouldNot(HaveOccurred())

	Ω(func() { o.MustValue() }).ShouldNot(Panic())
	Ω(o.MustValue()).Should(Equal(testValue))
}

func TestMultiSetValue(t *testing.T) {
	Default = NewGomegaWithT(t)
	var o Once

	err := o.SetValue(testValue)
	Ω(err).ShouldNot(HaveOccurred())

	err = o.SetValue(testValue)
	Ω(err).Should(HaveOccurred())
}

func TestMultiMustSetValue(t *testing.T) {
	Default = NewGomegaWithT(t)
	var o Once

	Ω(func() { o.MustSetValue(testValue) }).ShouldNot(Panic())
	Ω(func() { o.MustSetValue(testValue) }).Should(Panic())
}

func TestSyncMultiSetValue(t *testing.T) {
	Default = NewGomegaWithT(t)
	var o Once

	errorChannel := make(chan error)

	set := func() {
		err := o.SetValue(testValue)
		if err != nil {
			errorChannel <- err
		}
	}

	go set()
	go set()

	Eventually(errorChannel).Should(Receive())
}
