package oncecell

import (
	"testing"

	. "github.com/onsi/gomega"
)

const TEST_VALUE = "test_value"

func TestSetValue(t *testing.T) {
	Default = NewGomegaWithT(t)
	var o OnceString

	err := o.SetValue(TEST_VALUE)
	Ω(err).ShouldNot(HaveOccurred())

	v, err := o.Value()
	Ω(err).ShouldNot(HaveOccurred())
	Ω(v).Should(Equal(TEST_VALUE))
}

func TestGetValue(t *testing.T) {
	Default = NewGomegaWithT(t)
	var o OnceString

	err := o.SetValue(TEST_VALUE)

	v, err := o.Value()
	Ω(err).ShouldNot(HaveOccurred())
	Ω(v).Should(Equal(TEST_VALUE))
}

func TestMustSetValue(t *testing.T) {
	Default = NewGomegaWithT(t)
	var o OnceString

	Ω(func() { o.MustSetValue(TEST_VALUE) }).ShouldNot(Panic())
}

func TestMustGetValue(t *testing.T) {
	Default = NewGomegaWithT(t)
	var o OnceString

	err := o.SetValue(TEST_VALUE)
	Ω(err).ShouldNot(HaveOccurred())

	Ω(func() { o.MustValue() }).ShouldNot(Panic())
	Ω(o.MustValue()).Should(Equal(TEST_VALUE))
}

func TestMultiSetValue(t *testing.T) {
	Default = NewGomegaWithT(t)
	var o OnceString

	err := o.SetValue(TEST_VALUE)
	Ω(err).ShouldNot(HaveOccurred())

	err = o.SetValue(TEST_VALUE)
	Ω(err).Should(HaveOccurred())
}

func TestMultiMustSetValue(t *testing.T) {
	Default = NewGomegaWithT(t)
	var o OnceString

	Ω(func() { o.MustSetValue(TEST_VALUE) }).ShouldNot(Panic())
	Ω(func() { o.MustSetValue(TEST_VALUE) }).Should(Panic())
}

func TestSyncMultiSetValue(t *testing.T) {
	Default = NewGomegaWithT(t)
	var o OnceString

	errorChannel := make(chan error)

	set := func() {
		err := o.SetValue(TEST_VALUE)
		if err != nil {
			errorChannel <- err
		}
	}

	go set()
	go set()

	Eventually(errorChannel).Should(Receive())
}
