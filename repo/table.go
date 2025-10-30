package repo

import "fmt"

const MusicRoom = "MusicRoom"

var Host = fmt.Sprintf("%s.%s", MusicRoom, "Host")
var RoomGroup = fmt.Sprintf("%s.%s", MusicRoom, "RoomGroup")
var Room = fmt.Sprintf("%s.%s", MusicRoom, "Room")
var Reservation = fmt.Sprintf("%s.%s", MusicRoom, "Reservation")
var DaySlotException = fmt.Sprintf("%s.%s", MusicRoom, "DaySlotException")
var DaySlot = fmt.Sprintf("%s.%s", MusicRoom, "DaySlot")
var User = fmt.Sprintf("%s.%s", MusicRoom, "User")
