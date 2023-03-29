package menuparam

import (
	"context"

	"github.com/huuhoait/zero-admin-core/rpc/internal/svc"
	"github.com/huuhoait/zero-admin-core/rpc/internal/utils/errorhandler"
	"github.com/huuhoait/zero-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMenuParamByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetMenuParamByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMenuParamByIdLogic {
	return &GetMenuParamByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetMenuParamByIdLogic) GetMenuParamById(in *core.IDReq) (*core.MenuParamInfo, error) {
	result, err := l.svcCtx.DB.MenuParam.Get(l.ctx, in.Id)
	if err != nil {
		return nil, errorhandler.DefaultEntError(l.Logger, err, in)
	}

	return &core.MenuParamInfo{
		Id:        result.ID,
		CreatedAt: result.CreatedAt.UnixMilli(),
		UpdatedAt: result.UpdatedAt.UnixMilli(),
		Type:      result.Type,
		Key:       result.Key,
		Value:     result.Value,
	}, nil
}
