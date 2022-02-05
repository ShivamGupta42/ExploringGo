package files

import (
	"testing"
	"testing/fstest"
)

func TestFileParser(t *testing.T) {
	fs := fstest.MapFS{
		"hello world.md":  {Data: []byte("{ \"Title\" : \"Sample1\" }")},
		"hello-world2.md": {Data: []byte("{ \"Title\" : \"Sample2\" }")},
	}

	var posts []Post
	posts = FileParser(fs)

	fileTests := []struct {
		post  Post
		title string
	}{
		{posts[0], "Sample1"},
		{posts[1], "Sample2"},
	}

	for _, tt := range fileTests {
		t.Run("testing parsing of Title field", func(t *testing.T) {
			got := tt.post.Title
			want := tt.title

			if got != want {
				t.Errorf("Got %s , Wanted %s", got, want)
			}
		})
	}
}
