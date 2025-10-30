package model

import (
	"fmt"
	"musicRoomBookingbot/util"
	"time"
)

// 네이버나 카카오 OAuth 로그인
type Host struct {
	Id          int64  `form:"id" json:"id" db:"Id"`
	HostName    string `form:"hostName" json:"hostName" db:"HostName"`
	PhoneNo     string `form:"phoneNo" json:"phoneNo" db:"PhoneNo"`
	KakaoTalkId string `form:"kakaoTalkId" json:"kakaoTalkId" db:"KakaoTalkId"`
	CreatedAt   string `form:"createdAt" json:"createdAt" db:"CreatedAt"`
	Discard     int    `form:"discard" json:"discard" db:"Discard"`
}

type RoomFilter struct {
	UserIdList       []int64 `form:"userId" json:"userIdList"`             //특정 유저 아이디들
	RoomIdList       []int64 `form:"roomId" json:"roomIdList"`             //특정 룸 아이디들
	HostIdList       []int64 `form:"hostId" json:"hostIdList"`             //특정 호스트 아이디들
	GroupIdList      []int64 `form:"groupId" json:"groupIdList"`           //특정 룸 그룹 아이디들
	DayOfWeekStart   int     `form:"dayOfWeekStart" json:"dayOfWeekStart"` //예약 가능 요일 시작
	DayOfWeekEnd     int     `form:"dayOfWeekEnd" json:"dayOfWeekEnd"`     //예약 가능 요일 끝
	StartTime        string  `form:"startTime" json:"startTime"`           //예약 가능 시작 시간
	EndTime          string  `form:"endTime" json:"endTime"`               //예약 가능 끝 시간
	Keyword          string  `form:"keyword" json:"keyword"`
	OccupationStatus int     `form:"occupationStatus" json:"occupationStatus"` //0:빈방, 1:예약불가
}

type RoomGroup struct {
	Id        int64  `form:"id" json:"id" db:"Id"`
	GroupName string `form:"groupName" json:"groupName" db:"GroupName"`
	HostId    int64  `form:"hostId" json:"hostId" db:"HostId"`
	Address   string `form:"address" json:"address" db:"Address"`
	CreatedAt string `form:"createdAt" json:"createdAt" db:"CreatedAt"`
	Discard   int    `form:"discard" json:"discard" db:"Discard"`
}

type RoomDetail struct {
	Room               Room                `json:"room"`
	TimeSlotExceptions []TimeSlotException `json:"timeSlotExceptions"`
	TimeSlots          []TimeSlot          `json:"timeSlots"`
}

type Room struct {
	Id                      int64  `form:"id" json:"id" db:"Id"`
	GroupId                 int64  `form:"groupId" json:"groupId" db:"GroupId"`
	RoomName                string `form:"roomName" json:"roomName" db:"RoomName"`
	Discard                 int    `form:"discard" json:"discard" db:"Discard"`
	ReservableDaysMinOffset int    `form:"reservableDaysMinOffset" json:"reservableDaysMinOffset" db:"ReservableDaysMinOffset"`
	ReservableDaysMaxOffset int    `form:"reservableDaysMaxOffset" json:"reservableDaysMaxOffset" db:"ReservableDaysMaxOffset"`
}

type Reservation struct {
	Id              int64 `form:"id" json:"id" db:"Id"`
	UserId          int64 `form:"userId" json:"userId" db:"UserId"`
	RoomId          int64 `form:"roomId" json:"roomId" db:"RoomId"`
	StartTime       int64 `form:"startTime" json:"startTime" db:"StartTime"`
	EndTime         int64 `form:"endTime" json:"endTime" db:"EndTime"`
	CheckinTime     int64 `form:"checkinTime" json:"checkinTime" db:"CheckinTime"`
	CheckoutTime    int64 `form:"checkoutTime" json:"checkoutTime" db:"CheckoutTime"`
	ExtendedMinutes int   `form:"extendedMinutes" json:"extendedMinutes" db:"ExtendedMinutes"`
	Status          int   `form:"status" json:"status" db:"Status"`
	CancelReason    int   `form:"cancelReason" json:"cancelReason" db:"CancelReason"`
	Discard         int   `form:"discard" json:"discard" db:"Discard"`
}

type TimeSlotException struct {
	Id            int64  `form:"id" json:"id" db:"Id"`
	RoomId        int64  `form:"roomId" json:"roomId" db:"RoomId"`
	Date          string `form:"date" json:"date" db:"Date"` //YYYYMMDD
	DayOfWeek     int    `form:"dayOfWeek" json:"dayOfWeek" db:"DayOfWeek"`
	StartTime     string `form:"startTime" json:"startTime" db:"StartTime"`
	StartYYYYMMDD time.Time
	EndTime       string `form:"endTime" json:"endTime" db:"EndTime"`
	EndYYYYMMDD   time.Time
	Reason        int    `form:"reason" json:"reason" db:"Reason"`
	ReasonText    string `form:"reasonText" json:"reasonText" db:"ReasonText"`
	Discard       int    `form:"discard" json:"discard" db:"Discard"`
}

func (e *TimeSlotException) ParseStartAndEndTime() error {

	startStr := fmt.Sprintf("%s%s", e.Date, e.StartTime)
	start, err := time.ParseInLocation(util.YYYYMMDDhhmm, startStr, util.KST)
	if err != nil {
		return err
	}
	e.StartYYYYMMDD = start

	endStr := fmt.Sprintf("%s%s", e.Date, e.EndTime)
	end, err := time.ParseInLocation(util.YYYYMMDDhhmm, endStr, util.KST)
	if err != nil {
		return err
	}
	e.EndYYYYMMDD = end

	return nil
}

// 날짜 - 시작시각 HHMM 끝시각 HHMM
type TimeSlotExceptionDayMap map[string][]TimeSlotExceptionHHMM

type TimeSlotExceptionHHMM struct {
	Start time.Time
	End   time.Time
}

type TimeSlotsDetail struct {
	Id                      int64              `form:"id" json:"id" db:"Id"`
	StartTime               string             `form:"startTime" json:"startTime" db:"StartTime"`
	EndTime                 string             `form:"endTime" json:"endTime" db:"EndTime"`
	SplittedSlots           []SplittedTimeSlot `json:"splittedSlots"`
	DayOfWeek               int                `form:"dayOfWeek" json:"dayOfWeek" db:"DayOfWeek"`
	RoomId                  int64              `form:"roomId" json:"roomId" db:"RoomId"`
	RoomName                string             `form:"roomName" json:"roomName" db:"RoomName"`
	ReservationUnitMinutes  int                `form:"reservationUnitMinutes" json:"reservationUnitMinutes" db:"ReservationUnitMinutes"`
	ReservableDaysMinOffset int                `form:"reservableDaysMinOffset" json:"reservableDaysMinOffset" db:"ReservableDaysMinOffset"`
	ReservableDaysMaxOffset int                `form:"reservableDaysMaxOffset" json:"reservableDaysMaxOffset" db:"ReservableDaysMaxOffset"`
	GroupId                 int64              `form:"groupId" json:"groupId" db:"GroupId"`
	GroupName               string             `form:"groupName" json:"groupName" db:"GroupName"`
	Address                 string             `form:"address" json:"address" db:"Address"`
	Discard                 int                `form:"discard" json:"discard" db:"Discard"`
}

type TimeSlot struct {
	Id                     int64  `form:"id" json:"id" db:"Id"`
	RoomId                 int64  `form:"roomId" json:"roomId" db:"RoomId"`
	DayOfWeek              int    `form:"dayOfWeek" json:"dayOfWeek" db:"DayOfWeek"`
	StartTime              string `form:"startTime" json:"startTime" db:"StartTime"`
	StartTimeParsed        time.Time
	EndTime                string `form:"endTime" json:"endTime" db:"EndTime"`
	EndTimeParsed          time.Time
	ReservationUnitMinutes int `form:"reservationUnitMinutes" json:"reservationUnitMinutes" db:"ReservationUnitMinutes"`
	Discard                int `form:"discard" json:"discard" db:"Discard"`
}

func (ts *TimeSlot) ParseStartAndEndTime(yyyymmdd string) error {
	startStr := fmt.Sprintf("%s%s", yyyymmdd, ts.StartTime)
	start, err := time.ParseInLocation(util.YYYYMMDDhhmm, startStr, util.KST)
	if err != nil {
		return err
	}

	ts.StartTimeParsed = start

	endStr := fmt.Sprintf("%s%s", yyyymmdd, ts.EndTime)
	end, err := time.ParseInLocation(util.YYYYMMDDhhmm, endStr, util.KST)
	if err != nil {
		return err
	}
	ts.EndTimeParsed = end

	return nil
}

var Monday = 0
var Tuesday = 1
var Wednesday = 2
var Thursday = 3
var Friday = 4
var Saturday = 5
var Sunday = 6

type SplittedTimeSlot struct {
	StartTimeParsed time.Time
	StartTime       string `form:"startTime" json:"startTime"` //YYYYMMDDhhmm
	EndTimeParsed   time.Time
	EndTime         string `form:"endTime" json:"endTime"` //YYYYMMDDhhmm
}

type TimeSlotFilter struct {
	StartDateTime       string `form:"startDateTime" json:"startDateTime"`
	StartDateTimeParsed time.Time
	EndDateTime         string `form:"endDateTime" json:"endDateTime"`
	EndDateTimeParsed   time.Time
	Keyword             string  `form:"keyword" json:"keyword"`     //호스트, 지점, 연습실 이름
	HostIdList          []int64 `form:"hostId" json:"hostIdList"`   //특정 호스트 아이디들
	GroupIdList         []int64 `form:"groupId" json:"groupIdList"` //특정 룸 그룹 아이디들
	RoomIdList          []int64 `form:"roomId" json:"roomIdList"`   //특정 룸 아이디들

}

func (tf *TimeSlotFilter) ParseTime() error {
	currentTime := util.GetCurrentTime()

	tf.StartDateTimeParsed = currentTime

	if tf.StartDateTime == "" {
		tf.StartDateTime = currentTime.Format(util.YYYYMMDDhhmmss)
	}

	endDateTimeParsed, err := time.Parse(util.YYYYMMDDhhmmss, tf.EndDateTime)
	if err != nil {
		return err
	}
	tf.EndDateTimeParsed = endDateTimeParsed

	return nil
}

type User struct {
	UserId       int64  `form:"id" json:"id" db:"Id"`
	UserName     string `form:"userName" json:"userName" db:"UserName"`
	PhoneNo      string `form:"phoneNo" json:"phoneNo" db:"PhoneNo"`
	KakaoTalkId  string `form:"kakaoTalkId" json:"kakaoTalkId" db:"KakaoTalkId"`
	RegisterDate string `form:"registerDate" json:"registerDate" db:"RegisterDate"`
	Discard      int    `form:"discard" json:"discard" db:"Discard"`
}
