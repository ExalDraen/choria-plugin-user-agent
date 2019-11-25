package main

import (
	"fmt"
	"os/exec"

	"github.com/choria-io/go-external/agent"
)

// ListAction implements an agent.ActionHandler handler for listing the users
// that are logged in on a given POSIX system
func ListAction(request *agent.Request, reply *agent.Reply, config map[string]string) {

	out, err := exec.Command("/usr/bin/w", "--no-header", "--from").Output()
	if err != nil {
		reply.Statuscode = agent.Aborted
		reply.Statusmsg = fmt.Sprintf("could not list users %s#%s: %s", request.Agent, request.Action, err)
		return
	}

	reply.Data = map[string]string{
		"message": string(out),
	}
}
