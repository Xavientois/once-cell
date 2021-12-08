package oncecell

import (
	"testing"

	. "github.com/onsi/gomega"
)

var testUint64Value uint64 = 64

func TestSetUint64Value(t *testing.T) {
	Default = NewGomegaWithT(t)
	var o OnceUint64

	err := o.SetValue(testUint64Value)
	Ω(err).ShouldNot(HaveOccurred())

	v, err := o.Value()
	Ω(err).ShouldNot(HaveOccurred())
	Ω(v).Should(Equal(testUint64Value))
}

func TestGetUint64Value(t *testing.T) {
	Default = NewGomegaWithT(t)
	var o OnceUint64

	err := o.SetValue(testUint64Value)

	v, err := o.Value()
	Ω(err).ShouldNot(HaveOccurred())
	Ω(v).Should(Equal(testUint64Value))
}

func TestMustSetUint64Value(t *testing.T) {
	Default = NewGomegaWithT(t)
	var o OnceUint64

	Ω(func() { o.MustSetValue(testUint64Value) }).ShouldNot(Panic())
}

func TestMustGetUint64Value(t *testing.T) {
	Default = NewGomegaWithT(t)
	var o OnceUint64

	err := o.SetValue(testUint64Value)
	Ω(err).ShouldNot(HaveOccurred())

	Ω(func() { o.MustValue() }).ShouldNot(Panic())
	Ω(o.MustValue()).Should(Equal(testUint64Value))
}

func TestMultiSetUint64Value(t *testing.T) {
	Default = NewGomegaWithT(t)
	var o OnceUint64

	err := o.SetValue(testUint64Value)
	Ω(err).ShouldNot(HaveOccurred())

	err = o.SetValue(testUint64Value)
	Ω(err).Should(HaveOccurred())
}

func TestMultiMustSetUint64Value(t *testing.T) {
	Default = NewGomegaWithT(t)
	var o OnceUint64

	Ω(func() { o.MustSetValue(testUint64Value) }).ShouldNot(Panic())
	Ω(func() { o.MustSetValue(testUint64Value) }).Should(Panic())
}

func TestSyncMultiSetUint64Value(t *testing.T) {
	Default = NewGomegaWithT(t)
	var o OnceUint64

	errorChannel := make(chan error)

	set := func() {
		err := o.SetValue(testUint64Value)
		if err != nil {
			errorChannel <- err
		}
	}

	go set()
	go set()

	Eventually(errorChannel).Should(Receive())
}
