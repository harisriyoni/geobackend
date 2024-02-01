package gisbackend

import (
	"fmt"
	"testing"
)

var dbname = "gis"
var collname = "gis"

func TestGeoIntersects(t *testing.T) {
	mconn := SetConnection("MONGOSTRING", dbname)
	coordinates := Point{
		Coordinates: []float64{
			106.8200974960248, -6.163418436577771,
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
				{95.31123456789012, 5.553210987654321},
				{95.31133456789011, 5.553210987654321},
				{95.31133456789011, 5.553310987654321},
				{95.31123456789012, 5.553310987654321},
				{95.31123456789012, 5.553210987654321},
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
			95.30987654321098, 5.556789012345678,
		},
	}
	datagedung := Near(mconn, "gis", coordinates)
	fmt.Println(datagedung)
}

func TestNearSphere(t *testing.T) {
	mconn := SetConnection("MONGOSTRING", "gis")
	coordinates := Point{
		Coordinates: []float64{
			95.30987654321098, 5.556789012345678,
		},
	}
	datagedung := NearSphere(mconn, "gis", coordinates)
	fmt.Println(datagedung)
}

func TestBox(t *testing.T) {
	mconn := SetConnection("MONGOSTRING", "gis")
	coordinates := Polyline{
		Coordinates: [][]float64{
			{95.32345678901234, 5.567890123456789},
			{95.32355678901234, 5.567990123456789},
		},
	}
	datagedung := Box(mconn, collname, coordinates)
	fmt.Println(datagedung)
}
