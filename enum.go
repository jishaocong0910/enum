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

import (
	"reflect"
	"strings"
	"unsafe"
)

type iEnum[EL iEnumElem] interface {
	iEnum(EL)
}

// Enum is the enum collection that holds all enum elements
// and provides lookup operations. Embed it in your enum struct along
// with typed element fields, then initialize with NewEnum.
type Enum[EL iEnumElem] struct {
	elems       []EL
	idMap       map[string]EL
	elemNames   []string
	elemNameMap map[string]EL

	// UNDEFINED is the built-in zero-value element representing an undefined/missing enum value.
	UNDEFINED EL
}

func (e Enum[EL]) iEnum(EL) {}

// Elems returns all defined enum elements.
func (e Enum[EL]) Elems() []EL {
	return e.elems
}

// Names returns the names of all defined enum elements.
func (e Enum[EL]) Names() []string {
	return e.elemNames
}

// GetByName finds the enum element by its field name string.
// It returns the UNDEFINED element if not match.
func (e Enum[EL]) GetByName(str string) (el EL) {
	if v, ok := e.elemNameMap[str]; ok {
		return v
	}
	return
}

// GetByNameCI finds the enum element by its field name string (case-insensitive).
// It returns the UNDEFINED element if not match.
func (e Enum[EL]) GetByNameCI(str string) (el EL) {
	for _, v := range e.elems {
		if strings.EqualFold(v.Name(), str) {
			return v
		}
	}
	return
}

// NewEnum creates and initializes an enum collection. Parameter E must be a struct that
// embeds Enum[EL] and contains fields of type EL (exported or unexported).
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
			vf.FieldByName("EnumElem").Set(reflect.ValueOf(EnumElem{id: typeName + "." + tf.Name, name: tf.Name}))
			elems = append(elems, vf.Interface().(EL))
		} else {
			*(*EnumElem)(unsafe.Pointer(vf.FieldByName("EnumElem").UnsafeAddr())) = EnumElem{id: typeName + "." + tf.Name, name: tf.Name}
			elems = append(elems, *(*EL)(unsafe.Pointer(vf.UnsafeAddr())))
		}
	}

	idMap := make(map[string]EL, len(elems))
	elemNames := make([]string, 0, len(elems))
	elemNameMap := make(map[string]EL, len(elems))
	for _, elem := range elems {
		idMap[elem.ID()] = elem
		name := elem.Name()
		elemNames = append(elemNames, name)
		elemNameMap[name] = elem
	}

	*(*Enum[EL])(unsafe.Pointer(v.FieldByName("Enum").UnsafeAddr())) = Enum[EL]{
		elems:       elems,
		elemNames:   elemNames,
		elemNameMap: elemNameMap,
	}

	return v.Interface().(E)
}
