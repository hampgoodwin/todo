package pagination

import "testing"

func TestHasNextPage(t *testing.T) {
	type args struct {
		limit      int64
		totalItems int64
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "should-calculate-ok-when-there-are-more-items",
			args: args{
				limit:      20,
				totalItems: 30,
			},
			want: true,
		},
		{
			name: "should-calculate-ok-when-there-are-no-more-items",
			args: args{
				limit:      20,
				totalItems: 20,
			},
			want: true,
		},
		{
			name: "there-should-not-be-more-pages",
			args: args{
				limit:      20,
				totalItems: 15,
			},
			want: false,
		},
		{
			name: "if-limit-0-should-not-be-more-pages",
			args: args{
				limit:      0,
				totalItems: 555,
			},
			want: false,
		},
		{
			name: "if-total-items-is-0-no-more-pages",
			args: args{
				limit:      69,
				totalItems: 0,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HasNextPage(tt.args.limit, tt.args.totalItems); got != tt.want {
				t.Errorf("HasNextPage() = %v, want %v", got, tt.want)
			}
		})
	}
}
