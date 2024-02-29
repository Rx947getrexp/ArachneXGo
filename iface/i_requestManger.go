package iface

import "time"

type (
	IRequestManger interface {
		AddRequest(IUri) error
		GetNextRequest() IUri
		MarkURLAsProcess(string)
		SetRequestheaders(key, value string) IRequestManger
		SetRequestTimeout(time.Duration) IRequestManger
	}

	IUri interface {
		GetAddress() string
		GetMethod() string
	}
)
