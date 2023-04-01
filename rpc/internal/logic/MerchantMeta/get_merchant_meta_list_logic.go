package MerchantMeta

import (
	"context"

	"github.com/huuhoait/zero-admin-core/rpc/internal/svc"
	"github.com/huuhoait/zero-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMerchantMetaListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetMerchantMetaListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMerchantMetaListLogic {
	return &GetMerchantMetaListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetMerchantMetaListLogic) GetMerchantMetaList(in *core.MerchantMetaListReq) (*core.MerchantMetaListResp, error) {
	// todo: add your logic here and delete this line

	return &core.MerchantMetaListResp{}, nil
}
