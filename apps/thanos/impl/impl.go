package impl

import (
	"errors"
	"fmt"
	"net/http"

	"gitee.com/zuoyangs/go-rcp/config"
)

func getThanosData(encodedQuery string) (*http.Response, error) {

	config.Init()
	thanos_url, err := config.GetKey("thanos-config", "url")
	if err != nil {
		return nil, fmt.Errorf("Failed to get thanos URL: %v", err)
	}

	url := fmt.Sprintf("%s/api/v1/query?query=%s", thanos_url, encodedQuery)

	//fmt.Printf("\nthanos query url(apps/thanos/impl/impl.go):%s", url)

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %v", err)
	}

	if resp.Body == nil {
		return nil, errors.New("response body is empty")
	}
	return resp, nil
}
