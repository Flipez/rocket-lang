package object_test

import (
	"testing"
	"time"
)

func TestTimeModule(t *testing.T) {
	tests := []inputTestCase{
		{`Time.Layout`, time.Layout},
		{`Time.ANSIC`, time.ANSIC},
		{`Time.UnixDate`, time.UnixDate},
		{`Time.RubyDate`, time.RubyDate},
		{`Time.RFC822`, time.RFC822},
		{`Time.RFC822Z`, time.RFC822Z},
		{`Time.RFC850`, time.RFC850},
		{`Time.RFC1123`, time.RFC1123},
		{`Time.RFC1123Z`, time.RFC1123Z},
		{`Time.RFC3339`, time.RFC3339},
		{`Time.RFC3339Nano`, time.RFC3339Nano},
		{`Time.Kitchen`, time.Kitchen},
		{`Time.Stamp`, time.Stamp},
		{`Time.StampMilli`, time.StampMilli},
		{`Time.StampMicro`, time.StampMicro},
		{`Time.StampNano`, time.StampNano},
	}

	testInput(t, tests)
}

func TestTimeObjectMethods(t *testing.T) {
	tests := []inputTestCase{
		{`a = 1667144827; Time.format(a, "%a %b %e %M:%S %Y")`, "Sun Oct 30 47:07 2022"},
		{`a = 1667144827; Time.format(a, "Mon Jan _2 04:05 2006")`, "Sun Oct 30 47:07 2022"},
		{`a = "2022-03-23"; format = "2006-01-02"; Time.parse(a, format)`, 1647993600},
		{`a = "2022-03-23"; format = "09-01-02"; Time.parse(a, format)`, `Error while parsing time: parsing time "2022-03-23" as "09-01-02": cannot parse "2022-03-23" as "09-"`},
	}

	testInput(t, tests)
}
