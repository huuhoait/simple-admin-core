package MerchantMeta

import (
	"context"

	"github.com/huuhoait/zero-admin-core/rpc/internal/svc"
	"github.com/huuhoait/zero-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateMerchantMetaLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateMerchantMetaLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateMerchantMetaLogic {
	return &CreateMerchantMetaLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// MerchantMeta management
func (l *CreateMerchantMetaLogic) CreateMerchantMeta(in *core.MerchantMetaInfo) (*core.BaseIDResp, error) {
	// todo: add your logic here and delete this line

	return &core.BaseIDResp{}, nil
}
