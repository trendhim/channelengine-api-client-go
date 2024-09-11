/*
ChannelEngine Channel API

ChannelEngine API for channels

API version: 2.17.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package channel

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)


// RefundsAPIService RefundsAPI service
type RefundsAPIService service

type RefundsAPIRefundGetRequest struct {
	ctx context.Context
	ApiService *RefundsAPIService
	identifier string
	type_ *RefundIdentifier
}

// Specify whether to search by ID, Merchant Refund No or Channel Refund No
func (r RefundsAPIRefundGetRequest) Type_(type_ RefundIdentifier) RefundsAPIRefundGetRequest {
	r.type_ = &type_
	return r
}

func (r RefundsAPIRefundGetRequest) Execute() (*SingleOfIRefund, *http.Response, error) {
	return r.ApiService.RefundGetExecute(r)
}

/*
RefundGet [CLOSED BETA] Get refund by identifier

Gets a single refund by the given identifier<br /> <br />Beware, this endpoint is part of a closed beta and is only available for closed beta participants.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param identifier The identifier to search for
 @return RefundsAPIRefundGetRequest
*/
func (a *RefundsAPIService) RefundGet(ctx context.Context, identifier string) RefundsAPIRefundGetRequest {
	return RefundsAPIRefundGetRequest{
		ApiService: a,
		ctx: ctx,
		identifier: identifier,
	}
}

// Execute executes the request
//  @return SingleOfIRefund
func (a *RefundsAPIService) RefundGetExecute(r RefundsAPIRefundGetRequest) (*SingleOfIRefund, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *SingleOfIRefund
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RefundsAPIService.RefundGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2.1/refunds/channel/{identifier}"
	localVarPath = strings.Replace(localVarPath, "{"+"identifier"+"}", url.PathEscape(parameterValueToString(r.identifier, "identifier")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.type_ != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "type", r.type_, "")
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
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["apiKey"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarQueryParams.Add("apikey", key)
			}
		}
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
		if localVarHTTPResponse.StatusCode == 400 {
			var v ApiResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
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

type RefundsAPIRefundGetByFilterRequest struct {
	ctx context.Context
	ApiService *RefundsAPIService
	pageIndex *int32
	pageSize *int32
	channelRefundsFilterRequest *ChannelRefundsFilterRequest
}

func (r RefundsAPIRefundGetByFilterRequest) PageIndex(pageIndex int32) RefundsAPIRefundGetByFilterRequest {
	r.pageIndex = &pageIndex
	return r
}

func (r RefundsAPIRefundGetByFilterRequest) PageSize(pageSize int32) RefundsAPIRefundGetByFilterRequest {
	r.pageSize = &pageSize
	return r
}

// The filter
func (r RefundsAPIRefundGetByFilterRequest) ChannelRefundsFilterRequest(channelRefundsFilterRequest ChannelRefundsFilterRequest) RefundsAPIRefundGetByFilterRequest {
	r.channelRefundsFilterRequest = &channelRefundsFilterRequest
	return r
}

func (r RefundsAPIRefundGetByFilterRequest) Execute() (*SingleOfIRefund, *http.Response, error) {
	return r.ApiService.RefundGetByFilterExecute(r)
}

/*
RefundGetByFilter [CLOSED BETA] Get refunds by filter

Gets multiple refunds by the given filter<br /> <br />Beware, this endpoint is part of a closed beta and is only available for closed beta participants.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return RefundsAPIRefundGetByFilterRequest
*/
func (a *RefundsAPIService) RefundGetByFilter(ctx context.Context) RefundsAPIRefundGetByFilterRequest {
	return RefundsAPIRefundGetByFilterRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return SingleOfIRefund
func (a *RefundsAPIService) RefundGetByFilterExecute(r RefundsAPIRefundGetByFilterRequest) (*SingleOfIRefund, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *SingleOfIRefund
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RefundsAPIService.RefundGetByFilter")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2.1/refunds/channel"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.pageIndex != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pageIndex", r.pageIndex, "")
	}
	if r.pageSize != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pageSize", r.pageSize, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json-patch+json", "application/json", "application/*+json"}

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
	// body params
	localVarPostBody = r.channelRefundsFilterRequest
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["apiKey"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarQueryParams.Add("apikey", key)
			}
		}
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
		if localVarHTTPResponse.StatusCode == 400 {
			var v ApiResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
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

type RefundsAPIRefundProcessRequest struct {
	ctx context.Context
	ApiService *RefundsAPIService
	singleChannelProcessRefundRequest *SingleChannelProcessRefundRequest
}

// The refund to acknowledge
func (r RefundsAPIRefundProcessRequest) SingleChannelProcessRefundRequest(singleChannelProcessRefundRequest SingleChannelProcessRefundRequest) RefundsAPIRefundProcessRequest {
	r.singleChannelProcessRefundRequest = &singleChannelProcessRefundRequest
	return r
}

func (r RefundsAPIRefundProcessRequest) Execute() (*ApiResponse, *http.Response, error) {
	return r.ApiService.RefundProcessExecute(r)
}

/*
RefundProcess [CLOSED BETA] Process a refund

Marks a refund as processed<br /> <br />Beware, this endpoint is part of a closed beta and is only available for closed beta participants.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return RefundsAPIRefundProcessRequest
*/
func (a *RefundsAPIService) RefundProcess(ctx context.Context) RefundsAPIRefundProcessRequest {
	return RefundsAPIRefundProcessRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ApiResponse
func (a *RefundsAPIService) RefundProcessExecute(r RefundsAPIRefundProcessRequest) (*ApiResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ApiResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RefundsAPIService.RefundProcess")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2.1/refunds/channel/process"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json-patch+json", "application/json", "application/*+json"}

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
	// body params
	localVarPostBody = r.singleChannelProcessRefundRequest
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["apiKey"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarQueryParams.Add("apikey", key)
			}
		}
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