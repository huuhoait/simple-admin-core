package merchantmeta

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/huuhoait/zero-admin-core/api/internal/logic/merchantmeta"
	"github.com/huuhoait/zero-admin-core/api/internal/svc"
	"github.com/huuhoait/zero-admin-core/api/internal/types"
)

// swagger:route post /merchant_meta/list merchantmeta GetMerchantMetaList
//
// Get merchant meta list | 获取字典键值列表
//
// Get merchant meta list | 获取字典键值列表
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: MerchantMetaListReq
//
// Responses:
//  200: MerchantMetaListResp

func GetMerchantMetaListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.MerchantMetaListReq
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := merchantmeta.NewGetMerchantMetaListLogic(r.Context(), svcCtx)
		resp, err := l.GetMerchantMetaList(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
