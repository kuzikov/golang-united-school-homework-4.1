package reverse_string

import (
	"log"
	"testing"
)

func TestReverseString(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{input: "Hello!",
			expected: "!olleH"},
		{input: "abc",
			expected: "cba"},
		{input: "l0l",
			expected: "l0l",
		},
	}

	for _, test := range tests {
		if result := ReverseString(test.input); result != test.expected {
			log.Println(result)
			t.Errorf("Have: <%v>, want: <%v>", result, test.expected)
		}
	}
}
