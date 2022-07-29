# ESTemplate

## What is this?

ESTemplate exposes an HTTP endpoint to render .odt template files to PDF. It uses Go lang text/template engine: https://pkg.go.dev/text/template

## How it works?

Run the server:

```
$ go run main.go
```

Call the endpoint:

```
POST /templates/render
{ "templateUrl": "{link_to_your_template_file}.odt", "data": { "name": "John", "lastName": "Doe" } }
```

Response contains your .odt template rendered with data provided. 