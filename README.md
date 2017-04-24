# Daily Dog

![Daily Dog](https://i.imgur.com/0Uzt9VB.png)

> Get a Dog GIF delivered to a Slack channel every morning

### Running

### Building

* `docker build -t dailydog .` _\# If you want to set a global env for SLACK_WEBHOOK, edit the Dockerfile_
* `docker run -e "SLACK_WEBHOOK=https://hooks.slack.com/services/SOME_WEBHOOK/CREATED_IN/SLACK" dailydog:latest`

Without the `SLACK_WEBHOOK`, the app will fail to run.

### Deploying

You can use [now.sh]() to deploy Docker apps. You can use their initial service for free. However, there is a limit on the number of deploys and also you are assigned random URL for each deploy. So this would not be a good option for long term usage.

Use `now -e SLACK_WEBHOOK="https://hooks.slack.com/services/SOME_WEBHOOK/CREATED_IN/SLACK"` to deploy.

*Deploy Note:* I used [Zapier](https://zapier.com/), with the _schedule_ + _webhook_ zaps to ping this service each morning at a set time.

### Services

* Slack Webhook API
* [GIPHY Random API](https://github.com/Giphy/GiphyAPI#random-endpoint)
