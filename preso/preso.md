# Hello Go Lambda

.center.lambda[λ]
.center[![Gopher!](/preso/gopher-sprite.gif)]

--

???
I want to show you a Go Lambda. Then, I want to show you one that does something.
There will be Go, there will be Lambdas, there will be AWS SDK calls. Exciting stuff.

Why would you want to know about this? I don't know, do you want to know about this?
Probably! I mean, we are all nerds here right? If you write Go and use AWS then you definitely
want to know about it, maybe you're even using Go Lambdas already?

If you don't write Go, I hope the simplicity will win you over. If you're not using Lambda,
I've got nothing for you sorry. Lambda may or may not be right for you, who am I to say?
I'm not the implementation police.

---

# Performance

<br/><br/>
.center.img[![Gopher Star Wars](/preso/GOPHER_STAR_WARS.png)]

???
Performance! Everyone wants to know.
I'm going to get this out of the way right at the start: it's better.

Did I measure this? No. That would be a waste of effort.

Go is faster than Python and Node. Fight me.

If you want to race binaries against interpreters go right ahead, I'm good. There's an SDK
available for x-ray, go nuts.

---

# Performance

<br/><br/>
.center.img[![Gopher Mic Drop](/preso/GOPHER_MIC_DROP.png)]

???
I suppose I should state my bias if I haven't clearly enough already: I love Go, it's my favourite
programming language. The thing I love most about it is it's conciseness.

Less typing, easier reading.

Just watch Rob Pike talk about it, really, if he can't sell you on it no-one can.

---

# Hello Go Lambda

<br/><br/><br/><br/>
```sh
go get -d github.com/dedelala/hello-go-lambda
```

???

Cool? Cool. I need to give you a disclaimer: this presentation is all code. I'm wary of this
but I think good go speaks for itself.

If you're super keen you can go get this right now and have a poke through while I'm talking.

---

# Hello Go Lambda

```go
package main

import (
	"log"

	"github.com/aws/aws-lambda-go/lambda"
)

// Go handles a Lambda invocation.
func Go() {
	log.Println("Hello log world!")
}

func main() {
	lambda.Start(Go)
}
```

???
Using the lambda runtime for go is super easy. Here's a minimal lambda. All it does is write a log message.

When this runs on Lambda it will run like any other Go binary and enter at `main`.

---

# Hello Go Lambda

```go
package main

import (
	"log"

	"github.com/aws/aws-lambda-go/lambda"
)

// Go handles a Lambda invocation.
func Go() {
	log.Println("Hello log world!")
}

func `main()` {
	lambda.Start(Go)
}
```

???
The body of `main` has the wiring of the actual lambda invocation to a handler method.

---

# Hello Go Lambda

```go
package main

import (
	"log"

	"github.com/aws/aws-lambda-go/lambda"
)

// Go handles a Lambda invocation.
func Go() {
	log.Println("Hello log world!")
}

func main() {
	`lambda.Start(Go)`
}
```

???
`lambda.Start` blocks and does't return.

The handler method passed to `lambda.Start`,

---

# Hello Go Lambda

```go
package main

import (
	"log"

	"github.com/aws/aws-lambda-go/lambda"
)

// Go handles a Lambda invocation.
func `Go()` {
	log.Println("Hello log world!")
}

func main() {
	lambda.Start(Go)
}
```

???

in this case `Go`, may take arguments and
also return data and I will get to that soon enough.

---

# Hello Go Lambda

```go
package main

import (
	"log"

	"github.com/aws/aws-lambda-go/lambda"
)

// Go handles a Lambda invocation.
func Go() {
	log.Println("Hello log world!")
}

func main() {
	lambda.Start(Go)
}
```

???

Go's log package writes to stderr, there's nothing special going on there.

So then we build this...

---

# Hello Go Lambda

```go
package main

import (
	"log"

	"github.com/aws/aws-lambda-go/lambda"
)

// Go handles a Lambda invocation.
func Go() {
	log.Println("Hello log world!")
}

func main() {
	lambda.Start(Go)
}
```

```javascript
CGO_ENABLED=0 GOOS=linux go build
```

???
- The build command in the docs is _almost_ right


---

# Hello Go Lambda

```go
package main

import (
	"log"

	"github.com/aws/aws-lambda-go/lambda"
)

// Go handles a Lambda invocation.
func Go() {
	log.Println("Hello log world!")
}

func main() {
	lambda.Start(Go)
}
```

```javascript
`CGO_ENABLED=0` GOOS=linux go build
```

???
they left out the part where it's a static binary.

---

# Hello Go Lambda

```go
package main

import (
	"log"

	"github.com/aws/aws-lambda-go/lambda"
)

// Go handles a Lambda invocation.
func Go() {
	log.Println("Hello log world!")
}

func main() {
	lambda.Start(Go)
}
```

```javascript
CGO_ENABLED=0 GOOS=linux go build
```

???
Lambda users, question: have you ever had a mad fight with the Lambda execution environment?

And by mad fight I mean epic battle, like, Helm's Deep. You are Gandalf and the Lambda runtime is
orcs and you're just flailing about with a big stick hoping for the best.

Maybe there was a missing dependency, or some minor version difference that broke everything?

Did you have to pip or npm install the entire freakin internet into a zip file?

I have three words for you: go static binary. There is no documentation page for "how to package
your go lambda"... jussayin.

--

```javascript
zip hello.zip hello
```

???

Then all you gotta do is zip the binary. All dependencies built-in. Easy. I could drag half of github
into my project and those two steps will be exactly the same.

You can expect a static binary for a small lambda to fall in the 8-12 megabyte range, which is
probably a larger package than you're accustomed to *but* that's the trade. I'll pay.

They compress quite well, you should expect something like two thirds reduction.

There are even libraries that let you build a filesystem *inside your static binary*,
if you need external files for templates or whatever.

To test this, there's no package to mock the runtime I would just need to call the handler
from my tests.

Cool? Cool. Now let's look at some CloudFormation.

---

# CloudFormation

```yaml
Lambda:
  Type: "AWS::Lambda::Function"
  Properties:
    Description: Writes a log message.
    Handler: hello
    Role:
      Fn::GetAtt: LambdaRole.Arn
    Runtime: go1.x
    Timeout: 3
    Code:
      S3Bucket: dedelala-go-lambda
      S3Key: hello.zip
```

???
Here is the Lambda function. The whole template has the minimum required to invoke the lambda manually.

---

# CloudFormation

```yaml
Lambda:
  Type: "AWS::Lambda::Function"
  Properties:
    Description: Writes a log message.
    `Handler: hello`
    Role:
      Fn::GetAtt: LambdaRole.Arn
    Runtime: go1.x
    Timeout: 3
    Code:
      S3Bucket: dedelala-go-lambda
      S3Key: hello.zip
```

???
The handler, hello, is the name of the go binary in the zip.

---

# CloudFormation

```yaml
Lambda:
  Type: "AWS::Lambda::Function"
  Properties:
    Description: Writes a log message.
    Handler: hello
    `Role:`
      Fn::GetAtt: `LambdaRole.Arn`
    Runtime: go1.x
    Timeout: 3
    Code:
      S3Bucket: dedelala-go-lambda
      S3Key: hello.zip
```

???
The lambda role is a basic execution role with logging permission.

---

# CloudFormation

```yaml
Lambda:
  Type: "AWS::Lambda::Function"
  Properties:
    Description: Writes a log message.
    Handler: hello
    Role:
      Fn::GetAtt: LambdaRole.Arn
    `Runtime: go1.x`
    Timeout: 3
    Code:
      S3Bucket: dedelala-go-lambda
      S3Key: hello.zip
```

???
The runtime is go1.x hooray!

---

# CloudFormation

```yaml
Lambda:
  Type: "AWS::Lambda::Function"
  Properties:
    Description: Writes a log message.
    Handler: hello
    Role:
      Fn::GetAtt: LambdaRole.Arn
    Runtime: go1.x
    `Timeout: 3`
    Code:
      S3Bucket: dedelala-go-lambda
      S3Key: hello.zip
```

???
The function will run in a lot less than 3...

---

# CloudFormation

```yaml
Lambda:
  Type: "AWS::Lambda::Function"
  Properties:
    Description: Writes a log message.
    Handler: hello
    Role:
      Fn::GetAtt: LambdaRole.Arn
    Runtime: go1.x
    Timeout: 3
    `Code:`
      S3Bucket: dedelala-go-lambda
      S3Key: `hello.zip`
```

???
And the code is in the place!

---


# Hello Go Lambda

<br/><br/><br/><br/>
.center[![Gopher dance!](/preso/gopher-dance.gif)]

???
Cool! I have made a completely useless lambda. So how about a Lambda that does something?


---

# Do Something

<br/><br/><br/><br/>
.center[![Gopher dance!](/preso/gopher-dance.gif)]

???
I want a Lambda that can update Lambdas.
As the Lambda peeps will know, if I were to change the Lambda I just showed you and re-package
it, I would also need to update the function code, a choice between calling some giant
aws command or some voodoo with versioned buckets and references to object versions in
cloudformation.

So, to that end, how about a Lambda that triggers an update function code call from an S3 Event?

---

# Do Something

<br/><br/><br/><br/>
.center[.lambda[S3]![Gopher dance!](/preso/gopher-dance.gif).lambda[￫λ]]

???
Then all you gotta do is put the thing in a bucket and blammo: up to date.

Let's have a look...

---

# Do Something

```go
package main

import (
	"github.com/aws/aws-lambda-go/events"
	runtime "github.com/aws/aws-lambda-go/lambda"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/external"
	"github.com/aws/aws-sdk-go-v2/service/lambda"
)
```

???
So for starters we have a bunch more imports... the main thing to note is that there are two
packages called `lambda`.

---

# Do Something

```go
package main

import (
	"github.com/aws/aws-lambda-go/events"
	runtime "github.com/aws/aws-lambda-go/`lambda`"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/external"
	"github.com/aws/aws-sdk-go-v2/service/`lambda`"
)
```

???
Yes, AWS publishes two Go packages called `lambda`.

---

<br/><br/><br/><br/><br/><br/>
.center.img[![Picard Gopher!](/preso/Facepalm_Picard_Gopher.png)]

---

# Do Something

```go
package main

import (
	"github.com/aws/aws-lambda-go/events"
	runtime "github.com/aws/aws-lambda-go/`lambda`"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/external"
	"github.com/aws/aws-sdk-go-v2/service/`lambda`"
)
```

???
You will only run into this if you want to write a Lambda that does stuff to Lambdas.
I could and probably would split the sdk calling stuff into another file,
or even a package if this grew any larger.

That's a pattern my crew has been working at MYOB. We have been building out packages to handle
all the messy things we need to do along with simple command line clients.

Now that the go runtime is here it will be a relatively simple wiring excercise to turn everything
into Lambda. And like I said before, packaging those lambdas will be no different than packaging
the hello world that I just showed you. Thank Go for static binaries.

Anyway, two packages called lambda.

---

# Do Something

```go
package main

import (
	"github.com/aws/aws-lambda-go/events"
	`runtime` "github.com/aws/aws-lambda-go/lambda"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/external"
	"github.com/aws/aws-sdk-go-v2/service/lambda"
)
```

???
I've imported the one we already looked at as `runtime`.

---

# Do Something

```go
var svc *lambda.Lambda

func init() {
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		panic(err)
	}

	svc = lambda.New(cfg)
}

func main() {
	`runtime`.Start(Go)
}
```

???
So the one line in `main` changes to `runtime.Start`.

---

# Do Something

```go
var svc *lambda.Lambda

func `init()` {
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		panic(err)
	}

	svc = lambda.New(cfg)
}

func main() {
	runtime.Start(Go)
}
```

???
I have added an init function. Have you used init in go before? It's
a function that will run once when the package is initialised.

---

# Do Something

```go
var svc *lambda.Lambda

func init() {
	cfg, err := `external.LoadDefaultAWSConfig()`
	if err != nil {
		panic(err)
	}

	svc = lambda.New(cfg)
}

func main() {
	runtime.Start(Go)
}
```

???
The AWS config only needs to be loaded once so it makes sense to do it here.

I'm hesitant to panic, but init can't return an error so I have no choice.
If there was ever a reason to panic, missing AWS config in the Lambda runtime would be a good one.

---

# Do Something

```go
`var svc *lambda.Lambda`

func init() {
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		panic(err)
	}

	svc = `lambda.New(cfg)`
}

func main() {
	runtime.Start(Go)
}
```

???
I get a pointer to `lambda.Lambda` with `lambda.New` and stick it in a package variable.

Remember, the `lambda` we are referring to here is the *SDK service lambda*,
and not the *lambda runtime*. Two packages.

For every service package in the SDK there's an iface package that allows you to mock the service object.

For Lambda the package is lambda/lambdaiface. That's SDK service lambda. Two packages.

This is a pattern for the SDK, and the service objects are safe to use concurrently.

Using goroutines you can make thousands of calls in seconds. Maybe not so important
in this context, but trust me you will find an application for that.

---

# Go

```go
// Go handles a Lambda invocation for an S3 event.
func Go(ev events.S3Event) error {
	for _, r := range ev.Records {
		fn := r.S3.Object.Key
		fn = fn[:len(fn)-4]

		err := Update(fn, r.S3.Bucket.Name, r.S3.Object.Key)
		if err != nil {
			return err
		}
	}
	return nil
}
```

???
The handler function looks different too,


---

# Go

```go
// Go handles a Lambda invocation for an S3 event.
func Go(`ev events.S3Event`) error {
	for _, r := range ev.Records {
		fn := r.S3.Object.Key
		fn = fn[:len(fn)-4]

		err := Update(fn, r.S3.Bucket.Name, r.S3.Object.Key)
		if err != nil {
			return err
		}
	}
	return nil
}
```

???
look we are accepting a struct

---

# Go

```go
// Go handles a Lambda invocation for an S3 event.
func Go(ev events.S3Event) `error` {
	for _, r := range ev.Records {
		fn := r.S3.Object.Key
		fn = fn[:len(fn)-4]

		err := Update(fn, r.S3.Bucket.Name, r.S3.Object.Key)
		if err != nil {
			return err
		}
	}
	return nil
}
```

???
and possibly returning an
error, how exciting!

---

# Go

```go
// Go handles a Lambda invocation for an S3 event.
func Go(ev events.S3Event) error {
	for _, r := range ev.Records {
		fn := r.S3.Object.Key
		fn = fn[:len(fn)-4]

		err := Update(fn, r.S3.Bucket.Name, r.S3.Object.Key)
		if err != nil {
			return err
		}
	}
	return nil
}
```

???
If you look at the AWS examples for Go Lambdas, they'll have you writing structs to accept
whatever you like in no time.

The handler function is allowed to accept one object that can be unmarshaled from JSON.
Marshaling is go talk for serialiaing. You can return one object that can be marshaled into
JSON, in addition to an error.

For bonus points you may also accept a Go Context which can be used for graceful termination.

Any combination of those is allowed, full details are in the godoc.

---

# Do Something

```go
package main

import (
	`"github.com/aws/aws-lambda-go/events"`
	runtime "github.com/aws/aws-lambda-go/lambda"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/external"
	"github.com/aws/aws-sdk-go-v2/service/lambda"
)
```

???
Fortunately, AWS have supplied an `events` package with pre-defined structs for every
AWS event.

So clearly I will just use that.

There's also an `events/test` package with some functions to help with testing.

---

# Go

```go
// Go handles a Lambda invocation for an S3 event.
func Go(ev `events.S3Event`) error {
	for _, r := range ev.Records {
		fn := r.S3.Object.Key
		fn = fn[:len(fn)-4]

		err := Update(fn, r.S3.Bucket.Name, r.S3.Object.Key)
		if err != nil {
			return err
		}
	}
	return nil
}
```

???
Package `events` has an `S3Event`. Cool. Incidentally, the pre-defined
structs make it super easy to figure out what the heck you are going to get passed
when this thing runs. You can just look at the code.

---

# Go

```go
// Go handles a Lambda invocation for an S3 event.
func Go(ev events.S3Event) error {
	for _, r := range `ev.Records` {
		fn := r.S3.Object.Key
		fn = fn[:len(fn)-4]

		err := Update(fn, r.S3.Bucket.Name, r.S3.Object.Key)
		if err != nil {
			return err
		}
	}
	return nil
}
```

???
The S3 Event has an array of records

---

# Go

```go
// Go handles a Lambda invocation for an S3 event.
func Go(ev events.S3Event) error {
	for _, r := range ev.Records {
		`fn := r.S3.Object.Key`
		`fn = fn[:len(fn)-4]`

		err := Update(fn, r.S3.Bucket.Name, r.S3.Object.Key)
		if err != nil {
			return err
		}
	}
	return nil
}
```

???
for each of those I will crudely hack the last
four characters off the object key

---

# Go

```go
// Go handles a Lambda invocation for an S3 event.
func Go(ev events.S3Event) error {
	for _, r := range ev.Records {
		fn := r.S3.Object.Key
		fn = fn[:len(fn)-4]

		err := `Update(fn, r.S3.Bucket.Name, r.S3.Object.Key)`
		if err != nil {
			return err
		}
	}
	return nil
}
```

???
and blindly try to update a lambda with that name.

Hey, I never said it would be production grade... but it will work.

In CloudFormation the trigger is configured to only fire for `.zip`.

---

# Update

```go
// Update updates lambda function code with an S3 object.
func Update(f, b, k string) error {
	in := &lambda.UpdateFunctionCodeInput{
		FunctionName: aws.String(f),
		S3Bucket:     aws.String(b),
		S3Key:        aws.String(k),
	}

	req := svc.UpdateFunctionCodeRequest(in)
	_, err := req.Send()
	return err
}
```

???
Here's the code I'm calling to update the lambda.

The universal pattern for using the AWS SDK for Go is as follows:

---

# Update

```go
// Update updates lambda function code with an S3 object.
func Update(f, b, k string) error {
	in := `&lambda.UpdateFunctionCodeInput`{
		FunctionName: aws.String(f),
		S3Bucket:     aws.String(b),
		S3Key:        aws.String(k),
	}

	req := svc.UpdateFunctionCodeRequest(in)
	_, err := req.Send()
	return err
}
```

???
One. create a pointer to a cumbersome `Input` struct, in this case a `lambda.UpdateFunctionCodeInput`

---

# Update

```go
// Update updates lambda function code with an S3 object.
func Update(f, b, k string) error {
	in := &lambda.UpdateFunctionCodeInput{
		FunctionName: aws.String(f),
		S3Bucket:     aws.String(b),
		S3Key:        aws.String(k),
	}

	req := `svc.UpdateFunctionCodeRequest(in)`
	_, err := req.Send()
	return err
}
```

???
Two. generate a `Request` struct using the service object we created in init,
`req` is of type `lambda.UpdateFunctionCodeRequest`.

---

# Update

```go
// Update updates lambda function code with an S3 object.
func Update(f, b, k string) error {
	in := &lambda.UpdateFunctionCodeInput{
		FunctionName: aws.String(f),
		S3Bucket:     aws.String(b),
		S3Key:        aws.String(k),
	}

	req := svc.UpdateFunctionCodeRequest(in)
	_, err := `req.Send()`
	return err
}
```

???
Three. `Send` the request, get a lumbering great `Output` struct and an error back.

I'm discarding the output struct in this case and just returning the error.
If I kept the output it would be a pointer to a `lambda.UpdateFunctionCodeOutput`.

The SDK is *verbose*! I had to import three packages to do this one thing!

---

# Update

```go
// Update updates lambda function code with an S3 object.
func Update(f, b, k string) error {
	in := &lambda.UpdateFunctionCodeInput{
		FunctionName: `aws.String`(f),
		S3Bucket:     `aws.String`(b),
		S3Key:        `aws.String`(k),
	}

	req := svc.UpdateFunctionCodeRequest(in)
	_, err := req.Send()
	return err
}
```

???
They have even implemented their own primitive types.
It makes an interesting task of writing minimal code.

---

# Update

```go
// Update updates lambda function code with an S3 object.
func Update(f, b, k string) error {
	in := &lambda.UpdateFunctionCodeInput{
		FunctionName: aws.String(f),
		S3Bucket:     aws.String(b),
		S3Key:        aws.String(k),
	}

	req := svc.UpdateFunctionCodeRequest(in)
	_, err := req.Send()
	return err
}
```

???
In fact the main source file for sdk s3 I've been working with is 15,000 lines long.

It's all generated from schema.

Still, it's not too hard to write simple functions that wrangle the SDK.

---

# Packages!

`github.com/aws/aws-lambda-go` - the lambda runtime

`github.com/aws/aws-xray-sdk-go` - the x-ray sdk

`github.com/aws/aws-sdk-go-v2` - the aws sdk - beware all ye who enter here

.center.bigimg[![Gophers!](/preso/GOPHER_LEARN.png)]

???
So that's Go Lambdas! With a side of SDK. I hope you have enjoyed looking at slides upon slides
of code, somehow, or maybe at least got some inspo to do something great with these.

You can godoc.org slash any of the above for all the gory details. I recommend the godoc over AWS's
docs which is where you will end up if you go googling.

---

# Things!

`github.com/egonelbre/gophers` - for gophers

`github.com/ashleymcnamara/gophers` - for gophers

`gopherize.me` - for your own gopher

`github.com/dedelala/hello-go-lambda` - for this

.center.img[![Unicorn Gopher!](/preso/Unicorn_Gopher.png)]

???

Thanks!

