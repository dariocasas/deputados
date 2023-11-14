package http_client

import (
	"context"
	"errors"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/dariocasas/deputados3/src/configs"
	"github.com/dariocasas/deputados3/src/pkg/elapsed_time"
)

type HttpClient struct {
	config *configs.Config
}

func NewHttpClient(config *configs.Config) *HttpClient {
	return &HttpClient{
		config: config,
	}
}

func (c HttpClient) RequestPageWithRetry(input *HttpClientWithRetryInputDTO) *HttpClientOutputDTO {

	elapsedTime := elapsed_time.NewElapsedTime()
	retry := 0

	httpClientOutputDTO := c.RequestPage(
		&HttpClientInputDTO{
			Url:     input.Url,
			Timeout: input.Timeout,
		},
	)
	for {
		if httpClientOutputDTO.Error != nil {
			log.Printf("RequestPageWithRetry retry %s", httpClientOutputDTO.Error.Error())
		} else {
			if httpClientOutputDTO.Result != nil {
				if len(*httpClientOutputDTO.Result) > input.MinLen {
					break
				} else {
					log.Printf("RequestPageWithRetry retry len<MinLen")
				}
			} else {
				log.Printf("RequestPageWithRetry retry httpClientOutputDTO.Result==nil")
			}
		}

		retry++
		if retry >= input.Retry {
			break
		}
		httpClientOutputDTO = c.RequestPage(
			&HttpClientInputDTO{
				Url:     input.Url,
				Timeout: input.Timeout,
			},
		)
	}

	return &HttpClientOutputDTO{
		Duration: elapsedTime.Elapsed(),
		Error:    nil,
		Result:   httpClientOutputDTO.Result,
	}

}

func (c HttpClient) RequestPage(input *HttpClientInputDTO) *HttpClientOutputDTO {

	elapsedTime := elapsed_time.NewElapsedTime()

	timeout := input.Timeout
	if timeout == 0 {
		timeout = c.config.HttpClientTimeOutMilliSecs
	}

	ctx, cancel := context.WithTimeout(
		context.Background(),
		time.Duration(timeout)*time.Millisecond,
	)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, input.Url, nil)
	if err != nil {
		log.Printf("RequestPage: could not create request: %s\n", err)
		return &HttpClientOutputDTO{
			Duration: elapsedTime.Elapsed(),
			Error:    err,
			Result:   nil,
		}
	}
	req.Header.Set("Content-Type", "text/html charset=utf-8")

	client := http.Client{}

	res, err := client.Do(req)
	if err != nil {
		log.Printf("RequestPage: error making http request: %s\n", err)
		return &HttpClientOutputDTO{
			Duration: elapsedTime.Elapsed(),
			Error:    err,
			Result:   nil,
		}
	}
	defer res.Body.Close()

	if res.StatusCode == http.StatusOK {
		b, err := io.ReadAll(res.Body)
		if err != nil {
			return &HttpClientOutputDTO{
				Duration: elapsedTime.Elapsed(),
				Error:    err,
				Result:   nil,
			}
		}
		log.Printf("RequestPage: page len %d\n", len(b))

		return &HttpClientOutputDTO{
			Duration: elapsedTime.Elapsed(),
			Error:    nil,
			Result:   &b,
		}

	}

	log.Printf("RequestPage: error making http request: %s\n", errors.New(res.Status))
	return &HttpClientOutputDTO{
		Duration: elapsedTime.Elapsed(),
		Error:    errors.New(res.Status),
		Result:   nil,
	}
}

func (c HttpClient) RequestJson(input *HttpClientInputDTO) *HttpClientOutputDTO {

	elapsedTime := elapsed_time.NewElapsedTime()

	timeout := input.Timeout
	if timeout == 0 {
		timeout = c.config.HttpClientTimeOutMilliSecs
	}

	ctx, cancel := context.WithTimeout(
		context.Background(),
		time.Duration(timeout)*time.Millisecond,
	)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, input.Url, nil)
	if err != nil {
		// log.Printf("RequestJson: could not create request: %s\n", err)
		return &HttpClientOutputDTO{
			Duration: elapsedTime.Elapsed(),
			Error:    err,
			Result:   nil,
		}
	}
	req.Header.Set("Content-Type", "application/json")

	client := http.Client{}

	res, err := client.Do(req)
	if err != nil {
		// log.Printf("RequestJson: error making http request: %s\n", err)
		return &HttpClientOutputDTO{
			Duration: elapsedTime.Elapsed(),
			Error:    err,
			Result:   nil,
		}
	}
	defer res.Body.Close()

	if res.StatusCode == http.StatusOK {
		b, err := io.ReadAll(res.Body)
		if err != nil {
			return &HttpClientOutputDTO{
				Duration: elapsedTime.Elapsed(),
				Error:    err,
				Result:   nil,
			}
		}
		return &HttpClientOutputDTO{
			Duration: elapsedTime.Elapsed(),
			Error:    nil,
			Result:   &b,
		}

	}

	return &HttpClientOutputDTO{
		Duration: elapsedTime.Elapsed(),
		Error:    errors.New(res.Status),
		Result:   nil,
	}
}
