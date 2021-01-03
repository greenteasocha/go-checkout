package main

import (
	// "os"
	"fmt"
	"sort"
	// "runtime"
	// "reflect"
	// "time"
	// "sample-proj/add"
	"gopkg.in/src-d/go-git.v4"
	// . "gopkg.in/src-d/go-git.v4/_examples"
	"gopkg.in/src-d/go-git.v4/plumbing"
	// "github.com/go-git/go-git/v5/plumbing/object"

)

func main() {
	// defer func() {
    //     if r := recover(); r != nil {
    //         fmt.Fprintf(os.Stderr, "Panic: %v\n", r)
    //         for depth := 0; ; depth++ {
    //             pc, src, line, ok := runtime.Caller(depth)
    //             if !ok {
    //                 break
    //             }
    //             fmt.Fprintf(os.Stderr, " -> %d: %x: %s:%d\n", depth, pc, src, line)
    //         }
    //     }
	// }()
	
	dir := "/Users/k_abe/projects/panasonic/aws-kurashi-app-api"
	r, err := git.PlainOpen(dir)	
	if err != nil { panic(err) }
	fmt.Println("Directory Opened.")

	tagrefs, err := r.Branches()
	var tarr []*plumbing.Reference

	err = tagrefs.ForEach(func(t *plumbing.Reference) error {
		// h := t.Hash()
		tarr = append(tarr, t)
		return nil
	})

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

	w, err := r.Worktree()
	if err != nil {
		panic(err)
	}


	// ... checking out to commit
	// branch := tarr[0].Name()
	// commit := tarr[1].Hash().String()
	// fmt.Println(commit)
	name := tarr[1].Name()
	fmt.Println(name)
	// fmt.Println(branch)
	// // go := git.CheckoutOptions()
	// go := git.CheckoutOptions{
    //     Branch: branch
	// }

	gco := git.CheckoutOptions{
		// Hash: plumbing.NewHash(commit),
		Branch: name,
	}

	// fmt.Println(gco.Hash)
	// fmt.Println(gco.Branch)
	// fmt.Println(gco.Create)
	// fmt.Println(gco.Force)
	// fmt.Println(gco.Keep)

	err = w.Checkout(&gco)
	if err != nil {panic(err)}
    // handle error
}