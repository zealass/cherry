package cherryFacade

import (
	"google.golang.org/protobuf/proto"
)

type (
	SID        = string // session unique id
	UID        = int64  // 用户唯一id user unique id
	FrontendId = string // 前端节点id

	// INetwork 网络处理接口
	INetwork interface {
		SendRaw(bytes []byte)                                                        // write raw data to client
		RPC(nodeId string, route string, req proto.Message, rsp proto.Message) int32 // 调用remote rpc
		Response(mid uint, val interface{}, isError ...bool)                         // 回复消息到客户端
		Push(route string, val interface{})                                          // 推送消息对客户端
		Kick(reason interface{})                                                     // 踢下线
		RemoteAddr() string                                                          // 连接者的地址信息
		Close()                                                                      // 关闭接口
	}
)
