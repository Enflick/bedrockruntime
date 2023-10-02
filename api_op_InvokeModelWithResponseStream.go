// Code generated by smithy-go-codegen DO NOT EDIT.

package bedrockruntime

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	internalauth "github.com/aws/aws-sdk-go-v2/internal/auth"
	"github.com/aws/aws-sdk-go-v2/service/bedrockruntime/types"
	smithyendpoints "github.com/aws/smithy-go/endpoints"
	"github.com/aws/smithy-go/middleware"
	smithysync "github.com/aws/smithy-go/sync"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"sync"
)

// Invoke the specified Bedrock model to run inference using the input provided.
// Return the response in a stream. For more information, see Run inference (https://docs.aws.amazon.com/bedrock/latest/userguide/api-methods-run.html)
// in the Bedrock User Guide. For an example request and response, see Examples
// (after the Errors section).
func (c *Client) InvokeModelWithResponseStream(ctx context.Context, params *InvokeModelWithResponseStreamInput, optFns ...func(*Options)) (*InvokeModelWithResponseStreamOutput, error) {
	if params == nil {
		params = &InvokeModelWithResponseStreamInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "InvokeModelWithResponseStream", params, optFns, c.addOperationInvokeModelWithResponseStreamMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*InvokeModelWithResponseStreamOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type InvokeModelWithResponseStreamInput struct {

	// Inference input in the format specified by the content-type. To see the format
	// and content of this field for different models, refer to Inference parameters (https://docs.aws.amazon.com/bedrock/latest/userguide/model-parameters.html)
	// .
	//
	// This member is required.
	Body []byte

	// Id of the model to invoke using the streaming request.
	//
	// This member is required.
	ModelId *string

	// The desired MIME type of the inference body in the response. The default value
	// is application/json .
	Accept *string

	// The MIME type of the input data in the request. The default value is
	// application/json .
	ContentType *string

	noSmithyDocumentSerde
}

type InvokeModelWithResponseStreamOutput struct {

	// The MIME type of the inference result.
	//
	// This member is required.
	ContentType *string

	eventStream *InvokeModelWithResponseStreamEventStream

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

// GetStream returns the type to interact with the event stream.
func (o *InvokeModelWithResponseStreamOutput) GetStream() *InvokeModelWithResponseStreamEventStream {
	return o.eventStream
}

func (c *Client) addOperationInvokeModelWithResponseStreamMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpInvokeModelWithResponseStream{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpInvokeModelWithResponseStream{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
		return err
	}
	if err = addEventStreamInvokeModelWithResponseStreamMiddleware(stack, options); err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack, options); err != nil {
		return err
	}
	if err = addInvokeModelWithResponseStreamResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpInvokeModelWithResponseStreamValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opInvokeModelWithResponseStream(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecursionDetection(stack); err != nil {
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
	if err = addendpointDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opInvokeModelWithResponseStream(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "bedrock",
		OperationName: "InvokeModelWithResponseStream",
	}
}

// InvokeModelWithResponseStreamEventStream provides the event stream handling for the InvokeModelWithResponseStream operation.
//
// For testing and mocking the event stream this type should be initialized via
// the NewInvokeModelWithResponseStreamEventStream constructor function. Using the functional options
// to pass in nested mock behavior.
type InvokeModelWithResponseStreamEventStream struct {
	// ResponseStreamReader is the EventStream reader for the ResponseStream events.
	// This value is automatically set by the SDK when the API call is made Use this
	// member when unit testing your code with the SDK to mock out the EventStream
	// Reader.
	//
	// Must not be nil.
	Reader ResponseStreamReader

	done      chan struct{}
	closeOnce sync.Once
	err       *smithysync.OnceErr
}

// NewInvokeModelWithResponseStreamEventStream initializes an InvokeModelWithResponseStreamEventStream.
// This function should only be used for testing and mocking the InvokeModelWithResponseStreamEventStream
// stream within your application.
//
// The Reader member must be set before reading events from the stream.
func NewInvokeModelWithResponseStreamEventStream(optFns ...func(*InvokeModelWithResponseStreamEventStream)) *InvokeModelWithResponseStreamEventStream {
	es := &InvokeModelWithResponseStreamEventStream{
		done: make(chan struct{}),
		err:  smithysync.NewOnceErr(),
	}
	for _, fn := range optFns {
		fn(es)
	}
	return es
}

// Events returns a channel to read events from.
func (es *InvokeModelWithResponseStreamEventStream) Events() <-chan types.ResponseStream {
	return es.Reader.Events()
}

// Close closes the stream. This will also cause the stream to be closed.
// Close must be called when done using the stream API. Not calling Close
// may result in resource leaks.
//
// Will close the underlying EventStream writer and reader, and no more events can be
// sent or received.
func (es *InvokeModelWithResponseStreamEventStream) Close() error {
	es.closeOnce.Do(es.safeClose)
	return es.Err()
}

func (es *InvokeModelWithResponseStreamEventStream) safeClose() {
	close(es.done)

	es.Reader.Close()
}

// Err returns any error that occurred while reading or writing EventStream Events
// from the service API's response. Returns nil if there were no errors.
func (es *InvokeModelWithResponseStreamEventStream) Err() error {
	if err := es.err.Err(); err != nil {
		return err
	}

	if err := es.Reader.Err(); err != nil {
		return err
	}

	return nil
}

func (es *InvokeModelWithResponseStreamEventStream) waitStreamClose() {
	type errorSet interface {
		ErrorSet() <-chan struct{}
	}

	var outputErrCh <-chan struct{}
	if v, ok := es.Reader.(errorSet); ok {
		outputErrCh = v.ErrorSet()
	}
	var outputClosedCh <-chan struct{}
	if v, ok := es.Reader.(interface{ Closed() <-chan struct{} }); ok {
		outputClosedCh = v.Closed()
	}

	select {
	case <-es.done:
	case <-outputErrCh:
		es.err.SetError(es.Reader.Err())
		es.Close()

	case <-outputClosedCh:
		if err := es.Reader.Err(); err != nil {
			es.err.SetError(es.Reader.Err())
		}
		es.Close()

	}
}

type opInvokeModelWithResponseStreamResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver  builtInParameterResolver
}

func (*opInvokeModelWithResponseStreamResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opInvokeModelWithResponseStreamResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	if awsmiddleware.GetRequiresLegacyEndpoints(ctx) {
		return next.HandleSerialize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	if m.EndpointResolver == nil {
		return out, metadata, fmt.Errorf("expected endpoint resolver to not be nil")
	}

	params := EndpointParameters{}

	m.BuiltInResolver.ResolveBuiltIns(&params)

	var resolvedEndpoint smithyendpoints.Endpoint
	resolvedEndpoint, err = m.EndpointResolver.ResolveEndpoint(ctx, params)
	if err != nil {
		return out, metadata, fmt.Errorf("failed to resolve service endpoint, %w", err)
	}

	req.URL = &resolvedEndpoint.URI

	for k := range resolvedEndpoint.Headers {
		req.Header.Set(
			k,
			resolvedEndpoint.Headers.Get(k),
		)
	}

	authSchemes, err := internalauth.GetAuthenticationSchemes(&resolvedEndpoint.Properties)
	if err != nil {
		var nfe *internalauth.NoAuthenticationSchemesFoundError
		if errors.As(err, &nfe) {
			// if no auth scheme is found, default to sigv4
			signingName := "bedrock"
			signingRegion := m.BuiltInResolver.(*builtInResolver).Region
			ctx = awsmiddleware.SetSigningName(ctx, signingName)
			ctx = awsmiddleware.SetSigningRegion(ctx, signingRegion)

		}
		var ue *internalauth.UnSupportedAuthenticationSchemeSpecifiedError
		if errors.As(err, &ue) {
			return out, metadata, fmt.Errorf(
				"This operation requests signer version(s) %v but the client only supports %v",
				ue.UnsupportedSchemes,
				internalauth.SupportedSchemes,
			)
		}
	}

	for _, authScheme := range authSchemes {
		switch authScheme.(type) {
		case *internalauth.AuthenticationSchemeV4:
			v4Scheme, _ := authScheme.(*internalauth.AuthenticationSchemeV4)
			var signingName, signingRegion string
			if v4Scheme.SigningName == nil {
				signingName = "bedrock"
			} else {
				signingName = *v4Scheme.SigningName
			}
			if v4Scheme.SigningRegion == nil {
				signingRegion = m.BuiltInResolver.(*builtInResolver).Region
			} else {
				signingRegion = *v4Scheme.SigningRegion
			}
			if v4Scheme.DisableDoubleEncoding != nil {
				// The signer sets an equivalent value at client initialization time.
				// Setting this context value will cause the signer to extract it
				// and override the value set at client initialization time.
				ctx = internalauth.SetDisableDoubleEncoding(ctx, *v4Scheme.DisableDoubleEncoding)
			}
			ctx = awsmiddleware.SetSigningName(ctx, signingName)
			ctx = awsmiddleware.SetSigningRegion(ctx, signingRegion)
			break
		case *internalauth.AuthenticationSchemeV4A:
			v4aScheme, _ := authScheme.(*internalauth.AuthenticationSchemeV4A)
			if v4aScheme.SigningName == nil {
				v4aScheme.SigningName = aws.String("bedrock")
			}
			if v4aScheme.DisableDoubleEncoding != nil {
				// The signer sets an equivalent value at client initialization time.
				// Setting this context value will cause the signer to extract it
				// and override the value set at client initialization time.
				ctx = internalauth.SetDisableDoubleEncoding(ctx, *v4aScheme.DisableDoubleEncoding)
			}
			ctx = awsmiddleware.SetSigningName(ctx, *v4aScheme.SigningName)
			ctx = awsmiddleware.SetSigningRegion(ctx, v4aScheme.SigningRegionSet[0])
			break
		case *internalauth.AuthenticationSchemeNone:
			break
		}
	}

	return next.HandleSerialize(ctx, in)
}

func addInvokeModelWithResponseStreamResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
	return stack.Serialize.Insert(&opInvokeModelWithResponseStreamResolveEndpointMiddleware{
		EndpointResolver: options.EndpointResolverV2,
		BuiltInResolver: &builtInResolver{
			Region:       options.Region,
			UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
			UseFIPS:      options.EndpointOptions.UseFIPSEndpoint,
			Endpoint:     options.BaseEndpoint,
		},
	}, "ResolveEndpoint", middleware.After)
}
