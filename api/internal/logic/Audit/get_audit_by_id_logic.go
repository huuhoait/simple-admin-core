package audit

import (
	"context"

	"github.com/huuhoait/zero-admin-core/api/internal/svc"
	"github.com/huuhoait/zero-admin-core/api/internal/types"
	"github.com/huuhoait/zero-admin-core/rpc/types/core"
	"github.com/suyuan32/simple-admin-common/i18n"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAuditByIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetAuditByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAuditByIdLogic {
	return &GetAuditByIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetAuditByIdLogic) GetAuditById(req *types.IDReq) (resp *types.AuditInfoResp, err error) {
	data, err := l.svcCtx.CoreRpc.GetAuditById(l.ctx, &core.IDReq{Id: req.Id})
	if err != nil {
		return nil, err
	}
	var currentUser = l.ctx.Value("userId").(string)

	return &types.AuditInfoResp{
		BaseDataInfo: types.BaseDataInfo{
			Code: 0,
			Msg:  l.svcCtx.Trans.Trans(l.ctx, i18n.Success),
		},
		Data: types.AuditInfo{
			BaseIDInfo: types.BaseIDInfo{
				Id:        data.Id,
				CreatedAt: data.CreatedAt,
				UpdatedAt: data.UpdatedAt,
			},
			ObjectName:  data.ObjectName,
			ActionName:  data.ActionName,
			ChangedData: data.ChangedData,
			CreatedBy:   currentUser,
		},
	}, nil
}
