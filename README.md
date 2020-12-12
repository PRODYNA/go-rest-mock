
[![Project status](https://img.shields.io/badge/version-v0-green.svg)](https://github.com/prodyna/go-rest-mock/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/prodyna/go-rest-mock)](https://goreportcard.com/report/github.com/prodyna/go-rest-mock)
[![License](https://img.shields.io/dub/l/vibe-d.svg)](LICENSE)
[![Coverage Status](https://coveralls.io/repos/github/PRODYNA/go-rest-mock/badge.svg?branch=main)](https://coveralls.io/github/PRODYNA/go-rest-mock?branch=main)
# go-rest-mock
Simple, but powerful mock tool for ReST.

**Usage**
* Create a json file, in the test/data (see todos) directory
* Specify port and an identifier
* Add static content with static paths or use a templated path
* Set content type correctly
* Set the response with body as json
* Set the response content type
* Set the response status code
* Set headers if needed

## Build

Simply call ```go build -o grmock cmd/main.go``` or ```go run cmd/main.go```

## Supported features

* All HTTP methods for request and response
* All HTTP status Codes
* Matching of static paths in URL like /user/45566/account/4455
* Matching of template paths in URL like /user/{uid}/account/{aid}
* Multiple ports (one port for each mock configuration)
* Setting of response headers
* Setting of content type of the response
* Request Validation (Json)


## TODOs

* Flags for cli command, debug mode
* Configured response for errors
* Json validation flag for request
* Body with different content than json with property bodyRef
* Dockerfile
* More tests
* Using go template engine for dynamic responses
* Using go plugins for more dynamic responses

## Precedence of matching

1. Port, HTTP Method and content type must match.
2. If a static path matches, it will be used
3. If a template path matches, it will be used
4. If a "_default" path exists, it will be used
5. An error will be returned
 
## Build the mock server

``` go build -o mock cmd/main.go ```
 
## Start the mock server

``` ./mock --path ./data ```

## Example configuration

This can be called with e.g. http://localhost:9000/account/1234 in a browser.

```
{
  "id": "backend",
  "port": "9000",
  "paths": [
    {
      "method": "GET",
      "path": "/account/{id}",
      "contentType": "",
      "response": {
        "status": 200,
        "contentType": "application/json",
        "body": {
          "hello": "from backend 1"
        },
        "header": {
          "XX": "test 1"
        }
      }
    }
  ]
}

```

## Small load test for the mock tool

```
ab -n 100000 -c 1000 http://localhost:9000/account/1
```

```
This is ApacheBench, Version 2.3 <$Revision: 1843412 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking localhost (be patient)
Completed 10000 requests
...
Finished 100000 requests


Server Software:        
Server Hostname:        localhost
Server Port:            9000

Document Path:          /account/1
Document Length:        26 bytes

Concurrency Level:      1000
Time taken for tests:   5.921 seconds
Complete requests:      100000
Failed requests:        0
Total transferred:      14600000 bytes
HTML transferred:       2600000 bytes
Requests per second:    16889.24 [#/sec] (mean)
Time per request:       59.209 [ms] (mean)
Time per request:       0.059 [ms] (mean, across all concurrent requests)
Transfer rate:          2408.04 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0   27   4.1     27      50
Processing:    10   32   5.8     31      71
Waiting:        1   23   5.7     22      62
Total:         33   59   4.6     58      87

Percentage of the requests served within a certain time (ms)
  50%     58
  66%     59
  75%     61
  80%     62
  90%     65
  95%     68
  98%     71
  99%     72
 100%     87 (longest request)


```
