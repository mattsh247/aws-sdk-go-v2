module github.com/aws/aws-sdk-go-v2/service/ec2

go 1.15

require (
	github.com/aws/aws-sdk-go-v2 v1.20.1
	github.com/aws/aws-sdk-go-v2/internal/configsources v1.1.38
	github.com/aws/aws-sdk-go-v2/internal/endpoints/v2 v2.4.32
	github.com/aws/aws-sdk-go-v2/service/internal/presigned-url v1.9.32
	github.com/aws/smithy-go v1.14.1
	github.com/google/go-cmp v0.5.8
	github.com/jmespath/go-jmespath v0.4.0
)

replace github.com/aws/aws-sdk-go-v2 => ../../

replace github.com/aws/aws-sdk-go-v2/internal/configsources => ../../internal/configsources/

replace github.com/aws/aws-sdk-go-v2/internal/endpoints/v2 => ../../internal/endpoints/v2/

replace github.com/aws/aws-sdk-go-v2/service/internal/presigned-url => ../../service/internal/presigned-url/

replace github.com/aws/smithy-go => /Users/wty/Projects/aws-sdk-go-v2/../smithy-go
