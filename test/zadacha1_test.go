package test

import (
	"strconv"
	"testing"
)

func TestAbs(t *testing.T) {
	tests := []struct {
		name  string
		value float64
		want  float64
	}{
		{
			name:  "на положительное дано",
			value: 10,
			want:  10,
		},
		{
			name:  "на отрицательное дано",
			value: -10,
			want:  10,
		},
		{
			name:  "на ноль",
			value: 0,
			want:  0,
		},
		{
			name:  "на минимальное значение",
			value: -0.000000001,
			want:  0.000000001,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if sum := Abs(test.value); sum != test.want {
				sum1 := strconv.Itoa(int(sum))
				want := strconv.Itoa(int(test.want))
				t.Errorf("Sum() = " + sum1 + ", want %d" + want)
			}
		})
	}
}

// Abs возвращает абсолютное значение.
// Например: 3.1 => 3.1, -3.14 => 3.14, -0 => 0.
