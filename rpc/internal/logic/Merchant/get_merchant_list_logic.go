package merchant

import (
	"context"

	"github.com/huuhoait/zero-admin-core/rpc/internal/svc"
	"github.com/huuhoait/zero-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMerchantListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetMerchantListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMerchantListLogic {
	return &GetMerchantListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetMerchantListLogic) GetMerchantList(in *core.MerchantListReq) (*core.MerchantListResp, error) {
	var predicates []predicate.Merchant
	if in.CreatedBy != "" {
		predicates = append(predicates, merchant.CreatedByContains(in.CreatedBy))
	}
	if in.UpdatedBy != "" {
		predicates = append(predicates, merchant.UpdatedByContains(in.UpdatedBy))
	}
	if in.Name != "" {
		predicates = append(predicates, merchant.NameContains(in.Name))
	}
	result, err := l.svcCtx.DB.Merchant.Query().Where(predicates...).Page(l.ctx, in.Page, in.PageSize)

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	resp := &core.MerchantListResp{}
	resp.Total = result.PageDetails.Total

	for _, v := range result.List {
		resp.Data = append(resp.Data, &core.MerchantInfo{
			Id:        v.ID,
			CreatedAt: v.CreatedAt.UnixMilli(),
			UpdatedAt: v.UpdatedAt.UnixMilli(),
			CreatedBy: v.CreatedBy,
			UpdatedBy: v.UpdatedBy,
			Status:    uint32(v.Status),
			Sort:      v.Sort,
			Name:      v.Name,
			Leader:    v.Leader,
			Phone:     v.Phone,
			Email:     v.Email,
			Remark:    v.Remark,
			ParentId:  v.ParentID,
		})
	}

	return resp, nil
}
