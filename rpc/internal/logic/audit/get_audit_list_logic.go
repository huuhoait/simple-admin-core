package audit

import (
	"context"

	"github.com/huuhoait/zero-admin-core/rpc/ent"
	"github.com/huuhoait/zero-admin-core/rpc/ent/audit"
	"github.com/huuhoait/zero-admin-core/rpc/ent/predicate"
	"github.com/huuhoait/zero-admin-core/rpc/internal/svc"
	"github.com/huuhoait/zero-admin-core/rpc/internal/utils/errorhandler"
	"github.com/huuhoait/zero-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAuditListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetAuditListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAuditListLogic {
	return &GetAuditListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetAuditListLogic) GetAuditList(in *core.AuditListReq) (*core.AuditListResp, error) {

	var predicates []predicate.Audit
	if in.Name != "" {
		predicates = append(predicates, audit.ObjectName(in.Name))
	}

	result, err := l.svcCtx.DB.Audit.Query().Where(predicates...).Page(l.ctx, in.Page, in.PageSize, func(pager *ent.AuditPager) {
		pager.Order = ent.Asc(audit.FieldID)
	})
	if err != nil {
		return nil, errorhandler.DefaultEntError(l.Logger, err, in)
	}

	resp := &core.AuditListResp{}
	resp.Total = result.PageDetails.Total

	for _, v := range result.List {
		resp.Data = append(resp.Data, &core.AuditInfo{
			Id:          v.ID,
			CreatedAt:   v.CreatedAt.UnixMilli(),
			UpdatedAt:   v.UpdatedAt.UnixMilli(),
			ObjectName:  v.ObjectName,
			ActionName:  v.ActionName,
			ChangedData: v.ChangedData,
			CreatedBy:   v.CreatedBy,
		})
	}

	return resp, nil
}
