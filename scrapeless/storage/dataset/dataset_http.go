package dataset

import (
	"context"
	"github.com/scrapeless-ai/scrapeless-actor-sdk-go/env"
	"github.com/scrapeless-ai/scrapeless-actor-sdk-go/internal/code"
	"github.com/scrapeless-ai/scrapeless-actor-sdk-go/internal/remote/storage/storage_http"
	log "github.com/sirupsen/logrus"
)

type DSHttp struct {
	datasetId string
}

func NewDSHttp(datasetId ...string) Dataset {
	log.Info("dataset http init")
	if storage_http.Default() == nil {
		storage_http.Init()
	}
	dh := &DSHttp{datasetId: env.Env.DatasetId}
	if len(datasetId) > 0 {
		dh.datasetId = datasetId[0]
	}
	return dh
}

func (ds *DSHttp) ListDatasets(ctx context.Context, page int64, pageSize int64, desc bool) (*ListDatasetsResponse, error) {
	if page < 1 {
		page = 1
	}
	if pageSize < 10 {
		pageSize = 10
	}
	datasets, err := storage_http.Default().ListDatasets(ctx, &storage_http.ListDatasetsRequest{
		Page:     page,
		PageSize: pageSize,
		ActorId:  &env.Env.ActorId,
		RunId:    &env.Env.RunId,

		Desc: desc,
	})
	if err != nil {
		log.Errorf("failed to list datasets: %v\n", code.Format(err))
		return nil, code.Format(err)
	}
	var itemArray []DatasetInfo
	for _, item := range datasets.Items {
		itemArray = append(itemArray, DatasetInfo{
			Id:         item.Id,
			Name:       item.Name,
			ActorId:    item.ActorId,
			RunId:      item.RunId,
			Fields:     item.Fields,
			CreatedAt:  item.CreatedAt,
			UpdatedAt:  item.UpdatedAt,
			AccessedAt: item.AccessedAt,
		})
	}
	return &ListDatasetsResponse{
		Items: itemArray,
		Total: datasets.Total,
	}, nil
}

func (ds *DSHttp) CreateDataset(ctx context.Context, name string) (id string, err error) {
	dataset, err := storage_http.Default().CreateDataset(ctx, &storage_http.CreateDatasetRequest{
		Name:    name,
		ActorId: &env.Env.ActorId,
		RunId:   &env.Env.RunId,
	})
	if err != nil {
		log.Errorf("failed to create dataset: %v\n", code.Format(err))
		return "", code.Format(err)
	}
	return dataset.Id, nil
}

func (ds *DSHttp) UpdateDataset(ctx context.Context, name string) (bool, error) {
	ok, err := storage_http.Default().UpdateDataset(ctx, ds.datasetId, name)
	if err != nil {
		log.Errorf("failed to update dataset: %v\n", code.Format(err))
		return false, code.Format(err)
	}
	return ok, nil
}

func (ds *DSHttp) DelDataset(ctx context.Context) (bool, error) {
	ok, err := storage_http.Default().DelDataset(ctx, ds.datasetId)
	if err != nil {
		log.Errorf("failed to delete dataset: %v\n", code.Format(err))
		return false, code.Format(err)
	}
	return ok, nil
}

func (ds *DSHttp) addItemsWithId(ctx context.Context, items []map[string]any) (bool, error) {
	ok, err := storage_http.Default().AddDatasetItem(ctx, ds.datasetId, items)
	if err != nil {
		log.Errorf("failed to add items: %v\n", err)
		return false, code.Format(err)
	}
	return ok, nil
}

func (ds *DSHttp) getItemsWithId(ctx context.Context, page int, pageSize int, desc bool) (*ItemsResponse, error) {
	if page < 1 {
		page = 1
	}
	if pageSize < 10 {
		pageSize = 10
	}
	items, err := storage_http.Default().GetDataset(ctx, &storage_http.GetDataset{
		DatasetId: ds.datasetId,
		Desc:      desc,
		Page:      page,
		PageSize:  pageSize,
	})
	if err != nil {
		log.Errorf("failed to get items: %v\n", code.Format(err))
		return nil, code.Format(err)
	}
	var itemArray []map[string]any
	for _, item := range items.Items {
		itemArray = append(itemArray, item)
	}
	return &ItemsResponse{
		Items: itemArray,
		Total: items.Total,
	}, nil
}

func (ds *DSHttp) AddItems(ctx context.Context, items []map[string]any) (bool, error) {
	return ds.addItemsWithId(ctx, items)
}
func (ds *DSHttp) GetItems(ctx context.Context, page int, pageSize int, desc bool) (*ItemsResponse, error) {
	return ds.getItemsWithId(ctx, page, pageSize, desc)
}

func (ds *DSHttp) Close() error {
	return storage_http.Default().Close()
}
