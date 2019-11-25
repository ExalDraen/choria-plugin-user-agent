package main

import (
	"bufio"
	"fmt"
	"os/exec"

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

	var out []string
	for _, s := range sess {
		out = append(out, s.String())
	}
	reply.Data = map[string][]string{
		"sessions": out,
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
	var name, tty, host, login, idle, jcpu, pcpu, what string
	res := []Session{}
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Sscan(line, &name, &tty, &host, &login, &idle, &jcpu, &pcpu, &what)
		res = append(res, Session{loginName: name, remoteHost: host, command: what})
	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("reading `w` output: %s", err)
	}
	if err := cmd.Wait(); err != nil {
		return nil, err
	}
	return res, nil
}
