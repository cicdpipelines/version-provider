package main

import (
	"github.com/NoUseFreak/cicd/pkg/server"
	"github.com/cicdpipelines/version-provider/internal/version"
)

func main() {
	server.Serve("version", version.Provider())
}
