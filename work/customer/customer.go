//客户联系-客户管理
package customer

import (
	"encoding/json"
	"fmt"
	"github.com/yijizhichang/wechat-sdk/util"
	"github.com/yijizhichang/wechat-sdk/work/core"
)

const (
	GetQyCustomerListURL = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/list?access_token=%s&userid=%s"  //获取客户列表
	GetQyCustomerViewURL = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/get?access_token=%s&external_userid=%s&cursor=%s"  //获取客户详情
	GetQyCustomerViewBatchURL = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/batch/get_by_user?access_token=%s"  //批量获取客户详情
	UpdateQyCustomerRemarkURL = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/remark?access_token=%s"  //修改客户备注信息
	GetQyCustomerStrategyListURL = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/customer_strategy/list?access_token=%s"  //获取规则组列表
	GetQyCustomerStrategyViewURL = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/customer_strategy/get?access_token=%s"  //获取规则组详情
	GetQyCustomerStrategyRangeURL = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/customer_strategy/get_range?access_token=%s"  //获取规则组管理范围
	CreateQyCustomerStrategyURL = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/customer_strategy/create?access_token=%s"  //创建新的规则组
	UpdateQyCustomerStrategyURL = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/customer_strategy/edit?access_token=%s"  //编辑规则组及其管理范围
	DelQyCustomerStrategyURL = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/customer_strategy/del?access_token=%s"  //删除规则组
)

//Customer 客户管理
type Customer struct {
	*core.Context
}

//NewCustomer 实例化
func NewCustomer(context *core.Context) *Customer {
	cus := new(Customer)
	cus.Context = context
	return cus
}

//获取客户列表
type CusList struct {
	util.WxError
	ExternalUserid []string `json:"external_userid"`
}
func (c *Customer) GetQyCustomerList(accessToken, userid string)(result *CusList, err error){
	qyUrl := fmt.Sprintf(GetQyCustomerListURL, accessToken, userid)

	response, err := util.HTTPGet(qyUrl, c.ProxyUrl)
	if err != nil {
		return
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("GetQyCustomerList error : errcode=%d , errmsg=%s", result.ErrCode, result.ErrMsg)
	}
	return
}

//获取客户详情
type CusView struct {
	util.WxError
	ExternalContact struct{
		ExternalUserid  string  `json:"external_userid"`
		Name string `json:"name"`
		Position string `json:"position"`
		Avatar string `json:"avatar"`
		CorpName string `json:"corp_name"`
		CorpFullName string `json:"corp_full_name"`
		Type int64 `json:"type"`
		Gender int64 `json:"gender"`
		Unionid string `json:"unionid"`
		ExternalProfile struct{
			ExternalAttr []externalAttrItem `json:"external_attr"`
		} `json:"external_profile"`
	} `json:"external_contact"`
	FollowUser []followUserItem `json:"follow_user"`
	NextCursor string `json:"next_cursor"`
}
type externalAttrItem struct {
	Type int64 `json:"type"`
	Name string `json:"name"`
	Text struct{
		Value string `json:"value"`
	} `json:"text,omitempty"`
	Web struct{
		Url string `json:"url"`
		Title string `json:"title"`
	} `json:"web,omitempty"`
	Miniprogram struct{
		Appid string `json:"appid"`
		Pagepath string `json:"pagepath"`
		Title string `json:"title"`
	} `json:"miniprogram,omitempty"`
}
type followUserItem struct {
	Userid string `json:"userid"`
	Remark string `json:"remark"`
	Description string `json:"description"`
	Createtime int64 `json:"createtime"`
	Tags []tagItem `json:"tags"`
	RemarkCorpName string `json:"remark_corp_name"`
	RemarkMobiles []string `json:"remark_mobiles"`
	OperUserid string `json:"oper_userid"`
	AddWay int64 `json:"add_way"`
	State string `json:"state"`
}
type tagItem struct {
	GroupName string `json:"group_name"`
	TagName string `json:"tag_name"`
	TagId string `json:"tag_id"`
	Type int64 `json:"type"`
}
func (c *Customer) GetQyCustomerView(accessToken, externalUserid, cursor string)(result *CusView, err error){
	qyUrl := fmt.Sprintf(GetQyCustomerViewURL, accessToken, externalUserid, cursor)

	response, err := util.HTTPGet(qyUrl, c.ProxyUrl)
	if err != nil {
		return
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("GetQyCustomerView error : errcode=%d , errmsg=%s", result.ErrCode, result.ErrMsg)
	}
	return
}

//批量获取客户详情
type CusViewBatch struct {
	util.WxError
	ExternalContactList []externalContactItem `json:"external_contact_list"`
}
type externalContactItem struct {
	ExternalContact struct{
		ExternalUserid  string  `json:"external_userid"`
		Name string `json:"name"`
		Position string `json:"position"`
		Avatar string `json:"avatar"`
		CorpName string `json:"corp_name"`
		CorpFullName string `json:"corp_full_name"`
		Type int64 `json:"type"`
		Gender int64 `json:"gender"`
		Unionid string `json:"unionid"`
		ExternalProfile struct{
			ExternalAttr []externalAttrItem `json:"external_attr"`
		} `json:"external_profile"`
	} `json:"external_contact"`
	FollowInfo followInfoItem `json:"follow_info"`
}
type followInfoItem struct {
	Userid string `json:"userid"`
	Remark string `json:"remark"`
	Description string `json:"description"`
	Createtime int64 `json:"createtime"`
	Tags []tagItem `json:"tags"`
	RemarkCorpName string `json:"remark_corp_name"`
	RemarkMobiles []string `json:"remark_mobiles"`
	OperUserid string `json:"oper_userid"`
	AddWay int64 `json:"add_way"`
	State string `json:"state"`
}
type CusViewBatchReq struct {
	UseridList  []string  `json:"userid_list"`
	Cursor  string  `json:"cursor"`
	Limit   int     `json:"limit"`
}
func (c *Customer) GetQyCustomerViewBatch(accessToken string, req CusViewBatchReq)(result *CusViewBatch, err error){
	qyUrl := fmt.Sprintf(GetQyCustomerViewBatchURL, accessToken)

	response, err := util.PostJSON(qyUrl, req, c.ProxyUrl)

	if err != nil {
		return
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("GetQyCustomerViewBatch error : errcode=%d , errmsg=%s", result.ErrCode, result.ErrMsg)
	}
	return
}

//修改客户备注信息
type UpdateCusRemarkReq struct {
	Userid  string 	`json:"userid"`
	ExternalUserid string `json:"external_userid"`
	Remark string `json:"remark"`
	Description string `json:"description"`
	RemarkCompany string `json:"remark_company"`
	RemarkMobiles []string `json:"remark_mobiles"`
	RemarkPicMediaid string `json:"remark_pic_mediaid"`
}
func (c *Customer) UpdateQyCustomerRemark(accessToken string, req UpdateCusRemarkReq)(result *util.WxError, err error){
	qyUrl := fmt.Sprintf(UpdateQyCustomerRemarkURL, accessToken)

	response, err := util.PostJSON(qyUrl, req, c.ProxyUrl)

	if err != nil {
		return
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("UpdateQyCustomerRemark error : errcode=%d , errmsg=%s", result.ErrCode, result.ErrMsg)
	}
	return
}

//获取规则组列表
type CustomerStrategyReq struct {
	Cursor string `json:"cursor"`
	Limit  int32  `json:"limit"`
}
type CustomerStrategyRep struct {
	util.WxError
	Strategy []struct {
		StrategyId int32 `json:"strategy_id"`
	} `json:"strategy"`
	NextCursor string `json:"next_cursor"`
}
func (c *Customer) GetQyCustomerStrategyList(accessToken string, req CustomerStrategyReq)(result *CustomerStrategyRep, err error){
	qyUrl := fmt.Sprintf(GetQyCustomerStrategyListURL, accessToken)

	response, err := util.PostJSON(qyUrl, req, c.ProxyUrl)

	if err != nil {
		return
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("GetQyCustomerStrategyList error : errcode=%d , errmsg=%s", result.ErrCode, result.ErrMsg)
	}
	return
}

//获取规则组详情
type CustomerStrategyViewReq struct {
	StrategyId int `json:"strategy_id"`
}
type CustomerStrategyViewRep struct {
	util.WxError
	Strategy struct {
		StrategyId   int32      `json:"strategy_id"`
		ParentId     int32      `json:"parent_id"`
		StrategyName string     `json:"strategy_name"`
		CreateTime   int32      `json:"create_time"`
		AdminList    []string   `json:"admin_list"`
		Privilege    struct {
			ViewCustomerList        bool        `json:"view_customer_list"`
			ViewCustomerData        bool        `json:"view_customer_data"`
			ViewRoomList            bool        `json:"view_room_list"`
			ContactMe               bool        `json:"contact_me"`
			JoinRoom                bool        `json:"join_room"`
			ShareCustomer           bool        `json:"share_customer"`
			OperResignCustomer      bool        `json:"oper_resign_customer"`
			OperResignGroup         bool        `json:"oper_resign_group"`
			SendCustomerMsg         bool        `json:"send_customer_msg"`
			EditWelcomeMsg          bool        `json:"edit_welcome_msg"`
			ViewBehaviorData        bool        `json:"view_behavior_data"`
			ViewRoomData            bool        `json:"view_room_data"`
			SendGroupMsg            bool        `json:"send_group_msg"`
			RoomDeduplication       bool        `json:"room_deduplication"`
			RapidReply              bool        `json:"rapid_reply"`
			OnjobCustomerTransfer   bool        `json:"onjob_customer_transfer"`
			EditAntiSpamRule        bool        `json:"edit_anti_spam_rule"`
			ExportCustomerList      bool        `json:"export_customer_list"`
			ExportCustomerData      bool        `json:"export_customer_data"`
			ExportCustomerGroupList bool        `json:"export_customer_group_list"`
		} `json:"privilege"`
	} `json:"strategy"`
}
func (c *Customer) GetQyCustomerStrategyView(accessToken string, req CustomerStrategyViewReq)(result *CustomerStrategyViewRep, err error){
	qyUrl := fmt.Sprintf(GetQyCustomerStrategyViewURL, accessToken)

	response, err := util.PostJSON(qyUrl, req, c.ProxyUrl)

	if err != nil {
		return
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("GetQyCustomerStrategyView error : errcode=%d , errmsg=%s", result.ErrCode, result.ErrMsg)
	}
	return
}

//获取规则组管理范围
type CustomerStrategyRangeReq struct {
	StrategyId int32    `json:"strategy_id"`
	Cursor     string   `json:"cursor"`
	Limit      int32    `json:"limit"`
}
type CustomerStrategyRangeRep struct {
	util.WxError
	Range   []struct {
		Type    int32  `json:"type"`
		Userid  string `json:"userid,omitempty"`
		Partyid int32  `json:"partyid,omitempty"`
	} `json:"range"`
	NextCursor string `json:"next_cursor"`
}
func (c *Customer) GetQyCustomerStrategyRange(accessToken string, req CustomerStrategyRangeReq)(result *CustomerStrategyRangeRep, err error){
	qyUrl := fmt.Sprintf(GetQyCustomerStrategyRangeURL, accessToken)

	response, err := util.PostJSON(qyUrl, req, c.ProxyUrl)

	if err != nil {
		return
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("GetQyCustomerStrategyRange error : errcode=%d , errmsg=%s", result.ErrCode, result.ErrMsg)
	}
	return
}

//创建新的规则组
type CreateCustomerStrategyReq struct {
	ParentId     int32      `json:"parent_id"`
	StrategyName string   `json:"strategy_name"`
	AdminList    []string `json:"admin_list"`
	Privilege    struct {
		ViewCustomerList        bool        `json:"view_customer_list"`
		ViewCustomerData        bool        `json:"view_customer_data"`
		ViewRoomList            bool        `json:"view_room_list"`
		ContactMe               bool        `json:"contact_me"`
		JoinRoom                bool        `json:"join_room"`
		ShareCustomer           bool        `json:"share_customer"`
		OperResignCustomer      bool        `json:"oper_resign_customer"`
		SendCustomerMsg         bool	    `json:"send_customer_msg"`
		EditWelcomeMsg          bool        `json:"edit_welcome_msg"`
		ViewBehaviorData        bool        `json:"view_behavior_data"`
		ViewRoomData            bool        `json:"view_room_data"`
		SendGroupMsg            bool        `json:"send_group_msg"`
		RoomDeduplication       bool        `json:"room_deduplication"`
		RapidReply              bool        `json:"rapid_reply"`
		OnjobCustomerTransfer   bool        `json:"onjob_customer_transfer"`
		EditAntiSpamRule        bool        `json:"edit_anti_spam_rule"`
		ExportCustomerList      bool        `json:"export_customer_list"`
		ExportCustomerData      bool        `json:"export_customer_data"`
		ExportCustomerGroupList bool        `json:"export_customer_group_list"`
	} `json:"privilege"`
	Range []struct {
		Type    int32    `json:"type"`
		Userid  string `json:"userid,omitempty"`
		Partyid int32   `json:"partyid,omitempty"`
	} `json:"range"`
}
type CreateCustomerStrategyRep struct {
	util.WxError
	StrategyId int32    `json:"strategy_id"`
}
func (c *Customer) CreateQyCustomerStrategy(accessToken string, req CreateCustomerStrategyReq)(result *CreateCustomerStrategyRep, err error){
	qyUrl := fmt.Sprintf(CreateQyCustomerStrategyURL, accessToken)

	response, err := util.PostJSON(qyUrl, req, c.ProxyUrl)

	if err != nil {
		return
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("CreateQyCustomerStrategy error : errcode=%d , errmsg=%s", result.ErrCode, result.ErrMsg)
	}
	return
}

//编辑规则组及其管理范围
type UpdateQyCustomerStrategyReq struct {
	StrategyId   int32      `json:"strategy_id"`
	StrategyName string   `json:"strategy_name"`
	AdminList    []string `json:"admin_list"`
	Privilege    struct {
		ViewCustomerList        bool        `json:"view_customer_list"`
		ViewCustomerData        bool        `json:"view_customer_data"`
		ViewRoomList            bool        `json:"view_room_list"`
		ContactMe               bool        `json:"contact_me"`
		JoinRoom                bool        `json:"join_room"`
		ShareCustomer           bool        `json:"share_customer"`
		OperResignCustomer      bool        `json:"oper_resign_customer"`
		OperResignGroup         bool        `json:"oper_resign_group"`
		SendCustomerMsg         bool        `json:"send_customer_msg"`
		EditWelcomeMsg          bool        `json:"edit_welcome_msg"`
		ViewBehaviorData        bool        `json:"view_behavior_data"`
		ViewRoomData            bool        `json:"view_room_data"`
		SendGroupMsg            bool        `json:"send_group_msg"`
		RoomDeduplication       bool        `json:"room_deduplication"`
		RapidReply              bool        `json:"rapid_reply"`
		OnjobCustomerTransfer   bool        `json:"onjob_customer_transfer"`
		EditAntiSpamRule        bool        `json:"edit_anti_spam_rule"`
		ExportCustomerList      bool        `json:"export_customer_list"`
		ExportCustomerData      bool        `json:"export_customer_data"`
		ExportCustomerGroupList bool        `json:"export_customer_group_list"`
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
func (c *Customer) UpdateQyCustomerStrategy(accessToken string, req UpdateQyCustomerStrategyReq)(result *util.WxError, err error){
	qyUrl := fmt.Sprintf(UpdateQyCustomerStrategyURL, accessToken)

	response, err := util.PostJSON(qyUrl, req, c.ProxyUrl)

	if err != nil {
		return
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("UpdateQyCustomerStrategy error : errcode=%d , errmsg=%s", result.ErrCode, result.ErrMsg)
	}
	return
}

//删除规则组
type DelQyCustomerStrategyReq struct {
	StrategyId int32 `json:"strategy_id"`
}
func (c *Customer) DelQyCustomerStrategy(accessToken string, req DelQyCustomerStrategyReq)(result *util.WxError, err error){
	qyUrl := fmt.Sprintf(DelQyCustomerStrategyURL, accessToken)

	response, err := util.PostJSON(qyUrl, req, c.ProxyUrl)

	if err != nil {
		return
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("DelQyCustomerStrategy error : errcode=%d , errmsg=%s", result.ErrCode, result.ErrMsg)
	}
	return
}
