// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: code.proto

package proto

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/golang/protobuf/ptypes"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = ptypes.DynamicAny{}
)

// define the regex for a UUID once up-front
var _code_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on EventWithCode with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *EventWithCode) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Type

	if v, ok := interface{}(m.GetEventA()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return EventWithCodeValidationError{
				field:  "EventA",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetEventB()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return EventWithCodeValidationError{
				field:  "EventB",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetEventC()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return EventWithCodeValidationError{
				field:  "EventC",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// EventWithCodeValidationError is the validation error returned by
// EventWithCode.Validate if the designated constraints aren't met.
type EventWithCodeValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e EventWithCodeValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e EventWithCodeValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e EventWithCodeValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e EventWithCodeValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e EventWithCodeValidationError) ErrorName() string { return "EventWithCodeValidationError" }

// Error satisfies the builtin error interface
func (e EventWithCodeValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sEventWithCode.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = EventWithCodeValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = EventWithCodeValidationError{}

// Validate checks the field values on EventWithCodeA with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *EventWithCodeA) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for A

	return nil
}

// EventWithCodeAValidationError is the validation error returned by
// EventWithCodeA.Validate if the designated constraints aren't met.
type EventWithCodeAValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e EventWithCodeAValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e EventWithCodeAValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e EventWithCodeAValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e EventWithCodeAValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e EventWithCodeAValidationError) ErrorName() string { return "EventWithCodeAValidationError" }

// Error satisfies the builtin error interface
func (e EventWithCodeAValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sEventWithCodeA.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = EventWithCodeAValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = EventWithCodeAValidationError{}

// Validate checks the field values on EventWithCodeB with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *EventWithCodeB) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for A

	// no validation rules for B

	return nil
}

// EventWithCodeBValidationError is the validation error returned by
// EventWithCodeB.Validate if the designated constraints aren't met.
type EventWithCodeBValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e EventWithCodeBValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e EventWithCodeBValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e EventWithCodeBValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e EventWithCodeBValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e EventWithCodeBValidationError) ErrorName() string { return "EventWithCodeBValidationError" }

// Error satisfies the builtin error interface
func (e EventWithCodeBValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sEventWithCodeB.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = EventWithCodeBValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = EventWithCodeBValidationError{}

// Validate checks the field values on EventWithCodeC with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *EventWithCodeC) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for A

	// no validation rules for B

	if v, ok := interface{}(m.GetC()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return EventWithCodeCValidationError{
				field:  "C",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// EventWithCodeCValidationError is the validation error returned by
// EventWithCodeC.Validate if the designated constraints aren't met.
type EventWithCodeCValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e EventWithCodeCValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e EventWithCodeCValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e EventWithCodeCValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e EventWithCodeCValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e EventWithCodeCValidationError) ErrorName() string { return "EventWithCodeCValidationError" }

// Error satisfies the builtin error interface
func (e EventWithCodeCValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sEventWithCodeC.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = EventWithCodeCValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = EventWithCodeCValidationError{}

// Validate checks the field values on EventWithCodeC_Context with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *EventWithCodeC_Context) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for X

	// no validation rules for Y

	return nil
}

// EventWithCodeC_ContextValidationError is the validation error returned by
// EventWithCodeC_Context.Validate if the designated constraints aren't met.
type EventWithCodeC_ContextValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e EventWithCodeC_ContextValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e EventWithCodeC_ContextValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e EventWithCodeC_ContextValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e EventWithCodeC_ContextValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e EventWithCodeC_ContextValidationError) ErrorName() string {
	return "EventWithCodeC_ContextValidationError"
}

// Error satisfies the builtin error interface
func (e EventWithCodeC_ContextValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sEventWithCodeC_Context.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = EventWithCodeC_ContextValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = EventWithCodeC_ContextValidationError{}
