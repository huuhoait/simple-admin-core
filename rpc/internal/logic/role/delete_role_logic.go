package role

import (
	"context"

	"github.com/zeromicro/go-zero/core/errorx"

	"github.com/huuhoait/zero-admin-core/rpc/ent/role"
	"github.com/huuhoait/zero-admin-core/rpc/ent/user"

	"github.com/huuhoait/zero-admin-core/rpc/internal/svc"
	"github.com/huuhoait/zero-admin-core/rpc/internal/utils/errorhandler"
	"github.com/huuhoait/zero-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/huuhoait/zero-tools/i18n"
)

type DeleteRoleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteRoleLogic {
	return &DeleteRoleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteRoleLogic) DeleteRole(in *core.IDsReq) (*core.BaseResp, error) {
	count, err := l.svcCtx.DB.User.Query().Where(user.HasRolesWith(role.IDIn(in.Ids...))).Count(l.ctx)
	if err != nil {
		return nil, errorhandler.DefaultEntError(l.Logger, err, in)
	}

	if count != 0 {
		return nil, errorx.NewInvalidArgumentError("role.userExists")
	}

	_, err = l.svcCtx.DB.Role.Delete().Where(role.IDIn(in.Ids...)).Exec(l.ctx)
	if err != nil {
		return nil, errorhandler.DefaultEntError(l.Logger, err, in)
	}

	return &core.BaseResp{Msg: i18n.DeleteSuccess}, nil
}
