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

type GetAuditByIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	lang   string
}

func NewGetAuditByIdLogic(r *http.Request, svcCtx *svc.ServiceContext) *GetAuditByIdLogic {
	return &GetAuditByIdLogic{
		Logger: logx.WithContext(r.Context()),
		ctx:    r.Context(),
		svcCtx: svcCtx,
		lang:   r.Header.Get("Accept-Language"),
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
			BaseInfo: types.BaseInfo{
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
