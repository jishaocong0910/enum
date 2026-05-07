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

import (
	"reflect"
	"strings"
	"unsafe"
)

type iEnum[EL iEnumElem] interface {
	iEnum(EL)
}

type Enum[EL iEnumElem] struct {
	elems        []EL
	fieldNames   []string
	fieldNameMap map[string]EL

	// Built-in undefined enum
	UNDEFINED EL
}

func (e Enum[EL]) iEnum(EL) {}

// Elems returns all enum values.
func (e Enum[EL]) Elems() []EL {
	return e.elems
}

// Strings returns all enum value strings.
func (e Enum[EL]) Strings() []string {
	return e.fieldNames
}

// OfString finds the enum value by string.
func (e Enum[EL]) OfString(str string) (el EL) {
	if v, ok := e.fieldNameMap[str]; ok {
		return v
	}
	return
}

// OfStringCI finds the enum value by string (case-insensitive).
func (e Enum[EL]) OfStringCI(str string) (el EL) {
	for _, v := range e.elems {
		if strings.EqualFold(v.getFieldName(), str) {
			return v
		}
	}
	return
}

// NewEnum creates and initializes an enum instance.
func NewEnum[E iEnum[EL], EL iEnumElem](e E) E {
	t := reflect.TypeOf(&e).Elem()
	v := reflect.ValueOf(&e).Elem()
	elType := reflect.TypeOf((*EL)(nil)).Elem()
	typeName := elType.PkgPath() + "." + elType.Name()
	elems := make([]EL, 0, v.NumField()-1)

	for i := 0; i < v.NumField(); i++ {
		tf := t.Field(i)
		vf := v.Field(i)
		if !tf.Type.AssignableTo(elType) {
			continue
		}

		if tf.IsExported() {
			vf.FieldByName("EnumElem").Set(reflect.ValueOf(EnumElem{typeName: typeName, fieldName: tf.Name}))
			elems = append(elems, vf.Interface().(EL))
		} else {
			*(*EnumElem)(unsafe.Pointer(vf.FieldByName("EnumElem").UnsafeAddr())) = EnumElem{typeName: typeName, fieldName: tf.Name}
			elems = append(elems, *(*EL)(unsafe.Pointer(vf.UnsafeAddr())))
		}
	}

	fieldNames := make([]string, 0, len(elems))
	fieldNameMap := make(map[string]EL, len(elems))
	for _, elem := range elems {
		fieldName := elem.getFieldName()
		fieldNames = append(fieldNames, fieldName)
		fieldNameMap[fieldName] = elem
	}

	*(*Enum[EL])(unsafe.Pointer(v.FieldByName("Enum").UnsafeAddr())) = Enum[EL]{
		elems:        elems,
		fieldNames:   fieldNames,
		fieldNameMap: fieldNameMap,
	}

	return v.Interface().(E)
}
