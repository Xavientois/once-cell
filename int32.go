package oncecell

import (
	"errors"
	"sync/atomic"
)

// OnceInt32 wraps a int32 and enforces at runtime that the value is set at most once.
type OnceInt32 struct {
	value int32
	isSet uint32
}

// SetValue sets the value of the int32 to the provided value.
// This function will return an error if called more than once.
// This is thread-safe as the isSet flag is atomically set, thus ensuring that no two
// callers can check the flag at the same time.
func (o *OnceInt32) SetValue(value int32) error {
	if atomic.CompareAndSwapUint32(&o.isSet, 0, 1) {
		o.value = value
		return nil
	} else {
		return errors.New(setErrorMsg)
	}
}

// Value returns the value stored in the cell or an error if it has not yet been set
func (o *OnceInt32) Value() (retval int32, err error) {
	if atomic.LoadUint32(&o.isSet) == 1 {
		retval = o.value
	} else {
		err = errors.New(getErrorMsg)
	}
	return
}

// MustSetValue sets the value of the int32 to the provided value.
// This function will panic if the value has already been set
// This is thread-safe as the isSet flag is atomically set, thus ensuring that no two
// callers can check the flag at the same time.
func (o *OnceInt32) MustSetValue(value int32) {
	if atomic.CompareAndSwapUint32(&o.isSet, 0, 1) {
		o.value = value
	} else {
		panic(setErrorMsg)
	}
}

// MustValue returns the value stored in the cell or panics if it has not yet been set
func (o *OnceInt32) MustValue() int32 {
	if atomic.LoadUint32(&o.isSet) == 1 {
		return o.value
	} else {
		panic(getErrorMsg)
	}
}
