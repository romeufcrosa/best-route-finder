package services

import (
	"testing"

	. "github.com/onsi/gomega"
)

func Test(t *testing.T) {
	RegisterTestingT(t)

	testCases := []struct {
		desc string
	}{
		{
			desc: "",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {

		})
	}
}
