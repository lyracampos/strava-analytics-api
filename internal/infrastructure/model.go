package infrastructure

import (
	"time"

	"github.com/lyracampos/strava-analytics-api/internal/domain/entities"
)

type StravaActivity struct {
	Id                 int64
	UserId             int64 `json:"user_id"`
	Name               string
	Distance           float64
	MovingTime         float64 `json:"moving_time"`
	ElapsedTime        float64 `json:"elapsed_time"`
	TotalElevationGain float64 `json:"total_elevation_gain"`
	Type               string
	SportType          string    `json:"sport_type"`
	WorkoutType        int32     `json:"workout_type"`
	StartDate          time.Time `json:"start_date"`
	UtcOffset          float64   `json:"utc_offset"`
	AchievementCount   int32     `json:"achievement_count"`
	KudosCount         int32     `json:"kudos_count"`
	Trainer            bool
	Commute            bool
	GearId             string `json:"gear_id"`
	Manual             bool
	AverageSpeed       float32 `json:"average_speed"`
	MaxSpeed           float32 `json:"max_speed"`
	AverageWatts       float32 `json:"average_watts"`
}

func activitiesConverter(stravaActivities []*StravaActivity) []*entities.Activity {
	activitiesList := make([]*entities.Activity, 0, len(stravaActivities))
	for _, item := range stravaActivities {
		activity := entities.Activity{
			Id:                 item.Id,
			UserId:             item.UserId,
			Title:              item.Name,
			Distance:           item.Distance,
			MovingTime:         item.MovingTime,
			ElapsedTime:        item.ElapsedTime,
			TotalElevationGain: item.TotalElevationGain,
			Type:               item.Type,
			SportType:          item.SportType,
			WorkoutType:        item.WorkoutType,
			StartDate:          item.StartDate,
			UtcOffset:          item.UtcOffset,
			AchievementCount:   item.AchievementCount,
			KudosCount:         item.KudosCount,
			Trainer:            item.Trainer,
			Commute:            item.Commute,
			GearId:             item.GearId,
			Manual:             item.Manual,
			AverageSpeed:       item.AverageSpeed,
			MaxSpeed:           item.MaxSpeed,
			AverageWatts:       item.AverageWatts,
		}
		activitiesList = append(activitiesList, &activity)
	}
	return activitiesList
}
