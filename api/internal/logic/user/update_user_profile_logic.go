package user

import (
	"context"

	"github.com/huuhoait/zero-admin-core/api/internal/svc"
	"github.com/huuhoait/zero-admin-core/api/internal/types"
	"github.com/huuhoait/zero-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserProfileLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateUserProfileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserProfileLogic {
	return &UpdateUserProfileLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateUserProfileLogic) UpdateUserProfile(req *types.ProfileReq) (resp *types.BaseMsgResp, err error) {
	result, err := l.svcCtx.CoreRpc.UpdateUser(l.ctx, &core.UserInfo{
		Id:        l.ctx.Value("userId").(string),
		Nickname:  req.Nickname,
		Email:     req.Email,
		Mobile:    req.Mobile,
		Avatar:    req.Avatar,
		CreatedBy: l.ctx.Value("userId").(string),
	})
	if err != nil {
		return nil, err
	}

	return &types.BaseMsgResp{Msg: l.svcCtx.Trans.Trans(l.ctx, result.Msg)}, nil
}
