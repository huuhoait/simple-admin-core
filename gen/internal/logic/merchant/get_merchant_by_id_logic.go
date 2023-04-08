package merchant

import (
    "context"

    "github.com/huuhoait/zero-admin-core/internal/svc"
    "github.com/huuhoait/zero-admin-core/internal/types"
    "github.com/suyuan32/simple-admin-example-rpc/example"

    "github.com/suyuan32/simple-admin-common/i18n"
    "github.com/zeromicro/go-zero/core/logx"
)

type GetMerchantByIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetMerchantByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMerchantByIdLogic {
	return &GetMerchantByIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetMerchantByIdLogic) GetMerchantById(req *types.IDReq) (resp *types.MerchantInfoResp, err error) {
	data, err := l.svcCtx.merchantRpc.GetMerchantById(l.ctx, &__.IDReq{Id: req.Id})
	if err != nil {
		return nil, err
	}

	return &types.MerchantInfoResp{
		BaseDataInfo: types.BaseDataInfo{
			Code: 0,
			Msg:  l.svcCtx.Trans.Trans(l.ctx, i18n.Success),
		},
		Data: types.MerchantInfo{
            BaseInfo: types.BaseInfo{
                Id:        data.Id,
                CreatedAt: data.CreatedAt,
                UpdatedAt: data.UpdatedAt,
            },
        	Status: data.Status,
        	Sort: data.Sort,
        	Name: data.Name,
        	Leader: data.Leader,
        	Phone: data.Phone,
        	Email: data.Email,
        	Remark: data.Remark,
        	ParentId: data.ParentId,
        	CreatedBy: data.CreatedBy,
        	UpdatedBy: data.UpdatedBy,
		},
	}, nil
}

