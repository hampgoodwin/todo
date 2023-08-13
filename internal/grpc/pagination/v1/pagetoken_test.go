package pagination

import (
	"strings"
	"testing"

	"github.com/matryer/is"
	"google.golang.org/genproto/googleapis/example/library/v1"
)

func TestParsePageToken(t *testing.T) {
	t.Parallel()
	a := is.New(t)

	t.Run("invalid format", func(t *testing.T) {
		t.Parallel()
		request := &library.ListBooksRequest{
			Parent:    "shelves/1",
			PageSize:  10,
			PageToken: "invalid",
		}
		pageToken1, err := ParsePageToken(request)
		a.True(err != nil && strings.Contains(err.Error(), "decode"))
		a.Equal(PageToken{}, pageToken1)
	})

	t.Run("invalid checksum", func(t *testing.T) {
		t.Parallel()
		request := &library.ListBooksRequest{
			Parent:   "shelves/1",
			PageSize: 10,
			PageToken: EncodePageTokenStruct(&PageToken{
				RequestChecksum: 1234, // invalid
			}),
		}
		pageToken1, err := ParsePageToken(request)
		a.True(err != nil && strings.Contains(err.Error(), "checksum"))
		a.Equal(PageToken{}, pageToken1)
	})
}
