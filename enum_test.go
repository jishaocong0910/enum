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

package enum_test

import (
	"testing"

	"github.com/jishaocong0910/enum"
	"github.com/stretchr/testify/require"
)

type Animal struct {
	enum.Val
	Code int
}

var Animal_ = enum.NewType(struct {
	enum.Typ[Animal]
	cat,
	DOG,
	bird Animal
}{})

type Animal2 struct {
	enum.Val
}

var Animal2_ = enum.NewType(struct {
	enum.Typ[Animal2]
	cat,
	DOG Animal2
}{})

func TestEnum(t *testing.T) {
	r := require.New(t)
	r.True(Animal{}.Is(Animal{}))
	r.True(Animal_.UNDEFINED.Is(Animal{}))
	r.Equal([]string{"cat", "DOG", "bird"}, Animal_.Names())
	r.Equal("cat", Animal_.cat.Name)
	r.NotEqual("Cat", Animal_.cat.Name)
	r.False(Animal_.cat.Is((*Animal)(nil)))
	r.True(Animal_.OfName("cat").Is(Animal_.cat))
	r.False(Animal_.OfName("cat").Is(Animal2_.cat))
	r.True(Animal_.OfName("DOG").Is(Animal_.DOG))
	r.False(Animal_.OfName("DOG").Is(Animal2_.DOG))
	r.True(Animal_.OfName("bird").Is(Animal_.bird))
	r.Equal("cat", Animal_.cat.String())
	r.Equal("DOG", Animal_.DOG.String())
	r.Equal("bird", Animal_.bird.String())
	r.True(Animal_.cat.Is(Animal_.cat, Animal_.DOG))
	r.False(Animal_.cat.Is(Animal_.bird, Animal_.DOG))
	r.True(Animal_.cat.Not(Animal_.bird, Animal_.DOG))
	r.False(Animal_.cat.Not(Animal_.cat, Animal_.DOG))
	r.False(Animal_.OfName("SNAKE").IsPresent())
	r.True(Animal_.OfName("SNAKE").IsUndefined())
	r.False(Animal_.OfNameIgnoreCase("SNAKE").IsPresent())
	r.True(Animal_.OfNameIgnoreCase("SNAKE").IsUndefined())
	r.False(Animal_.OfName("BIRD").IsPresent())
	r.True(Animal_.OfName("BIRD").IsUndefined())
	r.True(Animal_.OfNameIgnoreCase("BIRD").IsPresent())
	r.False(Animal_.OfNameIgnoreCase("BIRD").IsUndefined())
	r.True(Animal_.UNDEFINED.Is(Animal_.UNDEFINED))
	r.Len(Animal_.Values(), 3)
	r.Equal("<undefined>", Animal{}.String())
}

type Country struct {
	enum.Val
	CountryCode string
	CallingCode string
}

type _Country struct {
	enum.Typ[Country]
	China,
	US,
	UK Country
	CountryCodes            enum.Keys[string]               `enum:"field=CountryCode"`
	OfCountryCode           enum.Directory[string, Country] `enum:"field=CountryCode"`
	OfCountryCodeIgnoreCase enum.DirectoryCI[Country]       `enum:"field=CountryCode"`
}

func TestKeys(t *testing.T) {
	//newType := enum.NewType(_Country{
	//	China: Country{CountryCode: "CN", CallingCode: "+86"},
	//	US:    Country{CountryCode: "US", CallingCode: "+1"},
	//	UK:    Country{CountryCode: "UK", CallingCode: "+44"},
	//})
}
