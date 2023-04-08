package merchant

import (
	"context"


	"github.com/huuhoait/zero-admin-core/internal/svc"
	"github.com/huuhoait/zero-admin-core/internal/utils/dberrorhandler"
    "github.com/huuhoait/zero-admin-core/core"

    "github.com/suyuan32/simple-admin-common/i18n"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateMerchantLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateMerchantLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateMerchantLogic {
	return &CreateMerchantLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateMerchantLogic) CreateMerchant(in *core.MerchantInfo) (*core.BaseIDResp, error) {
    result, err := l.svcCtx.DB.Merchant.Create().
			SetCreatedBy(in.CreatedBy).
			SetUpdatedBy(in.UpdatedBy).
			SetStatus(uint8(in.Status)).
			SetSort(in.Sort).
			SetName(in.Name).
			SetLeader(in.Leader).
			SetPhone(in.Phone).
			SetEmail(in.Email).
			SetRemark(in.Remark).
			SetParentID(in.ParentId).
			Save(l.ctx)

    if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

    return &core.BaseIDResp{Id: result.ID, Msg: i18n.CreateSuccess}, nil
}
