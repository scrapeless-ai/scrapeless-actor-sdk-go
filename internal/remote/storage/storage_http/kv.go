package storage_http

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/scrapeless-ai/scrapeless-actor-sdk-go/env"
	request2 "github.com/scrapeless-ai/scrapeless-actor-sdk-go/internal/remote/request"
	"github.com/tidwall/gjson"
	"net/http"
)

func (c *Client) ListNamespaces(ctx context.Context, page int, pageSize int, desc bool) (*KvNamespace, error) {
	body, err := request2.Request(ctx, request2.ReqInfo{
		Method: http.MethodGet,
		Url:    fmt.Sprintf("%s/api/v1/kv/namespaces?desc=%v&page=%d&pageSize=%d", env.ScrapelessApiHost, desc, page, pageSize),
		Body:   "",
		Headers: map[string]string{
			env.HTTPHeader: env.Token,
		},
	})
	if err != nil {
		return nil, err
	}
	var resp request2.RespInfo
	fmt.Println(body)
	err = json.Unmarshal([]byte(body), &resp)
	if err != nil {
		return nil, err
	}
	if resp.Err {
		return nil, fmt.Errorf("get dataset list err:%s", resp.Msg)
	}
	marshal, _ := json.Marshal(&resp.Data)
	var respData KvNamespace
	err = json.Unmarshal(marshal, &respData)
	if err != nil {
		return nil, err
	}
	return &respData, nil
}

func (c *Client) CreateNamespace(ctx context.Context, req *CreateKvNamespaceRequest) (string, error) {
	reqBody, err := json.Marshal(req)
	if err != nil {
		return "", err
	}
	body, err := request2.Request(ctx, request2.ReqInfo{
		Method: http.MethodPost,
		Url:    fmt.Sprintf("%s/api/v1/kv/namespaces", env.ScrapelessApiHost),
		Body:   string(reqBody),
		Headers: map[string]string{
			env.HTTPHeader: env.Token,
		},
	})
	if err != nil {
		return "", err
	}
	var resp request2.RespInfo
	err = json.Unmarshal([]byte(body), &resp)
	if err != nil {
		return "", err
	}
	if resp.Err {
		return "", fmt.Errorf("create namespace err:%s", resp.Msg)
	}
	id := gjson.Parse(body).Get("data.id").String()
	return id, nil
}
func (c *Client) GetNamespace(ctx context.Context, namespaceId string) (*KvNamespaceItem, error) {
	body, err := request2.Request(ctx, request2.ReqInfo{
		Method: http.MethodGet,
		Url:    fmt.Sprintf("%s/api/v1/kv/%s", env.ScrapelessApiHost, namespaceId),
		Headers: map[string]string{
			env.HTTPHeader: env.Token,
		},
	})
	if err != nil {
		return nil, err
	}
	var resp request2.RespInfo
	err = json.Unmarshal([]byte(body), &resp)
	if err != nil {
		return nil, err
	}
	if resp.Err {
		return nil, fmt.Errorf("get namespace err:%s", resp.Msg)
	}
	data := gjson.Parse(body).Get("data").String()
	var kvi KvNamespaceItem
	if err = json.Unmarshal([]byte(data), &kvi); err != nil {
		return nil, err
	}
	return &kvi, nil
}

func (c *Client) DelNamespace(ctx context.Context, namespaceId string) (bool, error) {
	body, err := request2.Request(ctx, request2.ReqInfo{
		Method: http.MethodDelete,
		Url:    fmt.Sprintf("%s/api/v1/kv/%s", env.ScrapelessApiHost, namespaceId),
		Headers: map[string]string{
			env.HTTPHeader: env.Token,
		},
	})
	if err != nil {
		return false, err
	}
	var resp request2.RespInfo
	err = json.Unmarshal([]byte(body), &resp)
	if err != nil {
		return false, err
	}
	if resp.Err {
		return false, fmt.Errorf("get namespace err:%s", resp.Msg)
	}
	ok := gjson.Parse(body).Get("data.success").Bool()

	return ok, nil
}

func (c *Client) RenameNamespace(ctx context.Context, namespaceId string, name string) (bool, error) {
	body, err := request2.Request(ctx, request2.ReqInfo{
		Method: http.MethodPut,
		Url:    fmt.Sprintf("%s/api/v1/kv/%s/rename", env.ScrapelessApiHost, namespaceId),
		Body:   fmt.Sprintf(`{"name":"%s"}`, name),
		Headers: map[string]string{
			env.HTTPHeader: env.Token,
		},
	})
	if err != nil {
		return false, err
	}
	var resp request2.RespInfo
	err = json.Unmarshal([]byte(body), &resp)
	if err != nil {
		return false, err
	}
	if resp.Err {
		return false, fmt.Errorf("get namespace err:%s", resp.Msg)
	}
	ok := gjson.Parse(body).Get("data.success").Bool()

	return ok, nil
}

func (c *Client) SetValue(ctx context.Context, req *SetValue) (bool, error) {
	reqBody := map[string]any{
		"expiration": req.Expiration,
		"key":        req.Key,
		"value":      req.Value,
	}
	reqBodyStr, err := json.Marshal(reqBody)
	if err != nil {
		return false, err
	}
	body, err := request2.Request(ctx, request2.ReqInfo{
		Method: http.MethodPut,
		Url:    fmt.Sprintf("%s/api/v1/kv/%s/key", env.ScrapelessApiHost, req.NamespaceId),
		Body:   string(reqBodyStr),
		Headers: map[string]string{
			env.HTTPHeader: env.Token,
		},
	})
	if err != nil {
		return false, err
	}
	var resp request2.RespInfo
	err = json.Unmarshal([]byte(body), &resp)
	if err != nil {
		return false, err
	}
	if resp.Err {
		return false, fmt.Errorf("get namespace err:%s", resp.Msg)
	}
	ok := gjson.Parse(body).Get("data.success").Bool()

	return ok, nil
}

func (c *Client) ListKeys(ctx context.Context, req *ListKeyInfo) (*KvKeys, error) {
	body, err := request2.Request(ctx, request2.ReqInfo{
		Method: http.MethodGet,
		Url:    fmt.Sprintf("%s/api/v1/kv/%s/keys?page=%d&pageSize=%d", env.ScrapelessApiHost, req.NamespaceId, req.Page, req.Size),
		Headers: map[string]string{
			env.HTTPHeader: env.Token,
		},
	})
	if err != nil {
		return nil, err
	}
	var resp request2.RespInfo
	err = json.Unmarshal([]byte(body), &resp)
	if err != nil {
		return nil, err
	}
	if resp.Err {
		return nil, fmt.Errorf("list Keys err:%s", resp.Msg)
	}
	marshal, _ := json.Marshal(&resp.Data)
	var respData KvKeys
	err = json.Unmarshal(marshal, &respData)
	if err != nil {
		return nil, err
	}
	return &respData, nil
}

func (c *Client) GetValue(ctx context.Context, namespaceId string, key string) (string, error) {
	body, err := request2.Request(ctx, request2.ReqInfo{
		Method: http.MethodGet,
		Url:    fmt.Sprintf("%s/api/v1/kv/%s/%s", env.ScrapelessApiHost, namespaceId, key),
		Headers: map[string]string{
			env.HTTPHeader: env.Token,
		},
	})
	if err != nil {
		return "", err
	}
	var resp request2.RespInfo
	err = json.Unmarshal([]byte(body), &resp)
	if err != nil {
		return "", err
	}
	if resp.Err {
		return "", fmt.Errorf("get value err:%s", resp.Msg)
	}
	data := gjson.Parse(body).Get("data").String()
	return data, nil
}

func (c *Client) DelValue(ctx context.Context, namespaceId string, key string) (bool, error) {
	body, err := request2.Request(ctx, request2.ReqInfo{
		Method: http.MethodDelete,
		Url:    fmt.Sprintf("%s/api/v1/kv/%s/%s", env.ScrapelessApiHost, namespaceId, key),
		Headers: map[string]string{
			env.HTTPHeader: env.Token,
		},
	})
	if err != nil {
		return false, err
	}
	var resp request2.RespInfo
	err = json.Unmarshal([]byte(body), &resp)
	if err != nil {
		return false, err
	}
	if resp.Err {
		return false, fmt.Errorf("list Keys err:%s", resp.Msg)
	}
	ok := gjson.Parse(body).Get("data.success").Bool()
	return ok, nil
}

func (c *Client) BulkSetValue(ctx context.Context, req *BulkSet) (int64, error) {
	reqBody, err := json.Marshal(req.Items)
	if err != nil {
		return 0, err
	}
	body, err := request2.Request(ctx, request2.ReqInfo{
		Method: http.MethodPost,
		Url:    fmt.Sprintf("%s/api/v1/kv/%s/bulk", env.ScrapelessApiHost, req.NamespaceId),
		Body:   fmt.Sprintf(`{"Items":%s}`, reqBody),
		Headers: map[string]string{
			env.HTTPHeader: env.Token,
		},
	})
	if err != nil {
		return 0, err
	}
	var resp request2.RespInfo
	err = json.Unmarshal([]byte(body), &resp)
	if err != nil {
		return 0, err
	}
	if resp.Err {
		return 0, fmt.Errorf("bulk set value err:%s", resp.Msg)
	}
	fmt.Println(body)
	successfulKeyCount := gjson.Parse(body).Get("data.successful_key_count").Int()
	return successfulKeyCount, nil
}

func (c *Client) BulkDelValue(ctx context.Context, namespaceId string, keys []string) (bool, error) {
	reqBody, err := json.Marshal(keys)
	if err != nil {
		return false, err
	}
	body, err := request2.Request(ctx, request2.ReqInfo{
		Method: http.MethodDelete,
		Url:    fmt.Sprintf("%s/api/v1/kv/%s/bulk", env.ScrapelessApiHost, namespaceId),
		Body:   fmt.Sprintf(`{"keys":%s}`, reqBody),
		Headers: map[string]string{
			env.HTTPHeader: env.Token,
		},
	})
	if err != nil {
		return false, err
	}
	var resp request2.RespInfo
	err = json.Unmarshal([]byte(body), &resp)
	if err != nil {
		return false, err
	}
	if resp.Err {
		return false, fmt.Errorf("bulk set value err:%s", resp.Msg)
	}
	ok := gjson.Parse(body).Get("data.success").Bool()
	return ok, nil
}
