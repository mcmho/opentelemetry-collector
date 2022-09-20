// Copyright The OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by "model/internal/cmd/pdatagen/main.go". DO NOT EDIT.
// To regenerate this file run "go run model/internal/cmd/pdatagen/main.go".

package pcommon

import (
	"go.opentelemetry.io/collector/pdata/internal"
	otlpcommon "go.opentelemetry.io/collector/pdata/internal/data/protogen/common/v1"
)

// InstrumentationScope is a message representing the instrumentation scope information.
//
// This is a reference type, if passed by value and callee modifies it the
// caller will see the modification.
//
// Must use NewInstrumentationScope function to create new instances.
// Important: zero-initialized instance is not valid for use.

type InstrumentationScope internal.InstrumentationScope

func newInstrumentationScope(orig *otlpcommon.InstrumentationScope) InstrumentationScope {
	return InstrumentationScope(internal.NewInstrumentationScope(orig))
}

func (ms InstrumentationScope) getOrig() *otlpcommon.InstrumentationScope {
	return internal.GetOrigInstrumentationScope(internal.InstrumentationScope(ms))
}

// NewInstrumentationScope creates a new empty InstrumentationScope.
//
// This must be used only in testing code. Users should use "AppendEmpty" when part of a Slice,
// OR directly access the member if this is embedded in another struct.
func NewInstrumentationScope() InstrumentationScope {
	return newInstrumentationScope(&otlpcommon.InstrumentationScope{})
}

// MoveTo moves all properties from the current struct to dest
// resetting the current instance to its zero value
func (ms InstrumentationScope) MoveTo(dest InstrumentationScope) {
	*dest.getOrig() = *ms.getOrig()
	*ms.getOrig() = otlpcommon.InstrumentationScope{}
}

// Name returns the name associated with this InstrumentationScope.
func (ms InstrumentationScope) Name() string {
	return ms.getOrig().Name
}

// SetName replaces the name associated with this InstrumentationScope.
func (ms InstrumentationScope) SetName(v string) {
	ms.getOrig().Name = v
}

// Version returns the version associated with this InstrumentationScope.
func (ms InstrumentationScope) Version() string {
	return ms.getOrig().Version
}

// SetVersion replaces the version associated with this InstrumentationScope.
func (ms InstrumentationScope) SetVersion(v string) {
	ms.getOrig().Version = v
}

// Attributes returns the Attributes associated with this InstrumentationScope.
func (ms InstrumentationScope) Attributes() Map {
	return Map(internal.NewMap(&ms.getOrig().Attributes))
}

// DroppedAttributesCount returns the droppedattributescount associated with this InstrumentationScope.
func (ms InstrumentationScope) DroppedAttributesCount() uint32 {
	return ms.getOrig().DroppedAttributesCount
}

// SetDroppedAttributesCount replaces the droppedattributescount associated with this InstrumentationScope.
func (ms InstrumentationScope) SetDroppedAttributesCount(v uint32) {
	ms.getOrig().DroppedAttributesCount = v
}

// CopyTo copies all properties from the current struct to the dest.
func (ms InstrumentationScope) CopyTo(dest InstrumentationScope) {
	dest.SetName(ms.Name())
	dest.SetVersion(ms.Version())
	ms.Attributes().CopyTo(dest.Attributes())
	dest.SetDroppedAttributesCount(ms.DroppedAttributesCount())
}

// Slice logically represents a slice of Value.
//
// This is a reference type. If passed by value and callee modifies it, the
// caller will see the modification.
//
// Must use NewSlice function to create new instances.
// Important: zero-initialized instance is not valid for use.
type Slice internal.Slice

func newSlice(orig *[]otlpcommon.AnyValue) Slice {
	return Slice(internal.NewSlice(orig))
}

func (ms Slice) getOrig() *[]otlpcommon.AnyValue {
	return internal.GetOrigSlice(internal.Slice(ms))
}

// NewSlice creates a Slice with 0 elements.
// Can use "EnsureCapacity" to initialize with a given capacity.
func NewSlice() Slice {
	orig := []otlpcommon.AnyValue(nil)
	return Slice(internal.NewSlice(&orig))
}

// Len returns the number of elements in the slice.
//
// Returns "0" for a newly instance created with "NewSlice()".
func (es Slice) Len() int {
	return len(*es.getOrig())
}

// At returns the element at the given index.
//
// This function is used mostly for iterating over all the values in the slice:
//   for i := 0; i < es.Len(); i++ {
//       e := es.At(i)
//       ... // Do something with the element
//   }
func (es Slice) At(ix int) Value {
	return newValue(&(*es.getOrig())[ix])
}

// CopyTo copies all elements from the current slice to the dest.
func (es Slice) CopyTo(dest Slice) {
	srcLen := es.Len()
	destCap := cap(*dest.getOrig())
	if srcLen <= destCap {
		(*dest.getOrig()) = (*dest.getOrig())[:srcLen:destCap]
	} else {
		(*dest.getOrig()) = make([]otlpcommon.AnyValue, srcLen)
	}

	for i := range *es.getOrig() {
		newValue(&(*es.getOrig())[i]).CopyTo(newValue(&(*dest.getOrig())[i]))
	}
}

// EnsureCapacity is an operation that ensures the slice has at least the specified capacity.
// 1. If the newCap <= cap then no change in capacity.
// 2. If the newCap > cap then the slice capacity will be expanded to equal newCap.
//
// Here is how a new Slice can be initialized:
//   es := NewSlice()
//   es.EnsureCapacity(4)
//   for i := 0; i < 4; i++ {
//       e := es.AppendEmpty()
//       // Here should set all the values for e.
//   }
func (es Slice) EnsureCapacity(newCap int) {
	oldCap := cap(*es.getOrig())
	if newCap <= oldCap {
		return
	}

	newOrig := make([]otlpcommon.AnyValue, len(*es.getOrig()), newCap)
	copy(newOrig, *es.getOrig())
	*es.getOrig() = newOrig
}

// AppendEmpty will append to the end of the slice an empty Value.
// It returns the newly added Value.
func (es Slice) AppendEmpty() Value {
	*es.getOrig() = append(*es.getOrig(), otlpcommon.AnyValue{})
	return es.At(es.Len() - 1)
}

// MoveAndAppendTo moves all elements from the current slice and appends them to the dest.
// The current slice will be cleared.
func (es Slice) MoveAndAppendTo(dest Slice) {
	if *dest.getOrig() == nil {
		// We can simply move the entire vector and avoid any allocations.
		*dest.getOrig() = *es.getOrig()
	} else {
		*dest.getOrig() = append(*dest.getOrig(), *es.getOrig()...)
	}
	*es.getOrig() = nil
}

// RemoveIf calls f sequentially for each element present in the slice.
// If f returns true, the element is removed from the slice.
func (es Slice) RemoveIf(f func(Value) bool) {
	newLen := 0
	for i := 0; i < len(*es.getOrig()); i++ {
		if f(es.At(i)) {
			continue
		}
		if newLen == i {
			// Nothing to move, element is at the right place.
			newLen++
			continue
		}
		(*es.getOrig())[newLen] = (*es.getOrig())[i]
		newLen++
	}
	// TODO: Prevent memory leak by erasing truncated values.
	*es.getOrig() = (*es.getOrig())[:newLen]
}
