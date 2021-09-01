package bmark

import (
	"sync"
	"testing"
)

type TripCreated struct {
	ID   string `json:"id"`
	Type string `json:"type"`
	Body struct {
		Module      int  `json:"module"`
		BookedBy    int  `json:"booked_by"`
		TripID      int  `json:"trip_id"`
		VehicleType int  `json:"vehicle_type"`
		PreBooking  bool `json:"pre_booking"`
		Passenger   struct {
			ID int `json:"id"`
		} `json:"passenger"`
		Driver struct {
			ID int `json:"id"`
		} `json:"driver"`
		Corporate struct {
			ID    int `json:"id"`
			DepID int `json:"dep_id"`
		} `json:"corporate"`
		Pickup struct {
			Time     int `json:"time"`
			Location []struct {
				Address string  `json:"address"`
				Lat     float64 `json:"lat"`
				Lng     float64 `json:"lng"`
			} `json:"location"`
		} `json:"pickup"`
		Drop struct {
			Location []struct {
				Address string  `json:"address"`
				Lat     float64 `json:"lat"`
				Lng     float64 `json:"lng"`
			} `json:"location"`
		} `json:"drop"`
		Promotion struct {
			Code string `json:"code"`
		} `json:"promotion"`
		Region struct {
			Ids []int `json:"ids"`
		} `json:"region"`
		Payment struct {
			PrimaryMethod   int `json:"primary_method"`
			SecondaryMethod int `json:"secondary_method"`
		} `json:"payment"`
		Comments struct {
			Remark      string `json:"remark"`
			DriverNotes string `json:"driver_notes"`
		} `json:"comments"`
		Filters struct {
			Driver struct {
				LanguageID int `json:"language_id"`
			} `json:"driver"`
			Vehicle struct {
				CompanyID int `json:"company_id"`
				BrandID   int `json:"brand_id"`
				ColorID   int `json:"color_id"`
			} `json:"vehicle"`
		} `json:"filters"`
		Surge struct {
			RegionID int `json:"region_id"`
			Value    int `json:"value"`
		} `json:"surge"`
		FareDetails struct {
			FareType         string  `json:"fare_type"`
			MinKm            float64 `json:"min_km"`
			MinFare          float64 `json:"min_fare"`
			AdditionalKmFare float64 `json:"additional_km_fare"`
			WaitingTimeFare  float64 `json:"waiting_time_fare"`
			FreeWaitingTime  int     `json:"free_waiting_time"`
			NightFare        float64 `json:"night_fare"`
			RideHours        float64 `json:"ride_hours"`
			ExtraRideFare    float64 `json:"extra_ride_fare"`
			DriverBata       float64 `json:"driver_bata"`
			TripType         int     `json:"trip_type"`
		} `json:"fare_details"`
	} `json:"body"`
	CreatedAt int64 `json:"created_at"`
	Expiry    int64 `json:"expiry"`
	Version   int   `json:"version"`
	TraceInfo struct {
		TraceID struct {
			High int64 `json:"high"`
			Low  int64 `json:"low"`
		} `json:"trace_id"`
		SpanID   int64 `json:"span_id"`
		ParentID int64 `json:"parent_id"`
		Sampled  bool  `json:"sampled"`
	} `json:"trace_info"`
}

var pool = &sync.Pool{
	New: func() interface{} {
		return new(TripCreated)
	},
}

var trip *TripCreated

func BenchmarkWithoutPool(b *testing.B) {
	for i := 0; i < b.N; i++ {
		trip = new(TripCreated)
		trip.Version = 1
		trip.CreatedAt = 1
	}
}

func BenchmarkWithPool(b *testing.B) {
	for i := 0; i < b.N; i++ {
		trip = pool.Get().(*TripCreated)
		trip.Version = 1
		trip.CreatedAt = 1

		// Returing the allocation to the pool once the task is done.
		pool.Put(trip)
	}
}
