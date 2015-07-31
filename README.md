# A Better Slackbot

[ ![Codeship Status for mhoc/jarvis](https://codeship.com/projects/c137cad0-f434-0132-da6b-46341c668533/status?branch=master)](https://codeship.com/projects/85567)

Or, a worse hubot, depending on how you see it. Probably that one.

# Installing and Running

0. `go get github.com/mhoc/jarvis`
1. Set up the API keys in the envvars listed in `config/env.go`. It will yell at you if you're missing one.
2. Set up `config.yaml` by the documentation in `config/yaml.go`. It will yell at you if you're missing something big.
3. Have a redis instance running on the machine with the URI stored in the env `REDIS_URI`.
4. `make` will install all the dependencies and start jarvis. Nice!

# Scripts

Directory contains some python scripts for managing jarvis development.
