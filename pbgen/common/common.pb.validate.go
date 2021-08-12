// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: common.proto

package common

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

// define the regex for a UUID once up-front
var _common_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on Money with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Money) Validate() error {
	if m == nil {
		return nil
	}

	if utf8.RuneCountInString(m.GetCurrencyCode()) != 3 {
		return MoneyValidationError{
			field:  "CurrencyCode",
			reason: "value length must be 3 runes",
		}

	}

	// no validation rules for Value

	return nil
}

// MoneyValidationError is the validation error returned by Money.Validate if
// the designated constraints aren't met.
type MoneyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e MoneyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e MoneyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e MoneyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e MoneyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e MoneyValidationError) ErrorName() string { return "MoneyValidationError" }

// Error satisfies the builtin error interface
func (e MoneyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sMoney.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = MoneyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = MoneyValidationError{}

// Validate checks the field values on PaymentIdentification with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *PaymentIdentification) Validate() error {
	if m == nil {
		return nil
	}

	if err := m._validateUuid(m.GetEndToEndId()); err != nil {
		return PaymentIdentificationValidationError{
			field:  "EndToEndId",
			reason: "value must be a valid UUID",
			cause:  err,
		}
	}

	if l := len(m.GetTransactionId()); l < 1 || l > 36 {
		return PaymentIdentificationValidationError{
			field:  "TransactionId",
			reason: "value length must be between 1 and 36 bytes, inclusive",
		}
	}

	if err := m._validateUuid(m.GetUetr()); err != nil {
		return PaymentIdentificationValidationError{
			field:  "Uetr",
			reason: "value must be a valid UUID",
			cause:  err,
		}
	}

	return nil
}

func (m *PaymentIdentification) _validateUuid(uuid string) error {
	if matched := _common_uuidPattern.MatchString(uuid); !matched {
		return errors.New("invalid uuid format")
	}

	return nil
}

// PaymentIdentificationValidationError is the validation error returned by
// PaymentIdentification.Validate if the designated constraints aren't met.
type PaymentIdentificationValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PaymentIdentificationValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PaymentIdentificationValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PaymentIdentificationValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PaymentIdentificationValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PaymentIdentificationValidationError) ErrorName() string {
	return "PaymentIdentificationValidationError"
}

// Error satisfies the builtin error interface
func (e PaymentIdentificationValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPaymentIdentification.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PaymentIdentificationValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PaymentIdentificationValidationError{}

// Validate checks the field values on PaymentTypeInformation with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *PaymentTypeInformation) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for ServiceLevel

	// no validation rules for LocalInstrument

	// no validation rules for PaymentChannel

	return nil
}

// PaymentTypeInformationValidationError is the validation error returned by
// PaymentTypeInformation.Validate if the designated constraints aren't met.
type PaymentTypeInformationValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PaymentTypeInformationValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PaymentTypeInformationValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PaymentTypeInformationValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PaymentTypeInformationValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PaymentTypeInformationValidationError) ErrorName() string {
	return "PaymentTypeInformationValidationError"
}

// Error satisfies the builtin error interface
func (e PaymentTypeInformationValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPaymentTypeInformation.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PaymentTypeInformationValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PaymentTypeInformationValidationError{}
