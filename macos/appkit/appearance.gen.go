// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [Appearance] class.
var AppearanceClass = _AppearanceClass{objc.GetClass("NSAppearance")}

type _AppearanceClass struct {
	objc.Class
}

// An interface definition for the [Appearance] class.
type IAppearance interface {
	objc.IObject
	BestMatchFromAppearancesWithNames(appearances []AppearanceName) AppearanceName
	PerformAsCurrentDrawingAppearance(block func())
	Name() AppearanceName
	AllowsVibrancy() bool
}

// An object that manages standard appearance attributes for UI elements in an app. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsappearance?language=objc
type Appearance struct {
	objc.Object
}

func AppearanceFrom(ptr unsafe.Pointer) Appearance {
	return Appearance{
		Object: objc.ObjectFrom(ptr),
	}
}

func (a_ Appearance) InitWithAppearanceNamedBundle(name AppearanceName, bundle foundation.IBundle) Appearance {
	rv := objc.Call[Appearance](a_, objc.Sel("initWithAppearanceNamed:bundle:"), name, objc.Ptr(bundle))
	return rv
}

// Creates an appearance object from the named appearance file located in the specified bundle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsappearance/1529131-initwithappearancenamed?language=objc
func Appearance_InitWithAppearanceNamedBundle(name AppearanceName, bundle foundation.IBundle) Appearance {
	return AppearanceClass.Alloc().InitWithAppearanceNamedBundle(name, bundle)
}

func (ac _AppearanceClass) Alloc() Appearance {
	rv := objc.Call[Appearance](ac, objc.Sel("alloc"))
	return rv
}

func Appearance_Alloc() Appearance {
	return AppearanceClass.Alloc()
}

func (ac _AppearanceClass) New() Appearance {
	rv := objc.Call[Appearance](ac, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewAppearance() Appearance {
	return AppearanceClass.New()
}

func (a_ Appearance) Init() Appearance {
	rv := objc.Call[Appearance](a_, objc.Sel("init"))
	return rv
}

// Returns the appearance name that most closely matches the current appearance object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsappearance/2980972-bestmatchfromappearanceswithname?language=objc
func (a_ Appearance) BestMatchFromAppearancesWithNames(appearances []AppearanceName) AppearanceName {
	rv := objc.Call[AppearanceName](a_, objc.Sel("bestMatchFromAppearancesWithNames:"), appearances)
	return rv
}

// Sets the appearance to be the active drawing appearance and perform the specified block. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsappearance/3674525-performascurrentdrawingappearanc?language=objc
func (a_ Appearance) PerformAsCurrentDrawingAppearance(block func()) {
	objc.Call[objc.Void](a_, objc.Sel("performAsCurrentDrawingAppearance:"), block)
}

// Creates an appearance object based on the name of one of the standard system appearances. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsappearance/1529612-appearancenamed?language=objc
func (ac _AppearanceClass) AppearanceNamed(name AppearanceName) Appearance {
	rv := objc.Call[Appearance](ac, objc.Sel("appearanceNamed:"), name)
	return rv
}

// Creates an appearance object based on the name of one of the standard system appearances. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsappearance/1529612-appearancenamed?language=objc
func Appearance_AppearanceNamed(name AppearanceName) Appearance {
	return AppearanceClass.AppearanceNamed(name)
}

// The name of the appearance. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsappearance/1528677-name?language=objc
func (a_ Appearance) Name() AppearanceName {
	rv := objc.Call[AppearanceName](a_, objc.Sel("name"))
	return rv
}

// The appearance that the system uses for color and asset resolution, and that’s active for drawing, usually from locking focus on a view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsappearance/3674524-currentdrawingappearance?language=objc
func (ac _AppearanceClass) CurrentDrawingAppearance() Appearance {
	rv := objc.Call[Appearance](ac, objc.Sel("currentDrawingAppearance"))
	return rv
}

// The appearance that the system uses for color and asset resolution, and that’s active for drawing, usually from locking focus on a view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsappearance/3674524-currentdrawingappearance?language=objc
func Appearance_CurrentDrawingAppearance() Appearance {
	return AppearanceClass.CurrentDrawingAppearance()
}

// Specifies whether the current appearance allows vibrancy. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsappearance/1524694-allowsvibrancy?language=objc
func (a_ Appearance) AllowsVibrancy() bool {
	rv := objc.Call[bool](a_, objc.Sel("allowsVibrancy"))
	return rv
}
