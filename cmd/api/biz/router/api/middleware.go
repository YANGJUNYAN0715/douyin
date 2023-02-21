// Code generated by hertz generator.

package Api

import (
	"context"
	"fmt"
	
	"github.com/YANGJUNYAN0715/douyin/tree/guo/cmd/api/biz/mw"
	"github.com/YANGJUNYAN0715/douyin/tree/guo/pkg/errno"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/middlewares/server/recovery"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/hertz-contrib/gzip"
	"github.com/hertz-contrib/requestid"
	// "go.opentelemetry.io/otel/trace"
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
	return []app.HandlerFunc{
		// use jwt mw
		mw.JwtMiddleware.MiddlewareFunc(),
	}
}

func _actionMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _message_ctionMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _chatMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _messagechatMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _publishMw() []app.HandlerFunc {
	// your code...
	return []app.HandlerFunc{
		// use jwt mw
		mw.JwtMiddleware.MiddlewareFunc(),
	}
}

func _action0Mw() []app.HandlerFunc {
	// your code...
	return nil
}

func _publish_ctionMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _listMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _publishlistMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _relationMw() []app.HandlerFunc {
	// your code...
	return []app.HandlerFunc{
		// use jwt mw
		mw.JwtMiddleware.MiddlewareFunc(),
	}
}

func _action1Mw() []app.HandlerFunc {
	// your code...
	return nil
}

func _relation_ctionMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _followMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _list0Mw() []app.HandlerFunc {
	// your code...
	return nil
}

func _relationfollowlistMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _followerMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _list1Mw() []app.HandlerFunc {
	// your code...
	return nil
}

func _relationfollowerlistMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _friendMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _list2Mw() []app.HandlerFunc {
	// your code...
	return nil
}

func _relationfriendlistMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _userMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _userinfoMw() []app.HandlerFunc {
	// your code...
	return []app.HandlerFunc{
		// use jwt mw
		mw.JwtMiddleware.MiddlewareFunc(),
	}
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
