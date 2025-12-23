package e

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type Animal struct {
	*EnumElem__
}

func TestEnum(t *testing.T) {
	enum := NewEnum(struct {
		*Enum__[Animal]
		cat, DOG, bird Animal
	}{
		cat:  Animal{},
		DOG:  Animal{},
		bird: Animal{},
	})

	r := require.New(t)
	r.True(Animal{}.Is(Animal{}))
	r.True(enum.Undefined().Is(enum.Undefined()))
	r.True(enum.OfString("cat").Is(enum.cat))
	r.True(enum.OfString("DOG").Is(enum.DOG))
	r.True(enum.OfString("bird").Is(enum.bird))
	r.Equal("cat", enum.cat.String())
	r.Equal("DOG", enum.DOG.String())
	r.Equal("bird", enum.bird.String())
	r.True(enum.cat.Is(enum.cat, enum.DOG))
	r.False(enum.cat.Is(enum.bird, enum.DOG))
	r.True(enum.cat.Not(enum.bird, enum.DOG))
	r.False(enum.cat.Not(enum.cat, enum.DOG))
	r.True(enum.OfString("SNAKE").IsUndefined())
	r.True(enum.OfStringCI("SNAKE").IsUndefined())
	r.True(enum.OfString("BIRD").IsUndefined())
	r.False(enum.OfStringCI("BIRD").IsUndefined())
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
		NewEnum(&animals_{})
	})

	r.PanicsWithValue("e.animals_.BIRD must be a struct value", func() {
		NewEnum(animals_{
			CAT:  Animal{},
			DOG:  Animal{},
			BIRD: &Animal{},
		})
	})
}
