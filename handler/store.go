package handler

import (
	"context"
	"github.com/wmsx/store_svc/models"
	proto "github.com/wmsx/store_svc/proto/store"
)

type StoreHandler struct{}

func (s *StoreHandler) SaveStoreInfo(ctx context.Context, req *proto.SaveStoreInfoRequest, res *proto.SaveStoreInfoResponse) error {
	var (
		objects []models.Object
	)
	if len(req.StoreInfos) == 0 {
		return nil
	}
	for _, storeInfo := range req.StoreInfos {
		object := models.Object{
			Bulk:       storeInfo.BulkName,
			ObjectName: storeInfo.ObjectName,
			Filename:   storeInfo.Filename,
			Size:       storeInfo.Size,
			Status:     0,
		}
		objects = append(objects, object)
	}

	if err := models.BatchAddObject(objects); err != nil {
		return err
	}
	name2IdMap := make(map[string]int64)
	for _, o := range objects {
		name2IdMap[o.Filename] = int64(o.ID)
	}
	res.Name2IdMap = name2IdMap
	return nil
}
