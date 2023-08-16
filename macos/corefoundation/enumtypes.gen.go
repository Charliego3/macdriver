// AUTO-GENERATED CODE, DO NOT MODIFY

package corefoundation

// Type used to represent a specific point in time relative to the absolute reference date of 1 Jan 2001 00:00:00 GMT. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/cfabsolutetime?language=objc
type AbsoluteTime TimeInterval

// A binary value of either 0 or 1. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/cfbit?language=objc
type Bit uint32

// Type that identifies a distinct reference number for a resource map. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/cfbundlerefnum?language=objc
type BundleRefNum int

// Flags that identify byte order. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/cfbyteorder?language=objc
type ByteOrder Index

// CFCalendarUnit constants are used to specify calendrical units, such as day or month, in various calendar calculations. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/cfcalendarunit?language=objc
type CalendarUnit OptionFlags

const (
	KCalendarUnitDay               CalendarUnit = 16
	KCalendarUnitEra               CalendarUnit = 2
	KCalendarUnitHour              CalendarUnit = 32
	KCalendarUnitMinute            CalendarUnit = 64
	KCalendarUnitMonth             CalendarUnit = 8
	KCalendarUnitQuarter           CalendarUnit = 2048
	KCalendarUnitSecond            CalendarUnit = 128
	KCalendarUnitWeek              CalendarUnit = 256
	KCalendarUnitWeekOfMonth       CalendarUnit = 4096
	KCalendarUnitWeekOfYear        CalendarUnit = 8192
	KCalendarUnitWeekday           CalendarUnit = 512
	KCalendarUnitWeekdayOrdinal    CalendarUnit = 1024
	KCalendarUnitYear              CalendarUnit = 4
	KCalendarUnitYearForWeekOfYear CalendarUnit = 16384
)

// Defines a predefined character set. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/cfcharactersetpredefinedset?language=objc
type CharacterSetPredefinedSet Index

const (
	KCharacterSetAlphaNumeric         CharacterSetPredefinedSet = 10
	KCharacterSetCapitalizedLetter    CharacterSetPredefinedSet = 13
	KCharacterSetControl              CharacterSetPredefinedSet = 1
	KCharacterSetDecimalDigit         CharacterSetPredefinedSet = 4
	KCharacterSetDecomposable         CharacterSetPredefinedSet = 9
	KCharacterSetIllegal              CharacterSetPredefinedSet = 12
	KCharacterSetLetter               CharacterSetPredefinedSet = 5
	KCharacterSetLowercaseLetter      CharacterSetPredefinedSet = 6
	KCharacterSetNewline              CharacterSetPredefinedSet = 15
	KCharacterSetNonBase              CharacterSetPredefinedSet = 8
	KCharacterSetPunctuation          CharacterSetPredefinedSet = 11
	KCharacterSetSymbol               CharacterSetPredefinedSet = 14
	KCharacterSetUppercaseLetter      CharacterSetPredefinedSet = 7
	KCharacterSetWhitespace           CharacterSetPredefinedSet = 2
	KCharacterSetWhitespaceAndNewline CharacterSetPredefinedSet = 3
)

// Constants returned by comparison functions, indicating whether a value is equal to, less than, or greater than another value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/cfcomparisonresult?language=objc
type ComparisonResult Index

const (
	KCompareEqualTo     ComparisonResult = 0
	KCompareGreaterThan ComparisonResult = 1
	KCompareLessThan    ComparisonResult = -1
)

// A CFOptionFlags type for specifying options for searching. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/cfdatasearchflags?language=objc
type DataSearchFlags OptionFlags

const (
	KDataSearchAnchored  DataSearchFlags = 2
	KDataSearchBackwards DataSearchFlags = 1
)

// Data type for predefined date and time format styles. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/cfdateformatterstyle?language=objc
type DateFormatterStyle Index

const (
	KDateFormatterFullStyle   DateFormatterStyle = 4
	KDateFormatterLongStyle   DateFormatterStyle = 3
	KDateFormatterMediumStyle DateFormatterStyle = 2
	KDateFormatterNoStyle     DateFormatterStyle = 0
	KDateFormatterShortStyle  DateFormatterStyle = 1
)

// Defines a type for the native file descriptor. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/cffiledescriptornativedescriptor?language=objc
type FileDescriptorNativeDescriptor int

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/cffilesecurityclearoptions?language=objc
type FileSecurityClearOptions OptionFlags

const (
	KFileSecurityClearAccessControlList FileSecurityClearOptions = 32
	KFileSecurityClearGroup             FileSecurityClearOptions = 2
	KFileSecurityClearGroupUUID         FileSecurityClearOptions = 16
	KFileSecurityClearMode              FileSecurityClearOptions = 4
	KFileSecurityClearOwner             FileSecurityClearOptions = 1
	KFileSecurityClearOwnerUUID         FileSecurityClearOptions = 8
)

// These option flags are used as a mask to indicate a specific set of fields in the CFGregorianDate or CFGregorianUnits structures. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/cfgregorianunitflags?language=objc
type GregorianUnitFlags OptionFlags

const (
	KGregorianAllUnits     GregorianUnitFlags = 16777215
	KGregorianUnitsDays    GregorianUnitFlags = 4
	KGregorianUnitsHours   GregorianUnitFlags = 8
	KGregorianUnitsMinutes GregorianUnitFlags = 16
	KGregorianUnitsMonths  GregorianUnitFlags = 2
	KGregorianUnitsSeconds GregorianUnitFlags = 32
	KGregorianUnitsYears   GregorianUnitFlags = 1
)

// A type for hash codes returned by the CFHash function. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/cfhashcode?language=objc
type HashCode int32

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/cfiso8601dateformatoptions?language=objc
type ISO8601DateFormatOptions OptionFlags

const (
	KISO8601DateFormatWithColonSeparatorInTime     ISO8601DateFormatOptions = 512
	KISO8601DateFormatWithColonSeparatorInTimeZone ISO8601DateFormatOptions = 1024
	KISO8601DateFormatWithDashSeparatorInDate      ISO8601DateFormatOptions = 256
	KISO8601DateFormatWithDay                      ISO8601DateFormatOptions = 16
	KISO8601DateFormatWithFractionalSeconds        ISO8601DateFormatOptions = 2048
	KISO8601DateFormatWithFullDate                 ISO8601DateFormatOptions = 275
	KISO8601DateFormatWithFullTime                 ISO8601DateFormatOptions = 1632
	KISO8601DateFormatWithInternetDateTime         ISO8601DateFormatOptions = 1907
	KISO8601DateFormatWithMonth                    ISO8601DateFormatOptions = 2
	KISO8601DateFormatWithSpaceBetweenDateAndTime  ISO8601DateFormatOptions = 128
	KISO8601DateFormatWithTime                     ISO8601DateFormatOptions = 32
	KISO8601DateFormatWithTimeZone                 ISO8601DateFormatOptions = 64
	KISO8601DateFormatWithWeekOfYear               ISO8601DateFormatOptions = 4
	KISO8601DateFormatWithYear                     ISO8601DateFormatOptions = 1
)

// Priority values used for kAXPriorityKey [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/cfindex?language=objc
type Index int32

const (
	KNotFound Index = -1
)

// These constants describe the text direction for a language. They are returned by the functions CFLocaleGetLanguageCharacterDirection and CFLocaleGetLanguageLineDirection. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/cflocalelanguagedirection?language=objc
type LocaleLanguageDirection Index

const (
	KLocaleLanguageDirectionBottomToTop LocaleLanguageDirection = 4
	KLocaleLanguageDirectionLeftToRight LocaleLanguageDirection = 1
	KLocaleLanguageDirectionRightToLeft LocaleLanguageDirection = 2
	KLocaleLanguageDirectionTopToBottom LocaleLanguageDirection = 3
	KLocaleLanguageDirectionUnknown     LocaleLanguageDirection = 0
)

// Suspension flags that indicate how distributed notifications should be handled when the receiving application is in the background. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/cfnotificationsuspensionbehavior?language=objc
type NotificationSuspensionBehavior Index

const (
	NotificationSuspensionBehaviorCoalesce           NotificationSuspensionBehavior = 2
	NotificationSuspensionBehaviorDeliverImmediately NotificationSuspensionBehavior = 4
	NotificationSuspensionBehaviorDrop               NotificationSuspensionBehavior = 1
	NotificationSuspensionBehaviorHold               NotificationSuspensionBehavior = 3
)

// Type for constants specifying how numbers should be parsed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/cfnumberformatteroptionflags?language=objc
type NumberFormatterOptionFlags OptionFlags

const (
	KNumberFormatterParseIntegersOnly NumberFormatterOptionFlags = 1
)

// Type for constants specifying how numbers should be padded. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/cfnumberformatterpadposition?language=objc
type NumberFormatterPadPosition Index

const (
	KNumberFormatterPadAfterPrefix  NumberFormatterPadPosition = 1
	KNumberFormatterPadAfterSuffix  NumberFormatterPadPosition = 3
	KNumberFormatterPadBeforePrefix NumberFormatterPadPosition = 0
	KNumberFormatterPadBeforeSuffix NumberFormatterPadPosition = 2
)

// These constants are used to specify how numbers should be rounded. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/cfnumberformatterroundingmode?language=objc
type NumberFormatterRoundingMode Index

const (
	KNumberFormatterRoundCeiling  NumberFormatterRoundingMode = 0
	KNumberFormatterRoundDown     NumberFormatterRoundingMode = 2
	KNumberFormatterRoundFloor    NumberFormatterRoundingMode = 1
	KNumberFormatterRoundHalfDown NumberFormatterRoundingMode = 5
	KNumberFormatterRoundHalfEven NumberFormatterRoundingMode = 4
	KNumberFormatterRoundHalfUp   NumberFormatterRoundingMode = 6
	KNumberFormatterRoundUp       NumberFormatterRoundingMode = 3
)

// Type for constants specifying a formatter style. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/cfnumberformatterstyle?language=objc
type NumberFormatterStyle Index

const (
	KNumberFormatterCurrencyAccountingStyle NumberFormatterStyle = 10
	KNumberFormatterCurrencyISOCodeStyle    NumberFormatterStyle = 8
	KNumberFormatterCurrencyPluralStyle     NumberFormatterStyle = 9
	KNumberFormatterCurrencyStyle           NumberFormatterStyle = 2
	KNumberFormatterDecimalStyle            NumberFormatterStyle = 1
	KNumberFormatterNoStyle                 NumberFormatterStyle = 0
	KNumberFormatterOrdinalStyle            NumberFormatterStyle = 6
	KNumberFormatterPercentStyle            NumberFormatterStyle = 3
	KNumberFormatterScientificStyle         NumberFormatterStyle = 4
	KNumberFormatterSpellOutStyle           NumberFormatterStyle = 5
)

// Flags used by CFNumber to indicate the data type of a value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/cfnumbertype?language=objc
type NumberType Index

const (
	KNumberCFIndexType   NumberType = 14
	KNumberCGFloatType   NumberType = 16
	KNumberCharType      NumberType = 7
	KNumberDoubleType    NumberType = 13
	KNumberFloat32Type   NumberType = 5
	KNumberFloat64Type   NumberType = 6
	KNumberFloatType     NumberType = 12
	KNumberIntType       NumberType = 9
	KNumberLongLongType  NumberType = 11
	KNumberLongType      NumberType = 10
	KNumberMaxType       NumberType = 16
	KNumberNSIntegerType NumberType = 15
	KNumberSInt16Type    NumberType = 2
	KNumberSInt32Type    NumberType = 3
	KNumberSInt64Type    NumberType = 4
	KNumberSInt8Type     NumberType = 1
	KNumberShortType     NumberType = 8
)

// A bitfield used for passing special allocation and other requests into Core Foundation functions. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/cfoptionflags?language=objc
type OptionFlags int32

// Specifies the format of a property list. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/cfpropertylistformat?language=objc
type PropertyListFormat Index

const (
	KPropertyListBinaryFormat_v1_0 PropertyListFormat = 200
	KPropertyListOpenStepFormat    PropertyListFormat = 1
	KPropertyListXMLFormat_v1_0    PropertyListFormat = 100
)

// Type for flags that determine the degree of mutability of newly created property lists. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/cfpropertylistmutabilityoptions?language=objc
type PropertyListMutabilityOptions OptionFlags

const (
	KPropertyListImmutable                  PropertyListMutabilityOptions = 0
	KPropertyListMutableContainers          PropertyListMutabilityOptions = 1
	KPropertyListMutableContainersAndLeaves PropertyListMutabilityOptions = 2
)

// Run loop activity stages in which run loop observers can be scheduled. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/cfrunloopactivity?language=objc
type RunLoopActivity OptionFlags

const (
	KRunLoopAfterWaiting  RunLoopActivity = 64
	KRunLoopAllActivities RunLoopActivity = 268435455
	KRunLoopBeforeSources RunLoopActivity = 4
	KRunLoopBeforeTimers  RunLoopActivity = 2
	KRunLoopBeforeWaiting RunLoopActivity = 32
	KRunLoopEntry         RunLoopActivity = 1
	KRunLoopExit          RunLoopActivity = 128
)

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/cfrunlooprunresult?language=objc
type RunLoopRunResult int32

const (
	KRunLoopRunFinished      RunLoopRunResult = 1
	KRunLoopRunHandledSource RunLoopRunResult = 4
	KRunLoopRunStopped       RunLoopRunResult = 2
	KRunLoopRunTimedOut      RunLoopRunResult = 3
)

// Types of socket activity that can cause the callback function of a CFSocket object to be called. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/cfsocketcallbacktype?language=objc
type SocketCallBackType OptionFlags

const (
	KSocketAcceptCallBack  SocketCallBackType = 2
	KSocketConnectCallBack SocketCallBackType = 4
	KSocketDataCallBack    SocketCallBackType = 3
	KSocketNoCallBack      SocketCallBackType = 0
	KSocketReadCallBack    SocketCallBackType = 1
	KSocketWriteCallBack   SocketCallBackType = 8
)

// Error codes for many CFSocket functions. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/cfsocketerror?language=objc
type SocketError Index

const (
	KSocketError   SocketError = -1
	KSocketSuccess SocketError = 0
	KSocketTimeout SocketError = -2
)

// Type for the platform-specific native socket handle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/cfsocketnativehandle?language=objc
type SocketNativeHandle int

// Defines constants for values returned in the domain field of the CFStreamError structure. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/cfstreamerrordomain?language=objc
type StreamErrorDomain Index

const (
	KStreamErrorDomainCustom      StreamErrorDomain = -1
	KStreamErrorDomainMacOSStatus StreamErrorDomain = 2
	KStreamErrorDomainPOSIX       StreamErrorDomain = 1
)

// Defines constants for stream-related events. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/cfstreameventtype?language=objc
type StreamEventType OptionFlags

const (
	KStreamEventCanAcceptBytes    StreamEventType = 4
	KStreamEventEndEncountered    StreamEventType = 16
	KStreamEventErrorOccurred     StreamEventType = 8
	KStreamEventHasBytesAvailable StreamEventType = 2
	KStreamEventNone              StreamEventType = 0
	KStreamEventOpenCompleted     StreamEventType = 1
)

// Constants that describe the status of a stream. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/cfstreamstatus?language=objc
type StreamStatus Index

const (
	KStreamStatusAtEnd   StreamStatus = 5
	KStreamStatusClosed  StreamStatus = 6
	KStreamStatusError   StreamStatus = 7
	KStreamStatusNotOpen StreamStatus = 0
	KStreamStatusOpen    StreamStatus = 2
	KStreamStatusOpening StreamStatus = 1
	KStreamStatusReading StreamStatus = 3
	KStreamStatusWriting StreamStatus = 4
)

// Encodings that are built-in on all platforms on which macOS runs. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/cfstringbuiltinencodings?language=objc
type StringBuiltInEncodings StringEncoding

const (
	KStringEncodingASCII         StringBuiltInEncodings = 1536
	KStringEncodingISOLatin1     StringBuiltInEncodings = 513
	KStringEncodingMacRoman      StringBuiltInEncodings = 0
	KStringEncodingNextStepLatin StringBuiltInEncodings = 2817
	KStringEncodingNonLossyASCII StringBuiltInEncodings = 3071
	KStringEncodingUTF16         StringBuiltInEncodings = 256
	KStringEncodingUTF16BE       StringBuiltInEncodings = 268435712
	KStringEncodingUTF16LE       StringBuiltInEncodings = 335544576
	KStringEncodingUTF32         StringBuiltInEncodings = 201326848
	KStringEncodingUTF32BE       StringBuiltInEncodings = 402653440
	KStringEncodingUTF32LE       StringBuiltInEncodings = 469762304
	KStringEncodingUTF8          StringBuiltInEncodings = 134217984
	KStringEncodingUnicode       StringBuiltInEncodings = 256
	KStringEncodingWindowsLatin1 StringBuiltInEncodings = 1280
)

// A CFOptionFlags type for specifying options for string comparison . [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/cfstringcompareflags?language=objc
type StringCompareFlags OptionFlags

const (
	KCompareAnchored             StringCompareFlags = 8
	KCompareBackwards            StringCompareFlags = 4
	KCompareCaseInsensitive      StringCompareFlags = 1
	KCompareDiacriticInsensitive StringCompareFlags = 128
	KCompareForcedOrdering       StringCompareFlags = 512
	KCompareLocalized            StringCompareFlags = 32
	KCompareNonliteral           StringCompareFlags = 16
	KCompareNumerically          StringCompareFlags = 64
	KCompareWidthInsensitive     StringCompareFlags = 256
)

// An integer type for constants used to specify supported string encodings in various CFString functions. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/cfstringencoding?language=objc
type StringEncoding uint32

// Index type for constants used to specify external string encodings. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/cfstringencodings?language=objc
type StringEncodings Index

const (
	KStringEncodingANSEL                   StringEncodings = 1537
	KStringEncodingBig5                    StringEncodings = 2563
	KStringEncodingBig5_E                  StringEncodings = 2569
	KStringEncodingBig5_HKSCS_1999         StringEncodings = 2566
	KStringEncodingCNS_11643_92_P1         StringEncodings = 1617
	KStringEncodingCNS_11643_92_P2         StringEncodings = 1618
	KStringEncodingCNS_11643_92_P3         StringEncodings = 1619
	KStringEncodingDOSArabic               StringEncodings = 1049
	KStringEncodingDOSBalticRim            StringEncodings = 1030
	KStringEncodingDOSCanadianFrench       StringEncodings = 1048
	KStringEncodingDOSChineseSimplif       StringEncodings = 1057
	KStringEncodingDOSChineseTrad          StringEncodings = 1059
	KStringEncodingDOSCyrillic             StringEncodings = 1043
	KStringEncodingDOSGreek                StringEncodings = 1029
	KStringEncodingDOSGreek1               StringEncodings = 1041
	KStringEncodingDOSGreek2               StringEncodings = 1052
	KStringEncodingDOSHebrew               StringEncodings = 1047
	KStringEncodingDOSIcelandic            StringEncodings = 1046
	KStringEncodingDOSJapanese             StringEncodings = 1056
	KStringEncodingDOSKorean               StringEncodings = 1058
	KStringEncodingDOSLatin1               StringEncodings = 1040
	KStringEncodingDOSLatin2               StringEncodings = 1042
	KStringEncodingDOSLatinUS              StringEncodings = 1024
	KStringEncodingDOSNordic               StringEncodings = 1050
	KStringEncodingDOSPortuguese           StringEncodings = 1045
	KStringEncodingDOSRussian              StringEncodings = 1051
	KStringEncodingDOSThai                 StringEncodings = 1053
	KStringEncodingDOSTurkish              StringEncodings = 1044
	KStringEncodingEBCDIC_CP037            StringEncodings = 3074
	KStringEncodingEBCDIC_US               StringEncodings = 3073
	KStringEncodingEUC_CN                  StringEncodings = 2352
	KStringEncodingEUC_JP                  StringEncodings = 2336
	KStringEncodingEUC_KR                  StringEncodings = 2368
	KStringEncodingEUC_TW                  StringEncodings = 2353
	KStringEncodingGBK_95                  StringEncodings = 1585
	KStringEncodingGB_18030_2000           StringEncodings = 1586
	KStringEncodingGB_2312_80              StringEncodings = 1584
	KStringEncodingHZ_GB_2312              StringEncodings = 2565
	KStringEncodingISOLatin10              StringEncodings = 528
	KStringEncodingISOLatin2               StringEncodings = 514
	KStringEncodingISOLatin3               StringEncodings = 515
	KStringEncodingISOLatin4               StringEncodings = 516
	KStringEncodingISOLatin5               StringEncodings = 521
	KStringEncodingISOLatin6               StringEncodings = 522
	KStringEncodingISOLatin7               StringEncodings = 525
	KStringEncodingISOLatin8               StringEncodings = 526
	KStringEncodingISOLatin9               StringEncodings = 527
	KStringEncodingISOLatinArabic          StringEncodings = 518
	KStringEncodingISOLatinCyrillic        StringEncodings = 517
	KStringEncodingISOLatinGreek           StringEncodings = 519
	KStringEncodingISOLatinHebrew          StringEncodings = 520
	KStringEncodingISOLatinThai            StringEncodings = 523
	KStringEncodingISO_2022_CN             StringEncodings = 2096
	KStringEncodingISO_2022_CN_EXT         StringEncodings = 2097
	KStringEncodingISO_2022_JP             StringEncodings = 2080
	KStringEncodingISO_2022_JP_1           StringEncodings = 2082
	KStringEncodingISO_2022_JP_2           StringEncodings = 2081
	KStringEncodingISO_2022_JP_3           StringEncodings = 2083
	KStringEncodingISO_2022_KR             StringEncodings = 2112
	KStringEncodingJIS_C6226_78            StringEncodings = 1572
	KStringEncodingJIS_X0201_76            StringEncodings = 1568
	KStringEncodingJIS_X0208_83            StringEncodings = 1569
	KStringEncodingJIS_X0208_90            StringEncodings = 1570
	KStringEncodingJIS_X0212_90            StringEncodings = 1571
	KStringEncodingKOI8_R                  StringEncodings = 2562
	KStringEncodingKOI8_U                  StringEncodings = 2568
	KStringEncodingKSC_5601_87             StringEncodings = 1600
	KStringEncodingKSC_5601_92_Johab       StringEncodings = 1601
	KStringEncodingMacArabic               StringEncodings = 4
	KStringEncodingMacArmenian             StringEncodings = 24
	KStringEncodingMacBengali              StringEncodings = 13
	KStringEncodingMacBurmese              StringEncodings = 19
	KStringEncodingMacCeltic               StringEncodings = 39
	KStringEncodingMacCentralEurRoman      StringEncodings = 29
	KStringEncodingMacChineseSimp          StringEncodings = 25
	KStringEncodingMacChineseTrad          StringEncodings = 2
	KStringEncodingMacCroatian             StringEncodings = 36
	KStringEncodingMacCyrillic             StringEncodings = 7
	KStringEncodingMacDevanagari           StringEncodings = 9
	KStringEncodingMacDingbats             StringEncodings = 34
	KStringEncodingMacEthiopic             StringEncodings = 28
	KStringEncodingMacExtArabic            StringEncodings = 31
	KStringEncodingMacFarsi                StringEncodings = 140
	KStringEncodingMacGaelic               StringEncodings = 40
	KStringEncodingMacGeorgian             StringEncodings = 23
	KStringEncodingMacGreek                StringEncodings = 6
	KStringEncodingMacGujarati             StringEncodings = 11
	KStringEncodingMacGurmukhi             StringEncodings = 10
	KStringEncodingMacHFS                  StringEncodings = 255
	KStringEncodingMacHebrew               StringEncodings = 5
	KStringEncodingMacIcelandic            StringEncodings = 37
	KStringEncodingMacInuit                StringEncodings = 236
	KStringEncodingMacJapanese             StringEncodings = 1
	KStringEncodingMacKannada              StringEncodings = 16
	KStringEncodingMacKhmer                StringEncodings = 20
	KStringEncodingMacKorean               StringEncodings = 3
	KStringEncodingMacLaotian              StringEncodings = 22
	KStringEncodingMacMalayalam            StringEncodings = 17
	KStringEncodingMacMongolian            StringEncodings = 27
	KStringEncodingMacOriya                StringEncodings = 12
	KStringEncodingMacRomanLatin1          StringEncodings = 2564
	KStringEncodingMacRomanian             StringEncodings = 38
	KStringEncodingMacSinhalese            StringEncodings = 18
	KStringEncodingMacSymbol               StringEncodings = 33
	KStringEncodingMacTamil                StringEncodings = 14
	KStringEncodingMacTelugu               StringEncodings = 15
	KStringEncodingMacThai                 StringEncodings = 21
	KStringEncodingMacTibetan              StringEncodings = 26
	KStringEncodingMacTurkish              StringEncodings = 35
	KStringEncodingMacUkrainian            StringEncodings = 152
	KStringEncodingMacVT100                StringEncodings = 252
	KStringEncodingMacVietnamese           StringEncodings = 30
	KStringEncodingNextStepJapanese        StringEncodings = 2818
	KStringEncodingShiftJIS                StringEncodings = 2561
	KStringEncodingShiftJIS_X0213          StringEncodings = 1576
	KStringEncodingShiftJIS_X0213_00       StringEncodings = 1576
	KStringEncodingShiftJIS_X0213_MenKuTen StringEncodings = 1577
	KStringEncodingUTF7                    StringEncodings = 67109120
	KStringEncodingUTF7_IMAP               StringEncodings = 2576
	KStringEncodingVISCII                  StringEncodings = 2567
	KStringEncodingWindowsArabic           StringEncodings = 1286
	KStringEncodingWindowsBalticRim        StringEncodings = 1287
	KStringEncodingWindowsCyrillic         StringEncodings = 1282
	KStringEncodingWindowsGreek            StringEncodings = 1283
	KStringEncodingWindowsHebrew           StringEncodings = 1285
	KStringEncodingWindowsKoreanJohab      StringEncodings = 1296
	KStringEncodingWindowsLatin2           StringEncodings = 1281
	KStringEncodingWindowsLatin5           StringEncodings = 1284
	KStringEncodingWindowsVietnamese       StringEncodings = 1288
)

// Unicode normalization forms as described in Unicode Technical Report #15. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/cfstringnormalizationform?language=objc
type StringNormalizationForm Index

const (
	KStringNormalizationFormC  StringNormalizationForm = 2
	KStringNormalizationFormD  StringNormalizationForm = 0
	KStringNormalizationFormKC StringNormalizationForm = 3
	KStringNormalizationFormKD StringNormalizationForm = 1
)

// Token types returned by CFStringTokenizerGoToTokenAtIndex and CFStringTokenizerAdvanceToNextToken. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/cfstringtokenizertokentype?language=objc
type StringTokenizerTokenType OptionFlags

const (
	KStringTokenizerTokenHasDerivedSubTokensMask StringTokenizerTokenType = 4
	KStringTokenizerTokenHasHasNumbersMask       StringTokenizerTokenType = 8
	KStringTokenizerTokenHasNonLettersMask       StringTokenizerTokenType = 16
	KStringTokenizerTokenHasSubTokensMask        StringTokenizerTokenType = 2
	KStringTokenizerTokenIsCJWordMask            StringTokenizerTokenType = 32
	KStringTokenizerTokenNone                    StringTokenizerTokenType = 0
	KStringTokenizerTokenNormal                  StringTokenizerTokenType = 1
)

// Type used to represent elapsed time in seconds. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/cftimeinterval?language=objc
type TimeInterval float64

const (
	KAbsoluteTimeIntervalSince1904 TimeInterval = 3061152000.000000
	KAbsoluteTimeIntervalSince1970 TimeInterval = 978307200.000000
)

// Index type for constants used to specify styles of time zone names. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/cftimezonenamestyle?language=objc
type TimeZoneNameStyle Index

const (
	KTimeZoneNameStyleDaylightSaving      TimeZoneNameStyle = 2
	KTimeZoneNameStyleGeneric             TimeZoneNameStyle = 4
	KTimeZoneNameStyleShortDaylightSaving TimeZoneNameStyle = 3
	KTimeZoneNameStyleShortGeneric        TimeZoneNameStyle = 5
	KTimeZoneNameStyleShortStandard       TimeZoneNameStyle = 1
	KTimeZoneNameStyleStandard            TimeZoneNameStyle = 0
)

// A type for unique, constant integer values that identify particular Core Foundation opaque types. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/cftypeid?language=objc
type TypeID int32

// Type for bookmark data creation options. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/cfurlbookmarkcreationoptions?language=objc
type URLBookmarkCreationOptions OptionFlags

const (
	KURLBookmarkCreationMinimalBookmarkMask              URLBookmarkCreationOptions = 512
	KURLBookmarkCreationPreferFileIDResolutionMask       URLBookmarkCreationOptions = 256
	KURLBookmarkCreationSecurityScopeAllowOnlyReadAccess URLBookmarkCreationOptions = 4096
	KURLBookmarkCreationSuitableForBookmarkFile          URLBookmarkCreationOptions = 1024
	KURLBookmarkCreationWithSecurityScope                URLBookmarkCreationOptions = 2048
	KURLBookmarkCreationWithoutImplicitSecurityScope     URLBookmarkCreationOptions = 536870912
)

// Type for bookmark file creation options. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/cfurlbookmarkfilecreationoptions?language=objc
type URLBookmarkFileCreationOptions OptionFlags

// Type for bookmark data resolution options. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/cfurlbookmarkresolutionoptions?language=objc
type URLBookmarkResolutionOptions OptionFlags

const (
	KBookmarkResolutionWithoutMountingMask              URLBookmarkResolutionOptions = 512
	KBookmarkResolutionWithoutUIMask                    URLBookmarkResolutionOptions = 256
	KURLBookmarkResolutionWithSecurityScope             URLBookmarkResolutionOptions = 1024
	KURLBookmarkResolutionWithoutImplicitStartAccessing URLBookmarkResolutionOptions = 32768
	KURLBookmarkResolutionWithoutMountingMask           URLBookmarkResolutionOptions = 512
	KURLBookmarkResolutionWithoutUIMask                 URLBookmarkResolutionOptions = 256
)

// The types of components in a URL. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/cfurlcomponenttype?language=objc
type URLComponentType Index

const (
	KURLComponentFragment          URLComponentType = 12
	KURLComponentHost              URLComponentType = 8
	KURLComponentNetLocation       URLComponentType = 2
	KURLComponentParameterString   URLComponentType = 10
	KURLComponentPassword          URLComponentType = 6
	KURLComponentPath              URLComponentType = 3
	KURLComponentPort              URLComponentType = 9
	KURLComponentQuery             URLComponentType = 11
	KURLComponentResourceSpecifier URLComponentType = 4
	KURLComponentScheme            URLComponentType = 1
	KURLComponentUser              URLComponentType = 5
	KURLComponentUserInfo          URLComponentType = 7
)

// Options for controlling enumerator behavior. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/cfurlenumeratoroptions?language=objc
type URLEnumeratorOptions OptionFlags

const (
	KURLEnumeratorDefaultBehavior             URLEnumeratorOptions = 0
	KURLEnumeratorDescendRecursively          URLEnumeratorOptions = 1
	KURLEnumeratorGenerateFileReferenceURLs   URLEnumeratorOptions = 4
	KURLEnumeratorGenerateRelativePathURLs    URLEnumeratorOptions = 64
	KURLEnumeratorIncludeDirectoriesPostOrder URLEnumeratorOptions = 32
	KURLEnumeratorIncludeDirectoriesPreOrder  URLEnumeratorOptions = 16
	KURLEnumeratorSkipInvisibles              URLEnumeratorOptions = 2
	KURLEnumeratorSkipPackageContents         URLEnumeratorOptions = 8
)

// Result codes from the CFURLEnumeratorGetNextURL function. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/cfurlenumeratorresult?language=objc
type URLEnumeratorResult Index

const (
	KURLEnumeratorDirectoryPostOrderSuccess URLEnumeratorResult = 4
	KURLEnumeratorEnd                       URLEnumeratorResult = 2
	KURLEnumeratorError                     URLEnumeratorResult = 3
	KURLEnumeratorSuccess                   URLEnumeratorResult = 1
)

// CFURL error codes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/cfurlerror?language=objc
type URLError Index

const (
	KURLImproperArgumentsError       URLError = -15
	KURLPropertyKeyUnavailableError  URLError = -17
	KURLRemoteHostUnavailableError   URLError = -14
	KURLResourceAccessViolationError URLError = -13
	KURLResourceNotFoundError        URLError = -12
	KURLTimeoutError                 URLError = -18
	KURLUnknownError                 URLError = -10
	KURLUnknownPropertyKeyError      URLError = -16
	KURLUnknownSchemeError           URLError = -11
)

// Options you can use to determine how CFURL functions parse a file system path name. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/cfurlpathstyle?language=objc
type URLPathStyle Index

const (
	KURLHFSPathStyle     URLPathStyle = 1
	KURLPOSIXPathStyle   URLPathStyle = 0
	KURLWindowsPathStyle URLPathStyle = 2
)

// The entity type identification codes that the parser uses to describe XML entities. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/cfxmlentitytypecode?language=objc
type XMLEntityTypeCode Index

const (
	KXMLEntityTypeCharacter      XMLEntityTypeCode = 4
	KXMLEntityTypeParameter      XMLEntityTypeCode = 0
	KXMLEntityTypeParsedExternal XMLEntityTypeCode = 2
	KXMLEntityTypeParsedInternal XMLEntityTypeCode = 1
	KXMLEntityTypeUnparsed       XMLEntityTypeCode = 3
)

// The various XML data type identification codes that the parser uses to describe XML structures. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/cfxmlnodetypecode?language=objc
type XMLNodeTypeCode Index

const (
	KXMLNodeTypeAttribute                XMLNodeTypeCode = 3
	KXMLNodeTypeAttributeListDeclaration XMLNodeTypeCode = 15
	KXMLNodeTypeCDATASection             XMLNodeTypeCode = 7
	KXMLNodeTypeComment                  XMLNodeTypeCode = 5
	KXMLNodeTypeDocument                 XMLNodeTypeCode = 1
	KXMLNodeTypeDocumentFragment         XMLNodeTypeCode = 8
	KXMLNodeTypeDocumentType             XMLNodeTypeCode = 11
	KXMLNodeTypeElement                  XMLNodeTypeCode = 2
	KXMLNodeTypeElementTypeDeclaration   XMLNodeTypeCode = 14
	KXMLNodeTypeEntity                   XMLNodeTypeCode = 9
	KXMLNodeTypeEntityReference          XMLNodeTypeCode = 10
	KXMLNodeTypeNotation                 XMLNodeTypeCode = 13
	KXMLNodeTypeProcessingInstruction    XMLNodeTypeCode = 4
	KXMLNodeTypeText                     XMLNodeTypeCode = 6
	KXMLNodeTypeWhitespace               XMLNodeTypeCode = 12
)

// Options you can use to control the parser's treatment of an XML document. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/cfxmlparseroptions?language=objc
type XMLParserOptions OptionFlags

const (
	KXMLParserAddImpliedAttributes    XMLParserOptions = 32
	KXMLParserAllOptions              XMLParserOptions = 16777215
	KXMLParserNoOptions               XMLParserOptions = 0
	KXMLParserReplacePhysicalEntities XMLParserOptions = 4
	KXMLParserResolveExternalEntities XMLParserOptions = 16
	KXMLParserSkipMetaData            XMLParserOptions = 2
	KXMLParserSkipWhitespace          XMLParserOptions = 8
	KXMLParserValidateDocument        XMLParserOptions = 1
)

// The various status and error flags that can be returned by the parser. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/cfxmlparserstatuscode?language=objc
type XMLParserStatusCode Index

const (
	KXMLErrorElementlessDocument            XMLParserStatusCode = 11
	KXMLErrorEncodingConversionFailure      XMLParserStatusCode = 3
	KXMLErrorMalformedCDSect                XMLParserStatusCode = 7
	KXMLErrorMalformedCharacterReference    XMLParserStatusCode = 13
	KXMLErrorMalformedCloseTag              XMLParserStatusCode = 8
	KXMLErrorMalformedComment               XMLParserStatusCode = 12
	KXMLErrorMalformedDTD                   XMLParserStatusCode = 5
	KXMLErrorMalformedDocument              XMLParserStatusCode = 10
	KXMLErrorMalformedName                  XMLParserStatusCode = 6
	KXMLErrorMalformedParsedCharacterData   XMLParserStatusCode = 14
	KXMLErrorMalformedProcessingInstruction XMLParserStatusCode = 4
	KXMLErrorMalformedStartTag              XMLParserStatusCode = 9
	KXMLErrorNoData                         XMLParserStatusCode = 15
	KXMLErrorUnexpectedEOF                  XMLParserStatusCode = 1
	KXMLErrorUnknownEncoding                XMLParserStatusCode = 2
	KXMLStatusParseInProgress               XMLParserStatusCode = -1
	KXMLStatusParseNotBegun                 XMLParserStatusCode = -2
	KXMLStatusParseSuccessful               XMLParserStatusCode = 0
)
