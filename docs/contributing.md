
I will accept practically any feature additions to Jarvis assuming they don't completely break the core of how jarvis operates. For example:

* Creating a command which provides a regex that overlaps with an existing command (nope!)
* Removing the imposed rate-limits (nope!)
* Crashes jarvis (nope!)
* Exposes a significant security vulnerability (only if its totally badass!)

# Building and Running

Check out the `deployment.md` document for more info on that.

In essence, you'll need `go` (latest prefered) and `gb` (go build tool). You'll also need to set up a `config.yaml` file with API keys to the many external services Jarvis can access. Finally, you'll need to have `redis` running at the URL specified in `config.yaml` and `docker` running on the same machine as Jarvis.

If you are missing any of these steps, its likely (not guaranteed, but likely) Jarvis will just refuse to start.

# Code

All of the jarvis source code is located in `src/`. `vendor/` contains all the libraries we use. We check that into version control because that's how GB rolls.

## `src/main`

This is the main method of jarvis. It calls out to a bunch of init functions in the `jarivs/` package to get initial setup done before starting a read-loop on the websocket and blocking itself with `select {}` (which is like `for {}` except it doesn't use CPU time).

## `jarvis/commands`

Contains all the command definitions. `util.Command` is an interface which every command has to implement. After being implemented, the command is "loaded" into a hash of command names to command objects in `handlers/command.go`.

Through the `util.Command` interface, Commands have to provide a lot of information about themselves, including documentation, regex patterns which they match on, and the functionality of the command. Help documentation is generated automatically based on the information you export.

### SubCommands

Every command is really just a collection of subcommands. Each subcommand has a regex pattern which maps to a function you execute. The function takes in two arguments; the incoming message and the regex which it matched on.

So why this distinction? Imagine it like this: A Command is a collection of subcommands and documentation which applies to all subcommands. If two subcommands are so different that they can't share the same documentation then it should be a completely separate command.

### Regex Patterns

The regex patterns commands export should be as precise as possible and include a start/end string delimiters (`^jarvis mycommand$`). This minimizes the number of assumptions you have to make in the command and should keep things pretty clean looking forward. You're welcome to have fun inside of the command and export as many commands as you want to give the best user experience possible; just remember that balancing "natural language support" and precise command matching can be difficult.

Some pre-processing is done on each command message before being checked against your regexes. You can assume the following when writing regex:

* The word `jarvis` will always look like that, even if the user writes `Jarvis`, `jarvis, `, or even `jarivs`.
* Trailing spaces will be removed

The regexes you provide are actually `util.Regex` objects, which is just a light wrapper I created around the normal library to make a few things easier. For example, to get subexpressions you just say `r.SubExpression("jarvis hello", 0)`.

SubExpressions are key to writing good jarvis commands. When doing those, make sure to name them. This is *purely* for the sake of documentation. Example: don't do `^jarvis mycommand (.+)$`; instead, write `^jarvis mycommand (?P<argument>.+)$`. When the user issues `jarvis help mycommand` they will get a printout of your regexes which includes the tags, so its easier for them to see how the command works.

## `jarvis/config`

Configuration information about jarvis, primarily sourced from `config.yaml`. Not that important.

## `jarvis/data`

This is the package which contains anything related to storing persistent data. Right now persistent data is stored in redis.

### Datum

A "datum" in jarvis-land is a piece of data users provide which I have given meaning. To have "meaning" means that its not enough to just *store* the data for the user, but also to make that data useful later. For example, a user can store their zipcode, and then the weather command can access that to use it as a default location.

Each datum has a master key which is appended with the user's userid. It also has a list of english "trigger phrases" which are like aliases for the master key. For example: datum `user-zipcode-` is aliased to be recognized by both `my zipcode` and `my zip code`. Thus if I type `jarvis remember my zipcode` I get a little bit of leeway there.

## `jarvis/handler`

A handler is a piece of code which takes an incoming message and does something with it. There are three handlers; the command handler which engages on any message that starts with `jarvis` except `jarvis help`, the help handler which engages on any message that starts with `jarvis help`, and the logging handler which just logs every message we get to the prompt.

### Command Handler

The command handler does a few things

1. It checks to ensure the incoming message is a command
2. It modifies the command text a bit to help the commands read it better (see Regex Patterns above)
3. It rate-limits the user sending the command so that they can't DDOS jarvis
4. It attempts to match the command to a regex, then executes it if it matches

### Help Handler

`jarvis help` is not an actual command, believe it or not. Reason being? Circular dependencies. Help functionality needs access to a list of every command, but the only place that list exists is in `handlers/command.go`. Thus if help were a command in `jarvis/commands`, `commands` would need to import `handlers` and `handlers` needs to import `commands`. No bueno.

Instead, `help` functionality is its own handler.

### Other

Handlers provide a channel through which messages are sent. They then "subscribe" to messages which come from the `ws` package.

There's also the `announce.go` shit in there. I didn't know where else to put it.

## `jarvis/log`

Just logging functionality.

## `jarvis/service`

A service is when we attempt to access anything external to jarvis. That could be an API (slack http api) or something on the machine (docker). The one thing in there which doesn't follow this pattern is `time_service.go`, which contains a bunch of highly complicated logic to convert times from `time.Time` objects to strings and back. Originally this was in `jarvis/util` but it got so complicated that I felt it deserved its own Service.

Each service should provide a type which should contain no fields. Its just there to namespace your service's functions. A command might use your service like `service.Slack{}.UsernameFromUserId()`. Forcing callers to provide arguments on service creation is a no-no; that stuff should be provided in `config.yaml` and the service accesses it automatically.

If your method call returns an error, that error should be *immediately* sendable back to slack.

## `jarvis/util`

Utility functions and structs and interfaces and stuff.

## `jarvis/ws`

Wraps websocket reading and writing. This shouldn't require much modifying.

But this _is_ how you send messages back to slack. You import `jarvis/ws` then just say `ws.SendMessage(message, channel)`. Its that simple.

Right now if a message is over 4000 characters then it will refuse to send and instead send an error back to slack. However, eventually I will add functionality to send a snippet instead.

# Writing A New Command

Dude its so simple.

1. Copy the template from `docs/template_command.go`.
2. CMD-F and find/replace all instances of 'MyCommand' with the UpperCased name of your command
3. In `Name()`, make sure the value that is returned is one word and completely lower-case (preferably no underscores, hyphens, or digits either)
4. Put your regex in `SubCommands()`. You can put as many as you want. The second argument is the function handler for the regex on match.
5. Provide a function which handles when the regex matches and probably eventually calls `ws.SendMessage()` to send something back.
6. Want to provide really complicated functionality? Put it in `jarvis/service` or `jarvis/util` to keep the command code clean.
7. Want to store persistent data? Put it in `jarvis/data`.
8. Want to spin off a new thread? If its fairly lightweight work, just do a goroutine! If its heavy heavy stuff, create a lambda function and call it.
9. Register your command in `handlers/commands.go`. Its a big ol' list in there, you can't miss it.
