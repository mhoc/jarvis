# A Better Slackbot

[ ![Codeship Status for mhoc/jarvis](https://codeship.com/projects/c137cad0-f434-0132-da6b-46341c668533/status?branch=master)](https://codeship.com/projects/85567)

Or, a worse hubot, depending on how you see it. Probably that one.

# Installing and Running

0. `go get github.com/mhoc/jarvis`
1. Set up `config.yaml` as per the documentation below. It will yell at you if you're missing something big.
2. Have a redis instance running on the machine with the URI stored in `config.yaml` under `redis`.
3. `make` will install all the dependencies. `make run` starts jarvis. Enjoy!

# Scripts

Directory contains some python scripts for managing jarvis development.

`python scripts/new_command.py Coolcommand` will create a template go file in `commands/` for you.

# Config.yaml

All environment config for jarvis is done through config.yaml. As such, this file will have secrets in it. An example file is provided below:

```
# The URI and passwords and such where redis is hosted at
redis: 'localhost:6379'

# API Tokens
tokens:
  
  # Slack API token, duh
  slack: token-goes-here
  
  # Darksky for weather api (https://developer.forecast.io)
  darksky: token-goes-here
  
  # Zipcode API (https://www.zipcodeapi.com/)
  zipcode: token-goes-here

# A list of admins by userid
# Admins have access to special commands (see commands/debug.go for some of them)
admins:
  - U01234567

# A human readable location name for jarvis. used in jarvis status.
location: My Cool Computer

# A channel blacklist.
# Any commands sent on these channels will be ignored
blacklist:
  - G0123456N

# A channel whitelist. 
# If this list has at least one element, any commands sent to any channel except those in the whitelist will be ignored
whitelist:
  - G083EQ05N

# Custom responses
# See commands/static.go
static:
  - key: michael
    value: I'm pretty sure michael is the coolest guy of all time, right?

```

# Example (Probably Bad) Deployment

* Follow the steps under installing and running.
* Using `supervisor` create the following in `/etc/supervisor/conf.d/jarvis.conf`:

```
[program:jarvis]
command=/root/go/src/github.com/mhoc/jarvis/jarvis
autostart=true
autorestart=true
startretries=10
user=root
```

* And in `etc/supervisor/conf.d/redis.conf`

```
[program:redis]
command=redis-server
autostart=true
autorestart=true
startretries=10
user=root
directory=/root/redis
```

* Issue `supervisorctl reload` to restart both jarvis and redis.

My entire continuous integration deployment script on codeship is:

```
ssh root@domain.com "cd go/src/github.com/mhoc/jarvis && git pull && supervisorctl reload"
```

# FAQ

I use the word "frequently" very lightly.

* What happens if two commands have overlapping regex definitions?

The answer is that it depends.

For example, if in `config.yaml` you include a `static` definition where `key` is `help`, then I can tell you that jarvis will output both the help text and also your static command.

That being said, providing cleaner command regex definitions is an area of improvement I am looking in to. Right now there are a few commands (see debug) where the command itself checks regexes against the input and might do nothing. There are others which specifically check that the word "jarvis" is in the input (weather.go, debug.go). There are others which don't do this. Its a mess that im going to improve.
