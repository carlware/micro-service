// +build unit

package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIDGenerator(t *testing.T) {
	id := IDGeneratorFunc()
	
	assert.NotNil(t, id)
	assert.Equal(t, len(id), 32)
}
