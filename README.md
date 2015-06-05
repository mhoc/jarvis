# A Better Slackbot

Or, a worse hubot, depending on how you see it.

# Writing Commands

Every js file in the commands directory contains a command. Wait, that's obvious.
Each command is formatted like so:

```
module.exports = {
  description: "Imagine jarvis himself is saying this. That's how this should sound.",
  match: [ /regex/, /array/ ],
  run: function(msg, respond) {}
}
```

Or

```
module.exports = [
  { see above },
  { oh here's another }
]
```

Depending on how many commands you want to include in that one file. Up to you.

### Match

`match` is an array of regex matches. You can look at examples in the provided
commands, but the gist is that if a given message matches a regex you provide,
the result of that match is provided in the slack message object under the
`_matchResult` key. This means you can do grouping very easily, which is
pretty damn cool.

If two commands provide matches which could both match a given message, consider
the behavior undefined. It will definitely match and execute one of them but
not both and it won't complain while doing so.

### Talking Back

The `respond` parameter on run is a function through which you provide a slackMsg
object. This object should at minimum have two fields:

```
{ "text": "The body of the message to send back.", "channel": "the channel id" }
```

It is adequate enough to simply modify the text element of the `msg` passed in
then reuse the rest.
