package MerchantMeta

import (
	"context"

	"github.com/huuhoait/zero-admin-core/rpc/internal/svc"
	"github.com/huuhoait/zero-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMerchantMetaByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetMerchantMetaByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMerchantMetaByIdLogic {
	return &GetMerchantMetaByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetMerchantMetaByIdLogic) GetMerchantMetaById(in *core.IDReq) (*core.MerchantMetaInfo, error) {
	// todo: add your logic here and delete this line

	return &core.MerchantMetaInfo{}, nil
}
