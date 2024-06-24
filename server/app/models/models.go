package models

import (
	"encoding/json"
	"fmt"
	"strings"
)

type CurrencyRate struct {
	Code       string `json:"code"`
	Codein     string `json:"codein"`
	Name       string `json:"name"`
	High       string `json:"high"`
	Low        string `json:"low"`
	VarBid     string `json:"varBid"`
	PctChange  string `json:"pctChange"`
	Bid        string `json:"bid"`
	Ask        string `json:"ask"`
	Timestamp  string `json:"timestamp"`
	CreateDate string `json:"create_date"`
}

func (cr *CurrencyRate) ConvertJSONToCurrencyRate(mainKey string, data []byte) error {
	partsKey := strings.Split(mainKey, "-")
	mainKey = strings.Join(partsKey, "")
	bytesCurrency, err := extractFirstPart(mainKey, data)
	if err != nil {
		return err
	}

	err = json.Unmarshal(bytesCurrency, &cr)
	if err != nil {
		return fmt.Errorf("cannot convert to CurrencyRate: %s", err.Error())
	}
	return nil
}

func extractFirstPart(mainKey string, data []byte) ([]byte, error) {
	var extract map[string]interface{}
	var bytes []byte

	err := json.Unmarshal(data, &extract)
	if err != nil {
		return bytes, fmt.Errorf("cannot extract the correct info: %s", err.Error())
	}

	value, ok := extract[mainKey]
	if !ok {
		var nameKey string
		for key := range extract {
			nameKey = key
			break
		}
		return bytes, fmt.Errorf("cannot find the first key, found key '%s', expect key '%s'", nameKey, mainKey)
	}

	bytes, err = json.Marshal(&value)
	if err != nil {
		return bytes, fmt.Errorf("cannot convert the content: %s", err.Error())
	}

	return bytes, nil
}

func (cr CurrencyRate) SaveOnTxt() []byte {
	return []byte(fmt.Sprintf("Dólar:{%s}", cr.Bid))
}
