// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [GridColumn] class.
var GridColumnClass = _GridColumnClass{objc.GetClass("NSGridColumn")}

type _GridColumnClass struct {
	objc.Class
}

// An interface definition for the [GridColumn] class.
type IGridColumn interface {
	objc.IObject
	MergeCellsInRange(range_ foundation.Range)
	CellAtIndex(index int) GridCell
	Width() float64
	SetWidth(value float64)
	IsHidden() bool
	SetHidden(value bool)
	LeadingPadding() float64
	SetLeadingPadding(value float64)
	GridView() GridView
	XPlacement() GridCellPlacement
	SetXPlacement(value GridCellPlacement)
	TrailingPadding() float64
	SetTrailingPadding(value float64)
	NumberOfCells() int
}

// A column within a grid view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgridcolumn?language=objc
type GridColumn struct {
	objc.Object
}

func GridColumnFrom(ptr unsafe.Pointer) GridColumn {
	return GridColumn{
		Object: objc.ObjectFrom(ptr),
	}
}

func (gc _GridColumnClass) Alloc() GridColumn {
	rv := objc.Call[GridColumn](gc, objc.Sel("alloc"))
	return rv
}

func GridColumn_Alloc() GridColumn {
	return GridColumnClass.Alloc()
}

func (gc _GridColumnClass) New() GridColumn {
	rv := objc.Call[GridColumn](gc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewGridColumn() GridColumn {
	return GridColumnClass.New()
}

func (g_ GridColumn) Init() GridColumn {
	rv := objc.Call[GridColumn](g_, objc.Sel("init"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgridcolumn/1639752-mergecellsinrange?language=objc
func (g_ GridColumn) MergeCellsInRange(range_ foundation.Range) {
	objc.Call[objc.Void](g_, objc.Sel("mergeCellsInRange:"), range_)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgridcolumn/1639683-cellatindex?language=objc
func (g_ GridColumn) CellAtIndex(index int) GridCell {
	rv := objc.Call[GridCell](g_, objc.Sel("cellAtIndex:"), index)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgridcolumn/1639679-width?language=objc
func (g_ GridColumn) Width() float64 {
	rv := objc.Call[float64](g_, objc.Sel("width"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgridcolumn/1639679-width?language=objc
func (g_ GridColumn) SetWidth(value float64) {
	objc.Call[objc.Void](g_, objc.Sel("setWidth:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgridcolumn/1639742-hidden?language=objc
func (g_ GridColumn) IsHidden() bool {
	rv := objc.Call[bool](g_, objc.Sel("isHidden"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgridcolumn/1639742-hidden?language=objc
func (g_ GridColumn) SetHidden(value bool) {
	objc.Call[objc.Void](g_, objc.Sel("setHidden:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgridcolumn/1639769-leadingpadding?language=objc
func (g_ GridColumn) LeadingPadding() float64 {
	rv := objc.Call[float64](g_, objc.Sel("leadingPadding"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgridcolumn/1639769-leadingpadding?language=objc
func (g_ GridColumn) SetLeadingPadding(value float64) {
	objc.Call[objc.Void](g_, objc.Sel("setLeadingPadding:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgridcolumn/1639675-gridview?language=objc
func (g_ GridColumn) GridView() GridView {
	rv := objc.Call[GridView](g_, objc.Sel("gridView"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgridcolumn/1639663-xplacement?language=objc
func (g_ GridColumn) XPlacement() GridCellPlacement {
	rv := objc.Call[GridCellPlacement](g_, objc.Sel("xPlacement"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgridcolumn/1639663-xplacement?language=objc
func (g_ GridColumn) SetXPlacement(value GridCellPlacement) {
	objc.Call[objc.Void](g_, objc.Sel("setXPlacement:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgridcolumn/1639670-trailingpadding?language=objc
func (g_ GridColumn) TrailingPadding() float64 {
	rv := objc.Call[float64](g_, objc.Sel("trailingPadding"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgridcolumn/1639670-trailingpadding?language=objc
func (g_ GridColumn) SetTrailingPadding(value float64) {
	objc.Call[objc.Void](g_, objc.Sel("setTrailingPadding:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgridcolumn/1639719-numberofcells?language=objc
func (g_ GridColumn) NumberOfCells() int {
	rv := objc.Call[int](g_, objc.Sel("numberOfCells"))
	return rv
}
