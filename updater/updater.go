package main

import (
	"github.com/aws/aws-lambda-go/events"
	runtime "github.com/aws/aws-lambda-go/lambda"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/external"
	"github.com/aws/aws-sdk-go-v2/service/lambda"
)

var svc *lambda.Lambda

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

func init() {
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		panic(err)
	}

	svc = lambda.New(cfg)
}

func main() {
	runtime.Start(Go)
}
