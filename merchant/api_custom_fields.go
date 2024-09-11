/*
ChannelEngine Merchant API

ChannelEngine API for merchants

API version: 2.17.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package merchant

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
)


// CustomFieldsAPIService CustomFieldsAPI service
type CustomFieldsAPIService service

type CustomFieldsAPICustomFieldsGetCustomFieldsRequest struct {
	ctx context.Context
	ApiService *CustomFieldsAPIService
	pageIndex *int32
	pageSize *int32
}

// A page index to get the items (starts from 0)
func (r CustomFieldsAPICustomFieldsGetCustomFieldsRequest) PageIndex(pageIndex int32) CustomFieldsAPICustomFieldsGetCustomFieldsRequest {
	r.pageIndex = &pageIndex
	return r
}

// Number of items to return (default 100)
func (r CustomFieldsAPICustomFieldsGetCustomFieldsRequest) PageSize(pageSize int32) CustomFieldsAPICustomFieldsGetCustomFieldsRequest {
	r.pageSize = &pageSize
	return r
}

func (r CustomFieldsAPICustomFieldsGetCustomFieldsRequest) Execute() (*CollectionOfCustomFieldResponse, *http.Response, error) {
	return r.ApiService.CustomFieldsGetCustomFieldsExecute(r)
}

/*
CustomFieldsGetCustomFields Gets custom fields

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return CustomFieldsAPICustomFieldsGetCustomFieldsRequest
*/
func (a *CustomFieldsAPIService) CustomFieldsGetCustomFields(ctx context.Context) CustomFieldsAPICustomFieldsGetCustomFieldsRequest {
	return CustomFieldsAPICustomFieldsGetCustomFieldsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return CollectionOfCustomFieldResponse
func (a *CustomFieldsAPIService) CustomFieldsGetCustomFieldsExecute(r CustomFieldsAPICustomFieldsGetCustomFieldsRequest) (*CollectionOfCustomFieldResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CollectionOfCustomFieldResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CustomFieldsAPIService.CustomFieldsGetCustomFields")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/custom-fields"

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
