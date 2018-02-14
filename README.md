# Hello Go Lambda

A presentation and supporting code for an introduction to Go Lambdas.

## Get

```
go get -d github.com/dedelala/hello-go-lambda
```

## Lambdas

### hello

A hello world that writes a log message and is manually invoked.

### updater

A lambda that updates lambdas from S3 events. The cloudformation for this is in *two stages* due to a circular dependency.

### make targets

- `[dir-name]` - go build command
- `package` - build and copy to s3
- `stack` - deploy stack
- `kill-stack`

### make variables

- `stack` - the stack name
- `target` - an s3 uri for where to put the deployment package

To deploy yourself, you will also need to change the Lambda Code location in cloudformation.

## Slides

The presentation uses a RemarkJS docker container, `cd preso && make preso` then point your browser at `localhost:1337`.
`SIGINT` to kill.

## Credit Where Credit is Due

- Egon Elbre's gophers: [github.com/egonelbre/gophers](https://github.com/egonelbre/gophers)
- Ashley McNamara's gophers: [github.com/ashleymcnamara/gophers](https://github.com/ashleymcnamara/gophers)
- Get your own gopher: [gopherize.me](https://gopherize.me)!
