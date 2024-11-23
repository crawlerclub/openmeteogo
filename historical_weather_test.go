package openmeteogo

import (
	"context"
	"encoding/json"
	"testing"
)

func TestHistoricalWeatherValidation(t *testing.T) {
	tests := []struct {
		name    string
		opts    *HistoricalOptions
		wantErr bool
	}{
		{
			name: "valid options",
			opts: &HistoricalOptions{
				Latitude:  41.902782,
				Longitude: 12.496366,
				StartDate: "2023-01-01",
				EndDate:   "2023-12-31",
			},
			wantErr: false,
		},
		{
			name: "invalid latitude",
			opts: &HistoricalOptions{
				Latitude:  91.0,
				Longitude: 12.496366,
				StartDate: "2023-01-01",
				EndDate:   "2023-12-31",
			},
			wantErr: true,
		},
		{
			name: "invalid date format",
			opts: &HistoricalOptions{
				Latitude:  41.902782,
				Longitude: 12.496366,
				StartDate: "2023/01/01",
				EndDate:   "2023-12-31",
			},
			wantErr: true,
		},
		{
			name: "end date before start date",
			opts: &HistoricalOptions{
				Latitude:  41.902782,
				Longitude: 12.496366,
				StartDate: "2023-12-31",
				EndDate:   "2023-01-01",
			},
			wantErr: true,
		},
		{
			name: "date range exceeds 1 year",
			opts: &HistoricalOptions{
				Latitude:  41.902782,
				Longitude: 12.496366,
				StartDate: "2023-01-01",
				EndDate:   "2024-01-02",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.opts.Validate()
			if (err != nil) != tt.wantErr {
				t.Errorf("HistoricalOptions.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestHistoricalWeatherArchive(t *testing.T) {

	client := NewClient(nil)

	opts := &HistoricalOptions{
		Latitude:  41.902782,
		Longitude: 12.496366,
		StartDate: "2023-01-01",
		EndDate:   "2023-01-02",
		Hourly: &[]OpenMeteoConst{
			HistoricalTemperature2m,
			HistoricalRelativeHumidity2m,
			HistoricalDewpoint2m,
			HistoricalApparentTemperature,
			HistoricalPrecipitation,
			HistoricalRain,
			HistoricalSnowfall,
			HistoricalCloudCover,
			HistoricalWindSpeed10m,
			HistoricalWindDirection10m,
		},
	}

	ctx := context.Background()
	got, err := client.HistoricalWeather.Archive(ctx, opts)
	if err != nil {
		t.Errorf("HistoricalWeather.Archive returned error: %v", err)
	}
	t.Logf("%+v", got)
	b, _ := json.Marshal(got)
	t.Logf("%s", b)
}
