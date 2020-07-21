package models

type HOrder struct {
	Id int `json:"id" gorm:"primary_key"`
	Num string `json:"num"`
	IdCard string `json:"idCard" gorm:"column:idCard"`
	OrderNo string `json:"orderNo" gorm:"column:orderNo"`
	RcptNo string `json:"rcptNo" gorm:"column:rcptNo"`
	TransNo string `json:"transNo" gorm:"column:transNo"`
	CreateTime int64 `json:"create_time" gorm:"column:create_time"`
	AppointTime int64 `json:"appoint_time" gorm:"column:appoint_time"`
	RefundTime int64 `json:"refund_time" gorm:"column:refund_time"`
	ApplyTime int64 `json:"apply_time" gorm:"column:apply_time"`
	CompleteTime int64 `json:"complete_time" gorm:"column:complete_time"`
	Type string `json:"type"  gorm:"column:type"`
	Openid string `json:"openid"`
	Status string `json:"status"`
	Bill string `json:"bill"`
	IfCheck string `json:"if_check" gorm:"column:if_check"`
}

func (HOrder) TableName() string {
	return "h_order"
}

