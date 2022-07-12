package helpers_test

import (
	helpers "main/helpers"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_stddev(t *testing.T) {
	var tests = []struct {
		nameOfTest string
		testArray  []int
		want       float64
		wantError  bool
	}{
		{
			nameOfTest: "Test_01 VALID STDDEV",
			testArray:  []int{7, 4, -2},
			want:       3.7416573867739413,
			wantError:  false,
		},
	}
	for _, tt := range tests {

		t.Run(tt.nameOfTest, func(t *testing.T) {
			result := helpers.Stddev(tt.testArray)
			if !tt.wantError {
				assert.Equal(t, result, tt.want)
			}

		})
	}
}
