/*
ChannelEngine Merchant API

ChannelEngine API for merchants

API version: 2.14.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package merchant

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"reflect"
)


// ProductBundleAPIService ProductBundleAPI service
type ProductBundleAPIService service

type ProductBundleAPIProductBundleGetByFilterRequest struct {
	ctx context.Context
	ApiService *ProductBundleAPIService
	search *string
	eanList *[]string
	merchantProductNoList *[]string
	page *int32
}

// Search product(s) by Name, MerchantProductNo, Ean or Brand&lt;br /&gt;This search is applied to the result after applying the other filters.
func (r ProductBundleAPIProductBundleGetByFilterRequest) Search(search string) ProductBundleAPIProductBundleGetByFilterRequest {
	r.search = &search
	return r
}

// Search products by submitting a list of EAN&#39;s.
func (r ProductBundleAPIProductBundleGetByFilterRequest) EanList(eanList []string) ProductBundleAPIProductBundleGetByFilterRequest {
	r.eanList = &eanList
	return r
}

// Search products by submitting a list of MerchantProductNo&#39;s.
func (r ProductBundleAPIProductBundleGetByFilterRequest) MerchantProductNoList(merchantProductNoList []string) ProductBundleAPIProductBundleGetByFilterRequest {
	r.merchantProductNoList = &merchantProductNoList
	return r
}

// The page to filter on. Starts at 1.
func (r ProductBundleAPIProductBundleGetByFilterRequest) Page(page int32) ProductBundleAPIProductBundleGetByFilterRequest {
	r.page = &page
	return r
}

func (r ProductBundleAPIProductBundleGetByFilterRequest) Execute() (*CollectionOfMerchantProductBundleResponse, *http.Response, error) {
	return r.ApiService.ProductBundleGetByFilterExecute(r)
}

/*
ProductBundleGetByFilter Gets product bundles

Gets the product bundle details.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ProductBundleAPIProductBundleGetByFilterRequest
*/
func (a *ProductBundleAPIService) ProductBundleGetByFilter(ctx context.Context) ProductBundleAPIProductBundleGetByFilterRequest {
	return ProductBundleAPIProductBundleGetByFilterRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return CollectionOfMerchantProductBundleResponse
func (a *ProductBundleAPIService) ProductBundleGetByFilterExecute(r ProductBundleAPIProductBundleGetByFilterRequest) (*CollectionOfMerchantProductBundleResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CollectionOfMerchantProductBundleResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProductBundleAPIService.ProductBundleGetByFilter")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/productbundles"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.search != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "search", r.search, "")
	}
	if r.eanList != nil {
		t := *r.eanList
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "eanList", s.Index(i).Interface(), "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "eanList", t, "multi")
		}
	}
	if r.merchantProductNoList != nil {
		t := *r.merchantProductNoList
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "merchantProductNoList", s.Index(i).Interface(), "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "merchantProductNoList", t, "multi")
		}
	}
	if r.page != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page, "")
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
