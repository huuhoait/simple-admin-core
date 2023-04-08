package merchant

import (
	"context"

    "github.com/huuhoait/zero-admin-core/ent/merchant"
    "github.com/huuhoait/zero-admin-core/internal/svc"
    "github.com/huuhoait/zero-admin-core/internal/utils/dberrorhandler"
    "github.com/huuhoait/zero-admin-core/core"

    "github.com/suyuan32/simple-admin-common/i18n"
    "github.com/zeromicro/go-zero/core/logx"
)

type DeleteMerchantLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteMerchantLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteMerchantLogic {
	return &DeleteMerchantLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteMerchantLogic) DeleteMerchant(in *core.IDsReq) (*core.BaseResp, error) {
	_, err := l.svcCtx.DB.Merchant.Delete().Where(merchant.IDIn(in.Ids...)).Exec(l.ctx)

    if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

    return &core.BaseResp{Msg: i18n.DeleteSuccess}, nil
}
