package models

import (
	"github.com/dchest/uniuri"
)

// IDGeneratedLength defines the length of the id string
const IDGeneratedLength = 32

// IDGeneratorFunc returns a randomly generated string useable as identifier
var IDGeneratorFunc = func() string {
	return uniuri.NewLen(IDGeneratedLength)
}
