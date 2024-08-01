// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package openai

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"reflect"

	"github.com/openai/openai-go/internal/apijson"
	"github.com/openai/openai-go/internal/param"
	"github.com/openai/openai-go/internal/requestconfig"
	"github.com/openai/openai-go/option"
	"github.com/openai/openai-go/packages/ssestream"
	"github.com/openai/openai-go/shared"
	"github.com/tidwall/gjson"
)

// BetaThreadService contains methods and other services that help with interacting
// with the openai API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBetaThreadService] method instead.
type BetaThreadService struct {
	Options  []option.RequestOption
	Runs     *BetaThreadRunService
	Messages *BetaThreadMessageService
}

// NewBetaThreadService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewBetaThreadService(opts ...option.RequestOption) (r *BetaThreadService) {
	r = &BetaThreadService{}
	r.Options = opts
	r.Runs = NewBetaThreadRunService(opts...)
	r.Messages = NewBetaThreadMessageService(opts...)
	return
}

// Create a thread.
func (r *BetaThreadService) New(ctx context.Context, body BetaThreadNewParams, opts ...option.RequestOption) (res *Thread, err error) {
	opts = append(r.Options[:], opts...)
	path := "threads"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieves a thread.
func (r *BetaThreadService) Get(ctx context.Context, threadID string, opts ...option.RequestOption) (res *Thread, err error) {
	opts = append(r.Options[:], opts...)
	if threadID == "" {
		err = errors.New("missing required thread_id parameter")
		return
	}
	path := fmt.Sprintf("threads/%s", threadID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Modifies a thread.
func (r *BetaThreadService) Update(ctx context.Context, threadID string, body BetaThreadUpdateParams, opts ...option.RequestOption) (res *Thread, err error) {
	opts = append(r.Options[:], opts...)
	if threadID == "" {
		err = errors.New("missing required thread_id parameter")
		return
	}
	path := fmt.Sprintf("threads/%s", threadID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Delete a thread.
func (r *BetaThreadService) Delete(ctx context.Context, threadID string, opts ...option.RequestOption) (res *ThreadDeleted, err error) {
	opts = append(r.Options[:], opts...)
	if threadID == "" {
		err = errors.New("missing required thread_id parameter")
		return
	}
	path := fmt.Sprintf("threads/%s", threadID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Create a thread and run it in one request.
func (r *BetaThreadService) NewAndRun(ctx context.Context, body BetaThreadNewAndRunParams, opts ...option.RequestOption) (res *Run, err error) {
	opts = append(r.Options[:], opts...)
	path := "threads/runs"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Create a thread and run it in one request.
func (r *BetaThreadService) NewAndRunStreaming(ctx context.Context, body BetaThreadNewAndRunParams, opts ...option.RequestOption) (stream *ssestream.Stream[AssistantStreamEvent]) {
	var (
		raw *http.Response
		err error
	)
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithJSONSet("stream", true)}, opts...)
	path := "threads/runs"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &raw, opts...)
	return ssestream.NewStream[AssistantStreamEvent](ssestream.NewDecoder(raw), err)
}

// An object describing the expected output of the model. If `json_object` only
// `function` type `tools` are allowed to be passed to the Run. If `text` the model
// can return text or any value needed.
type AssistantResponseFormat struct {
	// Must be one of `text` or `json_object`.
	Type AssistantResponseFormatType `json:"type"`
	JSON assistantResponseFormatJSON `json:"-"`
}

// assistantResponseFormatJSON contains the JSON metadata for the struct
// [AssistantResponseFormat]
type assistantResponseFormatJSON struct {
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AssistantResponseFormat) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r assistantResponseFormatJSON) RawJSON() string {
	return r.raw
}

func (r AssistantResponseFormat) implementsAssistantResponseFormatOptionUnion() {}

// Must be one of `text` or `json_object`.
type AssistantResponseFormatType string

const (
	AssistantResponseFormatTypeText       AssistantResponseFormatType = "text"
	AssistantResponseFormatTypeJSONObject AssistantResponseFormatType = "json_object"
)

func (r AssistantResponseFormatType) IsKnown() bool {
	switch r {
	case AssistantResponseFormatTypeText, AssistantResponseFormatTypeJSONObject:
		return true
	}
	return false
}

// An object describing the expected output of the model. If `json_object` only
// `function` type `tools` are allowed to be passed to the Run. If `text` the model
// can return text or any value needed.
type AssistantResponseFormatParam struct {
	// Must be one of `text` or `json_object`.
	Type param.Field[AssistantResponseFormatType] `json:"type"`
}

func (r AssistantResponseFormatParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AssistantResponseFormatParam) implementsAssistantResponseFormatOptionUnionParam() {}

// Specifies the format that the model must output. Compatible with
// [GPT-4o](https://platform.openai.com/docs/models/gpt-4o),
// [GPT-4 Turbo](https://platform.openai.com/docs/models/gpt-4-turbo-and-gpt-4),
// and all GPT-3.5 Turbo models since `gpt-3.5-turbo-1106`.
//
// Setting to `{ "type": "json_object" }` enables JSON mode, which guarantees the
// message the model generates is valid JSON.
//
// **Important:** when using JSON mode, you **must** also instruct the model to
// produce JSON yourself via a system or user message. Without this, the model may
// generate an unending stream of whitespace until the generation reaches the token
// limit, resulting in a long-running and seemingly "stuck" request. Also note that
// the message content may be partially cut off if `finish_reason="length"`, which
// indicates the generation exceeded `max_tokens` or the conversation exceeded the
// max context length.
//
// Union satisfied by [AssistantResponseFormatOptionString] or
// [AssistantResponseFormat].
type AssistantResponseFormatOptionUnion interface {
	implementsAssistantResponseFormatOptionUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AssistantResponseFormatOptionUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(AssistantResponseFormatOptionString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AssistantResponseFormat{}),
		},
	)
}

// `auto` is the default value
type AssistantResponseFormatOptionString string

const (
	AssistantResponseFormatOptionStringNone AssistantResponseFormatOptionString = "none"
	AssistantResponseFormatOptionStringAuto AssistantResponseFormatOptionString = "auto"
)

func (r AssistantResponseFormatOptionString) IsKnown() bool {
	switch r {
	case AssistantResponseFormatOptionStringNone, AssistantResponseFormatOptionStringAuto:
		return true
	}
	return false
}

func (r AssistantResponseFormatOptionString) implementsAssistantResponseFormatOptionUnion() {}

func (r AssistantResponseFormatOptionString) implementsAssistantResponseFormatOptionUnionParam() {}

// Specifies the format that the model must output. Compatible with
// [GPT-4o](https://platform.openai.com/docs/models/gpt-4o),
// [GPT-4 Turbo](https://platform.openai.com/docs/models/gpt-4-turbo-and-gpt-4),
// and all GPT-3.5 Turbo models since `gpt-3.5-turbo-1106`.
//
// Setting to `{ "type": "json_object" }` enables JSON mode, which guarantees the
// message the model generates is valid JSON.
//
// **Important:** when using JSON mode, you **must** also instruct the model to
// produce JSON yourself via a system or user message. Without this, the model may
// generate an unending stream of whitespace until the generation reaches the token
// limit, resulting in a long-running and seemingly "stuck" request. Also note that
// the message content may be partially cut off if `finish_reason="length"`, which
// indicates the generation exceeded `max_tokens` or the conversation exceeded the
// max context length.
//
// Satisfied by [AssistantResponseFormatOptionString],
// [AssistantResponseFormatParam].
type AssistantResponseFormatOptionUnionParam interface {
	implementsAssistantResponseFormatOptionUnionParam()
}

// Specifies a tool the model should use. Use to force the model to call a specific
// tool.
type AssistantToolChoice struct {
	// The type of the tool. If type is `function`, the function name must be set
	Type     AssistantToolChoiceType     `json:"type,required"`
	Function AssistantToolChoiceFunction `json:"function"`
	JSON     assistantToolChoiceJSON     `json:"-"`
}

// assistantToolChoiceJSON contains the JSON metadata for the struct
// [AssistantToolChoice]
type assistantToolChoiceJSON struct {
	Type        apijson.Field
	Function    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AssistantToolChoice) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r assistantToolChoiceJSON) RawJSON() string {
	return r.raw
}

func (r AssistantToolChoice) implementsAssistantToolChoiceOptionUnion() {}

// The type of the tool. If type is `function`, the function name must be set
type AssistantToolChoiceType string

const (
	AssistantToolChoiceTypeFunction        AssistantToolChoiceType = "function"
	AssistantToolChoiceTypeCodeInterpreter AssistantToolChoiceType = "code_interpreter"
	AssistantToolChoiceTypeFileSearch      AssistantToolChoiceType = "file_search"
)

func (r AssistantToolChoiceType) IsKnown() bool {
	switch r {
	case AssistantToolChoiceTypeFunction, AssistantToolChoiceTypeCodeInterpreter, AssistantToolChoiceTypeFileSearch:
		return true
	}
	return false
}

// Specifies a tool the model should use. Use to force the model to call a specific
// tool.
type AssistantToolChoiceParam struct {
	// The type of the tool. If type is `function`, the function name must be set
	Type     param.Field[AssistantToolChoiceType]          `json:"type,required"`
	Function param.Field[AssistantToolChoiceFunctionParam] `json:"function"`
}

func (r AssistantToolChoiceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AssistantToolChoiceParam) implementsAssistantToolChoiceOptionUnionParam() {}

type AssistantToolChoiceFunction struct {
	// The name of the function to call.
	Name string                          `json:"name,required"`
	JSON assistantToolChoiceFunctionJSON `json:"-"`
}

// assistantToolChoiceFunctionJSON contains the JSON metadata for the struct
// [AssistantToolChoiceFunction]
type assistantToolChoiceFunctionJSON struct {
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AssistantToolChoiceFunction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r assistantToolChoiceFunctionJSON) RawJSON() string {
	return r.raw
}

type AssistantToolChoiceFunctionParam struct {
	// The name of the function to call.
	Name param.Field[string] `json:"name,required"`
}

func (r AssistantToolChoiceFunctionParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Controls which (if any) tool is called by the model. `none` means the model will
// not call any tools and instead generates a message. `auto` is the default value
// and means the model can pick between generating a message or calling one or more
// tools. `required` means the model must call one or more tools before responding
// to the user. Specifying a particular tool like `{"type": "file_search"}` or
// `{"type": "function", "function": {"name": "my_function"}}` forces the model to
// call that tool.
//
// Union satisfied by [AssistantToolChoiceOptionString] or [AssistantToolChoice].
type AssistantToolChoiceOptionUnion interface {
	implementsAssistantToolChoiceOptionUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AssistantToolChoiceOptionUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(AssistantToolChoiceOptionString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AssistantToolChoice{}),
		},
	)
}

// `none` means the model will not call any tools and instead generates a message.
// `auto` means the model can pick between generating a message or calling one or
// more tools. `required` means the model must call one or more tools before
// responding to the user.
type AssistantToolChoiceOptionString string

const (
	AssistantToolChoiceOptionStringNone     AssistantToolChoiceOptionString = "none"
	AssistantToolChoiceOptionStringAuto     AssistantToolChoiceOptionString = "auto"
	AssistantToolChoiceOptionStringRequired AssistantToolChoiceOptionString = "required"
)

func (r AssistantToolChoiceOptionString) IsKnown() bool {
	switch r {
	case AssistantToolChoiceOptionStringNone, AssistantToolChoiceOptionStringAuto, AssistantToolChoiceOptionStringRequired:
		return true
	}
	return false
}

func (r AssistantToolChoiceOptionString) implementsAssistantToolChoiceOptionUnion() {}

func (r AssistantToolChoiceOptionString) implementsAssistantToolChoiceOptionUnionParam() {}

// Controls which (if any) tool is called by the model. `none` means the model will
// not call any tools and instead generates a message. `auto` is the default value
// and means the model can pick between generating a message or calling one or more
// tools. `required` means the model must call one or more tools before responding
// to the user. Specifying a particular tool like `{"type": "file_search"}` or
// `{"type": "function", "function": {"name": "my_function"}}` forces the model to
// call that tool.
//
// Satisfied by [AssistantToolChoiceOptionString], [AssistantToolChoiceParam].
type AssistantToolChoiceOptionUnionParam interface {
	implementsAssistantToolChoiceOptionUnionParam()
}

// Represents a thread that contains
// [messages](https://platform.openai.com/docs/api-reference/messages).
type Thread struct {
	// The identifier, which can be referenced in API endpoints.
	ID string `json:"id,required"`
	// The Unix timestamp (in seconds) for when the thread was created.
	CreatedAt int64 `json:"created_at,required"`
	// Set of 16 key-value pairs that can be attached to an object. This can be useful
	// for storing additional information about the object in a structured format. Keys
	// can be a maximum of 64 characters long and values can be a maxium of 512
	// characters long.
	Metadata interface{} `json:"metadata,required,nullable"`
	// The object type, which is always `thread`.
	Object ThreadObject `json:"object,required"`
	// A set of resources that are made available to the assistant's tools in this
	// thread. The resources are specific to the type of tool. For example, the
	// `code_interpreter` tool requires a list of file IDs, while the `file_search`
	// tool requires a list of vector store IDs.
	ToolResources ThreadToolResources `json:"tool_resources,required,nullable"`
	JSON          threadJSON          `json:"-"`
}

// threadJSON contains the JSON metadata for the struct [Thread]
type threadJSON struct {
	ID            apijson.Field
	CreatedAt     apijson.Field
	Metadata      apijson.Field
	Object        apijson.Field
	ToolResources apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *Thread) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threadJSON) RawJSON() string {
	return r.raw
}

// The object type, which is always `thread`.
type ThreadObject string

const (
	ThreadObjectThread ThreadObject = "thread"
)

func (r ThreadObject) IsKnown() bool {
	switch r {
	case ThreadObjectThread:
		return true
	}
	return false
}

// A set of resources that are made available to the assistant's tools in this
// thread. The resources are specific to the type of tool. For example, the
// `code_interpreter` tool requires a list of file IDs, while the `file_search`
// tool requires a list of vector store IDs.
type ThreadToolResources struct {
	CodeInterpreter ThreadToolResourcesCodeInterpreter `json:"code_interpreter"`
	FileSearch      ThreadToolResourcesFileSearch      `json:"file_search"`
	JSON            threadToolResourcesJSON            `json:"-"`
}

// threadToolResourcesJSON contains the JSON metadata for the struct
// [ThreadToolResources]
type threadToolResourcesJSON struct {
	CodeInterpreter apijson.Field
	FileSearch      apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ThreadToolResources) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threadToolResourcesJSON) RawJSON() string {
	return r.raw
}

type ThreadToolResourcesCodeInterpreter struct {
	// A list of [file](https://platform.openai.com/docs/api-reference/files) IDs made
	// available to the `code_interpreter` tool. There can be a maximum of 20 files
	// associated with the tool.
	FileIDs []string                               `json:"file_ids"`
	JSON    threadToolResourcesCodeInterpreterJSON `json:"-"`
}

// threadToolResourcesCodeInterpreterJSON contains the JSON metadata for the struct
// [ThreadToolResourcesCodeInterpreter]
type threadToolResourcesCodeInterpreterJSON struct {
	FileIDs     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreadToolResourcesCodeInterpreter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threadToolResourcesCodeInterpreterJSON) RawJSON() string {
	return r.raw
}

type ThreadToolResourcesFileSearch struct {
	// The
	// [vector store](https://platform.openai.com/docs/api-reference/vector-stores/object)
	// attached to this thread. There can be a maximum of 1 vector store attached to
	// the thread.
	VectorStoreIDs []string                          `json:"vector_store_ids"`
	JSON           threadToolResourcesFileSearchJSON `json:"-"`
}

// threadToolResourcesFileSearchJSON contains the JSON metadata for the struct
// [ThreadToolResourcesFileSearch]
type threadToolResourcesFileSearchJSON struct {
	VectorStoreIDs apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ThreadToolResourcesFileSearch) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threadToolResourcesFileSearchJSON) RawJSON() string {
	return r.raw
}

type ThreadDeleted struct {
	ID      string              `json:"id,required"`
	Deleted bool                `json:"deleted,required"`
	Object  ThreadDeletedObject `json:"object,required"`
	JSON    threadDeletedJSON   `json:"-"`
}

// threadDeletedJSON contains the JSON metadata for the struct [ThreadDeleted]
type threadDeletedJSON struct {
	ID          apijson.Field
	Deleted     apijson.Field
	Object      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreadDeleted) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threadDeletedJSON) RawJSON() string {
	return r.raw
}

type ThreadDeletedObject string

const (
	ThreadDeletedObjectThreadDeleted ThreadDeletedObject = "thread.deleted"
)

func (r ThreadDeletedObject) IsKnown() bool {
	switch r {
	case ThreadDeletedObjectThreadDeleted:
		return true
	}
	return false
}

type BetaThreadNewParams struct {
	// A list of [messages](https://platform.openai.com/docs/api-reference/messages) to
	// start the thread with.
	Messages param.Field[[]BetaThreadNewParamsMessage] `json:"messages"`
	// Set of 16 key-value pairs that can be attached to an object. This can be useful
	// for storing additional information about the object in a structured format. Keys
	// can be a maximum of 64 characters long and values can be a maxium of 512
	// characters long.
	Metadata param.Field[interface{}] `json:"metadata"`
	// A set of resources that are made available to the assistant's tools in this
	// thread. The resources are specific to the type of tool. For example, the
	// `code_interpreter` tool requires a list of file IDs, while the `file_search`
	// tool requires a list of vector store IDs.
	ToolResources param.Field[BetaThreadNewParamsToolResources] `json:"tool_resources"`
}

func (r BetaThreadNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BetaThreadNewParamsMessage struct {
	// The text contents of the message.
	Content param.Field[BetaThreadNewParamsMessagesContentUnion] `json:"content,required"`
	// The role of the entity that is creating the message. Allowed values include:
	//
	//   - `user`: Indicates the message is sent by an actual user and should be used in
	//     most cases to represent user-generated messages.
	//   - `assistant`: Indicates the message is generated by the assistant. Use this
	//     value to insert messages from the assistant into the conversation.
	Role param.Field[BetaThreadNewParamsMessagesRole] `json:"role,required"`
	// A list of files attached to the message, and the tools they should be added to.
	Attachments param.Field[[]BetaThreadNewParamsMessagesAttachment] `json:"attachments"`
	// Set of 16 key-value pairs that can be attached to an object. This can be useful
	// for storing additional information about the object in a structured format. Keys
	// can be a maximum of 64 characters long and values can be a maxium of 512
	// characters long.
	Metadata param.Field[interface{}] `json:"metadata"`
}

func (r BetaThreadNewParamsMessage) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The text contents of the message.
//
// Satisfied by [shared.UnionString],
// [BetaThreadNewParamsMessagesContentArrayOfContentParts].
type BetaThreadNewParamsMessagesContentUnion interface {
	ImplementsBetaThreadNewParamsMessagesContentUnion()
}

type BetaThreadNewParamsMessagesContentArrayOfContentParts []MessageContentPartParamUnion

func (r BetaThreadNewParamsMessagesContentArrayOfContentParts) ImplementsBetaThreadNewParamsMessagesContentUnion() {
}

// The role of the entity that is creating the message. Allowed values include:
//
//   - `user`: Indicates the message is sent by an actual user and should be used in
//     most cases to represent user-generated messages.
//   - `assistant`: Indicates the message is generated by the assistant. Use this
//     value to insert messages from the assistant into the conversation.
type BetaThreadNewParamsMessagesRole string

const (
	BetaThreadNewParamsMessagesRoleUser      BetaThreadNewParamsMessagesRole = "user"
	BetaThreadNewParamsMessagesRoleAssistant BetaThreadNewParamsMessagesRole = "assistant"
)

func (r BetaThreadNewParamsMessagesRole) IsKnown() bool {
	switch r {
	case BetaThreadNewParamsMessagesRoleUser, BetaThreadNewParamsMessagesRoleAssistant:
		return true
	}
	return false
}

type BetaThreadNewParamsMessagesAttachment struct {
	// The ID of the file to attach to the message.
	FileID param.Field[string] `json:"file_id"`
	// The tools to add this file to.
	Tools param.Field[[]BetaThreadNewParamsMessagesAttachmentsToolUnion] `json:"tools"`
}

func (r BetaThreadNewParamsMessagesAttachment) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BetaThreadNewParamsMessagesAttachmentsTool struct {
	// The type of tool being defined: `code_interpreter`
	Type param.Field[BetaThreadNewParamsMessagesAttachmentsToolsType] `json:"type,required"`
}

func (r BetaThreadNewParamsMessagesAttachmentsTool) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaThreadNewParamsMessagesAttachmentsTool) implementsBetaThreadNewParamsMessagesAttachmentsToolUnion() {
}

// Satisfied by [CodeInterpreterToolParam],
// [BetaThreadNewParamsMessagesAttachmentsToolsFileSearch],
// [BetaThreadNewParamsMessagesAttachmentsTool].
type BetaThreadNewParamsMessagesAttachmentsToolUnion interface {
	implementsBetaThreadNewParamsMessagesAttachmentsToolUnion()
}

type BetaThreadNewParamsMessagesAttachmentsToolsFileSearch struct {
	// The type of tool being defined: `file_search`
	Type param.Field[BetaThreadNewParamsMessagesAttachmentsToolsFileSearchType] `json:"type,required"`
}

func (r BetaThreadNewParamsMessagesAttachmentsToolsFileSearch) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaThreadNewParamsMessagesAttachmentsToolsFileSearch) implementsBetaThreadNewParamsMessagesAttachmentsToolUnion() {
}

// The type of tool being defined: `file_search`
type BetaThreadNewParamsMessagesAttachmentsToolsFileSearchType string

const (
	BetaThreadNewParamsMessagesAttachmentsToolsFileSearchTypeFileSearch BetaThreadNewParamsMessagesAttachmentsToolsFileSearchType = "file_search"
)

func (r BetaThreadNewParamsMessagesAttachmentsToolsFileSearchType) IsKnown() bool {
	switch r {
	case BetaThreadNewParamsMessagesAttachmentsToolsFileSearchTypeFileSearch:
		return true
	}
	return false
}

// The type of tool being defined: `code_interpreter`
type BetaThreadNewParamsMessagesAttachmentsToolsType string

const (
	BetaThreadNewParamsMessagesAttachmentsToolsTypeCodeInterpreter BetaThreadNewParamsMessagesAttachmentsToolsType = "code_interpreter"
	BetaThreadNewParamsMessagesAttachmentsToolsTypeFileSearch      BetaThreadNewParamsMessagesAttachmentsToolsType = "file_search"
)

func (r BetaThreadNewParamsMessagesAttachmentsToolsType) IsKnown() bool {
	switch r {
	case BetaThreadNewParamsMessagesAttachmentsToolsTypeCodeInterpreter, BetaThreadNewParamsMessagesAttachmentsToolsTypeFileSearch:
		return true
	}
	return false
}

// A set of resources that are made available to the assistant's tools in this
// thread. The resources are specific to the type of tool. For example, the
// `code_interpreter` tool requires a list of file IDs, while the `file_search`
// tool requires a list of vector store IDs.
type BetaThreadNewParamsToolResources struct {
	CodeInterpreter param.Field[BetaThreadNewParamsToolResourcesCodeInterpreter] `json:"code_interpreter"`
	FileSearch      param.Field[BetaThreadNewParamsToolResourcesFileSearch]      `json:"file_search"`
}

func (r BetaThreadNewParamsToolResources) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BetaThreadNewParamsToolResourcesCodeInterpreter struct {
	// A list of [file](https://platform.openai.com/docs/api-reference/files) IDs made
	// available to the `code_interpreter` tool. There can be a maximum of 20 files
	// associated with the tool.
	FileIDs param.Field[[]string] `json:"file_ids"`
}

func (r BetaThreadNewParamsToolResourcesCodeInterpreter) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BetaThreadNewParamsToolResourcesFileSearch struct {
	// The
	// [vector store](https://platform.openai.com/docs/api-reference/vector-stores/object)
	// attached to this thread. There can be a maximum of 1 vector store attached to
	// the thread.
	VectorStoreIDs param.Field[[]string] `json:"vector_store_ids"`
	// A helper to create a
	// [vector store](https://platform.openai.com/docs/api-reference/vector-stores/object)
	// with file_ids and attach it to this thread. There can be a maximum of 1 vector
	// store attached to the thread.
	VectorStores param.Field[[]BetaThreadNewParamsToolResourcesFileSearchVectorStore] `json:"vector_stores"`
}

func (r BetaThreadNewParamsToolResourcesFileSearch) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BetaThreadNewParamsToolResourcesFileSearchVectorStore struct {
	// The chunking strategy used to chunk the file(s). If not set, will use the `auto`
	// strategy.
	ChunkingStrategy param.Field[BetaThreadNewParamsToolResourcesFileSearchVectorStoresChunkingStrategyUnion] `json:"chunking_strategy"`
	// A list of [file](https://platform.openai.com/docs/api-reference/files) IDs to
	// add to the vector store. There can be a maximum of 10000 files in a vector
	// store.
	FileIDs param.Field[[]string] `json:"file_ids"`
	// Set of 16 key-value pairs that can be attached to a vector store. This can be
	// useful for storing additional information about the vector store in a structured
	// format. Keys can be a maximum of 64 characters long and values can be a maxium
	// of 512 characters long.
	Metadata param.Field[interface{}] `json:"metadata"`
}

func (r BetaThreadNewParamsToolResourcesFileSearchVectorStore) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The chunking strategy used to chunk the file(s). If not set, will use the `auto`
// strategy.
type BetaThreadNewParamsToolResourcesFileSearchVectorStoresChunkingStrategy struct {
	// Always `auto`.
	Type   param.Field[BetaThreadNewParamsToolResourcesFileSearchVectorStoresChunkingStrategyType] `json:"type,required"`
	Static param.Field[interface{}]                                                                `json:"static,required"`
}

func (r BetaThreadNewParamsToolResourcesFileSearchVectorStoresChunkingStrategy) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaThreadNewParamsToolResourcesFileSearchVectorStoresChunkingStrategy) implementsBetaThreadNewParamsToolResourcesFileSearchVectorStoresChunkingStrategyUnion() {
}

// The chunking strategy used to chunk the file(s). If not set, will use the `auto`
// strategy.
//
// Satisfied by
// [BetaThreadNewParamsToolResourcesFileSearchVectorStoresChunkingStrategyAuto],
// [BetaThreadNewParamsToolResourcesFileSearchVectorStoresChunkingStrategyStatic],
// [BetaThreadNewParamsToolResourcesFileSearchVectorStoresChunkingStrategy].
type BetaThreadNewParamsToolResourcesFileSearchVectorStoresChunkingStrategyUnion interface {
	implementsBetaThreadNewParamsToolResourcesFileSearchVectorStoresChunkingStrategyUnion()
}

// The default strategy. This strategy currently uses a `max_chunk_size_tokens` of
// `800` and `chunk_overlap_tokens` of `400`.
type BetaThreadNewParamsToolResourcesFileSearchVectorStoresChunkingStrategyAuto struct {
	// Always `auto`.
	Type param.Field[BetaThreadNewParamsToolResourcesFileSearchVectorStoresChunkingStrategyAutoType] `json:"type,required"`
}

func (r BetaThreadNewParamsToolResourcesFileSearchVectorStoresChunkingStrategyAuto) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaThreadNewParamsToolResourcesFileSearchVectorStoresChunkingStrategyAuto) implementsBetaThreadNewParamsToolResourcesFileSearchVectorStoresChunkingStrategyUnion() {
}

// Always `auto`.
type BetaThreadNewParamsToolResourcesFileSearchVectorStoresChunkingStrategyAutoType string

const (
	BetaThreadNewParamsToolResourcesFileSearchVectorStoresChunkingStrategyAutoTypeAuto BetaThreadNewParamsToolResourcesFileSearchVectorStoresChunkingStrategyAutoType = "auto"
)

func (r BetaThreadNewParamsToolResourcesFileSearchVectorStoresChunkingStrategyAutoType) IsKnown() bool {
	switch r {
	case BetaThreadNewParamsToolResourcesFileSearchVectorStoresChunkingStrategyAutoTypeAuto:
		return true
	}
	return false
}

type BetaThreadNewParamsToolResourcesFileSearchVectorStoresChunkingStrategyStatic struct {
	Static param.Field[BetaThreadNewParamsToolResourcesFileSearchVectorStoresChunkingStrategyStaticStatic] `json:"static,required"`
	// Always `static`.
	Type param.Field[BetaThreadNewParamsToolResourcesFileSearchVectorStoresChunkingStrategyStaticType] `json:"type,required"`
}

func (r BetaThreadNewParamsToolResourcesFileSearchVectorStoresChunkingStrategyStatic) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaThreadNewParamsToolResourcesFileSearchVectorStoresChunkingStrategyStatic) implementsBetaThreadNewParamsToolResourcesFileSearchVectorStoresChunkingStrategyUnion() {
}

type BetaThreadNewParamsToolResourcesFileSearchVectorStoresChunkingStrategyStaticStatic struct {
	// The number of tokens that overlap between chunks. The default value is `400`.
	//
	// Note that the overlap must not exceed half of `max_chunk_size_tokens`.
	ChunkOverlapTokens param.Field[int64] `json:"chunk_overlap_tokens,required"`
	// The maximum number of tokens in each chunk. The default value is `800`. The
	// minimum value is `100` and the maximum value is `4096`.
	MaxChunkSizeTokens param.Field[int64] `json:"max_chunk_size_tokens,required"`
}

func (r BetaThreadNewParamsToolResourcesFileSearchVectorStoresChunkingStrategyStaticStatic) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Always `static`.
type BetaThreadNewParamsToolResourcesFileSearchVectorStoresChunkingStrategyStaticType string

const (
	BetaThreadNewParamsToolResourcesFileSearchVectorStoresChunkingStrategyStaticTypeStatic BetaThreadNewParamsToolResourcesFileSearchVectorStoresChunkingStrategyStaticType = "static"
)

func (r BetaThreadNewParamsToolResourcesFileSearchVectorStoresChunkingStrategyStaticType) IsKnown() bool {
	switch r {
	case BetaThreadNewParamsToolResourcesFileSearchVectorStoresChunkingStrategyStaticTypeStatic:
		return true
	}
	return false
}

// Always `auto`.
type BetaThreadNewParamsToolResourcesFileSearchVectorStoresChunkingStrategyType string

const (
	BetaThreadNewParamsToolResourcesFileSearchVectorStoresChunkingStrategyTypeAuto   BetaThreadNewParamsToolResourcesFileSearchVectorStoresChunkingStrategyType = "auto"
	BetaThreadNewParamsToolResourcesFileSearchVectorStoresChunkingStrategyTypeStatic BetaThreadNewParamsToolResourcesFileSearchVectorStoresChunkingStrategyType = "static"
)

func (r BetaThreadNewParamsToolResourcesFileSearchVectorStoresChunkingStrategyType) IsKnown() bool {
	switch r {
	case BetaThreadNewParamsToolResourcesFileSearchVectorStoresChunkingStrategyTypeAuto, BetaThreadNewParamsToolResourcesFileSearchVectorStoresChunkingStrategyTypeStatic:
		return true
	}
	return false
}

type BetaThreadUpdateParams struct {
	// Set of 16 key-value pairs that can be attached to an object. This can be useful
	// for storing additional information about the object in a structured format. Keys
	// can be a maximum of 64 characters long and values can be a maxium of 512
	// characters long.
	Metadata param.Field[interface{}] `json:"metadata"`
	// A set of resources that are made available to the assistant's tools in this
	// thread. The resources are specific to the type of tool. For example, the
	// `code_interpreter` tool requires a list of file IDs, while the `file_search`
	// tool requires a list of vector store IDs.
	ToolResources param.Field[BetaThreadUpdateParamsToolResources] `json:"tool_resources"`
}

func (r BetaThreadUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A set of resources that are made available to the assistant's tools in this
// thread. The resources are specific to the type of tool. For example, the
// `code_interpreter` tool requires a list of file IDs, while the `file_search`
// tool requires a list of vector store IDs.
type BetaThreadUpdateParamsToolResources struct {
	CodeInterpreter param.Field[BetaThreadUpdateParamsToolResourcesCodeInterpreter] `json:"code_interpreter"`
	FileSearch      param.Field[BetaThreadUpdateParamsToolResourcesFileSearch]      `json:"file_search"`
}

func (r BetaThreadUpdateParamsToolResources) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BetaThreadUpdateParamsToolResourcesCodeInterpreter struct {
	// A list of [file](https://platform.openai.com/docs/api-reference/files) IDs made
	// available to the `code_interpreter` tool. There can be a maximum of 20 files
	// associated with the tool.
	FileIDs param.Field[[]string] `json:"file_ids"`
}

func (r BetaThreadUpdateParamsToolResourcesCodeInterpreter) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BetaThreadUpdateParamsToolResourcesFileSearch struct {
	// The
	// [vector store](https://platform.openai.com/docs/api-reference/vector-stores/object)
	// attached to this thread. There can be a maximum of 1 vector store attached to
	// the thread.
	VectorStoreIDs param.Field[[]string] `json:"vector_store_ids"`
}

func (r BetaThreadUpdateParamsToolResourcesFileSearch) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BetaThreadNewAndRunParams struct {
	// The ID of the
	// [assistant](https://platform.openai.com/docs/api-reference/assistants) to use to
	// execute this run.
	AssistantID param.Field[string] `json:"assistant_id,required"`
	// Override the default system message of the assistant. This is useful for
	// modifying the behavior on a per-run basis.
	Instructions param.Field[string] `json:"instructions"`
	// The maximum number of completion tokens that may be used over the course of the
	// run. The run will make a best effort to use only the number of completion tokens
	// specified, across multiple turns of the run. If the run exceeds the number of
	// completion tokens specified, the run will end with status `incomplete`. See
	// `incomplete_details` for more info.
	MaxCompletionTokens param.Field[int64] `json:"max_completion_tokens"`
	// The maximum number of prompt tokens that may be used over the course of the run.
	// The run will make a best effort to use only the number of prompt tokens
	// specified, across multiple turns of the run. If the run exceeds the number of
	// prompt tokens specified, the run will end with status `incomplete`. See
	// `incomplete_details` for more info.
	MaxPromptTokens param.Field[int64] `json:"max_prompt_tokens"`
	// Set of 16 key-value pairs that can be attached to an object. This can be useful
	// for storing additional information about the object in a structured format. Keys
	// can be a maximum of 64 characters long and values can be a maxium of 512
	// characters long.
	Metadata param.Field[interface{}] `json:"metadata"`
	// The ID of the [Model](https://platform.openai.com/docs/api-reference/models) to
	// be used to execute this run. If a value is provided here, it will override the
	// model associated with the assistant. If not, the model associated with the
	// assistant will be used.
	Model param.Field[BetaThreadNewAndRunParamsModel] `json:"model"`
	// Whether to enable
	// [parallel function calling](https://platform.openai.com/docs/guides/function-calling/parallel-function-calling)
	// during tool use.
	ParallelToolCalls param.Field[bool] `json:"parallel_tool_calls"`
	// Specifies the format that the model must output. Compatible with
	// [GPT-4o](https://platform.openai.com/docs/models/gpt-4o),
	// [GPT-4 Turbo](https://platform.openai.com/docs/models/gpt-4-turbo-and-gpt-4),
	// and all GPT-3.5 Turbo models since `gpt-3.5-turbo-1106`.
	//
	// Setting to `{ "type": "json_object" }` enables JSON mode, which guarantees the
	// message the model generates is valid JSON.
	//
	// **Important:** when using JSON mode, you **must** also instruct the model to
	// produce JSON yourself via a system or user message. Without this, the model may
	// generate an unending stream of whitespace until the generation reaches the token
	// limit, resulting in a long-running and seemingly "stuck" request. Also note that
	// the message content may be partially cut off if `finish_reason="length"`, which
	// indicates the generation exceeded `max_tokens` or the conversation exceeded the
	// max context length.
	ResponseFormat param.Field[AssistantResponseFormatOptionUnionParam] `json:"response_format"`
	// What sampling temperature to use, between 0 and 2. Higher values like 0.8 will
	// make the output more random, while lower values like 0.2 will make it more
	// focused and deterministic.
	Temperature param.Field[float64] `json:"temperature"`
	// If no thread is provided, an empty thread will be created.
	Thread param.Field[BetaThreadNewAndRunParamsThread] `json:"thread"`
	// Controls which (if any) tool is called by the model. `none` means the model will
	// not call any tools and instead generates a message. `auto` is the default value
	// and means the model can pick between generating a message or calling one or more
	// tools. `required` means the model must call one or more tools before responding
	// to the user. Specifying a particular tool like `{"type": "file_search"}` or
	// `{"type": "function", "function": {"name": "my_function"}}` forces the model to
	// call that tool.
	ToolChoice param.Field[AssistantToolChoiceOptionUnionParam] `json:"tool_choice"`
	// A set of resources that are used by the assistant's tools. The resources are
	// specific to the type of tool. For example, the `code_interpreter` tool requires
	// a list of file IDs, while the `file_search` tool requires a list of vector store
	// IDs.
	ToolResources param.Field[BetaThreadNewAndRunParamsToolResources] `json:"tool_resources"`
	// Override the tools the assistant can use for this run. This is useful for
	// modifying the behavior on a per-run basis.
	Tools param.Field[[]BetaThreadNewAndRunParamsToolUnion] `json:"tools"`
	// An alternative to sampling with temperature, called nucleus sampling, where the
	// model considers the results of the tokens with top_p probability mass. So 0.1
	// means only the tokens comprising the top 10% probability mass are considered.
	//
	// We generally recommend altering this or temperature but not both.
	TopP param.Field[float64] `json:"top_p"`
	// Controls for how a thread will be truncated prior to the run. Use this to
	// control the intial context window of the run.
	TruncationStrategy param.Field[BetaThreadNewAndRunParamsTruncationStrategy] `json:"truncation_strategy"`
}

func (r BetaThreadNewAndRunParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BetaThreadNewAndRunParamsModel string

const (
	BetaThreadNewAndRunParamsModelGPT4o               BetaThreadNewAndRunParamsModel = "gpt-4o"
	BetaThreadNewAndRunParamsModelGPT4o2024_05_13     BetaThreadNewAndRunParamsModel = "gpt-4o-2024-05-13"
	BetaThreadNewAndRunParamsModelGPT4oMini           BetaThreadNewAndRunParamsModel = "gpt-4o-mini"
	BetaThreadNewAndRunParamsModelGPT4oMini2024_07_18 BetaThreadNewAndRunParamsModel = "gpt-4o-mini-2024-07-18"
	BetaThreadNewAndRunParamsModelGPT4Turbo           BetaThreadNewAndRunParamsModel = "gpt-4-turbo"
	BetaThreadNewAndRunParamsModelGPT4Turbo2024_04_09 BetaThreadNewAndRunParamsModel = "gpt-4-turbo-2024-04-09"
	BetaThreadNewAndRunParamsModelGPT4_0125Preview    BetaThreadNewAndRunParamsModel = "gpt-4-0125-preview"
	BetaThreadNewAndRunParamsModelGPT4TurboPreview    BetaThreadNewAndRunParamsModel = "gpt-4-turbo-preview"
	BetaThreadNewAndRunParamsModelGPT4_1106Preview    BetaThreadNewAndRunParamsModel = "gpt-4-1106-preview"
	BetaThreadNewAndRunParamsModelGPT4VisionPreview   BetaThreadNewAndRunParamsModel = "gpt-4-vision-preview"
	BetaThreadNewAndRunParamsModelGPT4                BetaThreadNewAndRunParamsModel = "gpt-4"
	BetaThreadNewAndRunParamsModelGPT4_0314           BetaThreadNewAndRunParamsModel = "gpt-4-0314"
	BetaThreadNewAndRunParamsModelGPT4_0613           BetaThreadNewAndRunParamsModel = "gpt-4-0613"
	BetaThreadNewAndRunParamsModelGPT4_32k            BetaThreadNewAndRunParamsModel = "gpt-4-32k"
	BetaThreadNewAndRunParamsModelGPT4_32k0314        BetaThreadNewAndRunParamsModel = "gpt-4-32k-0314"
	BetaThreadNewAndRunParamsModelGPT4_32k0613        BetaThreadNewAndRunParamsModel = "gpt-4-32k-0613"
	BetaThreadNewAndRunParamsModelGPT3_5Turbo         BetaThreadNewAndRunParamsModel = "gpt-3.5-turbo"
	BetaThreadNewAndRunParamsModelGPT3_5Turbo16k      BetaThreadNewAndRunParamsModel = "gpt-3.5-turbo-16k"
	BetaThreadNewAndRunParamsModelGPT3_5Turbo0613     BetaThreadNewAndRunParamsModel = "gpt-3.5-turbo-0613"
	BetaThreadNewAndRunParamsModelGPT3_5Turbo1106     BetaThreadNewAndRunParamsModel = "gpt-3.5-turbo-1106"
	BetaThreadNewAndRunParamsModelGPT3_5Turbo0125     BetaThreadNewAndRunParamsModel = "gpt-3.5-turbo-0125"
	BetaThreadNewAndRunParamsModelGPT3_5Turbo16k0613  BetaThreadNewAndRunParamsModel = "gpt-3.5-turbo-16k-0613"
)

func (r BetaThreadNewAndRunParamsModel) IsKnown() bool {
	switch r {
	case BetaThreadNewAndRunParamsModelGPT4o, BetaThreadNewAndRunParamsModelGPT4o2024_05_13, BetaThreadNewAndRunParamsModelGPT4oMini, BetaThreadNewAndRunParamsModelGPT4oMini2024_07_18, BetaThreadNewAndRunParamsModelGPT4Turbo, BetaThreadNewAndRunParamsModelGPT4Turbo2024_04_09, BetaThreadNewAndRunParamsModelGPT4_0125Preview, BetaThreadNewAndRunParamsModelGPT4TurboPreview, BetaThreadNewAndRunParamsModelGPT4_1106Preview, BetaThreadNewAndRunParamsModelGPT4VisionPreview, BetaThreadNewAndRunParamsModelGPT4, BetaThreadNewAndRunParamsModelGPT4_0314, BetaThreadNewAndRunParamsModelGPT4_0613, BetaThreadNewAndRunParamsModelGPT4_32k, BetaThreadNewAndRunParamsModelGPT4_32k0314, BetaThreadNewAndRunParamsModelGPT4_32k0613, BetaThreadNewAndRunParamsModelGPT3_5Turbo, BetaThreadNewAndRunParamsModelGPT3_5Turbo16k, BetaThreadNewAndRunParamsModelGPT3_5Turbo0613, BetaThreadNewAndRunParamsModelGPT3_5Turbo1106, BetaThreadNewAndRunParamsModelGPT3_5Turbo0125, BetaThreadNewAndRunParamsModelGPT3_5Turbo16k0613:
		return true
	}
	return false
}

// If no thread is provided, an empty thread will be created.
type BetaThreadNewAndRunParamsThread struct {
	// A list of [messages](https://platform.openai.com/docs/api-reference/messages) to
	// start the thread with.
	Messages param.Field[[]BetaThreadNewAndRunParamsThreadMessage] `json:"messages"`
	// Set of 16 key-value pairs that can be attached to an object. This can be useful
	// for storing additional information about the object in a structured format. Keys
	// can be a maximum of 64 characters long and values can be a maxium of 512
	// characters long.
	Metadata param.Field[interface{}] `json:"metadata"`
	// A set of resources that are made available to the assistant's tools in this
	// thread. The resources are specific to the type of tool. For example, the
	// `code_interpreter` tool requires a list of file IDs, while the `file_search`
	// tool requires a list of vector store IDs.
	ToolResources param.Field[BetaThreadNewAndRunParamsThreadToolResources] `json:"tool_resources"`
}

func (r BetaThreadNewAndRunParamsThread) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BetaThreadNewAndRunParamsThreadMessage struct {
	// The text contents of the message.
	Content param.Field[BetaThreadNewAndRunParamsThreadMessagesContentUnion] `json:"content,required"`
	// The role of the entity that is creating the message. Allowed values include:
	//
	//   - `user`: Indicates the message is sent by an actual user and should be used in
	//     most cases to represent user-generated messages.
	//   - `assistant`: Indicates the message is generated by the assistant. Use this
	//     value to insert messages from the assistant into the conversation.
	Role param.Field[BetaThreadNewAndRunParamsThreadMessagesRole] `json:"role,required"`
	// A list of files attached to the message, and the tools they should be added to.
	Attachments param.Field[[]BetaThreadNewAndRunParamsThreadMessagesAttachment] `json:"attachments"`
	// Set of 16 key-value pairs that can be attached to an object. This can be useful
	// for storing additional information about the object in a structured format. Keys
	// can be a maximum of 64 characters long and values can be a maxium of 512
	// characters long.
	Metadata param.Field[interface{}] `json:"metadata"`
}

func (r BetaThreadNewAndRunParamsThreadMessage) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The text contents of the message.
//
// Satisfied by [shared.UnionString],
// [BetaThreadNewAndRunParamsThreadMessagesContentArrayOfContentParts].
type BetaThreadNewAndRunParamsThreadMessagesContentUnion interface {
	ImplementsBetaThreadNewAndRunParamsThreadMessagesContentUnion()
}

type BetaThreadNewAndRunParamsThreadMessagesContentArrayOfContentParts []MessageContentPartParamUnion

func (r BetaThreadNewAndRunParamsThreadMessagesContentArrayOfContentParts) ImplementsBetaThreadNewAndRunParamsThreadMessagesContentUnion() {
}

// The role of the entity that is creating the message. Allowed values include:
//
//   - `user`: Indicates the message is sent by an actual user and should be used in
//     most cases to represent user-generated messages.
//   - `assistant`: Indicates the message is generated by the assistant. Use this
//     value to insert messages from the assistant into the conversation.
type BetaThreadNewAndRunParamsThreadMessagesRole string

const (
	BetaThreadNewAndRunParamsThreadMessagesRoleUser      BetaThreadNewAndRunParamsThreadMessagesRole = "user"
	BetaThreadNewAndRunParamsThreadMessagesRoleAssistant BetaThreadNewAndRunParamsThreadMessagesRole = "assistant"
)

func (r BetaThreadNewAndRunParamsThreadMessagesRole) IsKnown() bool {
	switch r {
	case BetaThreadNewAndRunParamsThreadMessagesRoleUser, BetaThreadNewAndRunParamsThreadMessagesRoleAssistant:
		return true
	}
	return false
}

type BetaThreadNewAndRunParamsThreadMessagesAttachment struct {
	// The ID of the file to attach to the message.
	FileID param.Field[string] `json:"file_id"`
	// The tools to add this file to.
	Tools param.Field[[]BetaThreadNewAndRunParamsThreadMessagesAttachmentsToolUnion] `json:"tools"`
}

func (r BetaThreadNewAndRunParamsThreadMessagesAttachment) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BetaThreadNewAndRunParamsThreadMessagesAttachmentsTool struct {
	// The type of tool being defined: `code_interpreter`
	Type param.Field[BetaThreadNewAndRunParamsThreadMessagesAttachmentsToolsType] `json:"type,required"`
}

func (r BetaThreadNewAndRunParamsThreadMessagesAttachmentsTool) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaThreadNewAndRunParamsThreadMessagesAttachmentsTool) implementsBetaThreadNewAndRunParamsThreadMessagesAttachmentsToolUnion() {
}

// Satisfied by [CodeInterpreterToolParam],
// [BetaThreadNewAndRunParamsThreadMessagesAttachmentsToolsFileSearch],
// [BetaThreadNewAndRunParamsThreadMessagesAttachmentsTool].
type BetaThreadNewAndRunParamsThreadMessagesAttachmentsToolUnion interface {
	implementsBetaThreadNewAndRunParamsThreadMessagesAttachmentsToolUnion()
}

type BetaThreadNewAndRunParamsThreadMessagesAttachmentsToolsFileSearch struct {
	// The type of tool being defined: `file_search`
	Type param.Field[BetaThreadNewAndRunParamsThreadMessagesAttachmentsToolsFileSearchType] `json:"type,required"`
}

func (r BetaThreadNewAndRunParamsThreadMessagesAttachmentsToolsFileSearch) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaThreadNewAndRunParamsThreadMessagesAttachmentsToolsFileSearch) implementsBetaThreadNewAndRunParamsThreadMessagesAttachmentsToolUnion() {
}

// The type of tool being defined: `file_search`
type BetaThreadNewAndRunParamsThreadMessagesAttachmentsToolsFileSearchType string

const (
	BetaThreadNewAndRunParamsThreadMessagesAttachmentsToolsFileSearchTypeFileSearch BetaThreadNewAndRunParamsThreadMessagesAttachmentsToolsFileSearchType = "file_search"
)

func (r BetaThreadNewAndRunParamsThreadMessagesAttachmentsToolsFileSearchType) IsKnown() bool {
	switch r {
	case BetaThreadNewAndRunParamsThreadMessagesAttachmentsToolsFileSearchTypeFileSearch:
		return true
	}
	return false
}

// The type of tool being defined: `code_interpreter`
type BetaThreadNewAndRunParamsThreadMessagesAttachmentsToolsType string

const (
	BetaThreadNewAndRunParamsThreadMessagesAttachmentsToolsTypeCodeInterpreter BetaThreadNewAndRunParamsThreadMessagesAttachmentsToolsType = "code_interpreter"
	BetaThreadNewAndRunParamsThreadMessagesAttachmentsToolsTypeFileSearch      BetaThreadNewAndRunParamsThreadMessagesAttachmentsToolsType = "file_search"
)

func (r BetaThreadNewAndRunParamsThreadMessagesAttachmentsToolsType) IsKnown() bool {
	switch r {
	case BetaThreadNewAndRunParamsThreadMessagesAttachmentsToolsTypeCodeInterpreter, BetaThreadNewAndRunParamsThreadMessagesAttachmentsToolsTypeFileSearch:
		return true
	}
	return false
}

// A set of resources that are made available to the assistant's tools in this
// thread. The resources are specific to the type of tool. For example, the
// `code_interpreter` tool requires a list of file IDs, while the `file_search`
// tool requires a list of vector store IDs.
type BetaThreadNewAndRunParamsThreadToolResources struct {
	CodeInterpreter param.Field[BetaThreadNewAndRunParamsThreadToolResourcesCodeInterpreter] `json:"code_interpreter"`
	FileSearch      param.Field[BetaThreadNewAndRunParamsThreadToolResourcesFileSearch]      `json:"file_search"`
}

func (r BetaThreadNewAndRunParamsThreadToolResources) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BetaThreadNewAndRunParamsThreadToolResourcesCodeInterpreter struct {
	// A list of [file](https://platform.openai.com/docs/api-reference/files) IDs made
	// available to the `code_interpreter` tool. There can be a maximum of 20 files
	// associated with the tool.
	FileIDs param.Field[[]string] `json:"file_ids"`
}

func (r BetaThreadNewAndRunParamsThreadToolResourcesCodeInterpreter) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BetaThreadNewAndRunParamsThreadToolResourcesFileSearch struct {
	// The
	// [vector store](https://platform.openai.com/docs/api-reference/vector-stores/object)
	// attached to this thread. There can be a maximum of 1 vector store attached to
	// the thread.
	VectorStoreIDs param.Field[[]string] `json:"vector_store_ids"`
	// A helper to create a
	// [vector store](https://platform.openai.com/docs/api-reference/vector-stores/object)
	// with file_ids and attach it to this thread. There can be a maximum of 1 vector
	// store attached to the thread.
	VectorStores param.Field[[]BetaThreadNewAndRunParamsThreadToolResourcesFileSearchVectorStore] `json:"vector_stores"`
}

func (r BetaThreadNewAndRunParamsThreadToolResourcesFileSearch) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BetaThreadNewAndRunParamsThreadToolResourcesFileSearchVectorStore struct {
	// The chunking strategy used to chunk the file(s). If not set, will use the `auto`
	// strategy.
	ChunkingStrategy param.Field[BetaThreadNewAndRunParamsThreadToolResourcesFileSearchVectorStoresChunkingStrategyUnion] `json:"chunking_strategy"`
	// A list of [file](https://platform.openai.com/docs/api-reference/files) IDs to
	// add to the vector store. There can be a maximum of 10000 files in a vector
	// store.
	FileIDs param.Field[[]string] `json:"file_ids"`
	// Set of 16 key-value pairs that can be attached to a vector store. This can be
	// useful for storing additional information about the vector store in a structured
	// format. Keys can be a maximum of 64 characters long and values can be a maxium
	// of 512 characters long.
	Metadata param.Field[interface{}] `json:"metadata"`
}

func (r BetaThreadNewAndRunParamsThreadToolResourcesFileSearchVectorStore) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The chunking strategy used to chunk the file(s). If not set, will use the `auto`
// strategy.
type BetaThreadNewAndRunParamsThreadToolResourcesFileSearchVectorStoresChunkingStrategy struct {
	// Always `auto`.
	Type   param.Field[BetaThreadNewAndRunParamsThreadToolResourcesFileSearchVectorStoresChunkingStrategyType] `json:"type,required"`
	Static param.Field[interface{}]                                                                            `json:"static,required"`
}

func (r BetaThreadNewAndRunParamsThreadToolResourcesFileSearchVectorStoresChunkingStrategy) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaThreadNewAndRunParamsThreadToolResourcesFileSearchVectorStoresChunkingStrategy) implementsBetaThreadNewAndRunParamsThreadToolResourcesFileSearchVectorStoresChunkingStrategyUnion() {
}

// The chunking strategy used to chunk the file(s). If not set, will use the `auto`
// strategy.
//
// Satisfied by
// [BetaThreadNewAndRunParamsThreadToolResourcesFileSearchVectorStoresChunkingStrategyAuto],
// [BetaThreadNewAndRunParamsThreadToolResourcesFileSearchVectorStoresChunkingStrategyStatic],
// [BetaThreadNewAndRunParamsThreadToolResourcesFileSearchVectorStoresChunkingStrategy].
type BetaThreadNewAndRunParamsThreadToolResourcesFileSearchVectorStoresChunkingStrategyUnion interface {
	implementsBetaThreadNewAndRunParamsThreadToolResourcesFileSearchVectorStoresChunkingStrategyUnion()
}

// The default strategy. This strategy currently uses a `max_chunk_size_tokens` of
// `800` and `chunk_overlap_tokens` of `400`.
type BetaThreadNewAndRunParamsThreadToolResourcesFileSearchVectorStoresChunkingStrategyAuto struct {
	// Always `auto`.
	Type param.Field[BetaThreadNewAndRunParamsThreadToolResourcesFileSearchVectorStoresChunkingStrategyAutoType] `json:"type,required"`
}

func (r BetaThreadNewAndRunParamsThreadToolResourcesFileSearchVectorStoresChunkingStrategyAuto) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaThreadNewAndRunParamsThreadToolResourcesFileSearchVectorStoresChunkingStrategyAuto) implementsBetaThreadNewAndRunParamsThreadToolResourcesFileSearchVectorStoresChunkingStrategyUnion() {
}

// Always `auto`.
type BetaThreadNewAndRunParamsThreadToolResourcesFileSearchVectorStoresChunkingStrategyAutoType string

const (
	BetaThreadNewAndRunParamsThreadToolResourcesFileSearchVectorStoresChunkingStrategyAutoTypeAuto BetaThreadNewAndRunParamsThreadToolResourcesFileSearchVectorStoresChunkingStrategyAutoType = "auto"
)

func (r BetaThreadNewAndRunParamsThreadToolResourcesFileSearchVectorStoresChunkingStrategyAutoType) IsKnown() bool {
	switch r {
	case BetaThreadNewAndRunParamsThreadToolResourcesFileSearchVectorStoresChunkingStrategyAutoTypeAuto:
		return true
	}
	return false
}

type BetaThreadNewAndRunParamsThreadToolResourcesFileSearchVectorStoresChunkingStrategyStatic struct {
	Static param.Field[BetaThreadNewAndRunParamsThreadToolResourcesFileSearchVectorStoresChunkingStrategyStaticStatic] `json:"static,required"`
	// Always `static`.
	Type param.Field[BetaThreadNewAndRunParamsThreadToolResourcesFileSearchVectorStoresChunkingStrategyStaticType] `json:"type,required"`
}

func (r BetaThreadNewAndRunParamsThreadToolResourcesFileSearchVectorStoresChunkingStrategyStatic) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaThreadNewAndRunParamsThreadToolResourcesFileSearchVectorStoresChunkingStrategyStatic) implementsBetaThreadNewAndRunParamsThreadToolResourcesFileSearchVectorStoresChunkingStrategyUnion() {
}

type BetaThreadNewAndRunParamsThreadToolResourcesFileSearchVectorStoresChunkingStrategyStaticStatic struct {
	// The number of tokens that overlap between chunks. The default value is `400`.
	//
	// Note that the overlap must not exceed half of `max_chunk_size_tokens`.
	ChunkOverlapTokens param.Field[int64] `json:"chunk_overlap_tokens,required"`
	// The maximum number of tokens in each chunk. The default value is `800`. The
	// minimum value is `100` and the maximum value is `4096`.
	MaxChunkSizeTokens param.Field[int64] `json:"max_chunk_size_tokens,required"`
}

func (r BetaThreadNewAndRunParamsThreadToolResourcesFileSearchVectorStoresChunkingStrategyStaticStatic) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Always `static`.
type BetaThreadNewAndRunParamsThreadToolResourcesFileSearchVectorStoresChunkingStrategyStaticType string

const (
	BetaThreadNewAndRunParamsThreadToolResourcesFileSearchVectorStoresChunkingStrategyStaticTypeStatic BetaThreadNewAndRunParamsThreadToolResourcesFileSearchVectorStoresChunkingStrategyStaticType = "static"
)

func (r BetaThreadNewAndRunParamsThreadToolResourcesFileSearchVectorStoresChunkingStrategyStaticType) IsKnown() bool {
	switch r {
	case BetaThreadNewAndRunParamsThreadToolResourcesFileSearchVectorStoresChunkingStrategyStaticTypeStatic:
		return true
	}
	return false
}

// Always `auto`.
type BetaThreadNewAndRunParamsThreadToolResourcesFileSearchVectorStoresChunkingStrategyType string

const (
	BetaThreadNewAndRunParamsThreadToolResourcesFileSearchVectorStoresChunkingStrategyTypeAuto   BetaThreadNewAndRunParamsThreadToolResourcesFileSearchVectorStoresChunkingStrategyType = "auto"
	BetaThreadNewAndRunParamsThreadToolResourcesFileSearchVectorStoresChunkingStrategyTypeStatic BetaThreadNewAndRunParamsThreadToolResourcesFileSearchVectorStoresChunkingStrategyType = "static"
)

func (r BetaThreadNewAndRunParamsThreadToolResourcesFileSearchVectorStoresChunkingStrategyType) IsKnown() bool {
	switch r {
	case BetaThreadNewAndRunParamsThreadToolResourcesFileSearchVectorStoresChunkingStrategyTypeAuto, BetaThreadNewAndRunParamsThreadToolResourcesFileSearchVectorStoresChunkingStrategyTypeStatic:
		return true
	}
	return false
}

// A set of resources that are used by the assistant's tools. The resources are
// specific to the type of tool. For example, the `code_interpreter` tool requires
// a list of file IDs, while the `file_search` tool requires a list of vector store
// IDs.
type BetaThreadNewAndRunParamsToolResources struct {
	CodeInterpreter param.Field[BetaThreadNewAndRunParamsToolResourcesCodeInterpreter] `json:"code_interpreter"`
	FileSearch      param.Field[BetaThreadNewAndRunParamsToolResourcesFileSearch]      `json:"file_search"`
}

func (r BetaThreadNewAndRunParamsToolResources) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BetaThreadNewAndRunParamsToolResourcesCodeInterpreter struct {
	// A list of [file](https://platform.openai.com/docs/api-reference/files) IDs made
	// available to the `code_interpreter` tool. There can be a maximum of 20 files
	// associated with the tool.
	FileIDs param.Field[[]string] `json:"file_ids"`
}

func (r BetaThreadNewAndRunParamsToolResourcesCodeInterpreter) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BetaThreadNewAndRunParamsToolResourcesFileSearch struct {
	// The ID of the
	// [vector store](https://platform.openai.com/docs/api-reference/vector-stores/object)
	// attached to this assistant. There can be a maximum of 1 vector store attached to
	// the assistant.
	VectorStoreIDs param.Field[[]string] `json:"vector_store_ids"`
}

func (r BetaThreadNewAndRunParamsToolResourcesFileSearch) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BetaThreadNewAndRunParamsTool struct {
	// The type of tool being defined: `code_interpreter`
	Type       param.Field[BetaThreadNewAndRunParamsToolsType] `json:"type,required"`
	FileSearch param.Field[interface{}]                        `json:"file_search,required"`
	Function   param.Field[shared.FunctionDefinitionParam]     `json:"function"`
}

func (r BetaThreadNewAndRunParamsTool) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaThreadNewAndRunParamsTool) implementsBetaThreadNewAndRunParamsToolUnion() {}

// Satisfied by [CodeInterpreterToolParam], [FileSearchToolParam],
// [FunctionToolParam], [BetaThreadNewAndRunParamsTool].
type BetaThreadNewAndRunParamsToolUnion interface {
	implementsBetaThreadNewAndRunParamsToolUnion()
}

// The type of tool being defined: `code_interpreter`
type BetaThreadNewAndRunParamsToolsType string

const (
	BetaThreadNewAndRunParamsToolsTypeCodeInterpreter BetaThreadNewAndRunParamsToolsType = "code_interpreter"
	BetaThreadNewAndRunParamsToolsTypeFileSearch      BetaThreadNewAndRunParamsToolsType = "file_search"
	BetaThreadNewAndRunParamsToolsTypeFunction        BetaThreadNewAndRunParamsToolsType = "function"
)

func (r BetaThreadNewAndRunParamsToolsType) IsKnown() bool {
	switch r {
	case BetaThreadNewAndRunParamsToolsTypeCodeInterpreter, BetaThreadNewAndRunParamsToolsTypeFileSearch, BetaThreadNewAndRunParamsToolsTypeFunction:
		return true
	}
	return false
}

// Controls for how a thread will be truncated prior to the run. Use this to
// control the intial context window of the run.
type BetaThreadNewAndRunParamsTruncationStrategy struct {
	// The truncation strategy to use for the thread. The default is `auto`. If set to
	// `last_messages`, the thread will be truncated to the n most recent messages in
	// the thread. When set to `auto`, messages in the middle of the thread will be
	// dropped to fit the context length of the model, `max_prompt_tokens`.
	Type param.Field[BetaThreadNewAndRunParamsTruncationStrategyType] `json:"type,required"`
	// The number of most recent messages from the thread when constructing the context
	// for the run.
	LastMessages param.Field[int64] `json:"last_messages"`
}

func (r BetaThreadNewAndRunParamsTruncationStrategy) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The truncation strategy to use for the thread. The default is `auto`. If set to
// `last_messages`, the thread will be truncated to the n most recent messages in
// the thread. When set to `auto`, messages in the middle of the thread will be
// dropped to fit the context length of the model, `max_prompt_tokens`.
type BetaThreadNewAndRunParamsTruncationStrategyType string

const (
	BetaThreadNewAndRunParamsTruncationStrategyTypeAuto         BetaThreadNewAndRunParamsTruncationStrategyType = "auto"
	BetaThreadNewAndRunParamsTruncationStrategyTypeLastMessages BetaThreadNewAndRunParamsTruncationStrategyType = "last_messages"
)

func (r BetaThreadNewAndRunParamsTruncationStrategyType) IsKnown() bool {
	switch r {
	case BetaThreadNewAndRunParamsTruncationStrategyTypeAuto, BetaThreadNewAndRunParamsTruncationStrategyTypeLastMessages:
		return true
	}
	return false
}