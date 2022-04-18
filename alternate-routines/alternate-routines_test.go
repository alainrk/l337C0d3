package l337C0d3

import (
	"testing"
)

func Test_alternateRoutines(t *testing.T) {
	type args struct {
		i int
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name:    "Error lt 0",
			args:    args{i: -1},
			wantErr: true,
		},
		{
			name:    "Error gt Z",
			args:    args{i: 200},
			wantErr: true,
		},
		{
			name:    "Ok up to 2",
			args:    args{i: 2},
			wantErr: false,
			want:    "1A2B",
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, err := alternateRoutines(tt.args.i)
			if (err != nil) != tt.wantErr {
				t.Errorf("alternateRoutines() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("alternateRoutines() = %v, want %v", got, tt.want)
			}
		})
	}
}
