# A Better Slackbot

[ ![Codeship Status for mhoc/jarvis](https://codeship.com/projects/c137cad0-f434-0132-da6b-46341c668533/status?branch=master)](https://codeship.com/projects/85567)

Or, a worse hubot, depending on how you see it. Probably that one.

# Installing and Running

0. `go get github.com/mhoc/jarvis`

1. Set up the API keys in the envvars listed in `config/env.go`. It will yell at you if you're missing one.

2. Set up `config.yaml` by the documentation in `config/yaml.go`. It will yell at you if you're missing something big.

3. `make` will install dependencies, build, and run. Easy.

# Subscribing To Messages

The `handlers` package contains the various handlers which "subscribe" to messages from Slack. This is done through a call to `ws.SubscribeToMessages(chan util.IncomingSlackMessage)`. There is also `ws.SubscribeToAll(chan map[string]interface{})` to subscribe to every event, not just text messages.

Every new text message from slack is now sent to your channel and you can do whatever you want with it from the handler. Note that its possible (and likely) you might miss a message when Jarvis is first started.

Every handler should register itself in `handlers/init.go` (just look at how the others do it).

# Commands

Every command is a type which implements the `util.Command` interface. You can look at the other commands to see how it is done.

You can assume that your commmand's Execute() method is called on its own goroutine, so do whateverthehell you want. Even spawn other goroutines. Just call `ws.SendMessage()` when you're ready to send something back to slack. Multiple times or once. Doesn't matter.

Commands have to provide a list of "training words" on which new messages containing the word "jarvis" anywhere in them are compared against using a bayesian classification system. Anything with a match probability above some probability (currently 90%) is considered a match. Take a look at other commands to see examples, but you can make the strings as complete as you like.

Every command should register itself in `handlers/command.go`.
