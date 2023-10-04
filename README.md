# Golang Webhook

A simple app for Slack, implemented with [Go](https://go.dev/), that uses [incoming webhooks](https://api.slack.com/messaging/webhooks) to post jokes on Slack and to inform about changes on Bitcoin prices.

## Usage

To use this app, you'll need to:

1. Have admin access to a Slack Workspace.
2. Create a Slack app (which can be done [here](https://api.slack.com/apps/new)), selecting the option "From an app manifest" and then copying the content from the `manifest.json` file from this repository.
3. [Enable](https://api.slack.com/messaging/webhooks#enable_webhooks) incoming webooks on your workspace, if not already enabled.
4. [Create](https://api.slack.com/messaging/webhooks#create_a_webhook) a new incoming webhook and copy the Webhook URL.
5. On the `.env` file, replace the _\<insert-your-webhook-url-here\>_ text with your Webhook URL.
6. Run the application by executing `go run main`.

If you use the `--cron` parameter when running the application (e.g. `go run main --cron`) then the messages will only be send when some threshold are meet:

* Jokes will only be send if the previous jokes was sent more than a day ago.
* Bitcoin prices will only be send if either the previous price was sent more than a day ago or if the price changed more than a 10% compared to the previous time. 

This feature is useful if you want to configure a Cron job to execute the application periodically.

> Threshold values can be configured using the `.env` file.

## Dependencies

Jokes are fetched from [JokeAPI](https://sv443.net/jokeapi/v2/) while Bitcoin prices are obtained from [Coinbase](https://docs.cloud.coinbase.com/exchange/).

## License

Golang Webhook is free software; you can redistribute it and/or modify it under the terms of the Mozilla Public License v2.0. You should have received a copy of the MPL 2.0 along with this software, otherwise you can obtain one at http://mozilla.org/MPL/2.0/.
