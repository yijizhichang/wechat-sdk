//通讯录管理-成员管理
package company

import (
	"encoding/json"
	"fmt"
	"github.com/yijizhichang/wechat-sdk/util"
	"github.com/yijizhichang/wechat-sdk/work/core"
)

const (
	CreateUserURL = "https://qyapi.weixin.qq.com/cgi-bin/user/create?access_token=%s"  //创建成员
	GetUserURL = "https://qyapi.weixin.qq.com/cgi-bin/user/get?access_token=%s&userid=%s"  //获取成员
	UpdateUserURL = "https://qyapi.weixin.qq.com/cgi-bin/user/update?access_token=%s"  //更新成员
	DelUserURL = "https://qyapi.weixin.qq.com/cgi-bin/user/delete?access_token=%s&userid=%s"  //删除成员
	DelUserBatchURL = "https://qyapi.weixin.qq.com/cgi-bin/user/batchdelete?access_token=%s"  //批量删除成员
	GetUserSimpleListURL = "https://qyapi.weixin.qq.com/cgi-bin/user/simplelist?access_token=%s&department_id=%d&fetch_child=%d"   //获取部门成员
	GetUserListURL = "https://qyapi.weixin.qq.com/cgi-bin/user/list?access_token=%s&department_id=%d&fetch_child=%d"   //获取部门成员详情
	ConvertToOpenidURL = "https://qyapi.weixin.qq.com/cgi-bin/user/convert_to_openid?access_token=%s"   //userid与openid互换
	AuthSuccURL = "https://qyapi.weixin.qq.com/cgi-bin/user/authsucc?access_token=%s&userid=%s" //二次验证
	BatchInviteURL = "https://qyapi.weixin.qq.com/cgi-bin/batch/invite?access_token=%s" //邀请成员
	GetJoinCorpQrcodeURL = "https://qyapi.weixin.qq.com/cgi-bin/corp/get_join_qrcode?access_token=%s&size_type=%d"  //获取加入企业二维码
	GetActiveStatURL = "https://qyapi.weixin.qq.com/cgi-bin/user/get_active_stat?access_token=%s"  //获取企业活跃成员数
)

//Department 客户管理
type DepartmentUser struct {
	*core.Context
}

//NewDepartmentUser 实例化
func NewDepartmentUser(context *core.Context) *DepartmentUser {
	depu := new(DepartmentUser)
	depu.Context = context
	return depu
}

//创建用户

//获取用户

//更新用户

//删除用户

//批量删除成员

//获取部门成员
type simpleUserList struct {
	util.WxError
	Userlist [] userSimple `json:"userlist"`
}
type userSimple struct {
	Userid string `json:"userid"`
	Name string `json:"name"`
	Department []int32 `json:"department"`
	OpenUserid string `json:"open_userid"`
}
func (du *DepartmentUser) GetDepartmentSimpleUserList (accessToken string, departmentId int32, fetchChild int32) (result *simpleUserList, err error) {
	qyUrl := fmt.Sprintf(GetUserSimpleListURL, accessToken, departmentId, fetchChild)

	response, err := util.HTTPGet(qyUrl, du.ProxyUrl)
	if err != nil {
		return
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("GetDepartmentSimpleUserList error : errcode=%d , errmsg=%s", result.ErrCode, result.ErrMsg)
	}
	return
}

//获取部门成员详情
type userList struct {
	util.WxError
	Userlist [] userItem `json:"userlist"`
}
type userItem struct {
	Userid string `json:"userid"`
	Name string `json:"name"`
	Department []int32 `json:"department"`
	Order []int32 `json:"order"`
	Position string `json:"position"`
	Mobile string `json:"mobile"`
	Gender string `json:"gender"`
	Email string `json:"email"`
	IsLeaderInDept []int32 `json:"is_leader_in_dept"`
	Avatar string `json:"avatar"`
	ThumbAvatar string `json:"thumb_avatar"`
	Telephone string `json:"telephone"`
	Alias string `json:"alias"`
	Status int32 `json:"status"`
	Address string `json:"address"`
	HideMobile int32 `json:"hide_mobile"`
	EnglishName string `json:"english_name"`
	OpenUserid string `json:"open_userid"`
	MainDepartment int32 `json:"main_department"`
	Extattr struct{
		Attrs []attrItem `json:"attrs"`
	} `json:"extattr"`
	QrCode string `json:"qr_code"`
	ExternalPosition string `json:"external_position"`
	ExternalProfile struct{
		ExternalCorpName string `json:"external_corp_name"`
		ExternalAttr []attrItem `json:"external_attr"`
	} `json:"external_profile"`
}

type attrItem struct {
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
func (du *DepartmentUser) GetDepartmentUserList (accessToken string, departmentId int32, fetchChild int32) (result *userList, err error) {
	qyUrl := fmt.Sprintf(GetUserListURL, accessToken, departmentId, fetchChild)

	response, err := util.HTTPGet(qyUrl, du.ProxyUrl)
	if err != nil {
		return
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("GetDepartmentUserList error : errcode=%d , errmsg=%s", result.ErrCode, result.ErrMsg)
	}
	return
}

//userid与openid互换

//二次验证

//邀请成员

//获取加入企业二维码

//获取企业活跃成员数