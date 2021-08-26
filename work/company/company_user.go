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
type CreateUserReq struct {
	Userid         string `json:"userid"`
	Name           string `json:"name"`
	Alias          string `json:"alias"`
	Mobile         string `json:"mobile"`
	Department     []int32  `json:"department"`
	Order          []int32  `json:"order"`
	Position       string `json:"position"`
	Gender         string `json:"gender"`
	Email          string `json:"email"`
	IsLeaderInDept []int32  `json:"is_leader_in_dept"`
	Enable         int32    `json:"enable"`
	AvatarMediaid  string `json:"avatar_mediaid"`
	Telephone      string `json:"telephone"`
	Address        string `json:"address"`
	MainDepartment int32    `json:"main_department"`
	Extattr        struct {
		Attrs []struct {
			Type int32    `json:"type"`
			Name string `json:"name"`
			Text struct {
				Value string `json:"value"`
			} `json:"text,omitempty"`
			Web struct {
				Url   string `json:"url"`
				Title string `json:"title"`
			} `json:"web,omitempty"`
		} `json:"attrs"`
	} `json:"extattr"`
	ToInvite         bool   `json:"to_invite"`
	ExternalPosition string `json:"external_position"`
	ExternalProfile  struct {
		ExternalCorpName string `json:"external_corp_name"`
		WechatChannels   struct {
			Nickname string `json:"nickname"`
		} `json:"wechat_channels"`
		ExternalAttr []struct {
			Type int32    `json:"type"`
			Name string `json:"name"`
			Text struct {
				Value string `json:"value"`
			} `json:"text,omitempty"`
			Web struct {
				Url   string `json:"url"`
				Title string `json:"title"`
			} `json:"web,omitempty"`
			Miniprogram struct {
				Appid    string `json:"appid"`
				Pagepath string `json:"pagepath"`
				Title    string `json:"title"`
			} `json:"miniprogram,omitempty"`
		} `json:"external_attr"`
	} `json:"external_profile"`
}
func (du *DepartmentUser) CreateUser (accessToken string, req CreateUserReq) (result *util.WxError, err error) {
	qyUrl := fmt.Sprintf(CreateUserURL, accessToken)

	response, err := util.PostJSON(qyUrl, req, du.ProxyUrl)
	if err != nil {
		return
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("CreateUserURL error : errcode=%d , errmsg=%s", result.ErrCode, result.ErrMsg)
	}
	return
}

//获取用户
type UserView struct {
	util.WxError
	Userid         string `json:"userid"`
	Name           string `json:"name"`
	Department     []int32  `json:"department"`
	Order          []int32  `json:"order"`
	Position       string `json:"position"`
	Mobile         string `json:"mobile"`
	Gender         string `json:"gender"`
	Email          string `json:"email"`
	IsLeaderInDept []int32  `json:"is_leader_in_dept"`
	Avatar         string `json:"avatar"`
	ThumbAvatar    string `json:"thumb_avatar"`
	Telephone      string `json:"telephone"`
	Alias          string `json:"alias"`
	Address        string `json:"address"`
	OpenUserid     string `json:"open_userid"`
	MainDepartment int32    `json:"main_department"`
	Extattr        struct {
		Attrs []struct {
			Type int32    `json:"type"`
			Name string `json:"name"`
			Text struct {
				Value string `json:"value"`
			} `json:"text,omitempty"`
			Web struct {
				Url   string `json:"url"`
				Title string `json:"title"`
			} `json:"web,omitempty"`
		} `json:"attrs"`
	} `json:"extattr"`
	Status           int32    `json:"status"`
	QrCode           string `json:"qr_code"`
	ExternalPosition string `json:"external_position"`
	ExternalProfile  struct {
		ExternalCorpName string `json:"external_corp_name"`
		WechatChannels   struct {
			Nickname string `json:"nickname"`
			Status   int32    `json:"status"`
		} `json:"wechat_channels"`
		ExternalAttr []struct {
			Type int32    `json:"type"`
			Name string `json:"name"`
			Text struct {
				Value string `json:"value"`
			} `json:"text,omitempty"`
			Web struct {
				Url   string `json:"url"`
				Title string `json:"title"`
			} `json:"web,omitempty"`
			Miniprogram struct {
				Appid    string `json:"appid"`
				Pagepath string `json:"pagepath"`
				Title    string `json:"title"`
			} `json:"miniprogram,omitempty"`
		} `json:"external_attr"`
	} `json:"external_profile"`
}
func (du *DepartmentUser) GetUser (accessToken string, userid string) (result *UserView, err error) {
	qyUrl := fmt.Sprintf(GetUserURL, accessToken, userid)

	response, err := util.HTTPGet(qyUrl, du.ProxyUrl)
	if err != nil {
		return
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("GetUser error : errcode=%d , errmsg=%s", result.ErrCode, result.ErrMsg)
	}
	return
}

//更新用户
type UpdateUserReq struct {
	Userid         string `json:"userid"`
	Name           string `json:"name"`
	Department     []int32  `json:"department"`
	Order          []int32  `json:"order"`
	Position       string `json:"position"`
	Mobile         string `json:"mobile"`
	Gender         string `json:"gender"`
	Email          string `json:"email"`
	IsLeaderInDept []int32  `json:"is_leader_in_dept"`
	Enable         int32    `json:"enable"`
	AvatarMediaid  string `json:"avatar_mediaid"`
	Telephone      string `json:"telephone"`
	Alias          string `json:"alias"`
	Address        string `json:"address"`
	MainDepartment int32    `json:"main_department"`
	Extattr        struct {
		Attrs []struct {
			Type int32    `json:"type"`
			Name string `json:"name"`
			Text struct {
				Value string `json:"value"`
			} `json:"text,omitempty"`
			Web struct {
				Url   string `json:"url"`
				Title string `json:"title"`
			} `json:"web,omitempty"`
		} `json:"attrs"`
	} `json:"extattr"`
	ExternalPosition string `json:"external_position"`
	ExternalProfile  struct {
		ExternalCorpName string `json:"external_corp_name"`
		WechatChannels   struct {
			Nickname string `json:"nickname"`
		} `json:"wechat_channels"`
		ExternalAttr []struct {
			Type int32    `json:"type"`
			Name string `json:"name"`
			Text struct {
				Value string `json:"value"`
			} `json:"text,omitempty"`
			Web struct {
				Url   string `json:"url"`
				Title string `json:"title"`
			} `json:"web,omitempty"`
			Miniprogram struct {
				Appid    string `json:"appid"`
				Pagepath string `json:"pagepath"`
				Title    string `json:"title"`
			} `json:"miniprogram,omitempty"`
		} `json:"external_attr"`
	} `json:"external_profile"`
}
func (du *DepartmentUser) UpdateUser (accessToken string, req UpdateUserReq) (result *util.WxError, err error) {
	qyUrl := fmt.Sprintf(UpdateUserURL, accessToken)

	response, err := util.PostJSON(qyUrl, req, du.ProxyUrl)
	if err != nil {
		return
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("UpdateUser error : errcode=%d , errmsg=%s", result.ErrCode, result.ErrMsg)
	}
	return
}

//删除用户
func (du *DepartmentUser) DelUser (accessToken string, userid string) (result *UserView, err error) {
	qyUrl := fmt.Sprintf(DelUserURL, accessToken, userid)

	response, err := util.HTTPGet(qyUrl, du.ProxyUrl)
	if err != nil {
		return
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("DelUser error : errcode=%d , errmsg=%s", result.ErrCode, result.ErrMsg)
	}
	return
}

//批量删除成员
type DelUserBatchReq struct {
	Useridlist []string `json:"useridlist"`
}
func (du *DepartmentUser) DelUserBatch (accessToken string, req DelUserBatchReq) (result *util.WxError, err error) {
	qyUrl := fmt.Sprintf(DelUserBatchURL, accessToken)

	response, err := util.PostJSON(qyUrl, req, du.ProxyUrl)
	if err != nil {
		return
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("DelUserBatch error : errcode=%d , errmsg=%s", result.ErrCode, result.ErrMsg)
	}
	return
}

//获取部门成员
type SimpleUserList struct {
	util.WxError
	Userlist []struct {
		Userid     string `json:"userid"`
		Name       string `json:"name"`
		Department []int  `json:"department"`
		OpenUserid string `json:"open_userid"`
	} `json:"userlist"`
}
func (du *DepartmentUser) GetDepartmentSimpleUserList (accessToken string, departmentId int32, fetchChild int32) (result *SimpleUserList, err error) {
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
type UserList struct {
	util.WxError
	Userlist []struct {
		Userid         string `json:"userid"`
		Name           string `json:"name"`
		Department     []int32  `json:"department"`
		Order          []int32  `json:"order"`
		Position       string `json:"position"`
		Mobile         string `json:"mobile"`
		Gender         string `json:"gender"`
		Email          string `json:"email"`
		IsLeaderInDept []int32  `json:"is_leader_in_dept"`
		Avatar         string `json:"avatar"`
		ThumbAvatar    string `json:"thumb_avatar"`
		Telephone      string `json:"telephone"`
		Alias          string `json:"alias"`
		Status         int32    `json:"status"`
		Address        string `json:"address"`
		HideMobile     int32    `json:"hide_mobile"`
		EnglishName    string `json:"english_name"`
		OpenUserid     string `json:"open_userid"`
		MainDepartment int32    `json:"main_department"`
		Extattr        struct {
			Attrs []struct {
				Type int32    `json:"type"`
				Name string `json:"name"`
				Text struct {
					Value string `json:"value"`
				} `json:"text,omitempty"`
				Web struct {
					Url   string `json:"url"`
					Title string `json:"title"`
				} `json:"web,omitempty"`
			} `json:"attrs"`
		} `json:"extattr"`
		QrCode           string `json:"qr_code"`
		ExternalPosition string `json:"external_position"`
		ExternalProfile  struct {
			ExternalCorpName string `json:"external_corp_name"`
			ExternalAttr     []struct {
				Type int32    `json:"type"`
				Name string `json:"name"`
				Text struct {
					Value string `json:"value"`
				} `json:"text,omitempty"`
				Web struct {
					Url   string `json:"url"`
					Title string `json:"title"`
				} `json:"web,omitempty"`
				Miniprogram struct {
					Appid    string `json:"appid"`
					Pagepath string `json:"pagepath"`
					Title    string `json:"title"`
				} `json:"miniprogram,omitempty"`
			} `json:"external_attr"`
		} `json:"external_profile"`
	} `json:"userlist"`
}
func (du *DepartmentUser) GetDepartmentUserList (accessToken string, departmentId int32, fetchChild int32) (result *UserList, err error) {
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
type ConvertToOpenidReq struct {
	Userid string `json:"userid"`
}
type ConvertToOpenidRep struct {
	util.WxError
	Openid  string `json:"openid"`
}
func (du *DepartmentUser) ConvertToOpenid (accessToken string, req ConvertToOpenidReq) (result *ConvertToOpenidRep, err error) {
	qyUrl := fmt.Sprintf(ConvertToOpenidURL, accessToken)

	response, err := util.PostJSON(qyUrl, req, du.ProxyUrl)
	if err != nil {
		return
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("ConvertToOpenid error : errcode=%d , errmsg=%s", result.ErrCode, result.ErrMsg)
	}
	return
}

//二次验证
func (du *DepartmentUser) AuthSucc (accessToken string, userid string) (result *util.WxError, err error) {
	qyUrl := fmt.Sprintf(AuthSuccURL, accessToken, userid)

	response, err := util.HTTPGet(qyUrl, du.ProxyUrl)
	if err != nil {
		return
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("AuthSucc error : errcode=%d , errmsg=%s", result.ErrCode, result.ErrMsg)
	}
	return
}

//邀请成员
type BatchInviteReq struct {
	User  []string `json:"user"`
	Party []int32 `json:"party"`
	Tag   []int32 `json:"tag"`
}
type BatchInviteRep struct {
	util.WxError
	Invaliduser  []string `json:"invaliduser"`
	Invalidparty []int32 `json:"invalidparty"`
	Invalidtag   []int32 `json:"invalidtag"`
}
func (du *DepartmentUser) BatchInvite (accessToken string, req BatchInviteReq) (result *BatchInviteRep, err error) {
	qyUrl := fmt.Sprintf(BatchInviteURL, accessToken)

	response, err := util.PostJSON(qyUrl, req, du.ProxyUrl)
	if err != nil {
		return
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("BatchInvite error : errcode=%d , errmsg=%s", result.ErrCode, result.ErrMsg)
	}
	return
}

//获取加入企业二维码
type JoinCorpQrcodeRep struct {
	util.WxError
	JoinQrcode string `json:"join_qrcode"`
}
func (du *DepartmentUser) GetJoinCorpQrcode (accessToken string, sizeType int32) (result *JoinCorpQrcodeRep, err error) {
	qyUrl := fmt.Sprintf(GetJoinCorpQrcodeURL, accessToken, sizeType)

	response, err := util.HTTPGet(qyUrl, du.ProxyUrl)
	if err != nil {
		return
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("GetJoinCorpQrcode error : errcode=%d , errmsg=%s", result.ErrCode, result.ErrMsg)
	}
	return
}

//获取企业活跃成员数
type ActiveStatReq struct {
	Date string `json:"date"`
}
type ActiveStatRep struct {
	util.WxError
	ActiveCnt int32  `json:"active_cnt"`
}
func (du *DepartmentUser) GetActiveStat (accessToken string, req ActiveStatReq) (result *ActiveStatRep, err error) {
	qyUrl := fmt.Sprintf(GetActiveStatURL, accessToken)

	response, err := util.PostJSON(qyUrl, req, du.ProxyUrl)
	if err != nil {
		return
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("GetActiveStat error : errcode=%d , errmsg=%s", result.ErrCode, result.ErrMsg)
	}
	return
}
