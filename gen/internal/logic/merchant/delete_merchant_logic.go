package merchant

import (
	"context"

    "github.com/huuhoait/zero-admin-core/internal/svc"
	"github.com/huuhoait/zero-admin-core/internal/types"
	"github.com/suyuan32/simple-admin-example-rpc/example"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteMerchantLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteMerchantLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteMerchantLogic {
	return &DeleteMerchantLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteMerchantLogic) DeleteMerchant(req *types.IDsReq) (resp *types.BaseMsgResp, err error) {
	result, err := l.svcCtx.merchantRpc.DeleteMerchant(l.ctx, &__.IDsReq{
		Ids: req.Ids,
	})
	if err != nil {
		return nil, err
	}

	return &types.BaseMsgResp{Msg: l.svcCtx.Trans.Trans(l.ctx, result.Msg)}, nil
}
