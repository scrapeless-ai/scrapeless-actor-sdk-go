package scrapeless

import (
	"context"
	"testing"
)

func TestListDatasets(t *testing.T) {
	ac := New(WithStorage())
	ok, err := ac.Storage.GetDataset().ListDatasets(context.TODO(), -1, 0, false)
	t.Log(ok)
	t.Error(err)
	defer ac.Close()
}

func TestCreateDataset(t *testing.T) {
	ac := New(WithStorage())
	ds, _, err := ac.Storage.GetDataset().CreateDataset(context.TODO(), "sitmap")
	t.Log(ds)
	t.Error(err)
}

func TestUpdateDataset(t *testing.T) {
	ac := New(WithStorage())
	ds, _, err := ac.Storage.GetDataset().UpdateDataset(context.TODO(), "test")
	t.Log(ds)
	t.Error(err)
}

func TestDelDataset(t *testing.T) {
	ac := New(WithStorage())
	ds, err := ac.Storage.GetDataset().DelDataset(context.TODO())
	t.Log(ds)
	t.Error(err)
	defer ac.Close()
}

func TestGetItems(t *testing.T) {
	ac := New(WithStorage())
	ds, err := ac.Storage.GetDataset().GetItems(context.TODO(), 1, 100, true)
	t.Log(ds)
	t.Error(err)
}
