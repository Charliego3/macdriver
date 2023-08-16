// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [Screen] class.
var ScreenClass = _ScreenClass{objc.GetClass("NSScreen")}

type _ScreenClass struct {
	objc.Class
}

// An interface definition for the [Screen] class.
type IScreen interface {
	objc.IObject
	ConvertRectFromBacking(rect foundation.Rect) foundation.Rect
	ConvertRectToBacking(rect foundation.Rect) foundation.Rect
	CanRepresentDisplayGamut(displayGamut DisplayGamut) bool
	BackingAlignedRectOptions(rect foundation.Rect, options foundation.AlignmentOptions) foundation.Rect
	MinimumRefreshInterval() foundation.TimeInterval
	MaximumExtendedDynamicRangeColorComponentValue() float64
	VisibleFrame() foundation.Rect
	DisplayUpdateGranularity() foundation.TimeInterval
	MaximumRefreshInterval() foundation.TimeInterval
	LocalizedName() string
	MaximumPotentialExtendedDynamicRangeColorComponentValue() float64
	ColorSpace() ColorSpace
	SafeAreaInsets() foundation.EdgeInsets
	MaximumReferenceExtendedDynamicRangeColorComponentValue() float64
	LastDisplayUpdateTimestamp() foundation.TimeInterval
	DeviceDescription() map[DeviceDescriptionKey]objc.Object
	Depth() WindowDepth
	BackingScaleFactor() float64
	Frame() foundation.Rect
	SupportedWindowDepths() *WindowDepth
	MaximumFramesPerSecond() int
}

// An object that describes the attributes of a computer’s monitor or screen. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscreen?language=objc
type Screen struct {
	objc.Object
}

func ScreenFrom(ptr unsafe.Pointer) Screen {
	return Screen{
		Object: objc.ObjectFrom(ptr),
	}
}

func (sc _ScreenClass) Alloc() Screen {
	rv := objc.Call[Screen](sc, objc.Sel("alloc"))
	return rv
}

func Screen_Alloc() Screen {
	return ScreenClass.Alloc()
}

func (sc _ScreenClass) New() Screen {
	rv := objc.Call[Screen](sc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewScreen() Screen {
	return ScreenClass.New()
}

func (s_ Screen) Init() Screen {
	rv := objc.Call[Screen](s_, objc.Sel("init"))
	return rv
}

// Converts the rectangle from the device pixel aligned coordinates system of a screen. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscreen/1388364-convertrectfrombacking?language=objc
func (s_ Screen) ConvertRectFromBacking(rect foundation.Rect) foundation.Rect {
	rv := objc.Call[foundation.Rect](s_, objc.Sel("convertRectFromBacking:"), rect)
	return rv
}

// Converts the rectangle to the device pixel aligned coordinates system of a screen. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscreen/1388389-convertrecttobacking?language=objc
func (s_ Screen) ConvertRectToBacking(rect foundation.Rect) foundation.Rect {
	rv := objc.Call[foundation.Rect](s_, objc.Sel("convertRectToBacking:"), rect)
	return rv
}

// A Boolean value indicating whether the color space of the screen is capable of representing the specified display gamut. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscreen/2138325-canrepresentdisplaygamut?language=objc
func (s_ Screen) CanRepresentDisplayGamut(displayGamut DisplayGamut) bool {
	rv := objc.Call[bool](s_, objc.Sel("canRepresentDisplayGamut:"), displayGamut)
	return rv
}

// Converts a rectangle in global screen coordinates to a pixel aligned rectangle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscreen/1388381-backingalignedrect?language=objc
func (s_ Screen) BackingAlignedRectOptions(rect foundation.Rect, options foundation.AlignmentOptions) foundation.Rect {
	rv := objc.Call[foundation.Rect](s_, objc.Sel("backingAlignedRect:options:"), rect, options)
	return rv
}

// The shortest refresh interval that the screen supports. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscreen/3824747-minimumrefreshinterval?language=objc
func (s_ Screen) MinimumRefreshInterval() foundation.TimeInterval {
	rv := objc.Call[foundation.TimeInterval](s_, objc.Sel("minimumRefreshInterval"))
	return rv
}

// The current maximum color component value for the screen. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscreen/1388362-maximumextendeddynamicrangecolor?language=objc
func (s_ Screen) MaximumExtendedDynamicRangeColorComponentValue() float64 {
	rv := objc.Call[float64](s_, objc.Sel("maximumExtendedDynamicRangeColorComponentValue"))
	return rv
}

// The current location and dimensions of the visible screen. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscreen/1388369-visibleframe?language=objc
func (s_ Screen) VisibleFrame() foundation.Rect {
	rv := objc.Call[foundation.Rect](s_, objc.Sel("visibleFrame"))
	return rv
}

// The number of seconds between the screen’s supported update rates, for screens that support fixed update rates. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscreen/3824743-displayupdategranularity?language=objc
func (s_ Screen) DisplayUpdateGranularity() foundation.TimeInterval {
	rv := objc.Call[foundation.TimeInterval](s_, objc.Sel("displayUpdateGranularity"))
	return rv
}

// The largest refresh interval that the screen supports. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscreen/3824746-maximumrefreshinterval?language=objc
func (s_ Screen) MaximumRefreshInterval() foundation.TimeInterval {
	rv := objc.Call[foundation.TimeInterval](s_, objc.Sel("maximumRefreshInterval"))
	return rv
}

// The localized name of the display. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscreen/3228043-localizedname?language=objc
func (s_ Screen) LocalizedName() string {
	rv := objc.Call[string](s_, objc.Sel("localizedName"))
	return rv
}

// Returns a screen object representing the screen that can best represent color. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscreen/1388374-deepestscreen?language=objc
func (sc _ScreenClass) DeepestScreen() Screen {
	rv := objc.Call[Screen](sc, objc.Sel("deepestScreen"))
	return rv
}

// Returns a screen object representing the screen that can best represent color. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscreen/1388374-deepestscreen?language=objc
func Screen_DeepestScreen() Screen {
	return ScreenClass.DeepestScreen()
}

// The maximum possible color component value for the screen when it's in extended dynamic range (EDR) mode. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscreen/3180381-maximumpotentialextendeddynamicr?language=objc
func (s_ Screen) MaximumPotentialExtendedDynamicRangeColorComponentValue() float64 {
	rv := objc.Call[float64](s_, objc.Sel("maximumPotentialExtendedDynamicRangeColorComponentValue"))
	return rv
}

// The color space of the screen. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscreen/1388383-colorspace?language=objc
func (s_ Screen) ColorSpace() ColorSpace {
	rv := objc.Call[ColorSpace](s_, objc.Sel("colorSpace"))
	return rv
}

// The distances from the screen’s edges at which content isn’t obscured. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscreen/3882821-safeareainsets?language=objc
func (s_ Screen) SafeAreaInsets() foundation.EdgeInsets {
	rv := objc.Call[foundation.EdgeInsets](s_, objc.Sel("safeAreaInsets"))
	return rv
}

// The current maximum color component value for reference rendering to the screen. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscreen/3180382-maximumreferenceextendeddynamicr?language=objc
func (s_ Screen) MaximumReferenceExtendedDynamicRangeColorComponentValue() float64 {
	rv := objc.Call[float64](s_, objc.Sel("maximumReferenceExtendedDynamicRangeColorComponentValue"))
	return rv
}

// Returns an array of screen objects representing all of the screens available on the system. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscreen/1388393-screens?language=objc
func (sc _ScreenClass) Screens() []Screen {
	rv := objc.Call[[]Screen](sc, objc.Sel("screens"))
	return rv
}

// Returns an array of screen objects representing all of the screens available on the system. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscreen/1388393-screens?language=objc
func Screen_Screens() []Screen {
	return ScreenClass.Screens()
}

// The time of the last framebuffer update, expressed as the number of seconds since system startup. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscreen/3824744-lastdisplayupdatetimestamp?language=objc
func (s_ Screen) LastDisplayUpdateTimestamp() foundation.TimeInterval {
	rv := objc.Call[foundation.TimeInterval](s_, objc.Sel("lastDisplayUpdateTimestamp"))
	return rv
}

// The device dictionary for the screen. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscreen/1388360-devicedescription?language=objc
func (s_ Screen) DeviceDescription() map[DeviceDescriptionKey]objc.Object {
	rv := objc.Call[map[DeviceDescriptionKey]objc.Object](s_, objc.Sel("deviceDescription"))
	return rv
}

// Returns a Boolean value indicating whether each screen can have its own set of spaces. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscreen/1388365-screenshaveseparatespaces?language=objc
func (sc _ScreenClass) ScreensHaveSeparateSpaces() bool {
	rv := objc.Call[bool](sc, objc.Sel("screensHaveSeparateSpaces"))
	return rv
}

// Returns a Boolean value indicating whether each screen can have its own set of spaces. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscreen/1388365-screenshaveseparatespaces?language=objc
func Screen_ScreensHaveSeparateSpaces() bool {
	return ScreenClass.ScreensHaveSeparateSpaces()
}

// The current bit depth and colorspace information of the screen. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscreen/1388373-depth?language=objc
func (s_ Screen) Depth() WindowDepth {
	rv := objc.Call[WindowDepth](s_, objc.Sel("depth"))
	return rv
}

// The backing store pixel scale factor for the screen. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscreen/1388385-backingscalefactor?language=objc
func (s_ Screen) BackingScaleFactor() float64 {
	rv := objc.Call[float64](s_, objc.Sel("backingScaleFactor"))
	return rv
}

// The dimensions and location of the screen. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscreen/1388387-frame?language=objc
func (s_ Screen) Frame() foundation.Rect {
	rv := objc.Call[foundation.Rect](s_, objc.Sel("frame"))
	return rv
}

// Returns the screen object containing the window with the keyboard focus. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscreen/1388371-mainscreen?language=objc
func (sc _ScreenClass) MainScreen() Screen {
	rv := objc.Call[Screen](sc, objc.Sel("mainScreen"))
	return rv
}

// Returns the screen object containing the window with the keyboard focus. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscreen/1388371-mainscreen?language=objc
func Screen_MainScreen() Screen {
	return ScreenClass.MainScreen()
}

// A zero-terminated array of the window depths supported by the screen. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscreen/1388391-supportedwindowdepths?language=objc
func (s_ Screen) SupportedWindowDepths() *WindowDepth {
	rv := objc.Call[*WindowDepth](s_, objc.Sel("supportedWindowDepths"))
	return rv
}

// The maximum number of frames per second that the screen supports. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscreen/3824745-maximumframespersecond?language=objc
func (s_ Screen) MaximumFramesPerSecond() int {
	rv := objc.Call[int](s_, objc.Sel("maximumFramesPerSecond"))
	return rv
}
