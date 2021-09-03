package common_test

import (
	"fmt"
	"net/http"
	"testing"

	"allaboutapps.dev/aw/go-starter/internal/api"
	"allaboutapps.dev/aw/go-starter/internal/test"
	"github.com/stretchr/testify/assert"
)

func TestGeasdftSumRoute(t *testing.T) {
	cases := []struct {
		counter  string
		wantCode int
		want     string
	}{
		{"0", http.StatusOK, "0\n"},
		{"1", http.StatusOK, "1\n"},
		{"2", http.StatusOK, "3\n"},
		{"3", http.StatusOK, "6\n"},
		{"53", http.StatusOK, "1431\n"},
		{"4294967295", http.StatusOK, "9223372034707292160\n"},
		{"4294967296", http.StatusOK, "9223372039002259456\n"},
		{"9223372036854775807", http.StatusOK, "42535295865117307928310139910543638528\n"},

		{"string", http.StatusBadRequest, "Please provide an integer.\n"},
		{"2.5", http.StatusBadRequest, "Please provide an integer.\n"},

		{"-1", http.StatusBadRequest, "Please provide a positive integer.\n"},
		{"-100", http.StatusBadRequest, "Please provide a positive integer.\n"},
	}

	test.WithTestServer(t, func(s *api.Server) {
		for _, tc := range cases {
			path := fmt.Sprintf("/-/sum/%s?mgmt-secret=%s", tc.counter, s.Config.Management.Secret)

			res := test.PerformRequest(t, s, "GET", path, nil, nil)

			assert.Equal(t, tc.wantCode, res.Result().StatusCode)
			assert.Equal(t, tc.want, res.Body.String())
		}
	})
}
