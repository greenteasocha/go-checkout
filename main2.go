package main

import (
	"fmt"
	// "os"

	"gopkg.in/src-d/go-git.v4"
	. "gopkg.in/src-d/go-git.v4/_examples"
	// "gopkg.in/src-d/go-git.v4/plumbing"
)

// Basic example of how to checkout a specific commit.
func main() {

	// url := "https://github.com/git-fixtures/basic.git"
	// directory := "a"
	// commit := "master"
	// // Clone the given repository to the given directory
	// Info("git clone %s %s", url, directory)
	// r, err := git.PlainClone(directory, false, &git.CloneOptions{
	// 	URL: url,
	// })

	commit := "master"
	dir := "/Users/k_abe/projects/panasonic/aws-kurashi-app-api/.git"
	r, err := git.PlainOpen(dir)	

	CheckIfError(err)

	// ... retrieving the commit being pointed by HEAD
	Info("git show-ref --head HEAD")
	ref, err := r.Head()
	CheckIfError(err)
	fmt.Println(ref.Hash())

	w, err := r.Worktree()
	CheckIfError(err)

	// ... checking out to commit
	Info("git checkout %s", commit)
	fmt.Println(w)
	// err = w.Checkout(&git.CheckoutOptions{
	// 	Hash: plumbing.NewHash(commit),
	// })
	// CheckIfError(err)

	// // ... retrieving the commit being pointed by HEAD, it shows that the
	// // repository is pointing to the giving commit in detached mode
	// Info("git show-ref --head HEAD")
	// ref, err = r.Head()
	// CheckIfError(err)
	// fmt.Println(ref.Hash())
}