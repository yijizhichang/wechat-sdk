package menu

const (
	clickType            = "click"
	viewType             = "view"
	scanCodePushType     = "scancode_push"
	scanCodeWaitPushType = "scancode_waitmsg"
	picSYSPhotoType      = "pic_sysphoto"
	picPhotoOrAlbumType  = "pic_photo_or_album"
	picWeixinType        = "pic_weixin"
	locationSelectType   = "location_select"
	mediaIdType          = "media_id"
	viewLimitedType      = "view_limited"
	miniprogramType      = "miniprogram"
)

//Button 菜单按钮
type Button struct {
	Type       string    `json:"type,omitempty"`
	Name       string    `json:"name,omitempty"`
	Key        string    `json:"key,omitempty"`
	URL        string    `json:"url,omitempty"`
	MediaID    string    `json:"media_id,omitempty"`
	Appid      string    `json:"appid,omitempty"`    // 小程序的appid（仅认证公众号可配置）
	Pagepath   string    `json:"pagepath,omitempty"` // 小程序的页面路径
	SubButtons []*Button `json:"sub_button,omitempty"`
}

func SetButton(options ...ButtonOption) []*Button {
	var btn []*Button
	for _, option := range options {
		bt := new(Button)
		option(bt)
		btn = append(btn, bt)
	}
	return btn
}

// 选项模式
type ButtonOption func(btn *Button)

func WithSubButton(name string, subButton []*Button) ButtonOption {
	return func(btn *Button) {
		btn.Name = name
		btn.SubButtons = subButton
	}
}

// 为click类型
func WithClickButton(name, key string) ButtonOption {
	return func(btn *Button) {
		btn.Type = clickType
		btn.Name = name
		btn.Key = key
	}
}

// view类型
func WithViewButton(name, url string) ButtonOption {
	return func(btn *Button) {
		btn.Type = viewType
		btn.Name = name
		btn.URL = url
	}
}

// 扫码推事件
func WithScanCodePushButton(name, key string) ButtonOption {
	return func(btn *Button) {
		btn.Type = scanCodePushType
		btn.Name = name
		btn.Key = key
	}
}

// 扫码推事件且弹出"消息接收中"提示框
func WithScanCodeWaitMsgButton(name, key string) ButtonOption {
	return func(btn *Button) {
		btn.Type = scanCodeWaitPushType
		btn.Name = name
		btn.Key = key
	}
}

// 设置弹出系统拍照发图按钮
func WithPicSysPhotoButton(name, key string) ButtonOption {
	return func(btn *Button) {
		btn.Type = picSYSPhotoType
		btn.Name = name
		btn.Key = key
	}
}

// 设置弹出拍照或者相册发图类型按钮
func WithPicPhotoOrAlbumButton(name, key string) ButtonOption {
	return func(btn *Button) {
		btn.Type = picPhotoOrAlbumType
		btn.Name = name
		btn.Key = key
	}
}

// 设置弹出微信相册发图器类型按钮
func WithPicWeixinButton(name, key string) ButtonOption {
	return func(btn *Button) {
		btn.Type = picWeixinType
		btn.Name = name
		btn.Key = key
	}
}

// 设置 弹出地理位置选择器 类型按钮
func WithLocationSelectButton(name, key string) ButtonOption {
	return func(btn *Button) {
		btn.Type = locationSelectType
		btn.Name = name
		btn.Key = key
	}
}

// 设置下发消息(除文本消息) 类型按钮
func WithMediaIDButton(name, mediaID string) ButtonOption {
	return func(btn *Button) {
		btn.Type = mediaIdType
		btn.Name = name
		btn.MediaID = mediaID
	}
}

// 设置 跳转图文消息URL 类型按钮
func WithViewLimitedButton(name, mediaID string) ButtonOption {
	return func(btn *Button) {
		btn.Type = viewLimitedType
		btn.Name = name
		btn.MediaID = mediaID
	}
}

// 设置 跳转小程序 类型按钮
func WithMiniprogramButton(name, url, appid, pagepath string) ButtonOption {
	return func(btn *Button) {
		btn.Type = miniprogramType
		btn.Name = name
		btn.URL = url
		btn.Appid = appid
		btn.Pagepath = pagepath
	}
}
