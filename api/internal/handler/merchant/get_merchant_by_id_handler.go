package merchant

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/huuhoait/zero-admin-core/api/internal/logic/merchant"
	"github.com/huuhoait/zero-admin-core/api/internal/svc"
	"github.com/huuhoait/zero-admin-core/api/internal/types"
)

// swagger:route post /Merchant merchant GetMerchantById
//
// Get Merchant by ID | 通过ID获取部门
//
// Get Merchant by ID | 通过ID获取部门
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: IDReq
//
// Responses:
//  200: MerchantInfoResp

func GetMerchantByIdHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.IDReq
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := merchant.NewGetMerchantByIdLogic(r.Context(), svcCtx)
		resp, err := l.GetMerchantById(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
