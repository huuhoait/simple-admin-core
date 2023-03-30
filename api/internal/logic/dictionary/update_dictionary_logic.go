package dictionary

import (
	"context"

	"github.com/huuhoait/zero-admin-core/api/internal/svc"
	"github.com/huuhoait/zero-admin-core/api/internal/types"
	"github.com/huuhoait/zero-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateDictionaryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateDictionaryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateDictionaryLogic {
	return &UpdateDictionaryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateDictionaryLogic) UpdateDictionary(req *types.DictionaryInfo) (resp *types.BaseMsgResp, err error) {
	data, err := l.svcCtx.CoreRpc.UpdateDictionary(l.ctx,
		&core.DictionaryInfo{
			Id:        req.Id,
			Title:     req.Title,
			Name:      req.Name,
			Status:    req.Status,
			Desc:      req.Desc,
			UpdatedBy: l.ctx.Value("userId").(string),
		})
	if err != nil {
		return nil, err
	}
	return &types.BaseMsgResp{Msg: l.svcCtx.Trans.Trans(l.ctx, data.Msg)}, nil
}
