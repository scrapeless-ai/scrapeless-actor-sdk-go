package storage_http

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/scrapeless-ai/scrapeless-actor-sdk-go/env"
	request2 "github.com/scrapeless-ai/scrapeless-actor-sdk-go/internal/remote/request"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func (c *Client) ListDatasets(ctx context.Context, req *ListDatasetsRequest) (*ListDatasetsResponse, error) {
	body, err := request2.Request(ctx, request2.ReqInfo{
		Method: http.MethodGet,
		Url:    fmt.Sprintf("%s/api/v1/dataset?actorId=%s&desc=%v&page=%d&pageSize=%d&runId=%s", env.ScrapelessApiHost, *req.ActorId, req.Desc, req.Page, req.PageSize, *req.RunId),
		Body:   "",
		Headers: map[string]string{
			env.HTTPHeader: env.Token,
		},
	})
	log.Infof("list dataset body:%s\n", body)
	if err != nil {
		log.Errorf("list dataset err:%v\n", err)
		return nil, err
	}
	var resp request2.RespInfo
	err = json.Unmarshal([]byte(body), &resp)
	if err != nil {
		log.Error("unmarshal resp error :", err)
		return nil, err
	}
	if resp.Err {
		return nil, fmt.Errorf("get dataset list err:%s", resp.Msg)
	}
	marshal, _ := json.Marshal(&resp.Data)
	var respData ListDatasetsResponse
	err = json.Unmarshal(marshal, &respData)
	if err != nil {
		log.Error("unmarshal resp error :", err)
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
		Url:    fmt.Sprintf("%s/api/v1/dataset", env.ScrapelessApiHost),
		Body:   string(reqBody),
		Headers: map[string]string{
			env.HTTPHeader: env.Token,
		},
	})
	log.Infof("create dataset body:%s\n", body)
	if err != nil {
		log.Errorf("create dataset err:%v\n", err)
		return nil, err
	}
	var resp request2.RespInfo
	err = json.Unmarshal([]byte(body), &resp)
	if err != nil {
		log.Error("unmarshal resp error :", err)
		return nil, err
	}
	if resp.Err {
		return nil, fmt.Errorf("creat dataset err:%s", resp.Msg)
	}
	marshal, _ := json.Marshal(&resp.Data)
	var respData Dataset
	err = json.Unmarshal(marshal, &respData)
	if err != nil {
		log.Error("unmarshal resp error :", err)
		return nil, err
	}
	return &respData, nil
}

func (c *Client) UpdateDataset(ctx context.Context, datasetID, name string) (bool, error) {
	body, err := request2.Request(ctx, request2.ReqInfo{
		Method: http.MethodPut,
		Url:    fmt.Sprintf("%s/api/v1/dataset/%s", env.ScrapelessApiHost, datasetID),
		Body:   fmt.Sprintf(`{"name":"%s"}`, name),
		Headers: map[string]string{
			env.HTTPHeader: env.Token,
		},
	})
	log.Infof(`up dataset body:%s`, body)
	if err != nil {
		log.Errorf("up dataset err:%v\n", err)
		return false, err
	}
	var resp request2.RespInfo
	err = json.Unmarshal([]byte(body), &resp)
	if err != nil {
		log.Error("unmarshal resp error :", err)
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
		Url:    fmt.Sprintf("%s/api/v1/dataset/%s", env.ScrapelessApiHost, datasetID),
		Headers: map[string]string{
			env.HTTPHeader: env.Token,
		},
	})
	log.Infof("del dataset body:%s\n", body)
	if err != nil {
		log.Errorf("del dataset err:%v\n", err)
		return false, err
	}
	var resp request2.RespInfo
	err = json.Unmarshal([]byte(body), &resp)
	if err != nil {
		log.Error("unmarshal resp error :", err)
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
		Url:    fmt.Sprintf("%s/api/v1/dataset/%s/items?page=%d&pageSize=%d&desc=%v", env.ScrapelessApiHost, req.DatasetId, req.Page, req.PageSize, req.Desc),
		Headers: map[string]string{
			env.HTTPHeader: env.Token,
		},
	})
	log.Infof("get dataset body:%s\n", body)
	if err != nil {
		log.Errorf("get dataset err:%v\n", err)
		return nil, err
	}
	var resp request2.RespInfo
	err = json.Unmarshal([]byte(body), &resp)
	if err != nil {
		log.Error("unmarshal resp error :", err)
		return nil, err
	}
	if resp.Err {
		return nil, fmt.Errorf("get dataset item err:%s", resp.Msg)
	}
	marshal, _ := json.Marshal(&resp.Data)
	var respData DatasetItem
	err = json.Unmarshal(marshal, &respData)
	if err != nil {
		log.Error("unmarshal resp error :", err)
		return nil, err
	}
	return &respData, nil
}

func (c *Client) AddDatasetItem(ctx context.Context, datasetId string, data []map[string]any) (bool, error) {
	reqBody, err := json.Marshal(map[string]any{
		"items": data,
	})
	if err != nil {
		log.Errorf("marshal dataset item err:%v\n", err)
		return false, err
	}
	body, err := request2.Request(ctx, request2.ReqInfo{
		Method: http.MethodPost,
		Url:    fmt.Sprintf("%s/api/v1/dataset/%s/items", env.ScrapelessApiHost, datasetId),
		Body:   string(reqBody),
		Headers: map[string]string{
			env.HTTPHeader: env.Token,
		},
	})
	log.Infof("add dataset item body:%s\n", body)
	if err != nil {
		log.Errorf("add dataset item err:%v\n", err)
		return false, err
	}
	var resp request2.RespInfo
	err = json.Unmarshal([]byte(body), &resp)
	if err != nil {
		log.Errorf("unmarshal dataset item err:%v\n", err)
		return false, err
	}
	if resp.Err {
		log.Errorf("add dataset item err:%s\n", resp.Msg)
		return false, fmt.Errorf("add dataset item err:%s", resp.Msg)
	}
	return true, nil
}
