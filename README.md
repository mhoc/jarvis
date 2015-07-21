# A Better Slackbot

[ ![Codeship Status for mhoc/jarvis](https://codeship.com/projects/c137cad0-f434-0132-da6b-46341c668533/status?branch=master)](https://codeship.com/projects/85567)

Or, a worse hubot, depending on how you see it. Probably that one.

# Installing and Running

0. `go get github.com/mhoc/jarvis`

1. Set up the API keys in the envvars listed in `config/env.go`. It will yell at you if you're missing one.

2. Set up `config.yaml` by the documentation in `config/yaml.go`. It will yell at you if you're missing something big.

3. `make` will install dependencies, build, and run. Easy.

# Packages

I like using a lot of packages.

* `commands` is where command implementations go. Each one should provide a type which implements `util.Command` and should register itself in `handlers/command.go`.
* `config` contains logic for validating and providing configuration information from yaml and envvars.
* `data` contains a high-level interface to boltdb (not yet implemented).
* `handlers` are all goroutines which eat up messages from slack. They should each be started in `handlers/init.go`.
* `log` contains simple logging methods.
* `service` contains logic to get information from external sources, whether that be an API or the system itself like `git`.
* `util` has everything else that doesn't have a place and also has no inter-package dependencies besides `log` and `config`.
* `ws` is the websocket initialization, reading, and writing logic.

# Overview

Handlers provide channels to a method in `ws`. When a new slack message comes in, it is delivered to each channel. There should be a goroutine waiting to eat up the message on the other end. If not then it disappears.

One of those handlers is a simple print-to-console handler. The other is the command handler. It uses a bayesian classifier to match any message with the word `jarvis` in it to a command through the command's Learning Keywords it provides. Its not totally accurate.

# Commands

Commands can call `ws.SendMessage` to send a message back to slack at any time. They can spawn goroutines. They can store data in memory, or use boltdb to store something more persistent.

The learning keywords provided by each command should be short and unique. I think. Im not totally sure what's best actually but that seems to work well.

# Services

Services don't have an interface to abide by but there's a singleton pattern I use which I really like. Just go look at any of them and it should be obvious.
