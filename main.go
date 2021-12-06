package main

import (
	"flag"
	"fmt"
	"io/fs"
	"log"
	"net/http"

	"gomodules.xyz/memfs"
)

func main() {
	port := flag.String("p", "8100", "port to serve on")
	flag.Parse()

	rootFS := memfs.New()

	err := rootFS.MkdirAll("dir1/dir2", 0777)
	if err != nil {
		panic(err)
	}

	err = rootFS.WriteFile("dir1/dir2/f1.txt", []byte("incinerating-unsubstantial"), 0755)
	if err != nil {
		panic(err)
	}

	err = fs.WalkDir(rootFS, ".", func(path string, d fs.DirEntry, err error) error {
		fmt.Println(path)
		return nil
	})
	if err != nil {
		panic(err)
	}

	content, err := fs.ReadFile(rootFS, "dir1/dir2/f1.txt")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", content)

	http.Handle("/", http.FileServer(http.FS(rootFS)))
	log.Fatal(http.ListenAndServe(":"+*port, nil))
}
