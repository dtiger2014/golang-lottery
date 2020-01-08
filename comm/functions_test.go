package comm

import (
	"golang-lottery/conf"
	"time"
	"testing"
)


func TestNowUnix(t *testing.T) {
	nowUnix := NowUnix()

	cur := int(time.Now().In(conf.SysTimeLocation).Unix())
	if nowUnix != cur {
		t.Errorf("result %d, cur %d\n",nowUnix, cur)
	}
}

func TestFormatFromUnixTime(t *testing.T) {
	tests := []struct{
		t int64
		date string
	} {
		{1234567890, "2009-02-14 07:31:30"},
		{1553672819, "2019-03-27 15:46:59"},
		// {0,"2020-01-07 23:30:16"},
	}

	for _, test := range tests {
		result := FormatFromUnixTime(test.t)
		if result != test.date {
			t.Errorf("result %s, date %s", result, test.date)
		}
	}
}

func BenchmarkFormatFromUnixTime(b *testing.B) {
	t := int64(1234567890)
	date := "2009-02-14 07:31:30"

	for i := 0; i < b.N; i++ {
		result := FormatFromUnixTime(t)
		if result != date {
			b.Errorf("result %s, date %s", result, date)
		}
	}
	
}