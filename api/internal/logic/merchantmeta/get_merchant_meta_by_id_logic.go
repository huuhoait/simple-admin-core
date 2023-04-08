package merchantmeta

import (
	"context"

	"github.com/huuhoait/zero-admin-core/api/internal/svc"
	"github.com/huuhoait/zero-admin-core/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMerchantMetaByIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetMerchantMetaByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMerchantMetaByIdLogic {
	return &GetMerchantMetaByIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetMerchantMetaByIdLogic) GetMerchantMetaById(req *types.IDReq) (resp *types.MerchantMetaInfoResp, err error) {
	// todo: add your logic here and delete this line

	return
}
