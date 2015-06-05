# A Better Slackbot

Or, a worse hubot, depending on how you see it.

# Installing

1. Set up the API keys as outlined below

2. `npm install`

3. `npm start`

# APIs

This bot does and will use quite a few APIs. The only one that's required is the slack auth token, obviously. If you are missing an API key for any of the others, that command will just not function. API keys should be set as an env var with the name given. All APIs should and do have free tiers that are adequate for even heavy usage of Jarvis.

* `SLACK_AUTH_TOKEN`: The authorization token for slack. Create a new bot integration and it should give you a token. (REQUIRED).

* `DARK_SKY_API_TOKEN`: Used for weather information. http://developer.forecast.io.

* `ZIP_CODE_API_TOKEN`: Used to convert zip codes to locations for things like weather. http://zipcodeapi.com.

# Writing Commands

Every js file in the commands directory contains a command. Wait, that's obvious. Each command is formatted like so:

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

So the base element can be an object or an array of the objects, depending on how you want to organize the js files you create.

### Match

`match` is an array of regex matches. You can look at examples in the provided commands, but the gist is that if a given message matches a regex you provide, the result of that match is provided in the slack message object under the `_matchResult` key. This means you can do grouping very easily, which is pretty useful.

If two commands provide matches which could both match a given message, consider the behavior undefined. It will definitely match and execute one of them but not both and it won't complain while doing so.

### Talking Back

The `respond` parameter on run is a function through which you provide a slackMsg object. This object should at minimum have two fields:

```
{ "text": "The body of the message to send back.", "channel": "the channel id" }
```

It is adequate enough to simply modify the `text` element of the `msg` passed in then reuse the rest. If you want to post to a different channel then you'll need to know the channel Id and can modify the `channel` element. You can get the channelId for other channels through methods in `util/slack`.
