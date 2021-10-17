// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: spots/proto/spots_response.proto

package dto

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
)

// Validate checks the field values on SpotsAddResponse with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *SpotsAddResponse) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Status

	if v, ok := interface{}(m.GetData()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return SpotsAddResponseValidationError{
				field:  "Data",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// SpotsAddResponseValidationError is the validation error returned by
// SpotsAddResponse.Validate if the designated constraints aren't met.
type SpotsAddResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SpotsAddResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SpotsAddResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SpotsAddResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SpotsAddResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SpotsAddResponseValidationError) ErrorName() string { return "SpotsAddResponseValidationError" }

// Error satisfies the builtin error interface
func (e SpotsAddResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSpotsAddResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SpotsAddResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SpotsAddResponseValidationError{}

// Validate checks the field values on SpotsAddData with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *SpotsAddData) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	return nil
}

// SpotsAddDataValidationError is the validation error returned by
// SpotsAddData.Validate if the designated constraints aren't met.
type SpotsAddDataValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SpotsAddDataValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SpotsAddDataValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SpotsAddDataValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SpotsAddDataValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SpotsAddDataValidationError) ErrorName() string { return "SpotsAddDataValidationError" }

// Error satisfies the builtin error interface
func (e SpotsAddDataValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSpotsAddData.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SpotsAddDataValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SpotsAddDataValidationError{}

// Validate checks the field values on SpotsUpdateResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *SpotsUpdateResponse) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Status

	if v, ok := interface{}(m.GetData()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return SpotsUpdateResponseValidationError{
				field:  "Data",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// SpotsUpdateResponseValidationError is the validation error returned by
// SpotsUpdateResponse.Validate if the designated constraints aren't met.
type SpotsUpdateResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SpotsUpdateResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SpotsUpdateResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SpotsUpdateResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SpotsUpdateResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SpotsUpdateResponseValidationError) ErrorName() string {
	return "SpotsUpdateResponseValidationError"
}

// Error satisfies the builtin error interface
func (e SpotsUpdateResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSpotsUpdateResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SpotsUpdateResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SpotsUpdateResponseValidationError{}

// Validate checks the field values on SpotsUpdateData with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *SpotsUpdateData) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	// no validation rules for Name

	// no validation rules for Address

	if v, ok := interface{}(m.GetCoordinates()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return SpotsUpdateDataValidationError{
				field:  "Coordinates",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for Lighting

	// no validation rules for Friendly

	// no validation rules for Verified

	return nil
}

// SpotsUpdateDataValidationError is the validation error returned by
// SpotsUpdateData.Validate if the designated constraints aren't met.
type SpotsUpdateDataValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SpotsUpdateDataValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SpotsUpdateDataValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SpotsUpdateDataValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SpotsUpdateDataValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SpotsUpdateDataValidationError) ErrorName() string { return "SpotsUpdateDataValidationError" }

// Error satisfies the builtin error interface
func (e SpotsUpdateDataValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSpotsUpdateData.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SpotsUpdateDataValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SpotsUpdateDataValidationError{}

// Validate checks the field values on ResponseCoordinates with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ResponseCoordinates) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Lat

	// no validation rules for Long

	return nil
}

// ResponseCoordinatesValidationError is the validation error returned by
// ResponseCoordinates.Validate if the designated constraints aren't met.
type ResponseCoordinatesValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ResponseCoordinatesValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ResponseCoordinatesValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ResponseCoordinatesValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ResponseCoordinatesValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ResponseCoordinatesValidationError) ErrorName() string {
	return "ResponseCoordinatesValidationError"
}

// Error satisfies the builtin error interface
func (e ResponseCoordinatesValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sResponseCoordinates.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ResponseCoordinatesValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ResponseCoordinatesValidationError{}

// Validate checks the field values on SpotsGetResponse with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *SpotsGetResponse) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Status

	if v, ok := interface{}(m.GetData()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return SpotsGetResponseValidationError{
				field:  "Data",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// SpotsGetResponseValidationError is the validation error returned by
// SpotsGetResponse.Validate if the designated constraints aren't met.
type SpotsGetResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SpotsGetResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SpotsGetResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SpotsGetResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SpotsGetResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SpotsGetResponseValidationError) ErrorName() string { return "SpotsGetResponseValidationError" }

// Error satisfies the builtin error interface
func (e SpotsGetResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSpotsGetResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SpotsGetResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SpotsGetResponseValidationError{}

// Validate checks the field values on SpotsGetData with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *SpotsGetData) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	// no validation rules for Name

	// no validation rules for Address

	if v, ok := interface{}(m.GetCoordinates()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return SpotsGetDataValidationError{
				field:  "Coordinates",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for Lighting

	// no validation rules for Friendly

	// no validation rules for Verified

	return nil
}

// SpotsGetDataValidationError is the validation error returned by
// SpotsGetData.Validate if the designated constraints aren't met.
type SpotsGetDataValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SpotsGetDataValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SpotsGetDataValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SpotsGetDataValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SpotsGetDataValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SpotsGetDataValidationError) ErrorName() string { return "SpotsGetDataValidationError" }

// Error satisfies the builtin error interface
func (e SpotsGetDataValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSpotsGetData.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SpotsGetDataValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SpotsGetDataValidationError{}

// Validate checks the field values on SpotsGetAllResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *SpotsGetAllResponse) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Status

	for idx, item := range m.GetData() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return SpotsGetAllResponseValidationError{
					field:  fmt.Sprintf("Data[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// SpotsGetAllResponseValidationError is the validation error returned by
// SpotsGetAllResponse.Validate if the designated constraints aren't met.
type SpotsGetAllResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SpotsGetAllResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SpotsGetAllResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SpotsGetAllResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SpotsGetAllResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SpotsGetAllResponseValidationError) ErrorName() string {
	return "SpotsGetAllResponseValidationError"
}

// Error satisfies the builtin error interface
func (e SpotsGetAllResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSpotsGetAllResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SpotsGetAllResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SpotsGetAllResponseValidationError{}
