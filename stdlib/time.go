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
			Description: `Formats the given unix timestamp with the given layout.

[Go date and time formats](https://gosamples.dev/date-time-format-cheatsheet/) are natively supported.
You can also use some but not all [formats present in many other languages](https://apidock.com/ruby/Time/strftime) which are not fully supported.
Take a look at [the source](https://github.com/Flipez/rocket-lang/blob/main/stdlib/time.go) to see which formatters are supported.`,
			ArgPattern: object.Args(
				object.Arg(object.INTEGER_OBJ),
				object.Arg(object.STRING_OBJ),
			),
			ReturnPattern: object.Args(
				object.Arg(object.STRING_OBJ),
			),
			Example: `🚀 » Time.format(Time.unix(), "Mon Jan _2 15:04:05 2006")
» "Mon Oct 31 00:08:10 2022"
🚀 » Time.format(Time.unix(), "%a %b %e %H:%M:%S %Y")
» "Mon Oct 31 00:28:43 2022"`,
		},
		func(_ object.Environment, args ...object.Object) object.Object {
			unixTimestamp := args[0].(*object.Integer)
			timeFormat := args[1].(*object.String)

			time := time.Unix(unixTimestamp.Value, 0)

			return object.NewString(time.Format(convertTimeFormat(timeFormat)))
		})
	timeFunctions["parse"] = object.NewBuiltinFunction(
		"parse",
		object.MethodLayout{
			Description: `Parses a given string with the given format to a unix timestamp.

[Go date and time formats](https://gosamples.dev/date-time-format-cheatsheet/) are natively supported.
You can also use some but not all [formats present in many other languages](https://apidock.com/ruby/Time/strftime) which are not fully supported.
Take a look at [the source](https://github.com/Flipez/rocket-lang/blob/main/stdlib/time.go) to see which formatters are supported.`,
			ArgPattern: object.Args(
				object.Arg(object.STRING_OBJ),
				object.Arg(object.STRING_OBJ),
			),
			ReturnPattern: object.Args(
				object.Arg(object.STRING_OBJ),
			),
			Example: `🚀 » Time.parse("2022-03-23", "2006-01-02")
» 1647993600
🚀 » Time.parse("2022-03-23", "%Y-%m-%d")
» 1647993600`,
		},
		func(_ object.Environment, args ...object.Object) object.Object {
			timeString := args[0].(*object.String)
			timeFormat := args[1].(*object.String)

			timeParsed, err := time.Parse(convertTimeFormat(timeFormat), timeString.Value)
			if err != nil {
				return object.NewErrorFormat("Error while parsing time: %s", err)
			}

			return object.NewInteger(timeParsed.Unix())
		})
	timeFunctions["sleep"] = object.NewBuiltinFunction(
		"sleep",
		object.MethodLayout{
			Description: "Stops the RocketLang routine for at least the stated duration in seconds",
			ArgPattern: object.Args(
				object.Arg(object.INTEGER_OBJ),
			),
			ReturnPattern: object.Args(
				object.Arg(object.NIL_OBJ),
			),
			Example: `🚀 > Time.sleep(2)`,
		},
		func(_ object.Environment, args ...object.Object) object.Object {
			time.Sleep(time.Duration(args[0].(*object.Integer).Value) * time.Second)
			return object.NIL
		})
	timeFunctions["unix"] = object.NewBuiltinFunction(
		"unix",
		object.MethodLayout{
			Description: "Returns the current time as unix timestamp",
			ReturnPattern: object.Args(
				object.Arg(object.INTEGER_OBJ),
			),
			Example: `🚀 > Time.Unix()`,
		},
		func(_ object.Environment, args ...object.Object) object.Object {
			return object.NewInteger(time.Now().Unix())
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
