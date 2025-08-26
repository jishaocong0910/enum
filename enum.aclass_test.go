package e

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type Animal struct {
	*EnumElem__
}

func TestEnum(t *testing.T) {
	enum := NewEnum[Animal](struct {
		*Enum__[Animal]
		cat, DOG, bird Animal
	}{
		cat:  Animal{},
		DOG:  Animal{},
		bird: Animal{},
	})

	r := require.New(t)
	r.True(enum.Is(Animal{}, Animal{}))
	r.True(enum.Is(enum.OfString("cat"), enum.cat))
	r.True(enum.Is(enum.OfString("DOG"), enum.DOG))
	r.True(enum.Is(enum.OfString("bird"), enum.bird))
	r.Equal("cat", enum.cat.String())
	r.Equal("DOG", enum.DOG.String())
	r.Equal("bird", enum.bird.String())
	r.True(enum.Is(enum.cat, enum.cat, enum.DOG))
	r.False(enum.Is(enum.cat, enum.bird, enum.DOG))
	r.True(enum.Not(enum.cat, enum.bird, enum.DOG))
	r.False(enum.Not(enum.cat, enum.cat, enum.DOG))
	r.True(enum.OfString("SNAKE").IsUndefined())
	r.True(enum.OfStringIgnoreCase("SNAKE").IsUndefined())
	r.True(enum.OfString("BIRD").IsUndefined())
	r.False(enum.OfStringIgnoreCase("BIRD").IsUndefined())
	r.True(enum.Undefined().IsUndefined())

	values := enum.Elems()
	r.Len(values, 3)

	a := Animal{}
	r.Equal("<undefined>", a.String())
}

func TestEnumPanic(t *testing.T) {
	type animals_ struct {
		*Enum__[Animal]
		CAT  Animal
		DOG  Animal
		BIRD *Animal
	}

	r := require.New(t)
	r.PanicsWithValue("parameter \"e\" must be a struct value", func() {
		NewEnum[Animal](&animals_{})
	})

	r.PanicsWithValue("e.animals_.BIRD must be a struct value", func() {
		NewEnum[Animal](animals_{
			CAT:  Animal{},
			DOG:  Animal{},
			BIRD: &Animal{},
		})
	})
}
