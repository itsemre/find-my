# find-my

A (borderline hacky) Go client package for Apple's Find My service

## Installation

```bash
go get -u "github.com/itsemre/find-my"
```

## Usage 

##### NewClient

```go
// NewClient creates and initializes a new Find My client object
func NewClient(appleID, password string, extendedLogin bool) (*fmy.Client, error) 
```

##### GetUserInfo

```go
// GetUserInfo uses the provided Apple ID credentials and logs in to iCloud to return the user's information
func (c *fmy.Client) GetUserInfo() (*fmy.userInfo, error) 
```

##### GetFindMyInfo

```go
// GetFindMyInfo all information for the user's Find My service. It can optionally return the user's
// family members' information as well, with: getFamily = true
func (c *fmy.Client) GetFindMyInfo(getFamily bool) (*fmy.findMyInfo, error)
```

##### GetDeviceInfo

```go
// GetDeviceInfo returns the information of the given device ID
func (c *fmy.Client) GetDeviceInfo(id string) (*fmy.deviceInfo, error)
```

##### GetDeviceLocation

```go
// GetDeviceLocation returns the location information of the given device ID
func (c *fmy.Client) GetDeviceLocation(id string) (*fmy.deviceLocation, error)
```

##### AlertDevice

```go
// AlertDevice plays a sound on the device with the given ID
func (c *fmy.Client) AlertDevice(id string) error 
```

## Example

```go
package main

import (
    fmy "github.com/itsemre/find-my"
)

func main() {
    c, err := fmy.NewClient("<apple_id>", "<password>", true)
    if err != nil {
        panic(err)
    }

    dInfo, err := c.GetFindMyInfo(false) 
    if err != nil {
        panic(err)
    }

    for _, d := range dInfo.Content {
        err = c.AlertDevice(d.ID)
        if err != nil {
            panic(err)
        }
    }
}
```

The above example gets the IDs of all of the user's devices and uses them to play a sound on them.



---

> This project is based on the pre existing work done by [Matt Kruse](https://github.com/matt-kruse). For more information, check out this [repository](https://github.com/matt-kruse/find-my-iphone).