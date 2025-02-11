// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: api/operator/v1/operator.proto

package v1

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"sort"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"
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
	_ = anypb.Any{}
	_ = sort.Sort
)

// Validate checks the field values on ResolveAppealRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ResolveAppealRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ResolveAppealRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ResolveAppealRequestMultiError, or nil if none found.
func (m *ResolveAppealRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *ResolveAppealRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for AppealID

	// no validation rules for OperatorID

	if l := utf8.RuneCountInString(m.GetReason()); l < 1 || l > 250 {
		err := ResolveAppealRequestValidationError{
			field:  "Reason",
			reason: "value length must be between 1 and 250 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if l := utf8.RuneCountInString(m.GetContent()); l < 1 || l > 1024 {
		err := ResolveAppealRequestValidationError{
			field:  "Content",
			reason: "value length must be between 1 and 1024 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for Status

	if len(errors) > 0 {
		return ResolveAppealRequestMultiError(errors)
	}

	return nil
}

// ResolveAppealRequestMultiError is an error wrapping multiple validation
// errors returned by ResolveAppealRequest.ValidateAll() if the designated
// constraints aren't met.
type ResolveAppealRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ResolveAppealRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ResolveAppealRequestMultiError) AllErrors() []error { return m }

// ResolveAppealRequestValidationError is the validation error returned by
// ResolveAppealRequest.Validate if the designated constraints aren't met.
type ResolveAppealRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ResolveAppealRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ResolveAppealRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ResolveAppealRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ResolveAppealRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ResolveAppealRequestValidationError) ErrorName() string {
	return "ResolveAppealRequestValidationError"
}

// Error satisfies the builtin error interface
func (e ResolveAppealRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sResolveAppealRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ResolveAppealRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ResolveAppealRequestValidationError{}

// Validate checks the field values on ResolveAppealReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ResolveAppealReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ResolveAppealReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ResolveAppealReplyMultiError, or nil if none found.
func (m *ResolveAppealReply) ValidateAll() error {
	return m.validate(true)
}

func (m *ResolveAppealReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for AppealID

	// no validation rules for Status

	if len(errors) > 0 {
		return ResolveAppealReplyMultiError(errors)
	}

	return nil
}

// ResolveAppealReplyMultiError is an error wrapping multiple validation errors
// returned by ResolveAppealReply.ValidateAll() if the designated constraints
// aren't met.
type ResolveAppealReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ResolveAppealReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ResolveAppealReplyMultiError) AllErrors() []error { return m }

// ResolveAppealReplyValidationError is the validation error returned by
// ResolveAppealReply.Validate if the designated constraints aren't met.
type ResolveAppealReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ResolveAppealReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ResolveAppealReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ResolveAppealReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ResolveAppealReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ResolveAppealReplyValidationError) ErrorName() string {
	return "ResolveAppealReplyValidationError"
}

// Error satisfies the builtin error interface
func (e ResolveAppealReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sResolveAppealReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ResolveAppealReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ResolveAppealReplyValidationError{}
