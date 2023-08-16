// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [CollectionLayoutEdgeSpacing] class.
var CollectionLayoutEdgeSpacingClass = _CollectionLayoutEdgeSpacingClass{objc.GetClass("NSCollectionLayoutEdgeSpacing")}

type _CollectionLayoutEdgeSpacingClass struct {
	objc.Class
}

// An interface definition for the [CollectionLayoutEdgeSpacing] class.
type ICollectionLayoutEdgeSpacing interface {
	objc.IObject
	Leading() CollectionLayoutSpacing
	Trailing() CollectionLayoutSpacing
	Bottom() CollectionLayoutSpacing
	Top() CollectionLayoutSpacing
}

// An object that defines the space around the edges of items in a collection view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nscollectionlayoutedgespacing?language=objc
type CollectionLayoutEdgeSpacing struct {
	objc.Object
}

func CollectionLayoutEdgeSpacingFrom(ptr unsafe.Pointer) CollectionLayoutEdgeSpacing {
	return CollectionLayoutEdgeSpacing{
		Object: objc.ObjectFrom(ptr),
	}
}

func (cc _CollectionLayoutEdgeSpacingClass) SpacingForLeadingTopTrailingBottom(leading ICollectionLayoutSpacing, top ICollectionLayoutSpacing, trailing ICollectionLayoutSpacing, bottom ICollectionLayoutSpacing) CollectionLayoutEdgeSpacing {
	rv := objc.Call[CollectionLayoutEdgeSpacing](cc, objc.Sel("spacingForLeading:top:trailing:bottom:"), objc.Ptr(leading), objc.Ptr(top), objc.Ptr(trailing), objc.Ptr(bottom))
	return rv
}

// Creates an edge spacing object with the specified leading, top, trailing, and bottom spacing. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nscollectionlayoutedgespacing/3199068-spacingforleading?language=objc
func CollectionLayoutEdgeSpacing_SpacingForLeadingTopTrailingBottom(leading ICollectionLayoutSpacing, top ICollectionLayoutSpacing, trailing ICollectionLayoutSpacing, bottom ICollectionLayoutSpacing) CollectionLayoutEdgeSpacing {
	return CollectionLayoutEdgeSpacingClass.SpacingForLeadingTopTrailingBottom(leading, top, trailing, bottom)
}

func (cc _CollectionLayoutEdgeSpacingClass) Alloc() CollectionLayoutEdgeSpacing {
	rv := objc.Call[CollectionLayoutEdgeSpacing](cc, objc.Sel("alloc"))
	return rv
}

func CollectionLayoutEdgeSpacing_Alloc() CollectionLayoutEdgeSpacing {
	return CollectionLayoutEdgeSpacingClass.Alloc()
}

func (cc _CollectionLayoutEdgeSpacingClass) New() CollectionLayoutEdgeSpacing {
	rv := objc.Call[CollectionLayoutEdgeSpacing](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCollectionLayoutEdgeSpacing() CollectionLayoutEdgeSpacing {
	return CollectionLayoutEdgeSpacingClass.New()
}

func (c_ CollectionLayoutEdgeSpacing) Init() CollectionLayoutEdgeSpacing {
	rv := objc.Call[CollectionLayoutEdgeSpacing](c_, objc.Sel("init"))
	return rv
}

// The leading edge spacing value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nscollectionlayoutedgespacing/3199069-leading?language=objc
func (c_ CollectionLayoutEdgeSpacing) Leading() CollectionLayoutSpacing {
	rv := objc.Call[CollectionLayoutSpacing](c_, objc.Sel("leading"))
	return rv
}

// The trailing edge spacing value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nscollectionlayoutedgespacing/3199071-trailing?language=objc
func (c_ CollectionLayoutEdgeSpacing) Trailing() CollectionLayoutSpacing {
	rv := objc.Call[CollectionLayoutSpacing](c_, objc.Sel("trailing"))
	return rv
}

// The bottom edge spacing value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nscollectionlayoutedgespacing/3199065-bottom?language=objc
func (c_ CollectionLayoutEdgeSpacing) Bottom() CollectionLayoutSpacing {
	rv := objc.Call[CollectionLayoutSpacing](c_, objc.Sel("bottom"))
	return rv
}

// The top edge spacing value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nscollectionlayoutedgespacing/3199070-top?language=objc
func (c_ CollectionLayoutEdgeSpacing) Top() CollectionLayoutSpacing {
	rv := objc.Call[CollectionLayoutSpacing](c_, objc.Sel("top"))
	return rv
}
