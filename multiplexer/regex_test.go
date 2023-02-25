package multiplexer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRegexInt(t *testing.T) {
  findstr := regInt.FindString("/product/0.897")
  assert.Equal(t, "0", findstr)
}

func TestRegexFloat(t *testing.T) {
  matchFloat := regFloat.MatchString("0.8989701")
  assert.Equal(t, true, matchFloat)
  matchFloat = regFloat.MatchString("234")
  assert.Equal(t, false, matchFloat)
}

func TestRegexSlug(t *testing.T) {
  matchSlug := regSlug.MatchString("i-made-mei-sastra-jayadi")
  assert.Equal(t, true, matchSlug)
  matchSlug = regSlug.MatchString("I-Made-Mei-Sastra-Jayadi")
  assert.Equal(t, false, matchSlug)
  matchSlug = regSlug.MatchString("jayadi")
  assert.Equal(t, true, matchSlug)
}
