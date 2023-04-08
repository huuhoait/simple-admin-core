package merchantmeta

import (
	"context"


	"github.com/huuhoait/zero-admin-core/internal/svc"
	"github.com/huuhoait/zero-admin-core/internal/utils/dberrorhandler"
    "github.com/huuhoait/zero-admin-core/core"

    "github.com/suyuan32/simple-admin-common/i18n"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateMerchantMetaLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateMerchantMetaLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateMerchantMetaLogic {
	return &UpdateMerchantMetaLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateMerchantMetaLogic) UpdateMerchantMeta(in *core.MerchantMetaInfo) (*core.BaseResp, error) {
    err := l.svcCtx.DB.MerchantMeta.UpdateOneID(in.Id).
			SetNotEmptyCreatedBy(in.CreatedBy).
			SetNotEmptyUpdatedBy(in.UpdatedBy).
			SetNotEmptyStatus(uint8(in.Status)).
			SetNotEmptySort(in.Sort).
			SetNotEmptyTitle(in.Title).
			SetNotEmptyKey(in.Key).
			SetNotEmptyValue(in.Value).
			SetNotEmptyMerchantID(in.MerchantId).
			Exec(l.ctx)

    if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

    return &core.BaseResp{Msg: i18n.UpdateSuccess}, nil
}
