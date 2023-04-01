package MerchantMeta

import (
	"context"

	"github.com/huuhoait/zero-admin-core/rpc/internal/svc"
	"github.com/huuhoait/zero-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteMerchantMetaLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteMerchantMetaLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteMerchantMetaLogic {
	return &DeleteMerchantMetaLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteMerchantMetaLogic) DeleteMerchantMeta(in *core.IDsReq) (*core.BaseResp, error) {
	// todo: add your logic here and delete this line

	return &core.BaseResp{}, nil
}
