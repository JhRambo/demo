package routers

import (
	handler_account_AccountHttp "demo/gin/handlers/account_AccountHttp"
	handler_binary_BinaryHttp "demo/gin/handlers/binary_BinaryHttp"
	handler_hello_HelloHttp "demo/gin/handlers/hello_HelloHttp"
	handler_message_FeiShuService "demo/gin/handlers/message_FeiShuService"
	handler_message_PullService "demo/gin/handlers/message_PullService"
	handler_message_PushService "demo/gin/handlers/message_PushService"
	handler_msgpack_MsgpackHttp "demo/gin/handlers/msgpack_MsgpackHttp"
	handler_server_ServerHttp "demo/gin/handlers/server_ServerHttp"

	"github.com/gin-gonic/gin"
)

// 初始化路由
func InitRouters(r *gin.Engine) {
	r.POST("/v2/web/account/getVerificationCode", handler_account_AccountHttp.GetWebRegisterCode)
	r.POST("/v2/web/account/smsLogin", handler_account_AccountHttp.WebSMSLogin)
	r.POST("/v2/web/account/login", handler_account_AccountHttp.WebLogin)
	r.POST("/v2/web/account/logout", handler_account_AccountHttp.WebLogout)
	r.POST("/v2/web/account/changePwdCode", handler_account_AccountHttp.GetChangePwdCode)
	r.POST("/v2/web/account/checkcode", handler_account_AccountHttp.CheckCode)
	r.POST("/v2/web/account/changePwd", handler_account_AccountHttp.ChangePwdCode)
	r.POST("/v2/web/account/getEmailCode", handler_account_AccountHttp.GetEmailCode)
	r.POST("/v2/web/account/register", handler_account_AccountHttp.CheckEmail)
	r.POST("/v2/web/account/bypassAccountChangePwd", handler_account_AccountHttp.BypassAccountChangePWD)
	r.POST("/v2/web/account/info", handler_account_AccountHttp.AccountInfo)
	r.POST("/binary/uploadfile", handler_binary_BinaryHttp.UploadFile)
	r.POST("/hello/sayhello", handler_msgpack_MsgpackHttp.MsgPackProtocol)
	r.POST("/hello/saygoodbye", handler_hello_HelloHttp.SayGoodbye)
	r.POST("/v2/web/alarm/push", handler_message_FeiShuService.WebFsAlarmPush)
	r.POST("/v2/alarm/push", handler_message_FeiShuService.FsAlarmPush)
	r.POST("/v2/web/alarm/robot/set", handler_message_FeiShuService.WebFsAlarmRobotSet)
	r.POST("/v2/web/alarm/robot/get", handler_message_FeiShuService.WebFsAlarmRobotListGet)
	r.POST("/v2/web/alarm/robot/delete", handler_message_FeiShuService.WebFsAlarmRobotDelete)
	r.POST("/v2/message/invite", handler_message_PushService.PushInviteMessage)
	r.POST("/v2/message/proto1", handler_message_PullService.PullProtoMessage1)
	r.POST("/v2/message/proto2", handler_message_PullService.PullProtoMessage2)
	r.POST("/msgpack/protocol", handler_msgpack_MsgpackHttp.MsgPackProtocol)
	r.POST("/v2/web/server/list", handler_server_ServerHttp.GetServerList)
	r.POST("/v2/web/server/renderList", handler_server_ServerHttp.GetRenderServerList)
	r.POST("/v2/server/startVR", handler_server_ServerHttp.StartVR)
	r.POST("/v2/server/startPlanet", handler_server_ServerHttp.StartPlanet)
	r.POST("/v2/web/server/log/getList", handler_server_ServerHttp.GetLogFileList)
	r.POST("/v2/web/server/internalList", handler_server_ServerHttp.GetInternalList)
	r.POST("/v2/web/server/device/out", handler_server_ServerHttp.ServerDeviceOut)
	r.POST("/v2/web/server/stream/restart", handler_server_ServerHttp.ServersRestartSteam)
	r.POST("/v2/web/server/app/restart", handler_server_ServerHttp.ServersRestartApp)
	r.POST("/v2/server/device/addr", handler_server_ServerHttp.ServersAddr)
	r.POST("/v2/web/server/restart", handler_server_ServerHttp.RestartServer)
	r.POST("/v2/web/server/files", handler_server_ServerHttp.GetUpdateFiles)
	r.POST("/v2/web/server/update", handler_server_ServerHttp.UpdateServer)

}
