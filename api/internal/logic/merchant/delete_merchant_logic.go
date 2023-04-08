package merchant

import (
	"context"

	"github.com/huuhoait/zero-admin-core/api/internal/svc"
	"github.com/huuhoait/zero-admin-core/api/internal/types"

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
	// todo: add your logic here and delete this line

	return
}
