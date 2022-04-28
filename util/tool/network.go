/*
	@Author  johnny
	@Author  johnny.he@nextsmartship.com
	@Version v1.0.0
	@File    network
	@Date    2022/3/7 11:38 上午
	@Desc    http 请求（GET/POST）
*/

package tool

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"

	svrErr "github.com/NextSmartShip/common/util/processerror"
	"github.com/pkg/errors"
)

const (
	HTTPTimeout      = 30
	OuterHTTPTimeout = 180
)

func SendJSONRequestToAPI(url string, method string, postData interface{},
	headers map[string]string,
) ([]byte, error) {
	var (
		req           *http.Request
		jsonStr, data []byte
		err           error
	)

	// 30秒超时
	transport := &http.Transport{
		TLSNextProto: map[string]func(string, *tls.Conn) http.RoundTripper{},
	}

	// 10秒超时
	client := &http.Client{
		Timeout:   OuterHTTPTimeout * time.Second,
		Transport: transport,
	}

	var body io.Reader

	// 构造request
	if postData != nil {
		r, ok := postData.(string)
		if ok {
			body = strings.NewReader(r)
		} else {
			jsonStr, err = json.Marshal(postData)
			if err != nil {
				return data, errors.Wrap(err, "")
			}

			body = strings.NewReader(string(jsonStr))
		}
	}

	req, err = http.NewRequestWithContext(context.Background(), method, url, body)
	if err != nil {
		return data, errors.Wrap(err, "")
	}

	// 设置Header
	ProcessHTTPHeader(req, headers)

	req.Close = true

	// 发起请求
	rsp, err := client.Do(req)
	if err != nil {
		return data, errors.Wrap(err, "")
	}

	defer rsp.Body.Close()

	if rsp.StatusCode != http.StatusOK {
		return data, errors.Unwrap(svrErr.HTTPStatus(url, rsp.StatusCode, postData))
	}

	// 解析返回.
	data, err = ioutil.ReadAll(rsp.Body)
	if err == nil {
		return data, nil
	}

	if strings.Contains(err.Error(), "unexpected EOF") {
		return data, nil
	}

	return nil, errors.Wrap(err, "")
}

func SendGetRequest(url string, params map[string]interface{},
	headers map[string]string,
) ([]byte, error) {
	var (
		req  *http.Request
		data []byte
		err  error
	)

	client := &http.Client{}

	req, err = http.NewRequestWithContext(context.Background(), http.MethodGet, url, nil)
	if err != nil {
		return data, errors.Wrap(err, "")
	}

	// 设置Header.
	ProcessHTTPHeader(req, headers)

	// 设置Param.
	ProcessHTTPParam(req, params)

	rsp, err := client.Do(req)
	if err != nil {
		return data, errors.Wrap(err, "")
	}

	defer rsp.Body.Close()

	if rsp.StatusCode != http.StatusOK {
		return data, errors.Unwrap(svrErr.HTTPStatus(req.URL.String(), rsp.StatusCode, params))
	}

	// 解析返回
	data, err = ioutil.ReadAll(rsp.Body)

	return data, errors.Wrap(err, "")
}

func SendRequestToBaseAPI(url, method string, body io.Reader, headers map[string]string) ([]byte, error) {
	var (
		req  *http.Request
		data []byte
		err  error
		resp *http.Response
	)

	// 30秒超时
	// transport := &http.Transport{
	// 	TLSNextProto: map[string]func(string, *tls.Conn) http.RoundTripper{},
	// }

	client := &http.Client{
		Timeout: HTTPTimeout * time.Second,
		// Transport: transport,
	}

	req, err = http.NewRequestWithContext(context.Background(), method, url, body)
	if err != nil {
		return data, errors.Wrap(err, "")
	}

	// set Header
	ProcessHTTPHeader(req, headers)

	resp, err = client.Do(req)
	if err != nil {
		return data, errors.Wrap(err, "")
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		err = errors.Unwrap(svrErr.HTTPStatus(url, resp.StatusCode, body))

		return data, errors.Wrap(err, "")
	}

	data, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return data, errors.Wrap(err, "")
	}

	return data, errors.Wrap(err, "")
}

func ProcessHTTPHeader(req *http.Request, headers map[string]string) {
	if len(headers) > 0 {
		for key, value := range headers {
			req.Header.Set(key, value)
		}
	}
}

func ProcessHTTPParam(req *http.Request, params map[string]interface{}) {
	query := req.URL.Query()

	if len(params) > 0 {
		for key, value := range params {
			switch valueType := value.(type) {
			case string:
				query.Add(key, valueType)
			case int:
				val := strconv.Itoa(valueType)
				query.Add(key, val)
			case []string:
				for _, item := range valueType {
					query.Add(key, item)
				}
			}
		}
	}

	req.URL.RawQuery = query.Encode()
}
