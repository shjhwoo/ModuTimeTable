-- --------------------------------------------------------
-- 호스트:                          127.0.0.1
-- 서버 버전:                        8.0.43 - MySQL Community Server - GPL
-- 서버 OS:                        Linux
-- HeidiSQL 버전:                  11.3.0.6295
-- --------------------------------------------------------

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET NAMES utf8 */;
/*!50503 SET NAMES utf8mb4 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;


-- MusicRoom 데이터베이스 구조 내보내기
CREATE DATABASE IF NOT EXISTS `MusicRoom` /*!40100 DEFAULT CHARACTER SET utf8mb3 */ /*!80016 DEFAULT ENCRYPTION='N' */;
USE `MusicRoom`;

-- 테이블 MusicRoom.DaySlot 구조 내보내기
CREATE TABLE IF NOT EXISTS `DaySlot` (
  `Id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `RoomId` bigint unsigned NOT NULL DEFAULT '0',
  `DayOfWeek` tinyint unsigned NOT NULL DEFAULT '0',
  `StartTime` varchar(4) NOT NULL DEFAULT '',
  `EndTime` varchar(4) NOT NULL DEFAULT '',
  `ReservationUnitMinutes` tinyint unsigned NOT NULL DEFAULT '0',
  `Closed` tinyint(1) unsigned zerofill NOT NULL DEFAULT '0',
  `Discard` tinyint(1) unsigned zerofill NOT NULL DEFAULT '0',
  PRIMARY KEY (`Id`),
  KEY `RoomId` (`RoomId`),
  KEY `DayOfWeek` (`DayOfWeek`),
  KEY `StartTime` (`StartTime`),
  KEY `EndTime` (`EndTime`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8mb3 COMMENT='연습실 시간표 기본 정책 테이블';

-- 테이블 데이터 MusicRoom.DaySlot:~0 rows (대략적) 내보내기
DELETE FROM `DaySlot`;
/*!40000 ALTER TABLE `DaySlot` DISABLE KEYS */;
INSERT INTO `DaySlot` (`Id`, `RoomId`, `DayOfWeek`, `StartTime`, `EndTime`, `ReservationUnitMinutes`, `Closed`, `Discard`) VALUES
	(1, 1, 0, '0000', '0000', 0, 1, 0),
	(2, 1, 1, '0900', '1200', 60, 0, 0),
	(3, 1, 1, '1200', '2000', 30, 0, 0),
	(4, 1, 2, '0900', '2359', 60, 0, 0),
	(5, 1, 3, '0900', '2359', 60, 0, 0),
	(6, 1, 4, '0900', '2359', 60, 0, 0),
	(7, 1, 5, '0900', '2359', 60, 0, 0),
	(8, 1, 6, '0900', '2359', 60, 0, 0);
/*!40000 ALTER TABLE `DaySlot` ENABLE KEYS */;

-- 테이블 MusicRoom.DaySlotException 구조 내보내기
CREATE TABLE IF NOT EXISTS `DaySlotException` (
  `Id` bigint NOT NULL AUTO_INCREMENT,
  `RoomId` bigint NOT NULL DEFAULT '0',
  `StartDateTime` varchar(14) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '',
  `EndDateTime` varchar(14) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '',
  `Reason` tinyint unsigned NOT NULL DEFAULT '0',
  `ReasonText` tinytext CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci,
  `Discard` tinyint unsigned NOT NULL DEFAULT '0',
  PRIMARY KEY (`Id`),
  KEY `RoomId` (`RoomId`),
  KEY `StartTime` (`StartDateTime`) USING BTREE,
  KEY `EndTime` (`EndDateTime`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb3;

-- 테이블 데이터 MusicRoom.DaySlotException:~0 rows (대략적) 내보내기
DELETE FROM `DaySlotException`;
/*!40000 ALTER TABLE `DaySlotException` DISABLE KEYS */;
INSERT INTO `DaySlotException` (`Id`, `RoomId`, `StartDateTime`, `EndDateTime`, `Reason`, `ReasonText`, `Discard`) VALUES
	(1, 1, '20251105090000', '20251109160000', 0, '내부 인테리어 공사', 0);
/*!40000 ALTER TABLE `DaySlotException` ENABLE KEYS */;

-- 테이블 MusicRoom.Host 구조 내보내기
CREATE TABLE IF NOT EXISTS `Host` (
  `Id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `HostName` varchar(50) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '',
  `PhoneNo` varchar(11) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '',
  `KakaoTalkId` varchar(50) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '',
  `CreatedAt` varchar(14) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '',
  `Discard` tinyint unsigned NOT NULL DEFAULT '0',
  PRIMARY KEY (`Id`) USING BTREE,
  KEY `PhoneNo` (`PhoneNo`) USING BTREE,
  KEY `KakaoTalkId` (`KakaoTalkId`) USING BTREE,
  KEY `UserName` (`HostName`) USING BTREE,
  KEY `CreatedAt` (`CreatedAt`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb3 ROW_FORMAT=DYNAMIC COMMENT='호스트 테이블';

-- 테이블 데이터 MusicRoom.Host:~0 rows (대략적) 내보내기
DELETE FROM `Host`;
/*!40000 ALTER TABLE `Host` DISABLE KEYS */;
INSERT INTO `Host` (`Id`, `HostName`, `PhoneNo`, `KakaoTalkId`, `CreatedAt`, `Discard`) VALUES
	(1, '이든뮤직', '01098080785', 'shjhwoo', '20251029220815', 0);
/*!40000 ALTER TABLE `Host` ENABLE KEYS */;

-- 테이블 MusicRoom.Reservation 구조 내보내기
CREATE TABLE IF NOT EXISTS `Reservation` (
  `Id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `UserId` bigint unsigned NOT NULL DEFAULT '0',
  `RoomId` bigint unsigned NOT NULL DEFAULT '0',
  `StartDateTime` varchar(14) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '',
  `EndDateTime` varchar(14) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '',
  `CheckinTime` varchar(14) NOT NULL DEFAULT '',
  `CheckoutTime` varchar(14) NOT NULL DEFAULT '',
  `ExtendedMinutes` tinyint NOT NULL DEFAULT '0',
  `Status` tinyint unsigned NOT NULL DEFAULT '0',
  `CancelReason` tinyint unsigned NOT NULL DEFAULT '0',
  `Discard` tinyint unsigned NOT NULL DEFAULT '0',
  PRIMARY KEY (`Id`),
  KEY `UserId` (`UserId`),
  KEY `RoomId` (`RoomId`),
  KEY `CheckinTime` (`CheckinTime`),
  KEY `CheckoutTime` (`CheckoutTime`),
  KEY `StartTime` (`StartDateTime`) USING BTREE,
  KEY `EndTime` (`EndDateTime`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb3;

-- 테이블 데이터 MusicRoom.Reservation:~0 rows (대략적) 내보내기
DELETE FROM `Reservation`;
/*!40000 ALTER TABLE `Reservation` DISABLE KEYS */;
INSERT INTO `Reservation` (`Id`, `UserId`, `RoomId`, `StartDateTime`, `EndDateTime`, `CheckinTime`, `CheckoutTime`, `ExtendedMinutes`, `Status`, `CancelReason`, `Discard`) VALUES
	(1, 1, 1, '20251103153000', '20251103173000', '', '', 0, 0, 0, 0),
	(2, 2, 1, '20251104170000', '20251104180000', '', '', 0, 0, 0, 0);
/*!40000 ALTER TABLE `Reservation` ENABLE KEYS */;

-- 테이블 MusicRoom.Room 구조 내보내기
CREATE TABLE IF NOT EXISTS `Room` (
  `Id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `GroupId` bigint unsigned NOT NULL DEFAULT '0',
  `RoomName` varchar(50) NOT NULL DEFAULT '',
  `ReservableDaysMinOffset` tinyint unsigned NOT NULL DEFAULT '0',
  `ReservableDaysMaxOffset` tinyint unsigned NOT NULL DEFAULT '14',
  `Discard` tinyint unsigned NOT NULL DEFAULT '0',
  PRIMARY KEY (`Id`),
  KEY `RoomName` (`RoomName`),
  KEY `GroupId` (`GroupId`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb3;

-- 테이블 데이터 MusicRoom.Room:~0 rows (대략적) 내보내기
DELETE FROM `Room`;
/*!40000 ALTER TABLE `Room` DISABLE KEYS */;
INSERT INTO `Room` (`Id`, `GroupId`, `RoomName`, `ReservableDaysMinOffset`, `ReservableDaysMaxOffset`,  `Discard`) VALUES
	(1, 1, '18번드럼방', 0, 14, 0);
/*!40000 ALTER TABLE `Room` ENABLE KEYS */;

-- 테이블 MusicRoom.RoomGroup 구조 내보내기
CREATE TABLE IF NOT EXISTS `RoomGroup` (
  `Id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `HostId` bigint unsigned NOT NULL DEFAULT '0',
  `RoomGroupName` varchar(50) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '',
  `Address` varchar(500) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '',
  `CreatedAt` varchar(14) NOT NULL DEFAULT '',
  `Discard` tinyint unsigned NOT NULL DEFAULT '0',
  PRIMARY KEY (`Id`),
  KEY `HostId` (`HostId`),
  KEY `RoomGroupName` (`RoomGroupName`) USING BTREE,
  KEY `Address` (`Address`) USING BTREE,
  KEY `CreatedAt` (`CreatedAt`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb3 ROW_FORMAT=DYNAMIC;

-- 테이블 데이터 MusicRoom.RoomGroup:~0 rows (대략적) 내보내기
DELETE FROM `RoomGroup`;
/*!40000 ALTER TABLE `RoomGroup` DISABLE KEYS */;
INSERT INTO `RoomGroup` (`Id`, `HostId`, `RoomGroupName`, `Address`, `CreatedAt`, `Discard`) VALUES
	(1, 1, '부산대점', '부산대학로', '20251029220930', 0);
/*!40000 ALTER TABLE `RoomGroup` ENABLE KEYS */;

-- 테이블 MusicRoom.User 구조 내보내기
CREATE TABLE IF NOT EXISTS `User` (
  `Id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `UserName` varchar(50) NOT NULL DEFAULT '',
  `PhoneNo` varchar(11) NOT NULL DEFAULT '',
  `KakaoTalkId` varchar(50) NOT NULL DEFAULT '',
  `CreatedAt` varchar(14) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '',
  `Discard` tinyint unsigned NOT NULL DEFAULT '0',
  PRIMARY KEY (`Id`),
  KEY `UserName` (`UserName`),
  KEY `PhoneNo` (`PhoneNo`),
  KEY `KakaoTalkId` (`KakaoTalkId`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb3 COMMENT='사용자 테이블';

-- 테이블 데이터 MusicRoom.User:~0 rows (대략적) 내보내기
DELETE FROM `User`;
/*!40000 ALTER TABLE `User` DISABLE KEYS */;
INSERT INTO `User` (`Id`, `UserName`, `PhoneNo`, `KakaoTalkId`, `CreatedAt`, `Discard`) VALUES
	(1, '김수지', '01067891224', 'shjheee', '20251029221530', 0),
	(2, '홍길동', '01038841224', 'donggil56', '20251029171530', 0);
/*!40000 ALTER TABLE `User` ENABLE KEYS */;

/*!40101 SET SQL_MODE=IFNULL(@OLD_SQL_MODE, '') */;
/*!40014 SET FOREIGN_KEY_CHECKS=IFNULL(@OLD_FOREIGN_KEY_CHECKS, 1) */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40111 SET SQL_NOTES=IFNULL(@OLD_SQL_NOTES, 1) */;
