package merchant

import (
	"context"

    "github.com/huuhoait/zero-admin-core/internal/svc"
	"github.com/huuhoait/zero-admin-core/internal/types"
	"github.com/suyuan32/simple-admin-example-rpc/example"

	"github.com/suyuan32/simple-admin-common/i18n"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetMerchantListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetMerchantListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMerchantListLogic {
	return &GetMerchantListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetMerchantListLogic) GetMerchantList(req *types.MerchantListReq) (resp *types.MerchantListResp, err error) {
	data, err := l.svcCtx.merchantRpc.GetMerchantList(l.ctx,
		&__.MerchantListReq{
			Page:        req.Page,
			PageSize:    req.PageSize,
			MerchantId: req.MerchantId,
			Key: req.Key,
		})
	if err != nil {
		return nil, err
	}
	resp = &types.MerchantListResp{}
	resp.Msg = l.svcCtx.Trans.Trans(l.ctx, i18n.Success)
	resp.Data.Total = data.GetTotal()

	for _, v := range data.Data {
		resp.Data.Data = append(resp.Data.Data,
			types.MerchantInfo{
				BaseInfo: types.BaseInfo{
					Id:        v.Id,
					CreatedAt: v.CreatedAt,
					UpdatedAt: v.UpdatedAt,
				},
        	Status: v.Status,
        	Sort: v.Sort,
        	Name: v.Name,
        	Leader: v.Leader,
        	Phone: v.Phone,
        	Email: v.Email,
        	Remark: v.Remark,
        	ParentId: v.ParentId,
        	CreatedBy: v.CreatedBy,
        	UpdatedBy: v.UpdatedBy,
			})
	}
	return resp, nil
}
