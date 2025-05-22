package actor

import (
	"context"
	"encoding/json"
	"github.com/scrapeless-ai/scrapeless-actor-sdk-go/scrapeless/storage/object"
	"testing"
)

func TestStorageListBuckets(t *testing.T) {
	ac := New(WithStorage())
	value, err := ac.Storage.GetObject().ListBuckets(context.Background(), 1, 10)
	t.Log(value)
	t.Error(err)
}

func TestStorageCreateBucket(t *testing.T) {
	ac := New(WithStorage())
	value, _, err := ac.Storage.GetObject().CreateBucket(context.Background(), "test", "test")
	t.Log(value)
	t.Error(err)
}

func TestStorageDeleteBucket(t *testing.T) {
	ac := New(WithStorage())
	ok, err := ac.Storage.GetObject().DeleteBucket(context.Background())
	t.Log(ok)
	t.Error(err)
}

func TestStorageGetBucket(t *testing.T) {
	ac := New(WithStorage())
	ok, err := ac.Storage.GetObject().GetBucket(context.Background())
	t.Log(ok)
	t.Error(err)
}
func TestStorageListObjects(t *testing.T) {
	ac := New(WithStorage())
	objects, err := ac.Storage.GetObject().List(context.Background(), "", 0, 0)
	marshal, _ := json.Marshal(objects)
	t.Log(string(marshal))
	t.Error(err)
}

func TestStoragePutObject(t *testing.T) {
	oh := object.NewObjHttp("")
	objects, err := oh.Put(context.Background(), "test.json", []byte(`{"name":"jack"}`)) //e87b6a6c-516d-4b6b-86ab-1d47c9a7fd36
	t.Log(objects)
	t.Error(err)
}
