package menuparam

import (
	"context"

	"github.com/huuhoait/zero-admin-core/rpc/ent/menuparam"

	"github.com/huuhoait/zero-admin-core/rpc/internal/svc"
	"github.com/huuhoait/zero-admin-core/rpc/internal/utils/errorhandler"
	"github.com/huuhoait/zero-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/huuhoait/zero-tools/i18n"
)

type DeleteMenuParamLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteMenuParamLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteMenuParamLogic {
	return &DeleteMenuParamLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteMenuParamLogic) DeleteMenuParam(in *core.IDsReq) (*core.BaseResp, error) {
	_, err := l.svcCtx.DB.MenuParam.Delete().Where(menuparam.IDIn(in.Ids...)).Exec(l.ctx)
	if err != nil {
		return nil, errorhandler.DefaultEntError(l.Logger, err, in)
	}

	return &core.BaseResp{Msg: i18n.DeleteSuccess}, nil
}
