Monigor is a simple website monitor/scraper that posts results to Slack.

## Installation

`go get github.com/sekl/monigor`

There are currently only two dependencies: `github.com/PuerkitoBio/goquery` and `github.com/parnurzeal/gorequest`
If necessary, get them with `go get -d ./...` from the project root.

## Setup

Setup a webhook for your slack team and edit `config.json` with your details:

    {
        "webhook_url": "https://hooks.slack.com/services/XXX",
        "channel": "#general",
        "bot_name": "Monigor-Bot"
    }

List the sites you want to monitor and the elements you want to search for in `urls.json` (using CSS selectors).

## To Do:
- continuous monitoring
- support for more ways of finding elements or simple text
- adding/removing sites on demand
- other ways of alerting besides Slack
