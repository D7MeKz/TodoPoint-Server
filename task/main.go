package task

import (
	"os"
	"path/filepath"
	"todopoint/common/config/mongodb"
)

func main() {
	// init MongoDB
	client := mongodb.NewMongoClient(filepath.Dir(os.Args[0]))
	mongodb.SetClient(client)

}
