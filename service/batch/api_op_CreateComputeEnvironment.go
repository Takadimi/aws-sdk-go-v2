// Code generated by smithy-go-codegen DO NOT EDIT.

package batch

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/batch/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates an AWS Batch compute environment. You can create MANAGED or UNMANAGED
// compute environments. MANAGED compute environments can use Amazon EC2 or AWS
// Fargate resources. UNMANAGED compute environments can only use EC2 resources. In
// a managed compute environment, AWS Batch manages the capacity and instance types
// of the compute resources within the environment. This is based on the compute
// resource specification that you define or the launch template
// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-launch-templates.html)
// that you specify when you create the compute environment. You can choose either
// to use EC2 On-Demand Instances and EC2 Spot Instances, or to use Fargate and
// Fargate Spot capacity in your managed compute environment. You can optionally
// set a maximum price so that Spot Instances only launch when the Spot Instance
// price is less than a specified percentage of the On-Demand price. Multi-node
// parallel jobs are not supported on Spot Instances. In an unmanaged compute
// environment, you can manage your own EC2 compute resources and have a lot of
// flexibility with how you configure your compute resources. For example, you can
// use custom AMI. However, you need to verify that your AMI meets the Amazon ECS
// container instance AMI specification. For more information, see container
// instance AMIs
// (https://docs.aws.amazon.com/AmazonECS/latest/developerguide/container_instance_AMIs.html)
// in the Amazon Elastic Container Service Developer Guide. After you have created
// your unmanaged compute environment, you can use the DescribeComputeEnvironments
// operation to find the Amazon ECS cluster that's associated with it. Then,
// manually launch your container instances into that Amazon ECS cluster. For more
// information, see Launching an Amazon ECS container instance
// (https://docs.aws.amazon.com/AmazonECS/latest/developerguide/launch_container_instance.html)
// in the Amazon Elastic Container Service Developer Guide. AWS Batch doesn't
// upgrade the AMIs in a compute environment after it's created. For example, it
// doesn't update the AMIs when a newer version of the Amazon ECS-optimized AMI is
// available. Therefore, you're responsible for the management of the guest
// operating system (including updates and security patches) and any additional
// application software or utilities that you install on the compute resources. To
// use a new AMI for your AWS Batch jobs, complete these steps:
//
// * Create a new
// compute environment with the new AMI.
//
// * Add the compute environment to an
// existing job queue.
//
// * Remove the earlier compute environment from your job
// queue.
//
// * Delete the earlier compute environment.
func (c *Client) CreateComputeEnvironment(ctx context.Context, params *CreateComputeEnvironmentInput, optFns ...func(*Options)) (*CreateComputeEnvironmentOutput, error) {
	if params == nil {
		params = &CreateComputeEnvironmentInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateComputeEnvironment", params, optFns, addOperationCreateComputeEnvironmentMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateComputeEnvironmentOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Contains the parameters for CreateComputeEnvironment.
type CreateComputeEnvironmentInput struct {

	// The name for your compute environment. Up to 128 letters (uppercase and
	// lowercase), numbers, hyphens, and underscores are allowed.
	//
	// This member is required.
	ComputeEnvironmentName *string

	// The full Amazon Resource Name (ARN) of the IAM role that allows AWS Batch to
	// make calls to other AWS services on your behalf. For more information, see AWS
	// Batch service IAM role
	// (https://docs.aws.amazon.com/batch/latest/userguide/service_IAM_role.html) in
	// the AWS Batch User Guide. If your specified role has a path other than /, then
	// you must either specify the full role ARN (this is recommended) or prefix the
	// role name with the path. Depending on how you created your AWS Batch service
	// role, its ARN might contain the service-role path prefix. When you only specify
	// the name of the service role, AWS Batch assumes that your ARN doesn't use the
	// service-role path prefix. Because of this, we recommend that you specify the
	// full ARN of your service role when you create compute environments.
	//
	// This member is required.
	ServiceRole *string

	// The type of the compute environment: MANAGED or UNMANAGED. For more information,
	// see Compute Environments
	// (https://docs.aws.amazon.com/batch/latest/userguide/compute_environments.html)
	// in the AWS Batch User Guide.
	//
	// This member is required.
	Type types.CEType

	// Details about the compute resources managed by the compute environment. This
	// parameter is required for managed compute environments. For more information,
	// see Compute Environments
	// (https://docs.aws.amazon.com/batch/latest/userguide/compute_environments.html)
	// in the AWS Batch User Guide.
	ComputeResources *types.ComputeResource

	// The state of the compute environment. If the state is ENABLED, then the compute
	// environment accepts jobs from a queue and can scale out automatically based on
	// queues. If the state is ENABLED, then the AWS Batch scheduler can attempt to
	// place jobs from an associated job queue on the compute resources within the
	// environment. If the compute environment is managed, then it can scale its
	// instances out or in automatically, based on the job queue demand. If the state
	// is DISABLED, then the AWS Batch scheduler doesn't attempt to place jobs within
	// the environment. Jobs in a STARTING or RUNNING state continue to progress
	// normally. Managed compute environments in the DISABLED state don't scale out.
	// However, they scale in to minvCpus value after instances become idle.
	State types.CEState

	// The tags that you apply to the compute environment to help you categorize and
	// organize your resources. Each tag consists of a key and an optional value. For
	// more information, see Tagging AWS Resources
	// (https://docs.aws.amazon.com/general/latest/gr/aws_tagging.html) in AWS General
	// Reference. These tags can be updated or removed using the TagResource
	// (https://docs.aws.amazon.com/batch/latest/APIReference/API_TagResource.html) and
	// UntagResource
	// (https://docs.aws.amazon.com/batch/latest/APIReference/API_UntagResource.html)
	// API operations. These tags don't propagate to the underlying compute resources.
	Tags map[string]string
}

type CreateComputeEnvironmentOutput struct {

	// The Amazon Resource Name (ARN) of the compute environment.
	ComputeEnvironmentArn *string

	// The name of the compute environment. Up to 128 letters (uppercase and
	// lowercase), numbers, hyphens, and underscores are allowed.
	ComputeEnvironmentName *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateComputeEnvironmentMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateComputeEnvironment{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateComputeEnvironment{}, middleware.After)
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
	if err = addOpCreateComputeEnvironmentValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateComputeEnvironment(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateComputeEnvironment(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "batch",
		OperationName: "CreateComputeEnvironment",
	}
}
