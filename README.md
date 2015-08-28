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

# Command Structure (lol)

My definition of a command is "a set of functionality which can all adequately share a common help topic". Every command has a single help topic, but can define multiple regexes and, if it makes sense, entirely different execution paths for each regex.

First, the regexes. The regexes a command defines for itself should be FULL message matches. An example would be `^jarvis status$`. You can always assume that the first letter of the entire regex is lowercase (to account for mobile keyboard autocorrects). Technically there is nothing stopping you from defining a regex that looks like `status`, but in my opinion the former method makes more sense. For example, check out the `remember` command: what happens if we want to define a user storage key with the word "status" in it?

And moreover, I agree that if you only define, say, `status`, it makes for more NLP-like command capability because it would support invokations like "jarvis what is your status" or "jarvis tell me your status". But there's nothing stopping you from crafting a regex, or even multiple regexes, which emulatest he same behavior. It will never be as good, but making the match regexes as conservative as possible helps with... I guess you could call it namespace pollution, really.

These regexes are provided to the command handler (`handlers/command.go`) through the `Matches()` interface method, which returns a `[]util.CommandMatch`. This type is pretty simple; it just ties a Regex statement to a function which takes in a `util.IncomingSlackMessage`.

Check out `weather.go` for a perfect example of how to use this and create new commands. The weather command has one purpose: to provide the weather to users. But there's a lot of sub-behavior we need to model, and we model it by crafting different invokation regexes which each have different behavior.
