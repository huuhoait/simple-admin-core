package oauthprovider

import (
	"context"
	"net/http"
	"strings"
	"time"

	"github.com/huuhoait/zero-tools/utils/jwt"

	"github.com/huuhoait/zero-admin-core/api/internal/svc"
	"github.com/huuhoait/zero-admin-core/api/internal/types"
	"github.com/huuhoait/zero-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type OauthCallbackLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	r      *http.Request
}

func NewOauthCallbackLogic(r *http.Request, svcCtx *svc.ServiceContext) *OauthCallbackLogic {
	return &OauthCallbackLogic{
		Logger: logx.WithContext(r.Context()),
		ctx:    r.Context(),
		svcCtx: svcCtx,
		r:      r,
	}
}

func (l *OauthCallbackLogic) OauthCallback() (resp *types.CallbackResp, err error) {
	result, err := l.svcCtx.CoreRpc.OauthCallback(l.ctx, &core.CallbackReq{
		State: l.r.FormValue("state"),
		Code:  l.r.FormValue("code"),
	})
	if err != nil {
		return nil, err
	}

	token, err := jwt.NewJwtToken(l.svcCtx.Config.Auth.AccessSecret, result.Id, "roleId", time.Now().Unix(),
		l.svcCtx.Config.Auth.AccessExpire, result.RoleCodes)

	// add token into database
	expiredAt := time.Now().Add(time.Second * 259200).Unix()
	_, err = l.svcCtx.CoreRpc.CreateToken(l.ctx, &core.TokenInfo{
		Uuid:      result.Id,
		Token:     token,
		Source:    strings.Split(l.r.FormValue("state"), "-")[1],
		Status:    1,
		ExpiredAt: expiredAt,
	})

	if err != nil {
		return nil, err
	}

	return &types.CallbackResp{
		UserId: result.Id,
		Token:  token,
		Expire: uint64(time.Now().Add(time.Second * 259200).Unix()),
	}, nil
}
