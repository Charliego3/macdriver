// AUTO-GENERATED CODE, DO NOT MODIFY

package coreml

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [ModelDescription] class.
var ModelDescriptionClass = _ModelDescriptionClass{objc.GetClass("MLModelDescription")}

type _ModelDescriptionClass struct {
	objc.Class
}

// An interface definition for the [ModelDescription] class.
type IModelDescription interface {
	objc.IObject
	TrainingInputDescriptionsByName() map[string]FeatureDescription
	ParameterDescriptionsByKey() foundation.Dictionary
	IsUpdatable() bool
	Metadata() map[ModelMetadataKey]objc.Object
	PredictedFeatureName() string
	InputDescriptionsByName() map[string]FeatureDescription
	ClassLabels() []objc.Object
	PredictedProbabilitiesName() string
	OutputDescriptionsByName() map[string]FeatureDescription
}

// Information about a model, primarily the input and output format for each feature the model expects, and optional metadata. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlmodeldescription?language=objc
type ModelDescription struct {
	objc.Object
}

func ModelDescriptionFrom(ptr unsafe.Pointer) ModelDescription {
	return ModelDescription{
		Object: objc.ObjectFrom(ptr),
	}
}

func (mc _ModelDescriptionClass) Alloc() ModelDescription {
	rv := objc.Call[ModelDescription](mc, objc.Sel("alloc"))
	return rv
}

func ModelDescription_Alloc() ModelDescription {
	return ModelDescriptionClass.Alloc()
}

func (mc _ModelDescriptionClass) New() ModelDescription {
	rv := objc.Call[ModelDescription](mc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewModelDescription() ModelDescription {
	return ModelDescriptionClass.New()
}

func (m_ ModelDescription) Init() ModelDescription {
	rv := objc.Call[ModelDescription](m_, objc.Sel("init"))
	return rv
}

// A dictionary of the training input feature descriptions, which the model keys by the input’s name. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlmodeldescription/3180060-traininginputdescriptionsbyname?language=objc
func (m_ ModelDescription) TrainingInputDescriptionsByName() map[string]FeatureDescription {
	rv := objc.Call[map[string]FeatureDescription](m_, objc.Sel("trainingInputDescriptionsByName"))
	return rv
}

// A dictionary of the descriptions for the model’s parameters. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlmodeldescription/3333250-parameterdescriptionsbykey?language=objc
func (m_ ModelDescription) ParameterDescriptionsByKey() foundation.Dictionary {
	rv := objc.Call[foundation.Dictionary](m_, objc.Sel("parameterDescriptionsByKey"))
	return rv
}

// A Boolean value that indicates whether you can update the model with additional training. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlmodeldescription/3180059-isupdatable?language=objc
func (m_ ModelDescription) IsUpdatable() bool {
	rv := objc.Call[bool](m_, objc.Sel("isUpdatable"))
	return rv
}

// A dictionary of the model’s creation information, such as its description, author, version, and license. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlmodeldescription/2879386-metadata?language=objc
func (m_ ModelDescription) Metadata() map[ModelMetadataKey]objc.Object {
	rv := objc.Call[map[ModelMetadataKey]objc.Object](m_, objc.Sel("metadata"))
	return rv
}

// The name of the primary prediction feature output description. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlmodeldescription/2879390-predictedfeaturename?language=objc
func (m_ ModelDescription) PredictedFeatureName() string {
	rv := objc.Call[string](m_, objc.Sel("predictedFeatureName"))
	return rv
}

// A dictionary of input feature descriptions, which the model keys by the input’s name. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlmodeldescription/2879352-inputdescriptionsbyname?language=objc
func (m_ ModelDescription) InputDescriptionsByName() map[string]FeatureDescription {
	rv := objc.Call[map[string]FeatureDescription](m_, objc.Sel("inputDescriptionsByName"))
	return rv
}

// An array of labels, which can be either strings or a numbers, for classifier models. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlmodeldescription/3649657-classlabels?language=objc
func (m_ ModelDescription) ClassLabels() []objc.Object {
	rv := objc.Call[[]objc.Object](m_, objc.Sel("classLabels"))
	return rv
}

// The name of the feature output description for all probabilities of a prediction. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlmodeldescription/2879383-predictedprobabilitiesname?language=objc
func (m_ ModelDescription) PredictedProbabilitiesName() string {
	rv := objc.Call[string](m_, objc.Sel("predictedProbabilitiesName"))
	return rv
}

// A dictionary of output feature descriptions, which the model keys by the output’s name. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlmodeldescription/2879361-outputdescriptionsbyname?language=objc
func (m_ ModelDescription) OutputDescriptionsByName() map[string]FeatureDescription {
	rv := objc.Call[map[string]FeatureDescription](m_, objc.Sel("outputDescriptionsByName"))
	return rv
}
