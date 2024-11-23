package openmeteogo

import (
	"context"
	"fmt"
	"time"
)

type HistoricalWeatherService service

const (
	HistoricalTemperature2m       OpenMeteoConst = "temperature_2m"
	HistoricalRelativeHumidity2m  OpenMeteoConst = "relative_humidity_2m"
	HistoricalDewpoint2m          OpenMeteoConst = "dew_point_2m"
	HistoricalApparentTemperature OpenMeteoConst = "apparent_temperature"
	HistoricalPrecipitation       OpenMeteoConst = "precipitation"
	HistoricalRain                OpenMeteoConst = "rain"
	HistoricalSnowfall            OpenMeteoConst = "snowfall"
	HistoricalCloudCover          OpenMeteoConst = "cloud_cover"
	HistoricalWindSpeed10m        OpenMeteoConst = "wind_speed_10m"
	HistoricalWindDirection10m    OpenMeteoConst = "wind_direction_10m"
)

type HistoricalOptions struct {
	Options
	Latitude      float64           `url:"latitude"`
	Longitude     float64           `url:"longitude"`
	StartDate     string            `url:"start_date"` // YYYY-MM-DD
	EndDate       string            `url:"end_date"`   // YYYY-MM-DD
	Hourly        *[]OpenMeteoConst `url:"hourly,omitempty"`
	Daily         *[]OpenMeteoConst `url:"daily,omitempty"`
	Timezone      string            `url:"timezone,omitempty"`
	CellSelection string            `url:"cell_selection,omitempty"`
}

type HistoricalWeatherResponse struct {
	Latitude             float64                 `json:"latitude"`
	Longitude            float64                 `json:"longitude"`
	Generationtime_ms    float64                 `json:"generationtime_ms"`
	UtcOffsetSeconds     int                     `json:"utc_offset_seconds"`
	Timezone             string                  `json:"timezone"`
	TimezoneAbbreviation string                  `json:"timezone_abbreviation"`
	Elevation            float64                 `json:"elevation"`
	HourlyUnits          HistoricalUnitsResponse `json:"hourly_units,omitempty"`
	Hourly               HistoricalResponse      `json:"hourly,omitempty"`
	DailyUnits           HistoricalUnitsResponse `json:"daily_units,omitempty"`
	Daily                HistoricalResponse      `json:"daily,omitempty"`
}

type HistoricalUnitsResponse struct {
	Time               string  `json:"time"`
	Temperature2m      *string `json:"temperature_2m,omitempty"`
	RelativeHumidity2m *string `json:"relative_humidity_2m,omitempty"`
	Dewpoint2m         *string `json:"dew_point_2m,omitempty"`
	Precipitation      *string `json:"precipitation,omitempty"`
	Rain               *string `json:"rain,omitempty"`
	Snowfall           *string `json:"snowfall,omitempty"`
	CloudCover         *string `json:"cloud_cover,omitempty"`
	WindSpeed10m       *string `json:"wind_speed_10m,omitempty"`
	WindDirection10m   *string `json:"wind_direction_10m,omitempty"`
}

type HistoricalResponse struct {
	Time               []string   `json:"time"`
	Temperature2m      []*float64 `json:"temperature_2m,omitempty"`
	RelativeHumidity2m []*float64 `json:"relative_humidity_2m,omitempty"`
	Dewpoint2m         []*float64 `json:"dew_point_2m,omitempty"`
	Precipitation      []*float64 `json:"precipitation,omitempty"`
	Rain               []*float64 `json:"rain,omitempty"`
	Snowfall           []*float64 `json:"snowfall,omitempty"`
	CloudCover         []*float64 `json:"cloud_cover,omitempty"`
	WindSpeed10m       []*float64 `json:"wind_speed_10m,omitempty"`
	WindDirection10m   []*float64 `json:"wind_direction_10m,omitempty"`
}

// Validate checks if the options are valid
func (opts *HistoricalOptions) Validate() error {
	if opts.Latitude < -90 || opts.Latitude > 90 {
		return fmt.Errorf("latitude must be between -90 and 90")
	}
	if opts.Longitude < -180 || opts.Longitude > 180 {
		return fmt.Errorf("longitude must be between -180 and 180")
	}

	startDate, err := time.Parse("2006-01-02", opts.StartDate)
	if err != nil {
		return fmt.Errorf("invalid start_date format: %v", err)
	}

	endDate, err := time.Parse("2006-01-02", opts.EndDate)
	if err != nil {
		return fmt.Errorf("invalid end_date format: %v", err)
	}

	if endDate.Before(startDate) {
		return fmt.Errorf("end_date must be after start_date")
	}

	// Check if date range is within allowed limits (max 1 year)
	if endDate.Sub(startDate) > 365*24*time.Hour {
		return fmt.Errorf("date range must not exceed 1 year")
	}

	return nil
}

// Archive retrieves historical weather data for a specific time period
func (service *HistoricalWeatherService) Archive(ctx context.Context, opts *HistoricalOptions) (*HistoricalWeatherResponse, error) {
	if err := opts.Validate(); err != nil {
		return nil, err
	}

	u, err := addOptions("archive", opts)
	if err != nil {
		return nil, err
	}

	req, err := service.client.NewRequest("GET", service.client.HistoricalBaseURL, u, nil)
	if err != nil {
		return nil, err
	}

	result := new(HistoricalWeatherResponse)
	_, err = service.client.Do(ctx, req, result)
	return result, err
}
