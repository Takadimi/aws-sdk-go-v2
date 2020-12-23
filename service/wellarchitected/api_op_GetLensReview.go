// Code generated by smithy-go-codegen DO NOT EDIT.

package wellarchitected

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/wellarchitected/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Get lens review.
func (c *Client) GetLensReview(ctx context.Context, params *GetLensReviewInput, optFns ...func(*Options)) (*GetLensReviewOutput, error) {
	if params == nil {
		params = &GetLensReviewInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetLensReview", params, optFns, addOperationGetLensReviewMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetLensReviewOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Input to get lens review.
type GetLensReviewInput struct {

	// The alias of the lens, for example, serverless. Each lens is identified by its
	// LensSummary$LensAlias.
	//
	// This member is required.
	LensAlias *string

	// The ID assigned to the workload. This ID is unique within an AWS Region.
	//
	// This member is required.
	WorkloadId *string

	// The milestone number. A workload can have a maximum of 100 milestones.
	MilestoneNumber int32
}

// Output of a get lens review call.
type GetLensReviewOutput struct {

	// A lens review of a question.
	LensReview *types.LensReview

	// The milestone number. A workload can have a maximum of 100 milestones.
	MilestoneNumber int32

	// The ID assigned to the workload. This ID is unique within an AWS Region.
	WorkloadId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetLensReviewMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetLensReview{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetLensReview{}, middleware.After)
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
	if err = addOpGetLensReviewValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetLensReview(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetLensReview(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "wellarchitected",
		OperationName: "GetLensReview",
	}
}
