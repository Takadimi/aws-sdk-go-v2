// Code generated by smithy-go-codegen DO NOT EDIT.

package connect

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/connect/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// This API is in preview release for Amazon Connect and is subject to change.
// Provides information about the quick connects for the specified Amazon Connect
// instance.
func (c *Client) ListQuickConnects(ctx context.Context, params *ListQuickConnectsInput, optFns ...func(*Options)) (*ListQuickConnectsOutput, error) {
	if params == nil {
		params = &ListQuickConnectsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListQuickConnects", params, optFns, addOperationListQuickConnectsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListQuickConnectsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListQuickConnectsInput struct {

	// The identifier of the Amazon Connect instance.
	//
	// This member is required.
	InstanceId *string

	// The maximimum number of results to return per page.
	MaxResults int32

	// The token for the next set of results. Use the value returned in the previous
	// response in the next request to retrieve the next set of results.
	NextToken *string

	// The type of quick connect. In the Amazon Connect console, when you create a
	// quick connect, you are prompted to assign one of the following types: Agent
	// (USER), External (PHONE_NUMBER), or Queue (QUEUE).
	QuickConnectTypes []types.QuickConnectType
}

type ListQuickConnectsOutput struct {

	// If there are additional results, this is the token for the next set of results.
	NextToken *string

	// Information about the quick connects.
	QuickConnectSummaryList []types.QuickConnectSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListQuickConnectsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListQuickConnects{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListQuickConnects{}, middleware.After)
	if err != nil {
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
	if err = awsmiddleware.AddAttemptClockSkewMiddleware(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpListQuickConnectsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListQuickConnects(options.Region), middleware.Before); err != nil {
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
	return nil
}

// ListQuickConnectsAPIClient is a client that implements the ListQuickConnects
// operation.
type ListQuickConnectsAPIClient interface {
	ListQuickConnects(context.Context, *ListQuickConnectsInput, ...func(*Options)) (*ListQuickConnectsOutput, error)
}

var _ ListQuickConnectsAPIClient = (*Client)(nil)

// ListQuickConnectsPaginatorOptions is the paginator options for ListQuickConnects
type ListQuickConnectsPaginatorOptions struct {
	// The maximimum number of results to return per page.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListQuickConnectsPaginator is a paginator for ListQuickConnects
type ListQuickConnectsPaginator struct {
	options   ListQuickConnectsPaginatorOptions
	client    ListQuickConnectsAPIClient
	params    *ListQuickConnectsInput
	nextToken *string
	firstPage bool
}

// NewListQuickConnectsPaginator returns a new ListQuickConnectsPaginator
func NewListQuickConnectsPaginator(client ListQuickConnectsAPIClient, params *ListQuickConnectsInput, optFns ...func(*ListQuickConnectsPaginatorOptions)) *ListQuickConnectsPaginator {
	options := ListQuickConnectsPaginatorOptions{}
	if params.MaxResults != 0 {
		options.Limit = params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	if params == nil {
		params = &ListQuickConnectsInput{}
	}

	return &ListQuickConnectsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListQuickConnectsPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next ListQuickConnects page.
func (p *ListQuickConnectsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListQuickConnectsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.MaxResults = p.options.Limit

	result, err := p.client.ListQuickConnects(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextToken

	if p.options.StopOnDuplicateToken && prevToken != nil && p.nextToken != nil && *prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opListQuickConnects(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "connect",
		OperationName: "ListQuickConnects",
	}
}
