// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [DateIntervalFormatter] class.
var DateIntervalFormatterClass = _DateIntervalFormatterClass{objc.GetClass("NSDateIntervalFormatter")}

type _DateIntervalFormatterClass struct {
	objc.Class
}

// An interface definition for the [DateIntervalFormatter] class.
type IDateIntervalFormatter interface {
	IFormatter
	StringFromDateInterval(dateInterval IDateInterval) string
	StringFromDateToDate(fromDate IDate, toDate IDate) string
	Locale() Locale
	SetLocale(value ILocale)
	TimeZone() TimeZone
	SetTimeZone(value ITimeZone)
	TimeStyle() DateIntervalFormatterStyle
	SetTimeStyle(value DateIntervalFormatterStyle)
	DateStyle() DateIntervalFormatterStyle
	SetDateStyle(value DateIntervalFormatterStyle)
	Calendar() Calendar
	SetCalendar(value ICalendar)
	DateTemplate() string
	SetDateTemplate(value string)
}

// A formatter that creates string representations of time intervals. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdateintervalformatter?language=objc
type DateIntervalFormatter struct {
	Formatter
}

func DateIntervalFormatterFrom(ptr unsafe.Pointer) DateIntervalFormatter {
	return DateIntervalFormatter{
		Formatter: FormatterFrom(ptr),
	}
}

func (dc _DateIntervalFormatterClass) Alloc() DateIntervalFormatter {
	rv := objc.Call[DateIntervalFormatter](dc, objc.Sel("alloc"))
	return rv
}

func DateIntervalFormatter_Alloc() DateIntervalFormatter {
	return DateIntervalFormatterClass.Alloc()
}

func (dc _DateIntervalFormatterClass) New() DateIntervalFormatter {
	rv := objc.Call[DateIntervalFormatter](dc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewDateIntervalFormatter() DateIntervalFormatter {
	return DateIntervalFormatterClass.New()
}

func (d_ DateIntervalFormatter) Init() DateIntervalFormatter {
	rv := objc.Call[DateIntervalFormatter](d_, objc.Sel("init"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdateintervalformatter/1642848-stringfromdateinterval?language=objc
func (d_ DateIntervalFormatter) StringFromDateInterval(dateInterval IDateInterval) string {
	rv := objc.Call[string](d_, objc.Sel("stringFromDateInterval:"), objc.Ptr(dateInterval))
	return rv
}

// Returns a formatted string based on the specified start and end dates. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdateintervalformatter/1418368-stringfromdate?language=objc
func (d_ DateIntervalFormatter) StringFromDateToDate(fromDate IDate, toDate IDate) string {
	rv := objc.Call[string](d_, objc.Sel("stringFromDate:toDate:"), objc.Ptr(fromDate), objc.Ptr(toDate))
	return rv
}

// The locale to use when formatting date and time values. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdateintervalformatter/1409992-locale?language=objc
func (d_ DateIntervalFormatter) Locale() Locale {
	rv := objc.Call[Locale](d_, objc.Sel("locale"))
	return rv
}

// The locale to use when formatting date and time values. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdateintervalformatter/1409992-locale?language=objc
func (d_ DateIntervalFormatter) SetLocale(value ILocale) {
	objc.Call[objc.Void](d_, objc.Sel("setLocale:"), objc.Ptr(value))
}

// The time zone with which to specify time values. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdateintervalformatter/1410228-timezone?language=objc
func (d_ DateIntervalFormatter) TimeZone() TimeZone {
	rv := objc.Call[TimeZone](d_, objc.Sel("timeZone"))
	return rv
}

// The time zone with which to specify time values. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdateintervalformatter/1410228-timezone?language=objc
func (d_ DateIntervalFormatter) SetTimeZone(value ITimeZone) {
	objc.Call[objc.Void](d_, objc.Sel("setTimeZone:"), objc.Ptr(value))
}

// The style to use when formatting hour, minute, and second information. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdateintervalformatter/1415655-timestyle?language=objc
func (d_ DateIntervalFormatter) TimeStyle() DateIntervalFormatterStyle {
	rv := objc.Call[DateIntervalFormatterStyle](d_, objc.Sel("timeStyle"))
	return rv
}

// The style to use when formatting hour, minute, and second information. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdateintervalformatter/1415655-timestyle?language=objc
func (d_ DateIntervalFormatter) SetTimeStyle(value DateIntervalFormatterStyle) {
	objc.Call[objc.Void](d_, objc.Sel("setTimeStyle:"), value)
}

// The style to use when formatting day, month, and year information. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdateintervalformatter/1409519-datestyle?language=objc
func (d_ DateIntervalFormatter) DateStyle() DateIntervalFormatterStyle {
	rv := objc.Call[DateIntervalFormatterStyle](d_, objc.Sel("dateStyle"))
	return rv
}

// The style to use when formatting day, month, and year information. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdateintervalformatter/1409519-datestyle?language=objc
func (d_ DateIntervalFormatter) SetDateStyle(value DateIntervalFormatterStyle) {
	objc.Call[objc.Void](d_, objc.Sel("setDateStyle:"), value)
}

// The calendar to use for date values. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdateintervalformatter/1417984-calendar?language=objc
func (d_ DateIntervalFormatter) Calendar() Calendar {
	rv := objc.Call[Calendar](d_, objc.Sel("calendar"))
	return rv
}

// The calendar to use for date values. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdateintervalformatter/1417984-calendar?language=objc
func (d_ DateIntervalFormatter) SetCalendar(value ICalendar) {
	objc.Call[objc.Void](d_, objc.Sel("setCalendar:"), objc.Ptr(value))
}

// The template for formatting one date and time value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdateintervalformatter/1407373-datetemplate?language=objc
func (d_ DateIntervalFormatter) DateTemplate() string {
	rv := objc.Call[string](d_, objc.Sel("dateTemplate"))
	return rv
}

// The template for formatting one date and time value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdateintervalformatter/1407373-datetemplate?language=objc
func (d_ DateIntervalFormatter) SetDateTemplate(value string) {
	objc.Call[objc.Void](d_, objc.Sel("setDateTemplate:"), value)
}
