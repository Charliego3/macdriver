// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [CollectionViewLayout] class.
var CollectionViewLayoutClass = _CollectionViewLayoutClass{objc.GetClass("NSCollectionViewLayout")}

type _CollectionViewLayoutClass struct {
	objc.Class
}

// An interface definition for the [CollectionViewLayout] class.
type ICollectionViewLayout interface {
	objc.IObject
	IndexPathsToInsertForSupplementaryViewOfKind(elementKind CollectionViewSupplementaryElementKind) foundation.Set
	PrepareForTransitionFromLayout(oldLayout ICollectionViewLayout)
	RegisterNibForDecorationViewOfKind(nib INib, elementKind CollectionViewDecorationElementKind)
	PrepareLayout()
	LayoutAttributesForInterItemGapBeforeIndexPath(indexPath foundation.IIndexPath) CollectionViewLayoutAttributes
	InvalidateLayout()
	IndexPathsToInsertForDecorationViewOfKind(elementKind CollectionViewDecorationElementKind) foundation.Set
	InvalidationContextForBoundsChange(newBounds foundation.Rect) CollectionViewLayoutInvalidationContext
	ShouldInvalidateLayoutForPreferredLayoutAttributesWithOriginalAttributes(preferredAttributes ICollectionViewLayoutAttributes, originalAttributes ICollectionViewLayoutAttributes) bool
	InvalidationContextForPreferredLayoutAttributesWithOriginalAttributes(preferredAttributes ICollectionViewLayoutAttributes, originalAttributes ICollectionViewLayoutAttributes) CollectionViewLayoutInvalidationContext
	RegisterClassForDecorationViewOfKind(viewClass objc.IClass, elementKind CollectionViewDecorationElementKind)
	TargetContentOffsetForProposedContentOffset(proposedContentOffset foundation.Point) foundation.Point
	LayoutAttributesForElementsInRect(rect foundation.Rect) []CollectionViewLayoutAttributes
	LayoutAttributesForDecorationViewOfKindAtIndexPath(elementKind CollectionViewDecorationElementKind, indexPath foundation.IIndexPath) CollectionViewLayoutAttributes
	LayoutAttributesForDropTargetAtPoint(pointInCollectionView foundation.Point) CollectionViewLayoutAttributes
	InvalidateLayoutWithContext(context ICollectionViewLayoutInvalidationContext)
	PrepareForCollectionViewUpdates(updateItems []ICollectionViewUpdateItem)
	FinalLayoutAttributesForDisappearingDecorationElementOfKindAtIndexPath(elementKind CollectionViewDecorationElementKind, decorationIndexPath foundation.IIndexPath) CollectionViewLayoutAttributes
	IndexPathsToDeleteForDecorationViewOfKind(elementKind CollectionViewDecorationElementKind) foundation.Set
	IndexPathsToDeleteForSupplementaryViewOfKind(elementKind CollectionViewSupplementaryElementKind) foundation.Set
	FinalizeCollectionViewUpdates()
	PrepareForTransitionToLayout(newLayout ICollectionViewLayout)
	InitialLayoutAttributesForAppearingDecorationElementOfKindAtIndexPath(elementKind CollectionViewDecorationElementKind, decorationIndexPath foundation.IIndexPath) CollectionViewLayoutAttributes
	FinalizeLayoutTransition()
	PrepareForAnimatedBoundsChange(oldBounds foundation.Rect)
	LayoutAttributesForItemAtIndexPath(indexPath foundation.IIndexPath) CollectionViewLayoutAttributes
	LayoutAttributesForSupplementaryViewOfKindAtIndexPath(elementKind CollectionViewSupplementaryElementKind, indexPath foundation.IIndexPath) CollectionViewLayoutAttributes
	FinalizeAnimatedBoundsChange()
	ShouldInvalidateLayoutForBoundsChange(newBounds foundation.Rect) bool
	CollectionViewContentSize() foundation.Size
	CollectionView() CollectionView
}

// An abstract base class that you subclass and use to generate layout information for a collection view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewlayout?language=objc
type CollectionViewLayout struct {
	objc.Object
}

func CollectionViewLayoutFrom(ptr unsafe.Pointer) CollectionViewLayout {
	return CollectionViewLayout{
		Object: objc.ObjectFrom(ptr),
	}
}

func (cc _CollectionViewLayoutClass) Alloc() CollectionViewLayout {
	rv := objc.Call[CollectionViewLayout](cc, objc.Sel("alloc"))
	return rv
}

func CollectionViewLayout_Alloc() CollectionViewLayout {
	return CollectionViewLayoutClass.Alloc()
}

func (cc _CollectionViewLayoutClass) New() CollectionViewLayout {
	rv := objc.Call[CollectionViewLayout](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCollectionViewLayout() CollectionViewLayout {
	return CollectionViewLayoutClass.New()
}

func (c_ CollectionViewLayout) Init() CollectionViewLayout {
	rv := objc.Call[CollectionViewLayout](c_, objc.Sel("init"))
	return rv
}

// Returns the index paths for any supplementary views that the layout object wants to add to the collection view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewlayout/1534044-indexpathstoinsertforsupplementa?language=objc
func (c_ CollectionViewLayout) IndexPathsToInsertForSupplementaryViewOfKind(elementKind CollectionViewSupplementaryElementKind) foundation.Set {
	rv := objc.Call[foundation.Set](c_, objc.Sel("indexPathsToInsertForSupplementaryViewOfKind:"), elementKind)
	return rv
}

// Prepares the layout object to be installed in the collection view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewlayout/1533383-preparefortransitionfromlayout?language=objc
func (c_ CollectionViewLayout) PrepareForTransitionFromLayout(oldLayout ICollectionViewLayout) {
	objc.Call[objc.Void](c_, objc.Sel("prepareForTransitionFromLayout:"), objc.Ptr(oldLayout))
}

// Registers a nib file to use when creating the layout’s decoration views. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewlayout/1535405-registernib?language=objc
func (c_ CollectionViewLayout) RegisterNibForDecorationViewOfKind(nib INib, elementKind CollectionViewDecorationElementKind) {
	objc.Call[objc.Void](c_, objc.Sel("registerNib:forDecorationViewOfKind:"), objc.Ptr(nib), elementKind)
}

// Prepares the layout object to begin laying out content. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewlayout/1530434-preparelayout?language=objc
func (c_ CollectionViewLayout) PrepareLayout() {
	objc.Call[objc.Void](c_, objc.Sel("prepareLayout"))
}

// Returns layout attributes for the inter-item gap at the specified location in your layout. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewlayout/1528296-layoutattributesforinteritemgapb?language=objc
func (c_ CollectionViewLayout) LayoutAttributesForInterItemGapBeforeIndexPath(indexPath foundation.IIndexPath) CollectionViewLayoutAttributes {
	rv := objc.Call[CollectionViewLayoutAttributes](c_, objc.Sel("layoutAttributesForInterItemGapBeforeIndexPath:"), objc.Ptr(indexPath))
	return rv
}

// Invalidates all layout information and triggers a layout update. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewlayout/1534293-invalidatelayout?language=objc
func (c_ CollectionViewLayout) InvalidateLayout() {
	objc.Call[objc.Void](c_, objc.Sel("invalidateLayout"))
}

// Returns the index paths for any decoration views that the layout object wants to add to the collection view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewlayout/1528098-indexpathstoinsertfordecorationv?language=objc
func (c_ CollectionViewLayout) IndexPathsToInsertForDecorationViewOfKind(elementKind CollectionViewDecorationElementKind) foundation.Set {
	rv := objc.Call[foundation.Set](c_, objc.Sel("indexPathsToInsertForDecorationViewOfKind:"), elementKind)
	return rv
}

// Returns an invalidation context object that defines the portions of the layout that need to be updated. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewlayout/1534058-invalidationcontextforboundschan?language=objc
func (c_ CollectionViewLayout) InvalidationContextForBoundsChange(newBounds foundation.Rect) CollectionViewLayoutInvalidationContext {
	rv := objc.Call[CollectionViewLayoutInvalidationContext](c_, objc.Sel("invalidationContextForBoundsChange:"), newBounds)
	return rv
}

// Returns a Boolean indicating whether changes to a cell’s layout attributes trigger a larger layout update. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewlayout/1535536-shouldinvalidatelayoutforpreferr?language=objc
func (c_ CollectionViewLayout) ShouldInvalidateLayoutForPreferredLayoutAttributesWithOriginalAttributes(preferredAttributes ICollectionViewLayoutAttributes, originalAttributes ICollectionViewLayoutAttributes) bool {
	rv := objc.Call[bool](c_, objc.Sel("shouldInvalidateLayoutForPreferredLayoutAttributes:withOriginalAttributes:"), objc.Ptr(preferredAttributes), objc.Ptr(originalAttributes))
	return rv
}

// Returns an invalidation context object that defines the portions of the layout that need to be updated. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewlayout/1535343-invalidationcontextforpreferredl?language=objc
func (c_ CollectionViewLayout) InvalidationContextForPreferredLayoutAttributesWithOriginalAttributes(preferredAttributes ICollectionViewLayoutAttributes, originalAttributes ICollectionViewLayoutAttributes) CollectionViewLayoutInvalidationContext {
	rv := objc.Call[CollectionViewLayoutInvalidationContext](c_, objc.Sel("invalidationContextForPreferredLayoutAttributes:withOriginalAttributes:"), objc.Ptr(preferredAttributes), objc.Ptr(originalAttributes))
	return rv
}

// Registers a class to use when creating the layout’s decoration views. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewlayout/1533837-registerclass?language=objc
func (c_ CollectionViewLayout) RegisterClassForDecorationViewOfKind(viewClass objc.IClass, elementKind CollectionViewDecorationElementKind) {
	objc.Call[objc.Void](c_, objc.Sel("registerClass:forDecorationViewOfKind:"), objc.Ptr(viewClass), elementKind)
}

// Returns the offset value to use after an animated layout update or change. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewlayout/1535608-targetcontentoffsetforproposedco?language=objc
func (c_ CollectionViewLayout) TargetContentOffsetForProposedContentOffset(proposedContentOffset foundation.Point) foundation.Point {
	rv := objc.Call[foundation.Point](c_, objc.Sel("targetContentOffsetForProposedContentOffset:"), proposedContentOffset)
	return rv
}

// Returns the layout attribute objects for all items and views in the specified rectangle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewlayout/1524636-layoutattributesforelementsinrec?language=objc
func (c_ CollectionViewLayout) LayoutAttributesForElementsInRect(rect foundation.Rect) []CollectionViewLayoutAttributes {
	rv := objc.Call[[]CollectionViewLayoutAttributes](c_, objc.Sel("layoutAttributesForElementsInRect:"), rect)
	return rv
}

// Returns the layout attributes of the decoration view at the specified location in your layout. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewlayout/1527929-layoutattributesfordecorationvie?language=objc
func (c_ CollectionViewLayout) LayoutAttributesForDecorationViewOfKindAtIndexPath(elementKind CollectionViewDecorationElementKind, indexPath foundation.IIndexPath) CollectionViewLayoutAttributes {
	rv := objc.Call[CollectionViewLayoutAttributes](c_, objc.Sel("layoutAttributesForDecorationViewOfKind:atIndexPath:"), elementKind, objc.Ptr(indexPath))
	return rv
}

// Returns layout attributes for the drop target at the specified point. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewlayout/1535931-layoutattributesfordroptargetatp?language=objc
func (c_ CollectionViewLayout) LayoutAttributesForDropTargetAtPoint(pointInCollectionView foundation.Point) CollectionViewLayoutAttributes {
	rv := objc.Call[CollectionViewLayoutAttributes](c_, objc.Sel("layoutAttributesForDropTargetAtPoint:"), pointInCollectionView)
	return rv
}

// Invalidates specific parts of the layout using the specified context object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewlayout/1533282-invalidatelayoutwithcontext?language=objc
func (c_ CollectionViewLayout) InvalidateLayoutWithContext(context ICollectionViewLayoutInvalidationContext) {
	objc.Call[objc.Void](c_, objc.Sel("invalidateLayoutWithContext:"), objc.Ptr(context))
}

// Performs needed tasks before items are inserted, deleted, or moved within the collection view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewlayout/1533485-prepareforcollectionviewupdates?language=objc
func (c_ CollectionViewLayout) PrepareForCollectionViewUpdates(updateItems []ICollectionViewUpdateItem) {
	objc.Call[objc.Void](c_, objc.Sel("prepareForCollectionViewUpdates:"), updateItems)
}

// Returns the ending layout information for a decoration view being removed from the collection view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewlayout/1535917-finallayoutattributesfordisappea?language=objc
func (c_ CollectionViewLayout) FinalLayoutAttributesForDisappearingDecorationElementOfKindAtIndexPath(elementKind CollectionViewDecorationElementKind, decorationIndexPath foundation.IIndexPath) CollectionViewLayoutAttributes {
	rv := objc.Call[CollectionViewLayoutAttributes](c_, objc.Sel("finalLayoutAttributesForDisappearingDecorationElementOfKind:atIndexPath:"), elementKind, objc.Ptr(decorationIndexPath))
	return rv
}

// Returns index paths for any decoration views that the layout object wants to remove from the collection view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewlayout/1529635-indexpathstodeletefordecorationv?language=objc
func (c_ CollectionViewLayout) IndexPathsToDeleteForDecorationViewOfKind(elementKind CollectionViewDecorationElementKind) foundation.Set {
	rv := objc.Call[foundation.Set](c_, objc.Sel("indexPathsToDeleteForDecorationViewOfKind:"), elementKind)
	return rv
}

// Returns the index paths for any supplementary views that the layout object wants to remove from the collection view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewlayout/1529622-indexpathstodeleteforsupplementa?language=objc
func (c_ CollectionViewLayout) IndexPathsToDeleteForSupplementaryViewOfKind(elementKind CollectionViewSupplementaryElementKind) foundation.Set {
	rv := objc.Call[foundation.Set](c_, objc.Sel("indexPathsToDeleteForSupplementaryViewOfKind:"), elementKind)
	return rv
}

// Performs needed steps after items are inserted, deleted, or moved within a collection view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewlayout/1534473-finalizecollectionviewupdates?language=objc
func (c_ CollectionViewLayout) FinalizeCollectionViewUpdates() {
	objc.Call[objc.Void](c_, objc.Sel("finalizeCollectionViewUpdates"))
}

// Prepares the layout object to be uninstalled from the collection view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewlayout/1533945-preparefortransitiontolayout?language=objc
func (c_ CollectionViewLayout) PrepareForTransitionToLayout(newLayout ICollectionViewLayout) {
	objc.Call[objc.Void](c_, objc.Sel("prepareForTransitionToLayout:"), objc.Ptr(newLayout))
}

// Returns the starting layout information for a decoration view being added to the collection view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewlayout/1535003-initiallayoutattributesforappear?language=objc
func (c_ CollectionViewLayout) InitialLayoutAttributesForAppearingDecorationElementOfKindAtIndexPath(elementKind CollectionViewDecorationElementKind, decorationIndexPath foundation.IIndexPath) CollectionViewLayoutAttributes {
	rv := objc.Call[CollectionViewLayoutAttributes](c_, objc.Sel("initialLayoutAttributesForAppearingDecorationElementOfKind:atIndexPath:"), elementKind, objc.Ptr(decorationIndexPath))
	return rv
}

// Performs any final steps related to a layout transition before the transition animations actually occur. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewlayout/1534138-finalizelayouttransition?language=objc
func (c_ CollectionViewLayout) FinalizeLayoutTransition() {
	objc.Call[objc.Void](c_, objc.Sel("finalizeLayoutTransition"))
}

// Prepares the layout object for animated changes to the collection view’s bounds or for the insertion or deletion of items. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewlayout/1533678-prepareforanimatedboundschange?language=objc
func (c_ CollectionViewLayout) PrepareForAnimatedBoundsChange(oldBounds foundation.Rect) {
	objc.Call[objc.Void](c_, objc.Sel("prepareForAnimatedBoundsChange:"), oldBounds)
}

// Returns the layout attributes for the item at the specified index path. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewlayout/1534298-layoutattributesforitematindexpa?language=objc
func (c_ CollectionViewLayout) LayoutAttributesForItemAtIndexPath(indexPath foundation.IIndexPath) CollectionViewLayoutAttributes {
	rv := objc.Call[CollectionViewLayoutAttributes](c_, objc.Sel("layoutAttributesForItemAtIndexPath:"), objc.Ptr(indexPath))
	return rv
}

// Returns the layout attributes of the supplementary view at the specified location in your layout. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewlayout/1534287-layoutattributesforsupplementary?language=objc
func (c_ CollectionViewLayout) LayoutAttributesForSupplementaryViewOfKindAtIndexPath(elementKind CollectionViewSupplementaryElementKind, indexPath foundation.IIndexPath) CollectionViewLayoutAttributes {
	rv := objc.Call[CollectionViewLayoutAttributes](c_, objc.Sel("layoutAttributesForSupplementaryViewOfKind:atIndexPath:"), elementKind, objc.Ptr(indexPath))
	return rv
}

// Cleans up after any animated changes to the collection view’s bounds or after the insertion or deletion of items. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewlayout/1533163-finalizeanimatedboundschange?language=objc
func (c_ CollectionViewLayout) FinalizeAnimatedBoundsChange() {
	objc.Call[objc.Void](c_, objc.Sel("finalizeAnimatedBoundsChange"))
}

// Returns a Boolean indicating whether a bounds change triggers a layout update. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewlayout/1531047-shouldinvalidatelayoutforboundsc?language=objc
func (c_ CollectionViewLayout) ShouldInvalidateLayoutForBoundsChange(newBounds foundation.Rect) bool {
	rv := objc.Call[bool](c_, objc.Sel("shouldInvalidateLayoutForBoundsChange:"), newBounds)
	return rv
}

// Returns the class to use for layout attribute objects [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewlayout/1527555-layoutattributesclass?language=objc
func (cc _CollectionViewLayoutClass) LayoutAttributesClass() objc.Class {
	rv := objc.Call[objc.Class](cc, objc.Sel("layoutAttributesClass"))
	return rv
}

// Returns the class to use for layout attribute objects [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewlayout/1527555-layoutattributesclass?language=objc
func CollectionViewLayout_LayoutAttributesClass() objc.Class {
	return CollectionViewLayoutClass.LayoutAttributesClass()
}

// The width and height of the collection view’s contents. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewlayout/1532618-collectionviewcontentsize?language=objc
func (c_ CollectionViewLayout) CollectionViewContentSize() foundation.Size {
	rv := objc.Call[foundation.Size](c_, objc.Sel("collectionViewContentSize"))
	return rv
}

// Returns the class to use when creating an invalidation context object for the layout. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewlayout/1535858-invalidationcontextclass?language=objc
func (cc _CollectionViewLayoutClass) InvalidationContextClass() objc.Class {
	rv := objc.Call[objc.Class](cc, objc.Sel("invalidationContextClass"))
	return rv
}

// Returns the class to use when creating an invalidation context object for the layout. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewlayout/1535858-invalidationcontextclass?language=objc
func CollectionViewLayout_InvalidationContextClass() objc.Class {
	return CollectionViewLayoutClass.InvalidationContextClass()
}

// The collection view object currently using this layout. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewlayout/1533540-collectionview?language=objc
func (c_ CollectionViewLayout) CollectionView() CollectionView {
	rv := objc.Call[CollectionView](c_, objc.Sel("collectionView"))
	return rv
}
