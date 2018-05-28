# go-telenor-auth-example

A project to show a simple implementation of [go-telenor-auth](https://github.com/ExploratoryEngineering/go-telenor-auth). Excellent for debugging and testing your available APIs towards api.telenor.no.

## How to use
1. Download the newest [releases](https://github.com/ExploratoryEngineering/go-telenor-auth-example/releases) as an exectuable binary for Linux/OSX or Windows
2. Add a `config.json` file with the client id and client secret. See [configuration](#configuration) for more details

The server exposes the following endpoints:

*Open endpoints*
- `/`: Simple landing page which signals that everything is OK
- `/auth/login`: Go to this URL to log in as a user with the credentials towards api.telenor.no
- `/auth/logout`: Going to this URL will clear your current session

*Closed endpoints*
- `/api/*`: Will proxy a request towards the api.telenor.no
- `/secure/*`: An example path which is closed by default untill you authorize yourself through a login

The `/api/*` endpoint maps directly towards api.telenor.no and populates the request with a user session initiated through `/auth/login` with the configured credentials found in `config.json`. This allows for simple testing the live APIs without setting up a full production setup.

## Configuration
You can run the exectuable or built file by running

```bash
./go-telenor-auth
```

By default it will look for a configuration file named `config.json` which contains the following structure
```json
{
  "clientId": "<Your client ID from developer.telenor.no>",
  "clientSecret": "<Your client secret from developer.telenor.no>"
}
```

It will then host the server on :8080 and is available from the browser on localhost:8080.


### Overriding configuration
You can override the the configuration by using the following params

Configuration file name:
```bash
./go-telenor-auth -c <your-name-of-config.json>
```

Use id and secret as params:
```bash
./go-telenor-auth -apigee-client-id my-id apigee-client-secret my-secret
```

## Build

Fetch dependencies
```bash
go get -u
``` 

Build
```bash
go build -o go-telenor-auth
```