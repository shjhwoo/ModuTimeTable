package model

import (
	"musicRoomBookingbot/util"
	"time"
)

// 네이버나 카카오 OAuth 로그인
type Host struct {
	Id          *int64  `form:"id" json:"id" db:"Id"`
	HostName    *string `form:"hostName" json:"hostName" db:"HostName"`
	PhoneNo     *string `form:"phoneNo" json:"phoneNo" db:"PhoneNo"`
	KakaoTalkId *string `form:"kakaoTalkId" json:"kakaoTalkId" db:"KakaoTalkId"`
	CreatedAt   *string `form:"createdAt" json:"createdAt" db:"CreatedAt"`
	Discard     *int    `form:"discard" json:"discard" db:"Discard"`
}

type RoomFilter struct {
	UserIdList       []*int64 `form:"userId" json:"userIdList"`                 //특정 유저 아이디들
	RoomIdList       []*int64 `form:"roomId" json:"roomIdList"`                 //특정 룸 아이디들
	HostIdList       []*int64 `form:"hostId" json:"hostIdList"`                 //특정 호스트 아이디들
	GroupIdList      []*int64 `form:"groupId" json:"groupIdList"`               //특정 룸 그룹 아이디들
	DayOfWeekStart   *int     `form:"dayOfWeekStart" json:"dayOfWeekStart"`     //예약 가능 요일 시작
	DayOfWeekEnd     *int     `form:"dayOfWeekEnd" json:"dayOfWeekEnd"`         //예약 가능 요일 끝
	StartTime        *string  `form:"startTime" json:"startTime"`               //예약 가능 시작 시간
	EndTime          *string  `form:"endTime" json:"endTime"`                   //예약 가능 끝 시간
	HostNameLike     *string  `form:"hostNameLike" json:"hostNameLike"`         //호스트 이름. 예: xx음악학원, oo연습실 등
	GroupNameLike    *string  `form:"groupNameLike" json:"groupNameLike"`       //지점 이름
	RoomNameLike     *string  `form:"roomNameLike" json:"roomNameLike"`         //연습실 이름
	AddressLike      *string  `form:"addressLike" json:"addressLike"`           //주소
	OccupationStatus *int     `form:"occupationStatus" json:"occupationStatus"` //0:빈방, 1:예약불가
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
	Room               Room                `json:"room"`
	TimeSlotExceptions []TimeSlotException `json:"timeSlotExceptions"`
	TimeSlots          []TimeSlot          `json:"timeSlots"`
}

type Room struct {
	Id                      *int64  `form:"id" json:"id" db:"Id"`
	GroupId                 *int64  `form:"groupId" json:"groupId" db:"GroupId"`
	RoomName                *string `form:"roomName" json:"roomName" db:"RoomName"`
	Discard                 *int    `form:"discard" json:"discard" db:"Discard"`
	ReservableDaysMinOffset *int    `form:"reservableDaysMinOffset" json:"reservableDaysMinOffset" db:"reservableDaysMinOffset"`
	ReservableDaysMaxOffset *int    `form:"reservableDaysMaxOffset" json:"reservableDaysMaxOffset" db:"reservableDaysMaxOffset"`
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

type TimeSlotException struct {
	Id         *int64  `form:"id" json:"id" db:"Id"`
	RoomId     *int64  `form:"roomId" json:"roomId" db:"RoomId"`
	Date       *string `form:"date" json:"date" db:"Date"`
	DayOfWeek  *int    `form:"dayOfWeek" json:"dayOfWeek" db:"DayOfWeek"`
	StartTime  *string `form:"startTime" json:"startTime" db:"StartTime"`
	EndTime    *string `form:"endTime" json:"endTime" db:"EndTime"`
	Reason     *int    `form:"reason" json:"reason" db:"Reason"`
	ReasonText *string `form:"reasonText" json:"reasonText" db:"ReasonText"`
	Discard    *int    `form:"discard" json:"discard" db:"Discard"`
}

type TimeSlotsDetail struct {
	DayOfWeek *int    `form:"dayOfWeek" json:"dayOfWeek" db:"DayOfWeek"`
	StartTime *string `form:"startTime" json:"startTime" db:"StartTime"`
	EndTime   *string `form:"endTime" json:"endTime" db:"EndTime"`
	RoomId    *int64  `form:"roomId" json:"roomId" db:"RoomId"`
	RoomName  *string `form:"roomName" json:"roomName" db:"RoomName"`
	GroupId   *int64  `form:"groupId" json:"groupId" db:"GroupId"`
	GroupName *string `form:"groupName" json:"groupName" db:"GroupName"`
	Address   *string `form:"address" json:"address" db:"Address"`
}

type TimeSlot struct {
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

type TimeSlotFilter struct {
	StartDateTime       *string `form:"startDateTime" json:"startDateTime"`
	StartDateTimeParsed time.Time
	EndDateTime         *string `form:"endDateTime" json:"endDateTime"`
	EndDateTimeParsed   time.Time
	HostIdList          []*int64 `form:"hostId" json:"hostIdList"`   //특정 호스트 아이디들
	GroupIdList         []*int64 `form:"groupId" json:"groupIdList"` //특정 룸 그룹 아이디들
	RoomIdList          []*int64 `form:"roomId" json:"roomIdList"`   //특정 룸 아이디들
}

func (tf *TimeSlotFilter) ParseTime() error {
	currentTime := util.GetCurrentTime()

	tf.StartDateTimeParsed = currentTime

	if util.SafeStr(tf.StartDateTime) == "" {
		tf.StartDateTime = util.StringPtr(currentTime.Format(util.YYYYMMDDhhmmss))
	}

	endDateTimeParsed, err := time.Parse(util.YYYYMMDDhhmmss, util.SafeStr(tf.EndDateTime))
	if err != nil {
		return err
	}
	tf.EndDateTimeParsed = endDateTimeParsed

	return nil
}

type User struct {
	UserId       *int64  `form:"id" json:"id" db:"Id"`
	UserName     *string `form:"userName" json:"userName" db:"UserName"`
	PhoneNo      *string `form:"phoneNo" json:"phoneNo" db:"PhoneNo"`
	KakaoTalkId  *string `form:"kakaoTalkId" json:"kakaoTalkId" db:"KakaoTalkId"`
	RegisterDate *string `form:"registerDate" json:"registerDate" db:"RegisterDate"`
	Discard      *int    `form:"discard" json:"discard" db:"Discard"`
}
