package merchant

import (
	"context"

	"github.com/huuhoait/zero-admin-core/rpc/internal/svc"
	"github.com/huuhoait/zero-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateMerchantLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateMerchantLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateMerchantLogic {
	return &UpdateMerchantLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateMerchantLogic) UpdateMerchant(in *core.MerchantInfo) (*core.BaseResp, error) {
	err := l.svcCtx.DB.Merchant.UpdateOneID(in.Id).
		SetNotEmptyCreatedBy(in.CreatedBy).
		SetNotEmptyUpdatedBy(in.UpdatedBy).
		SetNotEmptyStatus(uint8(in.Status)).
		SetNotEmptySort(in.Sort).
		SetNotEmptyName(in.Name).
		SetNotEmptyLeader(in.Leader).
		SetNotEmptyPhone(in.Phone).
		SetNotEmptyEmail(in.Email).
		SetNotEmptyRemark(in.Remark).
		SetNotEmptyParentID(in.ParentId).
		Exec(l.ctx)

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	return &core.BaseResp{Msg: i18n.UpdateSuccess}, nil
}
