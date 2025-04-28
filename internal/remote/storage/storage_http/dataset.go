package storage_http

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/scrapeless-ai/scrapeless-actor-sdk-go/env"
	"github.com/scrapeless-ai/scrapeless-actor-sdk-go/internal/config"
	request2 "github.com/scrapeless-ai/scrapeless-actor-sdk-go/internal/remote/request"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func (c *Client) ListDatasets(ctx context.Context, req *ListDatasetsRequest) (*ListDatasetsResponse, error) {
	body, err := request2.Request(ctx, request2.ReqInfo{
		Method: http.MethodGet,
		Url:    fmt.Sprintf("%s/scrapeless/actor/api/v1/dataset?actorId=%s&desc=%v&page=%d&pageSize=%d&runId=%s", config.StorageServiceHost, *req.ActorId, req.Desc, req.Page, req.PageSize, *req.RunId),
		Body:   "",
		Headers: map[string]string{
			env.HTTPHeader: config.Token,
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
		return nil, fmt.Errorf("get dataset list err:%s", resp.Msg)
	}
	marshal, _ := json.Marshal(&resp.Data)
	var respData ListDatasetsResponse
	err = json.Unmarshal(marshal, &respData)
	if err != nil {
		return nil, err
	}
	return &respData, nil
}

func (c *Client) CreateDataset(ctx context.Context, req *CreateDatasetRequest) (*Dataset, error) {
	reqBody, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	body, err := request2.Request(ctx, request2.ReqInfo{
		Method: http.MethodPost,
		Url:    fmt.Sprintf("%s/scrapeless/actor/api/v1/dataset", config.StorageServiceHost),
		Body:   string(reqBody),
		Headers: map[string]string{
			env.HTTPHeader: config.Token,
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
		return nil, fmt.Errorf("creat dataset err:%s", resp.Msg)
	}
	marshal, _ := json.Marshal(&resp.Data)
	var respData Dataset
	err = json.Unmarshal(marshal, &respData)
	if err != nil {
		return nil, err
	}
	return &respData, nil
}

func (c *Client) UpdateDataset(ctx context.Context, datasetID, name string) (bool, error) {
	body, err := request2.Request(ctx, request2.ReqInfo{
		Method: http.MethodPut,
		Url:    fmt.Sprintf("%s/scrapeless/actor/api/v1/dataset/%s", config.StorageServiceHost, datasetID),
		Body:   fmt.Sprintf(`{"name":"%s"}`, name),
		Headers: map[string]string{
			env.HTTPHeader: config.Token,
		},
	})
	if err != nil {
		return false, err
	}
	fmt.Println(body)
	var resp request2.RespInfo
	err = json.Unmarshal([]byte(body), &resp)
	if err != nil {
		return false, err
	}
	if resp.Err {
		return false, fmt.Errorf("edit dataset err:%s", resp.Msg)
	}
	return true, nil
}
func (c *Client) DelDataset(ctx context.Context, datasetID string) (bool, error) {
	body, err := request2.Request(ctx, request2.ReqInfo{
		Method: http.MethodDelete,
		Url:    fmt.Sprintf("%s/scrapeless/actor/api/v1/dataset/%s", config.StorageServiceHost, datasetID),
		Headers: map[string]string{
			env.HTTPHeader: config.Token,
		},
	})
	if err != nil {
		return false, err
	}
	fmt.Println(body)
	var resp request2.RespInfo
	err = json.Unmarshal([]byte(body), &resp)
	if err != nil {
		return false, err
	}
	if resp.Err {
		return false, fmt.Errorf("del dataset err:%s", resp.Msg)
	}
	return true, nil
}

func (c *Client) GetDataset(ctx context.Context, req *GetDataset) (*DatasetItem, error) {
	body, err := request2.Request(ctx, request2.ReqInfo{
		Method: http.MethodGet,
		Url:    fmt.Sprintf("%s/scrapeless/actor/api/v1/dataset/%s/items?page=%d&pageSize=%d&desc=%v", config.StorageServiceHost, req.DatasetId, req.Page, req.PageSize, req.Desc),
		Headers: map[string]string{
			env.HTTPHeader: config.Token,
		},
	})
	if err != nil {
		return nil, err
	}
	fmt.Println(body)
	var resp request2.RespInfo
	err = json.Unmarshal([]byte(body), &resp)
	if err != nil {
		return nil, err
	}
	if resp.Err {
		return nil, fmt.Errorf("get dataset item err:%s", resp.Msg)
	}
	marshal, _ := json.Marshal(&resp.Data)
	var respData DatasetItem
	err = json.Unmarshal(marshal, &respData)
	if err != nil {
		return nil, err
	}
	return &respData, nil
}

func (c *Client) AddDatasetItem(ctx context.Context, datasetId string, data []map[string]any) (bool, error) {
	reqBody, err := json.Marshal(map[string]any{
		"items": data,
	})
	if err != nil {
		log.Error("marshal dataset item err:%v\n", err)
		return false, err
	}
	body, err := request2.Request(ctx, request2.ReqInfo{
		Method: http.MethodPost,
		Url:    fmt.Sprintf("%s/scrapeless/actor/api/v1/dataset/%s/items", config.StorageServiceHost, datasetId),
		Body:   string(reqBody),
		Headers: map[string]string{
			env.HTTPHeader: config.Token,
		},
	})
	if err != nil {
		return false, err
	}
	fmt.Println(body)
	var resp request2.RespInfo
	err = json.Unmarshal([]byte(body), &resp)
	if err != nil {
		return false, err
	}
	if resp.Err {
		return false, fmt.Errorf("add dataset item err:%s", resp.Msg)
	}
	return true, nil
}
