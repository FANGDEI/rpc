package main

import (
	"fmt"
	"grpc/blog"

	"google.golang.org/protobuf/encoding/protojson"
)

func main() {
	article := &blog.Article {
		Aid: 1,
		Title: "protobuf for golang",
		Views: 100,
	}

	// message to json
	jsonString := protojson.Format(article.ProtoReflect().Interface())
	fmt.Printf("jsonString: %v\n", jsonString)

	// json to message
	m := article.ProtoReflect().Interface()
	protojson.Unmarshal([]byte(jsonString), m)
	fmt.Printf("m: %v\n", m)
}