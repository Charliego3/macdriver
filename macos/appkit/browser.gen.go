// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [Browser] class.
var BrowserClass = _BrowserClass{objc.GetClass("NSBrowser")}

type _BrowserClass struct {
	objc.Class
}

// An interface definition for the [Browser] class.
type IBrowser interface {
	IControl
	SendAction() bool
	SetDefaultColumnWidth(columnWidth float64)
	ScrollRowToVisibleInColumn(row int, column int)
	ReloadDataForRowIndexesInColumn(rowIndexes foundation.IIndexSet, column int)
	Path() string
	SelectedRowInColumn(column int) int
	SelectedCellInColumn(column int) objc.Object
	ColumnContentWidthForColumnWidth(columnWidth float64) float64
	SelectRowInColumn(row int, column int)
	DefaultColumnWidth() float64
	PathToColumn(column int) string
	TitleFrameOfColumn(column int) foundation.Rect
	DrawTitleOfColumnInRect(column int, rect foundation.Rect)
	SetPath(path string) bool
	SelectedRowIndexesInColumn(column int) foundation.IndexSet
	DoClick(sender objc.IObject)
	ScrollColumnsLeftBy(shiftAmount int)
	LoadColumnZero()
	LoadedCellAtRowColumn(row int, col int) objc.Object
	IndexPathForColumn(column int) foundation.IndexPath
	ScrollColumnToVisible(column int)
	DoDoubleClick(sender objc.IObject)
	DraggingImageForRowsWithIndexesInColumnWithEventOffset(rowIndexes foundation.IIndexSet, column int, event IEvent, dragImageOffset foundation.PointPointer) Image
	SelectAll(sender objc.IObject)
	FrameOfColumn(column int) foundation.Rect
	ScrollColumnsRightBy(shiftAmount int)
	SetTitleOfColumn(string_ string, column int)
	IsLeafItem(item objc.IObject) bool
	ReloadColumn(column int)
	FrameOfInsideOfColumn(column int) foundation.Rect
	EditItemAtIndexPathWithEventSelect(indexPath foundation.IIndexPath, event IEvent, select_ bool)
	SelectRowIndexesInColumn(indexes foundation.IIndexSet, column int)
	TitleOfColumn(column int) string
	NoteHeightOfRowsWithIndexesChangedInColumn(indexSet foundation.IIndexSet, columnIndex int)
	ValidateVisibleColumns()
	CanDragRowsWithIndexesInColumnWithEvent(rowIndexes foundation.IIndexSet, column int, event IEvent) bool
	SetDraggingSourceOperationMaskForLocal(mask DragOperation, isLocal bool)
	ColumnWidthForColumnContentWidth(columnContentWidth float64) float64
	SetCellClass(factoryId objc.IClass)
	AddColumn()
	FrameOfRowInColumn(row int, column int) foundation.Rect
	GetRowColumnForPoint(row *int, column *int, point foundation.Point) bool
	Tile()
	ItemAtRowInColumn(row int, column int) objc.Object
	SetWidthOfColumn(columnWidth float64, columnIndex int)
	ItemAtIndexPath(indexPath foundation.IIndexPath) objc.Object
	WidthOfColumn(column int) float64
	ParentForItemsInColumn(column int) objc.Object
	DoubleAction() objc.Selector
	SetDoubleAction(value objc.Selector)
	SelectedCells() []Cell
	PrefersAllColumnUserResizing() bool
	SetPrefersAllColumnUserResizing(value bool)
	SelectedCell() objc.Object
	SelectionIndexPath() foundation.IndexPath
	SetSelectionIndexPath(value foundation.IIndexPath)
	AllowsEmptySelection() bool
	SetAllowsEmptySelection(value bool)
	RowHeight() float64
	SetRowHeight(value float64)
	AutohidesScroller() bool
	SetAutohidesScroller(value bool)
	ClickedColumn() int
	IsTitled() bool
	SetTitled(value bool)
	CellPrototype() objc.Object
	SetCellPrototype(value objc.IObject)
	LastVisibleColumn() int
	AllowsTypeSelect() bool
	SetAllowsTypeSelect(value bool)
	IsLoaded() bool
	ColumnResizingType() BrowserColumnResizingType
	SetColumnResizingType(value BrowserColumnResizingType)
	FirstVisibleColumn() int
	ReusesColumns() bool
	SetReusesColumns(value bool)
	Delegate() BrowserDelegateWrapper
	SetDelegate(value PBrowserDelegate)
	SetDelegateObject(valueObject objc.IObject)
	BackgroundColor() Color
	SetBackgroundColor(value IColor)
	AllowsBranchSelection() bool
	SetAllowsBranchSelection(value bool)
	SelectionIndexPaths() []foundation.IndexPath
	SetSelectionIndexPaths(value []foundation.IIndexPath)
	NumberOfVisibleColumns() int
	SendsActionOnArrowKeys() bool
	SetSendsActionOnArrowKeys(value bool)
	ClickedRow() int
	AllowsMultipleSelection() bool
	SetAllowsMultipleSelection(value bool)
	SelectedColumn() int
	TitleHeight() float64
	PathSeparator() string
	SetPathSeparator(value string)
	SeparatesColumns() bool
	SetSeparatesColumns(value bool)
	MaxVisibleColumns() int
	SetMaxVisibleColumns(value int)
	ColumnsAutosaveName() BrowserColumnsAutosaveName
	SetColumnsAutosaveName(value BrowserColumnsAutosaveName)
	LastColumn() int
	SetLastColumn(value int)
	MinColumnWidth() float64
	SetMinColumnWidth(value float64)
	TakesTitleFromPreviousColumn() bool
	SetTakesTitleFromPreviousColumn(value bool)
	HasHorizontalScroller() bool
	SetHasHorizontalScroller(value bool)
}

// An interface that displays a hierarchically organized list of data items that can be navigated and selected. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbrowser?language=objc
type Browser struct {
	Control
}

func BrowserFrom(ptr unsafe.Pointer) Browser {
	return Browser{
		Control: ControlFrom(ptr),
	}
}

func (bc _BrowserClass) Alloc() Browser {
	rv := objc.Call[Browser](bc, objc.Sel("alloc"))
	return rv
}

func Browser_Alloc() Browser {
	return BrowserClass.Alloc()
}

func (bc _BrowserClass) New() Browser {
	rv := objc.Call[Browser](bc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewBrowser() Browser {
	return BrowserClass.New()
}

func (b_ Browser) Init() Browser {
	rv := objc.Call[Browser](b_, objc.Sel("init"))
	return rv
}

func (b_ Browser) InitWithFrame(frameRect foundation.Rect) Browser {
	rv := objc.Call[Browser](b_, objc.Sel("initWithFrame:"), frameRect)
	return rv
}

// Initializes a control with the specified frame rectangle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontrol/1428900-initwithframe?language=objc
func Browser_InitWithFrame(frameRect foundation.Rect) Browser {
	return BrowserClass.Alloc().InitWithFrame(frameRect)
}

// Sends the action message to the target. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbrowser/1407675-sendaction?language=objc
func (b_ Browser) SendAction() bool {
	rv := objc.Call[bool](b_, objc.Sel("sendAction"))
	return rv
}

// Sets the default column width for new browser columns that do not otherwise have an initial width from defaults or the browser’s delegate. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbrowser/1407651-setdefaultcolumnwidth?language=objc
func (b_ Browser) SetDefaultColumnWidth(columnWidth float64) {
	objc.Call[objc.Void](b_, objc.Sel("setDefaultColumnWidth:"), columnWidth)
}

// Scrolls the specified row to be visible within the specified column. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbrowser/1407719-scrollrowtovisible?language=objc
func (b_ Browser) ScrollRowToVisibleInColumn(row int, column int) {
	objc.Call[objc.Void](b_, objc.Sel("scrollRowToVisible:inColumn:"), row, column)
}

// Updates the rows in the column with the specified column index with indexes in the specified set. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbrowser/1407805-reloaddataforrowindexes?language=objc
func (b_ Browser) ReloadDataForRowIndexesInColumn(rowIndexes foundation.IIndexSet, column int) {
	objc.Call[objc.Void](b_, objc.Sel("reloadDataForRowIndexes:inColumn:"), objc.Ptr(rowIndexes), column)
}

// Returns a string representing the browser’s current path. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbrowser/1407731-path?language=objc
func (b_ Browser) Path() string {
	rv := objc.Call[string](b_, objc.Sel("path"))
	return rv
}

// Returns the row index of the selected cell in the specified column. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbrowser/1407528-selectedrowincolumn?language=objc
func (b_ Browser) SelectedRowInColumn(column int) int {
	rv := objc.Call[int](b_, objc.Sel("selectedRowInColumn:"), column)
	return rv
}

// Returns the last (lowest) cell selected in the given column. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbrowser/1407592-selectedcellincolumn?language=objc
func (b_ Browser) SelectedCellInColumn(column int) objc.Object {
	rv := objc.Call[objc.Object](b_, objc.Sel("selectedCellInColumn:"), column)
	return rv
}

// Returns the content width for a given column width. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbrowser/1407522-columncontentwidthforcolumnwidth?language=objc
func (b_ Browser) ColumnContentWidthForColumnWidth(columnWidth float64) float64 {
	rv := objc.Call[float64](b_, objc.Sel("columnContentWidthForColumnWidth:"), columnWidth)
	return rv
}

// Selects the cell at the specified row and column index. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbrowser/1407560-selectrow?language=objc
func (b_ Browser) SelectRowInColumn(row int, column int) {
	objc.Call[objc.Void](b_, objc.Sel("selectRow:inColumn:"), row, column)
}

// Returns the default column width of the browser’s columns. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbrowser/1407615-defaultcolumnwidth?language=objc
func (b_ Browser) DefaultColumnWidth() float64 {
	rv := objc.Call[float64](b_, objc.Sel("defaultColumnWidth"))
	return rv
}

// Returns a string representing the path from the first column up to, but not including, the column at the given index. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbrowser/1407727-pathtocolumn?language=objc
func (b_ Browser) PathToColumn(column int) string {
	rv := objc.Call[string](b_, objc.Sel("pathToColumn:"), column)
	return rv
}

// Returns the bounds of the title frame for the specified column. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbrowser/1407544-titleframeofcolumn?language=objc
func (b_ Browser) TitleFrameOfColumn(column int) foundation.Rect {
	rv := objc.Call[foundation.Rect](b_, objc.Sel("titleFrameOfColumn:"), column)
	return rv
}

// Draws the title for the specified column within the given rectangle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbrowser/1407809-drawtitleofcolumn?language=objc
func (b_ Browser) DrawTitleOfColumnInRect(column int, rect foundation.Rect) {
	objc.Call[objc.Void](b_, objc.Sel("drawTitleOfColumn:inRect:"), column, rect)
}

// Sets the path to be displayed by the browser. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbrowser/1407751-setpath?language=objc
func (b_ Browser) SetPath(path string) bool {
	rv := objc.Call[bool](b_, objc.Sel("setPath:"), path)
	return rv
}

// Provides the indexes of the selected rows in a given column of the browser. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbrowser/1407807-selectedrowindexesincolumn?language=objc
func (b_ Browser) SelectedRowIndexesInColumn(column int) foundation.IndexSet {
	rv := objc.Call[foundation.IndexSet](b_, objc.Sel("selectedRowIndexesInColumn:"), column)
	return rv
}

// Responds to (single) mouse clicks in a column of the browser. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbrowser/1407600-doclick?language=objc
func (b_ Browser) DoClick(sender objc.IObject) {
	objc.Call[objc.Void](b_, objc.Sel("doClick:"), sender)
}

// Scrolls columns left by the specified number of columns. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbrowser/1407747-scrollcolumnsleftby?language=objc
func (b_ Browser) ScrollColumnsLeftBy(shiftAmount int) {
	objc.Call[objc.Void](b_, objc.Sel("scrollColumnsLeftBy:"), shiftAmount)
}

// Loads column 0; unloads previously loaded columns. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbrowser/1407745-loadcolumnzero?language=objc
func (b_ Browser) LoadColumnZero() {
	objc.Call[objc.Void](b_, objc.Sel("loadColumnZero"))
}

// Loads, if necessary, and returns the cell at the specified row and column location. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbrowser/1407581-loadedcellatrow?language=objc
func (b_ Browser) LoadedCellAtRowColumn(row int, col int) objc.Object {
	rv := objc.Call[objc.Object](b_, objc.Sel("loadedCellAtRow:column:"), row, col)
	return rv
}

// Returns the index path of the item whose children are displayed in the given column. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbrowser/1407739-indexpathforcolumn?language=objc
func (b_ Browser) IndexPathForColumn(column int) foundation.IndexPath {
	rv := objc.Call[foundation.IndexPath](b_, objc.Sel("indexPathForColumn:"), column)
	return rv
}

// Scrolls to make the specified column visible. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbrowser/1407516-scrollcolumntovisible?language=objc
func (b_ Browser) ScrollColumnToVisible(column int) {
	objc.Call[objc.Void](b_, objc.Sel("scrollColumnToVisible:"), column)
}

// Responds to double clicks in a column of the browser. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbrowser/1407800-dodoubleclick?language=objc
func (b_ Browser) DoDoubleClick(sender objc.IObject) {
	objc.Call[objc.Void](b_, objc.Sel("doDoubleClick:"), sender)
}

// Provides an image to represent dragged rows during a drag operation on the browser. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbrowser/1407532-draggingimageforrowswithindexes?language=objc
func (b_ Browser) DraggingImageForRowsWithIndexesInColumnWithEventOffset(rowIndexes foundation.IIndexSet, column int, event IEvent, dragImageOffset foundation.PointPointer) Image {
	rv := objc.Call[Image](b_, objc.Sel("draggingImageForRowsWithIndexes:inColumn:withEvent:offset:"), objc.Ptr(rowIndexes), column, objc.Ptr(event), dragImageOffset)
	return rv
}

// Selects all cells in the last column of the browser. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbrowser/1407717-selectall?language=objc
func (b_ Browser) SelectAll(sender objc.IObject) {
	objc.Call[objc.Void](b_, objc.Sel("selectAll:"), sender)
}

// Returns the rectangle containing the given column. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbrowser/1407620-frameofcolumn?language=objc
func (b_ Browser) FrameOfColumn(column int) foundation.Rect {
	rv := objc.Call[foundation.Rect](b_, objc.Sel("frameOfColumn:"), column)
	return rv
}

// Scrolls columns right by the specified number of columns. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbrowser/1407684-scrollcolumnsrightby?language=objc
func (b_ Browser) ScrollColumnsRightBy(shiftAmount int) {
	objc.Call[objc.Void](b_, objc.Sel("scrollColumnsRightBy:"), shiftAmount)
}

// Sets the title of the given column. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbrowser/1407575-settitle?language=objc
func (b_ Browser) SetTitleOfColumn(string_ string, column int) {
	objc.Call[objc.Void](b_, objc.Sel("setTitle:ofColumn:"), string_, column)
}

// Returns whether the specified item is a leaf item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbrowser/1407713-isleafitem?language=objc
func (b_ Browser) IsLeafItem(item objc.IObject) bool {
	rv := objc.Call[bool](b_, objc.Sel("isLeafItem:"), item)
	return rv
}

// Reloads the given column. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbrowser/1407764-reloadcolumn?language=objc
func (b_ Browser) ReloadColumn(column int) {
	objc.Call[objc.Void](b_, objc.Sel("reloadColumn:"), column)
}

// Returns the rectangle containing the specified column, not including borders. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbrowser/1407617-frameofinsideofcolumn?language=objc
func (b_ Browser) FrameOfInsideOfColumn(column int) foundation.Rect {
	rv := objc.Call[foundation.Rect](b_, objc.Sel("frameOfInsideOfColumn:"), column)
	return rv
}

// Begins editing the item at the specified path. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbrowser/1407558-edititematindexpath?language=objc
func (b_ Browser) EditItemAtIndexPathWithEventSelect(indexPath foundation.IIndexPath, event IEvent, select_ bool) {
	objc.Call[objc.Void](b_, objc.Sel("editItemAtIndexPath:withEvent:select:"), objc.Ptr(indexPath), objc.Ptr(event), select_)
}

// Specifies the selected rows in a given column of the browser. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbrowser/1407618-selectrowindexes?language=objc
func (b_ Browser) SelectRowIndexesInColumn(indexes foundation.IIndexSet, column int) {
	objc.Call[objc.Void](b_, objc.Sel("selectRowIndexes:inColumn:"), objc.Ptr(indexes), column)
}

// Returns the title displayed for the given column. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbrowser/1407794-titleofcolumn?language=objc
func (b_ Browser) TitleOfColumn(column int) string {
	rv := objc.Call[string](b_, objc.Sel("titleOfColumn:"), column)
	return rv
}

// Immediately retiles the browser’s columns using row heights specified by the browser’s delegate. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbrowser/1407608-noteheightofrowswithindexeschang?language=objc
func (b_ Browser) NoteHeightOfRowsWithIndexesChangedInColumn(indexSet foundation.IIndexSet, columnIndex int) {
	objc.Call[objc.Void](b_, objc.Sel("noteHeightOfRowsWithIndexesChanged:inColumn:"), objc.Ptr(indexSet), columnIndex)
}

// Validates the browser’s visible columns. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbrowser/1407725-validatevisiblecolumns?language=objc
func (b_ Browser) ValidateVisibleColumns() {
	objc.Call[objc.Void](b_, objc.Sel("validateVisibleColumns"))
}

// Indicates whether the browser can attempt to initiate a drag of the given rows for the given event. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbrowser/1407632-candragrowswithindexes?language=objc
func (b_ Browser) CanDragRowsWithIndexesInColumnWithEvent(rowIndexes foundation.IIndexSet, column int, event IEvent) bool {
	rv := objc.Call[bool](b_, objc.Sel("canDragRowsWithIndexes:inColumn:withEvent:"), objc.Ptr(rowIndexes), column, objc.Ptr(event))
	return rv
}

// Specifies the drag-operation mask for dragging operations with local or external destinations. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbrowser/1407559-setdraggingsourceoperationmask?language=objc
func (b_ Browser) SetDraggingSourceOperationMaskForLocal(mask DragOperation, isLocal bool) {
	objc.Call[objc.Void](b_, objc.Sel("setDraggingSourceOperationMask:forLocal:"), mask, isLocal)
}

// Returns the column width for the width of the given column’s content. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbrowser/1407648-columnwidthforcolumncontentwidth?language=objc
func (b_ Browser) ColumnWidthForColumnContentWidth(columnContentWidth float64) float64 {
	rv := objc.Call[float64](b_, objc.Sel("columnWidthForColumnContentWidth:"), columnContentWidth)
	return rv
}

// Sets the class of the cell to be used by the matrices in the columns of the browser. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbrowser/1407811-setcellclass?language=objc
func (b_ Browser) SetCellClass(factoryId objc.IClass) {
	objc.Call[objc.Void](b_, objc.Sel("setCellClass:"), objc.Ptr(factoryId))
}

// Adds a column to the right of the last column. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbrowser/1407723-addcolumn?language=objc
func (b_ Browser) AddColumn() {
	objc.Call[objc.Void](b_, objc.Sel("addColumn"))
}

// Returns the frame of the cell at the specified location, including the expandable arrow. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbrowser/1407622-frameofrow?language=objc
func (b_ Browser) FrameOfRowInColumn(row int, column int) foundation.Rect {
	rv := objc.Call[foundation.Rect](b_, objc.Sel("frameOfRow:inColumn:"), row, column)
	return rv
}

// Gets the row and column coordinates for the specified point, if a cell exists at that point. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbrowser/1407606-getrow?language=objc
func (b_ Browser) GetRowColumnForPoint(row *int, column *int, point foundation.Point) bool {
	rv := objc.Call[bool](b_, objc.Sel("getRow:column:forPoint:"), row, column, point)
	return rv
}

// Adjusts the various subviews of the browser—scrollers, columns, titles, and so on—without redrawing. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbrowser/1407788-tile?language=objc
func (b_ Browser) Tile() {
	objc.Call[objc.Void](b_, objc.Sel("tile"))
}

// Returns the item located at the specified row and column. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbrowser/1407701-itematrow?language=objc
func (b_ Browser) ItemAtRowInColumn(row int, column int) objc.Object {
	rv := objc.Call[objc.Object](b_, objc.Sel("itemAtRow:inColumn:"), row, column)
	return rv
}

// Sets the width of the specified column. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbrowser/1407679-setwidth?language=objc
func (b_ Browser) SetWidthOfColumn(columnWidth float64, columnIndex int) {
	objc.Call[objc.Void](b_, objc.Sel("setWidth:ofColumn:"), columnWidth, columnIndex)
}

// Returns the item at the specified index path. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbrowser/1407673-itematindexpath?language=objc
func (b_ Browser) ItemAtIndexPath(indexPath foundation.IIndexPath) objc.Object {
	rv := objc.Call[objc.Object](b_, objc.Sel("itemAtIndexPath:"), objc.Ptr(indexPath))
	return rv
}

// Removes the column configuration data stored under the given name from the application’s user defaults. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbrowser/1407668-removesavedcolumnswithautosavena?language=objc
func (bc _BrowserClass) RemoveSavedColumnsWithAutosaveName(name BrowserColumnsAutosaveName) {
	objc.Call[objc.Void](bc, objc.Sel("removeSavedColumnsWithAutosaveName:"), name)
}

// Removes the column configuration data stored under the given name from the application’s user defaults. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbrowser/1407668-removesavedcolumnswithautosavena?language=objc
func Browser_RemoveSavedColumnsWithAutosaveName(name BrowserColumnsAutosaveName) {
	BrowserClass.RemoveSavedColumnsWithAutosaveName(name)
}

// Returns the width of the specified column. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbrowser/1407778-widthofcolumn?language=objc
func (b_ Browser) WidthOfColumn(column int) float64 {
	rv := objc.Call[float64](b_, objc.Sel("widthOfColumn:"), column)
	return rv
}

// Returns the item that contains the children located in the specified column. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbrowser/1407570-parentforitemsincolumn?language=objc
func (b_ Browser) ParentForItemsInColumn(column int) objc.Object {
	rv := objc.Call[objc.Object](b_, objc.Sel("parentForItemsInColumn:"), column)
	return rv
}

// The browser’s double-click action method. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbrowser/1407698-doubleaction?language=objc
func (b_ Browser) DoubleAction() objc.Selector {
	rv := objc.Call[objc.Selector](b_, objc.Sel("doubleAction"))
	return rv
}

// The browser’s double-click action method. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbrowser/1407698-doubleaction?language=objc
func (b_ Browser) SetDoubleAction(value objc.Selector) {
	objc.Call[objc.Void](b_, objc.Sel("setDoubleAction:"), value)
}

// All cells selected in the rightmost column. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbrowser/1407588-selectedcells?language=objc
func (b_ Browser) SelectedCells() []Cell {
	rv := objc.Call[[]Cell](b_, objc.Sel("selectedCells"))
	return rv
}

// A Boolean that indicates whether the browser is set to resize all columns simultaneously rather than resizing a single column at a time. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbrowser/1407690-prefersallcolumnuserresizing?language=objc
func (b_ Browser) PrefersAllColumnUserResizing() bool {
	rv := objc.Call[bool](b_, objc.Sel("prefersAllColumnUserResizing"))
	return rv
}

// A Boolean that indicates whether the browser is set to resize all columns simultaneously rather than resizing a single column at a time. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbrowser/1407690-prefersallcolumnuserresizing?language=objc
func (b_ Browser) SetPrefersAllColumnUserResizing(value bool) {
	objc.Call[objc.Void](b_, objc.Sel("setPrefersAllColumnUserResizing:"), value)
}

// The last (rightmost and lowest) selected cell. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbrowser/1407730-selectedcell?language=objc
func (b_ Browser) SelectedCell() objc.Object {
	rv := objc.Call[objc.Object](b_, objc.Sel("selectedCell"))
	return rv
}

// The index path of the item selected in the browser. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbrowser/1407507-selectionindexpath?language=objc
func (b_ Browser) SelectionIndexPath() foundation.IndexPath {
	rv := objc.Call[foundation.IndexPath](b_, objc.Sel("selectionIndexPath"))
	return rv
}

// The index path of the item selected in the browser. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbrowser/1407507-selectionindexpath?language=objc
func (b_ Browser) SetSelectionIndexPath(value foundation.IIndexPath) {
	objc.Call[objc.Void](b_, objc.Sel("setSelectionIndexPath:"), objc.Ptr(value))
}

// A Boolean that indicates whether there can be nothing selected. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbrowser/1407585-allowsemptyselection?language=objc
func (b_ Browser) AllowsEmptySelection() bool {
	rv := objc.Call[bool](b_, objc.Sel("allowsEmptySelection"))
	return rv
}

// A Boolean that indicates whether there can be nothing selected. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbrowser/1407585-allowsemptyselection?language=objc
func (b_ Browser) SetAllowsEmptySelection(value bool) {
	objc.Call[objc.Void](b_, objc.Sel("setAllowsEmptySelection:"), value)
}

// The height of the browser’s rows. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbrowser/1407642-rowheight?language=objc
func (b_ Browser) RowHeight() float64 {
	rv := objc.Call[float64](b_, objc.Sel("rowHeight"))
	return rv
}

// The height of the browser’s rows. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbrowser/1407642-rowheight?language=objc
func (b_ Browser) SetRowHeight(value float64) {
	objc.Call[objc.Void](b_, objc.Sel("setRowHeight:"), value)
}

// A Boolean that indicates whether the browser automatically hides its scroller. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbrowser/1407696-autohidesscroller?language=objc
func (b_ Browser) AutohidesScroller() bool {
	rv := objc.Call[bool](b_, objc.Sel("autohidesScroller"))
	return rv
}

// A Boolean that indicates whether the browser automatically hides its scroller. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbrowser/1407696-autohidesscroller?language=objc
func (b_ Browser) SetAutohidesScroller(value bool) {
	objc.Call[objc.Void](b_, objc.Sel("setAutohidesScroller:"), value)
}

// The column number of the cell that the user clicked to display a context menu. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbrowser/1407590-clickedcolumn?language=objc
func (b_ Browser) ClickedColumn() int {
	rv := objc.Call[int](b_, objc.Sel("clickedColumn"))
	return rv
}

// A Boolean that indicates whether columns display titles. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbrowser/1407735-titled?language=objc
func (b_ Browser) IsTitled() bool {
	rv := objc.Call[bool](b_, objc.Sel("isTitled"))
	return rv
}

// A Boolean that indicates whether columns display titles. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbrowser/1407735-titled?language=objc
func (b_ Browser) SetTitled(value bool) {
	objc.Call[objc.Void](b_, objc.Sel("setTitled:"), value)
}

// The prototype NSCell for displaying items in the matrices in the columns of the browser. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbrowser/1407662-cellprototype?language=objc
func (b_ Browser) CellPrototype() objc.Object {
	rv := objc.Call[objc.Object](b_, objc.Sel("cellPrototype"))
	return rv
}

// The prototype NSCell for displaying items in the matrices in the columns of the browser. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbrowser/1407662-cellprototype?language=objc
func (b_ Browser) SetCellPrototype(value objc.IObject) {
	objc.Call[objc.Void](b_, objc.Sel("setCellPrototype:"), value)
}

// The index of the last visible column. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbrowser/1407610-lastvisiblecolumn?language=objc
func (b_ Browser) LastVisibleColumn() int {
	rv := objc.Call[int](b_, objc.Sel("lastVisibleColumn"))
	return rv
}

// A Boolean that indicates whether the browser allows keystroke-based selection (type select). [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbrowser/1407682-allowstypeselect?language=objc
func (b_ Browser) AllowsTypeSelect() bool {
	rv := objc.Call[bool](b_, objc.Sel("allowsTypeSelect"))
	return rv
}

// A Boolean that indicates whether the browser allows keystroke-based selection (type select). [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbrowser/1407682-allowstypeselect?language=objc
func (b_ Browser) SetAllowsTypeSelect(value bool) {
	objc.Call[objc.Void](b_, objc.Sel("setAllowsTypeSelect:"), value)
}

// A Boolean that indicates whether column 0 is loaded. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbrowser/1407758-loaded?language=objc
func (b_ Browser) IsLoaded() bool {
	rv := objc.Call[bool](b_, objc.Sel("isLoaded"))
	return rv
}

// A constant indicating the browser’s column resizing type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbrowser/1407694-columnresizingtype?language=objc
func (b_ Browser) ColumnResizingType() BrowserColumnResizingType {
	rv := objc.Call[BrowserColumnResizingType](b_, objc.Sel("columnResizingType"))
	return rv
}

// A constant indicating the browser’s column resizing type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbrowser/1407694-columnresizingtype?language=objc
func (b_ Browser) SetColumnResizingType(value BrowserColumnResizingType) {
	objc.Call[objc.Void](b_, objc.Sel("setColumnResizingType:"), value)
}

// The index of the first visible column. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbrowser/1407703-firstvisiblecolumn?language=objc
func (b_ Browser) FirstVisibleColumn() int {
	rv := objc.Call[int](b_, objc.Sel("firstVisibleColumn"))
	return rv
}

// A Boolean that indicates whether the browser reuses matrix objects after their columns are unloaded. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbrowser/1407792-reusescolumns?language=objc
func (b_ Browser) ReusesColumns() bool {
	rv := objc.Call[bool](b_, objc.Sel("reusesColumns"))
	return rv
}

// A Boolean that indicates whether the browser reuses matrix objects after their columns are unloaded. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbrowser/1407792-reusescolumns?language=objc
func (b_ Browser) SetReusesColumns(value bool) {
	objc.Call[objc.Void](b_, objc.Sel("setReusesColumns:"), value)
}

// The browser’s delegate. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbrowser/1407686-delegate?language=objc
func (b_ Browser) Delegate() BrowserDelegateWrapper {
	rv := objc.Call[BrowserDelegateWrapper](b_, objc.Sel("delegate"))
	return rv
}

// The browser’s delegate. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbrowser/1407686-delegate?language=objc
func (b_ Browser) SetDelegate(value PBrowserDelegate) {
	po0 := objc.WrapAsProtocol("NSBrowserDelegate", value)
	objc.SetAssociatedObject(b_, objc.AssociationKey("setDelegate"), po0, objc.ASSOCIATION_RETAIN)
	objc.Call[objc.Void](b_, objc.Sel("setDelegate:"), po0)
}

// The browser’s delegate. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbrowser/1407686-delegate?language=objc
func (b_ Browser) SetDelegateObject(valueObject objc.IObject) {
	objc.Call[objc.Void](b_, objc.Sel("setDelegate:"), objc.Ptr(valueObject))
}

// The browser’s background color. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbrowser/1407520-backgroundcolor?language=objc
func (b_ Browser) BackgroundColor() Color {
	rv := objc.Call[Color](b_, objc.Sel("backgroundColor"))
	return rv
}

// The browser’s background color. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbrowser/1407520-backgroundcolor?language=objc
func (b_ Browser) SetBackgroundColor(value IColor) {
	objc.Call[objc.Void](b_, objc.Sel("setBackgroundColor:"), objc.Ptr(value))
}

// A Boolean that indicates whether the user can select branch items. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbrowser/1407796-allowsbranchselection?language=objc
func (b_ Browser) AllowsBranchSelection() bool {
	rv := objc.Call[bool](b_, objc.Sel("allowsBranchSelection"))
	return rv
}

// A Boolean that indicates whether the user can select branch items. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbrowser/1407796-allowsbranchselection?language=objc
func (b_ Browser) SetAllowsBranchSelection(value bool) {
	objc.Call[objc.Void](b_, objc.Sel("setAllowsBranchSelection:"), value)
}

// An array containing the index paths of all items selected in the browser. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbrowser/1407536-selectionindexpaths?language=objc
func (b_ Browser) SelectionIndexPaths() []foundation.IndexPath {
	rv := objc.Call[[]foundation.IndexPath](b_, objc.Sel("selectionIndexPaths"))
	return rv
}

// An array containing the index paths of all items selected in the browser. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbrowser/1407536-selectionindexpaths?language=objc
func (b_ Browser) SetSelectionIndexPaths(value []foundation.IIndexPath) {
	objc.Call[objc.Void](b_, objc.Sel("setSelectionIndexPaths:"), value)
}

// The number of visible columns. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbrowser/1407700-numberofvisiblecolumns?language=objc
func (b_ Browser) NumberOfVisibleColumns() int {
	rv := objc.Call[int](b_, objc.Sel("numberOfVisibleColumns"))
	return rv
}

// A Boolean that indicates whether pressing an arrow key causes an action message to be sent. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbrowser/1407614-sendsactiononarrowkeys?language=objc
func (b_ Browser) SendsActionOnArrowKeys() bool {
	rv := objc.Call[bool](b_, objc.Sel("sendsActionOnArrowKeys"))
	return rv
}

// A Boolean that indicates whether pressing an arrow key causes an action message to be sent. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbrowser/1407614-sendsactiononarrowkeys?language=objc
func (b_ Browser) SetSendsActionOnArrowKeys(value bool) {
	objc.Call[objc.Void](b_, objc.Sel("setSendsActionOnArrowKeys:"), value)
}

// The row number of the cell that the user clicked to display a context menu. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbrowser/1407671-clickedrow?language=objc
func (b_ Browser) ClickedRow() int {
	rv := objc.Call[int](b_, objc.Sel("clickedRow"))
	return rv
}

// A Boolean that indicates whether the user can select multiple items. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbrowser/1407509-allowsmultipleselection?language=objc
func (b_ Browser) AllowsMultipleSelection() bool {
	rv := objc.Call[bool](b_, objc.Sel("allowsMultipleSelection"))
	return rv
}

// A Boolean that indicates whether the user can select multiple items. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbrowser/1407509-allowsmultipleselection?language=objc
func (b_ Browser) SetAllowsMultipleSelection(value bool) {
	objc.Call[objc.Void](b_, objc.Sel("setAllowsMultipleSelection:"), value)
}

// The index of the last column with a selected item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbrowser/1407568-selectedcolumn?language=objc
func (b_ Browser) SelectedColumn() int {
	rv := objc.Call[int](b_, objc.Sel("selectedColumn"))
	return rv
}

// The height of the column titles for the browser. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbrowser/1407709-titleheight?language=objc
func (b_ Browser) TitleHeight() float64 {
	rv := objc.Call[float64](b_, objc.Sel("titleHeight"))
	return rv
}

// The path separator. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbrowser/1407688-pathseparator?language=objc
func (b_ Browser) PathSeparator() string {
	rv := objc.Call[string](b_, objc.Sel("pathSeparator"))
	return rv
}

// The path separator. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbrowser/1407688-pathseparator?language=objc
func (b_ Browser) SetPathSeparator(value string) {
	objc.Call[objc.Void](b_, objc.Sel("setPathSeparator:"), value)
}

// A Boolean that indicates whether columns are separated by bezeled borders. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbrowser/1407653-separatescolumns?language=objc
func (b_ Browser) SeparatesColumns() bool {
	rv := objc.Call[bool](b_, objc.Sel("separatesColumns"))
	return rv
}

// A Boolean that indicates whether columns are separated by bezeled borders. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbrowser/1407653-separatescolumns?language=objc
func (b_ Browser) SetSeparatesColumns(value bool) {
	objc.Call[objc.Void](b_, objc.Sel("setSeparatesColumns:"), value)
}

// The maximum number of visible columns. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbrowser/1407749-maxvisiblecolumns?language=objc
func (b_ Browser) MaxVisibleColumns() int {
	rv := objc.Call[int](b_, objc.Sel("maxVisibleColumns"))
	return rv
}

// The maximum number of visible columns. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbrowser/1407749-maxvisiblecolumns?language=objc
func (b_ Browser) SetMaxVisibleColumns(value int) {
	objc.Call[objc.Void](b_, objc.Sel("setMaxVisibleColumns:"), value)
}

// The name used to automatically save the browser’s column configuration. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbrowser/1407650-columnsautosavename?language=objc
func (b_ Browser) ColumnsAutosaveName() BrowserColumnsAutosaveName {
	rv := objc.Call[BrowserColumnsAutosaveName](b_, objc.Sel("columnsAutosaveName"))
	return rv
}

// The name used to automatically save the browser’s column configuration. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbrowser/1407650-columnsautosavename?language=objc
func (b_ Browser) SetColumnsAutosaveName(value BrowserColumnsAutosaveName) {
	objc.Call[objc.Void](b_, objc.Sel("setColumnsAutosaveName:"), value)
}

// Returns the NSBrowserCell class. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbrowser/1407505-cellclass?language=objc
func (bc _BrowserClass) CellClass() objc.Class {
	rv := objc.Call[objc.Class](bc, objc.Sel("cellClass"))
	return rv
}

// Returns the NSBrowserCell class. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbrowser/1407505-cellclass?language=objc
func Browser_CellClass() objc.Class {
	return BrowserClass.CellClass()
}

// The index of the last column loaded. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbrowser/1407715-lastcolumn?language=objc
func (b_ Browser) LastColumn() int {
	rv := objc.Call[int](b_, objc.Sel("lastColumn"))
	return rv
}

// The index of the last column loaded. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbrowser/1407715-lastcolumn?language=objc
func (b_ Browser) SetLastColumn(value int) {
	objc.Call[objc.Void](b_, objc.Sel("setLastColumn:"), value)
}

// The minimum column width, in pixels. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbrowser/1407513-mincolumnwidth?language=objc
func (b_ Browser) MinColumnWidth() float64 {
	rv := objc.Call[float64](b_, objc.Sel("minColumnWidth"))
	return rv
}

// The minimum column width, in pixels. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbrowser/1407513-mincolumnwidth?language=objc
func (b_ Browser) SetMinColumnWidth(value float64) {
	objc.Call[objc.Void](b_, objc.Sel("setMinColumnWidth:"), value)
}

// A Boolean that indicates whether a column takes its title from the selected cell in the previous column. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbrowser/1407564-takestitlefrompreviouscolumn?language=objc
func (b_ Browser) TakesTitleFromPreviousColumn() bool {
	rv := objc.Call[bool](b_, objc.Sel("takesTitleFromPreviousColumn"))
	return rv
}

// A Boolean that indicates whether a column takes its title from the selected cell in the previous column. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbrowser/1407564-takestitlefrompreviouscolumn?language=objc
func (b_ Browser) SetTakesTitleFromPreviousColumn(value bool) {
	objc.Call[objc.Void](b_, objc.Sel("setTakesTitleFromPreviousColumn:"), value)
}

// A Boolean that indicates whether the browser has a horizontal scroller. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbrowser/1407551-hashorizontalscroller?language=objc
func (b_ Browser) HasHorizontalScroller() bool {
	rv := objc.Call[bool](b_, objc.Sel("hasHorizontalScroller"))
	return rv
}

// A Boolean that indicates whether the browser has a horizontal scroller. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbrowser/1407551-hashorizontalscroller?language=objc
func (b_ Browser) SetHasHorizontalScroller(value bool) {
	objc.Call[objc.Void](b_, objc.Sel("setHasHorizontalScroller:"), value)
}
