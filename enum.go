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
	OfName Directory[string, V]
	// OfNameIgnoreCase returns the enum value with the given name, or the zero value of V if absent.
	// It matches names after converting them to lowercase. If multiple names match,
	// it returns the first enum value in declaration order.
	OfNameIgnoreCase DirectoryCI[V]
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

// Directory looks up an enum value by a unique field value and returns the zero value if not found.
type Directory[K comparable, V enumValue] func(k K) V

// DirectoryCI looks up an enum value by a string field value and returns the zero value if not found.
// It matches values after converting them to lowercase. If multiple values match,
// it returns the first enum value in declaration order.
type DirectoryCI[V enumValue] func(k string) V

// NewType initializes and returns an enum type, assigning names and IDs to its values.
// T must be a struct embedding Typ[V], and V must be a struct embedding Val.
// Fields of type V declare enum values and may be exported or unexported.
// Custom fields on those values are preserved.
func NewType[T enumType[V], V enumValue](t T) T {
	tType := reflect.TypeFor[T]()
	tValue := reflect.ValueOf(&t).Elem()
	vType := reflect.TypeFor[V]()
	vTypeName := vType.PkgPath() + "." + vType.Name()
	values := initEmbedVals[V](tType, tValue, vType, vTypeName)
	initEmbedTyp(tValue, values)
	initTypeFields(tType, tValue, values)
	return tValue.Interface().(T)
}

type enumValue interface {
	enumValue()
	id() string
}

type enumType[V enumValue] interface {
	enumValue(V)
}

func initEmbedVals[V enumValue](tType reflect.Type, tValue reflect.Value, vType reflect.Type, vTypeName string) []V {
	values := make([]V, 0, tValue.NumField()-1)
	for i := 0; i < tValue.NumField(); i++ {
		field := tType.Field(i)
		if !field.Type.AssignableTo(vType) {
			continue
		}
		fieldValue := tValue.Field(i)
		val := Val{ID: vTypeName + "." + field.Name, Name: field.Name}
		var v V
		if field.IsExported() {
			fieldValue.FieldByName("Val").Set(reflect.ValueOf(val))
			v = fieldValue.Interface().(V)
		} else {
			*(*Val)(unsafe.Pointer(fieldValue.FieldByName("Val").UnsafeAddr())) = val
			v = *(*V)(unsafe.Pointer(fieldValue.UnsafeAddr()))
		}
		values = append(values, v)
	}
	return values
}

func initEmbedTyp[V enumValue](tValue reflect.Value, values []V) {
	names := make([]string, 0, len(values))
	for _, v := range values {
		names = append(names, reflect.ValueOf(v).FieldByName("Val").FieldByName("Name").Interface().(string))
	}

	tValue.FieldByName("Typ").Set(reflect.ValueOf(Typ[V]{
		values: values,
		Names: func() Keys[string] {
			return func() []string {
				return slices.Clone(names)
			}
		}(),
		OfName: func() Directory[string, V] {
			dir := make(map[string]V, len(values))
			for i, name := range names {
				dir[name] = values[i]
			}
			return func(k string) V {
				var v V
				if result, ok := dir[k]; ok {
					v = result
				}
				return v
			}
		}(),
		OfNameIgnoreCase: func() DirectoryCI[V] {
			dir := make(map[string]V, len(values))
			for i, name := range names {
				name = strings.ToLower(name)
				if _, ok := dir[name]; !ok {
					dir[name] = values[i]
				}
			}
			return func(k string) V {
				var v V
				if result, ok := dir[strings.ToLower(k)]; ok {
					v = result
				}
				return v
			}
		}(),
	}))
}

func initTypeFields[V enumValue](tType reflect.Type, tValue reflect.Value, values []V) {
	for i := 0; i < tValue.NumField(); i++ {
		field := tType.Field(i)
		if field.Anonymous {
			continue
		}
		fieldType := field.Type
		if fieldType.Kind() == reflect.Pointer {
			if isKeysType(fieldType.Elem()) {
				panic("enum.Keys不能支持指针")
			}
		}
		if !isKeysType(fieldType) {
			continue
		}
		buildKeys[string, V](values, field.Name)
	}
}

func isKeysType(t reflect.Type) bool {
	if t.PkgPath() == "github.com/jishaocong0910/enum" && strings.HasPrefix(t.Name(), "Keys[") {
		return true
	}
	return false
}

func buildKeys[K comparable, V enumValue](values []V, fieldName string) Keys[K] {
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

func buildDirectory[K comparable, V enumValue](values []V, fieldName string) Directory[K, V] {
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

func buildDirectoryCI[V enumValue](values []V, fieldName string) DirectoryCI[V] {
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
