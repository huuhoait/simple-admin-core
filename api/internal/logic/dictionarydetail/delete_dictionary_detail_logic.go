package dictionarydetail

import (
	"context"
	"fmt"

	"github.com/huuhoait/simple-admin-common/i18n"
	"github.com/zeromicro/go-zero/core/errorx"

	"github.com/huuhoait/simple-admin-core/api/internal/logic/dictionary"
	"github.com/huuhoait/simple-admin-core/api/internal/svc"
	"github.com/huuhoait/simple-admin-core/api/internal/types"
	"github.com/huuhoait/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteDictionaryDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteDictionaryDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteDictionaryDetailLogic {
	return &DeleteDictionaryDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteDictionaryDetailLogic) DeleteDictionaryDetail(req *types.IDsReq) (resp *types.BaseMsgResp, err error) {
	detailData, err := l.svcCtx.CoreRpc.GetDictionaryDetailById(l.ctx, &core.IDReq{Id: req.Ids[0]})
	if err != nil {
		return nil, err
	}

	dict, err := dictionary.NewGetDictionaryByIdLogic(l.ctx, l.svcCtx).GetDictionaryById(&types.IDReq{Id: *detailData.DictionaryId})
	if err != nil {
		return nil, err
	}

	if _, err := l.svcCtx.Redis.DelCtx(l.ctx, fmt.Sprintf("dict_%d", dict.Data.Name)); err != nil {
		logx.Errorw("failed to delete dictionary data in redis", logx.Field("detail", err))
		return nil, errorx.NewCodeInternalError(i18n.RedisError)
	}

	result, err := l.svcCtx.CoreRpc.DeleteDictionaryDetail(l.ctx, &core.IDsReq{
		Ids: req.Ids,
	})
	if err != nil {
		return nil, err
	}

	return &types.BaseMsgResp{Msg: l.svcCtx.Trans.Trans(l.ctx, result.Msg)}, nil
}
