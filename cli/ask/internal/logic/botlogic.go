package logic

import (
	"context"
	"github.com/kevwan/chatbot/cli/ask/internal/svc"
	"github.com/kevwan/chatbot/cli/ask/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type BotLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewBotLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BotLogic {
	return &BotLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *BotLogic) Bot(req *types.Request) (resp *types.Response, err error) {
	resp = &types.Response{}
	question := req.Message
	answers := l.svcCtx.ChatBot.GetResponse(question)
	if len(answers) == 0 {
		l.Logger.Infof("ask:%v, no answers", req.Message)
		return
	}
	for i, answer := range answers {
		resp.Message = append(resp.Message, answer.Content)
		if i >= l.svcCtx.Tops {
			break
		}
	}
	return
}
