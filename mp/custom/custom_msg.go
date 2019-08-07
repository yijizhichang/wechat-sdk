//聊天记录
package custom

import (
	"github.com/yijizhichang/wechat-sdk/util"
	"fmt"
	"encoding/json"
	"time"
)

const(
	GetMsgListURL   = "https://api.weixin.qq.com/customservice/msgrecord/getmsglist?access_token=%s"    //获取聊天记录
)


//获取聊天记录
type msgList struct {
	StartTime 	int64		`json:"starttime"`
	EndTime		int64		`json:"endtime"`
	MsgId 		int64 		`json:"msgid"`
	Number      int64		`json:"number"`
}

type msgListResult struct{
	util.WxError
	Number          int64			`json:"number"`
	MsgId			int64			`json:"msgid"`
	RecordList  	[] *msgInfo    `json:"recordlist"`
}

type msgInfo struct{
	Openid      string		`json:"openid"`
	Opercode    int64		`json:"opercode"`
	Text        string		`json:"text"`
	Time        int64		`json:"time"`
	Worker      string		`json:"worker"`
}
func (kf *Custom) GetMsgList(startTimeStr, endTimeStr string,  msgId, number int64) (result *msgListResult, err error) {
	accessToken, err := kf.GetAccessToken()
	if err != nil {
		return
	}
	wxUrl := fmt.Sprintf(GetMsgListURL, accessToken)

	postData := new(msgList)
	startTime,_ := time.Parse("2006-01-02 15:04:05", startTimeStr)
	endTime,_ := time.Parse("2006-01-02 15:04:05", endTimeStr)
	postData.StartTime = startTime.Unix()
	postData.EndTime = endTime.Unix()

	response, err := util.PostJSON(wxUrl, postData, kf.ProxyUrl)
	if err != nil {
		return
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("GetMsgList error : errcode=%v , errmsg=%v", result.ErrCode, string(result.ErrMsg))
	}
	return
}
