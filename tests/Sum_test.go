package tests

import (
	"github.com/navaneet-rao/learn-go/pkg"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsEqual(t *testing.T) {
	assert := assert.New(t)
	a := 100
	b := 100
	assert.Equal(200, basicSum.Sum(a, b), "sum done correctly!!")
}
