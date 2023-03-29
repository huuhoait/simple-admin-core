package user

import (
	"context"

	"github.com/huuhoait/zero-tools/utils/uuidx"

	"github.com/huuhoait/zero-tools/i18n"

	"github.com/huuhoait/zero-admin-core/rpc/ent/user"

	"github.com/huuhoait/zero-admin-core/rpc/internal/svc"
	"github.com/huuhoait/zero-admin-core/rpc/internal/utils/errorhandler"
	"github.com/huuhoait/zero-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteUserLogic {
	return &DeleteUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteUserLogic) DeleteUser(in *core.UUIDsReq) (*core.BaseResp, error) {
	_, err := l.svcCtx.DB.User.Delete().Where(user.IDIn(uuidx.ParseUUIDSlice(in.Ids)...)).Exec(l.ctx)
	if err != nil {
		return nil, errorhandler.DefaultEntError(l.Logger, err, in)
	}
	return &core.BaseResp{Msg: i18n.DeleteSuccess}, nil
}
