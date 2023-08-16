// AUTO-GENERATED CODE, DO NOT MODIFY

package coredata

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [ManagedObjectModel] class.
var ManagedObjectModelClass = _ManagedObjectModelClass{objc.GetClass("NSManagedObjectModel")}

type _ManagedObjectModelClass struct {
	objc.Class
}

// An interface definition for the [ManagedObjectModel] class.
type IManagedObjectModel interface {
	objc.IObject
	FetchRequestFromTemplateWithNameSubstitutionVariables(name string, variables map[string]objc.IObject) FetchRequest
	EntitiesForConfiguration(configuration string) []EntityDescription
	SetFetchRequestTemplateForName(fetchRequestTemplate IFetchRequest, name string)
	IsConfigurationCompatibleWithStoreMetadata(configuration string, metadata map[string]objc.IObject) bool
	FetchRequestTemplateForName(name string) FetchRequest
	SetEntitiesForConfiguration(entities []IEntityDescription, configuration string)
	VersionIdentifiers() foundation.Set
	SetVersionIdentifiers(value foundation.ISet)
	LocalizationDictionary() map[string]string
	SetLocalizationDictionary(value map[string]string)
	FetchRequestTemplatesByName() map[string]FetchRequest
	Entities() []EntityDescription
	SetEntities(value []IEntityDescription)
	EntitiesByName() map[string]EntityDescription
	EntityVersionHashesByName() map[string][]byte
	Configurations() []string
}

// A programmatic representation of the .xcdatamodeld file describing your objects. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobjectmodel?language=objc
type ManagedObjectModel struct {
	objc.Object
}

func ManagedObjectModelFrom(ptr unsafe.Pointer) ManagedObjectModel {
	return ManagedObjectModel{
		Object: objc.ObjectFrom(ptr),
	}
}

func (m_ ManagedObjectModel) InitWithContentsOfURL(url foundation.IURL) ManagedObjectModel {
	rv := objc.Call[ManagedObjectModel](m_, objc.Sel("initWithContentsOfURL:"), objc.Ptr(url))
	return rv
}

// Initializes the managed object model using the model file at the specified URL. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobjectmodel/1506225-initwithcontentsofurl?language=objc
func ManagedObjectModel_InitWithContentsOfURL(url foundation.IURL) ManagedObjectModel {
	return ManagedObjectModelClass.Alloc().InitWithContentsOfURL(url)
}

func (m_ ManagedObjectModel) Init() ManagedObjectModel {
	rv := objc.Call[ManagedObjectModel](m_, objc.Sel("init"))
	return rv
}

func (mc _ManagedObjectModelClass) Alloc() ManagedObjectModel {
	rv := objc.Call[ManagedObjectModel](mc, objc.Sel("alloc"))
	return rv
}

func ManagedObjectModel_Alloc() ManagedObjectModel {
	return ManagedObjectModelClass.Alloc()
}

func (mc _ManagedObjectModelClass) New() ManagedObjectModel {
	rv := objc.Call[ManagedObjectModel](mc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewManagedObjectModel() ManagedObjectModel {
	return ManagedObjectModelClass.New()
}

// Returns a copy of the fetch request template with the variables substituted by values from the substitutions dictionary. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobjectmodel/1506422-fetchrequestfromtemplatewithname?language=objc
func (m_ ManagedObjectModel) FetchRequestFromTemplateWithNameSubstitutionVariables(name string, variables map[string]objc.IObject) FetchRequest {
	rv := objc.Call[FetchRequest](m_, objc.Sel("fetchRequestFromTemplateWithName:substitutionVariables:"), name, variables)
	return rv
}

// Returns the entities of the model for a specified configuration. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobjectmodel/1506693-entitiesforconfiguration?language=objc
func (m_ ManagedObjectModel) EntitiesForConfiguration(configuration string) []EntityDescription {
	rv := objc.Call[[]EntityDescription](m_, objc.Sel("entitiesForConfiguration:"), configuration)
	return rv
}

// Associates the specified fetch request with the receiver using the given name. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobjectmodel/1506695-setfetchrequesttemplate?language=objc
func (m_ ManagedObjectModel) SetFetchRequestTemplateForName(fetchRequestTemplate IFetchRequest, name string) {
	objc.Call[objc.Void](m_, objc.Sel("setFetchRequestTemplate:forName:"), objc.Ptr(fetchRequestTemplate), name)
}

// Returns a Boolean value that indicates whether a given configuration in the model is compatible with given metadata from a persistent store. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobjectmodel/1506940-isconfiguration?language=objc
func (m_ ManagedObjectModel) IsConfigurationCompatibleWithStoreMetadata(configuration string, metadata map[string]objc.IObject) bool {
	rv := objc.Call[bool](m_, objc.Sel("isConfiguration:compatibleWithStoreMetadata:"), configuration, metadata)
	return rv
}

// Returns the fetch request with a specified name. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobjectmodel/1506369-fetchrequesttemplateforname?language=objc
func (m_ ManagedObjectModel) FetchRequestTemplateForName(name string) FetchRequest {
	rv := objc.Call[FetchRequest](m_, objc.Sel("fetchRequestTemplateForName:"), name)
	return rv
}

// Associates the specified entities with the model using the given configuration name. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobjectmodel/1506287-setentities?language=objc
func (m_ ManagedObjectModel) SetEntitiesForConfiguration(entities []IEntityDescription, configuration string) {
	objc.Call[objc.Void](m_, objc.Sel("setEntities:forConfiguration:"), entities, configuration)
}

// Returns, for the version information in given metadata, a model merged from a given array of models. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobjectmodel/1506856-modelbymergingmodels?language=objc
func (mc _ManagedObjectModelClass) ModelByMergingModelsForStoreMetadata(models []IManagedObjectModel, metadata map[string]objc.IObject) ManagedObjectModel {
	rv := objc.Call[ManagedObjectModel](mc, objc.Sel("modelByMergingModels:forStoreMetadata:"), models, metadata)
	return rv
}

// Returns, for the version information in given metadata, a model merged from a given array of models. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobjectmodel/1506856-modelbymergingmodels?language=objc
func ManagedObjectModel_ModelByMergingModelsForStoreMetadata(models []IManagedObjectModel, metadata map[string]objc.IObject) ManagedObjectModel {
	return ManagedObjectModelClass.ModelByMergingModelsForStoreMetadata(models, metadata)
}

// Returns a model created by merging all the models found in given bundles. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobjectmodel/1506704-mergedmodelfrombundles?language=objc
func (mc _ManagedObjectModelClass) MergedModelFromBundles(bundles []foundation.IBundle) ManagedObjectModel {
	rv := objc.Call[ManagedObjectModel](mc, objc.Sel("mergedModelFromBundles:"), bundles)
	return rv
}

// Returns a model created by merging all the models found in given bundles. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobjectmodel/1506704-mergedmodelfrombundles?language=objc
func ManagedObjectModel_MergedModelFromBundles(bundles []foundation.IBundle) ManagedObjectModel {
	return ManagedObjectModelClass.MergedModelFromBundles(bundles)
}

// The set of developer-defined version identifiers for the object model. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobjectmodel/1506268-versionidentifiers?language=objc
func (m_ ManagedObjectModel) VersionIdentifiers() foundation.Set {
	rv := objc.Call[foundation.Set](m_, objc.Sel("versionIdentifiers"))
	return rv
}

// The set of developer-defined version identifiers for the object model. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobjectmodel/1506268-versionidentifiers?language=objc
func (m_ ManagedObjectModel) SetVersionIdentifiers(value foundation.ISet) {
	objc.Call[objc.Void](m_, objc.Sel("setVersionIdentifiers:"), objc.Ptr(value))
}

// The localization dictionary of the model. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobjectmodel/1506846-localizationdictionary?language=objc
func (m_ ManagedObjectModel) LocalizationDictionary() map[string]string {
	rv := objc.Call[map[string]string](m_, objc.Sel("localizationDictionary"))
	return rv
}

// The localization dictionary of the model. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobjectmodel/1506846-localizationdictionary?language=objc
func (m_ ManagedObjectModel) SetLocalizationDictionary(value map[string]string) {
	objc.Call[objc.Void](m_, objc.Sel("setLocalizationDictionary:"), value)
}

// A dictionary of the receiver’s fetch request templates, keyed by name. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobjectmodel/1506580-fetchrequesttemplatesbyname?language=objc
func (m_ ManagedObjectModel) FetchRequestTemplatesByName() map[string]FetchRequest {
	rv := objc.Call[map[string]FetchRequest](m_, objc.Sel("fetchRequestTemplatesByName"))
	return rv
}

// The entities in the model. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobjectmodel/1506318-entities?language=objc
func (m_ ManagedObjectModel) Entities() []EntityDescription {
	rv := objc.Call[[]EntityDescription](m_, objc.Sel("entities"))
	return rv
}

// The entities in the model. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobjectmodel/1506318-entities?language=objc
func (m_ ManagedObjectModel) SetEntities(value []IEntityDescription) {
	objc.Call[objc.Void](m_, objc.Sel("setEntities:"), value)
}

// The entities of the model, keyed by name. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobjectmodel/1506203-entitiesbyname?language=objc
func (m_ ManagedObjectModel) EntitiesByName() map[string]EntityDescription {
	rv := objc.Call[map[string]EntityDescription](m_, objc.Sel("entitiesByName"))
	return rv
}

// The dictionary of the model’s entity names and their corresponding version hashes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobjectmodel/1506992-entityversionhashesbyname?language=objc
func (m_ ManagedObjectModel) EntityVersionHashesByName() map[string][]byte {
	rv := objc.Call[map[string][]byte](m_, objc.Sel("entityVersionHashesByName"))
	return rv
}

// All the available configuration names of the model. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobjectmodel/1506804-configurations?language=objc
func (m_ ManagedObjectModel) Configurations() []string {
	rv := objc.Call[[]string](m_, objc.Sel("configurations"))
	return rv
}
