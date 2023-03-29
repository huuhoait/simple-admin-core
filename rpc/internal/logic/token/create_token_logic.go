package token

import (
	"context"
	"time"

	"github.com/huuhoait/zero-tools/utils/uuidx"

	"github.com/huuhoait/zero-admin-core/rpc/internal/svc"
	"github.com/huuhoait/zero-admin-core/rpc/internal/utils/errorhandler"
	"github.com/huuhoait/zero-admin-core/rpc/types/core"

	"github.com/huuhoait/zero-tools/i18n"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateTokenLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateTokenLogic {
	return &CreateTokenLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateTokenLogic) CreateToken(in *core.TokenInfo) (*core.BaseUUIDResp, error) {
	result, err := l.svcCtx.DB.Token.Create().
		SetStatus(uint8(in.Status)).
		SetUUID(uuidx.ParseUUIDString(in.Uuid)).
		SetToken(in.Token).
		SetSource(in.Source).
		SetExpiredAt(time.Unix(in.ExpiredAt, 0)).
		Save(l.ctx)
	if err != nil {
		return nil, errorhandler.DefaultEntError(l.Logger, err, in)
	}

	return &core.BaseUUIDResp{Id: result.ID.String(), Msg: i18n.CreateSuccess}, nil
}
