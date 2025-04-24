package defaults

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type InnerTypes struct {
	Bool   bool   `default:"true"`
	String string `default:"hello"`
}

type Types struct {
	Bool       bool       `default:"true"`
	Int8       int8       `default:"10"`
	Int16      int16      `default:"10"`
	Int32      int32      `default:"10"`
	Int64      int64      `default:"10"`
	Int        int        `default:"10"`
	Uint8      uint8      `default:"10"`
	Uint16     uint16     `default:"10"`
	Uint32     uint32     `default:"10"`
	Uint64     uint64     `default:"10"`
	Uint       uint       `default:"10"`
	Uintptr    uintptr    `default:"10"`
	Float32    float32    `default:"10.1"`
	Float64    float64    `default:"10.1"`
	Complex64  complex64  `default:"1+2i"`
	Complex128 complex128 `default:"1+2i"`
	String     string     `default:"equal"`
	InnerTypes *InnerTypes
}

func TestTypes(t *testing.T) {
	b := &Types{}
	assert.Nil(t, Default(b))

	table := [][2]any{
		{b.Bool, true},
		{b.Int8, int8(10)},
		{b.Int16, int16(10)},
		{b.Int32, int32(10)},
		{b.Int64, int64(10)},
		{b.Int, int(10)},
		{b.Uint8, uint8(10)},
		{b.Uint16, uint16(10)},
		{b.Uint32, uint32(10)},
		{b.Uint64, uint64(10)},
		{b.Uint, uint(10)},
		{b.Uintptr, uintptr(10)},
		{b.Float32, float32(10.1)},
		{b.Float64, float64(10.1)},
		{b.Complex64, complex64(1 + 2i)},
		{b.Complex128, complex128(1 + 2i)},
		{b.String, "equal"},
		{b.InnerTypes.String, "hello"},
		{b.InnerTypes.Bool, true},
	}

	for x, item := range table {
		assert.Equal(t, item[0], item[1], "table item %d was not equal", x)
	}
}
