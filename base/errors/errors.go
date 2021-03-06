package errors

import (
	"bytes"
	"fmt"
)

// A special diagnosis error structure, which allows downstream call site to emit more notes after
// creation.
type DError struct {
	notes     []string // Reverse order
	rootCause error
}

// -----------------------------------------------------------------------------
// Conform error interface.
// -----------------------------------------------------------------------------
func (de *DError) Error() string { return de.String() }

// -----------------------------------------------------------------------------
// String().
//
// Formats the error stack into string.
// -----------------------------------------------------------------------------
func (de *DError) String() string {
	var buf bytes.Buffer

	fmt.Fprint(&buf, "\nDiagnosis Error:\n")

	for index := len(de.notes) - 1; index >= 0; index-- {
		fmt.Fprintf(&buf, "  > %v\n", de.notes[index])
	}

	fmt.Fprintf(&buf, "  @ %v\n", de.rootCause)
	return buf.String()
}

// -----------------------------------------------------------------------------
// Factory Methods.
// -----------------------------------------------------------------------------

// Wraps a note to the error and return as error interface.
//
// This is same as From(error).EmitNote(...)
func WrapNote(err error, sfmt string, args ...interface{}) error {
	if de, ok := err.(*DError); ok {
		note := fmt.Sprintf(sfmt, args...)
		de.notes = append(de.notes, note)
		return de
	} else {
		return From(err).EmitNote(sfmt, args...)
	}
}

// Creates a DError with root cause specified by the message.
func New(sfmt string, args ...interface{}) *DError {
	return &DError{
		rootCause: fmt.Errorf(sfmt, args...),
	}
}

// Creates a DError with root cause specified by `err`.
func From(err error) *DError {
	if de, ok := err.(*DError); ok {
		return de
	}

	return &DError{
		rootCause: err,
	}
}

// -----------------------------------------------------------------------------
// Public APIs.
// -----------------------------------------------------------------------------

// Emit one more note to the DError.
func (de *DError) EmitNote(sfmt string, args ...interface{}) *DError {
	note := fmt.Sprintf(sfmt, args...)
	de.notes = append(de.notes, note)
	return de
}
