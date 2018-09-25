// Package cloudtrace provides access to the Stackdriver Trace API.
//
// See https://cloud.google.com/trace
//
// Usage example:
//
//   import "github.com/gxed/google-api-go-client/cloudtrace/v1"
//   ...
//   cloudtraceService, err := cloudtrace.New(oauthHttpClient)
package cloudtrace

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

const apiId = "cloudtrace:v1"
const apiName = "cloudtrace"
const apiVersion = "v1"
const basePath = "https://cloudtrace.googleapis.com/"

// OAuth2 scopes used by this API.
const (
	// View and manage your data across Google Cloud Platform services
	CloudPlatformScope = "https://www.googleapis.com/auth/cloud-platform"

	// Write Trace data for a project or application
	TraceAppendScope = "https://www.googleapis.com/auth/trace.append"

	// Read Trace data for a project or application
	TraceReadonlyScope = "https://www.googleapis.com/auth/trace.readonly"
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
	rs.Traces = NewProjectsTracesService(s)
	return rs
}

type ProjectsService struct {
	s *Service

	Traces *ProjectsTracesService
}

func NewProjectsTracesService(s *Service) *ProjectsTracesService {
	rs := &ProjectsTracesService{s: s}
	return rs
}

type ProjectsTracesService struct {
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

// ListTracesResponse: The response message for the `ListTraces` method.
type ListTracesResponse struct {
	// NextPageToken: If defined, indicates that there are more traces that
	// match the request
	// and that this value should be passed to the next request to
	// continue
	// retrieving additional traces.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// Traces: List of trace records as specified by the view parameter.
	Traces []*Trace `json:"traces,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "NextPageToken") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "NextPageToken") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ListTracesResponse) MarshalJSON() ([]byte, error) {
	type NoMethod ListTracesResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Trace: A trace describes how long it takes for an application to
// perform an
// operation. It consists of a set of spans, each of which represent a
// single
// timed event within the operation.
type Trace struct {
	// ProjectId: Project ID of the Cloud project where the trace data is
	// stored.
	ProjectId string `json:"projectId,omitempty"`

	// Spans: Collection of spans in the trace.
	Spans []*TraceSpan `json:"spans,omitempty"`

	// TraceId: Globally unique identifier for the trace. This identifier is
	// a 128-bit
	// numeric value formatted as a 32-byte hex string. For
	// example,
	// `382d4f4c6b7bb2f4a972559d9085001d`.
	TraceId string `json:"traceId,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "ProjectId") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "ProjectId") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Trace) MarshalJSON() ([]byte, error) {
	type NoMethod Trace
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// TraceSpan: A span represents a single timed event within a trace.
// Spans can be nested
// and form a trace tree. Often, a trace contains a root span that
// describes the
// end-to-end latency of an operation and, optionally, one or more
// subspans for
// its suboperations. Spans do not need to be contiguous. There may be
// gaps
// between spans in a trace.
type TraceSpan struct {
	// EndTime: End time of the span in nanoseconds from the UNIX epoch.
	EndTime string `json:"endTime,omitempty"`

	// Kind: Distinguishes between spans generated in a particular context.
	// For example,
	// two spans with the same name may be distinguished using
	// `RPC_CLIENT`
	// and `RPC_SERVER` to identify queueing latency associated with the
	// span.
	//
	// Possible values:
	//   "SPAN_KIND_UNSPECIFIED" - Unspecified.
	//   "RPC_SERVER" - Indicates that the span covers server-side handling
	// of an RPC or other
	// remote network request.
	//   "RPC_CLIENT" - Indicates that the span covers the client-side
	// wrapper around an RPC or
	// other remote request.
	Kind string `json:"kind,omitempty"`

	// Labels: Collection of labels associated with the span. Label keys
	// must be less than
	// 128 bytes. Label values must be less than 16 kilobytes (10MB
	// for
	// `/stacktrace` values).
	//
	// Some predefined label keys exist, or you may create your own. When
	// creating
	// your own, we recommend the following formats:
	//
	// * `/category/product/key` for agents of well-known products (e.g.
	//   `/db/mongodb/read_size`).
	// * `short_host/path/key` for domain-specific keys (e.g.
	//   `foo.com/myproduct/bar`)
	//
	// Predefined labels include:
	//
	// *   `/agent`
	// *   `/component`
	// *   `/error/message`
	// *   `/error/name`
	// *   `/http/client_city`
	// *   `/http/client_country`
	// *   `/http/client_protocol`
	// *   `/http/client_region`
	// *   `/http/host`
	// *   `/http/method`
	// *   `/http/path`
	// *   `/http/redirected_url`
	// *   `/http/request/size`
	// *   `/http/response/size`
	// *   `/http/route`
	// *   `/http/status_code`
	// *   `/http/url`
	// *   `/http/user_agent`
	// *   `/pid`
	// *   `/stacktrace`
	// *   `/tid`
	Labels map[string]string `json:"labels,omitempty"`

	// Name: Name of the span. Must be less than 128 bytes. The span name is
	// sanitized
	// and displayed in the Stackdriver Trace tool in the
	// Google Cloud Platform Console.
	// The name may be a method name or some other per-call site name.
	// For the same executable and the same call point, a best practice
	// is
	// to use a consistent name, which makes it easier to
	// correlate
	// cross-trace spans.
	Name string `json:"name,omitempty"`

	// ParentSpanId: ID of the parent span, if any. Optional.
	ParentSpanId uint64 `json:"parentSpanId,omitempty,string"`

	// SpanId: Identifier for the span. Must be a 64-bit integer other than
	// 0 and
	// unique within a trace. For example, `2205310701640571284`.
	SpanId uint64 `json:"spanId,omitempty,string"`

	// StartTime: Start time of the span in nanoseconds from the UNIX epoch.
	StartTime string `json:"startTime,omitempty"`

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

func (s *TraceSpan) MarshalJSON() ([]byte, error) {
	type NoMethod TraceSpan
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Traces: List of new or updated traces.
type Traces struct {
	// Traces: List of traces.
	Traces []*Trace `json:"traces,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Traces") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Traces") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Traces) MarshalJSON() ([]byte, error) {
	type NoMethod Traces
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// method id "cloudtrace.projects.patchTraces":

type ProjectsPatchTracesCall struct {
	s          *Service
	projectId  string
	traces     *Traces
	urlParams_ gensupport.URLParams
	ctx_       context.Context
	header_    http.Header
}

// PatchTraces: Sends new traces to Stackdriver Trace or updates
// existing traces. If the ID
// of a trace that you send matches that of an existing trace, any
// fields
// in the existing trace and its spans are overwritten by the provided
// values,
// and any new fields provided are merged with the existing trace data.
// If the
// ID does not match, a new trace is created.
func (r *ProjectsService) PatchTraces(projectId string, traces *Traces) *ProjectsPatchTracesCall {
	c := &ProjectsPatchTracesCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.projectId = projectId
	c.traces = traces
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsPatchTracesCall) Fields(s ...googleapi.Field) *ProjectsPatchTracesCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsPatchTracesCall) Context(ctx context.Context) *ProjectsPatchTracesCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsPatchTracesCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsPatchTracesCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.traces)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/projects/{projectId}/traces")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("PATCH", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"projectId": c.projectId,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "cloudtrace.projects.patchTraces" call.
// Exactly one of *Empty or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Empty.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *ProjectsPatchTracesCall) Do(opts ...googleapi.CallOption) (*Empty, error) {
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
	//   "description": "Sends new traces to Stackdriver Trace or updates existing traces. If the ID\nof a trace that you send matches that of an existing trace, any fields\nin the existing trace and its spans are overwritten by the provided values,\nand any new fields provided are merged with the existing trace data. If the\nID does not match, a new trace is created.",
	//   "flatPath": "v1/projects/{projectId}/traces",
	//   "httpMethod": "PATCH",
	//   "id": "cloudtrace.projects.patchTraces",
	//   "parameterOrder": [
	//     "projectId"
	//   ],
	//   "parameters": {
	//     "projectId": {
	//       "description": "ID of the Cloud project where the trace data is stored.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/projects/{projectId}/traces",
	//   "request": {
	//     "$ref": "Traces"
	//   },
	//   "response": {
	//     "$ref": "Empty"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/trace.append"
	//   ]
	// }

}

// method id "cloudtrace.projects.traces.get":

type ProjectsTracesGetCall struct {
	s            *Service
	projectId    string
	traceId      string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// Get: Gets a single trace by its ID.
func (r *ProjectsTracesService) Get(projectId string, traceId string) *ProjectsTracesGetCall {
	c := &ProjectsTracesGetCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.projectId = projectId
	c.traceId = traceId
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsTracesGetCall) Fields(s ...googleapi.Field) *ProjectsTracesGetCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ProjectsTracesGetCall) IfNoneMatch(entityTag string) *ProjectsTracesGetCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsTracesGetCall) Context(ctx context.Context) *ProjectsTracesGetCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsTracesGetCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsTracesGetCall) doRequest(alt string) (*http.Response, error) {
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
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/projects/{projectId}/traces/{traceId}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"projectId": c.projectId,
		"traceId":   c.traceId,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "cloudtrace.projects.traces.get" call.
// Exactly one of *Trace or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Trace.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *ProjectsTracesGetCall) Do(opts ...googleapi.CallOption) (*Trace, error) {
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
	ret := &Trace{
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
	//   "description": "Gets a single trace by its ID.",
	//   "flatPath": "v1/projects/{projectId}/traces/{traceId}",
	//   "httpMethod": "GET",
	//   "id": "cloudtrace.projects.traces.get",
	//   "parameterOrder": [
	//     "projectId",
	//     "traceId"
	//   ],
	//   "parameters": {
	//     "projectId": {
	//       "description": "ID of the Cloud project where the trace data is stored.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "traceId": {
	//       "description": "ID of the trace to return.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/projects/{projectId}/traces/{traceId}",
	//   "response": {
	//     "$ref": "Trace"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/trace.readonly"
	//   ]
	// }

}

// method id "cloudtrace.projects.traces.list":

type ProjectsTracesListCall struct {
	s            *Service
	projectId    string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// List: Returns of a list of traces that match the specified filter
// conditions.
func (r *ProjectsTracesService) List(projectId string) *ProjectsTracesListCall {
	c := &ProjectsTracesListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.projectId = projectId
	return c
}

// EndTime sets the optional parameter "endTime": End of the time
// interval (inclusive) during which the trace data was
// collected from the application.
func (c *ProjectsTracesListCall) EndTime(endTime string) *ProjectsTracesListCall {
	c.urlParams_.Set("endTime", endTime)
	return c
}

// Filter sets the optional parameter "filter": An optional filter
// against labels for the request.
//
// By default, searches use prefix matching. To specify exact match,
// prepend
// a plus symbol (`+`) to the search term.
// Multiple terms are ANDed. Syntax:
//
// *   `root:NAME_PREFIX` or `NAME_PREFIX`: Return traces where any
// root
//     span starts with `NAME_PREFIX`.
// *   `+root:NAME` or `+NAME`: Return traces where any root span's name
// is
//     exactly `NAME`.
// *   `span:NAME_PREFIX`: Return traces where any span starts with
//     `NAME_PREFIX`.
// *   `+span:NAME`: Return traces where any span's name is exactly
//     `NAME`.
// *   `latency:DURATION`: Return traces whose overall latency is
//     greater or equal to than `DURATION`. Accepted units are
// nanoseconds
//     (`ns`), milliseconds (`ms`), and seconds (`s`). Default is `ms`.
// For
//     example, `latency:24ms` returns traces whose overall latency
//     is greater than or equal to 24 milliseconds.
// *   `label:LABEL_KEY`: Return all traces containing the specified
//     label key (exact match, case-sensitive) regardless of the
// key:value
//     pair's value (including empty values).
// *   `LABEL_KEY:VALUE_PREFIX`: Return all traces containing the
// specified
//     label key (exact match, case-sensitive) whose value starts with
//     `VALUE_PREFIX`. Both a key and a value must be specified.
// *   `+LABEL_KEY:VALUE`: Return all traces containing a key:value
// pair
//     exactly matching the specified text. Both a key and a value must
// be
//     specified.
// *   `method:VALUE`: Equivalent to `/http/method:VALUE`.
// *   `url:VALUE`: Equivalent to `/http/url:VALUE`.
func (c *ProjectsTracesListCall) Filter(filter string) *ProjectsTracesListCall {
	c.urlParams_.Set("filter", filter)
	return c
}

// OrderBy sets the optional parameter "orderBy": Field used to sort the
// returned traces.
// Can be one of the following:
//
// *   `trace_id`
// *   `name` (`name` field of root span in the trace)
// *   `duration` (difference between `end_time` and `start_time` fields
// of
//      the root span)
// *   `start` (`start_time` field of the root span)
//
// Descending order can be specified by appending `desc` to the sort
// field
// (for example, `name desc`).
//
// Only one sort field is permitted.
func (c *ProjectsTracesListCall) OrderBy(orderBy string) *ProjectsTracesListCall {
	c.urlParams_.Set("orderBy", orderBy)
	return c
}

// PageSize sets the optional parameter "pageSize": Maximum number of
// traces to return. If not specified or <= 0, the
// implementation selects a reasonable value.  The implementation
// may
// return fewer traces than the requested page size.
func (c *ProjectsTracesListCall) PageSize(pageSize int64) *ProjectsTracesListCall {
	c.urlParams_.Set("pageSize", fmt.Sprint(pageSize))
	return c
}

// PageToken sets the optional parameter "pageToken": Token identifying
// the page of results to return. If provided, use the
// value of the `next_page_token` field from a previous request.
func (c *ProjectsTracesListCall) PageToken(pageToken string) *ProjectsTracesListCall {
	c.urlParams_.Set("pageToken", pageToken)
	return c
}

// StartTime sets the optional parameter "startTime": Start of the time
// interval (inclusive) during which the trace data was
// collected from the application.
func (c *ProjectsTracesListCall) StartTime(startTime string) *ProjectsTracesListCall {
	c.urlParams_.Set("startTime", startTime)
	return c
}

// View sets the optional parameter "view": Type of data returned for
// traces in the list.  Default is
// `MINIMAL`.
//
// Possible values:
//   "VIEW_TYPE_UNSPECIFIED"
//   "MINIMAL"
//   "ROOTSPAN"
//   "COMPLETE"
func (c *ProjectsTracesListCall) View(view string) *ProjectsTracesListCall {
	c.urlParams_.Set("view", view)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsTracesListCall) Fields(s ...googleapi.Field) *ProjectsTracesListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ProjectsTracesListCall) IfNoneMatch(entityTag string) *ProjectsTracesListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsTracesListCall) Context(ctx context.Context) *ProjectsTracesListCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsTracesListCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsTracesListCall) doRequest(alt string) (*http.Response, error) {
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
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/projects/{projectId}/traces")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"projectId": c.projectId,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "cloudtrace.projects.traces.list" call.
// Exactly one of *ListTracesResponse or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *ListTracesResponse.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ProjectsTracesListCall) Do(opts ...googleapi.CallOption) (*ListTracesResponse, error) {
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
	ret := &ListTracesResponse{
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
	//   "description": "Returns of a list of traces that match the specified filter conditions.",
	//   "flatPath": "v1/projects/{projectId}/traces",
	//   "httpMethod": "GET",
	//   "id": "cloudtrace.projects.traces.list",
	//   "parameterOrder": [
	//     "projectId"
	//   ],
	//   "parameters": {
	//     "endTime": {
	//       "description": "End of the time interval (inclusive) during which the trace data was\ncollected from the application.",
	//       "format": "google-datetime",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "filter": {
	//       "description": "An optional filter against labels for the request.\n\nBy default, searches use prefix matching. To specify exact match, prepend\na plus symbol (`+`) to the search term.\nMultiple terms are ANDed. Syntax:\n\n*   `root:NAME_PREFIX` or `NAME_PREFIX`: Return traces where any root\n    span starts with `NAME_PREFIX`.\n*   `+root:NAME` or `+NAME`: Return traces where any root span's name is\n    exactly `NAME`.\n*   `span:NAME_PREFIX`: Return traces where any span starts with\n    `NAME_PREFIX`.\n*   `+span:NAME`: Return traces where any span's name is exactly\n    `NAME`.\n*   `latency:DURATION`: Return traces whose overall latency is\n    greater or equal to than `DURATION`. Accepted units are nanoseconds\n    (`ns`), milliseconds (`ms`), and seconds (`s`). Default is `ms`. For\n    example, `latency:24ms` returns traces whose overall latency\n    is greater than or equal to 24 milliseconds.\n*   `label:LABEL_KEY`: Return all traces containing the specified\n    label key (exact match, case-sensitive) regardless of the key:value\n    pair's value (including empty values).\n*   `LABEL_KEY:VALUE_PREFIX`: Return all traces containing the specified\n    label key (exact match, case-sensitive) whose value starts with\n    `VALUE_PREFIX`. Both a key and a value must be specified.\n*   `+LABEL_KEY:VALUE`: Return all traces containing a key:value pair\n    exactly matching the specified text. Both a key and a value must be\n    specified.\n*   `method:VALUE`: Equivalent to `/http/method:VALUE`.\n*   `url:VALUE`: Equivalent to `/http/url:VALUE`.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "orderBy": {
	//       "description": "Field used to sort the returned traces. Optional.\nCan be one of the following:\n\n*   `trace_id`\n*   `name` (`name` field of root span in the trace)\n*   `duration` (difference between `end_time` and `start_time` fields of\n     the root span)\n*   `start` (`start_time` field of the root span)\n\nDescending order can be specified by appending `desc` to the sort field\n(for example, `name desc`).\n\nOnly one sort field is permitted.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "pageSize": {
	//       "description": "Maximum number of traces to return. If not specified or \u003c= 0, the\nimplementation selects a reasonable value.  The implementation may\nreturn fewer traces than the requested page size. Optional.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "Token identifying the page of results to return. If provided, use the\nvalue of the `next_page_token` field from a previous request. Optional.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "projectId": {
	//       "description": "ID of the Cloud project where the trace data is stored.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "startTime": {
	//       "description": "Start of the time interval (inclusive) during which the trace data was\ncollected from the application.",
	//       "format": "google-datetime",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "view": {
	//       "description": "Type of data returned for traces in the list. Optional. Default is\n`MINIMAL`.",
	//       "enum": [
	//         "VIEW_TYPE_UNSPECIFIED",
	//         "MINIMAL",
	//         "ROOTSPAN",
	//         "COMPLETE"
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/projects/{projectId}/traces",
	//   "response": {
	//     "$ref": "ListTracesResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/trace.readonly"
	//   ]
	// }

}

// Pages invokes f for each page of results.
// A non-nil error returned from f will halt the iteration.
// The provided context supersedes any context provided to the Context method.
func (c *ProjectsTracesListCall) Pages(ctx context.Context, f func(*ListTracesResponse) error) error {
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
