package require_test

import (
	"encoding/json"
	"testing"

	"github.com/go-openapi/testify/v2/require"
)

func ExampleComparisonAssertionFunc() {
	t := &testing.T{} // provided by test

	t.Run("example", func(t *testing.T) {
		adder := func(x, y int) int {
			return x + y
		}

		type args struct {
			x int
			y int
		}

		tests := []struct {
			name      string
			args      args
			expect    int
			assertion require.ComparisonAssertionFunc
		}{
			{"2+2=4", args{2, 2}, 4, require.Equal},
			{"2+2!=5", args{2, 2}, 5, require.NotEqual},
			{"2+3==5", args{2, 3}, 5, require.Exactly},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				tt.assertion(t, tt.expect, adder(tt.args.x, tt.args.y))
			})
		}
	})
}

func ExampleValueAssertionFunc() {
	t := &testing.T{} // provided by test

	t.Run("example", func(t *testing.T) {
		dumbParse := func(input string) any {
			var x any
			_ = json.Unmarshal([]byte(input), &x)
			return x
		}

		tests := []struct {
			name      string
			arg       string
			assertion require.ValueAssertionFunc
		}{
			{"true is not nil", "true", require.NotNil},
			{"empty string is nil", "", require.Nil},
			{"zero is not nil", "0", require.NotNil},
			{"zero is zero", "0", require.Zero},
			{"false is zero", "false", require.Zero},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				tt.assertion(t, dumbParse(tt.arg))
			})
		}
	})
}

func ExampleBoolAssertionFunc() {
	t := &testing.T{} // provided by test

	t.Run("example", func(t *testing.T) {
		isOkay := func(x int) bool {
			return x >= 42
		}

		tests := []struct {
			name      string
			arg       int
			assertion require.BoolAssertionFunc
		}{
			{"-1 is bad", -1, require.False},
			{"42 is good", 42, require.True},
			{"41 is bad", 41, require.False},
			{"45 is cool", 45, require.True},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				tt.assertion(t, isOkay(tt.arg))
			})
		}
	})
}

func ExampleErrorAssertionFunc() {
	t := &testing.T{} // provided by test

	dumbParseNum := func(input string, v any) error {
		return json.Unmarshal([]byte(input), v)
	}

	t.Run("example", func(t *testing.T) {
		tests := []struct {
			name      string
			arg       string
			assertion require.ErrorAssertionFunc
		}{
			{"1.2 is number", "1.2", require.NoError},
			{"1.2.3 not number", "1.2.3", require.Error},
			{"true is not number", "true", require.Error},
			{"3 is number", "3", require.NoError},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				var x float64
				tt.assertion(t, dumbParseNum(tt.arg, &x))
			})
		}
	})
}
