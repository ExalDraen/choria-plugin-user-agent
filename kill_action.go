package main

import (
	"fmt"
	"os/exec"

	"github.com/choria-io/go-external/agent"
)

type killRequest struct {
	User string `json:"user"`
}

// KillAction implements an agent.ActionHandler handler for killing all the sessions
// of a particular user
func KillAction(request *agent.Request, reply *agent.Reply, config map[string]string) {
	req := &killRequest{}

	// parse the received request, sets appropriate errors on reply on failure
	if !request.ParseRequestData(req, reply) {
		return
	}

	err := killUserSessions(req.User)
	if err != nil {
		reply.Statuscode = agent.Aborted
		reply.Statusmsg = fmt.Sprintf("could not kill user sessions %s#%s(%s): %s",
			request.Agent, request.Action, req.User, err)
		return
	}
}

func killUserSessions(user string) error {
	err := exec.Command("/usr/bin/pkill", "-TERM", "-u", user).Run()
	if err != nil {
		return fmt.Errorf("pkill error: %s", err)
	}
	return nil
}
