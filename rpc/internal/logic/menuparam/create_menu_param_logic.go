package menuparam

import (
	"context"

	"github.com/huuhoait/zero-admin-core/rpc/internal/svc"
	"github.com/huuhoait/zero-admin-core/rpc/internal/utils/errorhandler"
	"github.com/huuhoait/zero-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/huuhoait/zero-tools/i18n"
)

type CreateMenuParamLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateMenuParamLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateMenuParamLogic {
	return &CreateMenuParamLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateMenuParamLogic) CreateMenuParam(in *core.MenuParamInfo) (*core.BaseIDResp, error) {
	result, err := l.svcCtx.DB.MenuParam.Create().
		SetType(in.Type).
		SetKey(in.Key).
		SetValue(in.Value).
		SetMenuID(in.MenuId).
		Save(l.ctx)
	if err != nil {
		return nil, errorhandler.DefaultEntError(l.Logger, err, in)
	}

	return &core.BaseIDResp{Id: result.ID, Msg: i18n.CreateSuccess}, nil
}
