package merchantmeta

import (
	"context"

	"github.com/huuhoait/zero-admin-core/ent/merchantmeta"
	"github.com/huuhoait/zero-admin-core/ent/predicate"
	"github.com/huuhoait/zero-admin-core/internal/svc"
	"github.com/huuhoait/zero-admin-core/internal/utils/dberrorhandler"
    "github.com/huuhoait/zero-admin-core/core"

    "github.com/zeromicro/go-zero/core/logx"
)

type GetMerchantMetaListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetMerchantMetaListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMerchantMetaListLogic {
	return &GetMerchantMetaListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetMerchantMetaListLogic) GetMerchantMetaList(in *core.MerchantMetaListReq) (*core.MerchantMetaListResp, error) {
	var predicates []predicate.MerchantMeta
	if in.CreatedBy != "" {
		predicates = append(predicates, merchantmeta.CreatedByContains(in.CreatedBy))
	}
	if in.UpdatedBy != "" {
		predicates = append(predicates, merchantmeta.UpdatedByContains(in.UpdatedBy))
	}
	if in.Title != "" {
		predicates = append(predicates, merchantmeta.TitleContains(in.Title))
	}
	result, err := l.svcCtx.DB.MerchantMeta.Query().Where(predicates...).Page(l.ctx, in.Page, in.PageSize)

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	resp := &core.MerchantMetaListResp{}
	resp.Total = result.PageDetails.Total

	for _, v := range result.List {
		resp.Data = append(resp.Data, &core.MerchantMetaInfo{
			Id:          v.ID,
			CreatedAt:   v.CreatedAt.UnixMilli(),
			UpdatedAt:   v.UpdatedAt.UnixMilli(),
			CreatedBy:	v.CreatedBy,
			UpdatedBy:	v.UpdatedBy,
			Status:	uint32(v.Status),
			Sort:	v.Sort,
			Title:	v.Title,
			Key:	v.Key,
			Value:	v.Value,
			MerchantId:	v.MerchantID,
		})
	}

	return resp, nil
}
