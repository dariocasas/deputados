package http_client

import "time"

type HttpClientInputDTO = struct {
	Url     string
	Timeout int
}

type HttpClientWithRetryInputDTO = struct {
	Url     string
	Timeout int
	MinLen  int
	Retry   int
}

type HttpClientOutputDTO = struct {
	Duration time.Duration
	Error    error
	Result   *[]byte
}

type HttpClientInterface interface {
	RequestPageWithRetry(input *HttpClientWithRetryInputDTO) *HttpClientOutputDTO
	RequestPage(input *HttpClientInputDTO) *HttpClientOutputDTO
	RequestJson(input *HttpClientInputDTO) *HttpClientOutputDTO
}
