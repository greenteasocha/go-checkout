package main

import (
	"os"
	"fmt"
	"reflect"    
	// "sample-proj/add"
	"github.com/go-git/go-git/v5"
)

func main() {
	r, err := git.PlainClone("/tmp/foo", false, &git.CloneOptions{
		URL:      "https://github.com/go-git/go-git",
		Progress: os.Stdout,
	})
	if err != nil {
		fmt.Println(reflect.TypeOf(err))
		if err == git.ErrRepositoryAlreadyExists {
			fmt.Println("repo was already cloned")
			dir := "/Users/k_abe/projects/panasonic/aws-kurashi-app-api/.git"
			r, err = git.PlainOpen(dir)
		}

		if err != nil { panic(err) }
	}

	branches, err := r.Branches()
	fmt.Println(branches.Next())
	br, err := branches.Next()
	fmt.Println(reflect.TypeOf(br))
	h := br.Name()
	fmt.Println(reflect.TypeOf(h))
	fmt.Println(h)

	// for i, v := range *branches {
	// 	fmt.Println(v)
	// }

	ref, err := r.Head()
	if err != nil {
		panic(err)
	}

	fmt.Println(ref)

// 	iter, _ := r.Commits()
//   	defer iter.Close()

//   	for {
// 		commit, err := iter.Next()
// 		if err != nil {
// 			if err == io.EOF {
// 				break
// 			}

// 			panic(err)
//   	}

//   	fmt.Println(commit)
//   }
    // a := add.Inp()
    // b := add.Inp()
    // fmt.Println(add.Add(a, b))
}