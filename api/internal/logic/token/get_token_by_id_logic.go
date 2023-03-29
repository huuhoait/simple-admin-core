package token

import (
	"context"

	"github.com/huuhoait/zero-admin-core/api/internal/svc"
	"github.com/huuhoait/zero-admin-core/api/internal/types"
	"github.com/huuhoait/zero-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/huuhoait/zero-tools/i18n"
)

type GetTokenByIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetTokenByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTokenByIdLogic {
	return &GetTokenByIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetTokenByIdLogic) GetTokenById(req *types.UUIDReq) (resp *types.TokenInfoResp, err error) {
	data, err := l.svcCtx.CoreRpc.GetTokenById(l.ctx, &core.UUIDReq{Id: req.Id})
	if err != nil {
		return nil, err
	}

	return &types.TokenInfoResp{
		BaseDataInfo: types.BaseDataInfo{
			Code: 0,
			Msg:  l.svcCtx.Trans.Trans(l.ctx, i18n.Success),
		},
		Data: types.TokenInfo{
			BaseUUIDInfo: types.BaseUUIDInfo{
				Id:        data.Id,
				CreatedAt: data.CreatedAt,
				UpdatedAt: data.UpdatedAt,
			},
			Status:    data.Status,
			Uuid:      data.Uuid,
			Token:     data.Token,
			Source:    data.Source,
			ExpiredAt: data.ExpiredAt,
		},
	}, nil
}
