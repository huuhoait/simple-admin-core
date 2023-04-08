package merchantmeta

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/huuhoait/zero-admin-core/api/internal/logic/merchantmeta"
	"github.com/huuhoait/zero-admin-core/api/internal/svc"
	"github.com/huuhoait/zero-admin-core/api/internal/types"
)

// swagger:route post /merchant_meta/update merchantmeta UpdateMerchantMeta
//
// Update merchant meta information | 更新字典键值
//
// Update merchant meta information | 更新字典键值
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: MerchantMetaInfo
//
// Responses:
//  200: BaseMsgResp

func UpdateMerchantMetaHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.MerchantMetaInfo
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := merchantmeta.NewUpdateMerchantMetaLogic(r.Context(), svcCtx)
		resp, err := l.UpdateMerchantMeta(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
