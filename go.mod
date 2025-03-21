module github.com/aws/aws-sdk-go-v2/service/bedrockruntime

go 1.20

require (
	github.com/aws/aws-sdk-go-v2 v1.16.5
	github.com/aws/aws-sdk-go-v2/aws/protocol/eventstream v1.6.2
	github.com/aws/aws-sdk-go-v2/internal/configsources v1.1.12
	github.com/aws/aws-sdk-go-v2/internal/endpoints/v2 v2.4.6
	github.com/aws/smithy-go v1.20.2
)

replace github.com/aws/aws-sdk-go-v2 => github.com/aws/aws-sdk-go-v2 v1.16.5

replace github.com/aws/aws-sdk-go-v2/aws/protocol/eventstream => github.com/aws/aws-sdk-go-v2/aws/protocol/eventstream v1.6.2

replace github.com/aws/aws-sdk-go-v2/internal/configsources => github.com/aws/aws-sdk-go-v2/internal/configsources v1.1.12

replace github.com/aws/aws-sdk-go-v2/internal/endpoints/v2 => github.com/aws/aws-sdk-go-v2/internal/endpoints/v2 v2.4.6

replace github.com/aws/smithy-go => github.com/aws/smithy-go v1.19.0
