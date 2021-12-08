package oncecell

import (
	"testing"

	. "github.com/onsi/gomega"
)

var testUint8Value uint8 = 8

func TestSetUint8Value(t *testing.T) {
	Default = NewGomegaWithT(t)
	var o OnceUint8

	err := o.SetValue(testUint8Value)
	Ω(err).ShouldNot(HaveOccurred())

	v, err := o.Value()
	Ω(err).ShouldNot(HaveOccurred())
	Ω(v).Should(Equal(testUint8Value))
}

func TestGetUint8Value(t *testing.T) {
	Default = NewGomegaWithT(t)
	var o OnceUint8

	err := o.SetValue(testUint8Value)

	v, err := o.Value()
	Ω(err).ShouldNot(HaveOccurred())
	Ω(v).Should(Equal(testUint8Value))
}

func TestMustSetUint8Value(t *testing.T) {
	Default = NewGomegaWithT(t)
	var o OnceUint8

	Ω(func() { o.MustSetValue(testUint8Value) }).ShouldNot(Panic())
}

func TestMustGetUint8Value(t *testing.T) {
	Default = NewGomegaWithT(t)
	var o OnceUint8

	err := o.SetValue(testUint8Value)
	Ω(err).ShouldNot(HaveOccurred())

	Ω(func() { o.MustValue() }).ShouldNot(Panic())
	Ω(o.MustValue()).Should(Equal(testUint8Value))
}

func TestMultiSetUint8Value(t *testing.T) {
	Default = NewGomegaWithT(t)
	var o OnceUint8

	err := o.SetValue(testUint8Value)
	Ω(err).ShouldNot(HaveOccurred())

	err = o.SetValue(testUint8Value)
	Ω(err).Should(HaveOccurred())
}

func TestMultiMustSetUint8Value(t *testing.T) {
	Default = NewGomegaWithT(t)
	var o OnceUint8

	Ω(func() { o.MustSetValue(testUint8Value) }).ShouldNot(Panic())
	Ω(func() { o.MustSetValue(testUint8Value) }).Should(Panic())
}

func TestSyncMultiSetUint8Value(t *testing.T) {
	Default = NewGomegaWithT(t)
	var o OnceUint8

	errorChannel := make(chan error)

	set := func() {
		err := o.SetValue(testUint8Value)
		if err != nil {
			errorChannel <- err
		}
	}

	go set()
	go set()

	Eventually(errorChannel).Should(Receive())
}
