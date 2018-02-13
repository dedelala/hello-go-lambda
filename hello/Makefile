bin = $(shell basename $(shell pwd))
target = s3://dedelala-go-lambda/$(bin).zip
stack = dedelala-hello

.PHONY: package clean stack kill-stack

package: $(bin).zip
	aws s3 cp $< $(target)

clean:
	rm -f $(bin) $(bin).zip

$(bin): *.go
	CGO_ENABLED=0 GOOS=linux go build

$(bin).zip: $(bin)
	zip $@ $<

stack:
	aws cloudformation deploy --template-file cfn.yml --stack-name $(stack) \
	--capabilities CAPABILITY_IAM
	aws cloudformation describe-stacks --stack-name $(stack) \
	--query "Stacks[].Outputs" --output text

kill-stack:
	aws cloudformation delete-stack --stack-name $(stack)
