package merchantmeta

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/huuhoait/zero-admin-core/api/internal/logic/merchantmeta"
	"github.com/huuhoait/zero-admin-core/api/internal/svc"
	"github.com/huuhoait/zero-admin-core/api/internal/types"
)

// swagger:route post /merchant_meta merchantmeta GetMerchantMetaById
//
// Get MerchantMeta by ID | 通过ID获取字典键值
//
// Get MerchantMeta by ID | 通过ID获取字典键值
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: IDReq
//
// Responses:
//  200: MerchantMetaInfoResp

func GetMerchantMetaByIdHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.IDReq
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := merchantmeta.NewGetMerchantMetaByIdLogic(r.Context(), svcCtx)
		resp, err := l.GetMerchantMetaById(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
