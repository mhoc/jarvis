
Deploying Jarvis is a multi-step process.

# 1. Register a bot with Slack

This will give you an authorization token which you will give Jarvis.

# 2. Register for the various APIs jarvis uses.

Look at `src/jarvis/config/yaml.go` for the guaranteed complete list, or otherwise check out these:

* Slack (obviously)
* ZipCode API
* DarkSky Weather
* AWS (for Lambda functionality)

# 3. Pull the code and build it

You'll need

* git
* golang
* gb (go build)

```
git clone http://github.com/mhoc/jarvis
make
```

# 4. Deploy redis

Jarvis uses redis to cache and store any persistent information it needs to. You can host it wherever you want.

# 5. Set up `config.yaml`

This file defines all the operational parameters under which jarvis is run, and it is required. Check out the `config.yaml` file in this folder for an example setup.

# 6. 
