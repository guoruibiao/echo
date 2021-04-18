# Echo

An useful forward proxy which support customize in CUID or UID dimension.

# modules

- UI
- Request parser
- Configuration parser
- Proxy

# how to usage

Here are four trigger positions rule, which are
```
CONFIG_TRIGGER_POSITION_HOST        = "HOST"
CONFIG_TRIGGER_POSITION_URI         = "URI"
CONFIG_TRIGGER_POSITION_QUERYSTRING = "QUERYSTRING"
CONFIG_TRIGGER_POSITION_POSTFORM    = "POSTFORM"
```

The target `request` will be customized if triggered by trigger words setting in `conf/conf.ini` at version-0.0.1

- `default`, the default host will be `httpbin.org`.

```
curl --location --request POST 'http://localhost:8080/proxy?name=tiger&triggerword=triggerget1' \
--header 'Content-Type: application/x-www-form-urlencoded' \
--data-urlencode 'cuid=tiger' \
--data-urlencode 'triggerword=triggerpost1'

{
  "args": {
    "name": "tiger",
    "triggerword": "triggerget1"
  },
  "data": "",
  "files": {},
  "form": {
    "cuid": "tiger",
    "triggerword": "triggerpost1"
  },
  "headers": {
    "Accept": "*/*",
    "Accept-Encoding": "gzip",
    "Content-Length": "35",
    "Content-Type": "application/x-www-form-urlencoded",
    "Host": "httpbin.org",
    "Http-Via": "10.254.5.91",
    "Http-X-Forwarded-For": "10.254.5.91",
    "Remote-Addr": "10.254.5.91",
    "User-Agent": "curl/7.71.1",
    "X-Amzn-Trace-Id": "Root=1-607bae51-03c5b8e5756002c4258ceae8",
    "X-Forward-For": "::1"
  },
  "json": null,
  "origin": "::1, 219.239.107.2",
  "url": "https://httpbin.org/post?name=tiger&triggerword=triggerget1"
}
```




- `trigger at query_string`, trigger position is `CONFIG_TRIGGER_POSITION_QUERYSTRING`, and trigger word is `triggerget`.
```go
// models/parser/configuration.go line:88
configurations.Items = append(configurations.Items, dao.ConfigItem{
			Type: library.CONFIG_TRIGGER_POSITION_QUERYSTRING,
			Trigger: "triggerget",
			TargetProtocol: library.PROTOCOL_HTTP,
			TargetHost: "helloworld.com",
			TargetPort: 80,
			TargetPath: "/",
		})
```

```bash
curl --location --request POST 'http://localhost:8080/proxy?name=tiger&triggerword=triggerget' \
--header 'Content-Type: application/x-www-form-urlencoded' \
--data-urlencode 'cuid=tiger' \
--data-urlencode 'triggerword=triggerpost1'
<!DOCTYPE html>
<html  lang="en" dir="ltr" prefix="content: http://purl.org/rss/1.0/modules/content/  dc: http://purl.org/dc/terms/  foaf: http://xmlns.com/foaf/0.1/  og: http://ogp.me/ns#  rdfs: http://www.w3.org/2000/01/rdf-schema#  schema: http://schema.org/  sioc: http://rdfs.org/sioc/ns#  sioct: http://rdfs.org/sioc/types#  skos: http://www.w3.org/2004/02/skos/core#  xsd: http://www.w3.org/2001/XMLSchema# ">
  <head>
    <meta charset="utf-8" />
<script>(function(i,s,o,g,r,a,m){i["GoogleAnalyticsObject"]=r;i[r]=i[r]||function(){(i[r].q=i[r].q||[]).push(arguments)},i[r].l=1*new Date();a=s.createElement(o),m=s.getElementsByTagName(o)[0];a.async=1
......
  <script type="text/javascript">window.NREUM||(NREUM={});NREUM.info={"beacon":"bam-cell.nr-data.net","licenseKey":"7bb5b98fa3","applicationID":"542748385","transactionName":"ZwBQbEJRWURZUEBRCV5Kc1tEWVhZF29wShNABF5kXl9TUmRwW1YSQgpeVFVCa3lXV1FuD1UScVdeREVYVF9RSlwKE1tdRw==","queueTime":0,"applicationTime":102,"atts":"S0dTGgpLSko=","errorBeacon":"bam-cell.nr-data.net","agent":""}</script></body>
</html>
```

- `trigger at postform`, trigger position is `CONFIG_TRIGGER_POSITION_POSTFORM`, and trigger word is `triggerpost`.
```go
// models/parser/configuration.go line:98
configurations.Items = append(configurations.Items, dao.ConfigItem{
			Type: library.CONFIG_TRIGGER_POSITION_POSTFORM,
			Trigger: "triggerpost",
			TargetProtocol: library.PROTOCOL_HTTPS,
			TargetHost: "httpbin.org",
			TargetPort: 443,
			TargetPath: "/post",
		})
```

```bash
curl --location --request POST 'http://localhost:8080/proxy?name=tiger&triggerword=triggerget1' \
--header 'Content-Type: application/x-www-form-urlencoded' \
--data-urlencode 'cuid=tiger' \
--data-urlencode 'triggerword=triggerpost'
{
  "args": {
    "name": "tiger",
    "triggerword": "triggerget1"
  },
  "data": "",
  "files": {},
  "form": {
    "cuid": "tiger",
    "triggerword": "triggerpost"
  },
  "headers": {
    "Accept": "*/*",
    "Accept-Encoding": "gzip",
    "Content-Length": "34",
    "Content-Type": "application/x-www-form-urlencoded",
    "Host": "httpbin.org",
    "Http-Via": "10.254.5.91",
    "Http-X-Forwarded-For": "10.254.5.91",
    "Remote-Addr": "10.254.5.91",
    "User-Agent": "curl/7.71.1",
    "X-Amzn-Trace-Id": "Root=1-607baf95-246d3bd73df13fe0393e5866"
  },
  "json": null,
  "origin": "::1, 219.239.107.2",
  "url": "https://httpbin.org/post?name=tiger&triggerword=triggerget1"
}
```

