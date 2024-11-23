package openmeteogo

import (
	"context"
	"encoding/json"
	"testing"
)

func TestHourlyAirQualityResponse(t *testing.T) {
	client := NewClient(nil)
	ctx := context.Background()

	opts := &HourlyAirQualityOptions{
		Latitude:     52.52,
		Longitude:    13.41,
		ForecastDays: 1,
		Hourly: &[]OpenMeteoConst{
			HourlyAirQualityPm10,
			HourlyAirQualityPm25,
		},
	}

	gotResponse, err := client.HourlyAirQuality.Forecast(ctx, opts)

	if err != nil {
		t.Fatal("Unexpected error")
	}

	t.Logf("%+v", gotResponse)
	b, _ := json.Marshal(gotResponse)
	t.Logf("%s", b)
}
