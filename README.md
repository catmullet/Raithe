![Raithe](https://github.com/catmullet/Raithe/blob/master/image%20(1).jpg)

## Lightweight | Persistent | Fast | Simple - Messaging Queue

### Register Clients
---
A Client is any service attempting to push or pop from the message queue.  
Contained within the root directory is a file agents_list.json.  This file contains all the agents that can register and is read in realtime.  So Adding an agent is easy.

#### Request
```json
{
	"agent_name":"basic"
}
```
#### Response
```json
{
    "success": true,
    "message": "",
    "security_token": {
        "agent_name": "basic",
        "token": "{{token}}"
    }
}
```

### Push Message
---
Pusing a message can be done by any agent.  Queues are the key and will be how agents identify which queue to pull from.
#### Request
```json
{
	"queue":"enrollment",
	"message":"Hello World",
	"security_token": {
        "agent_name": "basic",
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
---
Popping a message from the queue will grab the oldest message first.  Simple as that.
#### Request
```json
{
	"queue":"enrollment",
	"security_token": {
        "agent_name": "basic",
        "token": "{{token}}"
    }
}
```
#### Response
```json
{
    "queue": "enrollment",
    "body": {
        "message": "Hello World",
        "queue": "enrollment",
        "security_token": {
            "agent_name": "basic",
            "token": "{{token}}"
        }
    }
}
