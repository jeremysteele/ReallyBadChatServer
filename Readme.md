# Really Bad Chat Server
#### The worst chat server ever created by the hand of man

This server is bad. Like, really bad. Don't use it. Ever. Or may you should use it? I don't 
know, nor do I care. 

## The Horrible Protocol

Uses port 1234, because simplicity is secure. 

It uses some sort of TCP thing that listens on that super-secure port for random commands.

## Running it

`$ go run cmd\server\main.go`

## Command Reference

Use telnet to connect. 

- `SETNAME name` - Use a given name (32 char max). Default is anon
- `JOIN name` - Joins a room (32 char max).
- `QUIT` - Terminates the chat session. 
- `INFO` - Returns some sort of info about things.
- `MSG text` - Sends a message. Limit of 128 chars. 
- `LIST` - Lists all members in the room.