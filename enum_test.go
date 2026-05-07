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
	"testing"

	"github.com/stretchr/testify/require"
)

type Animal struct {
	EnumElem
}

type _Animal struct {
	Enum[Animal]
	cat,
	DOG,
	bird Animal
}

var Animal_ = NewEnum(_Animal{})

type Animal2 struct {
	EnumElem
}

type _Animal2 struct {
	Enum[Animal2]
	cat,
	DOG Animal2
}

var Animal2_ = NewEnum(_Animal2{})

func TestEnum(t *testing.T) {

	r := require.New(t)
	r.True(Animal{}.Is(Animal{}))
	r.True(Animal_.UNDEFINED.Is(Animal{}))
	r.True(Animal_.OfString("cat").Is(Animal_.cat))
	r.False(Animal_.OfString("cat").Is(Animal2_.cat))
	r.True(Animal_.OfString("DOG").Is(Animal_.DOG))
	r.False(Animal_.OfString("DOG").Is(Animal2_.DOG))
	r.True(Animal_.OfString("bird").Is(Animal_.bird))
	r.Equal("cat", Animal_.cat.String())
	r.Equal("DOG", Animal_.DOG.String())
	r.Equal("bird", Animal_.bird.String())
	r.True(Animal_.cat.Is(Animal_.cat, Animal_.DOG))
	r.False(Animal_.cat.Is(Animal_.bird, Animal_.DOG))
	r.True(Animal_.cat.Not(Animal_.bird, Animal_.DOG))
	r.False(Animal_.cat.Not(Animal_.cat, Animal_.DOG))
	r.False(Animal_.OfString("SNAKE").IsPresent())
	r.True(Animal_.OfString("SNAKE").IsUndefined())
	r.False(Animal_.OfStringCI("SNAKE").IsPresent())
	r.True(Animal_.OfStringCI("SNAKE").IsUndefined())
	r.False(Animal_.OfString("BIRD").IsPresent())
	r.True(Animal_.OfString("BIRD").IsUndefined())
	r.True(Animal_.OfStringCI("BIRD").IsPresent())
	r.False(Animal_.OfStringCI("BIRD").IsUndefined())
	r.True(Animal_.UNDEFINED.Is(Animal_.UNDEFINED))

	values := Animal_.Elems()
	r.Len(values, 3)

	a := Animal{}
	r.Equal("<undefined>", a.String())
}
