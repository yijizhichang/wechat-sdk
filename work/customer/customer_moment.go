//客户联系-客户朋友圈
package customer

import (
	"encoding/json"
	"fmt"
	"github.com/yijizhichang/wechat-sdk/util"
	"github.com/yijizhichang/wechat-sdk/work/core"
)

const (
	GetQyMomentListURL = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/get_moment_list?access_token=%s"  //获取企业全部的发表列表
	GetQyMomentTaskURL = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/get_moment_task?access_token=%s"  //获取客户朋友圈企业发表的列表
	GetQyMomentCustomerListURL = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/get_moment_customer_list?access_token=%s"  //获取客户朋友圈发表时选择的可见范围
	GetQyMomentSendResultURL = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/get_moment_send_result?access_token=%s"  //获取客户朋友圈发表后的可见客户列表
	GetQyMomentCommentsURL = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/get_moment_comments?access_token=%s"  //获取客户朋友圈的互动数据
	GetQyMomentStrategyListURL = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/moment_strategy/list?access_token=%s"  //获取规则组列表
	GetQyMomentStrategyViewURL = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/moment_strategy/get?access_token=%s"  //获取规则组详情
	GetQyMomentStrategyRangeURL = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/moment_strategy/get_range?access_token=%s"  //获取规则组管理范围
	CreateQyMomentStrategyURL = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/moment_strategy/create?access_token=%s"  //创建新的规则组
	UpdateQyMomentStrategyURL = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/moment_strategy/edit?access_token=%s"  //编辑规则组及其管理范围
	DelQyMomentStrategyURL = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/moment_strategy/del?access_token=%s"  //删除规则组

)

//CustomerMoment 客户朋友圈
type CustomerMoment struct {
	*core.Context
}

//NewCustomerMoment 实例化
func NewCustomerMoment(context *core.Context) *CustomerMoment {
	cmt := new(CustomerMoment)
	cmt.Context = context
	return cmt
}

//获取企业全部的发表列表
type MomentListReq struct {
	StartTime  int64    `json:"start_time"`
	EndTime    int64    `json:"end_time"`
	Creator    string `json:"creator"`
	FilterType int32    `json:"filter_type"`
	Cursor     string `json:"cursor"`
	Limit      int32    `json:"limit"`
}
type momentListRep struct {
	util.WxError
	NextCursor string `json:"next_cursor"`
	MomentList []struct {
		MomentId    string `json:"moment_id"`
		Creator     string `json:"creator"`
		CreateTime  string `json:"create_time"`
		CreateType  int32    `json:"create_type"`
		VisibleType int32    `json:"visible_type"`
		Text        struct {
			Content string `json:"content"`
		} `json:"text"`
		Image []struct {
			MediaId string `json:"media_id"`
		} `json:"image"`
		Video struct {
			MediaId      string `json:"media_id"`
			ThumbMediaId string `json:"thumb_media_id"`
		} `json:"video"`
		Link struct {
			Title string `json:"title"`
			Url   string `json:"url"`
		} `json:"link"`
		Location struct {
			Latitude  string `json:"latitude"`
			Longitude string `json:"longitude"`
			Name      string `json:"name"`
		} `json:"location"`
	} `json:"moment_list"`
}
func (cmt *CustomerMoment) GetQyMomentList(accessToken string, req MomentListReq)(result *momentListRep, err error){
	qyUrl := fmt.Sprintf(GetQyMomentListURL, accessToken)

	response, err := util.PostJSON(qyUrl, req, cmt.ProxyUrl)
	if err != nil {
		return
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("GetQyMomentList error : errcode=%d , errmsg=%s", result.ErrCode, result.ErrMsg)
	}
	return
}

//获取客户朋友圈企业发表的列表
type MomentTaskReq struct {
	MomentId string `json:"moment_id"`
	Cursor   string `json:"cursor"`
	Limit    int    `json:"limit"`
}
type momentTaskRep struct {
	util.WxError
	NextCursor string `json:"next_cursor"`
	TaskList   []struct {
		Userid        string `json:"userid"`
		PublishStatus int    `json:"publish_status"`
	} `json:"task_list"`
}
func (cmt *CustomerMoment) GetQyMomentTask(accessToken string, req MomentTaskReq)(result *momentTaskRep, err error){
	qyUrl := fmt.Sprintf(GetQyMomentTaskURL, accessToken)

	response, err := util.PostJSON(qyUrl, req, cmt.ProxyUrl)
	if err != nil {
		return
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("GetQyMomentTask error : errcode=%d , errmsg=%s", result.ErrCode, result.ErrMsg)
	}
	return
}

//获取客户朋友圈发表时选择的可见范围
type MomentCustomerListReq struct {
	MomentId string `json:"moment_id"`
	Userid   string `json:"userid"`
	Cursor   string `json:"cursor"`
	Limit    int32  `json:"limit"`
}
type momentCustomerListRep struct {
	util.WxError
	NextCursor   string `json:"next_cursor"`
	CustomerList []struct {
		Userid         string `json:"userid"`
		ExternalUserid string `json:"external_userid"`
	} `json:"customer_list"`
}
func (cmt *CustomerMoment) GetQyMomentCustomerList(accessToken string, req MomentCustomerListReq)(result *momentCustomerListRep, err error){
	qyUrl := fmt.Sprintf(GetQyMomentCustomerListURL, accessToken)

	response, err := util.PostJSON(qyUrl, req, cmt.ProxyUrl)
	if err != nil {
		return
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("GetQyMomentCustomerList error : errcode=%d , errmsg=%s", result.ErrCode, result.ErrMsg)
	}
	return
}

//获取客户朋友圈发表后的可见客户列表
type MomentSendResultReq struct {
	MomentId string `json:"moment_id"`
	Userid   string `json:"userid"`
	Cursor   string `json:"cursor"`
	Limit    int32  `json:"limit"`
}
type momentSendResultRep struct {
	util.WxError
	NextCursor   string `json:"next_cursor"`
	CustomerList []struct {
		ExternalUserid string `json:"external_userid"`
	} `json:"customer_list"`
}
func (cmt *CustomerMoment) GetQyMomentSendResult(accessToken string, req MomentSendResultReq)(result *momentSendResultRep, err error){
	qyUrl := fmt.Sprintf(GetQyMomentSendResultURL, accessToken)

	response, err := util.PostJSON(qyUrl, req, cmt.ProxyUrl)
	if err != nil {
		return
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("GetQyMomentSendResult error : errcode=%d , errmsg=%s", result.ErrCode, result.ErrMsg)
	}
	return
}

//获取客户朋友圈的互动数据
type MomentCommentsReq struct {
	MomentId string `json:"moment_id"`
	Userid   string `json:"userid"`
}
type momentCommentsRep struct {
	util.WxError
	CommentList []struct {
		ExternalUserid string `json:"external_userid,omitempty"`
		CreateTime     int64    `json:"create_time"`
		Userid         string `json:"userid,omitempty"`
	} `json:"comment_list"`
	LikeList []struct {
		ExternalUserid string `json:"external_userid,omitempty"`
		CreateTime     int64    `json:"create_time"`
		Userid         string `json:"userid,omitempty"`
	} `json:"like_list"`
}
func (cmt *CustomerMoment) GetQyMomentComments(accessToken string, req MomentCommentsReq)(result *momentCommentsRep, err error){
	qyUrl := fmt.Sprintf(GetQyMomentCommentsURL, accessToken)

	response, err := util.PostJSON(qyUrl, req, cmt.ProxyUrl)
	if err != nil {
		return
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("GetQyMomentComments error : errcode=%d , errmsg=%s", result.ErrCode, result.ErrMsg)
	}
	return
}

//获取规则组列表
type QyMomentStrategyReq struct {
	Cursor string `json:"cursor"`
	Limit  int32  `json:"limit"`
}
type qyMomentStrategyList struct {
	util.WxError
	Strategy []struct {
		StrategyId int32 `json:"strategy_id"`
	} `json:"strategy"`
	NextCursor string `json:"next_cursor"`
}
func (cmt *CustomerMoment) GetQyMomentStrategyList(accessToken string, req QyMomentStrategyReq)(result *qyMomentStrategyList, err error){
	qyUrl := fmt.Sprintf(GetQyMomentStrategyListURL, accessToken)

	response, err := util.PostJSON(qyUrl, req, cmt.ProxyUrl)
	if err != nil {
		return
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("GetQyMomentStrategyList error : errcode=%d , errmsg=%s", result.ErrCode, result.ErrMsg)
	}
	return
}

//获取规则组详情
type QyMomentStrategyViewReq struct {
	StrategyId int32 `json:"strategy_id"`
}
type qyMomentStrategyViewRep struct {
	util.WxError
	Strategy struct {
		StrategyId   int32    `json:"strategy_id"`
		ParentId     int32    `json:"parent_id"`
		StrategyName string   `json:"strategy_name"`
		CreateTime   int32    `json:"create_time"`
		AdminList    []string `json:"admin_list"`
		Privilege    struct {
			ViewMomentList bool `json:"view_moment_list"`
			SendMoment     bool `json:"send_moment"`
		} `json:"privilege"`
	} `json:"strategy"`
}
func (cmt *CustomerMoment) GetQyMomentStrategyView(accessToken string, req QyMomentStrategyViewReq)(result *qyMomentStrategyViewRep, err error){
	qyUrl := fmt.Sprintf(GetQyMomentStrategyViewURL, accessToken)

	response, err := util.PostJSON(qyUrl, req, cmt.ProxyUrl)
	if err != nil {
		return
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("GetQyMomentStrategyView error : errcode=%d , errmsg=%s", result.ErrCode, result.ErrMsg)
	}
	return
}

//获取规则组管理范围
type QyMomentStrategyRangeReq struct {
	StrategyId int32  `json:"strategy_id"`
	Cursor     string `json:"cursor"`
	Limit      int32  `json:"limit"`
}
type qyMomentStrategyRangeRep struct {
	util.WxError
	Range   []struct {
		Type    int32  `json:"type"`
		Userid  string `json:"userid,omitempty"`
		Partyid int32  `json:"partyid,omitempty"`
	} `json:"range"`
	NextCursor string `json:"next_cursor"`
}
func (cmt *CustomerMoment) GetQyMomentStrategyRange(accessToken string, req QyMomentStrategyRangeReq)(result *qyMomentStrategyRangeRep, err error){
	qyUrl := fmt.Sprintf(GetQyMomentStrategyRangeURL, accessToken)

	response, err := util.PostJSON(qyUrl, req, cmt.ProxyUrl)
	if err != nil {
		return
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("GetQyMomentStrategyRange error : errcode=%d , errmsg=%s", result.ErrCode, result.ErrMsg)
	}
	return
}

//创建新的规则组
type CreateQyMomentStrategyReq struct {
	ParentId     int32    `json:"parent_id"`
	StrategyName string   `json:"strategy_name"`
	AdminList    []string `json:"admin_list"`
	Privilege    struct {
		SendMoment     bool `json:"send_moment"`
		ViewMomentList bool `json:"view_moment_list"`
	} `json:"privilege"`
	Range []struct {
		Type    int32  `json:"type"`
		Userid  string `json:"userid,omitempty"`
		Partyid int32  `json:"partyid,omitempty"`
	} `json:"range"`
}
type createQyMomentStrategyRep struct {
	util.WxError
	StrategyId int    `json:"strategy_id"`
}
func (cmt *CustomerMoment) CreateQyMomentStrategy(accessToken string, req CreateQyMomentStrategyReq)(result *createQyMomentStrategyRep, err error){
	qyUrl := fmt.Sprintf(CreateQyMomentStrategyURL, accessToken)

	response, err := util.PostJSON(qyUrl, req, cmt.ProxyUrl)
	if err != nil {
		return
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("CreateQyMomentStrategy error : errcode=%d , errmsg=%s", result.ErrCode, result.ErrMsg)
	}
	return
}

//编辑规则组及其管理范围
type UpdateQyMomentStrategyReq struct {
	StrategyId   int32      `json:"strategy_id"`
	StrategyName string   `json:"strategy_name"`
	AdminList    []string `json:"admin_list"`
	Privilege    struct {
		ViewMomentList bool `json:"view_moment_list"`
		SendMoment     bool `json:"send_moment"`
	} `json:"privilege"`
	RangeAdd []struct {
		Type    int32    `json:"type"`
		Userid  string `json:"userid,omitempty"`
		Partyid int32    `json:"partyid,omitempty"`
	} `json:"range_add"`
	RangeDel []struct {
		Type    int32    `json:"type"`
		Userid  string `json:"userid,omitempty"`
		Partyid int32    `json:"partyid,omitempty"`
	} `json:"range_del"`
}
func (cmt *CustomerMoment) UpdateQyMomentStrategy(accessToken string, req UpdateQyMomentStrategyReq)(result *util.WxError, err error){
	qyUrl := fmt.Sprintf(UpdateQyMomentStrategyURL, accessToken)

	response, err := util.PostJSON(qyUrl, req, cmt.ProxyUrl)
	if err != nil {
		return
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("UpdateQyMomentStrategy error : errcode=%d , errmsg=%s", result.ErrCode, result.ErrMsg)
	}
	return
}

//删除规则组
type DelQyMomentStrategyReq struct {
	StrategyId int32 `json:"strategy_id"`
}
func (cmt *CustomerMoment) DelQyMomentStrategy(accessToken string, req DelQyMomentStrategyReq)(result *util.WxError, err error){
	qyUrl := fmt.Sprintf(DelQyMomentStrategyURL, accessToken)

	response, err := util.PostJSON(qyUrl, req, cmt.ProxyUrl)
	if err != nil {
		return
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("DelQyMomentStrategy error : errcode=%d , errmsg=%s", result.ErrCode, result.ErrMsg)
	}
	return
}
