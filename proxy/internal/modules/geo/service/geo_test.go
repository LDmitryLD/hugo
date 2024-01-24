package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGeo_SearchAddresses(t *testing.T) {
	geo := NewGeo()

	in := SearchAddressesIn{
		Query: "москва",
	}

	out := geo.SearchAddresses(in)

	assert.NotEmpty(t, out.Address)
	assert.Nil(t, out.Err)
}
