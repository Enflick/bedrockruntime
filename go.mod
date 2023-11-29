module github.com/aws/aws-sdk-go-v2/service/bedrockruntime

go 1.19

require (
	github.com/aws/aws-sdk-go-v2 v1.23.3
	github.com/aws/aws-sdk-go-v2/aws/protocol/eventstream v1.5.2
	github.com/aws/aws-sdk-go-v2/internal/configsources v1.2.6
	github.com/aws/aws-sdk-go-v2/internal/endpoints/v2 v2.5.6
	github.com/aws/smithy-go v1.18.0
	github.com/google/go-cmp v0.5.8
)

replace github.com/aws/aws-sdk-go-v2 => ../../

replace github.com/aws/aws-sdk-go-v2/aws/protocol/eventstream => ../../aws/protocol/eventstream/

replace github.com/aws/aws-sdk-go-v2/internal/configsources => ../../internal/configsources/

replace github.com/aws/aws-sdk-go-v2/internal/endpoints/v2 => ../../internal/endpoints/v2/
