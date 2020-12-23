// Code generated by smithy-go-codegen DO NOT EDIT.

package workmail

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Starts a mailbox export job to export MIME-format email messages and calendar
// items from the specified mailbox to the specified Amazon Simple Storage Service
// (Amazon S3) bucket. For more information, see Exporting mailbox content
// (https://docs.aws.amazon.com/workmail/latest/adminguide/mail-export.html) in the
// Amazon WorkMail Administrator Guide.
func (c *Client) StartMailboxExportJob(ctx context.Context, params *StartMailboxExportJobInput, optFns ...func(*Options)) (*StartMailboxExportJobOutput, error) {
	if params == nil {
		params = &StartMailboxExportJobInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StartMailboxExportJob", params, optFns, addOperationStartMailboxExportJobMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StartMailboxExportJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartMailboxExportJobInput struct {

	// The idempotency token for the client request.
	//
	// This member is required.
	ClientToken *string

	// The identifier of the user or resource associated with the mailbox.
	//
	// This member is required.
	EntityId *string

	// The Amazon Resource Name (ARN) of the symmetric AWS Key Management Service (AWS
	// KMS) key that encrypts the exported mailbox content.
	//
	// This member is required.
	KmsKeyArn *string

	// The identifier associated with the organization.
	//
	// This member is required.
	OrganizationId *string

	// The ARN of the AWS Identity and Access Management (IAM) role that grants write
	// permission to the S3 bucket.
	//
	// This member is required.
	RoleArn *string

	// The name of the S3 bucket.
	//
	// This member is required.
	S3BucketName *string

	// The S3 bucket prefix.
	//
	// This member is required.
	S3Prefix *string

	// The mailbox export job description.
	Description *string
}

type StartMailboxExportJobOutput struct {

	// The job ID.
	JobId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationStartMailboxExportJobMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpStartMailboxExportJob{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpStartMailboxExportJob{}, middleware.After)
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
	if err = addIdempotencyToken_opStartMailboxExportJobMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpStartMailboxExportJobValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStartMailboxExportJob(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpStartMailboxExportJob struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpStartMailboxExportJob) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpStartMailboxExportJob) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*StartMailboxExportJobInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *StartMailboxExportJobInput ")
	}

	if input.ClientToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opStartMailboxExportJobMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpStartMailboxExportJob{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opStartMailboxExportJob(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "workmail",
		OperationName: "StartMailboxExportJob",
	}
}
