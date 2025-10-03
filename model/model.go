package model

//네이버나 카카오 OAuth 로그인
type Host struct {
	Id          *int64  `form:"id" json:"id" db:"Id"`
	HostName    *string `form:"hostName" json:"hostName" db:"HostName"`
	PhoneNo     *string `form:"phoneNo" json:"phoneNo" db:"PhoneNo"`
	KakaoTalkId *string `form:"kakaoTalkId" json:"kakaoTalkId" db:"KakaoTalkId"`
	CreatedAt   *string `form:"createdAt" json:"createdAt" db:"CreatedAt"`
	Discard     *int    `form:"discard" json:"discard" db:"Discard"`
}

type RoomGroup struct {
	Id        *int64  `form:"id" json:"id" db:"Id"`
	GroupName *string `form:"groupName" json:"groupName" db:"GroupName"`
	HostId    *int64  `form:"hostId" json:"hostId" db:"HostId"`
	Address   *string `form:"address" json:"address" db:"Address"`
	CreatedAt *string `form:"createdAt" json:"createdAt" db:"CreatedAt"`
	Discard   *int    `form:"discard" json:"discard" db:"Discard"`
}

type RoomDetail struct {
	Room                     Room                     `json:"room"`
	BookingExceptionPolicies []BookingExceptionPolicy `json:"bookingExceptionPolicies"`
	BookingPolicies          []BookingPolicy          `json:"bookingPolicies"`
}

type Room struct {
	RoomId   *int64  `form:"id" json:"id" db:"Id"`
	GroupId  *int64  `form:"groupId" json:"groupId" db:"GroupId"`
	RoomName *string `form:"roomName" json:"roomName" db:"RoomName"`
	Discard  *int    `form:"discard" json:"discard" db:"Discard"`
}

type Reservation struct {
	Id              *int64 `form:"id" json:"id" db:"Id"`
	UserId          *int64 `form:"userId" json:"userId" db:"UserId"`
	RoomId          *int64 `form:"roomId" json:"roomId" db:"RoomId"`
	StartTime       *int64 `form:"startTime" json:"startTime" db:"StartTime"`
	EndTime         *int64 `form:"endTime" json:"endTime" db:"EndTime"`
	CheckinTime     *int64 `form:"checkinTime" json:"checkinTime" db:"CheckinTime"`
	CheckoutTime    *int64 `form:"checkoutTime" json:"checkoutTime" db:"CheckoutTime"`
	ExtendedMinutes *int   `form:"extendedMinutes" json:"extendedMinutes" db:"ExtendedMinutes"`
	Status          *int   `form:"status" json:"status" db:"Status"`
	CancelReason    *int   `form:"cancelReason" json:"cancelReason" db:"CancelReason"`
}

type BookingExceptionPolicy struct {
	Id         *int64  `form:"id" json:"id" db:"Id"`
	RoomId     *int64  `form:"roomId" json:"roomId" db:"RoomId"`
	Date       *string `form:"date" json:"date" db:"Date"`
	StartTime  *string `form:"startTime" json:"startTime" db:"StartTime"`
	EndTime    *string `form:"endTime" json:"endTime" db:"EndTime"`
	Reason     *int    `form:"reason" json:"reason" db:"Reason"`
	ReasonText *string `form:"reasonText" json:"reasonText" db:"ReasonText"`
	Discard    *int    `form:"discard" json:"discard" db:"Discard"`
}

type BookingPolicy struct {
	Id        *int64  `form:"id" json:"id" db:"Id"`
	RoomId    *int64  `form:"roomId" json:"roomId" db:"RoomId"`
	DayOfWeek *int    `form:"dayOfWeek" json:"dayOfWeek" db:"DayOfWeek"`
	StartTime *string `form:"startTime" json:"startTime" db:"StartTime"`
	EndTime   *string `form:"endTime" json:"endTime" db:"EndTime"`
	Discard   *int    `form:"discard" json:"discard" db:"Discard"`
}

var Monday = 0
var Tuesday = 1
var Wednesday = 2
var Thursday = 3
var Friday = 4
var Saturday = 5
var Sunday = 6

type User struct {
	UserId       *int64  `form:"id" json:"id" db:"Id"`
	UserName     *string `form:"userName" json:"userName" db:"UserName"`
	PhoneNo      *string `form:"phoneNo" json:"phoneNo" db:"PhoneNo"`
	KakaoTalkId  *string `form:"kakaoTalkId" json:"kakaoTalkId" db:"KakaoTalkId"`
	RegisterDate *string `form:"registerDate" json:"registerDate" db:"RegisterDate"`
	Discard      *int    `form:"discard" json:"discard" db:"Discard"`
}
