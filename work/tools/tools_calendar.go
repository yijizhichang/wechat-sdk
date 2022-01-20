/**
 * @Time: 2022/1/19 6:00 下午
 * @Author: soupzhb@gmail.com
 * @File: tools_schedule.go
 * @Software: GoLand
 */

package tools

import (
	"encoding/json"
	"fmt"
	"github.com/yijizhichang/wechat-sdk/util"
	"github.com/yijizhichang/wechat-sdk/work/core"
)

const (
	QyCalendarCreateURL = "https://qyapi.weixin.qq.com/cgi-bin/oa/calendar/add?access_token=%s"  //创建日历
	QyCalendarUpdateURL = "https://qyapi.weixin.qq.com/cgi-bin/oa/calendar/update?access_token=%s" //更新日历
	QyCalendarGetURL = "https://qyapi.weixin.qq.com/cgi-bin/oa/calendar/get?access_token=%s" //获取日历详情
	QyCalendarDelURL = "https://qyapi.weixin.qq.com/cgi-bin/oa/calendar/del?access_token=%s" //删除日历
	QyScheduleCreateURL = "https://qyapi.weixin.qq.com/cgi-bin/oa/schedule/add?access_token=%s" //创建日程
	QyScheduleUpdateURL = "https://qyapi.weixin.qq.com/cgi-bin/oa/schedule/update?access_token=%s" //更新日程
	QyScheduleGetURL = "https://qyapi.weixin.qq.com/cgi-bin/oa/schedule/get?access_token=%s" //获取日程详情
	QyScheduleDelURL = "https://qyapi.weixin.qq.com/cgi-bin/oa/schedule/del?access_token=%s" //取消日程
	QyScheduleListURL = "https://qyapi.weixin.qq.com/cgi-bin/oa/schedule/get_by_calendar?access_token=%s" //获取日历下的日程列表
	QyScheduleAddAttendeesURL = "https://qyapi.weixin.qq.com/cgi-bin/oa/schedule/add_attendees?access_token=%s" //新增日程参与者
	QyScheduleDelAttendeesURL = "https://qyapi.weixin.qq.com/cgi-bin/oa/schedule/del_attendees?access_token=%s" //删除日程参与者


)

//Calendar 日历
type Calendar struct {
	*core.Context
}

//NewCalendar 实例化
func NewCalendar(context *core.Context) *Calendar {
	c := new(Calendar)
	c.Context = context
	return c
}

//创建日历
type CreateCalendarReq struct {
	Calendar struct {
		Organizer    string `json:"organizer"`
		Readonly     int64  `json:"readonly"`
		SetAsDefault int64  `json:"set_as_default"`
		Summary      string `json:"summary"`
		Color        string `json:"color"`
		Description  string `json:"description"`
		Shares       []struct {
			Userid   string `json:"userid"`
			Readonly int64  `json:"readonly,omitempty"`
		} `json:"shares"`
	} `json:"calendar"`
	Agentid int64 `json:"agentid"`
}
type CreateCalendarRep struct {
	util.WxError
	CalId string `json:"cal_id"`
}
func (l *Calendar) CreateCalendar(accessToken string, req CreateCalendarReq)(result *CreateCalendarRep, err error){
	qyUrl := fmt.Sprintf(QyCalendarCreateURL, accessToken)

	response, err := util.PostJSON(qyUrl, req, l.ProxyUrl)

	if err != nil {
		return
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("CreateCalendar error : errcode=%d , errmsg=%s", result.ErrCode, result.ErrMsg)
	}
	return
}

//更新日历
type UpdateCalendarReq struct {
	Calendar struct {
		CalId       string `json:"cal_id"`
		Readonly    int64  `json:"readonly"`
		Summary     string `json:"summary"`
		Color       string `json:"color"`
		Description string `json:"description"`
		Shares      []struct {
			Userid   string `json:"userid"`
			Readonly int64  `json:"readonly,omitempty"`
		} `json:"shares"`
	} `json:"calendar"`
}
type UpdateCalendarRep struct {
	util.WxError
}
func (l *Calendar) UpdateCalendar(accessToken string, req UpdateCalendarReq)(result *UpdateCalendarRep, err error){
	qyUrl := fmt.Sprintf(QyCalendarUpdateURL, accessToken)

	response, err := util.PostJSON(qyUrl, req, l.ProxyUrl)

	if err != nil {
		return
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("UpdateCalendar error : errcode=%d , errmsg=%s", result.ErrCode, result.ErrMsg)
	}
	return
}

//获取日历详情
type GetCalendarReq struct {
	CalIdList []string `json:"cal_id_list"`
}
type GetCalendarRep struct {
	util.WxError
	CalendarList []struct {
		CalId       string `json:"cal_id"`
		Organizer   string `json:"organizer"`
		Readonly    int64  `json:"readonly"`
		Summary     string `json:"summary"`
		Color       string `json:"color"`
		Description string `json:"description"`
		Shares      []struct {
			Userid   string `json:"userid"`
			Readonly int64  `json:"readonly,omitempty"`
		} `json:"shares"`
	} `json:"calendar_list"`
}
func (l *Calendar) GetCalendar(accessToken string, req GetCalendarReq)(result *GetCalendarRep, err error){
	qyUrl := fmt.Sprintf(QyCalendarGetURL, accessToken)

	response, err := util.PostJSON(qyUrl, req, l.ProxyUrl)

	if err != nil {
		return
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("GetCalendar error : errcode=%d , errmsg=%s", result.ErrCode, result.ErrMsg)
	}
	return
}

//删除日历
type DelCalendarReq struct {
	CalId string `json:"cal_id"`
}
type DelCalendarRep struct {
	util.WxError
}
func (l *Calendar) DelCalendar(accessToken string, req DelCalendarReq)(result *DelCalendarRep, err error){
	qyUrl := fmt.Sprintf(QyCalendarDelURL, accessToken)

	response, err := util.PostJSON(qyUrl, req, l.ProxyUrl)

	if err != nil {
		return
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("DelCalendar error : errcode=%d , errmsg=%s", result.ErrCode, result.ErrMsg)
	}
	return
}


//创建日程
type CreateScheduleReq struct {
	Schedule struct {
		Organizer string `json:"organizer"`
		StartTime int64  `json:"start_time"`
		EndTime   int64  `json:"end_time"`
		Attendees []struct {
			Userid string `json:"userid"`
		} `json:"attendees"`
		Summary     string `json:"summary"`
		Description string `json:"description"`
		Reminders   struct {
			IsRemind              int64   `json:"is_remind"`
			RemindBeforeEventSecs int64   `json:"remind_before_event_secs"`
			IsRepeat              int64   `json:"is_repeat"`
			RepeatType            int64   `json:"repeat_type"`
			RepeatUntil           int64   `json:"repeat_until"`
			IsCustomRepeat        int64   `json:"is_custom_repeat"`
			RepeatInterval        int64   `json:"repeat_interval"`
			RepeatDayOfWeek       []int64 `json:"repeat_day_of_week"`
			RepeatDayOfMonth      []int64 `json:"repeat_day_of_month"`
			Timezone              int64   `json:"timezone"`
		} `json:"reminders"`
		Location string `json:"location"`
		CalId    string `json:"cal_id"`
	} `json:"schedule"`
	Agentid int64 `json:"agentid"`
}
type CreateScheduleRep struct {
	util.WxError
	ScheduleId string `json:"schedule_id"`
}
func (l *Calendar) CreateSchedule(accessToken string, req CreateScheduleReq)(result *CreateScheduleRep, err error){
	qyUrl := fmt.Sprintf(QyScheduleCreateURL, accessToken)

	response, err := util.PostJSON(qyUrl, req, l.ProxyUrl)

	if err != nil {
		return
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("CreateSchedule error : errcode=%d , errmsg=%s", result.ErrCode, result.ErrMsg)
	}
	return
}

//更新日程
type UpdateScheduleReq struct {
	Schedule struct {
		Organizer  string `json:"organizer"`
		ScheduleId string `json:"schedule_id"`
		StartTime  int64  `json:"start_time"`
		EndTime    int64  `json:"end_time"`
		Attendees  []struct {
			Userid string `json:"userid"`
		} `json:"attendees"`
		Summary     string `json:"summary"`
		Description string `json:"description"`
		Reminders   struct {
			IsRemind              int64   `json:"is_remind"`
			RemindBeforeEventSecs int64   `json:"remind_before_event_secs"`
			IsRepeat              int64   `json:"is_repeat"`
			RepeatType            int64   `json:"repeat_type"`
			RepeatUntil           int64   `json:"repeat_until"`
			IsCustomRepeat        int64   `json:"is_custom_repeat"`
			RepeatInterval        int64   `json:"repeat_interval"`
			RepeatDayOfWeek       []int64 `json:"repeat_day_of_week"`
			RepeatDayOfMonth      []int64 `json:"repeat_day_of_month"`
			Timezone              int64   `json:"timezone"`
		} `json:"reminders"`
		Location      string `json:"location"`
		SkipAttendees bool   `json:"skip_attendees "`
	} `json:"schedule"`
}
type UpdateScheduleRep struct {
	util.WxError
}
func (l *Calendar) UpdateSchedule(accessToken string, req UpdateScheduleReq)(result *UpdateScheduleRep, err error){
	qyUrl := fmt.Sprintf(QyScheduleUpdateURL, accessToken)

	response, err := util.PostJSON(qyUrl, req, l.ProxyUrl)

	if err != nil {
		return
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("UpdateSchedule error : errcode=%d , errmsg=%s", result.ErrCode, result.ErrMsg)
	}
	return
}

//获取日程详情
type GetScheduleReq struct {
	ScheduleIdList []string `json:"schedule_id_list"`
}
type GetScheduleRep struct {
	util.WxError
	ScheduleList []struct {
		ScheduleId string `json:"schedule_id"`
		Organizer  string `json:"organizer"`
		Attendees  []struct {
			Userid         string `json:"userid"`
			ResponseStatus int64  `json:"response_status"`
		} `json:"attendees"`
		Summary     string `json:"summary"`
		Description string `json:"description"`
		Reminders   struct {
			IsRemind              int64   `json:"is_remind"`
			IsRepeat              int64   `json:"is_repeat"`
			RemindBeforeEventSecs int64   `json:"remind_before_event_secs"`
			RemindTimeDiffs       []int64 `json:"remind_time_diffs"`
			RepeatType            int64   `json:"repeat_type"`
			RepeatUntil           int64   `json:"repeat_until"`
			IsCustomRepeat        int64   `json:"is_custom_repeat"`
			RepeatInterval        int64   `json:"repeat_interval"`
			RepeatDayOfWeek       []int64 `json:"repeat_day_of_week"`
			RepeatDayOfMonth      []int64 `json:"repeat_day_of_month"`
			Timezone              int64   `json:"timezone"`
			ExcludeTimeList       []struct {
				StartTime int64 `json:"start_time"`
			} `json:"exclude_time_list"`
		} `json:"reminders"`
		Location  string `json:"location"`
		CalId     string `json:"cal_id"`
		StartTime int64  `json:"start_time"`
		EndTime   int64  `json:"end_time"`
		Status    int64  `json:"status"`
	} `json:"schedule_list"`
}
func (l *Calendar) GetSchedule(accessToken string, req GetScheduleReq)(result *GetScheduleRep, err error){
	qyUrl := fmt.Sprintf(QyScheduleGetURL, accessToken)

	response, err := util.PostJSON(qyUrl, req, l.ProxyUrl)

	if err != nil {
		return
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("GetSchedule error : errcode=%d , errmsg=%s", result.ErrCode, result.ErrMsg)
	}
	return
}

//取消日程
type DelScheduleReq struct {
	ScheduleId string `json:"schedule_id"`
}
type DelScheduleRep struct {
	util.WxError
}
func (l *Calendar) DelSchedule(accessToken string, req DelScheduleReq)(result *DelScheduleRep, err error){
	qyUrl := fmt.Sprintf(QyScheduleDelURL, accessToken)

	response, err := util.PostJSON(qyUrl, req, l.ProxyUrl)

	if err != nil {
		return
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("DelSchedule error : errcode=%d , errmsg=%s", result.ErrCode, result.ErrMsg)
	}
	return
}

//获取日历下的日程列表
type ListScheduleReq struct {
	CalId  string `json:"cal_id"`
	Offset int64  `json:"offset"`
	Limit  int64  `json:"limit"`
}
type ListScheduleRep struct {
	util.WxError
	ScheduleList []struct {
		ScheduleId string `json:"schedule_id"`
		Sequence   int64    `json:"sequence"`
		Attendees  []struct {
			Userid         string `json:"userid"`
			ResponseStatus int64    `json:"response_status"`
		} `json:"attendees"`
		Summary     string `json:"summary"`
		Description string `json:"description"`
		Reminders   struct {
			IsRemind              int64   `json:"is_remind"`
			IsRepeat              int64   `json:"is_repeat"`
			RemindBeforeEventSecs int64   `json:"remind_before_event_secs"`
			RepeatType            int64   `json:"repeat_type"`
			RepeatUntil           int64   `json:"repeat_until"`
			IsCustomRepeat        int64   `json:"is_custom_repeat"`
			RepeatInterval        int64   `json:"repeat_interval"`
			RepeatDayOfWeek       []int64 `json:"repeat_day_of_week"`
			RepeatDayOfMonth      []int64 `json:"repeat_day_of_month"`
			Timezone              int64   `json:"timezone"`
		} `json:"reminders"`
		Location  string `json:"location"`
		StartTime int64  `json:"start_time"`
		EndTime   int64  `json:"end_time"`
		Status    int64  `json:"status"`
		CalId     string `json:"cal_id"`
	} `json:"schedule_list"`
}
func (l *Calendar) ListSchedule(accessToken string, req ListScheduleReq)(result *ListScheduleRep, err error){
	qyUrl := fmt.Sprintf(QyScheduleListURL, accessToken)

	response, err := util.PostJSON(qyUrl, req, l.ProxyUrl)

	if err != nil {
		return
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("ListSchedule error : errcode=%d , errmsg=%s", result.ErrCode, result.ErrMsg)
	}
	return
}

//新增日程参与者
type AddAttendeesScheduleReq struct {
	ScheduleId string `json:"schedule_id"`
	Attendees  []struct {
		Userid string `json:"userid"`
	} `json:"attendees"`
}
type AddAttendeesScheduleRep struct {
	util.WxError
}
func (l *Calendar) AddAttendeesSchedule(accessToken string, req AddAttendeesScheduleReq)(result *AddAttendeesScheduleRep, err error){
	qyUrl := fmt.Sprintf(QyScheduleAddAttendeesURL, accessToken)

	response, err := util.PostJSON(qyUrl, req, l.ProxyUrl)

	if err != nil {
		return
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("AddAttendeesSchedule error : errcode=%d , errmsg=%s", result.ErrCode, result.ErrMsg)
	}
	return
}

//删除日程参与者
type DelAttendeesScheduleReq struct {
	ScheduleId string `json:"schedule_id"`
	Attendees  []struct {
		Userid string `json:"userid"`
	} `json:"attendees"`
}
type DelAttendeesScheduleRep struct {
	util.WxError
}
func (l *Calendar) DelAttendeesSchedule(accessToken string, req DelAttendeesScheduleReq)(result *DelAttendeesScheduleRep, err error){
	qyUrl := fmt.Sprintf(QyScheduleDelAttendeesURL, accessToken)

	response, err := util.PostJSON(qyUrl, req, l.ProxyUrl)

	if err != nil {
		return
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("DelAttendeesSchedule error : errcode=%d , errmsg=%s", result.ErrCode, result.ErrMsg)
	}
	return
}
