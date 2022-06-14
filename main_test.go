package main

import (
	"testing"

	checkOddEven "github.com/bachtiarfr/my-go-library"
	checkOddEvenV2 "github.com/bachtiarfr/my-go-library/v2"
	assert "github.com/stretchr/testify/assert"
)
func TestCekBilanganGanjilGenapV1(t *testing.T) {
	tests :=
		[]struct {
			name     string
			expected string
			param  int
		}{
			{
				name	 : "Input 1 number #01",
				expected : "Ganjil",
				param	 : 1,
			},
			{
				name	 : "Input 1 number #02",
				expected : "Genap",
				param	 : 5,
			},
			{
				name	 : "Input 1 number #03",
				expected : "Ganjil",
				param	 : 3,
			},
		}

	assert := assert.New(t)

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := checkOddEven.CekGanjilGenap(test.param)
			assert.Equal(result, test.expected)
		})
	}
}

func TestCekBilanganGanjilGenapV2(t *testing.T) {
	tests :=
		[]struct {
			name     string
			expected []string
			param  []int
		}{
			{
				name	 : "#1",
				expected : []string {"ganjil", "genap", "ganjil", "genap", "ganjil"},
				param	 : []int{1, 2, 3, 4, 5},
			},
			{
				name	 : "#1",
				expected : []string {"genap", "ganjil", "genap", "genap", "ganjil"},
				param	 : []int{10, 27, 36, 4, 15},
			},
			{
				name	 : "#1",
				expected : []string {"genap", "genap", "ganjil", "genap", "ganjil"},
				param	 : []int{1, 2, 3, 4, 5},
			},
		}

	assert := assert.New(t)

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := checkOddEvenV2.CekGanjilGenap(test.param...)
			assert.Equal(result, test.expected)
		})
	}
}