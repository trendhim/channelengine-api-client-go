/*
ChannelEngine Channel API

Testing ShipmentAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package channel

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/trendhim/channelengine/channel"
)

func Test_channel_ShipmentAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ShipmentAPIService ShipmentIndex", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ShipmentAPI.ShipmentIndex(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ShipmentAPIService ShipmentShippingLabel", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var merchantShipmentNo string

		resp, httpRes, err := apiClient.ShipmentAPI.ShipmentShippingLabel(context.Background(), merchantShipmentNo).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
