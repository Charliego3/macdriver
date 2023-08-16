// AUTO-GENERATED CODE, DO NOT MODIFY

package coredata

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [PersistentStoreAsynchronousResult] class.
var PersistentStoreAsynchronousResultClass = _PersistentStoreAsynchronousResultClass{objc.GetClass("NSPersistentStoreAsynchronousResult")}

type _PersistentStoreAsynchronousResultClass struct {
	objc.Class
}

// An interface definition for the [PersistentStoreAsynchronousResult] class.
type IPersistentStoreAsynchronousResult interface {
	IPersistentStoreResult
	Cancel()
	OperationError() foundation.Error
	Progress() foundation.Progress
	ManagedObjectContext() ManagedObjectContext
}

// A concrete class used to represent the results of an asynchronous request. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistentstoreasynchronousresult?language=objc
type PersistentStoreAsynchronousResult struct {
	PersistentStoreResult
}

func PersistentStoreAsynchronousResultFrom(ptr unsafe.Pointer) PersistentStoreAsynchronousResult {
	return PersistentStoreAsynchronousResult{
		PersistentStoreResult: PersistentStoreResultFrom(ptr),
	}
}

func (pc _PersistentStoreAsynchronousResultClass) Alloc() PersistentStoreAsynchronousResult {
	rv := objc.Call[PersistentStoreAsynchronousResult](pc, objc.Sel("alloc"))
	return rv
}

func PersistentStoreAsynchronousResult_Alloc() PersistentStoreAsynchronousResult {
	return PersistentStoreAsynchronousResultClass.Alloc()
}

func (pc _PersistentStoreAsynchronousResultClass) New() PersistentStoreAsynchronousResult {
	rv := objc.Call[PersistentStoreAsynchronousResult](pc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewPersistentStoreAsynchronousResult() PersistentStoreAsynchronousResult {
	return PersistentStoreAsynchronousResultClass.New()
}

func (p_ PersistentStoreAsynchronousResult) Init() PersistentStoreAsynchronousResult {
	rv := objc.Call[PersistentStoreAsynchronousResult](p_, objc.Sel("init"))
	return rv
}

// Cancels the asynchronous fetch request. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistentstoreasynchronousresult/1404924-cancel?language=objc
func (p_ PersistentStoreAsynchronousResult) Cancel() {
	objc.Call[objc.Void](p_, objc.Sel("cancel"))
}

// An error that contains details if the asynchronous fetch request fails. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistentstoreasynchronousresult/1404904-operationerror?language=objc
func (p_ PersistentStoreAsynchronousResult) OperationError() foundation.Error {
	rv := objc.Call[foundation.Error](p_, objc.Sel("operationError"))
	return rv
}

// An object that reports progress for the asynchronous fetch request. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistentstoreasynchronousresult/1404920-progress?language=objc
func (p_ PersistentStoreAsynchronousResult) Progress() foundation.Progress {
	rv := objc.Call[foundation.Progress](p_, objc.Sel("progress"))
	return rv
}

// The managed object context for the result. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistentstoreasynchronousresult/1404916-managedobjectcontext?language=objc
func (p_ PersistentStoreAsynchronousResult) ManagedObjectContext() ManagedObjectContext {
	rv := objc.Call[ManagedObjectContext](p_, objc.Sel("managedObjectContext"))
	return rv
}
