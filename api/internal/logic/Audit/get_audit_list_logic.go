package audit

import (
	"context"

	"github.com/huuhoait/zero-admin-core/api/internal/svc"
	"github.com/huuhoait/zero-admin-core/api/internal/types"
	"github.com/huuhoait/zero-admin-core/rpc/types/core"
	"github.com/huuhoait/zero-tools/i18n"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAuditListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetAuditListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAuditListLogic {
	return &GetAuditListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetAuditListLogic) GetAuditList(req *types.AuditListReq) (resp *types.AuditListResp, err error) {
	data, err := l.svcCtx.CoreRpc.GetAuditList(l.ctx,
		&core.AuditListReq{
			Page:     req.Page,
			PageSize: req.PageSize,
			Name:     req.Name,
		})
	if err != nil {
		return nil, err
	}
	resp = &types.AuditListResp{}
	resp.Msg = l.svcCtx.Trans.Trans(l.ctx, i18n.Success)
	resp.Data.Total = data.GetTotal()

	for _, v := range data.Data {
		resp.Data.Data = append(resp.Data.Data,
			types.AuditInfo{
				BaseIDInfo: types.BaseIDInfo{
					Id:        v.Id,
					CreatedAt: v.CreatedAt,
					UpdatedAt: v.UpdatedAt,
				},
				ObjectName:  v.ObjectName,
				ActionName:  v.ActionName,
				ChangedData: v.ChangedData,
				CreatedBy:   v.CreatedBy,
			})
	}
	return resp, nil
}
