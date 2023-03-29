package menuparam

import (
	"context"

	"github.com/huuhoait/zero-admin-core/rpc/ent/menuparam"
	"github.com/huuhoait/zero-admin-core/rpc/ent/predicate"

	"github.com/huuhoait/zero-admin-core/rpc/internal/svc"
	"github.com/huuhoait/zero-admin-core/rpc/internal/utils/errorhandler"
	"github.com/huuhoait/zero-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMenuParamListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetMenuParamListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMenuParamListLogic {
	return &GetMenuParamListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetMenuParamListLogic) GetMenuParamList(in *core.MenuParamListReq) (*core.MenuParamListResp, error) {
	var predicates []predicate.MenuParam
	if in.MenuId != 0 {
		predicates = append(predicates, menuparam.MenuIDEQ(in.MenuId))
	}
	result, err := l.svcCtx.DB.MenuParam.Query().Where(predicates...).Page(l.ctx, in.Page, in.PageSize)
	if err != nil {
		return nil, errorhandler.DefaultEntError(l.Logger, err, in)
	}

	resp := &core.MenuParamListResp{}
	resp.Total = result.PageDetails.Total

	for _, v := range result.List {
		resp.Data = append(resp.Data, &core.MenuParamInfo{
			Id:        v.ID,
			CreatedAt: v.CreatedAt.UnixMilli(),
			UpdatedAt: v.UpdatedAt.UnixMilli(),
			Type:      v.Type,
			Key:       v.Key,
			Value:     v.Value,
		})
	}

	return resp, nil
}
