# Really Bad Chat Server
#### Now includes 100% more stuff

This server is bad. Like, really bad. Don't use it. Ever. Or may you should use it? I don't 
know, nor do I care. 

Why did I make this? A friend bet me $50 I couldn't create a decently-functional chat server in an hour.

## Amazing Features

- You can set your name. Completely new technology, patent pending.
- Multiple chat rooms! Also patent pending.
- It totally supports emojis. So poop away! 💩💩💩 (as long as your terminal supports them).

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