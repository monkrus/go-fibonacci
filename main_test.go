package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHandleFib(t *testing.T) {
	testCases := []struct {
		name string
		num  int
		want []byte
	}{

		{
			name: "zero",
			num:  0,
			want: []byte("0"),
		},

		{
			name: "one",
			num:  1,
			want: []byte("1"),
		},

		{
			name: "two",
			num:  2,
			want: []byte("1"),
		},

		{
			name: "three",
			num:  3,
			want: []byte("2"),
		},

		{
			name: "minus",
			num:  -3,
			want: []byte("-3"),
		},
	}

	handler := http.HandlerFunc(handleFib)
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			rec := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", fmt.Sprintf("/fib?num=%d", tc.num), nil)
			handler.ServeHTTP(rec, req)
			assert.Equal(t, tc.want, rec.Body.Bytes())
			assert.NotEqual(t, tc.name, rec.Body.Bytes())

		})
	}
}

func Test_handleFib(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			handleFib(tt.args.w, tt.args.r)
		})
	}
}
