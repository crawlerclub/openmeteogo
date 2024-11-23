package openmeteogo

import (
	"context"
	"encoding/json"
	"testing"
)

func TestCurrentAirQualityResponse(t *testing.T) {
	client := NewClient(nil)
	ctx := context.Background()

	opts := &CurrentAirQualityOptions{
		Latitude:     52.52,
		Longitude:    13.41,
		ForecastDays: 1,
		Domains:      AirQualityDomainEurope,
		Current: &[]OpenMeteoConst{
			CurrentAirQualityPm10,
			CurrentAirQualityPm25,
		},
	}

	gotResponse, err := client.CurrentAirQuality.Forecast(ctx, opts)

	if err != nil {
		t.Fatal("Unexpected error")
	}
	t.Logf("%+v", gotResponse)
	b, _ := json.Marshal(gotResponse)
	t.Logf("%s", b)
}
