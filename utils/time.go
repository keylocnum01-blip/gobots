package utils

import (
	"fmt"
	"time"
)

func FmtDuration(d time.Duration) string {
	d = d.Round(time.Second)
	x := d
	h := d / time.Hour
	d -= h * time.Hour
	m := d / time.Minute
	d -= m * time.Minute
	if x < 60*time.Second {
		return fmt.Sprintf("%v", x)
	} else if x < 3600*time.Second {
		return fmt.Sprintf("%02dMinutes's", m)
	} else if x < 86400*time.Second {
		return fmt.Sprintf("%02dHour's %02dMinute's", h%24, m)
	} else {
		return fmt.Sprintf("%02dDay's %02dHour's %02dMinute's", h/24, h%24, m)
	}
}
