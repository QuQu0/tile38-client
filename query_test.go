package t38c

import (
	"strings"
	"testing"
)

func TestInwQueryBuilder(t *testing.T) {
	search := &Search{}
	tests := []struct {
		Query    InwQueryBuilder
		Expected string
	}{
		{
			Query: search.Nearby("fleet", 10, 20, 30).
				Where("speed", 10, 20).
				Wherein("speed", 10, 20, 30).
				Match("abc*").
				Cursor(10).
				Format(FormatIDs).
				Limit(5),
			Expected: "NEARBY fleet WHERE speed 10 20 WHEREIN speed 3 10 20 30 MATCH abc* CURSOR 10 LIMIT 5 IDS POINT 10 20 30",
		},
		{
			Query: search.Intersects("fleet").
				Tile(10, 20, 30).
				Match("abc*"),
			Expected: "INTERSECTS fleet MATCH abc* TILE 10 20 30",
		},
	}

	for _, test := range tests {
		cmd := *test.Query.toCmd()
		actual := append([]string{cmd.Name}, cmd.Args...)
		expected := strings.Split(test.Expected, " ")
		if len(actual) != len(expected) {
			t.Fatalf("not equal: bad length (%d, %d)\n"+
				"expected: %s\n"+
				"actual  : %s\n", len(expected), len(actual), expected, actual)
		}
		for i := 0; i < len(actual); i++ {
			if actual[i] != expected[i] {
				t.Fatalf("not equal:\n"+
					"expected: %s\n"+
					"actual  : %s\n", expected, actual)
			}
		}
	}
}

func TestGeofenceQueryBuilder(t *testing.T) {
	geofence := &Geofence{}
	tests := []struct {
		Query    GeofenceQueryBuilder
		Expected string
	}{
		{
			Query: geofence.Nearby("fleet", 10, 20, 30).
				Actions(Enter, Exit, Cross).
				Clip().
				Commands(Set, Del).
				Cursor(5).
				Format(FormatHashes(5)),
			Expected: "NEARBY fleet CLIP CURSOR 5 FENCE DETECT enter,exit,cross COMMANDS set,del HASHES 5 POINT 10 20 30",
		},
		{
			Query: geofence.Roam("agent", "target", "*", 100).
				Distance().
				Wherein("price", 20, 30),
			Expected: "NEARBY agent DISTANCE WHEREIN price 2 20 30 FENCE ROAM target * 100",
		},
	}

	for _, test := range tests {
		cmd := *test.Query.toCmd()
		actual := append([]string{cmd.Name}, cmd.Args...)
		expected := strings.Split(test.Expected, " ")
		if len(actual) != len(expected) {
			t.Fatalf("not equal: bad length (%d, %d)\n"+
				"expected: %s\n"+
				"actual  : %s\n", len(expected), len(actual), expected, actual)
		}
		for i := 0; i < len(actual); i++ {
			if actual[i] != expected[i] {
				t.Fatalf("not equal:\n"+
					"expected: %s\n"+
					"actual  : %s\n", expected, actual)
			}
		}
	}
}

func TestSetQueryBuilder(t *testing.T) {
	keys := &Keys{}
	tests := []struct {
		Query    SetQueryBuilder
		Expected string
	}{
		{
			Query: keys.Set("agent", "47").
				PointZ(0, 0, -20).
				Field("age", 55).
				Expiration(60 * 60 * 24 * 365),
			Expected: "SET agent 47 EX 31536000 FIELD age 55 POINT 0 0 -20",
		},
	}

	for _, test := range tests {
		cmd := *test.Query.toCmd()
		actual := append([]string{cmd.Name}, cmd.Args...)
		expected := strings.Split(test.Expected, " ")
		if len(actual) != len(expected) {
			t.Fatalf("not equal: bad length (%d, %d)\n"+
				"expected: %s\n"+
				"actual  : %s\n", len(expected), len(actual), expected, actual)
		}
		for i := 0; i < len(actual); i++ {
			if actual[i] != expected[i] {
				t.Fatalf("not equal:\n"+
					"expected: %s\n"+
					"actual  : %s\n", expected, actual)
			}
		}
	}
}
