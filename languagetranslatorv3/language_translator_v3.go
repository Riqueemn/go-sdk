// Package languagetranslatorv3 : Operations and models for the LanguageTranslatorV3 service
package languagetranslatorv3

/**
 * Copyright 2018 IBM All Rights Reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

import (
	"fmt"
	"os"

	core "github.com/ibm-watson/go-sdk/core"
)

// LanguageTranslatorV3 : The LanguageTranslatorV3 service
type LanguageTranslatorV3 struct {
	service *core.WatsonService
}

// LanguageTranslatorV3Options : Service options
type LanguageTranslatorV3Options struct {
	Version        string
	URL            string
	Username       string
	Password       string
	IAMApiKey      string
	IAMAccessToken string
	IAMURL         string
}

// NewLanguageTranslatorV3 : Instantiate LanguageTranslatorV3
func NewLanguageTranslatorV3(options *LanguageTranslatorV3Options) (*LanguageTranslatorV3, error) {
	if options.URL == "" {
		options.URL = "https://gateway.watsonplatform.net/language-translator/api"
	}

	serviceOptions := &core.ServiceOptions{
		URL:            options.URL,
		Version:        options.Version,
		Username:       options.Username,
		Password:       options.Password,
		IAMApiKey:      options.IAMApiKey,
		IAMAccessToken: options.IAMAccessToken,
		IAMURL:         options.IAMURL,
	}
	service, serviceErr := core.NewWatsonService(serviceOptions, "language_translator")
	if serviceErr != nil {
		return nil, serviceErr
	}

	return &LanguageTranslatorV3{service: service}, nil
}

// Translate : Translate
func (languageTranslator *LanguageTranslatorV3) Translate(translateOptions *TranslateOptions) (*core.DetailedResponse, error) {
	if err := core.ValidateNotNil(translateOptions, "translateOptions cannot be nil"); err != nil {
		return nil, err
	}
	if err := core.ValidateStruct(translateOptions, "translateOptions"); err != nil {
		return nil, err
	}

	pathSegments := []string{"v3/translate"}
	pathParameters := []string{}

	builder := core.NewRequestBuilder(core.POST)
	builder.ConstructHTTPURL(languageTranslator.service.Options.URL, pathSegments, pathParameters)

	for headerName, headerValue := range translateOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	builder.AddQuery("version", languageTranslator.service.Options.Version)

	body := make(map[string]interface{})
	if translateOptions.Text != nil {
		body["text"] = translateOptions.Text
	}
	if translateOptions.ModelID != nil {
		body["model_id"] = translateOptions.ModelID
	}
	if translateOptions.Source != nil {
		body["source"] = translateOptions.Source
	}
	if translateOptions.Target != nil {
		body["target"] = translateOptions.Target
	}
	_, err := builder.SetBodyContentJSON(body)
	if err != nil {
		return nil, err
	}

	request, err := builder.Build()
	if err != nil {
		return nil, err
	}

	response, err := languageTranslator.service.Request(request, new(TranslationResult))
	return response, err
}

// GetTranslateResult : Cast result of Translate operation
func (languageTranslator *LanguageTranslatorV3) GetTranslateResult(response *core.DetailedResponse) *TranslationResult {
	result, ok := response.Result.(*TranslationResult)
	if ok {
		return result
	}
	return nil
}

// Identify : Identify language
func (languageTranslator *LanguageTranslatorV3) Identify(identifyOptions *IdentifyOptions) (*core.DetailedResponse, error) {
	if err := core.ValidateNotNil(identifyOptions, "identifyOptions cannot be nil"); err != nil {
		return nil, err
	}
	if err := core.ValidateStruct(identifyOptions, "identifyOptions"); err != nil {
		return nil, err
	}

	pathSegments := []string{"v3/identify"}
	pathParameters := []string{}

	builder := core.NewRequestBuilder(core.POST)
	builder.ConstructHTTPURL(languageTranslator.service.Options.URL, pathSegments, pathParameters)

	for headerName, headerValue := range identifyOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "text/plain")
	builder.AddQuery("version", languageTranslator.service.Options.Version)

	_, err := builder.SetBodyContent("text/plain", nil, nil, identifyOptions.Text)
	if err != nil {
		return nil, err
	}

	request, err := builder.Build()
	if err != nil {
		return nil, err
	}

	response, err := languageTranslator.service.Request(request, new(IdentifiedLanguages))
	return response, err
}

// GetIdentifyResult : Cast result of Identify operation
func (languageTranslator *LanguageTranslatorV3) GetIdentifyResult(response *core.DetailedResponse) *IdentifiedLanguages {
	result, ok := response.Result.(*IdentifiedLanguages)
	if ok {
		return result
	}
	return nil
}

// ListIdentifiableLanguages : List identifiable languages
func (languageTranslator *LanguageTranslatorV3) ListIdentifiableLanguages(listIdentifiableLanguagesOptions *ListIdentifiableLanguagesOptions) (*core.DetailedResponse, error) {
	if err := core.ValidateStruct(listIdentifiableLanguagesOptions, "listIdentifiableLanguagesOptions"); err != nil {
		return nil, err
	}

	pathSegments := []string{"v3/identifiable_languages"}
	pathParameters := []string{}

	builder := core.NewRequestBuilder(core.GET)
	builder.ConstructHTTPURL(languageTranslator.service.Options.URL, pathSegments, pathParameters)

	for headerName, headerValue := range listIdentifiableLanguagesOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddQuery("version", languageTranslator.service.Options.Version)

	request, err := builder.Build()
	if err != nil {
		return nil, err
	}

	response, err := languageTranslator.service.Request(request, new(IdentifiableLanguages))
	return response, err
}

// GetListIdentifiableLanguagesResult : Cast result of ListIdentifiableLanguages operation
func (languageTranslator *LanguageTranslatorV3) GetListIdentifiableLanguagesResult(response *core.DetailedResponse) *IdentifiableLanguages {
	result, ok := response.Result.(*IdentifiableLanguages)
	if ok {
		return result
	}
	return nil
}

// CreateModel : Create model
func (languageTranslator *LanguageTranslatorV3) CreateModel(createModelOptions *CreateModelOptions) (*core.DetailedResponse, error) {
	if err := core.ValidateNotNil(createModelOptions, "createModelOptions cannot be nil"); err != nil {
		return nil, err
	}
	if err := core.ValidateStruct(createModelOptions, "createModelOptions"); err != nil {
		return nil, err
	}
	if (createModelOptions.ForcedGlossary == nil) && (createModelOptions.ParallelCorpus == nil) {
		return nil, fmt.Errorf("At least one of forcedGlossary or parallelCorpus must be supplied")
	}

	pathSegments := []string{"v3/models"}
	pathParameters := []string{}

	builder := core.NewRequestBuilder(core.POST)
	builder.ConstructHTTPURL(languageTranslator.service.Options.URL, pathSegments, pathParameters)

	for headerName, headerValue := range createModelOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	builder.AddQuery("base_model_id", fmt.Sprint(*createModelOptions.BaseModelID))
	if createModelOptions.Name != nil {
		builder.AddQuery("name", fmt.Sprint(*createModelOptions.Name))
	}
	builder.AddQuery("version", languageTranslator.service.Options.Version)

	if createModelOptions.ForcedGlossary != nil {
		builder.AddFormData("forced_glossary", core.StringNilMapper(createModelOptions.ForcedGlossaryFilename),
			"application/octet-stream", createModelOptions.ForcedGlossary)
	}
	if createModelOptions.ParallelCorpus != nil {
		builder.AddFormData("parallel_corpus", core.StringNilMapper(createModelOptions.ParallelCorpusFilename),
			"application/octet-stream", createModelOptions.ParallelCorpus)
	}

	request, err := builder.Build()
	if err != nil {
		return nil, err
	}

	response, err := languageTranslator.service.Request(request, new(TranslationModel))
	return response, err
}

// GetCreateModelResult : Cast result of CreateModel operation
func (languageTranslator *LanguageTranslatorV3) GetCreateModelResult(response *core.DetailedResponse) *TranslationModel {
	result, ok := response.Result.(*TranslationModel)
	if ok {
		return result
	}
	return nil
}

// DeleteModel : Delete model
func (languageTranslator *LanguageTranslatorV3) DeleteModel(deleteModelOptions *DeleteModelOptions) (*core.DetailedResponse, error) {
	if err := core.ValidateNotNil(deleteModelOptions, "deleteModelOptions cannot be nil"); err != nil {
		return nil, err
	}
	if err := core.ValidateStruct(deleteModelOptions, "deleteModelOptions"); err != nil {
		return nil, err
	}

	pathSegments := []string{"v3/models"}
	pathParameters := []string{*deleteModelOptions.ModelID}

	builder := core.NewRequestBuilder(core.DELETE)
	builder.ConstructHTTPURL(languageTranslator.service.Options.URL, pathSegments, pathParameters)

	for headerName, headerValue := range deleteModelOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddQuery("version", languageTranslator.service.Options.Version)

	request, err := builder.Build()
	if err != nil {
		return nil, err
	}

	response, err := languageTranslator.service.Request(request, new(DeleteModelResult))
	return response, err
}

// GetDeleteModelResult : Cast result of DeleteModel operation
func (languageTranslator *LanguageTranslatorV3) GetDeleteModelResult(response *core.DetailedResponse) *DeleteModelResult {
	result, ok := response.Result.(*DeleteModelResult)
	if ok {
		return result
	}
	return nil
}

// GetModel : Get model details
func (languageTranslator *LanguageTranslatorV3) GetModel(getModelOptions *GetModelOptions) (*core.DetailedResponse, error) {
	if err := core.ValidateNotNil(getModelOptions, "getModelOptions cannot be nil"); err != nil {
		return nil, err
	}
	if err := core.ValidateStruct(getModelOptions, "getModelOptions"); err != nil {
		return nil, err
	}

	pathSegments := []string{"v3/models"}
	pathParameters := []string{*getModelOptions.ModelID}

	builder := core.NewRequestBuilder(core.GET)
	builder.ConstructHTTPURL(languageTranslator.service.Options.URL, pathSegments, pathParameters)

	for headerName, headerValue := range getModelOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddQuery("version", languageTranslator.service.Options.Version)

	request, err := builder.Build()
	if err != nil {
		return nil, err
	}

	response, err := languageTranslator.service.Request(request, new(TranslationModel))
	return response, err
}

// GetGetModelResult : Cast result of GetModel operation
func (languageTranslator *LanguageTranslatorV3) GetGetModelResult(response *core.DetailedResponse) *TranslationModel {
	result, ok := response.Result.(*TranslationModel)
	if ok {
		return result
	}
	return nil
}

// ListModels : List models
func (languageTranslator *LanguageTranslatorV3) ListModels(listModelsOptions *ListModelsOptions) (*core.DetailedResponse, error) {
	if err := core.ValidateStruct(listModelsOptions, "listModelsOptions"); err != nil {
		return nil, err
	}

	pathSegments := []string{"v3/models"}
	pathParameters := []string{}

	builder := core.NewRequestBuilder(core.GET)
	builder.ConstructHTTPURL(languageTranslator.service.Options.URL, pathSegments, pathParameters)

	for headerName, headerValue := range listModelsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	if listModelsOptions.Source != nil {
		builder.AddQuery("source", fmt.Sprint(*listModelsOptions.Source))
	}
	if listModelsOptions.Target != nil {
		builder.AddQuery("target", fmt.Sprint(*listModelsOptions.Target))
	}
	if listModelsOptions.DefaultModels != nil {
		builder.AddQuery("default", fmt.Sprint(*listModelsOptions.DefaultModels))
	}
	builder.AddQuery("version", languageTranslator.service.Options.Version)

	request, err := builder.Build()
	if err != nil {
		return nil, err
	}

	response, err := languageTranslator.service.Request(request, new(TranslationModels))
	return response, err
}

// GetListModelsResult : Cast result of ListModels operation
func (languageTranslator *LanguageTranslatorV3) GetListModelsResult(response *core.DetailedResponse) *TranslationModels {
	result, ok := response.Result.(*TranslationModels)
	if ok {
		return result
	}
	return nil
}

// CreateModelOptions : The createModel options.
type CreateModelOptions struct {

	// The model ID of the model to use as the base for customization. To see available models, use the `List models` method. Usually all IBM provided models are customizable. In addition, all your models that have been created via parallel corpus customization, can be further customized with a forced glossary.
	BaseModelID *string `json:"base_model_id" validate:"required"`

	// An optional model name that you can use to identify the model. Valid characters are letters, numbers, dashes, underscores, spaces and apostrophes. The maximum length is 32 characters.
	Name *string `json:"name,omitempty"`

	// A TMX file with your customizations. The customizations in the file completely overwrite the domain translaton data, including high frequency or high confidence phrase translations. You can upload only one glossary with a file size less than 10 MB per call. A forced glossary should contain single words or short phrases.
	ForcedGlossary *os.File `json:"forced_glossary,omitempty"`

	// The filename for forcedGlossary.
	ForcedGlossaryFilename *string `json:"forced_glossary_filename,omitempty"`

	// A TMX file with parallel sentences for source and target language. You can upload multiple parallel_corpus files in one request. All uploaded parallel_corpus files combined, your parallel corpus must contain at least 5,000 parallel sentences to train successfully.
	ParallelCorpus *os.File `json:"parallel_corpus,omitempty"`

	// The filename for parallelCorpus.
	ParallelCorpusFilename *string `json:"parallel_corpus_filename,omitempty"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// NewCreateModelOptions : Instantiate CreateModelOptions
func (languageTranslator *LanguageTranslatorV3) NewCreateModelOptions(baseModelID string) *CreateModelOptions {
	return &CreateModelOptions{
		BaseModelID: core.StringPtr(baseModelID),
	}
}

// SetBaseModelID : Allow user to set BaseModelID
func (options *CreateModelOptions) SetBaseModelID(param string) *CreateModelOptions {
	options.BaseModelID = core.StringPtr(param)
	return options
}

// SetName : Allow user to set Name
func (options *CreateModelOptions) SetName(param string) *CreateModelOptions {
	options.Name = core.StringPtr(param)
	return options
}

// SetForcedGlossary : Allow user to set ForcedGlossary
func (options *CreateModelOptions) SetForcedGlossary(param *os.File) *CreateModelOptions {
	options.ForcedGlossary = param
	return options
}

// SetForcedGlossaryFilename : Allow user to set ForcedGlossaryFilename
func (options *CreateModelOptions) SetForcedGlossaryFilename(param string) *CreateModelOptions {
	options.ForcedGlossaryFilename = core.StringPtr(param)
	return options
}

// SetParallelCorpus : Allow user to set ParallelCorpus
func (options *CreateModelOptions) SetParallelCorpus(param *os.File) *CreateModelOptions {
	options.ParallelCorpus = param
	return options
}

// SetParallelCorpusFilename : Allow user to set ParallelCorpusFilename
func (options *CreateModelOptions) SetParallelCorpusFilename(param string) *CreateModelOptions {
	options.ParallelCorpusFilename = core.StringPtr(param)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *CreateModelOptions) SetHeaders(param map[string]string) *CreateModelOptions {
	options.Headers = param
	return options
}

// DeleteModelOptions : The deleteModel options.
type DeleteModelOptions struct {

	// Model ID of the model to delete.
	ModelID *string `json:"model_id" validate:"required"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// NewDeleteModelOptions : Instantiate DeleteModelOptions
func (languageTranslator *LanguageTranslatorV3) NewDeleteModelOptions(modelID string) *DeleteModelOptions {
	return &DeleteModelOptions{
		ModelID: core.StringPtr(modelID),
	}
}

// SetModelID : Allow user to set ModelID
func (options *DeleteModelOptions) SetModelID(param string) *DeleteModelOptions {
	options.ModelID = core.StringPtr(param)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteModelOptions) SetHeaders(param map[string]string) *DeleteModelOptions {
	options.Headers = param
	return options
}

// DeleteModelResult : DeleteModelResult struct
type DeleteModelResult struct {

	// "OK" indicates that the model was successfully deleted.
	Status *string `json:"status" validate:"required"`
}

// GetModelOptions : The getModel options.
type GetModelOptions struct {

	// Model ID of the model to get.
	ModelID *string `json:"model_id" validate:"required"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// NewGetModelOptions : Instantiate GetModelOptions
func (languageTranslator *LanguageTranslatorV3) NewGetModelOptions(modelID string) *GetModelOptions {
	return &GetModelOptions{
		ModelID: core.StringPtr(modelID),
	}
}

// SetModelID : Allow user to set ModelID
func (options *GetModelOptions) SetModelID(param string) *GetModelOptions {
	options.ModelID = core.StringPtr(param)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *GetModelOptions) SetHeaders(param map[string]string) *GetModelOptions {
	options.Headers = param
	return options
}

// IdentifiableLanguage : IdentifiableLanguage struct
type IdentifiableLanguage struct {

	// The language code for an identifiable language.
	Language *string `json:"language" validate:"required"`

	// The name of the identifiable language.
	Name *string `json:"name" validate:"required"`
}

// IdentifiableLanguages : IdentifiableLanguages struct
type IdentifiableLanguages struct {

	// A list of all languages that the service can identify.
	Languages []IdentifiableLanguage `json:"languages" validate:"required"`
}

// IdentifiedLanguage : IdentifiedLanguage struct
type IdentifiedLanguage struct {

	// The language code for an identified language.
	Language *string `json:"language" validate:"required"`

	// The confidence score for the identified language.
	Confidence *float64 `json:"confidence" validate:"required"`
}

// IdentifiedLanguages : IdentifiedLanguages struct
type IdentifiedLanguages struct {

	// A ranking of identified languages with confidence scores.
	Languages []IdentifiedLanguage `json:"languages" validate:"required"`
}

// IdentifyOptions : The identify options.
type IdentifyOptions struct {

	// Input text in UTF-8 format.
	Text *string `json:"text" validate:"required"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// NewIdentifyOptions : Instantiate IdentifyOptions
func (languageTranslator *LanguageTranslatorV3) NewIdentifyOptions(text string) *IdentifyOptions {
	return &IdentifyOptions{
		Text: core.StringPtr(text),
	}
}

// SetText : Allow user to set Text
func (options *IdentifyOptions) SetText(param string) *IdentifyOptions {
	options.Text = core.StringPtr(param)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *IdentifyOptions) SetHeaders(param map[string]string) *IdentifyOptions {
	options.Headers = param
	return options
}

// ListIdentifiableLanguagesOptions : The listIdentifiableLanguages options.
type ListIdentifiableLanguagesOptions struct {

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// NewListIdentifiableLanguagesOptions : Instantiate ListIdentifiableLanguagesOptions
func (languageTranslator *LanguageTranslatorV3) NewListIdentifiableLanguagesOptions() *ListIdentifiableLanguagesOptions {
	return &ListIdentifiableLanguagesOptions{}
}

// SetHeaders : Allow user to set Headers
func (options *ListIdentifiableLanguagesOptions) SetHeaders(param map[string]string) *ListIdentifiableLanguagesOptions {
	options.Headers = param
	return options
}

// ListModelsOptions : The listModels options.
type ListModelsOptions struct {

	// Specify a language code to filter results by source language.
	Source *string `json:"source,omitempty"`

	// Specify a language code to filter results by target language.
	Target *string `json:"target,omitempty"`

	// If the default parameter isn't specified, the service will return all models (default and non-default) for each language pair. To return only default models, set this to `true`. To return only non-default models, set this to `false`. There is exactly one default model per language pair, the IBM provided base model.
	DefaultModels *bool `json:"default,omitempty"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// NewListModelsOptions : Instantiate ListModelsOptions
func (languageTranslator *LanguageTranslatorV3) NewListModelsOptions() *ListModelsOptions {
	return &ListModelsOptions{}
}

// SetSource : Allow user to set Source
func (options *ListModelsOptions) SetSource(param string) *ListModelsOptions {
	options.Source = core.StringPtr(param)
	return options
}

// SetTarget : Allow user to set Target
func (options *ListModelsOptions) SetTarget(param string) *ListModelsOptions {
	options.Target = core.StringPtr(param)
	return options
}

// SetDefaultModels : Allow user to set DefaultModels
func (options *ListModelsOptions) SetDefaultModels(param bool) *ListModelsOptions {
	options.DefaultModels = core.BoolPtr(param)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *ListModelsOptions) SetHeaders(param map[string]string) *ListModelsOptions {
	options.Headers = param
	return options
}

// TranslateOptions : The translate options.
type TranslateOptions struct {

	// Input text in UTF-8 encoding. Multiple entries will result in multiple translations in the response.
	Text []string `json:"text" validate:"required"`

	// Model ID of the translation model to use. If this is specified, the **source** and **target** parameters will be ignored. The method requires either a model ID or both the **source** and **target** parameters.
	ModelID *string `json:"model_id,omitempty"`

	// Language code of the source text language. Use with `target` as an alternative way to select a translation model. When `source` and `target` are set, and a model ID is not set, the system chooses a default model for the language pair (usually the model based on the news domain).
	Source *string `json:"source,omitempty"`

	// Language code of the translation target language. Use with source as an alternative way to select a translation model.
	Target *string `json:"target,omitempty"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// NewTranslateOptions : Instantiate TranslateOptions
func (languageTranslator *LanguageTranslatorV3) NewTranslateOptions(text []string) *TranslateOptions {
	return &TranslateOptions{
		Text: text,
	}
}

// SetText : Allow user to set Text
func (options *TranslateOptions) SetText(param []string) *TranslateOptions {
	options.Text = param
	return options
}

// SetModelID : Allow user to set ModelID
func (options *TranslateOptions) SetModelID(param string) *TranslateOptions {
	options.ModelID = core.StringPtr(param)
	return options
}

// SetSource : Allow user to set Source
func (options *TranslateOptions) SetSource(param string) *TranslateOptions {
	options.Source = core.StringPtr(param)
	return options
}

// SetTarget : Allow user to set Target
func (options *TranslateOptions) SetTarget(param string) *TranslateOptions {
	options.Target = core.StringPtr(param)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *TranslateOptions) SetHeaders(param map[string]string) *TranslateOptions {
	options.Headers = param
	return options
}

// Translation : Translation struct
type Translation struct {

	// Translation output in UTF-8.
	TranslationOutput *string `json:"translation" validate:"required"`
}

// TranslationModel : Response payload for models.
type TranslationModel struct {

	// A globally unique string that identifies the underlying model that is used for translation.
	ModelID *string `json:"model_id" validate:"required"`

	// Optional name that can be specified when the model is created.
	Name *string `json:"name,omitempty"`

	// Translation source language code.
	Source *string `json:"source,omitempty"`

	// Translation target language code.
	Target *string `json:"target,omitempty"`

	// Model ID of the base model that was used to customize the model. If the model is not a custom model, this will be an empty string.
	BaseModelID *string `json:"base_model_id,omitempty"`

	// The domain of the translation model.
	Domain *string `json:"domain,omitempty"`

	// Whether this model can be used as a base for customization. Customized models are not further customizable, and some base models are not customizable.
	Customizable *bool `json:"customizable,omitempty"`

	// Whether or not the model is a default model. A default model is the model for a given language pair that will be used when that language pair is specified in the source and target parameters.
	DefaultModel *bool `json:"default_model,omitempty"`

	// Either an empty string, indicating the model is not a custom model, or the ID of the service instance that created the model.
	Owner *string `json:"owner,omitempty"`

	// Availability of a model.
	Status *string `json:"status,omitempty"`
}

// TranslationModels : The response type for listing existing translation models.
type TranslationModels struct {

	// An array of available models.
	Models []TranslationModel `json:"models" validate:"required"`
}

// TranslationResult : TranslationResult struct
type TranslationResult struct {

	// Number of words in the input text.
	WordCount *int64 `json:"word_count" validate:"required"`

	// Number of characters in the input text.
	CharacterCount *int64 `json:"character_count" validate:"required"`

	// List of translation output in UTF-8, corresponding to the input text entries.
	Translations []Translation `json:"translations" validate:"required"`
}