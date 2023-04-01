package MerchantMeta

import (
	"context"

	"github.com/huuhoait/zero-admin-core/rpc/internal/svc"
	"github.com/huuhoait/zero-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateMerchantMetaLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateMerchantMetaLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateMerchantMetaLogic {
	return &UpdateMerchantMetaLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateMerchantMetaLogic) UpdateMerchantMeta(in *core.MerchantMetaInfo) (*core.BaseResp, error) {
	// todo: add your logic here and delete this line

	return &core.BaseResp{}, nil
}
