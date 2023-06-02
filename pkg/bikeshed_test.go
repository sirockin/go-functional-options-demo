package buildings

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func ExampleNewBikeShed() {
	bs, _ := NewBikeShed("default")
	fmt.Println(bs)
	// output: &{default Steel 5 2 Black}
}

func ExampleNewBikeShed_second() {
	bs, _ := NewBikeShed("Lewis' Lovely Bikeshed",
		MaterialOption(RecycledCoffeeGrounds),
		ColourOption(Green))
	fmt.Println(bs)
	// output: &{Lewis' Lovely Bikeshed RecycledCoffeeGrounds 5 2 Green}
}

func TestNewBikeShedHappy(t *testing.T) {
	testCases := []struct {
		name     string
		options  []OptionFunc
		expected BikeShed
	}{
		{
			name: "default",
			expected: BikeShed{
				name:     "default",
				length:   5,
				width:    2,
				material: Steel,
				colour:   Black,
			},
		},
		{
			name:    "specify colour",
			options: []OptionFunc{ColourOption(Blue)},
			expected: BikeShed{
				name:     "specify colour",
				length:   5,
				width:    2,
				material: Steel,
				colour:   Blue,
			},
		},
		{
			name:    "specify material",
			options: []OptionFunc{MaterialOption(Aluminium)},
			expected: BikeShed{
				name:     "specify material",
				length:   5,
				width:    2,
				material: Aluminium,
				colour:   Black,
			},
		},
		{
			name:    "specify length",
			options: []OptionFunc{LengthOption(10)},
			expected: BikeShed{
				name:     "specify length",
				length:   10,
				width:    2,
				material: Steel,
				colour:   Black,
			},
		},
		{
			name:    "specify width",
			options: []OptionFunc{WidthOption(3)},
			expected: BikeShed{
				name:     "specify width",
				length:   5,
				width:    3,
				material: Steel,
				colour:   Black,
			},
		},
		{
			name:    "specify multiple",
			options: []OptionFunc{WidthOption(3), LengthOption(8)},
			expected: BikeShed{
				name:     "specify multiple",
				length:   8,
				width:    3,
				material: Steel,
				colour:   Black,
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			got, err := NewBikeShed(testCase.name, testCase.options...)
			assert.NoError(t, err)
			assert.Equal(t, &testCase.expected, got)
		})
	}
}

func TestNewBikeShedUnhappy(t *testing.T) {
	testCases := []struct {
		name        string
		options     []OptionFunc
		expectedErr error
	}{
		{
			name:        "bricks not brown",
			options:     []OptionFunc{MaterialOption(Brick), ColourOption(White)},
			expectedErr: ErrBrickMustBeBrown,
		},
		{
			name:        "length too low",
			options:     []OptionFunc{LengthOption(3.9)},
			expectedErr: ErrLengthTooLow,
		},
		{
			name:        "length too high",
			options:     []OptionFunc{LengthOption(50.1)},
			expectedErr: ErrLengthTooHigh,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			got, err := NewBikeShed(testCase.name, testCase.options...)
			assert.Equal(t, testCase.expectedErr, err)
			assert.Nil(t, got)
		})
	}
}
