package call

import (
	"../LineThrift"
	"../thrift"
	"context"
)

var err error

func TalkService(AuthToken string) *LineThrift.TalkServiceClient {
	//fmt.Println("#### TalkService Initiated. ####")
	httpClient, _ := thrift.NewTHttpClient("https://legy-jp-addr-long.line.naver.jp/S4")
	buffer := thrift.NewTBufferedTransportFactory(4096)
	trans := httpClient.(*thrift.THttpClient)
	trans.SetHeader("User-Agent","Line/11.17.1")
	trans.SetHeader("X-Line-Application","ANDROID\t11.17.1\tAndroid Os\t9.1.1")
	trans.SetHeader("X-Line-Carrier","51089, 1-0")
	trans.SetHeader("X-Line-Access",AuthToken)
	buftrans, _ := buffer.GetTransport(trans)
	compactProtocol := thrift.NewTCompactProtocolFactory()
	return LineThrift.NewTalkServiceClientFactory(buftrans, compactProtocol)
}

func CallService(AuthToken string) *LineThrift.CallServiceClient {
	//fmt.Println("#### CallService Initiated. ####")
	httpClient, _ := thrift.NewTHttpClient("https://legy-jp-addr-long.line.naver.jp/V4")
	buffer := thrift.NewTBufferedTransportFactory(4096)
	trans := httpClient.(*thrift.THttpClient)
	trans.SetHeader("User-Agent","Line/11.17.1")
	trans.SetHeader("X-Line-Application","ANDROID\t11.17.1\tAndroid Os\t9.1.1")
	trans.SetHeader("X-Line-Carrier","51089, 1-0")
	trans.SetHeader("X-Line-Access",AuthToken)
	buftrans, _ := buffer.GetTransport(trans)
	compactProtocol := thrift.NewTCompactProtocolFactory()
	return LineThrift.NewCallServiceClientFactory(buftrans, compactProtocol)
}

func CreateGroup(name string,midlist []string,AuthToken string) (err error) {
	TS := TalkService(AuthToken)
	_,err = TS.CreateGroup(context.TODO(), int32(0), name, midlist)
	return err
}

func AcquireCallRoute(to string, AuthToken string) (r []string, err error) {
	TS := TalkService(AuthToken)
	res, err := TS.AcquireCallRoute(context.TODO(), to)
	return res, err
}

func AcquireGroupCallRoute(chatMid string, AuthToken string) (r *LineThrift.GroupCallRoute, err error) {
	CS := CallService(AuthToken)
	res, err := CS.AcquireGroupCallRoute(context.TODO(), chatMid, LineThrift.GroupCallMediaType_AUDIO)
	return res, err
}

func GetGroupCall(chatMid string, AuthToken string) (r *LineThrift.GroupCall, err error) {
	CS := CallService(AuthToken)
	res, err := CS.GetGroupCall(context.TODO(), chatMid)
	return res, err
}

func InviteIntoGroupCall(chatMid string, AuthToken string, memberMids []string) (err error) {
	CS := CallService(AuthToken)
	err = CS.InviteIntoGroupCall(context.TODO(), chatMid, memberMids, LineThrift.GroupCallMediaType_AUDIO)
	return err
}