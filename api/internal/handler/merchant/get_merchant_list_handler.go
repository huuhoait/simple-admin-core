package merchant

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/huuhoait/zero-admin-core/api/internal/logic/merchant"
	"github.com/huuhoait/zero-admin-core/api/internal/svc"
	"github.com/huuhoait/zero-admin-core/api/internal/types"
)

// swagger:route post /Merchant/list merchant GetMerchantList
//
// Get Merchant list | 获取部门列表
//
// Get Merchant list | 获取部门列表
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: MerchantListReq
//
// Responses:
//  200: MerchantListResp

func GetMerchantListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.MerchantListReq
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := merchant.NewGetMerchantListLogic(r.Context(), svcCtx)
		resp, err := l.GetMerchantList(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
