package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"frontendmasters.com/go/crypto/datatypes"
)

const apiUrl = "https://cex.io/api/ticker/%s/USD"

func GetRate(currency string) (*datatypes.Rate, error) {

	if len(currency) != 3 {
		return nil, fmt.Errorf("3 characters required: %v", len(currency))
	}

	res, err := http.Get(fmt.Sprintf(apiUrl, currency))

	if err != nil {
		return nil, err
	}

	var response CEXResponse

	if res.StatusCode == http.StatusOK {
		bodyBytes, err := io.ReadAll(res.Body)

		if err != nil {
			return nil, err
		}

		err = json.Unmarshal(bodyBytes, &response)
		// fmt.Printf("Res: %v", response)
		if err != nil {
			return nil, err
		}

	} else {
		return nil, fmt.Errorf("status code received: %v", res.StatusCode)
	}

	rate := datatypes.Rate{Currency: currency, Price: response.Bid}

	return &rate, nil

}
