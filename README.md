# Daily Dog

![Daily Dog](https://i.imgur.com/0Uzt9VB.png)

> Get a Dog GIF delivered to a Slack channel every morning

### Building with Docker

* `docker build -t dailydog .` _\# If you want to set a global env for SLACK_WEBHOOK, edit the Dockerfile_
* `docker run -e "SLACK_WEBHOOK=https://hooks.slack.com/services/SOME_WEBHOOK/CREATED_IN/SLACK" -p 3000:3000 dailydog:latest`

### Building with "go build"

Without the `SLACK_WEBHOOK` environment variable, **the app will fail to run**.

* `go build`
* `export SLACK_WEBHOOK=https://hooks.slack.com/services/SOME_WEBHOOK/CREATED_IN/SLACK && ./dailydog`

*With "go run":*

* `SLACK_WEBHOOK=https://hooks.slack.com/services/SOME_WEBHOOK/CREATED_IN/SLACK go run main.go`

By default, the app will run on port `3000`. Test with `curl localhost:3000` and you should see the response from the Slack webhook.

### Deploying

You can use [now.sh](https://zeit.co/now) to deploy Docker apps. You can use their initial service for free. However, there is a limit on the number of deploys and also you are assigned random URL for each deploy. So this would not be a good option for long term usage.

Use `now -e SLACK_WEBHOOK="https://hooks.slack.com/services/SOME_WEBHOOK/CREATED_IN/SLACK"` to deploy.

*Deploy Note:* I used [Zapier](https://zapier.com/), with the _schedule_ + _webhook_ zaps to ping this service each morning at a set time.

### Services

* [Slack Webhook API](https://api.slack.com/incoming-webhooks)
* [GIPHY Random API](https://github.com/Giphy/GiphyAPI#random-endpoint)
