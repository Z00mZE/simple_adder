package sum

import (
	"context"
	"testing"
)

func TestService_Sum(t *testing.T) {
	type args struct {
		a int64
		b int64
	}
	tests := []struct {
		name    string
		args    args
		want    int64
		wantErr bool
	}{
		{
			name: "test 1",
			args: args{
				a: 0,
				b: 0,
			},
			want:    0,
			wantErr: false,
		},
		{
			name: "test 1",
			args: args{
				a: 1,
				b: 2,
			},
			want:    3,
			wantErr: false,
		},
	}

	ctx := context.Background()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := UseCase{}
			got, err := s.Sum(ctx, tt.args.a, tt.args.b)
			if (err != nil) != tt.wantErr {
				t.Errorf("Sum() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Sum() got = %v, want %v", got, tt.want)
			}
		})
	}
}
