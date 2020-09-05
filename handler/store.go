package handler

import (
	"context"
	"github.com/micro/go-micro/v2/util/log"
	"github.com/wmsx/store_svc/models"
	proto "github.com/wmsx/store_svc/proto/store"
)

type StoreHandler struct{}

func (h *StoreHandler) GetByStoreIds(ctx context.Context,
	req *proto.GetByStoreIdsRequest, res *proto.GetByStoreIdsResponse) error {
	var (
		objects []*models.Object
		err     error
	)
	if objects, err = models.GetObjectsById(req.StoreIds); err != nil {
		log.Error("根据id查询object失败 err: ", err)
		return err
	}

	storeInfos := make([]*proto.StoreInfo, 0)
	for _, object := range objects {
		storeInfo := &proto.StoreInfo{
			Id:         object.ID,
			BulkName:   object.Bulk,
			ObjectName: object.ObjectName,
		}
		storeInfos = append(storeInfos, storeInfo)
	}
	res.StoreInfos = storeInfos
	return nil
}

func (h *StoreHandler) SaveStoreInfo(ctx context.Context, req *proto.SaveStoreInfoRequest, res *proto.SaveStoreInfoResponse) error {
	storeInfo := req.StoreInfo
	object := &models.Object{
		Bulk:       storeInfo.BulkName,
		ObjectName: storeInfo.ObjectName,
		Filename:   storeInfo.Filename,
		Size:       storeInfo.Size,
	}

	if err := models.AddObject(object); err != nil {
		log.Error("批量添加Object失败 err", err)
		return err
	}

	res.Name = object.Filename
	res.Id = object.ID
	return nil
}
