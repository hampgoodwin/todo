package pagination

import (
	"testing"

	"github.com/matryer/is"
	"google.golang.org/genproto/googleapis/example/library/v1"
)

func TestCalculateRequestChecksum(t *testing.T) {
	t.Parallel()
	for _, tt := range []struct {
		name     string
		request1 Request
		request2 Request
		equal    bool
	}{
		{
			name: "same request",
			request1: &library.ListBooksRequest{
				Parent:    "shelves/1",
				PageSize:  100,
				PageToken: "token",
			},
			request2: &library.ListBooksRequest{
				Parent:    "shelves/1",
				PageSize:  100,
				PageToken: "token",
			},
			equal: true,
		},
		{
			name: "different parents",
			request1: &library.ListBooksRequest{
				Parent:    "shelves/1",
				PageSize:  100,
				PageToken: "token",
			},
			request2: &library.ListBooksRequest{
				Parent:    "shelves/2",
				PageSize:  100,
				PageToken: "token",
			},
			equal: false,
		},
		{
			name: "different page sizes",
			request1: &library.ListBooksRequest{
				Parent:    "shelves/1",
				PageSize:  100,
				PageToken: "token",
			},
			request2: &library.ListBooksRequest{
				Parent:    "shelves/1",
				PageSize:  200,
				PageToken: "token",
			},
			equal: true,
		},
		{
			name: "different page tokens",
			request1: &library.ListBooksRequest{
				Parent:    "shelves/1",
				PageSize:  100,
				PageToken: "token",
			},
			request2: &library.ListBooksRequest{
				Parent:    "shelves/1",
				PageSize:  100,
				PageToken: "token2",
			},
			equal: true,
		},
	} {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			a := is.New(t)
			t.Parallel()
			checksum1, err := calculateRequestChecksum(tt.request1)
			a.NoErr(err)
			checksum2, err := calculateRequestChecksum(tt.request2)
			a.NoErr(err)
			if tt.equal {
				a.Equal(checksum1, checksum2)
			} else {
				a.True(checksum1 != checksum2)
			}
		})
	}
}
