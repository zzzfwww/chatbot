package handler

import (
	"net/http"

	"github.com/kevwan/chatbot/cli/ask/internal/logic"
	"github.com/kevwan/chatbot/cli/ask/internal/svc"
	"github.com/kevwan/chatbot/cli/ask/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func BotHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewBotLogic(r.Context(), svcCtx)
		resp, err := l.Bot(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
