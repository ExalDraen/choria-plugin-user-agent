package main

import (
	"github.com/choria-io/go-external/agent"
)

type echoRequest struct {
	Message string `json:"message"`
}

func echoAction(request *agent.Request, reply *agent.Reply, config map[string]string) {
	req := &echoRequest{}

	// parse the received request, sets appropriate errors on reply on failure
	if !request.ParseRequestData(req, reply) {
		return
	}

	reply.Data = map[string]string{
		"message": req.Message,
	}
}

func main() {
	userAgent := agent.NewAgent("user")
	defer userAgent.ProcessRequest()

	// action will be invoked on demand
	userAgent.MustRegisterAction("echo", echoAction)
	userAgent.MustRegisterAction("list", ListAction)
}
