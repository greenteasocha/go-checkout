module sample-proj

go 1.15

replace sample-proj/add => ./add

require (
	github.com/go-git/go-git/v5 v5.2.0 // indirect
	github.com/sergi/go-diff v1.1.0 // indirect
	golang.org/x/crypto v0.0.0-20201221181555-eec23a3978ad // indirect
	gopkg.in/sourcegraph/go-vcsurl.v1 v1.0.0-20131114132947-6b12603ea6fd // indirect
	gopkg.in/src-d/go-git.v3 v3.2.0 // indirect
	sample-proj/add v0.0.0-00010101000000-000000000000 // indirect
)
