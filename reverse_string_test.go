package reverse_string

import (
	"testing"
)

func Test_average(t *testing.T) {
	data := []struct {
		In       string
		Expected string
	}{
		{In: "Hello world!", Expected: "!dlrow olleH"},
	}
	for _, q := range data {
		got := ReverseString(q.In)
		if got != q.Expected {
			t.Logf("ReverseString expected: %s, got: %s", q.Expected, got)
			t.Fail()
		}
	}
}
