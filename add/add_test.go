package add

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	got := Add(2, 2)
	want := 6

	assert.Equal(t, want, got)	
}
