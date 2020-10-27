// Code generated by smithy-go-codegen DO NOT EDIT.

package personalizeevents

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/personalizeevents/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Adds one or more users to a Users dataset. For more information see
// importing-users.
func (c *Client) PutUsers(ctx context.Context, params *PutUsersInput, optFns ...func(*Options)) (*PutUsersOutput, error) {
	if params == nil {
		params = &PutUsersInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutUsers", params, optFns, addOperationPutUsersMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutUsersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutUsersInput struct {

	// The Amazon Resource Number (ARN) of the Users dataset you are adding the user or
	// users to.
	//
	// This member is required.
	DatasetArn *string

	// A list of user data.
	//
	// This member is required.
	Users []*types.User
}

type PutUsersOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationPutUsersMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpPutUsers{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpPutUsers{}, middleware.After)
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
	addOpPutUsersValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opPutUsers(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opPutUsers(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "personalize",
		OperationName: "PutUsers",
	}
}