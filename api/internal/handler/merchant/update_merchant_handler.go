package merchant

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/huuhoait/zero-admin-core/api/internal/logic/merchant"
	"github.com/huuhoait/zero-admin-core/api/internal/svc"
	"github.com/huuhoait/zero-admin-core/api/internal/types"
)

// swagger:route post /Merchant/update merchant UpdateMerchant
//
// Update Merchant information | 更新部门
//
// Update Merchant information | 更新部门
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: MerchantInfo
//
// Responses:
//  200: BaseMsgResp

func UpdateMerchantHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.MerchantInfo
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := merchant.NewUpdateMerchantLogic(r.Context(), svcCtx)
		resp, err := l.UpdateMerchant(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
