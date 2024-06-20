/*
ChannelEngine Merchant API

Testing ProductAPIService

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

func Test_merchant_ProductAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ProductAPIService ProductBulkDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ProductAPI.ProductBulkDelete(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProductAPIService ProductBulkPatch", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ProductAPI.ProductBulkPatch(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProductAPIService ProductBulkPatchExtraDataItems", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ProductAPI.ProductBulkPatchExtraDataItems(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProductAPIService ProductCreate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ProductAPI.ProductCreate(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProductAPIService ProductDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var merchantProductNo string

		resp, httpRes, err := apiClient.ProductAPI.ProductDelete(context.Background(), merchantProductNo).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProductAPIService ProductFreeze", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ProductAPI.ProductFreeze(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProductAPIService ProductGetByFilter", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ProductAPI.ProductGetByFilter(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProductAPIService ProductGetByMerchantProductNo", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var merchantProductNo string

		resp, httpRes, err := apiClient.ProductAPI.ProductGetByMerchantProductNo(context.Background(), merchantProductNo).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProductAPIService ProductPatch", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var merchantProductNo string

		resp, httpRes, err := apiClient.ProductAPI.ProductPatch(context.Background(), merchantProductNo).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProductAPIService ProductPatchExtraDataItems", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ProductAPI.ProductPatchExtraDataItems(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
