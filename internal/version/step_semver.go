package version

import (
	"fmt"

	cicd "github.com/NoUseFreak/cicd/pkg/helper"
	git "gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing/object"
)

func stepSemver() *cicd.Step {
	return &cicd.Step{
		Schema: map[string]*cicd.Schema{
			"type": {
				Type:     cicd.TypeString,
				Required: true,
			},
			"directory": {
				Type:     cicd.TypeString,
				Required: false,
			},
			"version": {
				Type:     cicd.TypeString,
				Computed: true,
			},
		},
		Run: func(data *cicd.RunArguments, ctx *cicd.Context) (*cicd.StepResponse, error) {
			fmt.Println("Generating version")

			v, err := gitCommitOffset(data.GetStringOrDefault("directory", ctx.Pwd))
			if err != nil {
				return nil, err
			}

			resp := cicd.StepResponse{
				Status: true,
			}
			resp.SetString("version", v)

			return &resp, nil
		},
	}
}

func gitCommitOffset(dir string) (string, error) {
	r, err := git.PlainOpen(dir)
	if err != nil {
		return "", err
	}
	ref, err := r.Head()
	if err != nil {
		return "", err
	}
	cIter, err := r.Log(&git.LogOptions{From: ref.Hash()})
	if err != nil {
		return "", err
	}

	count := 0
	err = cIter.ForEach(func(c *object.Commit) error {
		count++
		return nil
	})

	return fmt.Sprintf("%s.%s.%d", "1", "0", count), nil
}
