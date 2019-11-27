package main

import (
	"github.com/choria-io/go-external/agent"
)

func main() {
	userAgent := agent.NewAgent("user")
	defer userAgent.ProcessRequest()

	// action will be invoked on demand
	userAgent.MustRegisterAction("kill", KillAction)
	userAgent.MustRegisterAction("list", ListAction)
}
