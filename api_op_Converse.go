// Code generated by smithy-go-codegen DO NOT EDIT.

package bedrockruntime

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/Enflick/bedrockruntime/document"
	"github.com/Enflick/bedrockruntime/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Sends messages to the specified Amazon Bedrock model. Converse provides a
// consistent interface that works with all models that support messages. This
// allows you to write code once and use it with different models. Should a model
// have unique inference parameters, you can also pass those unique parameters to
// the model. For more information, see [Run inference]in the Bedrock User Guide.
//
// This operation requires permission for the bedrock:InvokeModel action.
//
// [Run inference]: https://docs.aws.amazon.com/bedrock/latest/userguide/api-methods-run.html
func (c *Client) Converse(ctx context.Context, params *ConverseInput, optFns ...func(*Options)) (*ConverseOutput, error) {
	if params == nil {
		params = &ConverseInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "Converse", params, optFns, c.addOperationConverseMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ConverseOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ConverseInput struct {

	// The messages that you want to send to the model.
	//
	// This member is required.
	Messages []types.Message

	// The identifier for the model that you want to call.
	//
	// The modelId to provide depends on the type of model that you use:
	//
	//   - If you use a base model, specify the model ID or its ARN. For a list of
	//   model IDs for base models, see [Amazon Bedrock base model IDs (on-demand throughput)]in the Amazon Bedrock User Guide.
	//
	//   - If you use a provisioned model, specify the ARN of the Provisioned
	//   Throughput. For more information, see [Run inference using a Provisioned Throughput]in the Amazon Bedrock User Guide.
	//
	//   - If you use a custom model, first purchase Provisioned Throughput for it.
	//   Then specify the ARN of the resulting provisioned model. For more information,
	//   see [Use a custom model in Amazon Bedrock]in the Amazon Bedrock User Guide.
	//
	// [Run inference using a Provisioned Throughput]: https://docs.aws.amazon.com/bedrock/latest/userguide/prov-thru-use.html
	// [Use a custom model in Amazon Bedrock]: https://docs.aws.amazon.com/bedrock/latest/userguide/model-customization-use.html
	// [Amazon Bedrock base model IDs (on-demand throughput)]: https://docs.aws.amazon.com/bedrock/latest/userguide/model-ids.html#model-ids-arns
	//
	// This member is required.
	ModelId *string

	// Additional inference parameters that the model supports, beyond the base set of
	// inference parameters that Converse supports in the inferenceConfig field. For
	// more information, see [Model parameters].
	//
	// [Model parameters]: https://docs.aws.amazon.com/bedrock/latest/userguide/model-parameters.html
	AdditionalModelRequestFields document.Interface

	// Additional model parameters field paths to return in the response. Converse
	// returns the requested fields as a JSON Pointer object in the
	// additionalModelResultFields field. The following is example JSON for
	// additionalModelResponseFieldPaths .
	//
	//     [ "/stop_sequence" ]
	//
	// For information about the JSON Pointer syntax, see the [Internet Engineering Task Force (IETF)] documentation.
	//
	// Converse rejects an empty JSON Pointer or incorrectly structured JSON Pointer
	// with a 400 error code. if the JSON Pointer is valid, but the requested field is
	// not in the model response, it is ignored by Converse .
	//
	// [Internet Engineering Task Force (IETF)]: https://datatracker.ietf.org/doc/html/rfc6901
	AdditionalModelResponseFieldPaths []string

	// Inference parameters to pass to the model. Converse supports a base set of
	// inference parameters. If you need to pass additional parameters that the model
	// supports, use the additionalModelRequestFields request field.
	InferenceConfig *types.InferenceConfiguration

	// A system prompt to pass to the model.
	System []types.SystemContentBlock

	// Configuration information for the tools that the model can use when generating
	// a response.
	//
	// This field is only supported by Anthropic Claude 3, Cohere Command R, Cohere
	// Command R+, and Mistral Large models.
	ToolConfig *types.ToolConfiguration

	noSmithyDocumentSerde
}

type ConverseOutput struct {

	// Metrics for the call to Converse .
	//
	// This member is required.
	Metrics *types.ConverseMetrics

	// The result from the call to Converse .
	//
	// This member is required.
	Output types.ConverseOutput

	// The reason why the model stopped generating output.
	//
	// This member is required.
	StopReason types.StopReason

	// The total number of tokens used in the call to Converse . The total includes the
	// tokens input to the model and the tokens generated by the model.
	//
	// This member is required.
	Usage *types.TokenUsage

	// Additional fields in the response that are unique to the model.
	AdditionalModelResponseFields document.Interface

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationConverseMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpConverse{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpConverse{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "Converse"); err != nil {
		return fmt.Errorf("add protocol finalizers: %v", err)
	}

	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = addClientRequestID(stack); err != nil {
		return err
	}
	if err = addComputeContentLength(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addComputePayloadSHA256(stack); err != nil {
		return err
	}
	if err = addRetry(stack, options); err != nil {
		return err
	}
	if err = addRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = addRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack, options); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addSetLegacyContextSigningOptionsMiddleware(stack); err != nil {
		return err
	}
	if err = addOpConverseValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opConverse(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRecursionDetection(stack); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	if err = addDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opConverse(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "Converse",
	}
}
