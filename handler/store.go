package handler

import (
	"context"
	"github.com/micro/go-micro/v2/util/log"
	"github.com/wmsx/store_svc/models"
	proto "github.com/wmsx/store_svc/proto/store"
)

type StoreHandler struct{}

func (h *StoreHandler) GetByObjectIds(ctx  context.Context,
	req *proto.GetByObjectIdsRequest, res *proto.GetByObjectIdsResponse) error {
	var (
		objects []*models.Object
		err error
	)
	if objects, err = models.GetObjectsById(req.ObjectIds); err != nil {
		log.Error("根据id查询object失败 err: ", err)
		return err
	}

	objectInfos := make([]*proto.ObjectInfo, 0)
	for _, object := range objects{
		objectInfo := &proto.ObjectInfo{
			Id:         object.ID,
			Bulk:       object.Bulk,
			ObjectName: object.ObjectName,
		}
		objectInfos = append(objectInfos, objectInfo)
	}
	res.ObjectInfos = objectInfos
	return nil
}

func (h *StoreHandler) SaveStoreInfo(ctx context.Context, req *proto.SaveStoreInfoRequest, res *proto.SaveStoreInfoResponse) error {
	var (
		objects []*models.Object
	)
	if len(req.StoreInfos) == 0 {
		return nil
	}
	for _, storeInfo := range req.StoreInfos {
		object := &models.Object{
			Bulk:       storeInfo.BulkName,
			ObjectName: storeInfo.ObjectName,
			Filename:   storeInfo.Filename,
			Size:       storeInfo.Size,
			Status:     1,
		}
		objects = append(objects, object)
	}

	if err := models.BatchAddObject(objects); err != nil {
		log.Error("批量添加Object失败 err", err)
		return err
	}
	name2IdMap := make(map[string]int64)
	for _, o := range objects {
		name2IdMap[o.Filename] = o.ID
	}
	res.Name2IdMap = name2IdMap
	return nil
}
