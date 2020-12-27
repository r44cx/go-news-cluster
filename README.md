# RSS Clustering

## (IN DEV)

## Configuration

File: `config.yml`

Example:
``` yaml
globals:
 interval: 5
 valid: 0.9
hosting:
 host: localhost
 port: 8001
feeds:
- URL: https://r44cx.zil/news.rss
  tag: "item"
  rate: 1800
- URL: https://r44cx.zil/blog.rss
  tag: "item"
  rate: 86400
``` 


## Build
`go build main.go`


## Make
