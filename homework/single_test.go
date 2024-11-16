package homework

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Single(t *testing.T) {
	a := NewSingleArray[int](3)
	err := a.Add(1, 0)
	assert.NoError(t, err)
	err = a.Add(3, 3)
	assert.NotNil(t, err)
}
