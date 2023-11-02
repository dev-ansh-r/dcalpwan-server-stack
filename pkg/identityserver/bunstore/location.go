package store

import "go.thethings.network/lorawan-stack/v3/pkg/ttnpb"

// Location can be embedded in other models to add location columns.
type Location struct {
	Latitude  float64 `bun:"latitude,nullzero"`
	Longitude float64 `bun:"longitude,nullzero"`
	Altitude  int32   `bun:"altitude,nullzero"`
	Accuracy  int32   `bun:"accuracy,nullzero"`
}

func locationFromPB(pb *ttnpb.Location) Location {
	if pb == nil {
		return Location{}
	}
	return Location{
		Latitude:  pb.Latitude,
		Longitude: pb.Longitude,
		Altitude:  pb.Altitude,
		Accuracy:  pb.Accuracy,
	}
}

func locationToPB(m Location) *ttnpb.Location {
	if m.Latitude == 0 && m.Longitude == 0 && m.Altitude == 0 && m.Accuracy == 0 {
		return nil
	}
	return &ttnpb.Location{
		Latitude:  m.Latitude,
		Longitude: m.Longitude,
		Altitude:  m.Altitude,
		Accuracy:  m.Accuracy,
		Source:    ttnpb.LocationSource_SOURCE_REGISTRY,
	}
}
