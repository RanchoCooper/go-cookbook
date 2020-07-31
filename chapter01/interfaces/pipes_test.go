package interfaces

import "testing"

func TestPipeExample(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{
			name: "base-case",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := PipeExample(); (err != nil) != tt.wantErr {
				t.Errorf("PipeExample() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
