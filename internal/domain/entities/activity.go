package entities

import (
	"time"
)

type Activity struct {
	Id                 int64
	UserId             int64
	Title              string
	Distance           float64
	MovingTime         float64
	ElapsedTime        float64
	TotalElevationGain float64
	Type               string
	SportType          string
	WorkoutType        int32
	StartDate          time.Time
	UtcOffset          float64
	AchievementCount   int32
	KudosCount         int32
	Trainer            bool
	Commute            bool
	GearId             string
	Manual             bool
	AverageSpeed       float32
	MaxSpeed           float32
	AverageWatts       float32
}
