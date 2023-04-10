package merchantmeta

import (
	"context"

	"github.com/huuhoait/zero-admin-core/rpc/internal/svc"
	"github.com/huuhoait/zero-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMerchantMetaByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetMerchantMetaByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMerchantMetaByIdLogic {
	return &GetMerchantMetaByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetMerchantMetaByIdLogic) GetMerchantMetaById(in *core.IDReq) (*core.MerchantMetaInfo, error) {
	result, err := l.svcCtx.DB.MerchantMeta.Get(l.ctx, in.Id)
	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	return &core.MerchantMetaInfo{
		Id:         result.ID,
		CreatedAt:  result.CreatedAt.UnixMilli(),
		UpdatedAt:  result.UpdatedAt.UnixMilli(),
		CreatedBy:  result.CreatedBy,
		UpdatedBy:  result.UpdatedBy,
		Status:     uint32(result.Status),
		Sort:       result.Sort,
		Title:      result.Title,
		Key:        result.Key,
		Value:      result.Value,
		MerchantId: result.MerchantID,
	}, nil
}
