// Code generated by smithy-go-codegen DO NOT EDIT.

package workmail

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/workmail/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Gets the effects of an organization's access control rules as they apply to a
// specified IPv4 address, access protocol action, or user ID.
func (c *Client) GetAccessControlEffect(ctx context.Context, params *GetAccessControlEffectInput, optFns ...func(*Options)) (*GetAccessControlEffectOutput, error) {
	if params == nil {
		params = &GetAccessControlEffectInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetAccessControlEffect", params, optFns, addOperationGetAccessControlEffectMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetAccessControlEffectOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetAccessControlEffectInput struct {

	// The access protocol action. Valid values include ActiveSync, AutoDiscover, EWS,
	// IMAP, SMTP, WindowsOutlook, and WebMail.
	//
	// This member is required.
	Action *string

	// The IPv4 address.
	//
	// This member is required.
	IpAddress *string

	// The identifier for the organization.
	//
	// This member is required.
	OrganizationId *string

	// The user ID.
	//
	// This member is required.
	UserId *string
}

type GetAccessControlEffectOutput struct {

	// The rule effect.
	Effect types.AccessControlRuleEffect

	// The rules that match the given parameters, resulting in an effect.
	MatchedRules []*string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetAccessControlEffectMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetAccessControlEffect{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetAccessControlEffect{}, middleware.After)
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
	addOpGetAccessControlEffectValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetAccessControlEffect(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opGetAccessControlEffect(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "workmail",
		OperationName: "GetAccessControlEffect",
	}
}