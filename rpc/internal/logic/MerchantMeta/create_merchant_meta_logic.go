package merchantmeta

import (
	"context"

	"github.com/huuhoait/zero-admin-core/rpc/internal/svc"
	"github.com/huuhoait/zero-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateMerchantMetaLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateMerchantMetaLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateMerchantMetaLogic {
	return &CreateMerchantMetaLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateMerchantMetaLogic) CreateMerchantMeta(in *core.MerchantMetaInfo) (*core.BaseIDResp, error) {
	result, err := l.svcCtx.DB.MerchantMeta.Create().
		SetCreatedBy(in.CreatedBy).
		SetUpdatedBy(in.UpdatedBy).
		SetStatus(uint8(in.Status)).
		SetSort(in.Sort).
		SetTitle(in.Title).
		SetKey(in.Key).
		SetValue(in.Value).
		SetMerchantID(in.MerchantId).
		Save(l.ctx)

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	return &core.BaseIDResp{Id: result.ID, Msg: i18n.CreateSuccess}, nil
}
