// AUTO-GENERATED CODE, DO NOT MODIFY

package cloudkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [ModifyRecordZonesOperation] class.
var ModifyRecordZonesOperationClass = _ModifyRecordZonesOperationClass{objc.GetClass("CKModifyRecordZonesOperation")}

type _ModifyRecordZonesOperationClass struct {
	objc.Class
}

// An interface definition for the [ModifyRecordZonesOperation] class.
type IModifyRecordZonesOperation interface {
	IDatabaseOperation
	RecordZonesToSave() []RecordZone
	SetRecordZonesToSave(value []IRecordZone)
	PerRecordZoneDeleteBlock() func(recordZoneID RecordZoneID, error foundation.Error)
	SetPerRecordZoneDeleteBlock(value func(recordZoneID RecordZoneID, error foundation.Error))
	RecordZoneIDsToDelete() []RecordZoneID
	SetRecordZoneIDsToDelete(value []IRecordZoneID)
	PerRecordZoneSaveBlock() func(recordZoneID RecordZoneID, recordZone RecordZone, error foundation.Error)
	SetPerRecordZoneSaveBlock(value func(recordZoneID RecordZoneID, recordZone RecordZone, error foundation.Error))
	ModifyRecordZonesCompletionBlock() func(savedRecordZones []RecordZone, deletedRecordZoneIDs []RecordZoneID, operationError foundation.Error)
	SetModifyRecordZonesCompletionBlock(value func(savedRecordZones []RecordZone, deletedRecordZoneIDs []RecordZoneID, operationError foundation.Error))
}

// An operation that modifies one or more record zones. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckmodifyrecordzonesoperation?language=objc
type ModifyRecordZonesOperation struct {
	DatabaseOperation
}

func ModifyRecordZonesOperationFrom(ptr unsafe.Pointer) ModifyRecordZonesOperation {
	return ModifyRecordZonesOperation{
		DatabaseOperation: DatabaseOperationFrom(ptr),
	}
}

func (m_ ModifyRecordZonesOperation) InitWithRecordZonesToSaveRecordZoneIDsToDelete(recordZonesToSave []IRecordZone, recordZoneIDsToDelete []IRecordZoneID) ModifyRecordZonesOperation {
	rv := objc.Call[ModifyRecordZonesOperation](m_, objc.Sel("initWithRecordZonesToSave:recordZoneIDsToDelete:"), recordZonesToSave, recordZoneIDsToDelete)
	return rv
}

// Creates an operation for modifying the specified record zones. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckmodifyrecordzonesoperation/1415167-initwithrecordzonestosave?language=objc
func ModifyRecordZonesOperation_InitWithRecordZonesToSaveRecordZoneIDsToDelete(recordZonesToSave []IRecordZone, recordZoneIDsToDelete []IRecordZoneID) ModifyRecordZonesOperation {
	return ModifyRecordZonesOperationClass.Alloc().InitWithRecordZonesToSaveRecordZoneIDsToDelete(recordZonesToSave, recordZoneIDsToDelete)
}

func (m_ ModifyRecordZonesOperation) Init() ModifyRecordZonesOperation {
	rv := objc.Call[ModifyRecordZonesOperation](m_, objc.Sel("init"))
	return rv
}

func (mc _ModifyRecordZonesOperationClass) Alloc() ModifyRecordZonesOperation {
	rv := objc.Call[ModifyRecordZonesOperation](mc, objc.Sel("alloc"))
	return rv
}

func ModifyRecordZonesOperation_Alloc() ModifyRecordZonesOperation {
	return ModifyRecordZonesOperationClass.Alloc()
}

func (mc _ModifyRecordZonesOperationClass) New() ModifyRecordZonesOperation {
	rv := objc.Call[ModifyRecordZonesOperation](mc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewModifyRecordZonesOperation() ModifyRecordZonesOperation {
	return ModifyRecordZonesOperationClass.New()
}

// The record zones to save to the database. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckmodifyrecordzonesoperation/1415171-recordzonestosave?language=objc
func (m_ ModifyRecordZonesOperation) RecordZonesToSave() []RecordZone {
	rv := objc.Call[[]RecordZone](m_, objc.Sel("recordZonesToSave"))
	return rv
}

// The record zones to save to the database. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckmodifyrecordzonesoperation/1415171-recordzonestosave?language=objc
func (m_ ModifyRecordZonesOperation) SetRecordZonesToSave(value []IRecordZone) {
	objc.Call[objc.Void](m_, objc.Sel("setRecordZonesToSave:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckmodifyrecordzonesoperation/3793702-perrecordzonedeleteblock?language=objc
func (m_ ModifyRecordZonesOperation) PerRecordZoneDeleteBlock() func(recordZoneID RecordZoneID, error foundation.Error) {
	rv := objc.Call[func(recordZoneID RecordZoneID, error foundation.Error)](m_, objc.Sel("perRecordZoneDeleteBlock"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckmodifyrecordzonesoperation/3793702-perrecordzonedeleteblock?language=objc
func (m_ ModifyRecordZonesOperation) SetPerRecordZoneDeleteBlock(value func(recordZoneID RecordZoneID, error foundation.Error)) {
	objc.Call[objc.Void](m_, objc.Sel("setPerRecordZoneDeleteBlock:"), value)
}

// The IDs of the record zones to delete permanently from the database. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckmodifyrecordzonesoperation/1415173-recordzoneidstodelete?language=objc
func (m_ ModifyRecordZonesOperation) RecordZoneIDsToDelete() []RecordZoneID {
	rv := objc.Call[[]RecordZoneID](m_, objc.Sel("recordZoneIDsToDelete"))
	return rv
}

// The IDs of the record zones to delete permanently from the database. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckmodifyrecordzonesoperation/1415173-recordzoneidstodelete?language=objc
func (m_ ModifyRecordZonesOperation) SetRecordZoneIDsToDelete(value []IRecordZoneID) {
	objc.Call[objc.Void](m_, objc.Sel("setRecordZoneIDsToDelete:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckmodifyrecordzonesoperation/3793703-perrecordzonesaveblock?language=objc
func (m_ ModifyRecordZonesOperation) PerRecordZoneSaveBlock() func(recordZoneID RecordZoneID, recordZone RecordZone, error foundation.Error) {
	rv := objc.Call[func(recordZoneID RecordZoneID, recordZone RecordZone, error foundation.Error)](m_, objc.Sel("perRecordZoneSaveBlock"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckmodifyrecordzonesoperation/3793703-perrecordzonesaveblock?language=objc
func (m_ ModifyRecordZonesOperation) SetPerRecordZoneSaveBlock(value func(recordZoneID RecordZoneID, recordZone RecordZone, error foundation.Error)) {
	objc.Call[objc.Void](m_, objc.Sel("setPerRecordZoneSaveBlock:"), value)
}

// The block to execute after CloudKit modifies all of the record zones. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckmodifyrecordzonesoperation/1415164-modifyrecordzonescompletionblock?language=objc
func (m_ ModifyRecordZonesOperation) ModifyRecordZonesCompletionBlock() func(savedRecordZones []RecordZone, deletedRecordZoneIDs []RecordZoneID, operationError foundation.Error) {
	rv := objc.Call[func(savedRecordZones []RecordZone, deletedRecordZoneIDs []RecordZoneID, operationError foundation.Error)](m_, objc.Sel("modifyRecordZonesCompletionBlock"))
	return rv
}

// The block to execute after CloudKit modifies all of the record zones. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckmodifyrecordzonesoperation/1415164-modifyrecordzonescompletionblock?language=objc
func (m_ ModifyRecordZonesOperation) SetModifyRecordZonesCompletionBlock(value func(savedRecordZones []RecordZone, deletedRecordZoneIDs []RecordZoneID, operationError foundation.Error)) {
	objc.Call[objc.Void](m_, objc.Sel("setModifyRecordZonesCompletionBlock:"), value)
}
