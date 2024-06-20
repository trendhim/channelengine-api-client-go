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
	"time"
	"reflect"
)


// NotificationAPIService NotificationAPI service
type NotificationAPIService service

type NotificationAPINotificationIndexRequest struct {
	ctx context.Context
	ApiService *NotificationAPIService
	fromDate *time.Time
	toDate *time.Time
	types *[]NotificationType
	merchantOrderNos *[]string
	channelOrderNos *[]string
	merchantReturnNos *[]string
	channelReturnNos *[]string
	merchantShipmentNos *[]string
	page *int32
}

// Filter on the notification date, starting from this date. This date is inclusive.
func (r NotificationAPINotificationIndexRequest) FromDate(fromDate time.Time) NotificationAPINotificationIndexRequest {
	r.fromDate = &fromDate
	return r
}

// Filter on the notification date, until this date. This date is exclusive.
func (r NotificationAPINotificationIndexRequest) ToDate(toDate time.Time) NotificationAPINotificationIndexRequest {
	r.toDate = &toDate
	return r
}

// Notification type(s) to filter on.
func (r NotificationAPINotificationIndexRequest) Types(types []NotificationType) NotificationAPINotificationIndexRequest {
	r.types = &types
	return r
}

// Filter on unique order reference used by the merchant.
func (r NotificationAPINotificationIndexRequest) MerchantOrderNos(merchantOrderNos []string) NotificationAPINotificationIndexRequest {
	r.merchantOrderNos = &merchantOrderNos
	return r
}

// Filter on unique order reference used by the channel.
func (r NotificationAPINotificationIndexRequest) ChannelOrderNos(channelOrderNos []string) NotificationAPINotificationIndexRequest {
	r.channelOrderNos = &channelOrderNos
	return r
}

// Filter on unique return reference used by the merchant.
func (r NotificationAPINotificationIndexRequest) MerchantReturnNos(merchantReturnNos []string) NotificationAPINotificationIndexRequest {
	r.merchantReturnNos = &merchantReturnNos
	return r
}

// Filter on unique return reference used by the channel.
func (r NotificationAPINotificationIndexRequest) ChannelReturnNos(channelReturnNos []string) NotificationAPINotificationIndexRequest {
	r.channelReturnNos = &channelReturnNos
	return r
}

// Filter on unique shipment reference used by the merchant.
func (r NotificationAPINotificationIndexRequest) MerchantShipmentNos(merchantShipmentNos []string) NotificationAPINotificationIndexRequest {
	r.merchantShipmentNos = &merchantShipmentNos
	return r
}

// The page to filter on. Starts at 1.
func (r NotificationAPINotificationIndexRequest) Page(page int32) NotificationAPINotificationIndexRequest {
	r.page = &page
	return r
}

func (r NotificationAPINotificationIndexRequest) Execute() (*CollectionOfMerchantNotificationResponse, *http.Response, error) {
	return r.ApiService.NotificationIndexExecute(r)
}

/*
NotificationIndex Gets notifications

Gets ChannelEngine notifications based on the available filters.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return NotificationAPINotificationIndexRequest
*/
func (a *NotificationAPIService) NotificationIndex(ctx context.Context) NotificationAPINotificationIndexRequest {
	return NotificationAPINotificationIndexRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return CollectionOfMerchantNotificationResponse
func (a *NotificationAPIService) NotificationIndexExecute(r NotificationAPINotificationIndexRequest) (*CollectionOfMerchantNotificationResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CollectionOfMerchantNotificationResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NotificationAPIService.NotificationIndex")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/notifications"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.fromDate != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fromDate", r.fromDate, "")
	}
	if r.toDate != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "toDate", r.toDate, "")
	}
	if r.types != nil {
		t := *r.types
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "types", s.Index(i).Interface(), "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "types", t, "multi")
		}
	}
	if r.merchantOrderNos != nil {
		t := *r.merchantOrderNos
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "merchantOrderNos", s.Index(i).Interface(), "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "merchantOrderNos", t, "multi")
		}
	}
	if r.channelOrderNos != nil {
		t := *r.channelOrderNos
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "channelOrderNos", s.Index(i).Interface(), "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "channelOrderNos", t, "multi")
		}
	}
	if r.merchantReturnNos != nil {
		t := *r.merchantReturnNos
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "merchantReturnNos", s.Index(i).Interface(), "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "merchantReturnNos", t, "multi")
		}
	}
	if r.channelReturnNos != nil {
		t := *r.channelReturnNos
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "channelReturnNos", s.Index(i).Interface(), "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "channelReturnNos", t, "multi")
		}
	}
	if r.merchantShipmentNos != nil {
		t := *r.merchantShipmentNos
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "merchantShipmentNos", s.Index(i).Interface(), "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "merchantShipmentNos", t, "multi")
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
