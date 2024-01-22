package test

import (
	"github.com/go-playground/assert/v2"
	"testing"
)

func TestAbsTestifi(t *testing.T) {
	tests := []struct {
		name  string
		value float64
		want  float64
	}{
		{
			name:  "negative value",
			value: -3.001,
			want:  3.001,
		},
		{
			name:  "small value",
			value: -0.00000001,
			want:  0.00000001,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// меняем на функцию Equal из пакета assert
			assert.Equal(t, test.want, Abs(test.value))
		})
	}
}
