// Code generated by smithy-go-codegen DO NOT EDIT.
package awsrestjson

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/awslabs/smithy-go/middleware"
	"github.com/awslabs/smithy-go/ptr"
	smithytesting "github.com/awslabs/smithy-go/testing"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

func TestClient_ConstantAndVariableQueryString_awsRestjson1Serialize(t *testing.T) {
	cases := map[string]struct {
		Params        *ConstantAndVariableQueryStringInput
		ExpectMethod  string
		ExpectURIPath string
		ExpectQuery   []smithytesting.QueryItem
		RequireQuery  []string
		ForbidQuery   []string
		ExpectHeader  http.Header
		RequireHeader []string
		ForbidHeader  []string
		BodyMediaType string
		BodyAssert    func(io.Reader) error
	}{
		// Mixes constant and variable query string parameters
		"RestJsonConstantAndVariableQueryStringMissingOneValue": {
			Params: &ConstantAndVariableQueryStringInput{
				Baz: ptr.String("bam"),
			},
			ExpectMethod:  "GET",
			ExpectURIPath: "/ConstantAndVariableQueryString",
			ExpectQuery: []smithytesting.QueryItem{
				{Key: "foo", Value: "bar"},
				{Key: "baz", Value: "bam"},
			},
			ForbidQuery: []string{
				"maybeSet",
			},
			BodyAssert: func(actual io.Reader) error {
				return smithytesting.CompareReaderEmpty(actual)
			},
		},
		// Mixes constant and variable query string parameters
		"RestJsonConstantAndVariableQueryStringAllValues": {
			Params: &ConstantAndVariableQueryStringInput{
				Baz:      ptr.String("bam"),
				MaybeSet: ptr.String("yes"),
			},
			ExpectMethod:  "GET",
			ExpectURIPath: "/ConstantAndVariableQueryString",
			ExpectQuery: []smithytesting.QueryItem{
				{Key: "foo", Value: "bar"},
				{Key: "baz", Value: "bam"},
				{Key: "maybeSet", Value: "yes"},
			},
			BodyAssert: func(actual io.Reader) error {
				return smithytesting.CompareReaderEmpty(actual)
			},
		},
	}
	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			var actualReq *http.Request
			client := New(Options{
				HTTPClient: smithyhttp.ClientDoFunc(func(r *http.Request) (*http.Response, error) {
					actualReq = r
					return &http.Response{
						StatusCode: 200,
						Header:     http.Header{},
						Body:       ioutil.NopCloser(strings.NewReader("")),
					}, nil
				}),
				APIOptions: []APIOptionFunc{
					func(s *middleware.Stack) error {
						s.Build.Clear()
						s.Finalize.Clear()
						return nil
					},
				},
				EndpointResolver: aws.EndpointResolverFunc(func(service, region string) (e aws.Endpoint, err error) {
					e.URL = "https://127.0.0.1"
					e.SigningRegion = "us-west-2"
					return e, err
				}),
				Region: "us-west-2"})
			result, err := client.ConstantAndVariableQueryString(context.Background(), c.Params)
			if err != nil {
				t.Fatalf("expect nil err, got %v", err)
			}
			if result == nil {
				t.Fatalf("expect not nil result")
			}
			if e, a := c.ExpectMethod, actualReq.Method; e != a {
				t.Errorf("expect %v method, got %v", e, a)
			}
			if e, a := c.ExpectURIPath, actualReq.URL.RawPath; e != a {
				t.Errorf("expect %v path, got %v", e, a)
			}
			queryItems := smithytesting.ParseRawQuery(actualReq.URL.RawQuery)
			smithytesting.AssertHasQuery(t, c.ExpectQuery, queryItems)
			smithytesting.AssertHasQueryKeys(t, c.RequireQuery, queryItems)
			smithytesting.AssertNotHaveQueryKeys(t, c.ForbidQuery, queryItems)
			smithytesting.AssertHasHeader(t, c.ExpectHeader, actualReq.Header)
			smithytesting.AssertHasHeaderKeys(t, c.RequireHeader, actualReq.Header)
			smithytesting.AssertNotHaveHeaderKeys(t, c.ForbidHeader, actualReq.Header)
			if actualReq.Body != nil {
				defer actualReq.Body.Close()
			}
			if c.BodyAssert != nil {
				if err := c.BodyAssert(actualReq.Body); err != nil {
					t.Errorf("expect body equal, got %v", err)
				}
			}
		})
	}
}
