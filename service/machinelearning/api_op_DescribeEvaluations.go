// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package machinelearning

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DescribeEvaluationsInput struct {
	_ struct{} `type:"structure"`

	// The equal to operator. The Evaluation results will have FilterVariable values
	// that exactly match the value specified with EQ.
	EQ *string `type:"string"`

	// Use one of the following variable to filter a list of Evaluation objects:
	//
	//    * CreatedAt - Sets the search criteria to the Evaluation creation date.
	//
	//    * Status - Sets the search criteria to the Evaluation status.
	//
	//    * Name - Sets the search criteria to the contents of Evaluation Name.
	//
	//    * IAMUser - Sets the search criteria to the user account that invoked
	//    an Evaluation.
	//
	//    * MLModelId - Sets the search criteria to the MLModel that was evaluated.
	//
	//    * DataSourceId - Sets the search criteria to the DataSource used in Evaluation.
	//
	//    * DataUri - Sets the search criteria to the data file(s) used in Evaluation.
	//    The URL can identify either a file or an Amazon Simple Storage Solution
	//    (Amazon S3) bucket or directory.
	FilterVariable EvaluationFilterVariable `type:"string" enum:"true"`

	// The greater than or equal to operator. The Evaluation results will have FilterVariable
	// values that are greater than or equal to the value specified with GE.
	GE *string `type:"string"`

	// The greater than operator. The Evaluation results will have FilterVariable
	// values that are greater than the value specified with GT.
	GT *string `type:"string"`

	// The less than or equal to operator. The Evaluation results will have FilterVariable
	// values that are less than or equal to the value specified with LE.
	LE *string `type:"string"`

	// The less than operator. The Evaluation results will have FilterVariable values
	// that are less than the value specified with LT.
	LT *string `type:"string"`

	// The maximum number of Evaluation to include in the result.
	Limit *int64 `min:"1" type:"integer"`

	// The not equal to operator. The Evaluation results will have FilterVariable
	// values not equal to the value specified with NE.
	NE *string `type:"string"`

	// The ID of the page in the paginated results.
	NextToken *string `type:"string"`

	// A string that is found at the beginning of a variable, such as Name or Id.
	//
	// For example, an Evaluation could have the Name 2014-09-09-HolidayGiftMailer.
	// To search for this Evaluation, select Name for the FilterVariable and any
	// of the following strings for the Prefix:
	//
	//    * 2014-09
	//
	//    * 2014-09-09
	//
	//    * 2014-09-09-Holiday
	Prefix *string `type:"string"`

	// A two-value parameter that determines the sequence of the resulting list
	// of Evaluation.
	//
	//    * asc - Arranges the list in ascending order (A-Z, 0-9).
	//
	//    * dsc - Arranges the list in descending order (Z-A, 9-0).
	//
	// Results are sorted by FilterVariable.
	SortOrder SortOrder `type:"string" enum:"true"`
}

// String returns the string representation
func (s DescribeEvaluationsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeEvaluationsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeEvaluationsInput"}
	if s.Limit != nil && *s.Limit < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("Limit", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Represents the query results from a DescribeEvaluations operation. The content
// is essentially a list of Evaluation.
type DescribeEvaluationsOutput struct {
	_ struct{} `type:"structure"`

	// The ID of the next page in the paginated results that indicates at least
	// one more page follows.
	NextToken *string `type:"string"`

	// A list of Evaluation that meet the search criteria.
	Results []Evaluation `type:"list"`
}

// String returns the string representation
func (s DescribeEvaluationsOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeEvaluations = "DescribeEvaluations"

// DescribeEvaluationsRequest returns a request value for making API operation for
// Amazon Machine Learning.
//
// Returns a list of DescribeEvaluations that match the search criteria in the
// request.
//
//    // Example sending a request using DescribeEvaluationsRequest.
//    req := client.DescribeEvaluationsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) DescribeEvaluationsRequest(input *DescribeEvaluationsInput) DescribeEvaluationsRequest {
	op := &aws.Operation{
		Name:       opDescribeEvaluations,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "Limit",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &DescribeEvaluationsInput{}
	}

	req := c.newRequest(op, input, &DescribeEvaluationsOutput{})
	return DescribeEvaluationsRequest{Request: req, Input: input, Copy: c.DescribeEvaluationsRequest}
}

// DescribeEvaluationsRequest is the request type for the
// DescribeEvaluations API operation.
type DescribeEvaluationsRequest struct {
	*aws.Request
	Input *DescribeEvaluationsInput
	Copy  func(*DescribeEvaluationsInput) DescribeEvaluationsRequest
}

// Send marshals and sends the DescribeEvaluations API request.
func (r DescribeEvaluationsRequest) Send(ctx context.Context) (*DescribeEvaluationsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeEvaluationsResponse{
		DescribeEvaluationsOutput: r.Request.Data.(*DescribeEvaluationsOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewDescribeEvaluationsRequestPaginator returns a paginator for DescribeEvaluations.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.DescribeEvaluationsRequest(input)
//   p := machinelearning.NewDescribeEvaluationsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewDescribeEvaluationsPaginator(req DescribeEvaluationsRequest) DescribeEvaluationsPaginator {
	return DescribeEvaluationsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *DescribeEvaluationsInput
				if req.Input != nil {
					tmp := *req.Input
					inCpy = &tmp
				}

				newReq := req.Copy(inCpy)
				newReq.SetContext(ctx)
				return newReq.Request, nil
			},
		},
	}
}

// DescribeEvaluationsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type DescribeEvaluationsPaginator struct {
	aws.Pager
}

func (p *DescribeEvaluationsPaginator) CurrentPage() *DescribeEvaluationsOutput {
	return p.Pager.CurrentPage().(*DescribeEvaluationsOutput)
}

// DescribeEvaluationsResponse is the response type for the
// DescribeEvaluations API operation.
type DescribeEvaluationsResponse struct {
	*DescribeEvaluationsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeEvaluations request.
func (r *DescribeEvaluationsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}