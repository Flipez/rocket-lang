package stdlib

import (
	"strings"
	"time"

	"github.com/flipez/rocket-lang/object"
)

var timeFunctions = map[string]*object.BuiltinFunction{}
var timeProperties = map[string]*object.BuiltinProperty{}

var timeFormatConversions = map[string]string{
	// year format
	"%Y": "2006", // Four-digit year
	"%y": "06",   // Two digit year

	// month format
	"%B":  "January", // Full month name
	"%b":  "Jan",     // Three-letter abbr. of the month
	"%m":  "01",      // Two digit month (leading 0)
	"%-m": "1",       //  At most two-digit month (without leading 0)

	// day format
	"%A":  "Monday", // Full weekday
	"%a":  "Mon",    // Three letter weekday
	"%d":  "02",     // Two-digit month day
	"%e":  "_2",     //  Two-character month day with leading space
	"%-d": "2",      // At most two-digit month day
	"%j":  "002",    //  Three digit day of the year

	// hour format
	"%H": "15", // Two-digit 24h format hour
	"%I": "03", // Two digit 12h format hour (with a leading 0 if necessary)
	"%l": "3",  // At most two-digit 12h format hour (without a leading 0)
	"%p": "PM", // AM/PM mark (uppercase)
	"%P": "pm", // AM/PM mark (lowercase)

	// minute format
	"%M": "04", // Two-digit minute (with a leading 0 if necessary)

	// second format
	"%S": "05", // Two-digit second (with a leading 0 if necessary)

	// time zone format
	"%Z":   "MST",       // Abbreviation of the time zone
	"%::z": "-07:00:00", // Numeric time zone offset with hours, minutes, and seconds separated by colon
	"%z":   "-0700",     // Numeric time zone offset with hours and minutes
	"%:z":  "-07:00",    // Numeric time zone offset with hours and minutes separated by colons
}

func init() {
	timeFunctions["format"] = object.NewBuiltinFunction(
		"format",
		object.MethodLayout{
			ArgPattern: object.Args(
				object.Arg(object.INTEGER_OBJ),
				object.Arg(object.STRING_OBJ),
			),
			ReturnPattern: object.Args(
				object.Arg(object.STRING_OBJ),
			),
		},
		func(_ object.Environment, args ...object.Object) object.Object {
			unixTimestamp := args[0].(*object.Integer)
			timeFormat := args[1].(*object.String)

			time := time.Unix(int64(unixTimestamp.Value), 0)

			return object.NewString(time.Format(convertTimeFormat(timeFormat)))
		})
	timeFunctions["parse"] = object.NewBuiltinFunction(
		"parse",
		object.MethodLayout{
			ArgPattern: object.Args(
				object.Arg(object.STRING_OBJ),
				object.Arg(object.STRING_OBJ),
			),
			ReturnPattern: object.Args(
				object.Arg(object.STRING_OBJ),
			),
		},
		func(_ object.Environment, args ...object.Object) object.Object {
			timeString := args[0].(*object.String)
			timeFormat := args[1].(*object.String)

			timeParsed, err := time.Parse(convertTimeFormat(timeFormat), timeString.Value)
			if err != nil {
				return object.NewErrorFormat("Error while parsing time: %s", err)
			}

			return object.NewInteger(int(timeParsed.Unix()))
		})
	timeFunctions["sleep"] = object.NewBuiltinFunction(
		"sleep",
		object.MethodLayout{
			ArgPattern: object.Args(
				object.Arg(object.INTEGER_OBJ),
			),
			ReturnPattern: object.Args(
				object.Arg(object.NIL_OBJ),
			),
		},
		func(_ object.Environment, args ...object.Object) object.Object {
			time.Sleep(time.Duration(args[0].(*object.Integer).Value) * time.Second)
			return object.NIL
		})
	timeFunctions["unix"] = object.NewBuiltinFunction(
		"unix",
		object.MethodLayout{
			ReturnPattern: object.Args(
				object.Arg(object.INTEGER_OBJ),
			),
		},
		func(_ object.Environment, args ...object.Object) object.Object {
			return object.NewInteger(int(time.Now().Unix()))
		})

	timeProperties["Layout"] = object.NewBuiltinProperty("Layout", object.NewString(time.Layout))
	timeProperties["ANSIC"] = object.NewBuiltinProperty("ANSIC", object.NewString(time.ANSIC))
	timeProperties["UnixDate"] = object.NewBuiltinProperty("UnixDate", object.NewString(time.UnixDate))
	timeProperties["RubyDate"] = object.NewBuiltinProperty("RubyDate", object.NewString(time.RubyDate))
	timeProperties["RFC822"] = object.NewBuiltinProperty("RFC822", object.NewString(time.RFC822))
	timeProperties["RFC822Z"] = object.NewBuiltinProperty("RFC822Z", object.NewString(time.RFC822Z))
	timeProperties["RFC850"] = object.NewBuiltinProperty("RFC850", object.NewString(time.RFC850))
	timeProperties["RFC1123"] = object.NewBuiltinProperty("RFC1123", object.NewString(time.RFC1123))
	timeProperties["RFC1123Z"] = object.NewBuiltinProperty("RFC1123Z", object.NewString(time.RFC1123Z))
	timeProperties["RFC3339"] = object.NewBuiltinProperty("RFC3339", object.NewString(time.RFC3339))
	timeProperties["RFC3339Nano"] = object.NewBuiltinProperty("RFC3339Nano", object.NewString(time.RFC3339Nano))
	timeProperties["Kitchen"] = object.NewBuiltinProperty("Kitchen", object.NewString(time.Kitchen))
	timeProperties["Stamp"] = object.NewBuiltinProperty("Stamp", object.NewString(time.Stamp))
	timeProperties["StampMilli"] = object.NewBuiltinProperty("StampMilli", object.NewString(time.StampMilli))
	timeProperties["StampMicro"] = object.NewBuiltinProperty("StampMicro", object.NewString(time.StampMicro))
	timeProperties["StampNano"] = object.NewBuiltinProperty("StampNano", object.NewString(time.StampNano))
}

func convertTimeFormat(format *object.String) string {
	timeFormat := format.Value
	for strfmtFormat, golangFormat := range timeFormatConversions {
		timeFormat = strings.ReplaceAll(timeFormat, strfmtFormat, golangFormat)
	}
	return timeFormat
}
