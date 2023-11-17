/*
ChannelEngine Merchant API

Testing CompetitionPriceAPIService

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

func Test_merchant_CompetitionPriceAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test CompetitionPriceAPIService CompetitionPricesGetBuyBoxPrices", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.CompetitionPriceAPI.CompetitionPricesGetBuyBoxPrices(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
