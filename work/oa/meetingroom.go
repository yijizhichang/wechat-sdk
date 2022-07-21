/**
 * @Time : 2022/7/21 10:13
 * @Author : soupzhb@gmail.com
 * @File : meetingroom.go
 * @Software: GoLand
 */

package oa

import (
	"encoding/json"
	"fmt"
	"github.com/yijizhichang/wechat-sdk/util"
	"github.com/yijizhichang/wechat-sdk/work/core"
)

const (
	CreateQyMeetingRoomURL         = "https://qyapi.weixin.qq.com/cgi-bin/oa/meetingroom/add?access_token=%s"                            //添加会议室
	GetQyMeetingRoomListURL        = "https://qyapi.weixin.qq.com/cgi-bin/oa/meetingroom/list?access_token=%s"                           //查询会议室
	EditQyMeetingRoomURL           = "https://qyapi.weixin.qq.com/cgi-bin/oa/meetingroom/edit?access_token=%s"                           //编辑会议室
	DelQyMeetingRoomURL            = "https://qyapi.weixin.qq.com/cgi-bin/oa/meetingroom/del?access_token=%s"                            //删除会议室
	GetQyMeetingRoomBookingListURL = "https://qyapi.weixin.qq.com/cgi-bin/oa/meetingroom/get_booking_info?access_token=%s"               //查询会议室的预定信息
	CreateQyMeetingRoomBookURL     = "https://qyapi.weixin.qq.com/cgi-bin/oa/meetingroom/book?access_token=%s"                           //预定会议室
	CancelQyMeetingRoomBookURL     = "https://qyapi.weixin.qq.com/cgi-bin/oa/meetingroom/cancel_book?access_token=%s"                    //取消预定会议室
	GetQyMeetingRoomBookingInfoURL = "https://qyapi.weixin.qq.com/cgi-bin/oa/meetingroom/get_booking_info_by_meeting_id?access_token=%s" //根据会议ID查询会议室的预定信息

)

//MeetingRoom 会议室管理
type MeetingRoom struct {
	*core.Context
}

//NewMeetingRoom 实例化
func NewMeetingRoom(context *core.Context) *MeetingRoom {
	mrm := new(MeetingRoom)
	mrm.Context = context
	return mrm
}

//添加会议室
type CreateMeetingRoomReq struct {
	Name       string `json:"name"`
	Capacity   int    `json:"capacity"`
	City       string `json:"city"`
	Building   string `json:"building"`
	Floor      string `json:"floor"`
	Equipment  []int  `json:"equipment"`
	Coordinate struct {
		Latitude  string `json:"latitude"`
		Longitude string `json:"longitude"`
	} `json:"coordinate"`
}

type CreateMeetingRoomRep struct {
	util.WxError
	MeetingroomId int `json:"meetingroom_id"`
}

func (mrm *MeetingRoom) CreateQyMeetingRoom(accessToken string, req CreateMeetingRoomReq) (result *CreateMeetingRoomRep, err error) {
	qyUrl := fmt.Sprintf(CreateQyMeetingRoomURL, accessToken)

	response, err := util.PostJSON(qyUrl, req, mrm.ProxyUrl)
	if err != nil {
		return
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("CreateQyMeetingRoom error : errcode=%d , errmsg=%s", result.ErrCode, result.ErrMsg)
	}
	return
}

//查询会议室
type GetQyMeetingRoomListReq struct {
	City      string `json:"city"`
	Building  string `json:"building"`
	Floor     string `json:"floor"`
	Equipment []int  `json:"equipment"`
}
type GetQyMeetingRoomListRep struct {
	util.WxError
	MeetingroomList []struct {
		MeetingroomId int    `json:"meetingroom_id"`
		Name          string `json:"name"`
		Capacity      int    `json:"capacity"`
		City          string `json:"city"`
		Building      string `json:"building"`
		Floor         string `json:"floor"`
		Equipment     []int  `json:"equipment"`
		Coordinate    struct {
			Latitude  string `json:"latitude"`
			Longitude string `json:"longitude"`
		} `json:"coordinate"`
		NeedApproval int `json:"need_approval"`
	} `json:"meetingroom_list"`
}

func (mrm *MeetingRoom) GetQyMeetingRoomList(accessToken string, req GetQyMeetingRoomListReq) (result *GetQyMeetingRoomListRep, err error) {
	qyUrl := fmt.Sprintf(GetQyMeetingRoomListURL, accessToken)

	response, err := util.PostJSON(qyUrl, req, mrm.ProxyUrl)
	if err != nil {
		return
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("GetQyMeetingRoomList error : errcode=%d , errmsg=%s", result.ErrCode, result.ErrMsg)
	}
	return
}

//编辑会议室
type EditQyMeetingRoomReq struct {
	MeetingroomId int    `json:"meetingroom_id"`
	Name          string `json:"name"`
	Capacity      int    `json:"capacity"`
	City          string `json:"city"`
	Building      string `json:"building"`
	Floor         string `json:"floor"`
	Equipment     []int  `json:"equipment"`
	Coordinate    struct {
		Latitude  string `json:"latitude"`
		Longitude string `json:"longitude"`
	} `json:"coordinate"`
}

func (mrm *MeetingRoom) EditQyMeetingRoom(accessToken string, req EditQyMeetingRoomReq) (result util.WxError, err error) {
	qyUrl := fmt.Sprintf(EditQyMeetingRoomURL, accessToken)

	response, err := util.PostJSON(qyUrl, req, mrm.ProxyUrl)
	if err != nil {
		return
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("EditQyMeetingRoom error : errcode=%d , errmsg=%s", result.ErrCode, result.ErrMsg)
	}
	return
}

//删除会议室
type DelQyMeetingRoomReq struct {
	MeetingroomId int `json:"meetingroom_id"`
}

func (mrm *MeetingRoom) DelQyMeetingRoom(accessToken string, req DelQyMeetingRoomReq) (result util.WxError, err error) {
	qyUrl := fmt.Sprintf(DelQyMeetingRoomURL, accessToken)

	response, err := util.PostJSON(qyUrl, req, mrm.ProxyUrl)
	if err != nil {
		return
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("DelQyMeetingRoom error : errcode=%d , errmsg=%s", result.ErrCode, result.ErrMsg)
	}
	return
}

//查询会议室的预定信息
type GetQyMeetingRoomBookingListReq struct {
	MeetingroomId int    `json:"meetingroom_id"`
	StartTime     int    `json:"start_time"`
	EndTime       int    `json:"end_time"`
	City          string `json:"city"`
	Building      string `json:"building"`
	Floor         string `json:"floor"`
}

type GetQyMeetingRoomBookingListRep struct {
	util.WxError
	BookingList []struct {
		MeetingroomId int `json:"meetingroom_id"`
		Schedule      []struct {
			MeetingId  string `json:"meeting_id"`
			ScheduleId string `json:"schedule_id"`
			StartTime  int    `json:"start_time"`
			EndTime    int    `json:"end_time"`
			Booker     string `json:"booker"`
			Status     int    `json:"status"`
		} `json:"schedule"`
	} `json:"booking_list"`
}

func (mrm *MeetingRoom) GetQyMeetingRoomBookingList(accessToken string, req GetQyMeetingRoomBookingListReq) (result *GetQyMeetingRoomBookingListRep, err error) {
	qyUrl := fmt.Sprintf(GetQyMeetingRoomBookingListURL, accessToken)

	response, err := util.PostJSON(qyUrl, req, mrm.ProxyUrl)
	if err != nil {
		return
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("GetQyMeetingRoomBookingList error : errcode=%d , errmsg=%s", result.ErrCode, result.ErrMsg)
	}
	return
}

//预定会议室
type CreateQyMeetingRoomBookReq struct {
	MeetingroomId int      `json:"meetingroom_id"`
	Subject       string   `json:"subject"`
	StartTime     int      `json:"start_time"`
	EndTime       int      `json:"end_time"`
	Booker        string   `json:"booker"`
	Attendees     []string `json:"attendees"`
}
type CreateQyMeetingRoomBookRep struct {
	util.WxError
	MeetingId  string `json:"meeting_id"`
	ScheduleId string `json:"schedule_id"`
}

func (mrm *MeetingRoom) CreateQyMeetingRoomBook(accessToken string, req CreateQyMeetingRoomBookReq) (result *CreateQyMeetingRoomBookRep, err error) {
	qyUrl := fmt.Sprintf(CreateQyMeetingRoomBookURL, accessToken)

	response, err := util.PostJSON(qyUrl, req, mrm.ProxyUrl)
	if err != nil {
		return
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("CreateQyMeetingRoomBook error : errcode=%d , errmsg=%s", result.ErrCode, result.ErrMsg)
	}
	return
}

//取消预定会议室
type CancelQyMeetingRoomBookReq struct {
	MeetingId    string `json:"meeting_id"`
	KeepSchedule int    `json:"keep_schedule"`
}

func (mrm *MeetingRoom) CancelQyMeetingRoomBook(accessToken string, req CancelQyMeetingRoomBookReq) (result util.WxError, err error) {
	qyUrl := fmt.Sprintf(CancelQyMeetingRoomBookURL, accessToken)

	response, err := util.PostJSON(qyUrl, req, mrm.ProxyUrl)
	if err != nil {
		return
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("CancelQyMeetingRoomBook error : errcode=%d , errmsg=%s", result.ErrCode, result.ErrMsg)
	}
	return
}

//根据会议ID查询会议室的预定信息
type GetQyMeetingRoomBookingInfoReq struct {
	MeetingroomId int    `json:"meetingroom_id"`
	MeetingId     string `json:"meeting_id"`
}
type GetQyMeetingRoomBookingInfoRep struct {
	util.WxError
	MeetingroomId int `json:"meetingroom_id"`
	Schedule      struct {
		MeetingId  string `json:"meeting_id"`
		ScheduleId string `json:"schedule_id"`
		StartTime  int    `json:"start_time"`
		EndTime    int    `json:"end_time"`
		Booker     string `json:"booker"`
		Status     int    `json:"status"`
	} `json:"schedule"`
}

func (mrm *MeetingRoom) GetQyMeetingRoomBookingInfo(accessToken string, req GetQyMeetingRoomBookingInfoReq) (result *GetQyMeetingRoomBookingInfoRep, err error) {
	qyUrl := fmt.Sprintf(GetQyMeetingRoomBookingInfoURL, accessToken)

	response, err := util.PostJSON(qyUrl, req, mrm.ProxyUrl)
	if err != nil {
		return
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("GetQyMeetingRoomBookingInfo error : errcode=%d , errmsg=%s", result.ErrCode, result.ErrMsg)
	}
	return
}
