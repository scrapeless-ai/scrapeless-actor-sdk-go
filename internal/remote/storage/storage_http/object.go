package storage_http

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/scrapeless-ai/scrapeless-actor-sdk-go/env"
	request2 "github.com/scrapeless-ai/scrapeless-actor-sdk-go/internal/remote/request"
	log "github.com/sirupsen/logrus"
	"github.com/tidwall/gjson"
	"io"
	"mime/multipart"
	"net/http"
)

func (c *Client) ListBuckets(ctx context.Context, page, size int) (*Object, error) {
	body, err := request2.Request(ctx, request2.ReqInfo{
		Method: http.MethodGet,
		Url:    fmt.Sprintf("%s/api/v1/object/buckets?page=%d&pageSize=%d", env.ScrapelessApiHost, page, size),
		Headers: map[string]string{
			env.HTTPHeader: env.Token,
		},
	})
	log.Infof("list buckets body :%s\n", body)
	if err != nil {
		log.Errorf("list buckets err:%v\n", err)
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
	var respData Object
	err = json.Unmarshal(marshal, &respData)
	if err != nil {
		log.Error("unmarshal resp error :", err)
		return nil, err
	}
	return &respData, nil
}

func (c *Client) CreateBucket(ctx context.Context, req *CreateBucketRequest) (string, error) {
	reqBody, err := json.Marshal(req)
	if err != nil {
		return "", err
	}
	body, err := request2.Request(ctx, request2.ReqInfo{
		Method: http.MethodPost,
		Url:    fmt.Sprintf("%s/api/v1/object/buckets", env.ScrapelessApiHost),
		Body:   string(reqBody),
		Headers: map[string]string{
			env.HTTPHeader: env.Token,
		},
	})
	log.Infof("create bucket body :%s\n", body)
	if err != nil {
		log.Errorf("create bucket err:%v\n", err)
		return "", err
	}
	var resp request2.RespInfo
	err = json.Unmarshal([]byte(body), &resp)
	if err != nil {
		log.Error("unmarshal resp error :", err)
		return "", err
	}
	if resp.Err {
		return "", fmt.Errorf("create bucket err:%s", resp.Msg)
	}
	id := gjson.Parse(body).Get("data.id").String()
	if id != "" {
		return id, nil
	}
	return "", fmt.Errorf("create bucket err:%s", resp.Msg)
}

func (c *Client) DeleteBucket(ctx context.Context, bucketId string) (bool, error) {
	body, err := request2.Request(ctx, request2.ReqInfo{
		Method: http.MethodDelete,
		Url:    fmt.Sprintf("%s/api/v1/object/buckets/%s", env.ScrapelessApiHost, bucketId),
		Headers: map[string]string{
			env.HTTPHeader: env.Token,
		},
	})
	log.Infof("del bucket body :%s\n", body)
	if err != nil {
		log.Errorf("del bucket err:%v\n", err)
		return false, err
	}
	var resp request2.RespInfo
	err = json.Unmarshal([]byte(body), &resp)
	if err != nil {
		log.Error("unmarshal resp error :", err)
		return false, err
	}
	if resp.Err {
		return false, fmt.Errorf("get bucket err:%s", resp.Msg)
	}
	if ok := gjson.Parse(body).Get("data.success").Bool(); ok {
		return ok, nil
	}
	return false, fmt.Errorf("del bucket err:%s", resp.Msg)
}
func (c *Client) GetBucket(ctx context.Context, bucketId string) (*Bucket, error) {
	body, err := request2.Request(ctx, request2.ReqInfo{
		Method: http.MethodGet,
		Url:    fmt.Sprintf("%s/api/v1/object/buckets/%s", env.ScrapelessApiHost, bucketId),
		Headers: map[string]string{
			env.HTTPHeader: env.Token,
		},
	})
	log.Infof("get bucket body :%s\n", body)
	if err != nil {
		log.Errorf("get bucket err:%v\n", err)
		return nil, err
	}
	var resp request2.RespInfo
	err = json.Unmarshal([]byte(body), &resp)
	if err != nil {
		log.Error("unmarshal resp error :", err)
		return nil, err
	}
	if resp.Err {
		return nil, fmt.Errorf("get bucket err:%s", resp.Msg)
	}
	marshal, _ := json.Marshal(&resp.Data)
	var respData Bucket
	err = json.Unmarshal(marshal, &respData)
	if err != nil {
		log.Error("unmarshal resp error :", err)
		return nil, err
	}
	return &respData, nil
}

func (c *Client) ListObjects(ctx context.Context, req *ListObjectsRequest) (*ObjectList, error) {
	body, err := request2.Request(ctx, request2.ReqInfo{
		Method: http.MethodGet,
		Url:    fmt.Sprintf("%s/api/v1/object/buckets/%s/objects", env.ScrapelessApiHost, req.BucketId),
		Headers: map[string]string{
			env.HTTPHeader: env.Token,
		},
	})
	log.Infof("list objects body :%s\n", body)
	if err != nil {
		log.Errorf("list objects err:%v\n", err)
		return nil, err
	}
	var resp request2.RespInfo
	err = json.Unmarshal([]byte(body), &resp)
	if err != nil {
		log.Error("unmarshal resp error :", err)
		return nil, err
	}
	if resp.Err {
		return nil, fmt.Errorf("get bucket err:%s", resp.Msg)
	}
	marshal, _ := json.Marshal(&resp.Data)
	var respData ObjectList
	err = json.Unmarshal(marshal, &respData)
	if err != nil {
		log.Error("unmarshal resp error :", err)
		return nil, err
	}
	return &respData, nil
}

func (c *Client) GetObject(ctx context.Context, req *ObjectRequest) ([]byte, error) {
	body, err := request2.Request(ctx, request2.ReqInfo{
		Method: http.MethodGet,
		Url:    fmt.Sprintf("%s/api/v1/object/buckets/%s/%s", env.ScrapelessApiHost, req.BucketId, req.ObjectId),
		Headers: map[string]string{
			env.HTTPHeader: env.Token,
		},
	})
	log.Infof("get object body :%s\n", body)
	if err != nil {
		log.Errorf("get object err:%v\n", err)
		return nil, err
	}
	var resp request2.RespInfo
	err = json.Unmarshal([]byte(body), &resp)
	if err != nil {
		log.Error("unmarshal resp error :", err)
		return []byte(body), nil
	}
	if resp.Err {
		return nil, fmt.Errorf("get object err:%s", resp.Msg)
	}
	return []byte(body), nil
}

func (c *Client) DeleteObject(ctx context.Context, req *ObjectRequest) (bool, error) {
	body, err := request2.Request(ctx, request2.ReqInfo{
		Method: http.MethodDelete,
		Url:    fmt.Sprintf("%s/api/v1/object/buckets/%s/%s", env.ScrapelessApiHost, req.BucketId, req.ObjectId),
		Headers: map[string]string{
			env.HTTPHeader: env.Token,
		},
	})
	log.Infof("del object body :%s\n", body)
	if err != nil {
		log.Errorf("del object err:%v\n", err)
		return false, err
	}
	var resp request2.RespInfo
	err = json.Unmarshal([]byte(body), &resp)
	if err != nil {
		log.Error("unmarshal resp error :", err)
		return false, err
	}
	if resp.Err {
		return false, fmt.Errorf("delete object err:%s", resp.Msg)
	}
	if ok := gjson.Parse(body).Get("data.success").Bool(); ok {
		return ok, nil
	}
	return false, fmt.Errorf("delete object err:%s", resp.Msg)
}

func (c *Client) PutObject(ctx context.Context, req *PutObjectRequest) (string, error) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, _ := writer.CreateFormFile("file", req.Filename)
	part.Write(req.Data)
	writer.WriteField("actorId", req.ActorId)
	writer.WriteField("runId", req.RunId)
	writer.Close()

	url := fmt.Sprintf("%s/api/v1/object/buckets/%s/object", env.ScrapelessApiHost, req.BucketId)
	request, _ := http.NewRequestWithContext(ctx, http.MethodPost, url, body)
	request.Header.Set("Content-Type", writer.FormDataContentType())
	request.Header.Set(env.HTTPHeader, env.Token)
	resp, err := c.client.Do(request)
	defer resp.Body.Close()
	all, _ := io.ReadAll(resp.Body)
	log.Infof("put object body :%s\n", string(all))
	if err != nil {
		log.Error("request error :", err)
		return "", err
	}
	var respInfo request2.RespInfo
	err = json.Unmarshal(all, &respInfo)
	if err != nil {
		log.Error("unmarshal resp error :", err)
		return "", err
	}
	if respInfo.Err {
		return "", fmt.Errorf("put object err:%s", respInfo.Msg)
	}
	objectId := gjson.Parse(string(all)).Get("data.object_id").String()
	if objectId == "" {
		return "", fmt.Errorf("put object err:%s", respInfo.Msg)
	}
	return objectId, nil
}
