package user

import (
	"context"

	"github.com/huuhoait/zero-tools/i18n"

	"github.com/huuhoait/zero-admin-core/api/internal/svc"
	"github.com/huuhoait/zero-admin-core/api/internal/types"
	"github.com/huuhoait/zero-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserInfoLogic) GetUserInfo() (resp *types.UserBaseIDInfoResp, err error) {
	user, err := l.svcCtx.CoreRpc.GetUserById(l.ctx,
		&core.UUIDReq{Id: l.ctx.Value("userId").(string)})
	if err != nil {
		return nil, err
	}

	return &types.UserBaseIDInfoResp{
		BaseDataInfo: types.BaseDataInfo{Msg: l.svcCtx.Trans.Trans(l.ctx, i18n.Success)},
		Data: types.UserBaseIDInfo{
			UUID:        user.Id,
			Username:    user.Username,
			Nickname:    user.Nickname,
			Avatar:      user.Avatar,
			HomePath:    user.HomePath,
			Description: user.Description,
		},
	}, nil
}
