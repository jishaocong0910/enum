// Copyright 2024-present jishaocong0910
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

package e

type iEnumElem interface {
	iEnumElem()
	id() string

	Name() string
	IsUndefined() bool
}

// EnumElem is the base type embedded in each enum element.
// It stores the type name and field name of the element.
type EnumElem struct {
	// the unique key of the enum element.
	ID   string
	name string
}

func (el EnumElem) iEnumElem() {}

func (el EnumElem) id() string {
	return el.ID
}

// String returns the string representation of the enum element.
func (el EnumElem) String() string {
	if el.name == "" {
		return "<undefined>"
	}
	return el.name
}

// Name returns the field name of the enum element.
func (el EnumElem) Name() string {
	return el.name
}

// Is reports whether the element equals any of the targets,
// matched by both field name and type name.
func (el EnumElem) Is(targets ...any) bool {
	for _, t := range targets {
		if e, ok := t.(iEnumElem); ok {
			if el.ID == e.id() {
				return true
			}
		}
	}
	return false
}

// Not reports whether the element does NOT equal any of the targets
// (inverse of Is).
func (el EnumElem) Not(targets ...any) bool {
	return !el.Is(targets...)
}

// IsUndefined reports whether the element represents an undefined value
// (i.e., it was not properly initialized).
func (el EnumElem) IsUndefined() bool {
	return el.ID == ""
}

// IsPresent reports whether the element is defined (not undefined).
func (el EnumElem) IsPresent() bool {
	return !el.IsUndefined()
}
