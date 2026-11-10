// Copyright (c) 2026 heathcetide. All rights reserved.

package models

import (
	"time"

	"github.com/heathcetide/infectious-disease/internal/constants"
)

// WorkEnvironmentInfo is monthly work-environment data for a user.
type WorkEnvironmentInfo struct {
	EnvironmentID uint `json:"environmentId" gorm:"primaryKey;autoIncrement"`
	UserID        uint `json:"userId" gorm:"index;not null"`

	YearMonth      string    `json:"yearMonth" gorm:"size:7;index"`
	SubmissionTime time.Time `json:"submissionTime"`

	WorkStationAltitude     int     `json:"workStationAltitude"`
	WorkStationName         string  `json:"workStationName" gorm:"size:200"`
	DormitoryBuildingNumber int     `json:"dormitoryBuildingNumber"`
	DormitoryRoomNumber     int     `json:"dormitoryRoomNumber"`
	DormitoryOccupants      int     `json:"dormitoryOccupants"`
	DormitoryArea           float64 `json:"dormitoryArea"`
	WindowOpenArea          float64 `json:"windowOpenArea"`
	VentilationCondition    string  `json:"ventilationCondition" gorm:"size:200"`

	AvgWorkRoomPeople           int `json:"avgWorkRoomPeople"`
	AvgWorkFaceToFace           int `json:"avgWorkFaceToFace"`
	AvgWorkPhysicalContact      int `json:"avgWorkPhysicalContact"`
	AvgBreakfastRoomPeople      int `json:"avgBreakfastRoomPeople"`
	AvgBreakfastFaceToFace      int `json:"avgBreakfastFaceToFace"`
	AvgBreakfastPhysicalContact int `json:"avgBreakfastPhysicalContact"`
	AvgLunchRoomPeople          int `json:"avgLunchRoomPeople"`
	AvgLunchFaceToFace          int `json:"avgLunchFaceToFace"`
	AvgLunchPhysicalContact     int `json:"avgLunchPhysicalContact"`
	AvgDinnerRoomPeople         int `json:"avgDinnerRoomPeople"`
	AvgDinnerFaceToFace         int `json:"avgDinnerFaceToFace"`
	AvgDinnerPhysicalContact    int `json:"avgDinnerPhysicalContact"`

	CollectiveActivitiesCount  int `json:"collectiveActivitiesCount"`
	AvgActivityRoomPeople      int `json:"avgActivityRoomPeople"`
	AvgActivityFaceToFace      int `json:"avgActivityFaceToFace"`
	AvgActivityPhysicalContact int `json:"avgActivityPhysicalContact"`

	AvgMonthlyPressure float64 `json:"avgMonthlyPressure"`
	MaxTemperature     float64 `json:"maxTemperature"`
	MinTemperature     float64 `json:"minTemperature"`
}

func (WorkEnvironmentInfo) TableName() string { return constants.TableWorkEnvironmentInfo }
