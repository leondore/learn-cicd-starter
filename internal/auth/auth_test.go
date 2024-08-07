package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	type testCase struct {
		input  http.Header
		out    string
		errMsg string
	}

	tests := []testCase{
		{
			input: http.Header{
				"Authorization": nil,
			},
			out:    "",
			errMsg: "no authorization header included",
		},
		{
			input: http.Header{
				"Authorization": []string{"Bearer asdfasdfasdfasdfasdfasdf"},
			},
			out:    "",
			errMsg: "malformed authorization header",
		},
		{
			input: http.Header{
				"Authorization": []string{"asdfasdfasdfasdfasdfasdf"},
			},
			out:    "",
			errMsg: "malformed authorization header",
		},
		{
			input: http.Header{
				"Authorization": []string{"ApiKey asdfasdfasdfasdfasdfasdf"},
			},
			out:    "asdfasdfasdfasdfasdfasdf",
			errMsg: "",
		},
	}

	for _, tc := range tests {
		got, err := GetAPIKey(tc.input)
		if got != tc.out {
			t.Fatalf("Expected %s. Got %s", tc.out, got)
		}

		if err != nil {
			errMsg := err.Error()

			if errMsg != tc.errMsg {
				t.Fatalf("Unexpected error: %s. Expected: %s", errMsg, tc.errMsg)
			}
		}
	}
}
