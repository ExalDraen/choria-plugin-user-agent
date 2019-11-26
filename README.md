# choria-plugin-user-agent

[![Go Report Card](https://goreportcard.com/badge/github.com/ExalDraen/choria-plugin-user-agent)](https://goreportcard.com/report/github.com/ExalDraen/choria-plugin-user-agent)

Go based agent plugin for the [choria orchestration system](https://choria.io)

## Supported Actions

### `list`

Lists users logged in on a given system (uses `w` under the covers).

### `kill`

Kills all of the sessions belonging to a given user (uses `pkill` under the covers):

```sh
 mco rpc user kill "user=alexander.hermes" -I /^csqdev-/
```
