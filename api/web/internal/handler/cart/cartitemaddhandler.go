package cart

import (
	"net/http"

	"github.com/feihua/zero-admin/api/web/internal/logic/cart"
	"github.com/feihua/zero-admin/api/web/internal/svc"
	"github.com/feihua/zero-admin/api/web/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func CartItemAddHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CartItemAddReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := cart.NewCartItemAddLogic(r.Context(), svcCtx)
		resp, err := l.CartItemAdd(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
