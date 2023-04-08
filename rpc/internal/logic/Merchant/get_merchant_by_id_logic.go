package merchant

import (
	"context"

	"github.com/huuhoait/zero-admin-core/internal/svc"
	"github.com/huuhoait/zero-admin-core/internal/utils/dberrorhandler"
	"github.com/huuhoait/zero-admin-core/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMerchantByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetMerchantByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMerchantByIdLogic {
	return &GetMerchantByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetMerchantByIdLogic) GetMerchantById(in *core.IDReq) (*core.MerchantInfo, error) {
	result, err := l.svcCtx.DB.Merchant.Get(l.ctx, in.Id)
	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	return &core.MerchantInfo{
		Id:          result.ID,
		CreatedAt:   result.CreatedAt.UnixMilli(),
		UpdatedAt:   result.UpdatedAt.UnixMilli(),
			CreatedBy:	result.CreatedBy,
			UpdatedBy:	result.UpdatedBy,
			Status:	uint32(result.Status),
			Sort:	result.Sort,
			Name:	result.Name,
			Leader:	result.Leader,
			Phone:	result.Phone,
			Email:	result.Email,
			Remark:	result.Remark,
			ParentId:	result.ParentID,
	}, nil
}

