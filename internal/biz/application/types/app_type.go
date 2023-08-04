package types

import "strings"

const (
	AppTypeBackend AppType = "BACKEND"
	AppTypeWeb     AppType = "WEB"
	AppTypeAndroid AppType = "ANDROID"
	AppTypeIOS     AppType = "IOS"
)

var AllAppTypes = []AppType{
	AppTypeBackend,
	AppTypeWeb,
	AppTypeAndroid,
	AppTypeIOS,
}

type AppType string

func (at AppType) String() string {
	return string(at)
}

func (at AppType) Upper() AppType {
	return AppType(strings.ToUpper(at.String()))
}

func (at AppType) IsSupported() bool {
	for _, sType := range AllAppTypes {
		if at.Upper() == sType {
			return true
		}
	}
	return false
}

func (at AppType) IsBackend() bool {
	return at.Upper() == AppTypeBackend
}

func (at AppType) IsWeb() bool {
	return at.Upper() == AppTypeWeb
}

func (at AppType) IsAndroid() bool {
	return at.Upper() == AppTypeAndroid
}

func (at AppType) IsIOS() bool {
	return at.Upper() == AppTypeIOS
}
