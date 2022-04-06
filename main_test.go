package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCarta_String(t *testing.T) {
	tests := map[string]struct {
		Valore         int
		Seme           string
		ExpectedString string
	}{
		"A di picche": {
			Valore:         1,
			Seme:           "picche",
			ExpectedString: "A",
		},
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
			c := &Carta{
				Valore: testCase.Valore,
				Seme:   testCase.Seme,
			}

			assert.Equal(st, testCase.ExpectedString, c.String())
		})
	}
}
