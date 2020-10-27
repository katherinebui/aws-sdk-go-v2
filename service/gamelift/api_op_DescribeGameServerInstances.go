// Code generated by smithy-go-codegen DO NOT EDIT.

package gamelift

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/gamelift/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// This operation is used with the Amazon GameLift FleetIQ solution and game server
// groups. Retrieves status information about the Amazon EC2 instances associated
// with a GameLift FleetIQ game server group. Use this operation to detect when
// instances are active or not available to host new game servers. If you are
// looking for instance configuration information, call DescribeGameServerGroup or
// access the corresponding Auto Scaling group properties. To request status for
// all instances in the game server group, provide a game server group ID only. To
// request status for specific instances, provide the game server group ID and one
// or more instance IDs. Use the pagination parameters to retrieve results in
// sequential segments. If successful, a collection of GameServerInstance objects
// is returned. This operation is not designed to be called with every game server
// claim request; this practice can cause you to exceed your API limit, which
// results in errors. Instead, as a best practice, cache the results and refresh
// your cache no more than once every 10 seconds. Learn more GameLift FleetIQ Guide
// (https://docs.aws.amazon.com/gamelift/latest/fleetiqguide/gsg-intro.html)
// Related operations
//
//     * CreateGameServerGroup
//
//     * ListGameServerGroups
//
//
// * DescribeGameServerGroup
//
//     * UpdateGameServerGroup
//
//     *
// DeleteGameServerGroup
//
//     * ResumeGameServerGroup
//
//     *
// SuspendGameServerGroup
//
//     * DescribeGameServerInstances
func (c *Client) DescribeGameServerInstances(ctx context.Context, params *DescribeGameServerInstancesInput, optFns ...func(*Options)) (*DescribeGameServerInstancesOutput, error) {
	if params == nil {
		params = &DescribeGameServerInstancesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeGameServerInstances", params, optFns, addOperationDescribeGameServerInstancesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeGameServerInstancesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeGameServerInstancesInput struct {

	// A unique identifier for the game server group. Use either the GameServerGroup
	// name or ARN value.
	//
	// This member is required.
	GameServerGroupName *string

	// The EC2 instance IDs that you want to retrieve status on. EC2 instance IDs use a
	// 17-character format, for example: i-1234567890abcdef0. To retrieve all instances
	// in the game server group, leave this parameter empty.
	InstanceIds []*string

	// The maximum number of results to return. Use this parameter with NextToken to
	// get results as a set of sequential segments.
	Limit *int32

	// A token that indicates the start of the next sequential segment of results. Use
	// the token returned with the previous call to this operation. To start at the
	// beginning of the result set, do not specify a value.
	NextToken *string
}

type DescribeGameServerInstancesOutput struct {

	// The collection of requested game server instances.
	GameServerInstances []*types.GameServerInstance

	// A token that indicates where to resume retrieving results on the next call to
	// this operation. If no token is returned, these results represent the end of the
	// list.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeGameServerInstancesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeGameServerInstances{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeGameServerInstances{}, middleware.After)
	if err != nil {
		return err
	}
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	addResolveEndpointMiddleware(stack, options)
	v4.AddComputePayloadSHA256Middleware(stack)
	addRetryMiddlewares(stack, options)
	addHTTPSignerV4Middleware(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	addClientUserAgent(stack)
	smithyhttp.AddErrorCloseResponseBodyMiddleware(stack)
	smithyhttp.AddCloseResponseBodyMiddleware(stack)
	addOpDescribeGameServerInstancesValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeGameServerInstances(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDescribeGameServerInstances(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "gamelift",
		OperationName: "DescribeGameServerInstances",
	}
}