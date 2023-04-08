package merchant

import (
	"context"

	"github.com/huuhoait/zero-admin-core/internal/svc"
	"github.com/huuhoait/zero-admin-core/internal/types"
	"github.com/suyuan32/simple-admin-example-rpc/example"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateMerchantLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateMerchantLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateMerchantLogic {
	return &CreateMerchantLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateMerchantLogic) CreateMerchant(req *types.MerchantInfo) (resp *types.BaseMsgResp, err error) {
	data, err := l.svcCtx.merchantRpc.CreateMerchant(l.ctx,
		&__.MerchantInfo{
			Id:          req.Id,
        	Status: req.Status,
        	Sort: req.Sort,
        	Name: req.Name,
        	Leader: req.Leader,
        	Phone: req.Phone,
        	Email: req.Email,
        	Remark: req.Remark,
        	ParentId: req.ParentId,
        	CreatedBy: req.CreatedBy,
        	UpdatedBy: req.UpdatedBy,
		})
	if err != nil {
		return nil, err
	}
	return &types.BaseMsgResp{Msg: l.svcCtx.Trans.Trans(l.ctx, data.Msg)}, nil
}
