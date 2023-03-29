package dictionary

import (
	"context"

	"github.com/huuhoait/zero-admin-core/rpc/internal/svc"
	"github.com/huuhoait/zero-admin-core/rpc/internal/utils/errorhandler"
	"github.com/huuhoait/zero-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/huuhoait/zero-tools/i18n"
)

type UpdateDictionaryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateDictionaryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateDictionaryLogic {
	return &UpdateDictionaryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateDictionaryLogic) UpdateDictionary(in *core.DictionaryInfo) (*core.BaseResp, error) {
	err := l.svcCtx.DB.Dictionary.UpdateOneID(in.Id).
		SetNotEmptyStatus(uint8(in.Status)).
		SetNotEmptyTitle(in.Title).
		SetNotEmptyName(in.Name).
		SetNotEmptyDesc(in.Desc).
		Exec(l.ctx)
	if err != nil {
		return nil, errorhandler.DefaultEntError(l.Logger, err, in)
	}

	return &core.BaseResp{Msg: i18n.UpdateSuccess}, nil
}
