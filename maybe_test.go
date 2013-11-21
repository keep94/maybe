// Copyright 2013 Travis Keep. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file or
// at http://opensource.org/licenses/BSD-3-Clause.

package maybe_test;

import (
  "github.com/keep94/maybe"
  "testing"
)

func TestMaybeInt(t *testing.T) {
  var m, c maybe.Int
  var v int = 7
  m.Set(v)
  if m != maybe.NewInt(v) {
    t.Error("Int Set broken.")
  }
  verifyString(t, "Just 7", m.String())
  m.Clear()
  if m != c {
    t.Error("Int Clear broken.")
  }
  verifyString(t, "Nothing", m.String())
}

func TestMaybeInt8(t *testing.T) {
  var m, c maybe.Int8
  var v int8 = 7
  m.Set(v)
  if m != maybe.NewInt8(v) {
    t.Error("Int8 Set broken.")
  }
  verifyString(t, "Just 7", m.String())
  m.Clear()
  if m != c {
    t.Error("Int8 Clear broken.")
  }
  verifyString(t, "Nothing", m.String())
}

func TestMaybeInt16(t *testing.T) {
  var m, c maybe.Int16
  var v int16 = 7
  m.Set(v)
  if m != maybe.NewInt16(v) {
    t.Error("Int16 Set broken.")
  }
  verifyString(t, "Just 7", m.String())
  m.Clear()
  if m != c {
    t.Error("Int16 Clear broken.")
  }
  verifyString(t, "Nothing", m.String())
}

func TestMaybeInt32(t *testing.T) {
  var m, c maybe.Int32
  var v int32 = 7
  m.Set(v)
  if m != maybe.NewInt32(v) {
    t.Error("Int32 Set broken.")
  }
  verifyString(t, "Just 7", m.String())
  m.Clear()
  if m != c {
    t.Error("Int32 Clear broken.")
  }
  verifyString(t, "Nothing", m.String())
}

func TestMaybeInt64(t *testing.T) {
  var m, c maybe.Int64
  var v int64 = 7
  m.Set(v)
  if m != maybe.NewInt64(v) {
    t.Error("Int64 Set broken.")
  }
  verifyString(t, "Just 7", m.String())
  m.Clear()
  if m != c {
    t.Error("Int64 Clear broken.")
  }
  verifyString(t, "Nothing", m.String())
}

func TestMaybeUint(t *testing.T) {
  var m, c maybe.Uint
  var v uint = 7
  m.Set(v)
  if m != maybe.NewUint(v) {
    t.Error("Uint Set broken.")
  }
  verifyString(t, "Just 7", m.String())
  m.Clear()
  if m != c {
    t.Error("Uint Clear broken.")
  }
  verifyString(t, "Nothing", m.String())
}

func TestMaybeUint8(t *testing.T) {
  var m, c  maybe.Uint8
  var v uint8 = 7
  m.Set(v)
  if m != maybe.NewUint8(v) {
    t.Error("Uint8 Set broken.")
  }
  verifyString(t, "Just 7", m.String())
  m.Clear()
  if m != c {
    t.Error("Uint8 Clear broken.")
  }
  verifyString(t, "Nothing", m.String())
}

func TestMaybeUint16(t *testing.T) {
  var m, c  maybe.Uint16
  var v uint16 = 7
  m.Set(v)
  if m != maybe.NewUint16(v) {
    t.Error("Uint16 Set broken.")
  }
  verifyString(t, "Just 7", m.String())
  m.Clear()
  if m != c {
    t.Error("Uint16 Clear broken.")
  }
  verifyString(t, "Nothing", m.String())
}

func TestMaybeUint32(t *testing.T) {
  var m, c  maybe.Uint32
  var v uint32 = 7
  m.Set(v)
  if m != maybe.NewUint32(v) {
    t.Error("Uint32 Set broken.")
  }
  verifyString(t, "Just 7", m.String())
  m.Clear()
  if m != c {
    t.Error("Uint32 Clear broken.")
  }
  verifyString(t, "Nothing", m.String())
}

func TestMaybeUint64(t *testing.T) {
  var m, c  maybe.Uint64
  var v uint64 = 7
  m.Set(v)
  if m != maybe.NewUint64(v) {
    t.Error("Uint64 Set broken.")
  }
  verifyString(t, "Just 7", m.String())
  m.Clear()
  if m != c {
    t.Error("Uint64 Clear broken.")
  }
  verifyString(t, "Nothing", m.String())
}

func TestMaybeFloat32(t *testing.T) {
  var m, c  maybe.Float32
  var v float32 = 7.0
  m.Set(v)
  if m != maybe.NewFloat32(v) {
    t.Error("Float32 Set broken.")
  }
  verifyString(t, "Just 7.000000", m.String())
  m.Clear()
  if m != c {
    t.Error("Float32 Clear broken.")
  }
  verifyString(t, "Nothing", m.String())
}

func TestMaybeFloat64(t *testing.T) {
  var m, c  maybe.Float64
  var v float64 = 7.0
  m.Set(v)
  if m != maybe.NewFloat64(v) {
    t.Error("Float64 Set broken.")
  }
  verifyString(t, "Just 7.000000", m.String())
  m.Clear()
  if m != c {
    t.Error("Float64 Clear broken.")
  }
  verifyString(t, "Nothing", m.String())
}

func TestMaybeBool(t *testing.T) {
  var m, c  maybe.Bool
  var v bool = true
  m.Set(v)
  if m != maybe.NewBool(v) {
    t.Error("Bool Set broken.")
  }
  verifyString(t, "Just true", m.String())
  m.Clear()
  if m != c {
    t.Error("Bool Clear broken.")
  }
  verifyString(t, "Nothing", m.String())
}

func TestMaybeString(t *testing.T) {
  var m, c  maybe.String
  var v string = "hello"
  m.Set(v)
  if m != maybe.NewString(v) {
    t.Error("String Set broken.")
  }
  verifyString(t, "Just hello", m.String())
  m.Clear()
  if m != c {
    t.Error("String Clear broken.")
  }
  verifyString(t, "Nothing", m.String())
}

func verifyString(t *testing.T, expected, actual string) {
  if expected != actual {
    t.Errorf("Expected %s, got %s", expected, actual);
  }
}
