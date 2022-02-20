package goalgo

import (
	"reflect"
	"testing"
	"time"
)

func Test_roster_SameBirthday(t *testing.T) {
	type fields struct {
		list []*student
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &roster{
				list: tt.fields.list,
			}
			if got := r.SameBirthday(); got != tt.want {
				t.Errorf("roster.SameBirthday() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPastDay(t *testing.T) {
	type args struct {
		year  int
		month int
		day   int
	}
	tests := []struct {
		name string
		args args
		// want time.Time
	}{
		// TODO: Add test cases.
		{"0, 0, 0", args{year: 0, month: 0, day: 0}},
		{"2 years ago", args{year: 2, month: 0, day: 0}},
		{"2 days ago", args{year: 0, month: 0, day: 2}},
		{"2 months ago", args{year: 0, month: 2, day: 0}},
		{"2, 2, 2", args{year: 2, month: 2, day: 2}},
		{"-1, -1, -1", args{year: -1, month: -1, day: -1}},
		{"50,50,50", args{year: 50, month: 50, day: 50}},
		{"500, 500, 500", args{year: 500, month: 500, day: 500}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			yearNow, monthNow, dayNow := Now.Date()

			yearNew := yearNow - tt.args.year
			monthNew := monthNow - time.Month(tt.args.month)
			dayNew := dayNow - tt.args.day

			want := time.Date(yearNew, monthNew, dayNew, Now.Hour(), Now.Minute(), Now.Second(), Now.Nanosecond(), time.UTC)

			if got := PastDay(tt.args.year, tt.args.month, tt.args.day); !reflect.DeepEqual(got, want) {
				t.Errorf("PastDay(%v/%v/%v) = %v, want %v", tt.args.month, tt.args.day, tt.args.year, got, want)
			}
		})
	}
}
