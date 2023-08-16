// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [DateComponentsFormatter] class.
var DateComponentsFormatterClass = _DateComponentsFormatterClass{objc.GetClass("NSDateComponentsFormatter")}

type _DateComponentsFormatterClass struct {
	objc.Class
}

// An interface definition for the [DateComponentsFormatter] class.
type IDateComponentsFormatter interface {
	IFormatter
	StringFromTimeInterval(ti TimeInterval) string
	StringFromDateComponents(components IDateComponents) string
	StringFromDateToDate(startDate IDate, endDate IDate) string
	UnitsStyle() DateComponentsFormatterUnitsStyle
	SetUnitsStyle(value DateComponentsFormatterUnitsStyle)
	IncludesTimeRemainingPhrase() bool
	SetIncludesTimeRemainingPhrase(value bool)
	AllowsFractionalUnits() bool
	SetAllowsFractionalUnits(value bool)
	MaximumUnitCount() int
	SetMaximumUnitCount(value int)
	CollapsesLargestUnit() bool
	SetCollapsesLargestUnit(value bool)
	ZeroFormattingBehavior() DateComponentsFormatterZeroFormattingBehavior
	SetZeroFormattingBehavior(value DateComponentsFormatterZeroFormattingBehavior)
	IncludesApproximationPhrase() bool
	SetIncludesApproximationPhrase(value bool)
	AllowedUnits() CalendarUnit
	SetAllowedUnits(value CalendarUnit)
	ReferenceDate() Date
	SetReferenceDate(value IDate)
	FormattingContext() FormattingContext
	SetFormattingContext(value FormattingContext)
	Calendar() Calendar
	SetCalendar(value ICalendar)
}

// A formatter that creates string representations of quantities of time. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdatecomponentsformatter?language=objc
type DateComponentsFormatter struct {
	Formatter
}

func DateComponentsFormatterFrom(ptr unsafe.Pointer) DateComponentsFormatter {
	return DateComponentsFormatter{
		Formatter: FormatterFrom(ptr),
	}
}

func (dc _DateComponentsFormatterClass) Alloc() DateComponentsFormatter {
	rv := objc.Call[DateComponentsFormatter](dc, objc.Sel("alloc"))
	return rv
}

func DateComponentsFormatter_Alloc() DateComponentsFormatter {
	return DateComponentsFormatterClass.Alloc()
}

func (dc _DateComponentsFormatterClass) New() DateComponentsFormatter {
	rv := objc.Call[DateComponentsFormatter](dc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewDateComponentsFormatter() DateComponentsFormatter {
	return DateComponentsFormatterClass.New()
}

func (d_ DateComponentsFormatter) Init() DateComponentsFormatter {
	rv := objc.Call[DateComponentsFormatter](d_, objc.Sel("init"))
	return rv
}

// Returns a localized string based on the specified date components and style option. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdatecomponentsformatter/1411422-localizedstringfromdatecomponent?language=objc
func (dc _DateComponentsFormatterClass) LocalizedStringFromDateComponentsUnitsStyle(components IDateComponents, unitsStyle DateComponentsFormatterUnitsStyle) string {
	rv := objc.Call[string](dc, objc.Sel("localizedStringFromDateComponents:unitsStyle:"), objc.Ptr(components), unitsStyle)
	return rv
}

// Returns a localized string based on the specified date components and style option. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdatecomponentsformatter/1411422-localizedstringfromdatecomponent?language=objc
func DateComponentsFormatter_LocalizedStringFromDateComponentsUnitsStyle(components IDateComponents, unitsStyle DateComponentsFormatterUnitsStyle) string {
	return DateComponentsFormatterClass.LocalizedStringFromDateComponentsUnitsStyle(components, unitsStyle)
}

// Returns a formatted string based on the specified number of seconds. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdatecomponentsformatter/1409040-stringfromtimeinterval?language=objc
func (d_ DateComponentsFormatter) StringFromTimeInterval(ti TimeInterval) string {
	rv := objc.Call[string](d_, objc.Sel("stringFromTimeInterval:"), ti)
	return rv
}

// Returns a formatted string based on the specified date component information. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdatecomponentsformatter/1407641-stringfromdatecomponents?language=objc
func (d_ DateComponentsFormatter) StringFromDateComponents(components IDateComponents) string {
	rv := objc.Call[string](d_, objc.Sel("stringFromDateComponents:"), objc.Ptr(components))
	return rv
}

// Returns a formatted string based on the time difference between two dates. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdatecomponentsformatter/1415967-stringfromdate?language=objc
func (d_ DateComponentsFormatter) StringFromDateToDate(startDate IDate, endDate IDate) string {
	rv := objc.Call[string](d_, objc.Sel("stringFromDate:toDate:"), objc.Ptr(startDate), objc.Ptr(endDate))
	return rv
}

// The formatting style for unit names. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdatecomponentsformatter/1413441-unitsstyle?language=objc
func (d_ DateComponentsFormatter) UnitsStyle() DateComponentsFormatterUnitsStyle {
	rv := objc.Call[DateComponentsFormatterUnitsStyle](d_, objc.Sel("unitsStyle"))
	return rv
}

// The formatting style for unit names. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdatecomponentsformatter/1413441-unitsstyle?language=objc
func (d_ DateComponentsFormatter) SetUnitsStyle(value DateComponentsFormatterUnitsStyle) {
	objc.Call[objc.Void](d_, objc.Sel("setUnitsStyle:"), value)
}

// A Boolean value indicating whether output strings reflect the amount of time remaining. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdatecomponentsformatter/1416416-includestimeremainingphrase?language=objc
func (d_ DateComponentsFormatter) IncludesTimeRemainingPhrase() bool {
	rv := objc.Call[bool](d_, objc.Sel("includesTimeRemainingPhrase"))
	return rv
}

// A Boolean value indicating whether output strings reflect the amount of time remaining. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdatecomponentsformatter/1416416-includestimeremainingphrase?language=objc
func (d_ DateComponentsFormatter) SetIncludesTimeRemainingPhrase(value bool) {
	objc.Call[objc.Void](d_, objc.Sel("setIncludesTimeRemainingPhrase:"), value)
}

// A Boolean indicating whether non-integer units may be used for values. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdatecomponentsformatter/1413084-allowsfractionalunits?language=objc
func (d_ DateComponentsFormatter) AllowsFractionalUnits() bool {
	rv := objc.Call[bool](d_, objc.Sel("allowsFractionalUnits"))
	return rv
}

// A Boolean indicating whether non-integer units may be used for values. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdatecomponentsformatter/1413084-allowsfractionalunits?language=objc
func (d_ DateComponentsFormatter) SetAllowsFractionalUnits(value bool) {
	objc.Call[objc.Void](d_, objc.Sel("setAllowsFractionalUnits:"), value)
}

// The maximum number of time units to include in the output string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdatecomponentsformatter/1416214-maximumunitcount?language=objc
func (d_ DateComponentsFormatter) MaximumUnitCount() int {
	rv := objc.Call[int](d_, objc.Sel("maximumUnitCount"))
	return rv
}

// The maximum number of time units to include in the output string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdatecomponentsformatter/1416214-maximumunitcount?language=objc
func (d_ DateComponentsFormatter) SetMaximumUnitCount(value int) {
	objc.Call[objc.Void](d_, objc.Sel("setMaximumUnitCount:"), value)
}

// A Boolean value indicating whether to collapse the largest unit into smaller units when a certain threshold is met. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdatecomponentsformatter/1410812-collapseslargestunit?language=objc
func (d_ DateComponentsFormatter) CollapsesLargestUnit() bool {
	rv := objc.Call[bool](d_, objc.Sel("collapsesLargestUnit"))
	return rv
}

// A Boolean value indicating whether to collapse the largest unit into smaller units when a certain threshold is met. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdatecomponentsformatter/1410812-collapseslargestunit?language=objc
func (d_ DateComponentsFormatter) SetCollapsesLargestUnit(value bool) {
	objc.Call[objc.Void](d_, objc.Sel("setCollapsesLargestUnit:"), value)
}

// The formatting style for units whose value is 0. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdatecomponentsformatter/1413749-zeroformattingbehavior?language=objc
func (d_ DateComponentsFormatter) ZeroFormattingBehavior() DateComponentsFormatterZeroFormattingBehavior {
	rv := objc.Call[DateComponentsFormatterZeroFormattingBehavior](d_, objc.Sel("zeroFormattingBehavior"))
	return rv
}

// The formatting style for units whose value is 0. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdatecomponentsformatter/1413749-zeroformattingbehavior?language=objc
func (d_ DateComponentsFormatter) SetZeroFormattingBehavior(value DateComponentsFormatterZeroFormattingBehavior) {
	objc.Call[objc.Void](d_, objc.Sel("setZeroFormattingBehavior:"), value)
}

// A Boolean value indicating whether the resulting phrase reflects an inexact time value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdatecomponentsformatter/1416387-includesapproximationphrase?language=objc
func (d_ DateComponentsFormatter) IncludesApproximationPhrase() bool {
	rv := objc.Call[bool](d_, objc.Sel("includesApproximationPhrase"))
	return rv
}

// A Boolean value indicating whether the resulting phrase reflects an inexact time value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdatecomponentsformatter/1416387-includesapproximationphrase?language=objc
func (d_ DateComponentsFormatter) SetIncludesApproximationPhrase(value bool) {
	objc.Call[objc.Void](d_, objc.Sel("setIncludesApproximationPhrase:"), value)
}

// The bitmask of calendrical units such as day and month to include in the output string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdatecomponentsformatter/1410216-allowedunits?language=objc
func (d_ DateComponentsFormatter) AllowedUnits() CalendarUnit {
	rv := objc.Call[CalendarUnit](d_, objc.Sel("allowedUnits"))
	return rv
}

// The bitmask of calendrical units such as day and month to include in the output string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdatecomponentsformatter/1410216-allowedunits?language=objc
func (d_ DateComponentsFormatter) SetAllowedUnits(value CalendarUnit) {
	objc.Call[objc.Void](d_, objc.Sel("setAllowedUnits:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdatecomponentsformatter/2878110-referencedate?language=objc
func (d_ DateComponentsFormatter) ReferenceDate() Date {
	rv := objc.Call[Date](d_, objc.Sel("referenceDate"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdatecomponentsformatter/2878110-referencedate?language=objc
func (d_ DateComponentsFormatter) SetReferenceDate(value IDate) {
	objc.Call[objc.Void](d_, objc.Sel("setReferenceDate:"), objc.Ptr(value))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdatecomponentsformatter/1414198-formattingcontext?language=objc
func (d_ DateComponentsFormatter) FormattingContext() FormattingContext {
	rv := objc.Call[FormattingContext](d_, objc.Sel("formattingContext"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdatecomponentsformatter/1414198-formattingcontext?language=objc
func (d_ DateComponentsFormatter) SetFormattingContext(value FormattingContext) {
	objc.Call[objc.Void](d_, objc.Sel("setFormattingContext:"), value)
}

// The default calendar to use when formatting date components. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdatecomponentsformatter/1407359-calendar?language=objc
func (d_ DateComponentsFormatter) Calendar() Calendar {
	rv := objc.Call[Calendar](d_, objc.Sel("calendar"))
	return rv
}

// The default calendar to use when formatting date components. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdatecomponentsformatter/1407359-calendar?language=objc
func (d_ DateComponentsFormatter) SetCalendar(value ICalendar) {
	objc.Call[objc.Void](d_, objc.Sel("setCalendar:"), objc.Ptr(value))
}
