package base62

import "testing"

func TestToBase62(t *testing.T) {
	type args struct {
		seed int64
	}
	testScenarios := []struct {
		name     string
		args     args
		expected string
	}{
		// TODO: Add test cases.
		{"Seed 1000 returns \"g8\"", args{seed: 1000}, "g8"},
	}
	for _, test := range testScenarios {
		t.Run(test.name, func(t *testing.T) {
			if got := ToBase62(test.args.seed); got != test.expected {
				t.Errorf("ToBase62() returns %#v but expected %#v", got, test.expected)
			}
		})
	}
}

func TestFromBase62(t *testing.T) {
	type args struct {
		base62Str string
	}
	testScenarios := []struct {
		name     string
		args     args
		expected int64
	}{
		// TODO: Add test cases.
		{"Base62String \"g8\" returns 1000", args{base62Str: "g8"}, 1000},
	}
	for _, test := range testScenarios {
		t.Run(test.name, func(t *testing.T) {
			if got := FromBase62(test.args.base62Str); got != test.expected {
				t.Errorf("FromBase62() returns %#v but expected %#v", got, test.expected)
			}
		})
	}
}
