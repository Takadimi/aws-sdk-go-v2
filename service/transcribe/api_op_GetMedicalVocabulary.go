// Code generated by smithy-go-codegen DO NOT EDIT.

package transcribe

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/transcribe/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Retrieves information about a medical vocabulary.
func (c *Client) GetMedicalVocabulary(ctx context.Context, params *GetMedicalVocabularyInput, optFns ...func(*Options)) (*GetMedicalVocabularyOutput, error) {
	if params == nil {
		params = &GetMedicalVocabularyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetMedicalVocabulary", params, optFns, c.addOperationGetMedicalVocabularyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetMedicalVocabularyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetMedicalVocabularyInput struct {

	// The name of the vocabulary that you want information about. The value is case
	// sensitive.
	//
	// This member is required.
	VocabularyName *string

	noSmithyDocumentSerde
}

type GetMedicalVocabularyOutput struct {

	// The location in Amazon S3 where the vocabulary is stored. Use this URI to get
	// the contents of the vocabulary. You can download your vocabulary from the URI
	// for a limited time.
	DownloadUri *string

	// If the VocabularyState is FAILED, this field contains information about why the
	// job failed.
	FailureReason *string

	// The valid language code for your vocabulary entries.
	LanguageCode types.LanguageCode

	// The date and time that the vocabulary was last modified with a text file
	// different from the one that was previously used.
	LastModifiedTime *time.Time

	// The name of the vocabulary returned by Amazon Transcribe Medical.
	VocabularyName *string

	// The processing state of the vocabulary. If the VocabularyState is READY then you
	// can use it in the StartMedicalTranscriptionJob operation.
	VocabularyState types.VocabularyState

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetMedicalVocabularyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetMedicalVocabulary{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetMedicalVocabulary{}, middleware.After)
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
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
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
	if err = addOpGetMedicalVocabularyValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetMedicalVocabulary(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetMedicalVocabulary(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "transcribe",
		OperationName: "GetMedicalVocabulary",
	}
}
