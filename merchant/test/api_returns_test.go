/*
ChannelEngine Merchant API

Testing ReturnsAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package merchant

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/trendhim/channelengine/merchant"
)

func Test_merchant_ReturnsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ReturnsAPIService ReturnAcknowledge", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ReturnsAPI.ReturnAcknowledge(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ReturnsAPIService ReturnAcknowledge_1", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ReturnsAPI.ReturnAcknowledge_0(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ReturnsAPIService ReturnCreate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ReturnsAPI.ReturnCreate(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ReturnsAPIService ReturnDeclareForMerchant", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ReturnsAPI.ReturnDeclareForMerchant(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ReturnsAPIService ReturnGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var identifier string

		resp, httpRes, err := apiClient.ReturnsAPI.ReturnGet(context.Background(), identifier).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ReturnsAPIService ReturnGetByFilter", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ReturnsAPI.ReturnGetByFilter(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ReturnsAPIService ReturnGetByMerchantOrderNo", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var merchantOrderNo string

		resp, httpRes, err := apiClient.ReturnsAPI.ReturnGetByMerchantOrderNo(context.Background(), merchantOrderNo).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ReturnsAPIService ReturnGetDeclaredByChannel", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ReturnsAPI.ReturnGetDeclaredByChannel(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ReturnsAPIService ReturnGetReturns", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ReturnsAPI.ReturnGetReturns(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ReturnsAPIService ReturnGetUnhandled", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ReturnsAPI.ReturnGetUnhandled(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ReturnsAPIService ReturnHandle", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ReturnsAPI.ReturnHandle(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ReturnsAPIService ReturnUpdateForMerchant", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ReturnsAPI.ReturnUpdateForMerchant(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}