// Copyright 2013 Travis Keep. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file or
// at http://opensource.org/licenses/BSD-3-Clause.

// Package maybe provides optional values as an alternative to using pointers
// which can lead to many small chunks on the heap.
// Programs using maybe instances should typically store and pass them as
// values, not pointers.
// Maybe instances fully support assignment (=) and equality (==) operators.
package maybe

import "fmt"

// Int represents an int value or nothing. The zero value is nothing.
type Int struct {
  // The int value 
  Value int
  // True if this instance represents an int, false if it represents nothing.
  Valid bool
}

// NewInt returns an instance representing the value x.
func NewInt(x int) Int {
  return Int{Value: x, Valid: true}
}

// Set makes this instance represent the value x.
func (m *Int) Set(x int) {
  m.Value = x
  m.Valid = true
}

// Clear makes this instance represent nothing.
func(m *Int) Clear() {
  m.Value = 0
  m.Valid = false
}

func (m Int) String() string {
  if !m.Valid {
    return "Nothing"
  }
  return fmt.Sprintf("Just %d", m.Value)
}

// Int8 represents an int8 value or nothing. The zero value is nothing.
type Int8 struct {
  // The int8 value 
  Value int8
  // True if this instance represents an int8, false if it represents nothing.
  Valid bool
}

// NewInt8 returns an instance representing the value x.
func NewInt8(x int8) Int8 {
  return Int8{Value: x, Valid: true}
}

// Set makes this instance represent the value x.
func (m *Int8) Set(x int8) {
  m.Value = x
  m.Valid = true
}

// Clear makes this instance represent nothing.
func(m *Int8) Clear() {
  m.Value = 0
  m.Valid = false
}

func (m Int8) String() string {
  if !m.Valid {
    return "Nothing"
  }
  return fmt.Sprintf("Just %d", m.Value)
}

// Int16 represents an int16 value or nothing. The zero value is nothing.
type Int16 struct {
  // The int16 value 
  Value int16
  // True if this instance represents an int16, false if it represents nothing.
  Valid bool
}

// NewInt16 returns an instance representing the value x.
func NewInt16(x int16) Int16 {
  return Int16{Value: x, Valid: true}
}

// Set makes this instance represent the value x.
func (m *Int16) Set(x int16) {
  m.Value = x
  m.Valid = true
}

// Clear makes this instance represent nothing.
func(m *Int16) Clear() {
  m.Value = 0
  m.Valid = false
}

func (m Int16) String() string {
  if !m.Valid {
    return "Nothing"
  }
  return fmt.Sprintf("Just %d", m.Value)
}

// Int32 represents an int32 value or nothing. The zero value is nothing.
type Int32 struct {
  // The int32 value 
  Value int32
  // True if this instance represents an int32, false if it represents nothing.
  Valid bool
}

// NewInt32 returns an instance representing the value x.
func NewInt32(x int32) Int32 {
  return Int32{Value: x, Valid: true}
}

// Set makes this instance represent the value x.
func (m *Int32) Set(x int32) {
  m.Value = x
  m.Valid = true
}

// Clear makes this instance represent nothing.
func(m *Int32) Clear() {
  m.Value = 0
  m.Valid = false
}

func (m Int32) String() string {
  if !m.Valid {
    return "Nothing"
  }
  return fmt.Sprintf("Just %d", m.Value)
}

// Int64 represents an int64 value or nothing. The zero value is nothing.
type Int64 struct {
  // The int64 value 
  Value int64
  // True if this instance represents an int64, false if it represents nothing.
  Valid bool
}

// NewInt64 returns an instance representing the value x.
func NewInt64(x int64) Int64 {
  return Int64{Value: x, Valid: true}
}

// Set makes this instance represent the value x.
func (m *Int64) Set(x int64) {
  m.Value = x
  m.Valid = true
}

// Clear makes this instance represent nothing.
func(m *Int64) Clear() {
  m.Value = 0
  m.Valid = false
}

func (m Int64) String() string {
  if !m.Valid {
    return "Nothing"
  }
  return fmt.Sprintf("Just %d", m.Value)
}

// Uint represents an uint value or nothing. The zero value is nothing.
type Uint struct {
  // The uint value 
  Value uint
  // True if this instance represents an uint, false if it represents nothing.
  Valid bool
}

// NewUint returns an instance representing the value x.
func NewUint(x uint) Uint {
  return Uint{Value: x, Valid: true}
}

// Set makes this instance represent the value x.
func (m *Uint) Set(x uint) {
  m.Value = x
  m.Valid = true
}

// Clear makes this instance represent nothing.
func(m *Uint) Clear() {
  m.Value = 0
  m.Valid = false
}

func (m Uint) String() string {
  if !m.Valid {
    return "Nothing"
  }
  return fmt.Sprintf("Just %d", m.Value)
}

// Uint8 represents an uint8 value or nothing. The zero value is nothing.
type Uint8 struct {
  // The uint8 value 
  Value uint8
  // True if this instance represents an uint8, false if it represents nothing.
  Valid bool
}

// NewUint8 returns an instance representing the value x.
func NewUint8(x uint8) Uint8 {
  return Uint8{Value: x, Valid: true}
}

// Set makes this instance represent the value x.
func (m *Uint8) Set(x uint8) {
  m.Value = x
  m.Valid = true
}

// Clear makes this instance represent nothing.
func(m *Uint8) Clear() {
  m.Value = 0
  m.Valid = false
}

func (m Uint8) String() string {
  if !m.Valid {
    return "Nothing"
  }
  return fmt.Sprintf("Just %d", m.Value)
}

// Uint16 represents an uint16 value or nothing. The zero value is nothing.
type Uint16 struct {
  // The uint16 value 
  Value uint16
  // True if this instance represents an uint16, false if it represents nothing.
  Valid bool
}

// NewUint16 returns an instance representing the value x.
func NewUint16(x uint16) Uint16 {
  return Uint16{Value: x, Valid: true}
}

// Set makes this instance represent the value x.
func (m *Uint16) Set(x uint16) {
  m.Value = x
  m.Valid = true
}

// Clear makes this instance represent nothing.
func(m *Uint16) Clear() {
  m.Value = 0
  m.Valid = false
}

func (m Uint16) String() string {
  if !m.Valid {
    return "Nothing"
  }
  return fmt.Sprintf("Just %d", m.Value)
}

// Uint32 represents an uint32 value or nothing. The zero value is nothing.
type Uint32 struct {
  // The uint32 value 
  Value uint32
  // True if this instance represents an uint32, false if it represents nothing.
  Valid bool
}

// NewUint32 returns an instance representing the value x.
func NewUint32(x uint32) Uint32 {
  return Uint32{Value: x, Valid: true}
}

// Set makes this instance represent the value x.
func (m *Uint32) Set(x uint32) {
  m.Value = x
  m.Valid = true
}

// Clear makes this instance represent nothing.
func(m *Uint32) Clear() {
  m.Value = 0
  m.Valid = false
}

func (m Uint32) String() string {
  if !m.Valid {
    return "Nothing"
  }
  return fmt.Sprintf("Just %d", m.Value)
}

// Uint64 represents an uint64 value or nothing. The zero value is nothing.
type Uint64 struct {
  // The uint64 value 
  Value uint64
  // True if this instance represents an uint64, false if it represents nothing.
  Valid bool
}

// NewUint64 returns an instance representing the value x.
func NewUint64(x uint64) Uint64 {
  return Uint64{Value: x, Valid: true}
}

// Set makes this instance represent the value x.
func (m *Uint64) Set(x uint64) {
  m.Value = x
  m.Valid = true
}

// Clear makes this instance represent nothing.
func(m *Uint64) Clear() {
  m.Value = 0
  m.Valid = false
}

func (m Uint64) String() string {
  if !m.Valid {
    return "Nothing"
  }
  return fmt.Sprintf("Just %d", m.Value)
}

// Float32 represents an float32 value or nothing. The zero value is nothing.
type Float32 struct {
  // The float32 value 
  Value float32
  // True if this instance represents an float32,
  // false if it represents nothing.
  Valid bool
}

// NewFloat32 returns an instance representing the value x.
func NewFloat32(x float32) Float32 {
  return Float32{Value: x, Valid: true}
}

// Set makes this instance represent the value x.
func (m *Float32) Set(x float32) {
  m.Value = x
  m.Valid = true
}

// Clear makes this instance represent nothing.
func(m *Float32) Clear() {
  m.Value = 0.0
  m.Valid = false
}

func (m Float32) String() string {
  if !m.Valid {
    return "Nothing"
  }
  return fmt.Sprintf("Just %f", m.Value)
}

// Float64 represents an float64 value or nothing. The zero value is nothing.
type Float64 struct {
  // The float64 value 
  Value float64
  // True if this instance represents an float64,
  // false if it represents nothing.
  Valid bool
}

// NewFloat64 returns an instance representing the value x.
func NewFloat64(x float64) Float64 {
  return Float64{Value: x, Valid: true}
}

// Set makes this instance represent the value x.
func (m *Float64) Set(x float64) {
  m.Value = x
  m.Valid = true
}

// Clear makes this instance represent nothing.
func(m *Float64) Clear() {
  m.Value = 0.0
  m.Valid = false
}

func (m Float64) String() string {
  if !m.Valid {
    return "Nothing"
  }
  return fmt.Sprintf("Just %f", m.Value)
}

// String represents an string value or nothing. The zero value is nothing.
type String struct {
  // The string value 
  Value string
  // True if this instance represents an string,
  // false if it represents nothing.
  Valid bool
}

// NewString returns an instance representing the value x.
func NewString(x string) String {
  return String{Value: x, Valid: true}
}

// Set makes this instance represent the value x.
func (m *String) Set(x string) {
  m.Value = x
  m.Valid = true
}

// Clear makes this instance represent nothing.
func(m *String) Clear() {
  m.Value = ""
  m.Valid = false
}

func (m String) String() string {
  if !m.Valid {
    return "Nothing"
  }
  return fmt.Sprintf("Just %s", m.Value)
}

// Bool represents an bool value or nothing. The zero value is nothing.
type Bool struct {
  // The bool value 
  Value bool
  // True if this instance represents an bool,
  // false if it represents nothing.
  Valid bool
}

// NewBool returns an instance representing the value x.
func NewBool(x bool) Bool {
  return Bool{Value: x, Valid: true}
}

// Set makes this instance represent the value x.
func (m *Bool) Set(x bool) {
  m.Value = x
  m.Valid = true
}

// Clear makes this instance represent nothing.
func(m *Bool) Clear() {
  m.Value = false
  m.Valid = false
}

func (m Bool) String() string {
  if !m.Valid {
    return "Nothing"
  }
  if m.Value {
    return "Just true"
  }
  return "Just false"
}


