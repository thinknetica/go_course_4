package main

import (
	"fmt"
	"go-core-4/20-NoSQL/2-KVStore/pkg/kvdb"
)

func main() {
	db := kvdb.New()
	db.SET("1", "test record")
	fmt.Println(db.GET("1"), db.GET("2"))
}
