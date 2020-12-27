package main

import (
	// "os"
	"fmt"
	"sort"
	// "reflect"
	// "time"
	// "sample-proj/add"
	"github.com/go-git/go-git/v5"
	// "github.com/go-git/go-git/v5/_examples"
	"github.com/go-git/go-git/v5/plumbing"
	// "github.com/go-git/go-git/v5/plumbing/object"

)

func main() {
	// r, err := git.PlainClone("/tmp/foo", false, &git.CloneOptions{
	// 	URL:      "https://github.com/go-git/go-git",
	// 	Progress: os.Stdout,
	// })
	// if err != nil {
	// 	fmt.Println(reflect.TypeOf(err))
	// 	if err == git.ErrRepositoryAlreadyExists {
	// 		fmt.Println("repo was already cloned")
	// 		dir := "/Users/k_abe/projects/panasonic/aws-kurashi-app-api/.git"
	// 		r, err = git.PlainOpen(dir)
	// 	}

	// 	if err != nil { panic(err) }
	// }

	dir := ""
	r, err := git.PlainOpen(dir)	
	if err != nil { panic(err) }
	fmt.Println("Directory Opened.")

	tagrefs, err := r.Branches()
	var tarr []*plumbing.Reference

	err = tagrefs.ForEach(func(t *plumbing.Reference) error {
		h := t.Hash()
		// commit, _ := r.CommitObject(h)
		tarr = append(tarr, t)
		return nil
	})


	// for i, v := range tarr {
	// 	c, _ := r.CommitObject(v.Hash())
	// 	fmt.Println(i, c.Author.When)
	// }

	// sort
	sort.Slice(tarr, func(i, j int) bool {
		c1, _ := r.CommitObject(tarr[i].Hash())
		t1 := c1.Author.When
		c2, _ := r.CommitObject(tarr[j].Hash())
		t2 := c2.Author.When
        return t1.Before(t2)
	})
	
	// show
	for i, v := range tarr {
		// c, _ := r.CommitObject(v.Hash())
		fmt.Printf("[%d] ", i)
		// fmt.Println(c.Author.When)
		fmt.Println(v.Name())
	}


	// fmt.Println(tarr[0].Before(tarr[1]))
	// fmt.Println(tarr[0].After(tarr[1]))









	// branches, err := r.Branches()
	// fmt.Println("\nA branch: ")
	
	// ref, err := branches.Next()
	// fmt.Println("\nBranch reference: ")
	// fmt.Println(reflect.TypeOf(ref))
	// fmt.Println(ref)

	// name := ref.String()
	// fmt.Println("\nName: ")
	// fmt.Println(reflect.TypeOf(name))
	// fmt.Println(name)

	// br, err := r.Branch(name)
	// fmt.Println("\nBranch: ")
	// fmt.Println(reflect.TypeOf(br))
	// fmt.Println(br)




	// for i, v := range *branches {
	// 	fmt.Println(v)
	// }

	// ref, err := r.Head()
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println(ref)

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