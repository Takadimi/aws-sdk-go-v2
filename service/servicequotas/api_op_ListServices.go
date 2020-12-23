// Code generated by smithy-go-codegen DO NOT EDIT.

package servicequotas

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/servicequotas/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists the names and codes for the services integrated with Service Quotas.
func (c *Client) ListServices(ctx context.Context, params *ListServicesInput, optFns ...func(*Options)) (*ListServicesOutput, error) {
	if params == nil {
		params = &ListServicesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListServices", params, optFns, addOperationListServicesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListServicesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListServicesInput struct {

	// The maximum number of results to return with a single call. To retrieve the
	// remaining results, if any, make another call with the token returned from this
	// call.
	MaxResults *int32

	// The token for the next page of results.
	NextToken *string
}

type ListServicesOutput struct {

	// The token to use to retrieve the next page of results. This value is null when
	// there are no more results to return.
	NextToken *string

	// Information about the services.
	Services []types.ServiceInfo

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListServicesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListServices{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListServices{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListServices(options.Region), middleware.Before); err != nil {
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

// ListServicesAPIClient is a client that implements the ListServices operation.
type ListServicesAPIClient interface {
	ListServices(context.Context, *ListServicesInput, ...func(*Options)) (*ListServicesOutput, error)
}

var _ ListServicesAPIClient = (*Client)(nil)

// ListServicesPaginatorOptions is the paginator options for ListServices
type ListServicesPaginatorOptions struct {
	// The maximum number of results to return with a single call. To retrieve the
	// remaining results, if any, make another call with the token returned from this
	// call.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListServicesPaginator is a paginator for ListServices
type ListServicesPaginator struct {
	options   ListServicesPaginatorOptions
	client    ListServicesAPIClient
	params    *ListServicesInput
	nextToken *string
	firstPage bool
}

// NewListServicesPaginator returns a new ListServicesPaginator
func NewListServicesPaginator(client ListServicesAPIClient, params *ListServicesInput, optFns ...func(*ListServicesPaginatorOptions)) *ListServicesPaginator {
	options := ListServicesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	if params == nil {
		params = &ListServicesInput{}
	}

	return &ListServicesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListServicesPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next ListServices page.
func (p *ListServicesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListServicesOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxResults = limit

	result, err := p.client.ListServices(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListServices(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "servicequotas",
		OperationName: "ListServices",
	}
}
