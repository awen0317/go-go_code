// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/admin/v3alpha/memory.proto

package envoy_admin_v3alpha

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
var _memory_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on Memory with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Memory) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Allocated

	// no validation rules for HeapSize

	// no validation rules for PageheapUnmapped

	// no validation rules for PageheapFree

	// no validation rules for TotalThreadCache

	return nil
}

// MemoryValidationError is the validation error returned by Memory.Validate if
// the designated constraints aren't met.
type MemoryValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e MemoryValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e MemoryValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e MemoryValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e MemoryValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e MemoryValidationError) ErrorName() string { return "MemoryValidationError" }

// Error satisfies the builtin error interface
func (e MemoryValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sMemory.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = MemoryValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = MemoryValidationError{}
