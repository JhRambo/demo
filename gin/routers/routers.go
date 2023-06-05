package routers

import (
	handler_account "demo/gin/handlers/account"
	handler_binary "demo/gin/handlers/binary"
	handler_hello "demo/gin/handlers/hello"
	handler_message "demo/gin/handlers/message"
	handler_msgpack "demo/gin/handlers/msgpack"
	handler_server "demo/gin/handlers/server"

	"github.com/gin-gonic/gin"
)

// 初始化路由
func InitRouters(r *gin.Engine) {
	r.POST("/v2/web/account/getVerificationCode", handler_account.GetWebRegisterCode)
	r.POST("/v2/web/account/smsLogin", handler_account.WebSMSLogin)
	r.POST("/v2/web/account/login", handler_account.WebLogin)
	r.POST("/v2/web/account/logout", handler_account.WebLogout)
	r.POST("/v2/web/account/changePwdCode", handler_account.GetChangePwdCode)
	r.POST("/v2/web/account/checkcode", handler_account.CheckCode)
	r.POST("/v2/web/account/changePwd", handler_account.ChangePwdCode)
	r.POST("/v2/web/account/getEmailCode", handler_account.GetEmailCode)
	r.POST("/v2/web/account/register", handler_account.CheckEmail)
	r.POST("/v2/web/account/bypassAccountChangePwd", handler_account.BypassAccountChangePWD)
	r.POST("/v2/web/account/info", handler_account.AccountInfo)
	r.POST("/binary/uploadfile", handler_binary.UploadFile)
	r.POST("/hello/sayhello", handler_msgpack.MsgPackProtocol)
	r.POST("/hello/saygoodbye", handler_hello.SayGoodbye)
	r.POST("/v2/message/invite", handler_message.PushInviteMessage)
	r.POST("/v2/web/alarm/push", handler_message.WebFsAlarmPush)
	r.POST("/v2/alarm/push", handler_message.FsAlarmPush)
	r.POST("/v2/web/alarm/robot/set", handler_message.WebFsAlarmRobotSet)
	r.POST("/v2/web/alarm/robot/get", handler_message.WebFsAlarmRobotListGet)
	r.POST("/v2/web/alarm/robot/delete", handler_message.WebFsAlarmRobotDelete)
	r.POST("/msgpack/protocol", handler_msgpack.MsgPackProtocol)
	r.POST("/v2/web/server/list", handler_server.GetServerList)
	r.POST("/v2/web/server/renderList", handler_server.GetRenderServerList)
	r.POST("/v2/server/startVR", handler_server.StartVR)
	r.POST("/v2/server/startPlanet", handler_server.StartPlanet)
	r.POST("/v2/web/server/log/getList", handler_server.GetLogFileList)
	r.POST("/v2/web/server/internalList", handler_server.GetInternalList)
	r.POST("/v2/web/server/device/out", handler_server.ServerDeviceOut)
	r.POST("/v2/web/server/stream/restart", handler_server.ServersRestartSteam)
	r.POST("/v2/web/server/app/restart", handler_server.ServersRestartApp)
	r.POST("/v2/server/device/addr", handler_server.ServersAddr)
	r.POST("/v2/web/server/restart", handler_server.RestartServer)
	r.POST("/v2/web/server/files", handler_server.GetUpdateFiles)
	r.POST("/v2/web/server/update", handler_server.UpdateServer)

}
