package token

import (
	"context"

	"github.com/huuhoait/zero-tools/utils/uuidx"

	"github.com/huuhoait/zero-admin-core/rpc/internal/svc"
	"github.com/huuhoait/zero-admin-core/rpc/internal/utils/errorhandler"
	"github.com/huuhoait/zero-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetTokenByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetTokenByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTokenByIdLogic {
	return &GetTokenByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetTokenByIdLogic) GetTokenById(in *core.UUIDReq) (*core.TokenInfo, error) {
	result, err := l.svcCtx.DB.Token.Get(l.ctx, uuidx.ParseUUIDString(in.Id))
	if err != nil {
		return nil, errorhandler.DefaultEntError(l.Logger, err, in)
	}

	return &core.TokenInfo{
		Id:        result.ID.String(),
		CreatedAt: result.CreatedAt.UnixMilli(),
		UpdatedAt: result.UpdatedAt.UnixMilli(),
		Status:    uint32(result.Status),
		Uuid:      result.UUID.String(),
		Token:     result.Token,
		Source:    result.Source,
		ExpiredAt: result.ExpiredAt.UnixMilli(),
	}, nil
}
