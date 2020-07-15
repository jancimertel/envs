package envs

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestParseBool(t *testing.T) {
	validIns := []string{
		"true",
		"TRUE",
		"True",
		"TruE",
		"1",
		"false",
		"FALSE",
		"False",
		"FalSE",
		"0",
	}
	validOuts := []bool{
		true,
		true,
		true,
		true,
		true,
		false,
		false,
		false,
		false,
		false,
	}
	invalidIns := []string{
		"",
		"truer",
		"random",
		"3",
	}
	for index, in := range validIns {
		out, err := parseBool(in)
		assert.Equal(t, out, validOuts[index], "Valid outs should be equal")
		assert.Nil(t, err, "Error should be nil")
	}
	for _, in := range invalidIns {
		out, err := parseBool(in)
		assert.Equal(t, out, false, "Invalid our shoudl always return false")
		assert.Error(t, err, "Error should not be nil")
	}
}

func TestMustHaveBool(t *testing.T) {
	type test struct {
		BoolTrue bool
		BoolFalse bool
	}
	_ = os.Setenv("BoolTrue", "1")
	_ = os.Setenv("BoolFalse", "0")

	inst := &test{}
	assert.Nil(t, MustHave(inst))
	assert.True(t, inst.BoolTrue)
	assert.False(t, inst.BoolFalse)
}

func TestMustHaveString(t *testing.T) {
	wantedString := "wanted"
	type test struct {
		String string
	}
	_ = os.Setenv("String", wantedString)

	inst := &test{}
	assert.Nil(t, MustHave(inst))
	assert.Equal(t, wantedString, inst.String)
}

func TestMustHaveInt(t *testing.T) {
	wantedNumber := 9
	type test struct {
		Int int
		Int8 int8
		Int16 int16
		Int32 int32
		Int64 int64
	}
	_ = os.Setenv("Int", fmt.Sprintf("%d", wantedNumber))
	_ = os.Setenv("Int8", fmt.Sprintf("%d", wantedNumber))
	_ = os.Setenv("Int16", fmt.Sprintf("%d", wantedNumber))
	_ = os.Setenv("Int32", fmt.Sprintf("%d", wantedNumber))
	_ = os.Setenv("Int64", fmt.Sprintf("%d", wantedNumber))

	inst := &test{}
	assert.Nil(t, MustHave(inst))
	assert.Equal(t, wantedNumber, inst.Int)
	assert.Equal(t, int8(wantedNumber), inst.Int8)
	assert.Equal(t, int16(wantedNumber), inst.Int16)
	assert.Equal(t, int32(wantedNumber), inst.Int32)
	assert.Equal(t, int64(wantedNumber), inst.Int64)
}

func TestMustHaveUint(t *testing.T) {
	wantedNumber := uint(9)
	type test struct {
		Uint uint
		Uint8 uint8
		Uint16 uint16
		Uint32 uint32
		Uint64 uint64
	}
	_ = os.Setenv("Uint", fmt.Sprintf("%d", wantedNumber))
	_ = os.Setenv("Uint8", fmt.Sprintf("%d", wantedNumber))
	_ = os.Setenv("Uint16", fmt.Sprintf("%d", wantedNumber))
	_ = os.Setenv("Uint32", fmt.Sprintf("%d", wantedNumber))
	_ = os.Setenv("Uint64", fmt.Sprintf("%d", wantedNumber))

	inst := &test{}
	assert.Nil(t, MustHave(inst))
	assert.Equal(t, wantedNumber, inst.Uint)
	assert.Equal(t, uint8(wantedNumber), inst.Uint8)
	assert.Equal(t, uint16(wantedNumber), inst.Uint16)
	assert.Equal(t, uint32(wantedNumber), inst.Uint32)
	assert.Equal(t, uint64(wantedNumber), inst.Uint64)
}

func TestMustHaveFloat(t *testing.T) {
	wantedNumber := float32(9.5)
	type test struct {
		Float32 float32
		Float64 float64
	}
	_ = os.Setenv("Float32", fmt.Sprintf("%f", wantedNumber))
	_ = os.Setenv("Float64", fmt.Sprintf("%f", wantedNumber))

	inst := &test{}
	assert.Nil(t, MustHave(inst))
	assert.Equal(t, wantedNumber, inst.Float32)
	assert.Equal(t, float64(wantedNumber), inst.Float64)
}

func TestMustHaveStruct(t *testing.T) {
	type test struct {
	}

	inst := test{}
	assert.Error(t, MustHave(inst))
}
