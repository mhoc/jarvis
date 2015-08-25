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

All environment config for jarvis is done through config.yaml. As such, this file will have secrets in it. An example file is provided in the repo with all the available keys set.

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
