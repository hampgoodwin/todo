package pagination

import (
	"testing"

	"github.com/matryer/is"
)

func Test_PageTokenStruct(t *testing.T) {
	t.Parallel()
	type pageToken struct {
		Int    int
		String string
	}
	for _, tt := range []struct {
		name string
		in   pageToken
	}{
		{
			name: "all set",
			in: pageToken{
				Int:    42,
				String: "foo",
			},
		},
		{
			name: "default value",
			in: pageToken{
				String: "foo",
			},
		},
	} {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			a := is.New(t)

			t.Parallel()
			str := EncodePageTokenStruct(tt.in)
			var out pageToken
			a.NoErr(DecodePageTokenStruct(str, &out))
			a.Equal(tt.in, out)
		})
	}
}
