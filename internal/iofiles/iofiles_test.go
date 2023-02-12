package iofiles_test

import (
	"testing"

	"github.com/edualb-challenge/treebabel/internal/iofiles"
)

func TestGetFirstLine(t *testing.T) {
	type input struct {
		pathFile string
	}
	type testCase struct {
		name   string
		in     input
		out    []byte
		expErr bool
	}

	tests := []testCase{
		{
			name: "should return first line when valid file path",
			in: input{
				pathFile: "../../testdata/treebabel/challenge-input.json",
			},
			out: []byte(`{"timestamp": "2018-12-26 18:11:08.509654","translation_id": "5aa5b2f39f7254a75aa5","source_language": "en","target_language": "fr","client_name": "airliberty","event_name": "translation_delivered","nr_words": 30, "duration": 20}`),
		},
		{
			name: "should return error when empty file",
			in: input{
				pathFile: "../../testdata/treebabel/empty.json",
			},
			expErr: true,
		},
		{
			name: "should return error when invalid file path",
			in: input{
				pathFile: "invalid/path",
			},
			expErr: true,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			got, err := iofiles.GetFirstLine(tc.in.pathFile)
			if err != nil && !tc.expErr {
				t.Errorf("unexpected error, got %v wants nil", err)
				t.FailNow()
			}
			if err == nil && tc.expErr {
				t.Error("unexpected no error value, got nil wants error")
				t.FailNow()
			}

			if string(got) != string(tc.out) {
				t.Errorf("unexpected bytes value,\ngot: %s\nwants: %s", string(got), string(tc.out))
			}
		})
	}
}

func TestGetLastLine(t *testing.T) {
	type input struct {
		pathFile string
	}
	type testCase struct {
		name   string
		in     input
		out    []byte
		expErr bool
	}

	tests := []testCase{
		{
			name: "should return first line when valid file path",
			in: input{
				pathFile: "../../testdata/treebabel/challenge-input.json",
			},
			out: []byte(`{"timestamp": "2018-12-26 18:23:19.903159","translation_id": "5aa5b2f39f7254a75bb3","source_language": "en","target_language": "fr","client_name": "taxi-eats","event_name": "translation_delivered","nr_words": 100, "duration": 54}`),
		},
		{
			name: "should return error when empty file",
			in: input{
				pathFile: "../../testdata/treebabel/empty.json",
			},
			expErr: true,
		},
		{
			name: "should return error when invalid file path",
			in: input{
				pathFile: "invalid/path",
			},
			expErr: true,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			got, err := iofiles.GetLastLine(tc.in.pathFile)
			if err != nil && !tc.expErr {
				t.Errorf("unexpected error, got %v wants nil", err)
				t.FailNow()
			}
			if err == nil && tc.expErr {
				t.Error("unexpected no error value, got nil wants error")
				t.FailNow()
			}

			if string(got) != string(tc.out) {
				t.Errorf("unexpected bytes value,\ngot: %s\nwants: %s", string(got), string(tc.out))
			}
		})
	}
}
