package merchantmeta

import (
	"context"

	"github.com/huuhoait/zero-admin-core/api/internal/svc"
	"github.com/huuhoait/zero-admin-core/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteMerchantMetaLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteMerchantMetaLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteMerchantMetaLogic {
	return &DeleteMerchantMetaLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteMerchantMetaLogic) DeleteMerchantMeta(req *types.IDsReq) (resp *types.BaseMsgResp, err error) {
	// todo: add your logic here and delete this line

	return
}
