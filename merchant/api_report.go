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
	"strings"
	"os"
)


// ReportAPIService ReportAPI service
type ReportAPIService service

type ReportAPIReportCreateSettlementsReportRequest struct {
	ctx context.Context
	ApiService *ReportAPIService
	merchantCreateSettlementsReportRequest *MerchantCreateSettlementsReportRequest
}

// To provide settlementIds and type of report SUMMARY or DETAILED.
func (r ReportAPIReportCreateSettlementsReportRequest) MerchantCreateSettlementsReportRequest(merchantCreateSettlementsReportRequest MerchantCreateSettlementsReportRequest) ReportAPIReportCreateSettlementsReportRequest {
	r.merchantCreateSettlementsReportRequest = &merchantCreateSettlementsReportRequest
	return r
}

func (r ReportAPIReportCreateSettlementsReportRequest) Execute() (*MerchantCreateReportResponse, *http.Response, error) {
	return r.ApiService.ReportCreateSettlementsReportExecute(r)
}

/*
ReportCreateSettlementsReport Creates a settlement report

Creates a settlement report based on the **Settlement ID** provided. There are 3 types of reports:<br />**SUMMARY** - a high level financial overview.<br />**DETAILED** - a detailed report containing all transactions.<br />**CUSTOM_JSON** - a report grouped by orders, you can name the csv columns with a JSON file. This JSON file can be defined<br />in the settlement export plugin.<br /> <br />All the settlements are automatically acknowledged if that was not already the case.<br />**NB:** depending on the number of transactions within the settlement, it can take a few minutes for the report to be generated.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ReportAPIReportCreateSettlementsReportRequest
*/
func (a *ReportAPIService) ReportCreateSettlementsReport(ctx context.Context) ReportAPIReportCreateSettlementsReportRequest {
	return ReportAPIReportCreateSettlementsReportRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return MerchantCreateReportResponse
func (a *ReportAPIService) ReportCreateSettlementsReportExecute(r ReportAPIReportCreateSettlementsReportRequest) (*MerchantCreateReportResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *MerchantCreateReportResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ReportAPIService.ReportCreateSettlementsReport")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/reports/settlements"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.merchantCreateSettlementsReportRequest == nil {
		return localVarReturnValue, nil, reportError("merchantCreateSettlementsReportRequest is required and must be specified")
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
	localVarPostBody = r.merchantCreateSettlementsReportRequest
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
		if localVarHTTPResponse.StatusCode == 404 {
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

type ReportAPIReportGetReportRequest struct {
	ctx context.Context
	ApiService *ReportAPIService
	reportId string
}

func (r ReportAPIReportGetReportRequest) Execute() (**os.File, *http.Response, error) {
	return r.ApiService.ReportGetReportExecute(r)
}

/*
ReportGetReport Gets a settlement report

Gets a settlement report based on the **Report ID** provided. The generated report is a CSV file with a semicolon (;) as a delimiter.<br />If a field has a comma (,) then it is enclosed in quotes ("").

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param reportId 
 @return ReportAPIReportGetReportRequest
*/
func (a *ReportAPIService) ReportGetReport(ctx context.Context, reportId string) ReportAPIReportGetReportRequest {
	return ReportAPIReportGetReportRequest{
		ApiService: a,
		ctx: ctx,
		reportId: reportId,
	}
}

// Execute executes the request
//  @return *os.File
func (a *ReportAPIService) ReportGetReportExecute(r ReportAPIReportGetReportRequest) (**os.File, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  **os.File
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ReportAPIService.ReportGetReport")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/reports/{reportId}"
	localVarPath = strings.Replace(localVarPath, "{"+"reportId"+"}", url.PathEscape(parameterValueToString(r.reportId, "reportId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"text/csv", "application/json"}

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
		if localVarHTTPResponse.StatusCode == 404 {
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

type ReportAPIReportGetStatusRequest struct {
	ctx context.Context
	ApiService *ReportAPIService
	reportId string
}

func (r ReportAPIReportGetStatusRequest) Execute() (*MerchantGetReportStatusResponse, *http.Response, error) {
	return r.ApiService.ReportGetStatusExecute(r)
}

/*
ReportGetStatus Gets the status of a settlement report

Returns a report status based on the **Report ID** provided. There are four statuses:<br />**IN_PROGRESS** - the report is still being created.<br />**DONE** - the report has been created.<br />**FAILED** - the report creation failed.<br />**NOT_FOUND** - the Report ID was not found.<br /> <br />**NB:** if the status is **DONE**, the response contains a URL with a download path.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param reportId 
 @return ReportAPIReportGetStatusRequest
*/
func (a *ReportAPIService) ReportGetStatus(ctx context.Context, reportId string) ReportAPIReportGetStatusRequest {
	return ReportAPIReportGetStatusRequest{
		ApiService: a,
		ctx: ctx,
		reportId: reportId,
	}
}

// Execute executes the request
//  @return MerchantGetReportStatusResponse
func (a *ReportAPIService) ReportGetStatusExecute(r ReportAPIReportGetStatusRequest) (*MerchantGetReportStatusResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *MerchantGetReportStatusResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ReportAPIService.ReportGetStatus")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/reports/{reportId}/status"
	localVarPath = strings.Replace(localVarPath, "{"+"reportId"+"}", url.PathEscape(parameterValueToString(r.reportId, "reportId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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
