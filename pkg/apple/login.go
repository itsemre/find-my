package apple

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// c := http.Client{
// 	// incorporate jar
// 	&http.Transport{
// 	}
// }

const (
	loginURL = "https://setup.icloud.com/setup/ws/1/login"
	origin   = "https://www.icloud.com"
)

func Login(appleID, password string, extededLogin bool) (*LoginResponseBody, error) {
	requestBody, err := json.Marshal(LoginRequestBody{
		AppleID:       appleID,
		Password:      password,
		ExtendedLogin: extededLogin,
	})
	if err != nil {
		return nil, fmt.Errorf("can not marshal http request body: %w", err)
	}

	request, err := http.NewRequest("POST", loginURL, bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, fmt.Errorf("can not create http request: %w", err)
	}
	request.Header.Set("Origin", origin)
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, fmt.Errorf("can not send http request: %w", err)
	} else if response.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(response.Body)
		return nil, fmt.Errorf("can not send http request: %s", string(body))
	}

	var responseBody LoginResponseBody
	err = json.NewDecoder(response.Body).Decode(&responseBody)
	if err != nil {
		return nil, fmt.Errorf("can not decode http response body: %w", err)
	}
	return &responseBody, nil
}
