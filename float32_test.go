package oncecell

import (
	"testing"

	. "github.com/onsi/gomega"
)

var testFloat32Value float32 = 0.32

func TestSetFloat32Value(t *testing.T) {
	Default = NewGomegaWithT(t)
	var o OnceFloat32

	err := o.SetValue(testFloat32Value)
	Ω(err).ShouldNot(HaveOccurred())

	v, err := o.Value()
	Ω(err).ShouldNot(HaveOccurred())
	Ω(v).Should(Equal(testFloat32Value))
}

func TestGetFloat32Value(t *testing.T) {
	Default = NewGomegaWithT(t)
	var o OnceFloat32

	err := o.SetValue(testFloat32Value)

	v, err := o.Value()
	Ω(err).ShouldNot(HaveOccurred())
	Ω(v).Should(Equal(testFloat32Value))
}

func TestMustSetFloat32Value(t *testing.T) {
	Default = NewGomegaWithT(t)
	var o OnceFloat32

	Ω(func() { o.MustSetValue(testFloat32Value) }).ShouldNot(Panic())
}

func TestMustGetFloat32Value(t *testing.T) {
	Default = NewGomegaWithT(t)
	var o OnceFloat32

	err := o.SetValue(testFloat32Value)
	Ω(err).ShouldNot(HaveOccurred())

	Ω(func() { o.MustValue() }).ShouldNot(Panic())
	Ω(o.MustValue()).Should(Equal(testFloat32Value))
}

func TestMultiSetFloat32Value(t *testing.T) {
	Default = NewGomegaWithT(t)
	var o OnceFloat32

	err := o.SetValue(testFloat32Value)
	Ω(err).ShouldNot(HaveOccurred())

	err = o.SetValue(testFloat32Value)
	Ω(err).Should(HaveOccurred())
}

func TestMultiMustSetFloat32Value(t *testing.T) {
	Default = NewGomegaWithT(t)
	var o OnceFloat32

	Ω(func() { o.MustSetValue(testFloat32Value) }).ShouldNot(Panic())
	Ω(func() { o.MustSetValue(testFloat32Value) }).Should(Panic())
}

func TestSyncMultiSetFloat32Value(t *testing.T) {
	Default = NewGomegaWithT(t)
	var o OnceFloat32

	errorChannel := make(chan error)

	set := func() {
		err := o.SetValue(testFloat32Value)
		if err != nil {
			errorChannel <- err
		}
	}

	go set()
	go set()

	Eventually(errorChannel).Should(Receive())
}
