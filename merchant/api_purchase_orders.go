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
	"time"
)


// PurchaseOrdersAPIService PurchaseOrdersAPI service
type PurchaseOrdersAPIService service

type PurchaseOrdersAPIAcknowledgeRequest struct {
	ctx context.Context
	ApiService *PurchaseOrdersAPIService
	singleMerchantAcknowledgePurchaseOrderLinesRequest *SingleMerchantAcknowledgePurchaseOrderLinesRequest
}

// Model for purchase order and lines data to be acknowledged with the channel.
func (r PurchaseOrdersAPIAcknowledgeRequest) SingleMerchantAcknowledgePurchaseOrderLinesRequest(singleMerchantAcknowledgePurchaseOrderLinesRequest SingleMerchantAcknowledgePurchaseOrderLinesRequest) PurchaseOrdersAPIAcknowledgeRequest {
	r.singleMerchantAcknowledgePurchaseOrderLinesRequest = &singleMerchantAcknowledgePurchaseOrderLinesRequest
	return r
}

func (r PurchaseOrdersAPIAcknowledgeRequest) Execute() (*ApiResponse, *http.Response, error) {
	return r.ApiService.AcknowledgeExecute(r)
}

/*
Acknowledge Acknowledges lines of a purchase order

Creates line acknowledgements (i.e., accepted, backordered, rejected) for a purchase order.<br />Request will be accepted and data persisted only if all validations passed.<br />Any validation messages and errors will be returned in a HTTP 4xx response.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return PurchaseOrdersAPIAcknowledgeRequest
*/
func (a *PurchaseOrdersAPIService) Acknowledge(ctx context.Context) PurchaseOrdersAPIAcknowledgeRequest {
	return PurchaseOrdersAPIAcknowledgeRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ApiResponse
func (a *PurchaseOrdersAPIService) AcknowledgeExecute(r PurchaseOrdersAPIAcknowledgeRequest) (*ApiResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ApiResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PurchaseOrdersAPIService.Acknowledge")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/purchase-orders/lines/acknowledge"

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
	localVarPostBody = r.singleMerchantAcknowledgePurchaseOrderLinesRequest
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
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 409 {
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

type PurchaseOrdersAPICreateRequest struct {
	ctx context.Context
	ApiService *PurchaseOrdersAPIService
	singleMerchantCreatePurchaseOrderShipmentRequest *SingleMerchantCreatePurchaseOrderShipmentRequest
}

func (r PurchaseOrdersAPICreateRequest) SingleMerchantCreatePurchaseOrderShipmentRequest(singleMerchantCreatePurchaseOrderShipmentRequest SingleMerchantCreatePurchaseOrderShipmentRequest) PurchaseOrdersAPICreateRequest {
	r.singleMerchantCreatePurchaseOrderShipmentRequest = &singleMerchantCreatePurchaseOrderShipmentRequest
	return r
}

func (r PurchaseOrdersAPICreateRequest) Execute() (*ApiResponse, *http.Response, error) {
	return r.ApiService.CreateExecute(r)
}

/*
Create Create a purchase order shipment.

One shipments can ship (parts of) multiple orders

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return PurchaseOrdersAPICreateRequest
*/
func (a *PurchaseOrdersAPIService) Create(ctx context.Context) PurchaseOrdersAPICreateRequest {
	return PurchaseOrdersAPICreateRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ApiResponse
func (a *PurchaseOrdersAPIService) CreateExecute(r PurchaseOrdersAPICreateRequest) (*ApiResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ApiResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PurchaseOrdersAPIService.Create")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/purchase-orders/shipments"

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
	localVarPostBody = r.singleMerchantCreatePurchaseOrderShipmentRequest
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

type PurchaseOrdersAPIGetByFilterRequest struct {
	ctx context.Context
	ApiService *PurchaseOrdersAPIService
	channelId *int32
	identifiersIdentifierType *PurchaseOrderShipmentIdentifierTypeValue
	identifiersModels *[]string
	shippedDateRangeFromDate *time.Time
	shippedDateRangeToDate *time.Time
	createDateRangeFromDate *time.Time
	createDateRangeToDate *time.Time
	updateDateRangeFromDate *time.Time
	updateDateRangeToDate *time.Time
	billOfLadingNumber *string
	carrierName *string
	pageIndex *int32
	pageSize *int32
}

// The identifier of the channel
func (r PurchaseOrdersAPIGetByFilterRequest) ChannelId(channelId int32) PurchaseOrdersAPIGetByFilterRequest {
	r.channelId = &channelId
	return r
}

// The type of identifier: which identifier to filter on
func (r PurchaseOrdersAPIGetByFilterRequest) IdentifiersIdentifierType(identifiersIdentifierType PurchaseOrderShipmentIdentifierTypeValue) PurchaseOrdersAPIGetByFilterRequest {
	r.identifiersIdentifierType = &identifiersIdentifierType
	return r
}

// The value (of the selected type) to filter on
func (r PurchaseOrdersAPIGetByFilterRequest) IdentifiersModels(identifiersModels []string) PurchaseOrdersAPIGetByFilterRequest {
	r.identifiersModels = &identifiersModels
	return r
}

func (r PurchaseOrdersAPIGetByFilterRequest) ShippedDateRangeFromDate(shippedDateRangeFromDate time.Time) PurchaseOrdersAPIGetByFilterRequest {
	r.shippedDateRangeFromDate = &shippedDateRangeFromDate
	return r
}

func (r PurchaseOrdersAPIGetByFilterRequest) ShippedDateRangeToDate(shippedDateRangeToDate time.Time) PurchaseOrdersAPIGetByFilterRequest {
	r.shippedDateRangeToDate = &shippedDateRangeToDate
	return r
}

func (r PurchaseOrdersAPIGetByFilterRequest) CreateDateRangeFromDate(createDateRangeFromDate time.Time) PurchaseOrdersAPIGetByFilterRequest {
	r.createDateRangeFromDate = &createDateRangeFromDate
	return r
}

func (r PurchaseOrdersAPIGetByFilterRequest) CreateDateRangeToDate(createDateRangeToDate time.Time) PurchaseOrdersAPIGetByFilterRequest {
	r.createDateRangeToDate = &createDateRangeToDate
	return r
}

func (r PurchaseOrdersAPIGetByFilterRequest) UpdateDateRangeFromDate(updateDateRangeFromDate time.Time) PurchaseOrdersAPIGetByFilterRequest {
	r.updateDateRangeFromDate = &updateDateRangeFromDate
	return r
}

func (r PurchaseOrdersAPIGetByFilterRequest) UpdateDateRangeToDate(updateDateRangeToDate time.Time) PurchaseOrdersAPIGetByFilterRequest {
	r.updateDateRangeToDate = &updateDateRangeToDate
	return r
}

// The Bill of Lading number. Multiple shipments can have the same Bill of Lading number
func (r PurchaseOrdersAPIGetByFilterRequest) BillOfLadingNumber(billOfLadingNumber string) PurchaseOrdersAPIGetByFilterRequest {
	r.billOfLadingNumber = &billOfLadingNumber
	return r
}

// The name of the carrier
func (r PurchaseOrdersAPIGetByFilterRequest) CarrierName(carrierName string) PurchaseOrdersAPIGetByFilterRequest {
	r.carrierName = &carrierName
	return r
}

func (r PurchaseOrdersAPIGetByFilterRequest) PageIndex(pageIndex int32) PurchaseOrdersAPIGetByFilterRequest {
	r.pageIndex = &pageIndex
	return r
}

func (r PurchaseOrdersAPIGetByFilterRequest) PageSize(pageSize int32) PurchaseOrdersAPIGetByFilterRequest {
	r.pageSize = &pageSize
	return r
}

func (r PurchaseOrdersAPIGetByFilterRequest) Execute() (*CollectionOfIPurchaseOrderShipmentByFilter, *http.Response, error) {
	return r.ApiService.GetByFilterExecute(r)
}

/*
GetByFilter Gets purchase order shipments by filter

Gets purchase order shipments based on the available filters.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return PurchaseOrdersAPIGetByFilterRequest
*/
func (a *PurchaseOrdersAPIService) GetByFilter(ctx context.Context) PurchaseOrdersAPIGetByFilterRequest {
	return PurchaseOrdersAPIGetByFilterRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return CollectionOfIPurchaseOrderShipmentByFilter
func (a *PurchaseOrdersAPIService) GetByFilterExecute(r PurchaseOrdersAPIGetByFilterRequest) (*CollectionOfIPurchaseOrderShipmentByFilter, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CollectionOfIPurchaseOrderShipmentByFilter
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PurchaseOrdersAPIService.GetByFilter")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/purchase-orders/shipments/merchant"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.channelId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "channelId", r.channelId, "")
	}
	if r.identifiersIdentifierType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "identifiers.identifierType", r.identifiersIdentifierType, "")
	}
	if r.identifiersModels != nil {
		t := *r.identifiersModels
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "identifiers.models", s.Index(i).Interface(), "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "identifiers.models", t, "multi")
		}
	}
	if r.shippedDateRangeFromDate != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "shippedDateRange.fromDate", r.shippedDateRangeFromDate, "")
	}
	if r.shippedDateRangeToDate != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "shippedDateRange.toDate", r.shippedDateRangeToDate, "")
	}
	if r.createDateRangeFromDate != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "createDateRange.fromDate", r.createDateRangeFromDate, "")
	}
	if r.createDateRangeToDate != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "createDateRange.toDate", r.createDateRangeToDate, "")
	}
	if r.updateDateRangeFromDate != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "updateDateRange.fromDate", r.updateDateRangeFromDate, "")
	}
	if r.updateDateRangeToDate != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "updateDateRange.toDate", r.updateDateRangeToDate, "")
	}
	if r.billOfLadingNumber != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "billOfLadingNumber", r.billOfLadingNumber, "")
	}
	if r.carrierName != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "carrierName", r.carrierName, "")
	}
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

type PurchaseOrdersAPIGetByFilter_0Request struct {
	ctx context.Context
	ApiService *PurchaseOrdersAPIService
	identifiersIdentifierType *PurchaseOrderIdentifierType
	identifiersModels *[]string
	statuses *[]ModulesPurchaseOrderStatus
	orderDateRangeFromDate *time.Time
	orderDateRangeToDate *time.Time
	createDateRangeFromDate *time.Time
	createDateRangeToDate *time.Time
	updateDateRangeFromDate *time.Time
	updateDateRangeToDate *time.Time
	channelIds *[]int32
	type_ *ModulesPurchaseOrderType
	pageIndex *int32
	pageSize *int32
}

// The type of identifier: which identifier to filter on
func (r PurchaseOrdersAPIGetByFilter_0Request) IdentifiersIdentifierType(identifiersIdentifierType PurchaseOrderIdentifierType) PurchaseOrdersAPIGetByFilter_0Request {
	r.identifiersIdentifierType = &identifiersIdentifierType
	return r
}

// The value (of the selected type) to filter on
func (r PurchaseOrdersAPIGetByFilter_0Request) IdentifiersModels(identifiersModels []string) PurchaseOrdersAPIGetByFilter_0Request {
	r.identifiersModels = &identifiersModels
	return r
}

func (r PurchaseOrdersAPIGetByFilter_0Request) Statuses(statuses []ModulesPurchaseOrderStatus) PurchaseOrdersAPIGetByFilter_0Request {
	r.statuses = &statuses
	return r
}

func (r PurchaseOrdersAPIGetByFilter_0Request) OrderDateRangeFromDate(orderDateRangeFromDate time.Time) PurchaseOrdersAPIGetByFilter_0Request {
	r.orderDateRangeFromDate = &orderDateRangeFromDate
	return r
}

func (r PurchaseOrdersAPIGetByFilter_0Request) OrderDateRangeToDate(orderDateRangeToDate time.Time) PurchaseOrdersAPIGetByFilter_0Request {
	r.orderDateRangeToDate = &orderDateRangeToDate
	return r
}

func (r PurchaseOrdersAPIGetByFilter_0Request) CreateDateRangeFromDate(createDateRangeFromDate time.Time) PurchaseOrdersAPIGetByFilter_0Request {
	r.createDateRangeFromDate = &createDateRangeFromDate
	return r
}

func (r PurchaseOrdersAPIGetByFilter_0Request) CreateDateRangeToDate(createDateRangeToDate time.Time) PurchaseOrdersAPIGetByFilter_0Request {
	r.createDateRangeToDate = &createDateRangeToDate
	return r
}

func (r PurchaseOrdersAPIGetByFilter_0Request) UpdateDateRangeFromDate(updateDateRangeFromDate time.Time) PurchaseOrdersAPIGetByFilter_0Request {
	r.updateDateRangeFromDate = &updateDateRangeFromDate
	return r
}

func (r PurchaseOrdersAPIGetByFilter_0Request) UpdateDateRangeToDate(updateDateRangeToDate time.Time) PurchaseOrdersAPIGetByFilter_0Request {
	r.updateDateRangeToDate = &updateDateRangeToDate
	return r
}

func (r PurchaseOrdersAPIGetByFilter_0Request) ChannelIds(channelIds []int32) PurchaseOrdersAPIGetByFilter_0Request {
	r.channelIds = &channelIds
	return r
}

func (r PurchaseOrdersAPIGetByFilter_0Request) Type_(type_ ModulesPurchaseOrderType) PurchaseOrdersAPIGetByFilter_0Request {
	r.type_ = &type_
	return r
}

func (r PurchaseOrdersAPIGetByFilter_0Request) PageIndex(pageIndex int32) PurchaseOrdersAPIGetByFilter_0Request {
	r.pageIndex = &pageIndex
	return r
}

func (r PurchaseOrdersAPIGetByFilter_0Request) PageSize(pageSize int32) PurchaseOrdersAPIGetByFilter_0Request {
	r.pageSize = &pageSize
	return r
}

func (r PurchaseOrdersAPIGetByFilter_0Request) Execute() (*CollectionOfIPurchaseOrderByFilter, *http.Response, error) {
	return r.ApiService.GetByFilter_1Execute(r)
}

/*
GetByFilter_0 Gets purchase orders by filter

Gets purchase orders based on the available filters.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return PurchaseOrdersAPIGetByFilter_0Request
*/
func (a *PurchaseOrdersAPIService) GetByFilter_1(ctx context.Context) PurchaseOrdersAPIGetByFilter_0Request {
	return PurchaseOrdersAPIGetByFilter_0Request{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return CollectionOfIPurchaseOrderByFilter
func (a *PurchaseOrdersAPIService) GetByFilter_1Execute(r PurchaseOrdersAPIGetByFilter_0Request) (*CollectionOfIPurchaseOrderByFilter, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CollectionOfIPurchaseOrderByFilter
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PurchaseOrdersAPIService.GetByFilter_1")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/purchase-orders"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.identifiersIdentifierType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "identifiers.identifierType", r.identifiersIdentifierType, "")
	}
	if r.identifiersModels != nil {
		t := *r.identifiersModels
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "identifiers.models", s.Index(i).Interface(), "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "identifiers.models", t, "multi")
		}
	}
	if r.statuses != nil {
		t := *r.statuses
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "statuses", s.Index(i).Interface(), "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "statuses", t, "multi")
		}
	}
	if r.orderDateRangeFromDate != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "orderDateRange.fromDate", r.orderDateRangeFromDate, "")
	}
	if r.orderDateRangeToDate != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "orderDateRange.toDate", r.orderDateRangeToDate, "")
	}
	if r.createDateRangeFromDate != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "createDateRange.fromDate", r.createDateRangeFromDate, "")
	}
	if r.createDateRangeToDate != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "createDateRange.toDate", r.createDateRangeToDate, "")
	}
	if r.updateDateRangeFromDate != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "updateDateRange.fromDate", r.updateDateRangeFromDate, "")
	}
	if r.updateDateRangeToDate != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "updateDateRange.toDate", r.updateDateRangeToDate, "")
	}
	if r.channelIds != nil {
		t := *r.channelIds
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "channelIds", s.Index(i).Interface(), "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "channelIds", t, "multi")
		}
	}
	if r.type_ != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "type", r.type_, "")
	}
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

type PurchaseOrdersAPIUpdateRequest struct {
	ctx context.Context
	ApiService *PurchaseOrdersAPIService
	singleMerchantUpdatePurchaseOrderShipmentRequest *SingleMerchantUpdatePurchaseOrderShipmentRequest
}

func (r PurchaseOrdersAPIUpdateRequest) SingleMerchantUpdatePurchaseOrderShipmentRequest(singleMerchantUpdatePurchaseOrderShipmentRequest SingleMerchantUpdatePurchaseOrderShipmentRequest) PurchaseOrdersAPIUpdateRequest {
	r.singleMerchantUpdatePurchaseOrderShipmentRequest = &singleMerchantUpdatePurchaseOrderShipmentRequest
	return r
}

func (r PurchaseOrdersAPIUpdateRequest) Execute() (*ApiResponse, *http.Response, error) {
	return r.ApiService.UpdateExecute(r)
}

/*
Update Update a purchase order shipment.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return PurchaseOrdersAPIUpdateRequest
*/
func (a *PurchaseOrdersAPIService) Update(ctx context.Context) PurchaseOrdersAPIUpdateRequest {
	return PurchaseOrdersAPIUpdateRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ApiResponse
func (a *PurchaseOrdersAPIService) UpdateExecute(r PurchaseOrdersAPIUpdateRequest) (*ApiResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ApiResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PurchaseOrdersAPIService.Update")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/purchase-orders/shipments"

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
	localVarPostBody = r.singleMerchantUpdatePurchaseOrderShipmentRequest
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
