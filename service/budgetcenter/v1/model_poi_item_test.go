package v1

import (
	"encoding/json"
	"testing"
)

func TestPoiItem_JSONUnmarshal(t *testing.T) {
	tests := []struct {
		name       string
		jsonStr    string
		wantCity   string
		wantCityId int32
		wantFlat   float64
		wantLabel  string
	}{
		{
			name:       "full fields",
			jsonStr:    `{"city":"北京","city_id":1,"city_adcode":"110000","flat":39.9,"flng":116.4,"poi_range":500,"label":"国贸"}`,
			wantCity:   "北京",
			wantCityId: 1,
			wantFlat:   39.9,
			wantLabel:  "国贸",
		},
		{
			name:       "empty object",
			jsonStr:    `{}`,
			wantCity:   "",
			wantCityId: 0,
			wantFlat:   0,
			wantLabel:  "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var item PoiItem
			if err := json.Unmarshal([]byte(tt.jsonStr), &item); err != nil {
				t.Fatalf("unmarshal failed: %v", err)
			}
			if tt.wantCity != "" {
				if item.City == nil || *item.City != tt.wantCity {
					t.Errorf("City = %v, want %v", item.City, tt.wantCity)
				}
			} else {
				if item.City != nil {
					t.Errorf("City = %v, want nil", item.City)
				}
			}
			if tt.wantCityId != 0 {
				if item.CityId == nil || *item.CityId != tt.wantCityId {
					t.Errorf("CityId = %v, want %v", item.CityId, tt.wantCityId)
				}
			}
			if tt.wantFlat != 0 {
				if item.Flat == nil || *item.Flat != tt.wantFlat {
					t.Errorf("Flat = %v, want %v", item.Flat, tt.wantFlat)
				}
			}
			if tt.wantLabel != "" {
				if item.Label == nil || *item.Label != tt.wantLabel {
					t.Errorf("Label = %v, want %v", item.Label, tt.wantLabel)
				}
			}
		})
	}
}

func TestPoiItemBuilder(t *testing.T) {
	item := NewPoiItemBuilder().
		City("北京").
		CityId(1).
		CityAdcode("110000").
		Flat(39.9).
		Flng(116.4).
		PoiRange(500).
		Label("国贸").
		Build()

	if item.City == nil || *item.City != "北京" {
		t.Errorf("City = %v, want 北京", item.City)
	}
	if item.CityId == nil || *item.CityId != 1 {
		t.Errorf("CityId = %v, want 1", item.CityId)
	}
	if item.CityAdcode == nil || *item.CityAdcode != "110000" {
		t.Errorf("CityAdcode = %v, want 110000", item.CityAdcode)
	}
	if item.Flat == nil || *item.Flat != 39.9 {
		t.Errorf("Flat = %v, want 39.9", item.Flat)
	}
	if item.Flng == nil || *item.Flng != 116.4 {
		t.Errorf("Flng = %v, want 116.4", item.Flng)
	}
	if item.PoiRange == nil || *item.PoiRange != 500 {
		t.Errorf("PoiRange = %v, want 500", item.PoiRange)
	}
	if item.Label == nil || *item.Label != "国贸" {
		t.Errorf("Label = %v, want 国贸", item.Label)
	}
}
