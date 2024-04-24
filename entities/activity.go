package entities

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
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

type ListParams struct {
	Before  time.Time
	After   time.Time
	Page    int
	PerPage int
}

func List(listParams ListParams) ([]*Activity, error) {
	stravaUrl := fmt.Sprintf(
		"https://www.strava.com/api/v3/athlete/activities?after=%d&before=%d",
		listParams.After.Unix(),
		listParams.Before.Unix(),
	)
	req, err := http.NewRequest(http.MethodGet, stravaUrl, nil)
	if err != nil {
		return []*Activity{}, fmt.Errorf("erro ao tentar criar Request: %w", err)
	}

	req.Header.Add(
		"Authorization",
		"Bearer ee16a0f42fdcbb0c51f4c57ec6ef5f1f3d329f8d",
	)

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return []*Activity{}, fmt.Errorf("erro ao fazer o Request: %w", err)
	}
	if res.StatusCode != http.StatusOK {
		return []*Activity{}, fmt.Errorf("erro ao fazer o Request: %w", err)
	}

	resJson, err := io.ReadAll(res.Body)
	if err != nil {
		return []*Activity{}, fmt.Errorf("erro ao carregar o response: %w", err)
	}

	var stravaActivitiesList []*StravaActivity
	err = json.Unmarshal(resJson, &stravaActivitiesList)
	if err != nil {
		return []*Activity{}, fmt.Errorf("erro ao interpretar o response: %w", err)
	}
	activitiesList := activitiesConverter(stravaActivitiesList)
	return activitiesList, nil
}

func activitiesConverter(stravaActivitiesList []*StravaActivity) []*Activity {
	activitiesList := make([]*Activity, 0, len(stravaActivitiesList))
	for _, item := range stravaActivitiesList {
		activity := Activity{
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
