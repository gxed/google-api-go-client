// Package firestore provides access to the Cloud Firestore API.
//
// This package is DEPRECATED. Use package cloud.google.com/go/firestore instead.
//
// See https://cloud.google.com/firestore
//
// Usage example:
//
//   import "github.com/gxed/google-api-go-client/firestore/v1beta2"
//   ...
//   firestoreService, err := firestore.New(oauthHttpClient)
package firestore

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	context "golang.org/x/net/context"
	ctxhttp "golang.org/x/net/context/ctxhttp"
	gensupport "github.com/gxed/google-api-go-client/gensupport"
	googleapi "github.com/gxed/google-api-go-client/googleapi"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// Always reference these packages, just in case the auto-generated code
// below doesn't.
var _ = bytes.NewBuffer
var _ = strconv.Itoa
var _ = fmt.Sprintf
var _ = json.NewDecoder
var _ = io.Copy
var _ = url.Parse
var _ = gensupport.MarshalJSON
var _ = googleapi.Version
var _ = errors.New
var _ = strings.Replace
var _ = context.Canceled
var _ = ctxhttp.Do

const apiId = "firestore:v1beta2"
const apiName = "firestore"
const apiVersion = "v1beta2"
const basePath = "https://firestore.googleapis.com/"

// OAuth2 scopes used by this API.
const (
	// View and manage your data across Google Cloud Platform services
	CloudPlatformScope = "https://www.googleapis.com/auth/cloud-platform"

	// View and manage your Google Cloud Datastore data
	DatastoreScope = "https://www.googleapis.com/auth/datastore"
)

func New(client *http.Client) (*Service, error) {
	if client == nil {
		return nil, errors.New("client is nil")
	}
	s := &Service{client: client, BasePath: basePath}
	s.Projects = NewProjectsService(s)
	return s, nil
}

type Service struct {
	client    *http.Client
	BasePath  string // API endpoint base URL
	UserAgent string // optional additional User-Agent fragment

	Projects *ProjectsService
}

func (s *Service) userAgent() string {
	if s.UserAgent == "" {
		return googleapi.UserAgent
	}
	return googleapi.UserAgent + " " + s.UserAgent
}

func NewProjectsService(s *Service) *ProjectsService {
	rs := &ProjectsService{s: s}
	rs.Databases = NewProjectsDatabasesService(s)
	return rs
}

type ProjectsService struct {
	s *Service

	Databases *ProjectsDatabasesService
}

func NewProjectsDatabasesService(s *Service) *ProjectsDatabasesService {
	rs := &ProjectsDatabasesService{s: s}
	rs.CollectionGroup = NewProjectsDatabasesCollectionGroupService(s)
	rs.CollectionGroups = NewProjectsDatabasesCollectionGroupsService(s)
	return rs
}

type ProjectsDatabasesService struct {
	s *Service

	CollectionGroup *ProjectsDatabasesCollectionGroupService

	CollectionGroups *ProjectsDatabasesCollectionGroupsService
}

func NewProjectsDatabasesCollectionGroupService(s *Service) *ProjectsDatabasesCollectionGroupService {
	rs := &ProjectsDatabasesCollectionGroupService{s: s}
	rs.Indexes = NewProjectsDatabasesCollectionGroupIndexesService(s)
	return rs
}

type ProjectsDatabasesCollectionGroupService struct {
	s *Service

	Indexes *ProjectsDatabasesCollectionGroupIndexesService
}

func NewProjectsDatabasesCollectionGroupIndexesService(s *Service) *ProjectsDatabasesCollectionGroupIndexesService {
	rs := &ProjectsDatabasesCollectionGroupIndexesService{s: s}
	return rs
}

type ProjectsDatabasesCollectionGroupIndexesService struct {
	s *Service
}

func NewProjectsDatabasesCollectionGroupsService(s *Service) *ProjectsDatabasesCollectionGroupsService {
	rs := &ProjectsDatabasesCollectionGroupsService{s: s}
	rs.Fields = NewProjectsDatabasesCollectionGroupsFieldsService(s)
	rs.Indexes = NewProjectsDatabasesCollectionGroupsIndexesService(s)
	return rs
}

type ProjectsDatabasesCollectionGroupsService struct {
	s *Service

	Fields *ProjectsDatabasesCollectionGroupsFieldsService

	Indexes *ProjectsDatabasesCollectionGroupsIndexesService
}

func NewProjectsDatabasesCollectionGroupsFieldsService(s *Service) *ProjectsDatabasesCollectionGroupsFieldsService {
	rs := &ProjectsDatabasesCollectionGroupsFieldsService{s: s}
	return rs
}

type ProjectsDatabasesCollectionGroupsFieldsService struct {
	s *Service
}

func NewProjectsDatabasesCollectionGroupsIndexesService(s *Service) *ProjectsDatabasesCollectionGroupsIndexesService {
	rs := &ProjectsDatabasesCollectionGroupsIndexesService{s: s}
	return rs
}

type ProjectsDatabasesCollectionGroupsIndexesService struct {
	s *Service
}

// Empty: A generic empty message that you can re-use to avoid defining
// duplicated
// empty messages in your APIs. A typical example is to use it as the
// request
// or the response type of an API method. For instance:
//
//     service Foo {
//       rpc Bar(google.protobuf.Empty) returns
// (google.protobuf.Empty);
//     }
//
// The JSON representation for `Empty` is empty JSON object `{}`.
type Empty struct {
	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`
}

// GoogleFirestoreAdminV1beta2Field: Represents a single field in the
// database.
//
// Fields are grouped by their "Collection Group", which represent
// all
// collections in the database with the same id.
type GoogleFirestoreAdminV1beta2Field struct {
	// IndexConfig: The index configuration for this field. If unset, field
	// indexing will
	// revert to the configuration defined by the `ancestor_field`.
	// To
	// explicitly remove all indexes for this field, specify an index
	// config
	// with an empty list of indexes.
	IndexConfig *GoogleFirestoreAdminV1beta2IndexConfig `json:"indexConfig,omitempty"`

	// Name: A field name of the
	// form
	// `projects/{project_id}/databases/{database_id}/collectionGroups/{
	// collection_id}/fields/{field_path}`
	//
	// A field path may be a simple field name, e.g. `address` or a path to
	// fields
	// within map_value , e.g. `address.city`,
	// or a special field path. The only valid special field is `*`,
	// which
	// represents any field.
	//
	// Field paths may be quoted using ` (backtick). The only character that
	// needs
	// to be escaped within a quoted field path is the backtick character
	// itself,
	// escaped using a backslash. Special characters in field paths
	// that
	// must be quoted include: `*`, `.`,
	// ``` (backtick), `[`, `]`, as well as any ascii symbolic
	// characters.
	//
	// Examples:
	// (Note: Comments here are written in markdown syntax, so there is an
	//  additional layer of backticks to represent a code
	// block)
	// `\`address.city\`` represents a field named `address.city`, not the
	// map key
	// `city` in the field `address`.
	// `\`*\`` represents a field named `*`, not any field.
	//
	// A special `Field` contains the default indexing settings for all
	// fields.
	// This field's resource name
	// is:
	// `projects/{project_id}/databases/{database_id}/collectionGroups/__
	// default__/fields/*`
	// Indexes defined on this `Field` will be applied to all fields which
	// do not
	// have their own `Field` index configuration.
	Name string `json:"name,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "IndexConfig") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "IndexConfig") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleFirestoreAdminV1beta2Field) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleFirestoreAdminV1beta2Field
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleFirestoreAdminV1beta2FieldOperationMetadata: Metadata for
// google.longrunning.Operation results from
// FirestoreAdmin.UpdateField.
type GoogleFirestoreAdminV1beta2FieldOperationMetadata struct {
	// BytesProgress: The progress, in bytes, of this operation.
	BytesProgress *GoogleFirestoreAdminV1beta2Progress `json:"bytesProgress,omitempty"`

	// DocumentProgress: The progress, in documents, of this operation.
	DocumentProgress *GoogleFirestoreAdminV1beta2Progress `json:"documentProgress,omitempty"`

	// EndTime: The time this operation completed. Will be unset if
	// operation still in
	// progress.
	EndTime string `json:"endTime,omitempty"`

	// Field: The field resource that this operation is acting on. For
	// example:
	// `projects/{project_id}/databases/{database_id}/collectionGrou
	// ps/{collection_id}/fields/{field_path}`
	Field string `json:"field,omitempty"`

	// IndexConfigDeltas: A list of IndexConfigDelta, which describe the
	// intent of this
	// operation.
	IndexConfigDeltas []*GoogleFirestoreAdminV1beta2IndexConfigDelta `json:"indexConfigDeltas,omitempty"`

	// StartTime: The time this operation started.
	StartTime string `json:"startTime,omitempty"`

	// State: The state of the operation.
	//
	// Possible values:
	//   "OPERATION_STATE_UNSPECIFIED" - Unspecified.
	//   "INITIALIZING" - Request is being prepared for processing.
	//   "PROCESSING" - Request is actively being processed.
	//   "CANCELLING" - Request is in the process of being cancelled after
	// user called
	// google.longrunning.Operations.CancelOperation on the operation.
	//   "FINALIZING" - Request has been processed and is in its
	// finalization stage.
	//   "SUCCESSFUL" - Request has completed successfully.
	//   "FAILED" - Request has finished being processed, but encountered an
	// error.
	//   "CANCELLED" - Request has finished being cancelled after user
	// called
	// google.longrunning.Operations.CancelOperation.
	State string `json:"state,omitempty"`

	// ForceSendFields is a list of field names (e.g. "BytesProgress") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "BytesProgress") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleFirestoreAdminV1beta2FieldOperationMetadata) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleFirestoreAdminV1beta2FieldOperationMetadata
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleFirestoreAdminV1beta2Index: Cloud Firestore indexes enable
// simple and complex queries against
// documents in a database.
type GoogleFirestoreAdminV1beta2Index struct {
	// Fields: The fields supported by this index.
	//
	// For composite indexes, this is always 2 or more fields.
	// The last field entry is always for the field path `__name__`. If,
	// on
	// creation, `__name__` was not specified as the last field, it will be
	// added
	// automatically with the same direction as that of the last field
	// defined. If
	// the final field in a composite index is not directional, the
	// `__name__`
	// will be ordered ASCENDING (unless explicitly specified).
	//
	// For single field indexes, this will always be exactly one entry with
	// a
	// field path equal to the field path of the associated field.
	Fields []*GoogleFirestoreAdminV1beta2IndexField `json:"fields,omitempty"`

	// Name: Output only.
	// A server defined name for this index.
	// The form of this name for composite indexes will
	// be:
	// `projects/{project_id}/databases/{database_id}/collectionGroups/{c
	// ollection_id}/indexes/{composite_index_id}`
	// For single field indexes, this field will be empty.
	Name string `json:"name,omitempty"`

	// QueryScope: Indexes with a collection query scope specified allow
	// queries
	// against a collection that is the child of a specific document,
	// specified at
	// query time, and that has the same collection id.
	//
	// Indexes with a collection group query scope specified allow queries
	// against
	// all collections descended from a specific document, specified at
	// query
	// time, and that have the same collection id as this index.
	//
	// Possible values:
	//   "QUERY_SCOPE_UNSPECIFIED" - The query scope is unspecified. Not a
	// valid option.
	//   "COLLECTION" - Indexes with a collection query scope specified
	// allow queries
	// against a collection that is the child of a specific document,
	// specified
	// at query time, and that has the collection id specified by the index.
	QueryScope string `json:"queryScope,omitempty"`

	// State: Output only.
	// The serving state of the index.
	//
	// Possible values:
	//   "STATE_UNSPECIFIED" - The state is unspecified.
	//   "CREATING" - The index is being created.
	// There is an active long-running operation for the index.
	// The index is updated when writing a document.
	// Some index data may exist.
	//   "READY" - The index is ready to be used.
	// The index is updated when writing a document.
	// The index is fully populated from all stored documents it applies to.
	//   "NEEDS_REPAIR" - The index was being created, but something went
	// wrong.
	// There is no active long-running operation for the index,
	// and the most recently finished long-running operation failed.
	// The index is not updated when writing a document.
	// Some index data may exist.
	// Use the google.longrunning.Operations API to determine why the
	// operation
	// that last attempted to create this index failed, then re-create
	// the
	// index.
	State string `json:"state,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Fields") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Fields") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleFirestoreAdminV1beta2Index) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleFirestoreAdminV1beta2Index
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleFirestoreAdminV1beta2IndexConfig: The index configuration for
// this field.
type GoogleFirestoreAdminV1beta2IndexConfig struct {
	// AncestorField: Output only.
	// Specifies the resource name of the `Field` from which this
	// field's
	// index configuration is set (when `uses_ancestor_config` is true),
	// or from which it *would* be set if this field had no index
	// configuration
	// (when `uses_ancestor_config` is false).
	AncestorField string `json:"ancestorField,omitempty"`

	// Indexes: The indexes supported for this field.
	Indexes []*GoogleFirestoreAdminV1beta2Index `json:"indexes,omitempty"`

	// Reverting: Output only
	// When true, the `Field`'s index configuration is in the process of
	// being
	// reverted. Once complete, the index config will transition to the
	// same
	// state as the field specified by `ancestor_field`, at which
	// point
	// `uses_ancestor_config` will be `true` and `reverting` will be
	// `false`.
	Reverting bool `json:"reverting,omitempty"`

	// UsesAncestorConfig: Output only.
	// When true, the `Field`'s index configuration is set from
	// the
	// configuration specified by the `ancestor_field`.
	// When false, the `Field`'s index configuration is defined explicitly.
	UsesAncestorConfig bool `json:"usesAncestorConfig,omitempty"`

	// ForceSendFields is a list of field names (e.g. "AncestorField") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "AncestorField") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleFirestoreAdminV1beta2IndexConfig) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleFirestoreAdminV1beta2IndexConfig
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleFirestoreAdminV1beta2IndexConfigDelta: Information about an
// index configuration change.
type GoogleFirestoreAdminV1beta2IndexConfigDelta struct {
	// ChangeType: Specifies how the index is changing.
	//
	// Possible values:
	//   "CHANGE_TYPE_UNSPECIFIED" - The type of change is not specified or
	// known.
	//   "ADD" - The single field index is being added.
	//   "REMOVE" - The single field index is being removed.
	ChangeType string `json:"changeType,omitempty"`

	// Index: The index being changed.
	Index *GoogleFirestoreAdminV1beta2Index `json:"index,omitempty"`

	// ForceSendFields is a list of field names (e.g. "ChangeType") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "ChangeType") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleFirestoreAdminV1beta2IndexConfigDelta) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleFirestoreAdminV1beta2IndexConfigDelta
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleFirestoreAdminV1beta2IndexField: A field in an index.
// The field_path describes which field is indexed, the value_mode
// describes
// how the field value is indexed.
type GoogleFirestoreAdminV1beta2IndexField struct {
	// ArrayConfig: Indicates that this field supports operations on
	// `array_value`s.
	//
	// Possible values:
	//   "ARRAY_CONFIG_UNSPECIFIED" - The index does not support additional
	// array queries.
	//   "CONTAINS" - The index supports array containment queries.
	ArrayConfig string `json:"arrayConfig,omitempty"`

	// FieldPath: Can be __name__.
	// For single field indexes, this must match the name of the field or
	// may
	// be omitted.
	FieldPath string `json:"fieldPath,omitempty"`

	// Order: Indicates that this field supports ordering by the specified
	// order or
	// comparing using =, <, <=, >, >=.
	//
	// Possible values:
	//   "ORDER_UNSPECIFIED" - The ordering is unspecified. Not a valid
	// option.
	//   "ASCENDING" - The field is ordered by ascending field value.
	//   "DESCENDING" - The field is ordered by descending field value.
	Order string `json:"order,omitempty"`

	// ForceSendFields is a list of field names (e.g. "ArrayConfig") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "ArrayConfig") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleFirestoreAdminV1beta2IndexField) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleFirestoreAdminV1beta2IndexField
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleFirestoreAdminV1beta2IndexOperationMetadata: Metadata for
// google.longrunning.Operation results from
// FirestoreAdmin.CreateIndex.
type GoogleFirestoreAdminV1beta2IndexOperationMetadata struct {
	// EndTime: The time this operation completed. Will be unset if
	// operation still in
	// progress.
	EndTime string `json:"endTime,omitempty"`

	// Index: The index resource that this operation is acting on. For
	// example:
	// `projects/{project_id}/databases/{database_id}/collectionGrou
	// ps/{collection_id}/indexes/{index_id}`
	Index string `json:"index,omitempty"`

	// ProgressBytes: The progress, in bytes, of this operation.
	ProgressBytes *GoogleFirestoreAdminV1beta2Progress `json:"progressBytes,omitempty"`

	// ProgressDocuments: The progress, in documents, of this operation.
	ProgressDocuments *GoogleFirestoreAdminV1beta2Progress `json:"progressDocuments,omitempty"`

	// StartTime: The time this operation started.
	StartTime string `json:"startTime,omitempty"`

	// State: The state of the operation.
	//
	// Possible values:
	//   "OPERATION_STATE_UNSPECIFIED" - Unspecified.
	//   "INITIALIZING" - Request is being prepared for processing.
	//   "PROCESSING" - Request is actively being processed.
	//   "CANCELLING" - Request is in the process of being cancelled after
	// user called
	// google.longrunning.Operations.CancelOperation on the operation.
	//   "FINALIZING" - Request has been processed and is in its
	// finalization stage.
	//   "SUCCESSFUL" - Request has completed successfully.
	//   "FAILED" - Request has finished being processed, but encountered an
	// error.
	//   "CANCELLED" - Request has finished being cancelled after user
	// called
	// google.longrunning.Operations.CancelOperation.
	State string `json:"state,omitempty"`

	// ForceSendFields is a list of field names (e.g. "EndTime") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "EndTime") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleFirestoreAdminV1beta2IndexOperationMetadata) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleFirestoreAdminV1beta2IndexOperationMetadata
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleFirestoreAdminV1beta2ListFieldsResponse: The response for
// FirestoreAdmin.ListFields.
type GoogleFirestoreAdminV1beta2ListFieldsResponse struct {
	// Fields: The requested fields.
	Fields []*GoogleFirestoreAdminV1beta2Field `json:"fields,omitempty"`

	// NextPageToken: A page token that may be used to request another page
	// of results. If blank,
	// this is the last page.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Fields") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Fields") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleFirestoreAdminV1beta2ListFieldsResponse) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleFirestoreAdminV1beta2ListFieldsResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleFirestoreAdminV1beta2ListIndexesResponse: The response for
// FirestoreAdmin.ListIndexes.
type GoogleFirestoreAdminV1beta2ListIndexesResponse struct {
	// Indexes: The requested indexes.
	Indexes []*GoogleFirestoreAdminV1beta2Index `json:"indexes,omitempty"`

	// NextPageToken: A page token that may be used to request another page
	// of results. If blank,
	// this is the last page.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Indexes") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Indexes") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleFirestoreAdminV1beta2ListIndexesResponse) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleFirestoreAdminV1beta2ListIndexesResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleFirestoreAdminV1beta2Progress: Describes the progress of the
// operation.
// Unit of work is generic and must be interpreted based on where
// Progress
// is used.
type GoogleFirestoreAdminV1beta2Progress struct {
	// CompletedWork: The amount of work completed.
	CompletedWork int64 `json:"completedWork,omitempty,string"`

	// EstimatedWork: The amount of work estimated.
	EstimatedWork int64 `json:"estimatedWork,omitempty,string"`

	// ForceSendFields is a list of field names (e.g. "CompletedWork") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "CompletedWork") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleFirestoreAdminV1beta2Progress) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleFirestoreAdminV1beta2Progress
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleLongrunningOperation: This resource represents a long-running
// operation that is the result of a
// network API call.
type GoogleLongrunningOperation struct {
	// Done: If the value is `false`, it means the operation is still in
	// progress.
	// If `true`, the operation is completed, and either `error` or
	// `response` is
	// available.
	Done bool `json:"done,omitempty"`

	// Error: The error result of the operation in case of failure or
	// cancellation.
	Error *Status `json:"error,omitempty"`

	// Metadata: Service-specific metadata associated with the operation.
	// It typically
	// contains progress information and common metadata such as create
	// time.
	// Some services might not provide such metadata.  Any method that
	// returns a
	// long-running operation should document the metadata type, if any.
	Metadata googleapi.RawMessage `json:"metadata,omitempty"`

	// Name: The server-assigned name, which is only unique within the same
	// service that
	// originally returns it. If you use the default HTTP mapping,
	// the
	// `name` should have the format of `operations/some/unique/name`.
	Name string `json:"name,omitempty"`

	// Response: The normal response of the operation in case of success.
	// If the original
	// method returns no data on success, such as `Delete`, the response
	// is
	// `google.protobuf.Empty`.  If the original method is
	// standard
	// `Get`/`Create`/`Update`, the response should be the resource.  For
	// other
	// methods, the response should have the type `XxxResponse`, where
	// `Xxx`
	// is the original method name.  For example, if the original method
	// name
	// is `TakeSnapshot()`, the inferred response type
	// is
	// `TakeSnapshotResponse`.
	Response googleapi.RawMessage `json:"response,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Done") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Done") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleLongrunningOperation) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleLongrunningOperation
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Status: The `Status` type defines a logical error model that is
// suitable for different
// programming environments, including REST APIs and RPC APIs. It is
// used by
// [gRPC](https://github.com/grpc). The error model is designed to
// be:
//
// - Simple to use and understand for most users
// - Flexible enough to meet unexpected needs
//
// # Overview
//
// The `Status` message contains three pieces of data: error code, error
// message,
// and error details. The error code should be an enum value
// of
// google.rpc.Code, but it may accept additional error codes if needed.
// The
// error message should be a developer-facing English message that
// helps
// developers *understand* and *resolve* the error. If a localized
// user-facing
// error message is needed, put the localized message in the error
// details or
// localize it in the client. The optional error details may contain
// arbitrary
// information about the error. There is a predefined set of error
// detail types
// in the package `google.rpc` that can be used for common error
// conditions.
//
// # Language mapping
//
// The `Status` message is the logical representation of the error
// model, but it
// is not necessarily the actual wire format. When the `Status` message
// is
// exposed in different client libraries and different wire protocols,
// it can be
// mapped differently. For example, it will likely be mapped to some
// exceptions
// in Java, but more likely mapped to some error codes in C.
//
// # Other uses
//
// The error model and the `Status` message can be used in a variety
// of
// environments, either with or without APIs, to provide a
// consistent developer experience across different
// environments.
//
// Example uses of this error model include:
//
// - Partial errors. If a service needs to return partial errors to the
// client,
//     it may embed the `Status` in the normal response to indicate the
// partial
//     errors.
//
// - Workflow errors. A typical workflow has multiple steps. Each step
// may
//     have a `Status` message for error reporting.
//
// - Batch operations. If a client uses batch request and batch
// response, the
//     `Status` message should be used directly inside batch response,
// one for
//     each error sub-response.
//
// - Asynchronous operations. If an API call embeds asynchronous
// operation
//     results in its response, the status of those operations should
// be
//     represented directly using the `Status` message.
//
// - Logging. If some API errors are stored in logs, the message
// `Status` could
//     be used directly after any stripping needed for security/privacy
// reasons.
type Status struct {
	// Code: The status code, which should be an enum value of
	// google.rpc.Code.
	Code int64 `json:"code,omitempty"`

	// Details: A list of messages that carry the error details.  There is a
	// common set of
	// message types for APIs to use.
	Details []googleapi.RawMessage `json:"details,omitempty"`

	// Message: A developer-facing error message, which should be in
	// English. Any
	// user-facing error message should be localized and sent in
	// the
	// google.rpc.Status.details field, or localized by the client.
	Message string `json:"message,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Code") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Code") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Status) MarshalJSON() ([]byte, error) {
	type NoMethod Status
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// method id "firestore.projects.databases.collectionGroup.indexes.create":

type ProjectsDatabasesCollectionGroupIndexesCreateCall struct {
	s                                *Service
	parent                           string
	googlefirestoreadminv1beta2index *GoogleFirestoreAdminV1beta2Index
	urlParams_                       gensupport.URLParams
	ctx_                             context.Context
	header_                          http.Header
}

// Create: Creates a composite index. This returns a
// google.longrunning.Operation
// which may be used to track the status of the creation. The metadata
// for
// the operation will be the type IndexOperationMetadata.
func (r *ProjectsDatabasesCollectionGroupIndexesService) Create(parent string, googlefirestoreadminv1beta2index *GoogleFirestoreAdminV1beta2Index) *ProjectsDatabasesCollectionGroupIndexesCreateCall {
	c := &ProjectsDatabasesCollectionGroupIndexesCreateCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.parent = parent
	c.googlefirestoreadminv1beta2index = googlefirestoreadminv1beta2index
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsDatabasesCollectionGroupIndexesCreateCall) Fields(s ...googleapi.Field) *ProjectsDatabasesCollectionGroupIndexesCreateCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsDatabasesCollectionGroupIndexesCreateCall) Context(ctx context.Context) *ProjectsDatabasesCollectionGroupIndexesCreateCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsDatabasesCollectionGroupIndexesCreateCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsDatabasesCollectionGroupIndexesCreateCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.googlefirestoreadminv1beta2index)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta2/{+parent}/indexes")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"parent": c.parent,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "firestore.projects.databases.collectionGroup.indexes.create" call.
// Exactly one of *GoogleLongrunningOperation or error will be non-nil.
// Any non-2xx status code is an error. Response headers are in either
// *GoogleLongrunningOperation.ServerResponse.Header or (if a response
// was returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ProjectsDatabasesCollectionGroupIndexesCreateCall) Do(opts ...googleapi.CallOption) (*GoogleLongrunningOperation, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &GoogleLongrunningOperation{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Creates a composite index. This returns a google.longrunning.Operation\nwhich may be used to track the status of the creation. The metadata for\nthe operation will be the type IndexOperationMetadata.",
	//   "flatPath": "v1beta2/projects/{projectsId}/databases/{databasesId}/collectionGroup/{collectionGroupId}/indexes",
	//   "httpMethod": "POST",
	//   "id": "firestore.projects.databases.collectionGroup.indexes.create",
	//   "parameterOrder": [
	//     "parent"
	//   ],
	//   "parameters": {
	//     "parent": {
	//       "description": "A parent name of the form\n`projects/{project_id}/databases/{database_id}/collectionGroups/{collection_id}`",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/databases/[^/]+/collectionGroup/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta2/{+parent}/indexes",
	//   "request": {
	//     "$ref": "GoogleFirestoreAdminV1beta2Index"
	//   },
	//   "response": {
	//     "$ref": "GoogleLongrunningOperation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/datastore"
	//   ]
	// }

}

// method id "firestore.projects.databases.collectionGroup.indexes.list":

type ProjectsDatabasesCollectionGroupIndexesListCall struct {
	s            *Service
	parent       string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// List: Lists composite indexes.
func (r *ProjectsDatabasesCollectionGroupIndexesService) List(parent string) *ProjectsDatabasesCollectionGroupIndexesListCall {
	c := &ProjectsDatabasesCollectionGroupIndexesListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.parent = parent
	return c
}

// Filter sets the optional parameter "filter": The filter to apply to
// list results.
func (c *ProjectsDatabasesCollectionGroupIndexesListCall) Filter(filter string) *ProjectsDatabasesCollectionGroupIndexesListCall {
	c.urlParams_.Set("filter", filter)
	return c
}

// PageSize sets the optional parameter "pageSize": The number of
// results to return.
func (c *ProjectsDatabasesCollectionGroupIndexesListCall) PageSize(pageSize int64) *ProjectsDatabasesCollectionGroupIndexesListCall {
	c.urlParams_.Set("pageSize", fmt.Sprint(pageSize))
	return c
}

// PageToken sets the optional parameter "pageToken": A page token,
// returned from a previous call to
// FirestoreAdmin.ListIndexes, that may be used to get the next
// page of results.
func (c *ProjectsDatabasesCollectionGroupIndexesListCall) PageToken(pageToken string) *ProjectsDatabasesCollectionGroupIndexesListCall {
	c.urlParams_.Set("pageToken", pageToken)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsDatabasesCollectionGroupIndexesListCall) Fields(s ...googleapi.Field) *ProjectsDatabasesCollectionGroupIndexesListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ProjectsDatabasesCollectionGroupIndexesListCall) IfNoneMatch(entityTag string) *ProjectsDatabasesCollectionGroupIndexesListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsDatabasesCollectionGroupIndexesListCall) Context(ctx context.Context) *ProjectsDatabasesCollectionGroupIndexesListCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsDatabasesCollectionGroupIndexesListCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsDatabasesCollectionGroupIndexesListCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta2/{+parent}/indexes")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"parent": c.parent,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "firestore.projects.databases.collectionGroup.indexes.list" call.
// Exactly one of *GoogleFirestoreAdminV1beta2ListIndexesResponse or
// error will be non-nil. Any non-2xx status code is an error. Response
// headers are in either
// *GoogleFirestoreAdminV1beta2ListIndexesResponse.ServerResponse.Header
// or (if a response was returned at all) in
// error.(*googleapi.Error).Header. Use googleapi.IsNotModified to check
// whether the returned error was because http.StatusNotModified was
// returned.
func (c *ProjectsDatabasesCollectionGroupIndexesListCall) Do(opts ...googleapi.CallOption) (*GoogleFirestoreAdminV1beta2ListIndexesResponse, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &GoogleFirestoreAdminV1beta2ListIndexesResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Lists composite indexes.",
	//   "flatPath": "v1beta2/projects/{projectsId}/databases/{databasesId}/collectionGroup/{collectionGroupId}/indexes",
	//   "httpMethod": "GET",
	//   "id": "firestore.projects.databases.collectionGroup.indexes.list",
	//   "parameterOrder": [
	//     "parent"
	//   ],
	//   "parameters": {
	//     "filter": {
	//       "description": "The filter to apply to list results.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "pageSize": {
	//       "description": "The number of results to return.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "A page token, returned from a previous call to\nFirestoreAdmin.ListIndexes, that may be used to get the next\npage of results.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "parent": {
	//       "description": "A parent name of the form\n`projects/{project_id}/databases/{database_id}/collectionGroups/{collection_id}`",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/databases/[^/]+/collectionGroup/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta2/{+parent}/indexes",
	//   "response": {
	//     "$ref": "GoogleFirestoreAdminV1beta2ListIndexesResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/datastore"
	//   ]
	// }

}

// Pages invokes f for each page of results.
// A non-nil error returned from f will halt the iteration.
// The provided context supersedes any context provided to the Context method.
func (c *ProjectsDatabasesCollectionGroupIndexesListCall) Pages(ctx context.Context, f func(*GoogleFirestoreAdminV1beta2ListIndexesResponse) error) error {
	c.ctx_ = ctx
	defer c.PageToken(c.urlParams_.Get("pageToken")) // reset paging to original point
	for {
		x, err := c.Do()
		if err != nil {
			return err
		}
		if err := f(x); err != nil {
			return err
		}
		if x.NextPageToken == "" {
			return nil
		}
		c.PageToken(x.NextPageToken)
	}
}

// method id "firestore.projects.databases.collectionGroups.fields.get":

type ProjectsDatabasesCollectionGroupsFieldsGetCall struct {
	s            *Service
	name         string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// Get: Gets the metadata and configuration for a Field.
func (r *ProjectsDatabasesCollectionGroupsFieldsService) Get(name string) *ProjectsDatabasesCollectionGroupsFieldsGetCall {
	c := &ProjectsDatabasesCollectionGroupsFieldsGetCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsDatabasesCollectionGroupsFieldsGetCall) Fields(s ...googleapi.Field) *ProjectsDatabasesCollectionGroupsFieldsGetCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ProjectsDatabasesCollectionGroupsFieldsGetCall) IfNoneMatch(entityTag string) *ProjectsDatabasesCollectionGroupsFieldsGetCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsDatabasesCollectionGroupsFieldsGetCall) Context(ctx context.Context) *ProjectsDatabasesCollectionGroupsFieldsGetCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsDatabasesCollectionGroupsFieldsGetCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsDatabasesCollectionGroupsFieldsGetCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta2/{+name}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "firestore.projects.databases.collectionGroups.fields.get" call.
// Exactly one of *GoogleFirestoreAdminV1beta2Field or error will be
// non-nil. Any non-2xx status code is an error. Response headers are in
// either *GoogleFirestoreAdminV1beta2Field.ServerResponse.Header or (if
// a response was returned at all) in error.(*googleapi.Error).Header.
// Use googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ProjectsDatabasesCollectionGroupsFieldsGetCall) Do(opts ...googleapi.CallOption) (*GoogleFirestoreAdminV1beta2Field, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &GoogleFirestoreAdminV1beta2Field{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Gets the metadata and configuration for a Field.",
	//   "flatPath": "v1beta2/projects/{projectsId}/databases/{databasesId}/collectionGroups/{collectionGroupsId}/fields/{fieldsId}",
	//   "httpMethod": "GET",
	//   "id": "firestore.projects.databases.collectionGroups.fields.get",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "A name of the form\n`projects/{project_id}/databases/{database_id}/collectionGroups/{collection_id}/fields/{field_id}`",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/databases/[^/]+/collectionGroups/[^/]+/fields/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta2/{+name}",
	//   "response": {
	//     "$ref": "GoogleFirestoreAdminV1beta2Field"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/datastore"
	//   ]
	// }

}

// method id "firestore.projects.databases.collectionGroups.fields.list":

type ProjectsDatabasesCollectionGroupsFieldsListCall struct {
	s            *Service
	parent       string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// List: Lists the field configuration and metadata for this
// database.
//
// Currently, FirestoreAdmin.ListFields only supports listing
// fields
// that have been explicitly overridden. To issue this query,
// call
// FirestoreAdmin.ListFields with the filter set
// to
// `indexConfig.usesAncestorConfig:false`.
func (r *ProjectsDatabasesCollectionGroupsFieldsService) List(parent string) *ProjectsDatabasesCollectionGroupsFieldsListCall {
	c := &ProjectsDatabasesCollectionGroupsFieldsListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.parent = parent
	return c
}

// Filter sets the optional parameter "filter": The filter to apply to
// list results. Currently,
// FirestoreAdmin.ListFields only supports listing fields
// that have been explicitly overridden. To issue this query,
// call
// FirestoreAdmin.ListFields with the filter set
// to
// `indexConfig.usesAncestorConfig:false`.
func (c *ProjectsDatabasesCollectionGroupsFieldsListCall) Filter(filter string) *ProjectsDatabasesCollectionGroupsFieldsListCall {
	c.urlParams_.Set("filter", filter)
	return c
}

// PageSize sets the optional parameter "pageSize": The number of
// results to return.
func (c *ProjectsDatabasesCollectionGroupsFieldsListCall) PageSize(pageSize int64) *ProjectsDatabasesCollectionGroupsFieldsListCall {
	c.urlParams_.Set("pageSize", fmt.Sprint(pageSize))
	return c
}

// PageToken sets the optional parameter "pageToken": A page token,
// returned from a previous call to
// FirestoreAdmin.ListFields, that may be used to get the next
// page of results.
func (c *ProjectsDatabasesCollectionGroupsFieldsListCall) PageToken(pageToken string) *ProjectsDatabasesCollectionGroupsFieldsListCall {
	c.urlParams_.Set("pageToken", pageToken)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsDatabasesCollectionGroupsFieldsListCall) Fields(s ...googleapi.Field) *ProjectsDatabasesCollectionGroupsFieldsListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ProjectsDatabasesCollectionGroupsFieldsListCall) IfNoneMatch(entityTag string) *ProjectsDatabasesCollectionGroupsFieldsListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsDatabasesCollectionGroupsFieldsListCall) Context(ctx context.Context) *ProjectsDatabasesCollectionGroupsFieldsListCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsDatabasesCollectionGroupsFieldsListCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsDatabasesCollectionGroupsFieldsListCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta2/{+parent}/fields")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"parent": c.parent,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "firestore.projects.databases.collectionGroups.fields.list" call.
// Exactly one of *GoogleFirestoreAdminV1beta2ListFieldsResponse or
// error will be non-nil. Any non-2xx status code is an error. Response
// headers are in either
// *GoogleFirestoreAdminV1beta2ListFieldsResponse.ServerResponse.Header
// or (if a response was returned at all) in
// error.(*googleapi.Error).Header. Use googleapi.IsNotModified to check
// whether the returned error was because http.StatusNotModified was
// returned.
func (c *ProjectsDatabasesCollectionGroupsFieldsListCall) Do(opts ...googleapi.CallOption) (*GoogleFirestoreAdminV1beta2ListFieldsResponse, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &GoogleFirestoreAdminV1beta2ListFieldsResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Lists the field configuration and metadata for this database.\n\nCurrently, FirestoreAdmin.ListFields only supports listing fields\nthat have been explicitly overridden. To issue this query, call\nFirestoreAdmin.ListFields with the filter set to\n`indexConfig.usesAncestorConfig:false`.",
	//   "flatPath": "v1beta2/projects/{projectsId}/databases/{databasesId}/collectionGroups/{collectionGroupsId}/fields",
	//   "httpMethod": "GET",
	//   "id": "firestore.projects.databases.collectionGroups.fields.list",
	//   "parameterOrder": [
	//     "parent"
	//   ],
	//   "parameters": {
	//     "filter": {
	//       "description": "The filter to apply to list results. Currently,\nFirestoreAdmin.ListFields only supports listing fields\nthat have been explicitly overridden. To issue this query, call\nFirestoreAdmin.ListFields with the filter set to\n`indexConfig.usesAncestorConfig:false`.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "pageSize": {
	//       "description": "The number of results to return.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "A page token, returned from a previous call to\nFirestoreAdmin.ListFields, that may be used to get the next\npage of results.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "parent": {
	//       "description": "A parent name of the form\n`projects/{project_id}/databases/{database_id}/collectionGroups/{collection_id}`",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/databases/[^/]+/collectionGroups/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta2/{+parent}/fields",
	//   "response": {
	//     "$ref": "GoogleFirestoreAdminV1beta2ListFieldsResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/datastore"
	//   ]
	// }

}

// Pages invokes f for each page of results.
// A non-nil error returned from f will halt the iteration.
// The provided context supersedes any context provided to the Context method.
func (c *ProjectsDatabasesCollectionGroupsFieldsListCall) Pages(ctx context.Context, f func(*GoogleFirestoreAdminV1beta2ListFieldsResponse) error) error {
	c.ctx_ = ctx
	defer c.PageToken(c.urlParams_.Get("pageToken")) // reset paging to original point
	for {
		x, err := c.Do()
		if err != nil {
			return err
		}
		if err := f(x); err != nil {
			return err
		}
		if x.NextPageToken == "" {
			return nil
		}
		c.PageToken(x.NextPageToken)
	}
}

// method id "firestore.projects.databases.collectionGroups.fields.patch":

type ProjectsDatabasesCollectionGroupsFieldsPatchCall struct {
	s                                *Service
	name                             string
	googlefirestoreadminv1beta2field *GoogleFirestoreAdminV1beta2Field
	urlParams_                       gensupport.URLParams
	ctx_                             context.Context
	header_                          http.Header
}

// Patch: Updates a field configuration. Currently, field updates apply
// only to
// single field index configuration. However, calls
// to
// FirestoreAdmin.UpdateField should provide a field mask to
// avoid
// changing any configuration that the caller isn't aware of. The field
// mask
// should be specified as: `{ paths: "index_config" }`.
//
// This call returns a google.longrunning.Operation which may be used
// to
// track the status of the field update. The metadata for
// the operation will be the type FieldOperationMetadata.
//
// To configure the default field settings for the database, use
// the special `Field` with resource
// name:
// `projects/{project_id}/databases/{database_id}/collectionGroups/
// __default__/fields/*`.
func (r *ProjectsDatabasesCollectionGroupsFieldsService) Patch(name string, googlefirestoreadminv1beta2field *GoogleFirestoreAdminV1beta2Field) *ProjectsDatabasesCollectionGroupsFieldsPatchCall {
	c := &ProjectsDatabasesCollectionGroupsFieldsPatchCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	c.googlefirestoreadminv1beta2field = googlefirestoreadminv1beta2field
	return c
}

// UpdateMask sets the optional parameter "updateMask": A mask, relative
// to the field. If specified, only configuration specified
// by this field_mask will be updated in the field.
func (c *ProjectsDatabasesCollectionGroupsFieldsPatchCall) UpdateMask(updateMask string) *ProjectsDatabasesCollectionGroupsFieldsPatchCall {
	c.urlParams_.Set("updateMask", updateMask)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsDatabasesCollectionGroupsFieldsPatchCall) Fields(s ...googleapi.Field) *ProjectsDatabasesCollectionGroupsFieldsPatchCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsDatabasesCollectionGroupsFieldsPatchCall) Context(ctx context.Context) *ProjectsDatabasesCollectionGroupsFieldsPatchCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsDatabasesCollectionGroupsFieldsPatchCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsDatabasesCollectionGroupsFieldsPatchCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.googlefirestoreadminv1beta2field)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta2/{+name}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("PATCH", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "firestore.projects.databases.collectionGroups.fields.patch" call.
// Exactly one of *GoogleLongrunningOperation or error will be non-nil.
// Any non-2xx status code is an error. Response headers are in either
// *GoogleLongrunningOperation.ServerResponse.Header or (if a response
// was returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ProjectsDatabasesCollectionGroupsFieldsPatchCall) Do(opts ...googleapi.CallOption) (*GoogleLongrunningOperation, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &GoogleLongrunningOperation{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Updates a field configuration. Currently, field updates apply only to\nsingle field index configuration. However, calls to\nFirestoreAdmin.UpdateField should provide a field mask to avoid\nchanging any configuration that the caller isn't aware of. The field mask\nshould be specified as: `{ paths: \"index_config\" }`.\n\nThis call returns a google.longrunning.Operation which may be used to\ntrack the status of the field update. The metadata for\nthe operation will be the type FieldOperationMetadata.\n\nTo configure the default field settings for the database, use\nthe special `Field` with resource name:\n`projects/{project_id}/databases/{database_id}/collectionGroups/__default__/fields/*`.",
	//   "flatPath": "v1beta2/projects/{projectsId}/databases/{databasesId}/collectionGroups/{collectionGroupsId}/fields/{fieldsId}",
	//   "httpMethod": "PATCH",
	//   "id": "firestore.projects.databases.collectionGroups.fields.patch",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "A field name of the form\n`projects/{project_id}/databases/{database_id}/collectionGroups/{collection_id}/fields/{field_path}`\n\nA field path may be a simple field name, e.g. `address` or a path to fields\nwithin map_value , e.g. `address.city`,\nor a special field path. The only valid special field is `*`, which\nrepresents any field.\n\nField paths may be quoted using ` (backtick). The only character that needs\nto be escaped within a quoted field path is the backtick character itself,\nescaped using a backslash. Special characters in field paths that\nmust be quoted include: `*`, `.`,\n``` (backtick), `[`, `]`, as well as any ascii symbolic characters.\n\nExamples:\n(Note: Comments here are written in markdown syntax, so there is an\n additional layer of backticks to represent a code block)\n`\\`address.city\\`` represents a field named `address.city`, not the map key\n`city` in the field `address`.\n`\\`*\\`` represents a field named `*`, not any field.\n\nA special `Field` contains the default indexing settings for all fields.\nThis field's resource name is:\n`projects/{project_id}/databases/{database_id}/collectionGroups/__default__/fields/*`\nIndexes defined on this `Field` will be applied to all fields which do not\nhave their own `Field` index configuration.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/databases/[^/]+/collectionGroups/[^/]+/fields/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "updateMask": {
	//       "description": "A mask, relative to the field. If specified, only configuration specified\nby this field_mask will be updated in the field.",
	//       "format": "google-fieldmask",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta2/{+name}",
	//   "request": {
	//     "$ref": "GoogleFirestoreAdminV1beta2Field"
	//   },
	//   "response": {
	//     "$ref": "GoogleLongrunningOperation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/datastore"
	//   ]
	// }

}

// method id "firestore.projects.databases.collectionGroups.indexes.delete":

type ProjectsDatabasesCollectionGroupsIndexesDeleteCall struct {
	s          *Service
	name       string
	urlParams_ gensupport.URLParams
	ctx_       context.Context
	header_    http.Header
}

// Delete: Deletes a composite index.
func (r *ProjectsDatabasesCollectionGroupsIndexesService) Delete(name string) *ProjectsDatabasesCollectionGroupsIndexesDeleteCall {
	c := &ProjectsDatabasesCollectionGroupsIndexesDeleteCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsDatabasesCollectionGroupsIndexesDeleteCall) Fields(s ...googleapi.Field) *ProjectsDatabasesCollectionGroupsIndexesDeleteCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsDatabasesCollectionGroupsIndexesDeleteCall) Context(ctx context.Context) *ProjectsDatabasesCollectionGroupsIndexesDeleteCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsDatabasesCollectionGroupsIndexesDeleteCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsDatabasesCollectionGroupsIndexesDeleteCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta2/{+name}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("DELETE", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "firestore.projects.databases.collectionGroups.indexes.delete" call.
// Exactly one of *Empty or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Empty.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *ProjectsDatabasesCollectionGroupsIndexesDeleteCall) Do(opts ...googleapi.CallOption) (*Empty, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Empty{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Deletes a composite index.",
	//   "flatPath": "v1beta2/projects/{projectsId}/databases/{databasesId}/collectionGroups/{collectionGroupsId}/indexes/{indexesId}",
	//   "httpMethod": "DELETE",
	//   "id": "firestore.projects.databases.collectionGroups.indexes.delete",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "A name of the form\n`projects/{project_id}/databases/{database_id}/collectionGroups/{collection_id}/indexes/{index_id}`",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/databases/[^/]+/collectionGroups/[^/]+/indexes/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta2/{+name}",
	//   "response": {
	//     "$ref": "Empty"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/datastore"
	//   ]
	// }

}

// method id "firestore.projects.databases.collectionGroups.indexes.get":

type ProjectsDatabasesCollectionGroupsIndexesGetCall struct {
	s            *Service
	name         string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// Get: Gets a composite index.
func (r *ProjectsDatabasesCollectionGroupsIndexesService) Get(name string) *ProjectsDatabasesCollectionGroupsIndexesGetCall {
	c := &ProjectsDatabasesCollectionGroupsIndexesGetCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsDatabasesCollectionGroupsIndexesGetCall) Fields(s ...googleapi.Field) *ProjectsDatabasesCollectionGroupsIndexesGetCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ProjectsDatabasesCollectionGroupsIndexesGetCall) IfNoneMatch(entityTag string) *ProjectsDatabasesCollectionGroupsIndexesGetCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsDatabasesCollectionGroupsIndexesGetCall) Context(ctx context.Context) *ProjectsDatabasesCollectionGroupsIndexesGetCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsDatabasesCollectionGroupsIndexesGetCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsDatabasesCollectionGroupsIndexesGetCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta2/{+name}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "firestore.projects.databases.collectionGroups.indexes.get" call.
// Exactly one of *GoogleFirestoreAdminV1beta2Index or error will be
// non-nil. Any non-2xx status code is an error. Response headers are in
// either *GoogleFirestoreAdminV1beta2Index.ServerResponse.Header or (if
// a response was returned at all) in error.(*googleapi.Error).Header.
// Use googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ProjectsDatabasesCollectionGroupsIndexesGetCall) Do(opts ...googleapi.CallOption) (*GoogleFirestoreAdminV1beta2Index, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &GoogleFirestoreAdminV1beta2Index{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Gets a composite index.",
	//   "flatPath": "v1beta2/projects/{projectsId}/databases/{databasesId}/collectionGroups/{collectionGroupsId}/indexes/{indexesId}",
	//   "httpMethod": "GET",
	//   "id": "firestore.projects.databases.collectionGroups.indexes.get",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "A name of the form\n`projects/{project_id}/databases/{database_id}/collectionGroups/{collection_id}/indexes/{index_id}`",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/databases/[^/]+/collectionGroups/[^/]+/indexes/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta2/{+name}",
	//   "response": {
	//     "$ref": "GoogleFirestoreAdminV1beta2Index"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/datastore"
	//   ]
	// }

}
