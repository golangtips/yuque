package internal

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/golangtips/yuque/sdk/intf"
)

var _ intf.IYuQue = (*YuQue)(nil)

type YuQue struct {
	UserAgent string //应用名称
	baseURL   string
	token     string
	namespace string
	client    *http.Client
}

func NewYuQue(baseURL, token, userAgent, namespace string) *YuQue {

	// GetRepoDocList 请求会比较久，适当增加超时时长
	client := &http.Client{
		Timeout: 30 * time.Second,
	}

	return &YuQue{
		UserAgent: userAgent,
		baseURL:   baseURL,
		token:     token,
		namespace: namespace,
		client:    client,
	}
}

func (y *YuQue) GetRepoDocList(ctx context.Context, request *intf.GetRepoDocListRequest) (*intf.GetRepoDocListResponse, error) {

	namespace := y.namespace
	if request.Namespace != "" {
		namespace = request.Namespace
	}

	url := fmt.Sprintf("%s/repos/%s/docs", y.baseURL, namespace)

	req := y.buildHTTPRequest("GET", url, nil)
	q := req.URL.Query()
	if request.Offset > 0 {
		q.Add("offset", strconv.Itoa(request.Offset))
	}

	if request.Limit > 0 {
		q.Add("limit", strconv.Itoa(request.Limit))
	}

	req.URL.RawQuery = q.Encode()
	resp, err := y.client.Do(req)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("ioutil: %w", err)
	}

	//log.Println(string(body))

	var response intf.GetRepoDocListResponse
	if err = json.Unmarshal(body, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

func (y *YuQue) GetRepoDocDetail(_ context.Context, request *intf.GetRepoDocDetailRequest) (*intf.GetRepoDocDetailResponse, error) {

	namespace := y.namespace
	if request.Namespace != "" {
		namespace = request.Namespace
	}

	url := fmt.Sprintf("%s/repos/%s/docs/%s", y.baseURL, namespace, request.Slug)
	req := y.buildHTTPRequest("GET", url, nil)

	resp, err := y.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("ioutil: %w", err)
	}

	var detail intf.GetRepoDocDetailResponse
	if err = json.Unmarshal(body, &detail); err != nil {
		return nil, err
	}

	return &detail, nil
}

func (y *YuQue) Search(ctx context.Context, request *intf.SearchRequest) (*intf.SearchResponse, error) {

	return &intf.SearchResponse{
		//
	}, nil
}

// buildHTTPRequest 辅助函数
func (y *YuQue) buildHTTPRequest(method, url string, body io.Reader) *http.Request {
	req, _ := http.NewRequest(method, url, body)
	req.Header.Add("User-Agent", y.UserAgent)
	req.Header.Add("X-Auth-Token", y.token)
	return req
}
