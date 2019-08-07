//图文消息留言管理接口
package media

import (
	"fmt"
	"github.com/yijizhichang/wechat-sdk/util"
	"encoding/json"
)

const (
	OpenCommentURL       	 = "https://api.weixin.qq.com/cgi-bin/comment/open?access_token=%s"          //打开已群发文章评论
	CloseCommentURL          = "https://api.weixin.qq.com/cgi-bin/comment/close?access_token=%s"         //关闭已群发文章评论
	GetCommentListURL        = "https://api.weixin.qq.com/cgi-bin/comment/list?access_token=%s"   		 //查看指定文章的评论数据
	MarkElectCommentURL 	 = "https://api.weixin.qq.com/cgi-bin/comment/markelect?access_token=%s"     //将评论标记精选
	UnMarkElectCommentURL    = "https://api.weixin.qq.com/cgi-bin/comment/unmarkelect?access_token=%s"   //将评论取消精选
	DeleteCommentURL     	 = "https://api.weixin.qq.com/cgi-bin/comment/delete?access_token=%s" 		 //删除评论
	ReplyCommentURL      	 = "https://api.weixin.qq.com/cgi-bin/comment/reply/add?access_token=%s"     //回复评论
	DeleteReplyCommentURL    = "https://api.weixin.qq.com/cgi-bin/comment/reply/delete?access_token=%s"  //删除回复
)


//打开已群发文章评论
type reqOpenComment struct{
	MsgDataId		uint32		`json:"msg_data_id"`    //群发返回的msg_data_id
	Index 			uint32		`json:"index"`          //多图文时，用来指定第几篇图文，从0开始，不带默认操作该msg_data_id的第一篇图文
}

func (media *Media) OpenComment(msgDataId, index uint32) (result util.WxError, err error)  {
	accessToken, err := media.GetAccessToken()
	if err != nil {
		return
	}
	wxUrl := fmt.Sprintf(OpenCommentURL, accessToken)

	postData := new(reqOpenComment)
	postData.MsgDataId = msgDataId
	postData.Index = index

	response, err := util.PostJSON(wxUrl, postData, media.ProxyUrl)

	if err != nil {
		return
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("OpenComment error : errcode=%v , errmsg=%v", result.ErrCode, result.ErrMsg)
	}
	return
}



//关闭已群发文章评论
type reqCloseComment struct{
	MsgDataId		uint32		`json:"msg_data_id"`    //群发返回的msg_data_id
	Index 			uint32		`json:"index"`          //多图文时，用来指定第几篇图文，从0开始，不带默认操作该msg_data_id的第一篇图文
}

func (media *Media) CloseComment(msgDataId, index uint32) (result util.WxError, err error)  {
	accessToken, err := media.GetAccessToken()
	if err != nil {
		return
	}
	wxUrl := fmt.Sprintf(CloseCommentURL, accessToken)

	postData := new(reqCloseComment)
	postData.MsgDataId = msgDataId
	postData.Index = index

	response, err := util.PostJSON(wxUrl, postData, media.ProxyUrl)

	if err != nil {
		return
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("CloseComment error : errcode=%v , errmsg=%v", result.ErrCode, result.ErrMsg)
	}
	return
}



//查看指定文章的评论数据
type reqCommentList struct{
	MsgDataId		uint32		`json:"msg_data_id"`    //群发返回的msg_data_id
	Index 			uint32		`json:"index"`          //多图文时，用来指定第几篇图文，从0开始，不带默认操作该msg_data_id的第一篇图文
	Begin 			uint32		`json:"begin"`			//起始位置
	Count           uint32		`json:"count"`			//获取数目（>=50会被拒绝）
	Type            uint32		`json:"type"`			//type=0 普通评论&精选评论 type=1 普通评论 type=2 精选评论
}

type commentList struct{
	util.WxError
	Total 	 uint32		`json:"total"`
	Comment  [] *comment	`json:"comment"`
}

type comment struct {
	UserCommentId	 uint32		`json:"user_comment_id"`
	CreateTime       uint32		`json:"create_time"`
	Content 		 string		`json:"content"`
	CommentType      uint32		`json:"comment_type"`
	Openid           string		`json:"openid"`
}

func (media *Media) GetCommentList(msgDataId, index, begin, count, commentType uint32) (result *commentList, err error)  {
	accessToken, err := media.GetAccessToken()
	if err != nil {
		return
	}
	wxUrl := fmt.Sprintf(GetCommentListURL, accessToken)

	postData := new(reqCommentList)
	postData.MsgDataId = msgDataId
	postData.Index = index
	postData.Begin = begin
	postData.Count = count
	postData.Type = commentType

	response, err := util.PostJSON(wxUrl, postData, media.ProxyUrl)

	if err != nil {
		return
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("GetCommentList error : errcode=%v , errmsg=%v", result.ErrCode, result.ErrMsg)
	}
	return
}



//将评论标记精选
type reqMarkElectComment struct{
	MsgDataId		uint32		`json:"msg_data_id"`    	//群发返回的msg_data_id
	Index 			uint32		`json:"index"`          	//多图文时，用来指定第几篇图文，从0开始，不带默认操作该msg_data_id的第一篇图文
	UserCommentId   uint32		`json:"user_comment_id"`  	//用户评论id
}

func (media *Media) MarkElectComment(msgDataId, index, userCommentId uint32) (result util.WxError, err error)  {
	accessToken, err := media.GetAccessToken()
	if err != nil {
		return
	}
	wxUrl := fmt.Sprintf(MarkElectCommentURL, accessToken)

	postData := new(reqMarkElectComment)
	postData.MsgDataId = msgDataId
	postData.Index = index
	postData.UserCommentId = userCommentId

	response, err := util.PostJSON(wxUrl, postData, media.ProxyUrl)

	if err != nil {
		return
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("MarkElectComment error : errcode=%v , errmsg=%v", result.ErrCode, result.ErrMsg)
	}
	return
}



//将评论取消精选
type reqUnMarkElectComment struct{
	MsgDataId		uint32		`json:"msg_data_id"`    	//群发返回的msg_data_id
	Index 			uint32		`json:"index"`          	//多图文时，用来指定第几篇图文，从0开始，不带默认操作该msg_data_id的第一篇图文
	UserCommentId   uint32		`json:"user_comment_id"`  	//用户评论id
}

func (media *Media) UnMarkElectComment(msgDataId, index, userCommentId uint32) (result util.WxError, err error)  {
	accessToken, err := media.GetAccessToken()
	if err != nil {
		return
	}
	wxUrl := fmt.Sprintf(UnMarkElectCommentURL, accessToken)

	postData := new(reqUnMarkElectComment)
	postData.MsgDataId = msgDataId
	postData.Index = index
	postData.UserCommentId = userCommentId

	response, err := util.PostJSON(wxUrl, postData, media.ProxyUrl)

	if err != nil {
		return
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("UnMarkElectComment error : errcode=%v , errmsg=%v", result.ErrCode, result.ErrMsg)
	}
	return
}



//删除评论
type reqDeleteComment struct{
	MsgDataId		uint32		`json:"msg_data_id"`    	//群发返回的msg_data_id
	Index 			uint32		`json:"index"`          	//多图文时，用来指定第几篇图文，从0开始，不带默认操作该msg_data_id的第一篇图文
	UserCommentId   uint32		`json:"user_comment_id"`  	//用户评论id
}

func (media *Media) DeleteComment(msgDataId, index, userCommentId uint32) (result util.WxError, err error)  {
	accessToken, err := media.GetAccessToken()
	if err != nil {
		return
	}
	wxUrl := fmt.Sprintf(DeleteCommentURL, accessToken)

	postData := new(reqDeleteComment)
	postData.MsgDataId = msgDataId
	postData.Index = index
	postData.UserCommentId = userCommentId

	response, err := util.PostJSON(wxUrl, postData, media.ProxyUrl)

	if err != nil {
		return
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("DeleteComment error : errcode=%v , errmsg=%v", result.ErrCode, result.ErrMsg)
	}
	return
}




//回复评论
type reqReplyComment struct{
	MsgDataId		uint32		`json:"msg_data_id"`    	//群发返回的msg_data_id
	Index 			uint32		`json:"index"`          	//多图文时，用来指定第几篇图文，从0开始，不带默认操作该msg_data_id的第一篇图文
	UserCommentId   uint32		`json:"user_comment_id"`  	//用户评论id
	Content         string		`json:"content"`			//回复内容
}

func (media *Media) ReplayComment(msgDataId, index, userCommentId uint32, content string) (result util.WxError, err error)  {
	accessToken, err := media.GetAccessToken()
	if err != nil {
		return
	}
	wxUrl := fmt.Sprintf(ReplyCommentURL, accessToken)

	postData := new(reqReplyComment)
	postData.MsgDataId = msgDataId
	postData.Index = index
	postData.UserCommentId = userCommentId
	postData.Content = content

	response, err := util.PostJSON(wxUrl, postData, media.ProxyUrl)

	if err != nil {
		return
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("ReplayComment error : errcode=%v , errmsg=%v", result.ErrCode, result.ErrMsg)
	}
	return
}



//删除回复
type reqDeleteReplyComment struct{
	MsgDataId		uint32		`json:"msg_data_id"`    	//群发返回的msg_data_id
	Index 			uint32		`json:"index"`          	//多图文时，用来指定第几篇图文，从0开始，不带默认操作该msg_data_id的第一篇图文
	UserCommentId   uint32		`json:"user_comment_id"`  	//用户评论id
}

func (media *Media) DeleteReplayComment(msgDataId, index, userCommentId uint32) (result util.WxError, err error)  {
	accessToken, err := media.GetAccessToken()
	if err != nil {
		return
	}
	wxUrl := fmt.Sprintf(DeleteReplyCommentURL, accessToken)

	postData := new(reqDeleteReplyComment)
	postData.MsgDataId = msgDataId
	postData.Index = index
	postData.UserCommentId = userCommentId

	response, err := util.PostJSON(wxUrl, postData, media.ProxyUrl)

	if err != nil {
		return
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("DeleteReplayComment error : errcode=%v , errmsg=%v", result.ErrCode, result.ErrMsg)
	}
	return
}

