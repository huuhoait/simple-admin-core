package token

import (
	"context"
	"time"

	"github.com/huuhoait/simple-admin-common/enum/common"
	"github.com/huuhoait/simple-admin-common/i18n"
	"github.com/huuhoait/simple-admin-common/msg/logmsg"
	"github.com/huuhoait/simple-admin-common/utils/pointy"
	"github.com/huuhoait/simple-admin-common/utils/uuidx"
	"github.com/zeromicro/go-zero/core/errorx"

	"github.com/huuhoait/simple-admin-core/rpc/internal/svc"
	"github.com/huuhoait/simple-admin-core/rpc/internal/utils/errorhandler"
	"github.com/huuhoait/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateTokenLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateTokenLogic {
	return &UpdateTokenLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateTokenLogic) UpdateToken(in *core.TokenInfo) (*core.BaseResp, error) {
	token, err := l.svcCtx.DB.Token.UpdateOneID(uuidx.ParseUUIDString(*in.Id)).
		SetNotNilStatus(pointy.GetStatusPointer(in.Status)).
		SetNotNilSource(in.Source).
		Save(l.ctx)
	if err != nil {
		return nil, errorhandler.DefaultEntError(l.Logger, err, in)
	}

	if uint8(*in.Status) == common.StatusBanned {
		expiredTime := int(token.ExpiredAt.Unix() - time.Now().Unix())
		if expiredTime > 0 {
			err = l.svcCtx.Redis.Setex("token_"+token.Token, "1", expiredTime)
			if err != nil {
				logx.Errorw(logmsg.RedisError, logx.Field("detail", err.Error()))
				return nil, errorx.NewInternalError(i18n.RedisError)
			}
		}
	} else if uint8(*in.Status) == common.StatusNormal {
		_, err := l.svcCtx.Redis.Del("token_" + token.Token)
		if err != nil {
			logx.Errorw(logmsg.RedisError, logx.Field("detail", err.Error()))
			return nil, errorx.NewInternalError(i18n.RedisError)
		}
	}

	return &core.BaseResp{Msg: i18n.UpdateSuccess}, nil
}
