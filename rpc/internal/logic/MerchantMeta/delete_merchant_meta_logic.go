package merchantmeta

import (
	"context"

	"github.com/huuhoait/zero-admin-core/core"
	"github.com/huuhoait/zero-admin-core/internal/svc"
	"github.com/huuhoait/zero-admin-core/internal/utils/dberrorhandler"
	"github.com/huuhoait/zero-admin-core/rpc/ent/merchantmeta"

	"github.com/suyuan32/simple-admin-common/i18n"
	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteMerchantMetaLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteMerchantMetaLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteMerchantMetaLogic {
	return &DeleteMerchantMetaLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteMerchantMetaLogic) DeleteMerchantMeta(in *core.IDsReq) (*core.BaseResp, error) {
	_, err := l.svcCtx.DB.MerchantMeta.Delete().Where(merchantmeta.IDIn(in.Ids...)).Exec(l.ctx)

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	return &core.BaseResp{Msg: i18n.DeleteSuccess}, nil
}
