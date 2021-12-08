package oncecell

import (
	"errors"
	"sync/atomic"
)

const GET_ERROR_MSG = "Tried to get oncecell value before it is set"
const SET_ERROR_MSG = "Tried to set oncecell multiple times"

/// OnceString wraps a string and enforces at runtime that the value is set at most once.
type OnceString struct {
	value string
	isSet uint32
}

/// SetValue sets the value of the string to the provided value.
/// This function will return an error if called more than once.
/// This is thread-safe as the isSet flag is atomically set, thus ensuring that no two
/// callers can check the flag at the same time.
func (o *OnceString) SetValue(value string) error {
	if atomic.CompareAndSwapUint32(&o.isSet, 0, 1) {
		o.value = value
		return nil
	} else {
		return errors.New(SET_ERROR_MSG)
	}
}

/// Value returns the value stored in the cell or an error if it has not yet been set
func (o *OnceString) Value() (retval string, err error) {
	if atomic.LoadUint32(&o.isSet) == 1 {
		retval = o.value
	} else {
		err = errors.New(GET_ERROR_MSG)
	}
	return
}

/// MustSetValue sets the value of the string to the provided value.
/// This function will panic if the value has already been set
/// This is thread-safe as the isSet flag is atomically set, thus ensuring that no two
/// callers can check the flag at the same time.
func (o *OnceString) MustSetValue(value string) {
	if atomic.CompareAndSwapUint32(&o.isSet, 0, 1) {
		o.value = value
	} else {
		panic(SET_ERROR_MSG)
	}
}

/// MustValue returns the value stored in the cell or panics if it has not yet been set
func (o *OnceString) MustValue() string {
	if atomic.LoadUint32(&o.isSet) == 1 {
		return o.value
	} else {
		panic(GET_ERROR_MSG)
	}
}
