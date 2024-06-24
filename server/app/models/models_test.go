package models_test

import (
	"reflect"
	"server-go/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestModels(t *testing.T) {
	correctJSONMock := `{"USDBRL":{"code":"USD","codein":"BRL","name":"Dólar Americano/Real Brasileiro","high":"5.4259","low":"5.4194","varBid":"0.0059","pctChange":"0.11","bid":"5.425","ask":"5.4258","timestamp":"1718659799","create_date":"2024-06-17 18:29:59"}}`

	wrongCurrencyJSONMock := `{"EURBRL":{"code":"EUR","codein":"BRL","name":"Euro/Real Brasileiro","high":"5.82","low":"5.82","varBid":"0","pctChange":"0","bid":"5.816","ask":"5.824","timestamp":"1718670273","create_date":"2024-06-17 21:24:33"}}`

	testsCases := []struct {
		name    string
		msg     string
		strJson string
		mainKey string
	}{
		{
			name:    "test success",
			msg:     "",
			strJson: correctJSONMock,
			mainKey: "USDBRL",
		},
		{
			name:    "test error cannot extract",
			msg:     "cannot extract the correct info: invalid character '\\'' looking for beginning of object key string",
			strJson: "{'asasasa'",
			mainKey: "USDBRL",
		},
		{
			name:    "test error cannot extract the correct key",
			msg:     "cannot find the first key, found key 'EURBRL', expect key 'USDBRL'",
			strJson: wrongCurrencyJSONMock,
			mainKey: "USDBRL",
		},
	}

	for _, tt := range testsCases {
		t.Run(tt.name, func(t *testing.T) {
			var cr models.CurrencyRate
			err := cr.ConvertJSONToCurrencyRate(tt.mainKey, []byte(tt.strJson))
			if err != nil {
				assert.NotNil(t, err)
				assert.Equal(t, tt.msg, err.Error())
				return
			}
			equals := reflect.DeepEqual(cr, models.CurrencyRate{
				Code:       "USD",
				Codein:     "BRL",
				Name:       "Dólar Americano/Real Brasileiro",
				High:       "5.4259",
				Low:        "5.4194",
				VarBid:     "0.0059",
				PctChange:  "0.11",
				Bid:        "5.425",
				Ask:        "5.4258",
				Timestamp:  "1718659799",
				CreateDate: "2024-06-17 18:29:59",
			})
			assert.True(t, equals)
		})
	}
}
