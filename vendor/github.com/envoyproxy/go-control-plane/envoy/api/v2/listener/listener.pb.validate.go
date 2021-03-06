// Code generated by protoc-gen-validate
// source: envoy/api/v2/listener/listener.proto
// DO NOT EDIT!!!

package listener

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

	"github.com/gogo/protobuf/types"
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
	_ = types.DynamicAny{}
)

// Validate checks the field values on Filter with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Filter) Validate() error {
	if m == nil {
		return nil
	}

	if len(m.GetName()) < 1 {
		return FilterValidationError{
			Field:  "Name",
			Reason: "value length must be at least 1 bytes",
		}
	}

	if v, ok := interface{}(m.GetConfig()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return FilterValidationError{
				Field:  "Config",
				Reason: "embedded message failed validation",
				Cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetDeprecatedV1()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return FilterValidationError{
				Field:  "DeprecatedV1",
				Reason: "embedded message failed validation",
				Cause:  err,
			}
		}
	}

	return nil
}

// FilterValidationError is the validation error returned by Filter.Validate if
// the designated constraints aren't met.
type FilterValidationError struct {
	Field  string
	Reason string
	Cause  error
	Key    bool
}

// Error satisfies the builtin error interface
func (e FilterValidationError) Error() string {
	cause := ""
	if e.Cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.Cause)
	}

	key := ""
	if e.Key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sFilter.%s: %s%s",
		key,
		e.Field,
		e.Reason,
		cause)
}

var _ error = FilterValidationError{}

// Validate checks the field values on FilterChainMatch with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *FilterChainMatch) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetPrefixRanges() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return FilterChainMatchValidationError{
					Field:  fmt.Sprintf("PrefixRanges[%v]", idx),
					Reason: "embedded message failed validation",
					Cause:  err,
				}
			}
		}

	}

	// no validation rules for AddressSuffix

	if v, ok := interface{}(m.GetSuffixLen()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return FilterChainMatchValidationError{
				Field:  "SuffixLen",
				Reason: "embedded message failed validation",
				Cause:  err,
			}
		}
	}

	for idx, item := range m.GetSourcePrefixRanges() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return FilterChainMatchValidationError{
					Field:  fmt.Sprintf("SourcePrefixRanges[%v]", idx),
					Reason: "embedded message failed validation",
					Cause:  err,
				}
			}
		}

	}

	for idx, item := range m.GetSourcePorts() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return FilterChainMatchValidationError{
					Field:  fmt.Sprintf("SourcePorts[%v]", idx),
					Reason: "embedded message failed validation",
					Cause:  err,
				}
			}
		}

	}

	if v, ok := interface{}(m.GetDestinationPort()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return FilterChainMatchValidationError{
				Field:  "DestinationPort",
				Reason: "embedded message failed validation",
				Cause:  err,
			}
		}
	}

	// no validation rules for TransportProtocol

	return nil
}

// FilterChainMatchValidationError is the validation error returned by
// FilterChainMatch.Validate if the designated constraints aren't met.
type FilterChainMatchValidationError struct {
	Field  string
	Reason string
	Cause  error
	Key    bool
}

// Error satisfies the builtin error interface
func (e FilterChainMatchValidationError) Error() string {
	cause := ""
	if e.Cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.Cause)
	}

	key := ""
	if e.Key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sFilterChainMatch.%s: %s%s",
		key,
		e.Field,
		e.Reason,
		cause)
}

var _ error = FilterChainMatchValidationError{}

// Validate checks the field values on FilterChain with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *FilterChain) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetFilterChainMatch()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return FilterChainValidationError{
				Field:  "FilterChainMatch",
				Reason: "embedded message failed validation",
				Cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetTlsContext()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return FilterChainValidationError{
				Field:  "TlsContext",
				Reason: "embedded message failed validation",
				Cause:  err,
			}
		}
	}

	for idx, item := range m.GetFilters() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return FilterChainValidationError{
					Field:  fmt.Sprintf("Filters[%v]", idx),
					Reason: "embedded message failed validation",
					Cause:  err,
				}
			}
		}

	}

	if v, ok := interface{}(m.GetUseProxyProto()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return FilterChainValidationError{
				Field:  "UseProxyProto",
				Reason: "embedded message failed validation",
				Cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetMetadata()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return FilterChainValidationError{
				Field:  "Metadata",
				Reason: "embedded message failed validation",
				Cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetTransportSocket()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return FilterChainValidationError{
				Field:  "TransportSocket",
				Reason: "embedded message failed validation",
				Cause:  err,
			}
		}
	}

	return nil
}

// FilterChainValidationError is the validation error returned by
// FilterChain.Validate if the designated constraints aren't met.
type FilterChainValidationError struct {
	Field  string
	Reason string
	Cause  error
	Key    bool
}

// Error satisfies the builtin error interface
func (e FilterChainValidationError) Error() string {
	cause := ""
	if e.Cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.Cause)
	}

	key := ""
	if e.Key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sFilterChain.%s: %s%s",
		key,
		e.Field,
		e.Reason,
		cause)
}

var _ error = FilterChainValidationError{}

// Validate checks the field values on ListenerFilter with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *ListenerFilter) Validate() error {
	if m == nil {
		return nil
	}

	if len(m.GetName()) < 1 {
		return ListenerFilterValidationError{
			Field:  "Name",
			Reason: "value length must be at least 1 bytes",
		}
	}

	if v, ok := interface{}(m.GetConfig()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ListenerFilterValidationError{
				Field:  "Config",
				Reason: "embedded message failed validation",
				Cause:  err,
			}
		}
	}

	return nil
}

// ListenerFilterValidationError is the validation error returned by
// ListenerFilter.Validate if the designated constraints aren't met.
type ListenerFilterValidationError struct {
	Field  string
	Reason string
	Cause  error
	Key    bool
}

// Error satisfies the builtin error interface
func (e ListenerFilterValidationError) Error() string {
	cause := ""
	if e.Cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.Cause)
	}

	key := ""
	if e.Key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListenerFilter.%s: %s%s",
		key,
		e.Field,
		e.Reason,
		cause)
}

var _ error = ListenerFilterValidationError{}

// Validate checks the field values on Filter_DeprecatedV1 with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *Filter_DeprecatedV1) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Type

	return nil
}

// Filter_DeprecatedV1ValidationError is the validation error returned by
// Filter_DeprecatedV1.Validate if the designated constraints aren't met.
type Filter_DeprecatedV1ValidationError struct {
	Field  string
	Reason string
	Cause  error
	Key    bool
}

// Error satisfies the builtin error interface
func (e Filter_DeprecatedV1ValidationError) Error() string {
	cause := ""
	if e.Cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.Cause)
	}

	key := ""
	if e.Key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sFilter_DeprecatedV1.%s: %s%s",
		key,
		e.Field,
		e.Reason,
		cause)
}

var _ error = Filter_DeprecatedV1ValidationError{}
