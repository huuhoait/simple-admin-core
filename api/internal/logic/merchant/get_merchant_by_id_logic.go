package merchant

import (
	"context"

	"github.com/huuhoait/zero-admin-core/api/internal/svc"
	"github.com/huuhoait/zero-admin-core/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMerchantByIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetMerchantByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMerchantByIdLogic {
	return &GetMerchantByIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetMerchantByIdLogic) GetMerchantById(req *types.IDReq) (resp *types.MerchantInfoResp, err error) {
	// todo: add your logic here and delete this line

	return
}
