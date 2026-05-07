/*
 * Copyright 2024-present jishaocong0910
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package e

const undefined = "<undefined>"

type iEnumElem interface {
	iEnumElem()
	getTypeName() string
	getFieldName() string

	String() string
	IsUndefined() bool
}

type EnumElem struct {
	typeName  string
	fieldName string
}

func (el EnumElem) iEnumElem() {}
func (el EnumElem) getTypeName() string {
	if el.typeName == "" {
		return undefined
	}
	return el.typeName
}

func (el EnumElem) getFieldName() string {
	if el.fieldName == "" {
		return undefined
	}
	return el.fieldName
}

// String returns the string representation of the enum value, which is the same as the field name in the enum collection, thus unique.
func (el EnumElem) String() string {
	return el.getFieldName()
}

// Is checks if the enum value matches any of the specified targets
func (el EnumElem) Is(targets ...any) bool {
	for _, t := range targets {
		if e, ok := t.(iEnumElem); ok {
			if el.getFieldName() == e.getFieldName() && el.getTypeName() == e.getTypeName() {
				return true
			}
		}
	}
	return false
}

// Not is the opposite of Is
func (el EnumElem) Not(targets ...any) bool {
	return !el.Is(targets...)
}

// IsPresent checks if the enum value is present (not undefined)
func (el EnumElem) IsPresent() bool {
	return el.getFieldName() != undefined
}

// IsUndefined checks if the enum value is undefined
func (el EnumElem) IsUndefined() bool {
	return el.getFieldName() == undefined
}
