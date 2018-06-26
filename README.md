![Raithe](https://raw.githubusercontent.com/catmullet/Raithe/master/docs/img/banner.jpg)
[![Maintainability](https://api.codeclimate.com/v1/badges/94e11fd3b812339047c5/maintainability)](https://codeclimate.com/github/catmullet/Raithe/maintainability)
[![Go Report Card](https://goreportcard.com/badge/github.com/catmullet/Raithe)](https://goreportcard.com/report/github.com/catmullet/Raithe)
[![GoDoc](https://godoc.org/github.com/catmullet/Raithe?status.svg)](https://godoc.org/github.com/catmullet/Raithe)

# _The Compact and Persistent Messaging Queue_
## Getting Started
#### Run Go Get
```json
go get github.com/catmullet/Raithe
```
#### Modifying your env file
Your env file simply contains the port you want to run from and the redis configuration.
```json
# Essential Configurations
PORT=8021

# Redis
REDIS_URL=127.0.0.1:6379
REDIS_PASSWORD=""
REDIS_DB=0
```
#### Fire It Up
```json
go run raithe.go
```
![Raithe](https://raw.githubusercontent.com/catmullet/Raithe/master/docs/img/Raithe_cmd.png)
#### Your Agents List
Your agents list will contain all agents that can register as a "producer" or "consumer", both are the same to Raithe.
Contained within your agents_list.json of the root folder you will see something like this.  Just list the names of the agents and each agent will need to call the register path _(/auth/register)_ to stake thier claim on that agent name.
```json
{
  "agents": [
    "test"
  ]
}
```
## Playing Around
### Register Clients
A Client is any service attempting to push or pop from the message queue.  
Contained within the root directory is a file agents_list.json.  This file contains all the agents that can register and is read in realtime.  So Adding an agent is easy.

#### Request _/auth/register_
```json
{
	"agent_name":"{{agent name}}"
}
```
#### Response
```json
{
    "success": true,
    "message": "",
    "security_token": {
        "agent_name": "{{agent name}}",
        "token": "{{token}}"
    }
}
```

### Push Message
Pusing a message can be done by any agent.  Queues are the key and will be how agents identify which queue to pull from.
#### Request _/queue/push_
```json
{
	"queue":"test",
	"message":"Hello World",
	"security_token": {
        	"agent_name": "{{agent name}}",
        	"token": "{{token}}"
    }
}
```
#### Response
```json
{
    "success": true
}
```
### Pop Message
Popping a message from the queue will grab the oldest message first.  Simple as that.
#### Request _/queue/pop_
```json
{
	"queue":"test",
	"security_token": {
        	"agent_name": "{{agent name}}",
        	"token": "{{token}}"
    }
}
```
#### Response
```json
{
    "queue": "test",
    "message": "Hello World"
}
