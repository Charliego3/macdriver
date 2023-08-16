// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [Pasteboard] class.
var PasteboardClass = _PasteboardClass{objc.GetClass("NSPasteboard")}

type _PasteboardClass struct {
	objc.Class
}

// An interface definition for the [Pasteboard] class.
type IPasteboard interface {
	objc.IObject
	StringForType(dataType PasteboardType) string
	ReadFileContentsTypeToFile(type_ PasteboardType, filename string) string
	ReadObjectsForClassesOptions(classArray []objc.IClass, options map[PasteboardReadingOptionKey]objc.IObject) []objc.Object
	IndexOfPasteboardItem(pasteboardItem IPasteboardItem) uint
	AddTypesOwner(newTypes []PasteboardType, newOwner objc.IObject) int
	WriteFileContents(filename string) bool
	ReadFileWrapper() foundation.FileWrapper
	WriteObjects(objects []PPasteboardWriting) bool
	ReleaseGlobally()
	WriteFileWrapper(wrapper foundation.IFileWrapper) bool
	SetPropertyListForType(plist objc.IObject, dataType PasteboardType) bool
	SetStringForType(string_ string, dataType PasteboardType) bool
	CanReadObjectForClassesOptions(classArray []objc.IClass, options map[PasteboardReadingOptionKey]objc.IObject) bool
	DataForType(dataType PasteboardType) []byte
	ClearContents() int
	SetDataForType(data []byte, dataType PasteboardType) bool
	PrepareForNewContentsWithOptions(options PasteboardContentsOptions) int
	AvailableTypeFromArray(types []PasteboardType) PasteboardType
	DeclareTypesOwner(newTypes []PasteboardType, newOwner objc.IObject) int
	CanReadItemWithDataConformingToTypes(types []string) bool
	PropertyListForType(dataType PasteboardType) objc.Object
	Name() PasteboardName
	Types() []PasteboardType
	ChangeCount() int
	PasteboardItems() []PasteboardItem
}

// An object that transfers data to and from the pasteboard server. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspasteboard?language=objc
type Pasteboard struct {
	objc.Object
}

func PasteboardFrom(ptr unsafe.Pointer) Pasteboard {
	return Pasteboard{
		Object: objc.ObjectFrom(ptr),
	}
}

func (pc _PasteboardClass) Alloc() Pasteboard {
	rv := objc.Call[Pasteboard](pc, objc.Sel("alloc"))
	return rv
}

func Pasteboard_Alloc() Pasteboard {
	return PasteboardClass.Alloc()
}

func (pc _PasteboardClass) New() Pasteboard {
	rv := objc.Call[Pasteboard](pc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewPasteboard() Pasteboard {
	return PasteboardClass.New()
}

func (p_ Pasteboard) Init() Pasteboard {
	rv := objc.Call[Pasteboard](p_, objc.Sel("init"))
	return rv
}

// Returns a concatenation of the strings for the specified type from all the items in the receiver that contain the type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspasteboard/1533566-stringfortype?language=objc
func (p_ Pasteboard) StringForType(dataType PasteboardType) string {
	rv := objc.Call[string](p_, objc.Sel("stringForType:"), dataType)
	return rv
}

// Returns the pasteboard with the specified name. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspasteboard/1531026-pasteboardwithname?language=objc
func (pc _PasteboardClass) PasteboardWithName(name PasteboardName) Pasteboard {
	rv := objc.Call[Pasteboard](pc, objc.Sel("pasteboardWithName:"), name)
	return rv
}

// Returns the pasteboard with the specified name. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspasteboard/1531026-pasteboardwithname?language=objc
func Pasteboard_PasteboardWithName(name PasteboardName) Pasteboard {
	return PasteboardClass.PasteboardWithName(name)
}

// Reads data representing a file’s contents from the receiver and writes it to the specified file. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspasteboard/1533575-readfilecontentstype?language=objc
func (p_ Pasteboard) ReadFileContentsTypeToFile(type_ PasteboardType, filename string) string {
	rv := objc.Call[string](p_, objc.Sel("readFileContentsType:toFile:"), type_, filename)
	return rv
}

// Reads from the receiver objects that best match the specified array of classes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspasteboard/1524454-readobjectsforclasses?language=objc
func (p_ Pasteboard) ReadObjectsForClassesOptions(classArray []objc.IClass, options map[PasteboardReadingOptionKey]objc.IObject) []objc.Object {
	rv := objc.Call[[]objc.Object](p_, objc.Sel("readObjectsForClasses:options:"), classArray, options)
	return rv
}

// Returns the index of the specified pasteboard item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspasteboard/1525176-indexofpasteboarditem?language=objc
func (p_ Pasteboard) IndexOfPasteboardItem(pasteboardItem IPasteboardItem) uint {
	rv := objc.Call[uint](p_, objc.Sel("indexOfPasteboardItem:"), objc.Ptr(pasteboardItem))
	return rv
}

// Adds promises for the specified types to the first pasteboard item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspasteboard/1533580-addtypes?language=objc
func (p_ Pasteboard) AddTypesOwner(newTypes []PasteboardType, newOwner objc.IObject) int {
	rv := objc.Call[int](p_, objc.Sel("addTypes:owner:"), newTypes, newOwner)
	return rv
}

// Creates a new pasteboard object that supplies the specified data in as many types as possible based on the available filter services. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspasteboard/1530706-pasteboardbyfilteringdata?language=objc
func (pc _PasteboardClass) PasteboardByFilteringDataOfType(data []byte, type_ PasteboardType) Pasteboard {
	rv := objc.Call[Pasteboard](pc, objc.Sel("pasteboardByFilteringData:ofType:"), data, type_)
	return rv
}

// Creates a new pasteboard object that supplies the specified data in as many types as possible based on the available filter services. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspasteboard/1530706-pasteboardbyfilteringdata?language=objc
func Pasteboard_PasteboardByFilteringDataOfType(data []byte, type_ PasteboardType) Pasteboard {
	return PasteboardClass.PasteboardByFilteringDataOfType(data, type_)
}

// Writes the contents of the specified file to the pasteboard. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspasteboard/1531224-writefilecontents?language=objc
func (p_ Pasteboard) WriteFileContents(filename string) bool {
	rv := objc.Call[bool](p_, objc.Sel("writeFileContents:"), filename)
	return rv
}

// Reads data representing a file’s contents from the receiver and returns it as a file wrapper. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspasteboard/1524779-readfilewrapper?language=objc
func (p_ Pasteboard) ReadFileWrapper() foundation.FileWrapper {
	rv := objc.Call[foundation.FileWrapper](p_, objc.Sel("readFileWrapper"))
	return rv
}

// Writes an array of objects to the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspasteboard/1525945-writeobjects?language=objc
func (p_ Pasteboard) WriteObjects(objects []PPasteboardWriting) bool {
	rv := objc.Call[bool](p_, objc.Sel("writeObjects:"), objects)
	return rv
}

// Releases the receiver’s resources in the pasteboard server. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspasteboard/1527044-releaseglobally?language=objc
func (p_ Pasteboard) ReleaseGlobally() {
	objc.Call[objc.Void](p_, objc.Sel("releaseGlobally"))
}

// Writes the serialized contents of the specified file wrapper to the pasteboard. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspasteboard/1527279-writefilewrapper?language=objc
func (p_ Pasteboard) WriteFileWrapper(wrapper foundation.IFileWrapper) bool {
	rv := objc.Call[bool](p_, objc.Sel("writeFileWrapper:"), objc.Ptr(wrapper))
	return rv
}

// Sets the given property list as the representation for the specified type for the first item on the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspasteboard/1530774-setpropertylist?language=objc
func (p_ Pasteboard) SetPropertyListForType(plist objc.IObject, dataType PasteboardType) bool {
	rv := objc.Call[bool](p_, objc.Sel("setPropertyList:forType:"), plist, dataType)
	return rv
}

// Sets the given string as the representation for the specified type for the first item on the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspasteboard/1528225-setstring?language=objc
func (p_ Pasteboard) SetStringForType(string_ string, dataType PasteboardType) bool {
	rv := objc.Call[bool](p_, objc.Sel("setString:forType:"), string_, dataType)
	return rv
}

// Creates and returns a new pasteboard with a name that is guaranteed to be unique with respect to other pasteboards in the system. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspasteboard/1528936-pasteboardwithuniquename?language=objc
func (pc _PasteboardClass) PasteboardWithUniqueName() Pasteboard {
	rv := objc.Call[Pasteboard](pc, objc.Sel("pasteboardWithUniqueName"))
	return rv
}

// Creates and returns a new pasteboard with a name that is guaranteed to be unique with respect to other pasteboards in the system. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspasteboard/1528936-pasteboardwithuniquename?language=objc
func Pasteboard_PasteboardWithUniqueName() Pasteboard {
	return PasteboardClass.PasteboardWithUniqueName()
}

// Creates a new pasteboard object that supplies the specified file in as many types as possible based on the available filter services. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspasteboard/1532744-pasteboardbyfilteringfile?language=objc
func (pc _PasteboardClass) PasteboardByFilteringFile(filename string) Pasteboard {
	rv := objc.Call[Pasteboard](pc, objc.Sel("pasteboardByFilteringFile:"), filename)
	return rv
}

// Creates a new pasteboard object that supplies the specified file in as many types as possible based on the available filter services. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspasteboard/1532744-pasteboardbyfilteringfile?language=objc
func Pasteboard_PasteboardByFilteringFile(filename string) Pasteboard {
	return PasteboardClass.PasteboardByFilteringFile(filename)
}

// Returns a Boolean value that indicates whether the receiver contains any items that can be represented as an instance of any class in a given array. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspasteboard/1533360-canreadobjectforclasses?language=objc
func (p_ Pasteboard) CanReadObjectForClassesOptions(classArray []objc.IClass, options map[PasteboardReadingOptionKey]objc.IObject) bool {
	rv := objc.Call[bool](p_, objc.Sel("canReadObjectForClasses:options:"), classArray, options)
	return rv
}

// Returns the data for the specified type from the first item in the receiver that contains the type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspasteboard/1531810-datafortype?language=objc
func (p_ Pasteboard) DataForType(dataType PasteboardType) []byte {
	rv := objc.Call[[]byte](p_, objc.Sel("dataForType:"), dataType)
	return rv
}

// Clears the existing contents of the pasteboard. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspasteboard/1533599-clearcontents?language=objc
func (p_ Pasteboard) ClearContents() int {
	rv := objc.Call[int](p_, objc.Sel("clearContents"))
	return rv
}

// Returns the data types that can be converted to the specified type using the available filter services. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspasteboard/1533612-typesfilterableto?language=objc
func (pc _PasteboardClass) TypesFilterableTo(type_ PasteboardType) []PasteboardType {
	rv := objc.Call[[]PasteboardType](pc, objc.Sel("typesFilterableTo:"), type_)
	return rv
}

// Returns the data types that can be converted to the specified type using the available filter services. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspasteboard/1533612-typesfilterableto?language=objc
func Pasteboard_TypesFilterableTo(type_ PasteboardType) []PasteboardType {
	return PasteboardClass.TypesFilterableTo(type_)
}

// Sets the data as the representation for the specified type for the first item on the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspasteboard/1531214-setdata?language=objc
func (p_ Pasteboard) SetDataForType(data []byte, dataType PasteboardType) bool {
	rv := objc.Call[bool](p_, objc.Sel("setData:forType:"), data, dataType)
	return rv
}

// Prepares the pasteboard to receive new contents, removing the existing pasteboard contents. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspasteboard/2344960-preparefornewcontentswithoptions?language=objc
func (p_ Pasteboard) PrepareForNewContentsWithOptions(options PasteboardContentsOptions) int {
	rv := objc.Call[int](p_, objc.Sel("prepareForNewContentsWithOptions:"), options)
	return rv
}

// Creates a new pasteboard object that supplies the specified pasteboard data in as many types as possible based on the available filter services. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspasteboard/1530088-pasteboardbyfilteringtypesinpast?language=objc
func (pc _PasteboardClass) PasteboardByFilteringTypesInPasteboard(pboard IPasteboard) Pasteboard {
	rv := objc.Call[Pasteboard](pc, objc.Sel("pasteboardByFilteringTypesInPasteboard:"), objc.Ptr(pboard))
	return rv
}

// Creates a new pasteboard object that supplies the specified pasteboard data in as many types as possible based on the available filter services. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspasteboard/1530088-pasteboardbyfilteringtypesinpast?language=objc
func Pasteboard_PasteboardByFilteringTypesInPasteboard(pboard IPasteboard) Pasteboard {
	return PasteboardClass.PasteboardByFilteringTypesInPasteboard(pboard)
}

// Scans the specified types for a type that the receiver supports. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspasteboard/1526078-availabletypefromarray?language=objc
func (p_ Pasteboard) AvailableTypeFromArray(types []PasteboardType) PasteboardType {
	rv := objc.Call[PasteboardType](p_, objc.Sel("availableTypeFromArray:"), types)
	return rv
}

// Prepares the receiver for a change in its contents by declaring the new types of data it will contain and a new owner. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspasteboard/1533561-declaretypes?language=objc
func (p_ Pasteboard) DeclareTypesOwner(newTypes []PasteboardType, newOwner objc.IObject) int {
	rv := objc.Call[int](p_, objc.Sel("declareTypes:owner:"), newTypes, newOwner)
	return rv
}

// Returns a Boolean value that indicates whether the receiver contains any items that conform to the specified UTIs. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspasteboard/1533576-canreaditemwithdataconformingtot?language=objc
func (p_ Pasteboard) CanReadItemWithDataConformingToTypes(types []string) bool {
	rv := objc.Call[bool](p_, objc.Sel("canReadItemWithDataConformingToTypes:"), types)
	return rv
}

// Returns the property list for the specified type from the first item in the receiver that contains the type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspasteboard/1528588-propertylistfortype?language=objc
func (p_ Pasteboard) PropertyListForType(dataType PasteboardType) objc.Object {
	rv := objc.Call[objc.Object](p_, objc.Sel("propertyListForType:"), dataType)
	return rv
}

// The receiver’s name. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspasteboard/1529388-name?language=objc
func (p_ Pasteboard) Name() PasteboardName {
	rv := objc.Call[PasteboardName](p_, objc.Sel("name"))
	return rv
}

// An array of the receiver’s supported data types. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspasteboard/1529599-types?language=objc
func (p_ Pasteboard) Types() []PasteboardType {
	rv := objc.Call[[]PasteboardType](p_, objc.Sel("types"))
	return rv
}

// The receiver’s change count. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspasteboard/1533544-changecount?language=objc
func (p_ Pasteboard) ChangeCount() int {
	rv := objc.Call[int](p_, objc.Sel("changeCount"))
	return rv
}

// An array that contains all the items held by the pasteboard. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspasteboard/1529995-pasteboarditems?language=objc
func (p_ Pasteboard) PasteboardItems() []PasteboardItem {
	rv := objc.Call[[]PasteboardItem](p_, objc.Sel("pasteboardItems"))
	return rv
}

// The shared pasteboard object to use for general content. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspasteboard/1530091-generalpasteboard?language=objc
func (pc _PasteboardClass) GeneralPasteboard() Pasteboard {
	rv := objc.Call[Pasteboard](pc, objc.Sel("generalPasteboard"))
	return rv
}

// The shared pasteboard object to use for general content. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspasteboard/1530091-generalpasteboard?language=objc
func Pasteboard_GeneralPasteboard() Pasteboard {
	return PasteboardClass.GeneralPasteboard()
}
