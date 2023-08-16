// AUTO-GENERATED CODE, DO NOT MODIFY

package webkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [ContentWorld] class.
var ContentWorldClass = _ContentWorldClass{objc.GetClass("WKContentWorld")}

type _ContentWorldClass struct {
	objc.Class
}

// An interface definition for the [ContentWorld] class.
type IContentWorld interface {
	objc.IObject
	Name() string
}

// An object that defines a scope of execution for JavaScript code, and which you use to prevent conflicts between different scripts. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkcontentworld?language=objc
type ContentWorld struct {
	objc.Object
}

func ContentWorldFrom(ptr unsafe.Pointer) ContentWorld {
	return ContentWorld{
		Object: objc.ObjectFrom(ptr),
	}
}

func (cc _ContentWorldClass) Alloc() ContentWorld {
	rv := objc.Call[ContentWorld](cc, objc.Sel("alloc"))
	return rv
}

func ContentWorld_Alloc() ContentWorld {
	return ContentWorldClass.Alloc()
}

func (cc _ContentWorldClass) New() ContentWorld {
	rv := objc.Call[ContentWorld](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewContentWorld() ContentWorld {
	return ContentWorldClass.New()
}

func (c_ ContentWorld) Init() ContentWorld {
	rv := objc.Call[ContentWorld](c_, objc.Sel("init"))
	return rv
}

// Returns the custom content world with the specified name. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkcontentworld/3552419-worldwithname?language=objc
func (cc _ContentWorldClass) WorldWithName(name string) ContentWorld {
	rv := objc.Call[ContentWorld](cc, objc.Sel("worldWithName:"), name)
	return rv
}

// Returns the custom content world with the specified name. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkcontentworld/3552419-worldwithname?language=objc
func ContentWorld_WorldWithName(name string) ContentWorld {
	return ContentWorldClass.WorldWithName(name)
}

// The content world for the current webpage’s content. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkcontentworld/3552418-pageworld?language=objc
func (cc _ContentWorldClass) PageWorld() ContentWorld {
	rv := objc.Call[ContentWorld](cc, objc.Sel("pageWorld"))
	return rv
}

// The content world for the current webpage’s content. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkcontentworld/3552418-pageworld?language=objc
func ContentWorld_PageWorld() ContentWorld {
	return ContentWorldClass.PageWorld()
}

// The default world for clients. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkcontentworld/3552416-defaultclientworld?language=objc
func (cc _ContentWorldClass) DefaultClientWorld() ContentWorld {
	rv := objc.Call[ContentWorld](cc, objc.Sel("defaultClientWorld"))
	return rv
}

// The default world for clients. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkcontentworld/3552416-defaultclientworld?language=objc
func ContentWorld_DefaultClientWorld() ContentWorld {
	return ContentWorldClass.DefaultClientWorld()
}

// The name of a custom content world. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkcontentworld/3552417-name?language=objc
func (c_ ContentWorld) Name() string {
	rv := objc.Call[string](c_, objc.Sel("name"))
	return rv
}
