// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cloudfront

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type ListDistributionsByOriginRequestPolicyIdInput struct {
	_ struct{} `type:"structure"`

	// Use this field when paginating results to indicate where to begin in your
	// list of distribution IDs. The response includes distribution IDs in the list
	// that occur after the marker. To get the next page of the list, set this field’s
	// value to the value of NextMarker from the current page’s response.
	Marker *string `location:"querystring" locationName:"Marker" type:"string"`

	// The maximum number of distribution IDs that you want in the response.
	MaxItems *int64 `location:"querystring" locationName:"MaxItems" type:"integer"`

	// The ID of the origin request policy whose associated distribution IDs you
	// want to list.
	//
	// OriginRequestPolicyId is a required field
	OriginRequestPolicyId *string `location:"uri" locationName:"OriginRequestPolicyId" type:"string" required:"true"`
}

// String returns the string representation
func (s ListDistributionsByOriginRequestPolicyIdInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListDistributionsByOriginRequestPolicyIdInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListDistributionsByOriginRequestPolicyIdInput"}

	if s.OriginRequestPolicyId == nil {
		invalidParams.Add(aws.NewErrParamRequired("OriginRequestPolicyId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListDistributionsByOriginRequestPolicyIdInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.OriginRequestPolicyId != nil {
		v := *s.OriginRequestPolicyId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "OriginRequestPolicyId", protocol.StringValue(v), metadata)
	}
	if s.Marker != nil {
		v := *s.Marker

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "Marker", protocol.StringValue(v), metadata)
	}
	if s.MaxItems != nil {
		v := *s.MaxItems

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "MaxItems", protocol.Int64Value(v), metadata)
	}
	return nil
}

type ListDistributionsByOriginRequestPolicyIdOutput struct {
	_ struct{} `type:"structure" payload:"DistributionIdList"`

	// A list of distribution IDs.
	DistributionIdList *DistributionIdList `type:"structure"`
}

// String returns the string representation
func (s ListDistributionsByOriginRequestPolicyIdOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListDistributionsByOriginRequestPolicyIdOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.DistributionIdList != nil {
		v := s.DistributionIdList

		metadata := protocol.Metadata{}
		e.SetFields(protocol.PayloadTarget, "DistributionIdList", v, metadata)
	}
	return nil
}

const opListDistributionsByOriginRequestPolicyId = "ListDistributionsByOriginRequestPolicyId2020_05_31"

// ListDistributionsByOriginRequestPolicyIdRequest returns a request value for making API operation for
// Amazon CloudFront.
//
// Gets a list of distribution IDs for distributions that have a cache behavior
// that’s associated with the specified origin request policy.
//
// You can optionally specify the maximum number of items to receive in the
// response. If the total number of items in the list exceeds the maximum that
// you specify, or the default maximum, the response is paginated. To get the
// next page of items, send a subsequent request that specifies the NextMarker
// value from the current response as the Marker value in the subsequent request.
//
//    // Example sending a request using ListDistributionsByOriginRequestPolicyIdRequest.
//    req := client.ListDistributionsByOriginRequestPolicyIdRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cloudfront-2020-05-31/ListDistributionsByOriginRequestPolicyId
func (c *Client) ListDistributionsByOriginRequestPolicyIdRequest(input *ListDistributionsByOriginRequestPolicyIdInput) ListDistributionsByOriginRequestPolicyIdRequest {
	op := &aws.Operation{
		Name:       opListDistributionsByOriginRequestPolicyId,
		HTTPMethod: "GET",
		HTTPPath:   "/2020-05-31/distributionsByOriginRequestPolicyId/{OriginRequestPolicyId}",
	}

	if input == nil {
		input = &ListDistributionsByOriginRequestPolicyIdInput{}
	}

	req := c.newRequest(op, input, &ListDistributionsByOriginRequestPolicyIdOutput{})

	return ListDistributionsByOriginRequestPolicyIdRequest{Request: req, Input: input, Copy: c.ListDistributionsByOriginRequestPolicyIdRequest}
}

// ListDistributionsByOriginRequestPolicyIdRequest is the request type for the
// ListDistributionsByOriginRequestPolicyId API operation.
type ListDistributionsByOriginRequestPolicyIdRequest struct {
	*aws.Request
	Input *ListDistributionsByOriginRequestPolicyIdInput
	Copy  func(*ListDistributionsByOriginRequestPolicyIdInput) ListDistributionsByOriginRequestPolicyIdRequest
}

// Send marshals and sends the ListDistributionsByOriginRequestPolicyId API request.
func (r ListDistributionsByOriginRequestPolicyIdRequest) Send(ctx context.Context) (*ListDistributionsByOriginRequestPolicyIdResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListDistributionsByOriginRequestPolicyIdResponse{
		ListDistributionsByOriginRequestPolicyIdOutput: r.Request.Data.(*ListDistributionsByOriginRequestPolicyIdOutput),
		response: &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ListDistributionsByOriginRequestPolicyIdResponse is the response type for the
// ListDistributionsByOriginRequestPolicyId API operation.
type ListDistributionsByOriginRequestPolicyIdResponse struct {
	*ListDistributionsByOriginRequestPolicyIdOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListDistributionsByOriginRequestPolicyId request.
func (r *ListDistributionsByOriginRequestPolicyIdResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}