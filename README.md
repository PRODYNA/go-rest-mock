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


## TODOs

* Flags for cli command, path for json files, debug mode etc.
* Configured resposne for errors
* Json validation flag for request
* Body with different content than json with property bodyRef
* Dockerfile
* More tests
* Using go template engine for dynamic responses
* Using go plugins (go 1.8 needd) for more dynamic responses

## Precedence of matching

1. Port, HTTP Method and content type must match.
2. If a static path matches, it will be used
3. If a template path matches, it will be used
4. If a "_default" path exists, it will be used
5. Error is returned
 

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
