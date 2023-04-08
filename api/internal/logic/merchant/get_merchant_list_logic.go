package merchant

import (
	"context"

	"github.com/huuhoait/zero-admin-core/api/internal/svc"
	"github.com/huuhoait/zero-admin-core/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMerchantListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetMerchantListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMerchantListLogic {
	return &GetMerchantListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetMerchantListLogic) GetMerchantList(req *types.MerchantListReq) (resp *types.MerchantListResp, err error) {
	// todo: add your logic here and delete this line

	return
}
