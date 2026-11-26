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

package enum

import (
	"reflect"
	"slices"
	"strings"
	"unsafe"
)

// Val is the base type embedded in each enum value.
// NewType initializes its name and ID from the declaring field.
// The enum value type and its other embedded types must not declare a Name field
// that conflicts with Val.Name.
type Val struct {
	// ID identifies the enum value by package path, value type name, and field name.
	ID string
	// Name is the field name used to declare the enum value.
	Name string
}

func (v Val) enumValue() {}

func (v Val) id() string {
	return v.ID
}

// String returns Name, or "<undefined>" if Name is empty.
func (v Val) String() string {
	if v.Name == "" {
		return "<undefined>"
	}
	return v.Name
}

// Is reports whether any target is an enum value with the same ID.
// Nil pointers and targets that are not enum values are ignored.
func (v Val) Is(targets ...any) bool {
	for _, t := range targets {
		if e, ok := t.(enumValue); ok {
			if rv := reflect.ValueOf(e); rv.Kind() == reflect.Pointer && rv.IsNil() {
				continue
			}
			if v.ID == e.id() {
				return true
			}
		}
	}
	return false
}

// Not reports whether none of the targets match the value according to Is.
func (v Val) Not(targets ...any) bool {
	return !v.Is(targets...)
}

// IsUndefined reports whether ID is empty, as it is for a zero value.
func (v Val) IsUndefined() bool {
	return v.ID == ""
}

// IsPresent reports whether ID is nonempty.
func (v Val) IsPresent() bool {
	return !v.IsUndefined()
}

// Typ provides access to the values of an enum type and lookup by name.
// Embed Typ[V] in a struct with fields of type V, then initialize it with NewType.
type Typ[V enumValue] struct {
	values []V
	// Names returns the declared enum value names in declaration order.
	Names Keys[string]
	// OfName returns the enum value with the given name, or the zero value of V if absent.
	OfName Dir[string, V]
	// OfNameIgnoreCase returns the enum value with the given name, or the zero value of V if absent.
	// It matches names after converting them to lowercase. If multiple names match,
	// it returns the first enum value in declaration order.
	OfNameIgnoreCase DirCI[V]
	// UNDEFINED is the zero value of V, representing an undefined enum value.
	UNDEFINED V
}

func (t Typ[V]) enumValue(V) {}

// Values returns the declared enum values in declaration order, excluding UNDEFINED.
func (t Typ[V]) Values() []V {
	return slices.Clone(t.values)
}

// Keys returns the unique field values of an enum's values in declaration order.
type Keys[K comparable] func() []K

// Dir looks up an enum value by a unique field value and returns the zero value if not found.
type Dir[K comparable, V enumValue] func(k K) V

// DirCI looks up an enum value by a string field value and returns the zero value if not found.
// It matches values after converting them to lowercase. If multiple values match,
// it returns the first enum value in declaration order.
type DirCI[V enumValue] func(k string) V

// NewType initializes and returns an enum type, assigning names and IDs to its values.
// T must be a struct embedding Typ[V], and V must be a struct embedding Val.
// Fields of type V declare enum values and may be exported or unexported.
// Custom fields on those values are preserved.
func NewType[T enumType[V], V enumValue](t T) T {
	tType := reflect.TypeFor[T]()
	tValue := reflect.ValueOf(&t).Elem()
	vType := reflect.TypeFor[V]()
	vTypeName := vType.PkgPath() + "." + vType.Name()
	values := _buildValues[V](tType, tValue, vType, vTypeName)
	return _buildType[T](tValue, values)
}

type enumValue interface {
	enumValue()
	id() string
}

type enumType[V enumValue] interface {
	enumValue(V)
}

func _buildValues[V enumValue](tType reflect.Type, tValue reflect.Value, vType reflect.Type, vTypeName string) []V {
	values := make([]V, 0, tValue.NumField()-1)
	for i := 0; i < tValue.NumField(); i++ {
		if v, ok := _buildValue[V](tType.Field(i), tValue.Field(i), vType, vTypeName); ok {
			values = append(values, v)
		}
	}
	return values
}

func _buildValue[V enumValue](field reflect.StructField, fieldValue reflect.Value, vType reflect.Type, vTypeName string) (v V, ok bool) {
	if !field.Type.AssignableTo(vType) {
		return
	}
	if field.IsExported() {
		fieldValue.FieldByName("Val").Set(reflect.ValueOf(Val{ID: vTypeName + "." + field.Name, Name: field.Name}))
		v = fieldValue.Interface().(V)
	} else {
		*(*Val)(unsafe.Pointer(fieldValue.FieldByName("Val").UnsafeAddr())) = Val{ID: vTypeName + "." + field.Name, Name: field.Name}
		v = *(*V)(unsafe.Pointer(fieldValue.UnsafeAddr()))
	}
	return v, true
}

func _buildType[T enumType[V], V enumValue](etValue reflect.Value, values []V) T {
	*(*Typ[V])(unsafe.Pointer(etValue.FieldByName("Typ").UnsafeAddr())) = Typ[V]{
		values:           values,
		Names:            _buildKeys[string](values, "Name"),
		OfName:           _buildDir[string](values, "Name"),
		OfNameIgnoreCase: _buildDirCI(values, "Name"),
	}
	return etValue.Interface().(T)
}

func _buildKeys[K comparable, V enumValue](values []V, fieldName string) Keys[K] {
	if field, ok := reflect.TypeFor[V]().FieldByName(fieldName); ok {
		if field.Type != reflect.TypeFor[K]() {
			panic("类型不匹配")
		}
	} else {
		panic("字段找不到")
	}

	exists := make(map[K]V, len(values))
	keys := make([]K, 0, len(values))
	for _, v := range values {
		key := reflect.ValueOf(v).FieldByName(fieldName).Interface().(K)
		if _, ok := exists[key]; !ok {
			exists[key] = v
			keys = append(keys, key)
		} else {
			panic("重复值")
		}
	}

	return func() []K {
		return slices.Clone(keys)
	}
}

func _buildDir[K comparable, V enumValue](values []V, fieldName string) Dir[K, V] {
	if field, ok := reflect.TypeFor[V]().FieldByName(fieldName); ok {
		if field.Type != reflect.TypeFor[K]() {
			panic("类型不匹配")
		}
	} else {
		panic("字段找不到")
	}

	dir := make(map[K]V, len(values))
	for _, v := range values {
		key := reflect.ValueOf(v).FieldByName(fieldName).Interface().(K)
		if _, ok := dir[key]; !ok {
			dir[key] = v
		} else {
			panic("重复值")
		}
	}

	return func(k K) V {
		var v V
		if result, ok := dir[k]; ok {
			v = result
		}
		return v
	}
}

func _buildDirCI[V enumValue](values []V, fieldName string) DirCI[V] {
	if field, ok := reflect.TypeFor[V]().FieldByName(fieldName); ok {
		if field.Type != reflect.TypeFor[string]() {
			panic("必须是string")
		}
	} else {
		panic("字段找不到")
	}

	dir := make(map[string]V, len(values))
	for _, v := range values {
		key := reflect.ValueOf(v).FieldByName(fieldName).Interface().(string)
		key = strings.ToLower(key)
		if _, ok := dir[key]; !ok {
			dir[key] = v
		}
	}

	return func(k string) V {
		var v V
		if result, ok := dir[strings.ToLower(k)]; ok {
			v = result
		}
		return v
	}
}
