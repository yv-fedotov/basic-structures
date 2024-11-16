package homework

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_VectorArray(t *testing.T) {
	v := NewVectorArray[int](4)
	v.Add(10)
	v.Add(20)
	v.Add(30)
	assert.Equal(t, 3, len(v.array))

	_, err := v.Remove(3)
	assert.NotNil(t, err)
	element, err := v.Remove(2)
	assert.Equal(t, 30, element)

	assert.Equal(t, 2, len(v.array))

}
