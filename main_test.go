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
				name	 : "#01",
				expected : "Ganjil",
				param	 : 1,
			},
			{
				name	 : "#02",
				expected : "Genap",
				param	 : 2,
			},
			{
				name	 : "#03",
				expected : "Ganjil",
				param	 : 3,
			},
			{
				name	 : "#04",
				expected : "Genap",
				param	 : 4,
			},
			{
				name	 : "#05",
				expected : "Ganjil",
				param	 : 5,
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
				name	 : "#01",
				expected : []string {"ganjil", "genap", "ganjil", "genap", "ganjil"},
				param	 : []int{1, 2, 3, 4, 5},
			},
			{
				name	 : "#02",
				expected : []string {"genap", "genap", "ganjil", "genap", "ganjil"},
				param	 : []int{12, 22, 39, 44, 5},
			},
			{
				name	 : "#03",
				expected : []string {"genap", "genap", "genap", "genap", "ganjil"},
				param	 : []int{30, 22, 12, 44, 5},
			},
			{
				name	 : "#04",
				expected : []string {"ganjil", "genap", "genap", "genap", "ganjil"},
				param	 : []int{11, 22, 2, 44, 15},
			},
			{
				name	 : "#05",
				expected : []string {"ganjil", "genap", "ganjil", "genap", "ganjil"},
				param	 : []int{5, 4, 3, 2, 1},
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

func BenchmarkCekGanjilGenapV2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		checkOddEvenV2.CekGanjilGenap(1, 2, 3, 4, 5)
	}
}