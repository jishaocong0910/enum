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
	r.Equal("cat", Animal_.cat.Name())
	r.NotEqual("Cat", Animal_.cat.Name())
	r.True(Animal_.GetByName("cat").Is(Animal_.cat))
	r.False(Animal_.GetByName("cat").Is(Animal2_.cat))
	r.True(Animal_.GetByName("DOG").Is(Animal_.DOG))
	r.False(Animal_.GetByName("DOG").Is(Animal2_.DOG))
	r.True(Animal_.GetByName("bird").Is(Animal_.bird))
	r.Equal("cat", Animal_.cat.String())
	r.Equal("DOG", Animal_.DOG.String())
	r.Equal("bird", Animal_.bird.String())
	r.True(Animal_.cat.Is(Animal_.cat, Animal_.DOG))
	r.False(Animal_.cat.Is(Animal_.bird, Animal_.DOG))
	r.True(Animal_.cat.Not(Animal_.bird, Animal_.DOG))
	r.False(Animal_.cat.Not(Animal_.cat, Animal_.DOG))
	r.False(Animal_.GetByName("SNAKE").IsPresent())
	r.True(Animal_.GetByName("SNAKE").IsUndefined())
	r.False(Animal_.GetByNameCI("SNAKE").IsPresent())
	r.True(Animal_.GetByNameCI("SNAKE").IsUndefined())
	r.False(Animal_.GetByName("BIRD").IsPresent())
	r.True(Animal_.GetByName("BIRD").IsUndefined())
	r.True(Animal_.GetByNameCI("BIRD").IsPresent())
	r.False(Animal_.GetByNameCI("BIRD").IsUndefined())
	r.True(Animal_.UNDEFINED.Is(Animal_.UNDEFINED))
	r.Equal([]string{"cat", "DOG", "bird"}, Animal_.Names())

	values := Animal_.Elems()
	r.Len(values, 3)

	a := Animal{}
	r.Equal("<undefined>", a.String())
}
