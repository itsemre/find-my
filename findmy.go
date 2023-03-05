package fmy

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	pjar "github.com/juju/persistent-cookiejar"
)

const (
	origin = "https://www.icloud.com"

	loginURL = "https://setup.icloud.com/setup/ws/1/login"

	findMyEndpoint     = "/fmipservice/client/web/initClient"
	findMyAppName      = "iCloud Find (Web)"
	findMyAppVersion   = "2.0"
	findMyApiVersion   = "3.0"
	findMyTimezone     = "US/Eastern"
	findMyInactiveTime = 3571

	playSoundEndpoint = "/fmipservice/client/web/playSound"
	playSoundSubject  = "Amazon Echo Find My iPhone Alert"
)

type Client struct {
	AppleID       string
	Password      string
	ExtendedLogin bool

	UserInfo   *userInfo
	FindMyInfo *findMyInfo

	httpClient *http.Client
	jar        *pjar.Jar

	findMyURL *url.URL
	alertURL  *url.URL
}

// NewClient creates and initializes a new Find My client object
func NewClient(appleID, password string, extendedLogin bool) (*Client, error) {
	if appleID == "" || password == "" {
		return nil, errors.New("please provide Apple ID credentials")
	}

	// Create a cookie jar
	j, err := pjar.New(nil)
	if err != nil {
		return nil, fmt.Errorf("error creating cookie jar: %s", err.Error())
	}

	// Initialize client with cookie jar and Apple ID credentials
	c := Client{
		AppleID:       appleID,
		Password:      password,
		ExtendedLogin: extendedLogin,
		UserInfo:      nil,
		FindMyInfo:    nil,
		httpClient: &http.Client{
			Jar: j,
		},
		jar:       j,
		findMyURL: nil,
	}
	return &c, nil
}

// post creates and sends a POST request and returns the response decoded into the given object
func (c *Client) post(url string, decodeObject, body any) (any, error) {
	// Encode body to JSON
	jsonBody, err := json.Marshal(body)
	if err != nil {
		return nil, fmt.Errorf("error marshaling http request body: %s", err.Error())
	}

	// Create HTTP request
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, fmt.Errorf("error creating HTTP request: %s", err.Error())
	}

	// Send HTTP request
	resp, err := c.sendRequest(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// If theres nothing to decode return with the HTTP response
	if decodeObject == nil {
		return resp, nil
	}

	// Decode response into the given object
	err = json.NewDecoder(resp.Body).Decode(decodeObject)
	if err != nil {
		return nil, fmt.Errorf("error decoding HTTP response body: %s", err.Error())
	}
	return decodeObject, nil
}

// sendRequest sets the required request headers prior to sending it
func (c *Client) sendRequest(r *http.Request) (*http.Response, error) {
	// Format cookies
	cookies := ""
	for _, cookie := range c.jar.AllCookies() {
		cookies += fmt.Sprintf("%s=\"%s\";", cookie.Name, cookie.Value)
	}
	// Set headers
	r.Header.Set("Origin", origin)
	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("Cookie", cookies)
	// Send request
	resp, err := c.httpClient.Do(r)
	if resp.StatusCode != http.StatusOK {
		if err == nil {
			body, _ := ioutil.ReadAll(resp.Body)
			err = errors.New(string(body))
		}
		return nil, fmt.Errorf("HTTP request failed with status %d and error: %s", resp.StatusCode, err)
	}
	return resp, nil
}

// GetUserInfo uses the provided Apple ID credentials and logs in to iCloud to return the user's information
func (c *Client) GetUserInfo() (*userInfo, error) {
	// Make request
	resp, err := c.post(loginURL, &userInfo{}, struct {
		AppleID       string `json:"apple_id"`
		Password      string `json:"password"`
		ExtendedLogin bool   `json:"extended_login"`
	}{
		c.AppleID,
		c.Password,
		c.ExtendedLogin,
	})
	if err != nil {
		return nil, err
	}
	// Cast the response to *userInfo
	c.UserInfo = resp.(*userInfo)
	return c.UserInfo, nil
}

// GetFindMyInfo all information for the user's Find My service. It can optionally return the user's
// family members' information as well, with: getFamily = true
func (c *Client) GetFindMyInfo(getFamily bool) (*findMyInfo, error) {
	// If the user information does not exist, retrieve it
	if c.UserInfo == nil {
		_, err := c.GetUserInfo()
		if err != nil {
			return nil, err
		}
	}

	// Check if the user's Find My service is enabled
	if c.UserInfo.Webservices.Findme.Status != "active" {
		return nil, fmt.Errorf("error: the Find My service is not active for user %s (%s)", c.UserInfo.DsInfo.FullName, c.UserInfo.DsInfo.AppleID)
	}

	// Construct the Find My URL
	if c.findMyURL == nil {
		u, err := url.Parse(c.UserInfo.Webservices.Findme.URL)
		if err != nil {
			return nil, fmt.Errorf("error parsing Find My URL: %s", err.Error())
		}
		u = u.JoinPath(findMyEndpoint)
		c.findMyURL = u
	}

	// Make request
	resp, err := c.post(c.findMyURL.String(), &findMyInfo{}, struct {
		Context interface{} `json:"clientContext"`
	}{
		Context: struct {
			AppName      string `json:"appName"`
			AppVersion   string `json:"appVersion"`
			ApiVersion   string `json:"apiVersion"`
			Timezone     string `json:"timezone"`
			InactiveTime int    `json:"inactiveTime"`
			GetFamily    bool   `json:"fmly"`
		}{
			findMyAppName,
			findMyAppVersion,
			findMyApiVersion,
			findMyTimezone,
			findMyInactiveTime,
			getFamily,
		},
	})
	if err != nil {
		return nil, err
	}
	// Cast the response to *userInfo
	c.FindMyInfo = resp.(*findMyInfo)
	return c.FindMyInfo, nil
}

// GetDeviceInfo returns the information of the given device ID
func (c *Client) GetDeviceInfo(id string) (*deviceInfo, error) {
	// If the device information does not exist, retrieve it
	if c.FindMyInfo == nil {
		_, err := c.GetFindMyInfo(false)
		if err != nil {
			return nil, err
		}
	}
	// Match the given device ID in the list of devices in user's account
	for _, d := range c.FindMyInfo.Content {
		if d.ID == id {
			return &d, nil
		}
	}
	return nil, fmt.Errorf("error: could not find device ID '%s'", id)
}

// GetDeviceLocation returns the location information of the given device ID
func (c *Client) GetDeviceLocation(id string) (*deviceLocation, error) {
	dInfo, err := c.GetDeviceInfo(id)
	if err != nil {
		return nil, err
	}
	return &dInfo.Location, nil
}

// AlertDevice plays a sound on the device with the given ID
func (c *Client) AlertDevice(id string) error {
	// If the user information does not exist, retrieve it
	if c.UserInfo == nil {
		_, err := c.GetUserInfo()
		if err != nil {
			return err
		}
	}

	// Construct the alert URL
	if c.alertURL == nil {
		u, err := url.Parse(c.UserInfo.Webservices.Findme.URL)
		if err != nil {
			return fmt.Errorf("error parsing alert URL: %s", err.Error())
		}
		u = u.JoinPath(playSoundEndpoint)
		c.alertURL = u
	}

	// Make request
	_, err := c.post(c.alertURL.String(), nil, struct {
		Subject string `json:"subject"`
		Device  string `json:"device"`
	}{
		playSoundSubject,
		id,
	})
	if err != nil {
		return err
	}
	return nil

}
