package main

import (
	"encoding/json"
	"fmt"
	"grpc/user"

	"google.golang.org/protobuf/proto"
)

func main() {
	article := &user.Article{
		Aid:   1,
		Title: "protobuf for golang",
		Views: 100,
	}
	// 序列化为二进制数据
	bytes, _ := proto.Marshal(article.ProtoReflect().Interface())
	fmt.Printf("bytes: %v\n", bytes)

	// 反序列化
	otherArticle := &user.Article{}
	proto.Unmarshal(bytes, otherArticle.ProtoReflect().Interface());
	fmt.Printf("otherArticle: %v\n", otherArticle)

	// otherArticle 转为 json
	bytes, _ = json.Marshal(otherArticle)
	fmt.Println(string(bytes))
}