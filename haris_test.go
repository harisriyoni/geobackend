package geobackend

import (
	"fmt"
	"testing"
)

var dbname = "gis"
var collname = "gis"

func TestGeoIntersects(t *testing.T) {
	mconn := SetConnection("MONGOSTRING", dbname)
	coordinates := Polygon{
		Coordinates: [][][]float64{
			{
				{106.80393025390346, -6.1789031084327775},
				{106.85007925717223, -6.177405314931775},
				{106.84728059992784, -6.190286152748811},
				{106.80777316130363, -6.18985446949516},
				{106.80393025390346, -6.1789031084327775},
			},
		},
	}
	datagedung := GeoIntersects(mconn, collname, coordinates)
	fmt.Println(datagedung)
}

func TestGeoWithin(t *testing.T) {
	mconn := SetConnection("MONGOSTRING", dbname)
	coordinates := Polygon{
		Coordinates: [][][]float64{
			{
				{106.83515985012605, -6.166125656657883},
				{106.83582623098323, -6.1678923886587711},
				{106.83893600831055, -6.1663096915163464},
				{106.83834366977129, -6.1641380761234075},
				{106.83515985012605, -6.166125656657883},
			},
		},
	}
	datagedung := GeoWithin(mconn, collname, coordinates)
	fmt.Println(datagedung)
}

func TestNear(t *testing.T) {
	mconn := SetConnection2dsphere("MONGOSTRING", "gis", "gis")
	coordinates := Point{
		Coordinates: []float64{
			106.80899259801674, -6.168340926357345,
		},
	}
	datagedung := Near(mconn, "gis", coordinates)
	fmt.Println(datagedung)
}

func TestNearSphere(t *testing.T) {
	mconn := SetConnection("MONGOSTRING", "gis")
	coordinates := Point{
		Coordinates: []float64{
			106.81781953215608, -6.167583300161709,
		},
	}
	datagedung := NearSphere(mconn, "gis", coordinates)
	fmt.Println(datagedung)
}

func TestBox(t *testing.T) {
	mconn := SetConnection("MONGOSTRING", "gis")
	coordinates := Polyline{
		Coordinates: [][]float64{
			{106.8200974960248, -6.163418436577771},
			{106.82086127204684, -6.167848024969331},
		},
	}
	datagedung := Box(mconn, collname, coordinates)
	fmt.Println(datagedung)
}
