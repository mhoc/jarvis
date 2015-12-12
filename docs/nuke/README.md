
# Command: `nuke`

This command requires you to also have a lambda function defined. I have the source code for it below.

```
var http = require('https')
console.log('Loading function');

exports.handler = function(event, context) {
    console.log('Received event:', JSON.stringify(event, null, 2));
    options = {
    	hostname: 'slack.com',
    	port: 443,
    	path: "/api/chat.postMessage?token=" + event.token + "&channel=" + event.channel + "&text=" + encodeURIComponent(event.text),
    	method: 'POST',
    	headers: {
		    'Content-Type': 'application/json',
	    }
    }
    setInterval(function() {
        req = http.request(options, function(res) {
        	console.log('STATUS: ' + res.statusCode);
        	console.log('HEADERS: ' + JSON.stringify(res.headers));
          	res.setEncoding('utf8');
          	res.on('data', function (chunk) {
          	    console.log(chunk)
            // 	context.success()
          	})
          	res.on('end', function() {
            	console.log('No more data in response.')
            // 	context.success()
          	})
        })
        req.end()
    }, 500)
};
```

Yup, that has no `context.success()` call. It will keep repeating for however long you set your lambda instance to timeout at. That's what we want, yo. Maximum bills.
