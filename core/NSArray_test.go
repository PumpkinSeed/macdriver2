package core_test

import (
	"testing"

	"github.com/PumpkinSeed/macdriver2/core"
	"github.com/stretchr/testify/assert"
)

func TestNSArraySize(t *testing.T) {
	arr := core.NSArray_WithObjects(core.String("a"), core.String("b"), core.String("c"))
	assert.EqualValues(t, 3, arr.Count())
	assert.Equal(t, []string{"a", "b", "c"}, arr.Strings())
}
