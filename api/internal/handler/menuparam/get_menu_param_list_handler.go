package menuparam

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/huuhoait/zero-admin-core/api/internal/logic/menuparam"
	"github.com/huuhoait/zero-admin-core/api/internal/svc"
	"github.com/huuhoait/zero-admin-core/api/internal/types"
)

// swagger:route post /menu_param/list menuparam GetMenuParamList
//
// Get menu param list | 获取菜单参数列表
//
// Get menu param list | 获取菜单参数列表
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: MenuParamListReq
//
// Responses:
//  200: MenuParamListResp

func GetMenuParamListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.MenuParamListReq
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := menuparam.NewGetMenuParamListLogic(r.Context(), svcCtx)
		resp, err := l.GetMenuParamList(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
