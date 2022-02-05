package files

import (
	"encoding/json"
	"fmt"
	"testing/fstest"
)

type Post struct {
	Title string
	Desc  string
	Text  string
}

func FileParser(fsMap fstest.MapFS) []Post {

	posts := make([]Post, 0)

	for fileName, data := range fsMap {
		fmt.Printf("Processing file %s \n", fileName)
		var post Post
		json.Unmarshal([]byte(data.Data), &post)
		posts = append(posts, post)
	}
	return posts
}
