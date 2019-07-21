package version

import (
	cicd "github.com/NoUseFreak/cicd/pkg/helper"
)

func Provider() *cicd.Provider {
	return &cicd.Provider{
		StepMap: map[string]*cicd.Step{
			"version_semver": stepSemver(),
		},
	}
}
