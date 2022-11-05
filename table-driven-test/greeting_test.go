package main

import (
	"testing"
	"time"
)

func TestDo(t *testing.T) {
	tests := map[string]struct {
		in   string
		want string
	}{
		"3:59":  {"3:59", "こんばんは"},
		"4:00":  {"4:00", "おはよう"},
		"9:59":  {"9:59", "おはよう"},
		"10:00": {"10:00", "こんにちは"},
		"16:59": {"16:59", "こんにちは"},
		"17:00": {"17:00", "こんばんは"},
	}
	for tc, tt := range tests {
		tt := tt
		t.Run(tc, func(t *testing.T) {
			if got := Do(toDate(t, tt.in)); got != tt.want {
				t.Errorf("Do() = %v, want %v", got, tt.want)
			}
		})
	}
}

func toDate(t *testing.T, testTime string) time.Time {
	t.Helper()
	d, err := time.Parse("15:04", testTime)
	if err != nil {
		t.Fatalf("toDate: %v", err)
	}
	return d
}
