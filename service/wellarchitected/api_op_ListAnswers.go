// Code generated by smithy-go-codegen DO NOT EDIT.

package wellarchitected

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/wellarchitected/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// List of answers.
func (c *Client) ListAnswers(ctx context.Context, params *ListAnswersInput, optFns ...func(*Options)) (*ListAnswersOutput, error) {
	if params == nil {
		params = &ListAnswersInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListAnswers", params, optFns, addOperationListAnswersMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListAnswersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Input to list answers.
type ListAnswersInput struct {

	// The alias of the lens, for example, serverless. Each lens is identified by its
	// LensSummary$LensAlias.
	//
	// This member is required.
	LensAlias *string

	// The ID assigned to the workload. This ID is unique within an AWS Region.
	//
	// This member is required.
	WorkloadId *string

	// The maximum number of results to return for this request.
	MaxResults int32

	// The milestone number. A workload can have a maximum of 100 milestones.
	MilestoneNumber int32

	// The token to use to retrieve the next set of results.
	NextToken *string

	// The ID used to identify a pillar, for example, security. A pillar is identified
	// by its PillarReviewSummary$PillarId.
	PillarId *string
}

// Output of a list answers call.
type ListAnswersOutput struct {

	// List of answer summaries of lens review in a workload.
	AnswerSummaries []types.AnswerSummary

	// The alias of the lens, for example, serverless. Each lens is identified by its
	// LensSummary$LensAlias.
	LensAlias *string

	// The milestone number. A workload can have a maximum of 100 milestones.
	MilestoneNumber int32

	// The token to use to retrieve the next set of results.
	NextToken *string

	// The ID assigned to the workload. This ID is unique within an AWS Region.
	WorkloadId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListAnswersMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListAnswers{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListAnswers{}, middleware.After)
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
	if err = addOpListAnswersValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListAnswers(options.Region), middleware.Before); err != nil {
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

// ListAnswersAPIClient is a client that implements the ListAnswers operation.
type ListAnswersAPIClient interface {
	ListAnswers(context.Context, *ListAnswersInput, ...func(*Options)) (*ListAnswersOutput, error)
}

var _ ListAnswersAPIClient = (*Client)(nil)

// ListAnswersPaginatorOptions is the paginator options for ListAnswers
type ListAnswersPaginatorOptions struct {
	// The maximum number of results to return for this request.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListAnswersPaginator is a paginator for ListAnswers
type ListAnswersPaginator struct {
	options   ListAnswersPaginatorOptions
	client    ListAnswersAPIClient
	params    *ListAnswersInput
	nextToken *string
	firstPage bool
}

// NewListAnswersPaginator returns a new ListAnswersPaginator
func NewListAnswersPaginator(client ListAnswersAPIClient, params *ListAnswersInput, optFns ...func(*ListAnswersPaginatorOptions)) *ListAnswersPaginator {
	options := ListAnswersPaginatorOptions{}
	if params.MaxResults != 0 {
		options.Limit = params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	if params == nil {
		params = &ListAnswersInput{}
	}

	return &ListAnswersPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListAnswersPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next ListAnswers page.
func (p *ListAnswersPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListAnswersOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.MaxResults = p.options.Limit

	result, err := p.client.ListAnswers(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListAnswers(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "wellarchitected",
		OperationName: "ListAnswers",
	}
}
