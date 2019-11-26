package main

import (
	"bufio"
	"fmt"
	"os/exec"
	"strings"

	"github.com/choria-io/go-external/agent"
)

// Session describes a single logged in session
type Session struct {
	loginName  string
	remoteHost string
	command    string
}

// String returns a string representation of a session
func (s Session) String() string {
	return fmt.Sprintf("%s<%s>: %s", s.loginName, s.remoteHost, s.command)
}

// ListAction implements an agent.ActionHandler handler for listing the users
// that are logged in on a given POSIX system
func ListAction(request *agent.Request, reply *agent.Reply, config map[string]string) {

	sess, err := ListUserSessions()
	if err != nil {
		reply.Statuscode = agent.Aborted
		reply.Statusmsg = fmt.Sprintf("could not list users %s#%s: %s", request.Agent, request.Action, err)
		return
	}

	// If there are no user sessions, don't send data.
	if sess != nil {
		var out []string
		for _, s := range sess {
			out = append(out, s.String())
		}
		reply.Data = map[string][]string{
			"sessions": out,
		}
	}
}

// ListUserSessions lists the users logged in sessions
func ListUserSessions() ([]Session, error) {
	cmd := exec.Command("/usr/bin/w", "--no-header")
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return nil, err
	}
	if err := cmd.Start(); err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(stdout)
	res := []Session{}
	// Scan over each line of w's output and pull out the bits we want
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)
		// The command is the 8th field and may contain spaces
		res = append(res, Session{loginName: fields[0], remoteHost: fields[2], command: strings.Join(fields[7:], " ")})
	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("reading `w` output: %s", err)
	}
	if err := cmd.Wait(); err != nil {
		return nil, err
	}
	return res, nil
}
