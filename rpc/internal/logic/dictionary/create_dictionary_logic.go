package dictionary

import (
	"context"

	"github.com/huuhoait/simple-admin-common/utils/pointy"

	"github.com/huuhoait/simple-admin-core/rpc/internal/svc"
	"github.com/huuhoait/simple-admin-core/rpc/internal/utils/errorhandler"
	"github.com/huuhoait/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/huuhoait/simple-admin-common/i18n"
)

type CreateDictionaryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateDictionaryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateDictionaryLogic {
	return &CreateDictionaryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateDictionaryLogic) CreateDictionary(in *core.DictionaryInfo) (*core.BaseIDResp, error) {
	result, err := l.svcCtx.DB.Dictionary.Create().
		SetNotNilStatus(pointy.GetStatusPointer(in.Status)).
		SetNotNilTitle(in.Title).
		SetNotNilName(in.Name).
		SetNotNilDesc(in.Desc).
		Save(l.ctx)
	if err != nil {
		return nil, errorhandler.DefaultEntError(l.Logger, err, in)
	}

	return &core.BaseIDResp{Id: result.ID, Msg: i18n.CreateSuccess}, nil
}
