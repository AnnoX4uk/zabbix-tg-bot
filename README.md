#
## Introduction
This is a very simple application to send messages via telegram bot. 

## Requriments

You need to register new telegrambot with [botfather](https://telegram.me/BotFather) and add one command to your bot: "getchat"
The next, you create wery simple ```config.json``` with appsettings and paste TG token to it
```json
{
    "Chatid": #somevalue in int64,
    "Token":"YOUR_TOKEN_HERE",
    "ListenPort": 8080
}

```

Also you can set this variables as env:
```bash
export CHATID="chatid"
export TOKEN="token"
export PORT="port"  #default port is 8080
```

## First run

When you place your ```config.json``` in project folder, just run build and ther run application
```bash
 go build main.go 
```

Then find your bot, connect him to the new chat and send command ```/start``` and ```/getchat```. Bot send chat-id for you, what you need paste to your ```config.json```. Restart the bot.

## Test it & enjoy

Application listen port from your config and wait for POST messages with json. Right now, bot accept very simple json
```json
 {"message": "Your message here" }
```
You can test it with curl

```bash
 curl -X POST -d '{"message":"your message here"}' http://localhost:8080
```

Or make alertscript for zabbix like this:

```bash
#!/bin/bash

url=${1}
theme=${2}
message=${3}

str=`jq --null-input --arg message "${3}" '{"message":$message}'`

curl -X POST -d "${str}" "${url}"
```