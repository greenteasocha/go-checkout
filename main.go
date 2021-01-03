package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"

	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
)

func main() {
	// open
	dir := "./"
	op := git.PlainOpenOptions{ DetectDotGit: true, }
	r, err := git.PlainOpenWithOptions(dir, &op)	
	if err != nil { panic(err) }

	// enumerate
	var tarr []*plumbing.Reference
	tagrefs, err := r.Branches()
	err = tagrefs.ForEach(func(t *plumbing.Reference) error {
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
		fmt.Printf("[%d] ", i)
		fmt.Println(v.Name())
	}

	// select
	fmt.Printf(">> Select branch number you move on: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	num, err := strconv.Atoi(scanner.Text())
	if err != nil { panic(err) }

	// checkout
	name := tarr[num].Name()
	gco := git.CheckoutOptions{ Branch: name, }
 
	w, err := r.Worktree()
	if err != nil { panic(err) }
	err = w.Checkout(&gco)
	if err != nil { panic(err) }

	// end
	fmt.Printf("You moved on: %s\n", name)
}