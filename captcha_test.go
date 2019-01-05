package captcha_test

import (
	"testing"

	"github.com/jo-apichai/captcha"
)

func TestCreate(t *testing.T) {
	testCases := []struct {
		input [4]int
		want  string
	}{
		{
			input: [4]int{1, 1, 1, 1},
			want:  "1 + ONE",
		},
		{
			input: [4]int{2, 2, 2, 2},
			want:  "TWO - 2",
		},
		{
			input: [4]int{1, 3, 3, 3},
			want:  "3 * THREE",
		},
	}

	for _, c := range testCases {
		i := c.input

		if r := captcha.Create(i[0], i[1], i[2], i[3]); c.want != r {
			t.Errorf("Expect: %s, Got: %s", c.want, r)
		}
	}
}
