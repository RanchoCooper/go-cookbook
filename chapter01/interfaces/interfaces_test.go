package interfaces

import (
	"bytes"
	"io"
	"testing"
)

func TestCopy(t *testing.T) {
	type args struct {
		in io.ReadSeeker
	}
	tests := []struct {
		name    string
		args    args
		wantOut string
		wantErr bool
	}{
		{
			name: "base-case",
			args: args{bytes.NewReader([]byte("example"))},
			wantOut: "exampleexample",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := &bytes.Buffer{}
			err := Copy(tt.args.in, out)
			if (err != nil) != tt.wantErr {
				t.Errorf("Copy() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotOut := out.String(); gotOut != tt.wantOut {
				t.Errorf("Copy() gotOut = %v, want %v", gotOut, tt.wantOut)
			}
		})
	}
}
