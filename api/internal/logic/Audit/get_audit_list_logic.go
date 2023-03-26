package Audit

import (
	"context"
	"net/http"

	"github.com/suyuan32/simple-admin-common/i18n"
	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAuditListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	lang   string
}

func NewGetAuditListLogic(r *http.Request, svcCtx *svc.ServiceContext) *GetAuditListLogic {
	return &GetAuditListLogic{
		Logger: logx.WithContext(r.Context()),
		ctx:    r.Context(),
		svcCtx: svcCtx,
		lang:   r.Header.Get("Accept-Language"),
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
				BaseInfo: types.BaseInfo{
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
