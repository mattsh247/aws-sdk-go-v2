module github.com/aws/aws-sdk-go-v2/config

go 1.15

require (
	github.com/aws/aws-sdk-go-v2 v0.31.1-0.20210105194811-58b543144e2a
	github.com/aws/aws-sdk-go-v2/credentials v0.2.0
	github.com/aws/aws-sdk-go-v2/feature/ec2/imds v0.1.0
	github.com/aws/aws-sdk-go-v2/service/sts v0.31.0
	github.com/aws/smithy-go v0.5.1-0.20210104190327-c7045c94c1ec
	github.com/google/go-cmp v0.5.4
)

replace (
	github.com/aws/aws-sdk-go-v2 => ../
	github.com/aws/aws-sdk-go-v2/credentials => ../credentials/
	github.com/aws/aws-sdk-go-v2/feature/ec2/imds => ../feature/ec2/imds/
	github.com/aws/aws-sdk-go-v2/service/sts => ../service/sts/
)
