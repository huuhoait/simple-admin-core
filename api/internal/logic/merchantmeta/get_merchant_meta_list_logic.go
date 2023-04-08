package merchantmeta

import (
	"context"

	"github.com/huuhoait/zero-admin-core/api/internal/svc"
	"github.com/huuhoait/zero-admin-core/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMerchantMetaListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetMerchantMetaListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMerchantMetaListLogic {
	return &GetMerchantMetaListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetMerchantMetaListLogic) GetMerchantMetaList(req *types.MerchantMetaListReq) (resp *types.MerchantMetaListResp, err error) {
	// todo: add your logic here and delete this line

	return
}
