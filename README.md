# A Better Slackbot

Or, a worse hubot, depending on how you see it.

# Writing Commands

Every js file in the commands directory contains a command. Wait, that's obvious.
Each command is formatted like so:

```
{
  match: [],
  run: function(msg) {}
}
```

So your command file could look like

```
module.exports = {
  match: [ ... ],
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

When crafting a match string you have access to three special variables.

* `{name}` is the name of the bot as defined in config.json

* `{postfix}` is essentially just "whatever is left at the end of the message"

* `{word}` can be repeated multiple times and looks for space-deliniated words

A few examples:

"{name} help me" would match against "jarvis help me" if the 
