package logic

import (
	"context"

	"github.com/huuhoait/zero-admin-core/rpc/internal/svc"
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
	// todo: add your logic here and delete this line

	return &core.AuditListResp{}, nil
}
