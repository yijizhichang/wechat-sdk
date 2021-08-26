//客户联系-统计管理
package customer

import (
	"encoding/json"
	"fmt"
	"github.com/yijizhichang/wechat-sdk/util"
	"github.com/yijizhichang/wechat-sdk/work/core"
)

const (
	GetCustomerBehaviorDataURL = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/get_user_behavior_data?access_token=%s"  //获取「联系客户统计」数据
	GetCustomerGroupChatDataURL = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/groupchat/statistic?access_token=%s"  //获取「群聊数据统计」数据 按群主聚合的方式
	GetCustomerGroupChatDataByDayURL = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/groupchat/statistic_group_by_day?access_token=%s"  //获取「群聊数据统计」数据 按自然日聚合的方式

)

//CustomerData
type CustomerData struct {
	*core.Context
}

//NewCustomer 实例化
func NewCustomerData(context *core.Context) *CustomerData {
	cd := new(CustomerData)
	cd.Context = context
	return cd
}

//获取「联系客户统计」数据
type CustomerBehaviorDataReq struct {
	Userid []string `json:"userid"`
	Partyid []int32 `json:"partyid"`
	StartTime int64 `json:"start_time"`
	EndTime int64 `json:"end_time"`
}
type CustomerBehaviorDataRep struct {
	util.WxError
	BehaviorData []behaviorData `json:"behavior_data"`
}
type behaviorData struct {
	StatTime int64 `json:"stat_time"`
	ChatCnt int64 `json:"chat_cnt"`
	MessageCnt int64 `json:"message_cnt"`
	ReplyPercentage float64 `json:"reply_percentage"`
	AvgReplyTime float64 `json:"avg_reply_time"`
	NegativeFeedbackCnt int64 `json:"negative_feedback_cnt"`
	NewApplyCnt int64 `json:"new_apply_cnt"`
	NewContactCnt int64 `json:"new_contact_cnt"`
}
func (cd *CustomerData) GetCustomerBehaviorData(accessToken string, req CustomerBehaviorDataReq)(result *CustomerBehaviorDataRep, err error){
	qyUrl := fmt.Sprintf(GetCustomerBehaviorDataURL, accessToken)

	response, err := util.PostJSON(qyUrl, req, cd.ProxyUrl)
	if err != nil {
		return
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("GetCustomerBehaviorData error : errcode=%d , errmsg=%s", result.ErrCode, result.ErrMsg)
	}
	return
}

//获取「群聊数据统计」数据
type CustomerGroupChatDataReq struct {
	DayBeginTime int64 `json:"day_begin_time"`
	DayEndTime   int64 `json:"day_end_time"`
	OwnerFilter  struct {
		UseridList []string `json:"userid_list"`
	} `json:"owner_filter"`
	OrderBy  int32 `json:"order_by"`
	OrderAsc int32 `json:"order_asc"`
	Offset   int32 `json:"offset"`
	Limit    int32 `json:"limit"`
}
type CustomerGroupChatDataRep struct {
	util.WxError
	Total      int64    `json:"total"`
	NextOffset int64    `json:"next_offset"`
	Items      []struct {
		Owner string `json:"owner"`
		Data  struct {
			NewChatCnt            int64 `json:"new_chat_cnt"`
			ChatTotal             int64 `json:"chat_total"`
			ChatHasMsg            int64 `json:"chat_has_msg"`
			NewMemberCnt          int64 `json:"new_member_cnt"`
			MemberTotal           int64 `json:"member_total"`
			MemberHasMsg          int64 `json:"member_has_msg"`
			MsgTotal              int64 `json:"msg_total"`
			MigrateTraineeChatCnt int64 `json:"migrate_trainee_chat_cnt"`
		} `json:"data"`
	} `json:"items"`
}
func (cd *CustomerData) GetCustomerGroupChatData(accessToken string, req CustomerGroupChatDataReq)(result *CustomerGroupChatDataRep, err error){
	qyUrl := fmt.Sprintf(GetCustomerGroupChatDataURL, accessToken)

	response, err := util.PostJSON(qyUrl, req, cd.ProxyUrl)
	if err != nil {
		return
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("GetCustomerGroupChatData error : errcode=%d , errmsg=%s", result.ErrCode, result.ErrMsg)
	}
	return
}

//获取「群聊数据统计」数据 按自然日聚合的方式
type CustomerGroupChatByDayDataReq struct {
	DayBeginTime int64 `json:"day_begin_time"`
	DayEndTime   int64 `json:"day_end_time"`
	OwnerFilter  struct {
		UseridList []string `json:"userid_list"`
	} `json:"owner_filter"`
}
type CustomerGroupChatByDayDataRep struct {
	util.WxError
	Items   []struct {
		StatTime int64 `json:"stat_time"`
		Data     struct {
			NewChatCnt            int32 `json:"new_chat_cnt"`
			ChatTotal             int32 `json:"chat_total"`
			ChatHasMsg            int32 `json:"chat_has_msg"`
			NewMemberCnt          int32 `json:"new_member_cnt"`
			MemberTotal           int32 `json:"member_total"`
			MemberHasMsg          int32 `json:"member_has_msg"`
			MsgTotal              int32 `json:"msg_total"`
			MigrateTraineeChatCnt int32 `json:"migrate_trainee_chat_cnt"`
		} `json:"data"`
	} `json:"items"`
}
func (cd *CustomerData) GetCustomerGroupChatByDayData(accessToken string, req CustomerGroupChatByDayDataReq)(result *CustomerGroupChatByDayDataRep, err error){
	qyUrl := fmt.Sprintf(GetCustomerGroupChatDataByDayURL, accessToken)

	response, err := util.PostJSON(qyUrl, req, cd.ProxyUrl)
	if err != nil {
		return
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("GetCustomerGroupChatByDayData error : errcode=%d , errmsg=%s", result.ErrCode, result.ErrMsg)
	}
	return
}

