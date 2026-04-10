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
	getFieldName() string

	String() string
	IsUndefined() bool
}

type EnumElem struct {
	fieldName string
}

func (el EnumElem) iEnumElem() {}

func (el EnumElem) getFieldName() string {
	if el.fieldName == "" {
		return undefined
	}
	return el.fieldName
}

// String 返回枚举值字符串形式，与枚举集合中的字段名相同，因此具有唯一性。
func (el EnumElem) String() string {
	return el.getFieldName()
}

// Is 判断是否存在指定枚举值
func (el EnumElem) Is(targets ...any) bool {
	for _, t := range targets {
		if e, ok := t.(iEnumElem); ok {
			if el.getFieldName() == e.getFieldName() {
				return true
			}
		}
	}
	return false
}

// Not 与Is方法相反
func (el EnumElem) Not(targets ...any) bool {
	return !el.Is(targets...)
}

// IsUndefined 是否未定义的枚举
func (el EnumElem) IsUndefined() bool {
	return el.getFieldName() == undefined
}
