package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCarta_String(t *testing.T) {
	tests := map[string]struct {
		Valore         int
		Seme           string
		ExpectedString string
	}{
		"Q di fiori": {
			Valore:         12,
			Seme:           "fiori",
			ExpectedString: "Q",
		},
		"K di cuori": {
			Valore:         13,
			Seme:           "cuori",
			ExpectedString: "K",
		},
	}

	for testName, testCase := range tests {
		t.Run(testName, func(st *testing.T) {
			c := Carta{
				Valore: testCase.Valore,
				Seme:   testCase.Seme,
			}

			assert.Equal(st, testCase.ExpectedString, c.String())
		})
	}
}
