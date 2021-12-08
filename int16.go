package oncecell

import (
	"errors"
	"sync/atomic"
)

// OnceInt16 wraps a int16 and enforces at runtime that the value is set at most once.
type OnceInt16 struct {
	value int16
	isSet uint32
}

// SetValue sets the value of the int16 to the provided value.
// This function will return an error if called more than once.
// This is thread-safe as the isSet flag is atomically set, thus ensuring that no two
// callers can check the flag at the same time.
func (o *OnceInt16) SetValue(value int16) error {
	if atomic.CompareAndSwapUint32(&o.isSet, 0, 1) {
		o.value = value
		return nil
	} else {
		return errors.New(setErrorMsg)
	}
}

// Value returns the value stored in the cell or an error if it has not yet been set
func (o *OnceInt16) Value() (retval int16, err error) {
	if atomic.LoadUint32(&o.isSet) == 1 {
		retval = o.value
	} else {
		err = errors.New(getErrorMsg)
	}
	return
}

// MustSetValue sets the value of the int16 to the provided value.
// This function will panic if the value has already been set
// This is thread-safe as the isSet flag is atomically set, thus ensuring that no two
// callers can check the flag at the same time.
func (o *OnceInt16) MustSetValue(value int16) {
	if atomic.CompareAndSwapUint32(&o.isSet, 0, 1) {
		o.value = value
	} else {
		panic(setErrorMsg)
	}
}

// MustValue returns the value stored in the cell or panics if it has not yet been set
func (o *OnceInt16) MustValue() int16 {
	if atomic.LoadUint32(&o.isSet) == 1 {
		return o.value
	} else {
		panic(getErrorMsg)
	}
}
