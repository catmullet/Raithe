![Raithe](https://github.com/catmullet/Raithe/blob/master/image%20(1).jpg)

# Raithe
---
## Lightweight | Persistent | Fast | Simple - Messaging Queue

### Register Clients
---
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
