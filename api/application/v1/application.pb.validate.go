// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: application/v1/application.proto

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

// Validate checks the field values on Application with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Application) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Application with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in ApplicationMultiError, or
// nil if none found.
func (m *Application) ValidateAll() error {
	return m.validate(true)
}

func (m *Application) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for Code

	// no validation rules for Name

	// no validation rules for AppType

	// no validation rules for Organization

	if all {
		switch v := interface{}(m.GetDevelopmentInfo()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ApplicationValidationError{
					field:  "DevelopmentInfo",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ApplicationValidationError{
					field:  "DevelopmentInfo",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetDevelopmentInfo()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ApplicationValidationError{
				field:  "DevelopmentInfo",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetTestInfo()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ApplicationValidationError{
					field:  "TestInfo",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ApplicationValidationError{
					field:  "TestInfo",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetTestInfo()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ApplicationValidationError{
				field:  "TestInfo",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	for idx, item := range m.GetNotificationInfos() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ApplicationValidationError{
						field:  fmt.Sprintf("NotificationInfos[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ApplicationValidationError{
						field:  fmt.Sprintf("NotificationInfos[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ApplicationValidationError{
					field:  fmt.Sprintf("NotificationInfos[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	// no validation rules for Desc

	// no validation rules for GmtCreate

	// no validation rules for GmtModify

	// no validation rules for CreateUser

	// no validation rules for ModifyUser

	if len(errors) > 0 {
		return ApplicationMultiError(errors)
	}

	return nil
}

// ApplicationMultiError is an error wrapping multiple validation errors
// returned by Application.ValidateAll() if the designated constraints aren't met.
type ApplicationMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ApplicationMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ApplicationMultiError) AllErrors() []error { return m }

// ApplicationValidationError is the validation error returned by
// Application.Validate if the designated constraints aren't met.
type ApplicationValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ApplicationValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ApplicationValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ApplicationValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ApplicationValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ApplicationValidationError) ErrorName() string { return "ApplicationValidationError" }

// Error satisfies the builtin error interface
func (e ApplicationValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sApplication.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ApplicationValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ApplicationValidationError{}

// Validate checks the field values on CreateAppRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *CreateAppRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateAppRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateAppRequestMultiError, or nil if none found.
func (m *CreateAppRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateAppRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if l := utf8.RuneCountInString(m.GetName()); l < 3 || l > 64 {
		err := CreateAppRequestValidationError{
			field:  "Name",
			reason: "value length must be between 3 and 64 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for Code

	// no validation rules for AppType

	if l := utf8.RuneCountInString(m.GetOrganization()); l < 3 || l > 64 {
		err := CreateAppRequestValidationError{
			field:  "Organization",
			reason: "value length must be between 3 and 64 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if all {
		switch v := interface{}(m.GetDevelopmentInfo()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CreateAppRequestValidationError{
					field:  "DevelopmentInfo",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CreateAppRequestValidationError{
					field:  "DevelopmentInfo",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetDevelopmentInfo()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CreateAppRequestValidationError{
				field:  "DevelopmentInfo",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetTestInfo()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CreateAppRequestValidationError{
					field:  "TestInfo",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CreateAppRequestValidationError{
					field:  "TestInfo",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetTestInfo()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CreateAppRequestValidationError{
				field:  "TestInfo",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	for idx, item := range m.GetNotificationInfos() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, CreateAppRequestValidationError{
						field:  fmt.Sprintf("NotificationInfos[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, CreateAppRequestValidationError{
						field:  fmt.Sprintf("NotificationInfos[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return CreateAppRequestValidationError{
					field:  fmt.Sprintf("NotificationInfos[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	// no validation rules for Desc

	if len(errors) > 0 {
		return CreateAppRequestMultiError(errors)
	}

	return nil
}

// CreateAppRequestMultiError is an error wrapping multiple validation errors
// returned by CreateAppRequest.ValidateAll() if the designated constraints
// aren't met.
type CreateAppRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateAppRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateAppRequestMultiError) AllErrors() []error { return m }

// CreateAppRequestValidationError is the validation error returned by
// CreateAppRequest.Validate if the designated constraints aren't met.
type CreateAppRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateAppRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateAppRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateAppRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateAppRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateAppRequestValidationError) ErrorName() string { return "CreateAppRequestValidationError" }

// Error satisfies the builtin error interface
func (e CreateAppRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateAppRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateAppRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateAppRequestValidationError{}

// Validate checks the field values on ListAppRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *ListAppRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListAppRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in ListAppRequestMultiError,
// or nil if none found.
func (m *ListAppRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *ListAppRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for Code

	// no validation rules for Name

	// no validation rules for AppType

	// no validation rules for Organization

	// no validation rules for DevelopmentLanguage

	// no validation rules for PageIndex

	// no validation rules for PageSize

	if len(errors) > 0 {
		return ListAppRequestMultiError(errors)
	}

	return nil
}

// ListAppRequestMultiError is an error wrapping multiple validation errors
// returned by ListAppRequest.ValidateAll() if the designated constraints
// aren't met.
type ListAppRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListAppRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListAppRequestMultiError) AllErrors() []error { return m }

// ListAppRequestValidationError is the validation error returned by
// ListAppRequest.Validate if the designated constraints aren't met.
type ListAppRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListAppRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListAppRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListAppRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListAppRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListAppRequestValidationError) ErrorName() string { return "ListAppRequestValidationError" }

// Error satisfies the builtin error interface
func (e ListAppRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListAppRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListAppRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListAppRequestValidationError{}

// Validate checks the field values on ListAppResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *ListAppResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListAppResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListAppResponseMultiError, or nil if none found.
func (m *ListAppResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *ListAppResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetData() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ListAppResponseValidationError{
						field:  fmt.Sprintf("Data[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ListAppResponseValidationError{
						field:  fmt.Sprintf("Data[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ListAppResponseValidationError{
					field:  fmt.Sprintf("Data[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if all {
		switch v := interface{}(m.GetPageInfo()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ListAppResponseValidationError{
					field:  "PageInfo",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ListAppResponseValidationError{
					field:  "PageInfo",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetPageInfo()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ListAppResponseValidationError{
				field:  "PageInfo",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return ListAppResponseMultiError(errors)
	}

	return nil
}

// ListAppResponseMultiError is an error wrapping multiple validation errors
// returned by ListAppResponse.ValidateAll() if the designated constraints
// aren't met.
type ListAppResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListAppResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListAppResponseMultiError) AllErrors() []error { return m }

// ListAppResponseValidationError is the validation error returned by
// ListAppResponse.Validate if the designated constraints aren't met.
type ListAppResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListAppResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListAppResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListAppResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListAppResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListAppResponseValidationError) ErrorName() string { return "ListAppResponseValidationError" }

// Error satisfies the builtin error interface
func (e ListAppResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListAppResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListAppResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListAppResponseValidationError{}

// Validate checks the field values on GetAppRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *GetAppRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetAppRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in GetAppRequestMultiError, or
// nil if none found.
func (m *GetAppRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetAppRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Code

	if len(errors) > 0 {
		return GetAppRequestMultiError(errors)
	}

	return nil
}

// GetAppRequestMultiError is an error wrapping multiple validation errors
// returned by GetAppRequest.ValidateAll() if the designated constraints
// aren't met.
type GetAppRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetAppRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetAppRequestMultiError) AllErrors() []error { return m }

// GetAppRequestValidationError is the validation error returned by
// GetAppRequest.Validate if the designated constraints aren't met.
type GetAppRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetAppRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetAppRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetAppRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetAppRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetAppRequestValidationError) ErrorName() string { return "GetAppRequestValidationError" }

// Error satisfies the builtin error interface
func (e GetAppRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetAppRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetAppRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetAppRequestValidationError{}

// Validate checks the field values on DeleteAppRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *DeleteAppRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DeleteAppRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DeleteAppRequestMultiError, or nil if none found.
func (m *DeleteAppRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *DeleteAppRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Code

	if len(errors) > 0 {
		return DeleteAppRequestMultiError(errors)
	}

	return nil
}

// DeleteAppRequestMultiError is an error wrapping multiple validation errors
// returned by DeleteAppRequest.ValidateAll() if the designated constraints
// aren't met.
type DeleteAppRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeleteAppRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeleteAppRequestMultiError) AllErrors() []error { return m }

// DeleteAppRequestValidationError is the validation error returned by
// DeleteAppRequest.Validate if the designated constraints aren't met.
type DeleteAppRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteAppRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteAppRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteAppRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteAppRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteAppRequestValidationError) ErrorName() string { return "DeleteAppRequestValidationError" }

// Error satisfies the builtin error interface
func (e DeleteAppRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteAppRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteAppRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteAppRequestValidationError{}

// Validate checks the field values on EmptyResponse with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *EmptyResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on EmptyResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in EmptyResponseMultiError, or
// nil if none found.
func (m *EmptyResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *EmptyResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return EmptyResponseMultiError(errors)
	}

	return nil
}

// EmptyResponseMultiError is an error wrapping multiple validation errors
// returned by EmptyResponse.ValidateAll() if the designated constraints
// aren't met.
type EmptyResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m EmptyResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m EmptyResponseMultiError) AllErrors() []error { return m }

// EmptyResponseValidationError is the validation error returned by
// EmptyResponse.Validate if the designated constraints aren't met.
type EmptyResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e EmptyResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e EmptyResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e EmptyResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e EmptyResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e EmptyResponseValidationError) ErrorName() string { return "EmptyResponseValidationError" }

// Error satisfies the builtin error interface
func (e EmptyResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sEmptyResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = EmptyResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = EmptyResponseValidationError{}
