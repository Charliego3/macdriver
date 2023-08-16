// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [GridCell] class.
var GridCellClass = _GridCellClass{objc.GetClass("NSGridCell")}

type _GridCellClass struct {
	objc.Class
}

// An interface definition for the [GridCell] class.
type IGridCell interface {
	objc.IObject
	RowAlignment() GridRowAlignment
	SetRowAlignment(value GridRowAlignment)
	ContentView() View
	SetContentView(value IView)
	CustomPlacementConstraints() []LayoutConstraint
	SetCustomPlacementConstraints(value []ILayoutConstraint)
	YPlacement() GridCellPlacement
	SetYPlacement(value GridCellPlacement)
	Column() GridColumn
	Row() GridRow
	XPlacement() GridCellPlacement
	SetXPlacement(value GridCellPlacement)
}

// An individual content area within a grid view, typically at the intersection of a row and a column. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgridcell?language=objc
type GridCell struct {
	objc.Object
}

func GridCellFrom(ptr unsafe.Pointer) GridCell {
	return GridCell{
		Object: objc.ObjectFrom(ptr),
	}
}

func (gc _GridCellClass) Alloc() GridCell {
	rv := objc.Call[GridCell](gc, objc.Sel("alloc"))
	return rv
}

func GridCell_Alloc() GridCell {
	return GridCellClass.Alloc()
}

func (gc _GridCellClass) New() GridCell {
	rv := objc.Call[GridCell](gc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewGridCell() GridCell {
	return GridCellClass.New()
}

func (g_ GridCell) Init() GridCell {
	rv := objc.Call[GridCell](g_, objc.Sel("init"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgridcell/1823686-rowalignment?language=objc
func (g_ GridCell) RowAlignment() GridRowAlignment {
	rv := objc.Call[GridRowAlignment](g_, objc.Sel("rowAlignment"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgridcell/1823686-rowalignment?language=objc
func (g_ GridCell) SetRowAlignment(value GridRowAlignment) {
	objc.Call[objc.Void](g_, objc.Sel("setRowAlignment:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgridcell/1639681-emptycontentview?language=objc
func (gc _GridCellClass) EmptyContentView() View {
	rv := objc.Call[View](gc, objc.Sel("emptyContentView"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgridcell/1639681-emptycontentview?language=objc
func GridCell_EmptyContentView() View {
	return GridCellClass.EmptyContentView()
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgridcell/1639721-contentview?language=objc
func (g_ GridCell) ContentView() View {
	rv := objc.Call[View](g_, objc.Sel("contentView"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgridcell/1639721-contentview?language=objc
func (g_ GridCell) SetContentView(value IView) {
	objc.Call[objc.Void](g_, objc.Sel("setContentView:"), objc.Ptr(value))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgridcell/1639717-customplacementconstraints?language=objc
func (g_ GridCell) CustomPlacementConstraints() []LayoutConstraint {
	rv := objc.Call[[]LayoutConstraint](g_, objc.Sel("customPlacementConstraints"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgridcell/1639717-customplacementconstraints?language=objc
func (g_ GridCell) SetCustomPlacementConstraints(value []ILayoutConstraint) {
	objc.Call[objc.Void](g_, objc.Sel("setCustomPlacementConstraints:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgridcell/1639737-yplacement?language=objc
func (g_ GridCell) YPlacement() GridCellPlacement {
	rv := objc.Call[GridCellPlacement](g_, objc.Sel("yPlacement"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgridcell/1639737-yplacement?language=objc
func (g_ GridCell) SetYPlacement(value GridCellPlacement) {
	objc.Call[objc.Void](g_, objc.Sel("setYPlacement:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgridcell/1639747-column?language=objc
func (g_ GridCell) Column() GridColumn {
	rv := objc.Call[GridColumn](g_, objc.Sel("column"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgridcell/1639763-row?language=objc
func (g_ GridCell) Row() GridRow {
	rv := objc.Call[GridRow](g_, objc.Sel("row"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgridcell/1639710-xplacement?language=objc
func (g_ GridCell) XPlacement() GridCellPlacement {
	rv := objc.Call[GridCellPlacement](g_, objc.Sel("xPlacement"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgridcell/1639710-xplacement?language=objc
func (g_ GridCell) SetXPlacement(value GridCellPlacement) {
	objc.Call[objc.Void](g_, objc.Sel("setXPlacement:"), value)
}
