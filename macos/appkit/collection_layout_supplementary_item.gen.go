// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [CollectionLayoutSupplementaryItem] class.
var CollectionLayoutSupplementaryItemClass = _CollectionLayoutSupplementaryItemClass{objc.GetClass("NSCollectionLayoutSupplementaryItem")}

type _CollectionLayoutSupplementaryItemClass struct {
	objc.Class
}

// An interface definition for the [CollectionLayoutSupplementaryItem] class.
type ICollectionLayoutSupplementaryItem interface {
	ICollectionLayoutItem
	ContainerAnchor() CollectionLayoutAnchor
	ItemAnchor() CollectionLayoutAnchor
	ElementKind() string
	ZIndex() int
	SetZIndex(value int)
}

// An object used to add an extra visual decoration to an item in a collection view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nscollectionlayoutsupplementaryitem?language=objc
type CollectionLayoutSupplementaryItem struct {
	CollectionLayoutItem
}

func CollectionLayoutSupplementaryItemFrom(ptr unsafe.Pointer) CollectionLayoutSupplementaryItem {
	return CollectionLayoutSupplementaryItem{
		CollectionLayoutItem: CollectionLayoutItemFrom(ptr),
	}
}

func (cc _CollectionLayoutSupplementaryItemClass) SupplementaryItemWithLayoutSizeElementKindContainerAnchorItemAnchor(layoutSize ICollectionLayoutSize, elementKind string, containerAnchor ICollectionLayoutAnchor, itemAnchor ICollectionLayoutAnchor) CollectionLayoutSupplementaryItem {
	rv := objc.Call[CollectionLayoutSupplementaryItem](cc, objc.Sel("supplementaryItemWithLayoutSize:elementKind:containerAnchor:itemAnchor:"), objc.Ptr(layoutSize), elementKind, objc.Ptr(containerAnchor), objc.Ptr(itemAnchor))
	return rv
}

// Creates a supplementary item of the specified size and element kind, an anchor relative to a container, and an anchor relative to an item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nscollectionlayoutsupplementaryitem/3213900-supplementaryitemwithlayoutsize?language=objc
func CollectionLayoutSupplementaryItem_SupplementaryItemWithLayoutSizeElementKindContainerAnchorItemAnchor(layoutSize ICollectionLayoutSize, elementKind string, containerAnchor ICollectionLayoutAnchor, itemAnchor ICollectionLayoutAnchor) CollectionLayoutSupplementaryItem {
	return CollectionLayoutSupplementaryItemClass.SupplementaryItemWithLayoutSizeElementKindContainerAnchorItemAnchor(layoutSize, elementKind, containerAnchor, itemAnchor)
}

func (cc _CollectionLayoutSupplementaryItemClass) Alloc() CollectionLayoutSupplementaryItem {
	rv := objc.Call[CollectionLayoutSupplementaryItem](cc, objc.Sel("alloc"))
	return rv
}

func CollectionLayoutSupplementaryItem_Alloc() CollectionLayoutSupplementaryItem {
	return CollectionLayoutSupplementaryItemClass.Alloc()
}

func (cc _CollectionLayoutSupplementaryItemClass) New() CollectionLayoutSupplementaryItem {
	rv := objc.Call[CollectionLayoutSupplementaryItem](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCollectionLayoutSupplementaryItem() CollectionLayoutSupplementaryItem {
	return CollectionLayoutSupplementaryItemClass.New()
}

func (c_ CollectionLayoutSupplementaryItem) Init() CollectionLayoutSupplementaryItem {
	rv := objc.Call[CollectionLayoutSupplementaryItem](c_, objc.Sel("init"))
	return rv
}

func (cc _CollectionLayoutSupplementaryItemClass) ItemWithLayoutSize(layoutSize ICollectionLayoutSize) CollectionLayoutSupplementaryItem {
	rv := objc.Call[CollectionLayoutSupplementaryItem](cc, objc.Sel("itemWithLayoutSize:"), objc.Ptr(layoutSize))
	return rv
}

// Creates an item of the specified size. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nscollectionlayoutitem/3213871-itemwithlayoutsize?language=objc
func CollectionLayoutSupplementaryItem_ItemWithLayoutSize(layoutSize ICollectionLayoutSize) CollectionLayoutSupplementaryItem {
	return CollectionLayoutSupplementaryItemClass.ItemWithLayoutSize(layoutSize)
}

// The anchor between the supplementary item and the container it's attached to. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nscollectionlayoutsupplementaryitem/3199109-containeranchor?language=objc
func (c_ CollectionLayoutSupplementaryItem) ContainerAnchor() CollectionLayoutAnchor {
	rv := objc.Call[CollectionLayoutAnchor](c_, objc.Sel("containerAnchor"))
	return rv
}

// The anchor between the supplementary item and the item it's attached to. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nscollectionlayoutsupplementaryitem/3199113-itemanchor?language=objc
func (c_ CollectionLayoutSupplementaryItem) ItemAnchor() CollectionLayoutAnchor {
	rv := objc.Call[CollectionLayoutAnchor](c_, objc.Sel("itemAnchor"))
	return rv
}

// A string that identifies the type of supplementary item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nscollectionlayoutsupplementaryitem/3199110-elementkind?language=objc
func (c_ CollectionLayoutSupplementaryItem) ElementKind() string {
	rv := objc.Call[string](c_, objc.Sel("elementKind"))
	return rv
}

// The vertical stacking order of the supplementary item in relation to other items in the section. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nscollectionlayoutsupplementaryitem/3199114-zindex?language=objc
func (c_ CollectionLayoutSupplementaryItem) ZIndex() int {
	rv := objc.Call[int](c_, objc.Sel("zIndex"))
	return rv
}

// The vertical stacking order of the supplementary item in relation to other items in the section. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nscollectionlayoutsupplementaryitem/3199114-zindex?language=objc
func (c_ CollectionLayoutSupplementaryItem) SetZIndex(value int) {
	objc.Call[objc.Void](c_, objc.Sel("setZIndex:"), value)
}
