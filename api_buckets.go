/*
OpenBuckets API

The OpenBuckets web-based tool is a powerful utility that allows users to quickly locate open buckets in cloud storage systems through a simple query. In addition, it provides a convenient way to search for various file types across these open buckets, making it an essential tool for security professionals, researchers, and anyone interested in discovering exposed data. This Postman collection aims to showcase the capabilities of OpenBuckets by providing a set of API requests that demonstrate how to leverage its features. By following this collection, you'll learn how to utilize OpenBuckets to identify open buckets and search for specific file types within them.

API version: 1.0.0
Contact: support@openbuckets.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openbuckets

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
)


// BucketsAPIService BucketsAPI service
type BucketsAPIService service

type ApiSearchBucketsRequest struct {
	ctx context.Context
	ApiService *BucketsAPIService
	keywords *string
	type_ *string
	exact *float32
	start *float32
	limit *float32
	order *string
	direction *string
}

// the search keywords to filter bucket names (e.g., \&quot;abg\&quot;)
func (r ApiSearchBucketsRequest) Keywords(keywords string) ApiSearchBucketsRequest {
	r.keywords = &keywords
	return r
}

// the type of bucket to filter (e.g., aws,dos,azure,gcp)
func (r ApiSearchBucketsRequest) Type_(type_ string) ApiSearchBucketsRequest {
	r.type_ = &type_
	return r
}

// whether to perform an exact match for the keywords (0 for false, 1 for true)
func (r ApiSearchBucketsRequest) Exact(exact float32) ApiSearchBucketsRequest {
	r.exact = &exact
	return r
}

// starting index for pagination
func (r ApiSearchBucketsRequest) Start(start float32) ApiSearchBucketsRequest {
	r.start = &start
	return r
}

// number of search results to return per page
func (r ApiSearchBucketsRequest) Limit(limit float32) ApiSearchBucketsRequest {
	r.limit = &limit
	return r
}

// the sorting field for the search results (e.g., \&quot;fileCount\&quot; for sorting by file count)
func (r ApiSearchBucketsRequest) Order(order string) ApiSearchBucketsRequest {
	r.order = &order
	return r
}

// the sorting direction for the search results (e.g., \&quot;asc\&quot; for ascending)
func (r ApiSearchBucketsRequest) Direction(direction string) ApiSearchBucketsRequest {
	r.direction = &direction
	return r
}

func (r ApiSearchBucketsRequest) Execute() (*BucketSearchResults, *http.Response, error) {
	return r.ApiService.SearchBucketsExecute(r)
}

/*
SearchBuckets Search Buckets

This request enables you to perform a targeted search for buckets within the OpenBuckets database using advanced filters. You can narrow down the search based on various criteria such as keywords, bucket type, exact match, sorting, and pagination.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiSearchBucketsRequest
*/
func (a *BucketsAPIService) SearchBuckets(ctx context.Context) ApiSearchBucketsRequest {
	return ApiSearchBucketsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return BucketSearchResults
func (a *BucketsAPIService) SearchBucketsExecute(r ApiSearchBucketsRequest) (*BucketSearchResults, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *BucketSearchResults
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BucketsAPIService.SearchBuckets")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/buckets"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.keywords != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "keywords", r.keywords, "")
	}
	if r.type_ != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "type", r.type_, "")
	}
	if r.exact != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "exact", r.exact, "")
	}
	if r.start != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "start", r.start, "")
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
	}
	if r.order != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "order", r.order, "")
	}
	if r.direction != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "direction", r.direction, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
