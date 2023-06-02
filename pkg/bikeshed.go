package buildings

import (
	"errors"
)

type BikeShed struct {
	name     string
	material Material
	length   float32
	width    float32
	colour   Colour
}

const (
	minBikeShedLength = 4
	maxBikeShedLength = 50
	minBikeShedWidth  = 2
	maxBikeShedWidth  = 4
)

type Material int

const (
	Wood Material = iota
	Steel
	Aluminium
	Brick
	RecycledCoffeeGrounds
)

type Colour int

const (
	Black Colour = iota
	White
	Blue
	Green
	Brown
)

func NewBikeShed(name string, options ...OptionFunc) (*BikeShed, error) {
	// Set defaults and fixed param
	ret := BikeShed{
		name:     name,
		length:   5,
		width:    2,
		material: Steel,
		colour:   Black,
	}

	// Apply the options and validate individual values
	for _, option := range options {
		if err := option(&ret); err != nil {
			return nil, err
		}
	}

	// Composite validation
	if ret.material == Brick && ret.colour != Brown {
		return nil, ErrBrickMustBeBrown
	}

	return &ret, nil
}

var (
	ErrLengthTooLow     = errors.New("length too low")
	ErrLengthTooHigh    = errors.New("length too high")
	ErrWidthTooLow      = errors.New("width too low")
	ErrWidthTooHigh     = errors.New("width too high")
	ErrBrickMustBeBrown = errors.New("brick must be brown")
)

type OptionFunc func(*BikeShed) error

func MaterialOption(material Material) OptionFunc {
	return func(bs *BikeShed) error {
		bs.material = material
		return nil
	}
}

func ColourOption(colour Colour) OptionFunc {
	return func(bs *BikeShed) error {
		bs.colour = colour
		return nil
	}
}

func LengthOption(length float32) OptionFunc {
	return func(bs *BikeShed) error {
		if length < minBikeShedLength {
			return ErrLengthTooLow
		}
		if length > maxBikeShedLength {
			return ErrLengthTooHigh
		}
		bs.length = length
		return nil
	}
}

func WidthOption(width float32) OptionFunc {
	return func(bs *BikeShed) error {
		if width < minBikeShedWidth {
			return ErrWidthTooLow
		}
		if width > maxBikeShedWidth {
			return ErrWidthTooHigh
		}
		bs.width = width
		return nil
	}
}
