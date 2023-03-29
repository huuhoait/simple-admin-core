package audit

import (
	"context"

	"github.com/huuhoait/zero-admin-core/rpc/internal/svc"
	"github.com/huuhoait/zero-admin-core/rpc/internal/utils/errorhandler"
	"github.com/huuhoait/zero-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAuditByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetAuditByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAuditByIdLogic {
	return &GetAuditByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetAuditByIdLogic) GetAuditById(in *core.IDReq) (*core.AuditInfo, error) {
	result, err := l.svcCtx.DB.Audit.Get(l.ctx, in.Id)
	if err != nil {
		return nil, errorhandler.DefaultEntError(l.Logger, err, in)
	}

	return &core.AuditInfo{
		Id:          result.ID,
		CreatedAt:   result.CreatedAt.UnixMilli(),
		UpdatedAt:   result.UpdatedAt.UnixMilli(),
		ObjectName:  result.ObjectName,
		ActionName:  result.ActionName,
		ChangedData: result.ChangedData,
		CreatedBy:   result.CreatedBy,
	}, nil

}
