// Code generated by hertz generator.

package Api

import (
	"context"
	"fmt"

	// "github.com/YANGJUNYAN0715/douyin/tree/guo/cmd/api/biz/mw"
	"github.com/YANGJUNYAN0715/douyin/tree/guo/pkg/errno"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/middlewares/server/recovery"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/hertz-contrib/gzip"
	"github.com/hertz-contrib/requestid"
)

func rootMw() []app.HandlerFunc {
	// your code...
	return []app.HandlerFunc{
		// use recovery mw
		recovery.Recovery(recovery.WithRecoveryHandler(
			func(ctx context.Context, c *app.RequestContext, err interface{}, stack []byte) {
				hlog.SystemLogger().CtxErrorf(ctx, "[Recovery] err=%v\nstack=%s", err, stack)
				c.JSON(consts.StatusInternalServerError, utils.H{
					"code":    errno.ServiceErr.ErrCode,
					"message": fmt.Sprintf("[Recovery] err=%v\nstack=%s", err, stack),
				})
			},
		)),
		// use requestid mw
		requestid.New(),
		// use gzip mw
		gzip.Gzip(gzip.DefaultCompression),
	}
}

func _douyinMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _messageMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _actionMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _message_ctionmessageMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _chatMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _messagechatmessageMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _userMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _loginMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _loginuserMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _registerMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _registeruserMw() []app.HandlerFunc {
	// your code...
	return nil
}
