package main

import (
	"bytes"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_decode(t *testing.T) {
	t.Run("decode", func(t *testing.T) {
		t.Run("should return a Torrent struct", func(t *testing.T) {
			torrent, _ := decode(bytes.NewReader([]byte(`d8:announce32:http://example.com:6969/announce7:comment9:"Comment"10:created by8:John Doe13:creation datei1702236381e4:infod6:lengthi550809600e4:name9:image.iso12:piece lengthi262144e6:pieces7:abcdefgee`)))
			if diff := cmp.Diff(Torrent{
				Announce: "http://example.com:6969/announce",
				Info: Info{
					Pieces:      "abcdefg",
					PieceLength: 262144,
					Length:      550809600,
					Name:        "image.iso",
				},
			}, *torrent); diff != "" {
				t.Errorf("mismatch (-want +got):\n%s", diff)
			}
		})
	})
}
