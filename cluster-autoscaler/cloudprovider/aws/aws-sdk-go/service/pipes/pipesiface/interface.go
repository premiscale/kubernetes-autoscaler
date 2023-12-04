// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package pipesiface provides an interface to enable mocking the Amazon EventBridge Pipes service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package pipesiface

import (
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/aws/aws-sdk-go/aws"
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/aws/aws-sdk-go/aws/request"
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/aws/aws-sdk-go/service/pipes"
)

// PipesAPI provides an interface to enable mocking the
// pipes.Pipes service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//	// myFunc uses an SDK service client to make a request to
//	// Amazon EventBridge Pipes.
//	func myFunc(svc pipesiface.PipesAPI) bool {
//	    // Make svc.CreatePipe request
//	}
//
//	func main() {
//	    sess := session.New()
//	    svc := pipes.New(sess)
//
//	    myFunc(svc)
//	}
//
// In your _test.go file:
//
//	// Define a mock struct to be used in your unit tests of myFunc.
//	type mockPipesClient struct {
//	    pipesiface.PipesAPI
//	}
//	func (m *mockPipesClient) CreatePipe(input *pipes.CreatePipeInput) (*pipes.CreatePipeOutput, error) {
//	    // mock response/functionality
//	}
//
//	func TestMyFunc(t *testing.T) {
//	    // Setup Test
//	    mockSvc := &mockPipesClient{}
//
//	    myfunc(mockSvc)
//
//	    // Verify myFunc's functionality
//	}
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters. Its suggested to use the pattern above for testing, or using
// tooling to generate mocks to satisfy the interfaces.
type PipesAPI interface {
	CreatePipe(*pipes.CreatePipeInput) (*pipes.CreatePipeOutput, error)
	CreatePipeWithContext(aws.Context, *pipes.CreatePipeInput, ...request.Option) (*pipes.CreatePipeOutput, error)
	CreatePipeRequest(*pipes.CreatePipeInput) (*request.Request, *pipes.CreatePipeOutput)

	DeletePipe(*pipes.DeletePipeInput) (*pipes.DeletePipeOutput, error)
	DeletePipeWithContext(aws.Context, *pipes.DeletePipeInput, ...request.Option) (*pipes.DeletePipeOutput, error)
	DeletePipeRequest(*pipes.DeletePipeInput) (*request.Request, *pipes.DeletePipeOutput)

	DescribePipe(*pipes.DescribePipeInput) (*pipes.DescribePipeOutput, error)
	DescribePipeWithContext(aws.Context, *pipes.DescribePipeInput, ...request.Option) (*pipes.DescribePipeOutput, error)
	DescribePipeRequest(*pipes.DescribePipeInput) (*request.Request, *pipes.DescribePipeOutput)

	ListPipes(*pipes.ListPipesInput) (*pipes.ListPipesOutput, error)
	ListPipesWithContext(aws.Context, *pipes.ListPipesInput, ...request.Option) (*pipes.ListPipesOutput, error)
	ListPipesRequest(*pipes.ListPipesInput) (*request.Request, *pipes.ListPipesOutput)

	ListPipesPages(*pipes.ListPipesInput, func(*pipes.ListPipesOutput, bool) bool) error
	ListPipesPagesWithContext(aws.Context, *pipes.ListPipesInput, func(*pipes.ListPipesOutput, bool) bool, ...request.Option) error

	ListTagsForResource(*pipes.ListTagsForResourceInput) (*pipes.ListTagsForResourceOutput, error)
	ListTagsForResourceWithContext(aws.Context, *pipes.ListTagsForResourceInput, ...request.Option) (*pipes.ListTagsForResourceOutput, error)
	ListTagsForResourceRequest(*pipes.ListTagsForResourceInput) (*request.Request, *pipes.ListTagsForResourceOutput)

	StartPipe(*pipes.StartPipeInput) (*pipes.StartPipeOutput, error)
	StartPipeWithContext(aws.Context, *pipes.StartPipeInput, ...request.Option) (*pipes.StartPipeOutput, error)
	StartPipeRequest(*pipes.StartPipeInput) (*request.Request, *pipes.StartPipeOutput)

	StopPipe(*pipes.StopPipeInput) (*pipes.StopPipeOutput, error)
	StopPipeWithContext(aws.Context, *pipes.StopPipeInput, ...request.Option) (*pipes.StopPipeOutput, error)
	StopPipeRequest(*pipes.StopPipeInput) (*request.Request, *pipes.StopPipeOutput)

	TagResource(*pipes.TagResourceInput) (*pipes.TagResourceOutput, error)
	TagResourceWithContext(aws.Context, *pipes.TagResourceInput, ...request.Option) (*pipes.TagResourceOutput, error)
	TagResourceRequest(*pipes.TagResourceInput) (*request.Request, *pipes.TagResourceOutput)

	UntagResource(*pipes.UntagResourceInput) (*pipes.UntagResourceOutput, error)
	UntagResourceWithContext(aws.Context, *pipes.UntagResourceInput, ...request.Option) (*pipes.UntagResourceOutput, error)
	UntagResourceRequest(*pipes.UntagResourceInput) (*request.Request, *pipes.UntagResourceOutput)

	UpdatePipe(*pipes.UpdatePipeInput) (*pipes.UpdatePipeOutput, error)
	UpdatePipeWithContext(aws.Context, *pipes.UpdatePipeInput, ...request.Option) (*pipes.UpdatePipeOutput, error)
	UpdatePipeRequest(*pipes.UpdatePipeInput) (*request.Request, *pipes.UpdatePipeOutput)
}

var _ PipesAPI = (*pipes.Pipes)(nil)
