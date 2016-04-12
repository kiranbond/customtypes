// Copyright 2015 ZeroStack, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// RawMessage is a data type that can be used to partially unmarshal a JSON
// struct. Adding additional methods on []byte provides a mechanism to
// work with partially unmarshaled JSON that is then unmarshaled on demand into
// intended data type at the actual usage. This is used in the invoke package
// at github.com/zerostackinc/invoke

package customtypes

import "errors"

// RawMessage is a raw encoded JSON object.
// It implements json.Marshaler and json.Unmarshaler like json.RawMessage,
// but also proto.Marshaler and proto.Unmarshaler.
type RawMessage []byte

// Size returns size of the RawMessage
func (m *RawMessage) Size() int {
  return len((*m))
}

// Equal checks the two RawMessages are equal
func (m *RawMessage) Equal(data RawMessage) bool {
  if len(*m) != len(data) {
    return false
  }
  for ii, b := range *m {
    if b != data[ii] {
      return false
    }
  }
  return true
}

// MarshalJSON returns *m as the JSON encoding of m.
func (m *RawMessage) MarshalJSON() ([]byte, error) {
  return *m, nil
}

// UnmarshalJSON sets *m to a copy of data.
func (m *RawMessage) UnmarshalJSON(data []byte) error {
  if m == nil {
    return errors.New("RawMessage: UnmarshalJSON on nil pointer")
  }
  *m = append((*m)[0:0], data...)
  return nil
}

// Marshal implements proto.Marshaler.
func (m *RawMessage) Marshal() ([]byte, error) {
  return *m, nil
}

// MarshalTo implements proto.Marshaler.
func (m *RawMessage) MarshalTo(data []byte) (int, error) {
  data = *m
  return len(*m), nil
}

// Unmarshal implements proto.Unmarshaler.
func (m *RawMessage) Unmarshal(data []byte) error {
  if m == nil {
    return errors.New("RawMessage: Unmarshal on nil pointer")
  }
  *m = append((*m)[0:0], data...)
  return nil
}

// String implements a stringer method for printing.
func (m *RawMessage) String() string {
  if m == nil {
    return ""
  }
  return string(*m)
}
