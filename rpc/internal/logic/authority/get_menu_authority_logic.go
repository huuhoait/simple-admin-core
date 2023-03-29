package authority

import (
	"context"

	"github.com/huuhoait/zero-admin-core/rpc/ent/role"

	"github.com/huuhoait/zero-admin-core/rpc/internal/svc"
	"github.com/huuhoait/zero-admin-core/rpc/internal/utils/errorhandler"
	"github.com/huuhoait/zero-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMenuAuthorityLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetMenuAuthorityLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMenuAuthorityLogic {
	return &GetMenuAuthorityLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetMenuAuthorityLogic) GetMenuAuthority(in *core.IDReq) (*core.RoleMenuAuthorityResp, error) {
	menus, err := l.svcCtx.DB.Role.Query().Where(role.ID(in.Id)).QueryMenus().IDs(l.ctx)
	if err != nil {
		return nil, errorhandler.DefaultEntError(l.Logger, err, in)
	}

	return &core.RoleMenuAuthorityResp{MenuId: menus}, nil
}
