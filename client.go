package goeve

import (
	"context"
	"errors"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/go-resty/resty/v2"
)

// DefaultBaseURL is the default base URL of the NetEase EVE Swagger Interface (ESI).
// The old host esi.evepc.163.com has been migrated to ali-esi.evepc.163.com.
// The API specification is published at https://ali-esi.evepc.163.com/latest/swagger.json
//
// DefaultBaseURL 是网易 EVE Swagger Interface (ESI) 的默认基础地址。
// 旧地址 esi.evepc.163.com 已迁移至 ali-esi.evepc.163.com。
// 接口规范发布于 https://ali-esi.evepc.163.com/latest/swagger.json
const DefaultBaseURL = "https://ali-esi.evepc.163.com/latest"

// Client is the EVE Online ESI API client.
// Client 是 EVE Online ESI API 客户端。
type Client struct {
	restyClient *resty.Client
	baseURL     string
	tokenSource *TokenSource
	httpClient  *http.Client
	language    string
}

// Option configures a Client. / Option 用于配置 Client。
type Option func(*Client)

// errBodyRequired is returned when a required request body is nil.
// errBodyRequired 在必需请求体为 nil 时返回。
var errBodyRequired = errors.New("goeve: request body is required")

// NewClient creates a new ESI API client. Apply options to customize it:
//
//	client := goeve.NewClient(
//	    goeve.WithDebug(true),
//	    goeve.WithAuthToken("access-token"),
//	)
//
// NewClient 创建一个新的 ESI API 客户端，可通过选项定制。
func NewClient(options ...Option) *Client {
	c := &Client{
		baseURL:     DefaultBaseURL,
		restyClient: resty.New(),
	}
	c.restyClient.SetBaseURL(c.baseURL)
	for _, option := range options {
		option(c)
	}
	if c.httpClient != nil {
		c.restyClient = resty.NewWithClient(c.httpClient)
		c.restyClient.SetBaseURL(c.baseURL)
	}
	return c
}

// WithBaseURL sets a custom base URL for the client.
// WithBaseURL 设置客户端的自定义基础地址。
func WithBaseURL(baseURL string) Option {
	return func(c *Client) {
		c.baseURL = baseURL
		c.restyClient.SetBaseURL(baseURL)
	}
}

// WithHTTPClient sets a custom underlying net/http client.
// WithHTTPClient 设置自定义的底层 net/http 客户端。
func WithHTTPClient(httpClient *http.Client) Option {
	return func(c *Client) {
		c.httpClient = httpClient
	}
}

// WithDebug enables or disables debug mode of the underlying resty client.
// WithDebug 开启或关闭底层 resty 客户端的调试模式。
func WithDebug(debug bool) Option {
	return func(c *Client) {
		c.restyClient.SetDebug(debug)
	}
}

// WithTimeout sets the timeout for each request.
// WithTimeout 设置每个请求的超时时间。
func WithTimeout(timeout time.Duration) Option {
	return func(c *Client) {
		c.restyClient.SetTimeout(timeout)
	}
}

// WithTokenSource attaches a TokenSource to the client. Before every request the
// current access token is obtained from the source, refreshing it automatically
// when it has expired. It is ignored for requests that carry their own token
// via the WithAuthToken RequestOption.
//
// WithTokenSource 为客户端挂载 TokenSource。每次请求前都会从该源获取当前
// 访问令牌，过期时自动使用 refresh_token 续期。对于通过 WithAuthToken
// 请求选项自带令牌的请求不生效。
//
// WithTokenSource 为客户端挂载 TokenSource。每次请求前都会从该源获取当前
// 访问令牌，过期时自动使用 refresh_token 续期。
func WithTokenSource(ts *TokenSource) Option {
	return func(c *Client) {
		c.tokenSource = ts
	}
}

// WithLanguage sets the default Accept-Language header sent with every request
// (e.g. goeve.LanguageChinese for Simplified Chinese). It can be overridden per
// request with the WithAcceptLanguage RequestOption.
//
// WithLanguage 设置每个请求默认携带的 Accept-Language 响应语言
// （如 goeve.LanguageChinese 简体中文）。可用 WithAcceptLanguage 请求选项按请求覆盖。
func WithLanguage(language string) Option {
	return func(c *Client) {
		c.language = language
	}
}

// RestyClient returns the underlying resty client for advanced customization.
// RestyClient 返回底层 resty 客户端，用于高级定制。
func (c *Client) RestyClient() *resty.Client {
	return c.restyClient
}

// send executes an HTTP request against the ESI API and handles errors uniformly.
// result must be a non-nil pointer for requests expected to return a body.
//
// send 统一执行对 ESI API 的 HTTP 请求并处理错误。
// 对于预期返回响应体的请求，result 必须是非 nil 指针。
func (c *Client) send(ctx context.Context, method, path string, pathParams map[string]string, query url.Values, headers map[string]string, body, result any) (*resty.Response, error) {
	request := c.restyClient.NewRequest()
	request.SetContext(ctx)
	if _, ok := headers["Authorization"]; !ok && c.tokenSource != nil {
		token, err := c.tokenSource.Token(ctx)
		if err != nil {
			return nil, err
		}
		if token != "" {
			request.SetAuthToken(token)
		}
	}
	if len(pathParams) > 0 {
		request.SetPathParams(pathParams)
	}
	if len(query) > 0 {
		request.SetQueryParamsFromValues(query)
	}
	for key, value := range headers {
		request.SetHeader(key, value)
	}
	if _, ok := headers["Accept-Language"]; !ok && c.language != "" {
		request.SetHeader("Accept-Language", c.language)
	}
	if body != nil {
		request.SetBody(body)
	}
	if result != nil {
		request.SetResult(result)
	}
	request.SetError(&APIError{})

	response, err := request.Execute(method, path)
	if err != nil {
		return response, err
	}
	// 304 Not Modified: not an error, the caller compares the ETag itself.
	// 304 Not Modified：不是错误，由调用方自行比较 ETag。
	if response.StatusCode() == http.StatusNotModified {
		return response, nil
	}
	if response.IsError() {
		if apiErr, ok := response.Error().(*APIError); ok {
			apiErr.StatusCode = response.StatusCode()
			if apiErr.Message == "" {
				apiErr.RawBody = response.String()
			}
			return response, apiErr
		}
		return response, &APIError{
			StatusCode: response.StatusCode(),
			RawBody:    response.String(),
		}
	}
	return response, nil
}

// get / post / put / delete are thin wrappers around send used by the module methods.
// get / post / put / delete 是供各模块方法使用的 send 快捷封装。

func (c *Client) get(ctx context.Context, path string, pathParams map[string]string, query url.Values, headers map[string]string, result any) error {
	_, err := c.send(ctx, http.MethodGet, path, pathParams, query, headers, nil, result)
	return err
}

func (c *Client) post(ctx context.Context, path string, pathParams map[string]string, query url.Values, headers map[string]string, body, result any) error {
	_, err := c.send(ctx, http.MethodPost, path, pathParams, query, headers, body, result)
	return err
}

func (c *Client) put(ctx context.Context, path string, pathParams map[string]string, query url.Values, headers map[string]string, body any) error {
	_, err := c.send(ctx, http.MethodPut, path, pathParams, query, headers, body, nil)
	return err
}

func (c *Client) delete(ctx context.Context, path string, pathParams map[string]string, query url.Values, headers map[string]string) error {
	_, err := c.send(ctx, http.MethodDelete, path, pathParams, query, headers, nil, nil)
	return err
}

// requestOptions accumulates per-request query parameters and headers.
// requestOptions 汇总每次请求可覆盖的查询参数与请求头。
type requestOptions struct {
	query   url.Values
	headers map[string]string
}

// RequestOption customizes a single request: pagination, ETag, language and
// endpoint-specific query parameters. Pass them variadically to any API method.
//
// RequestOption 用于定制单次请求：分页、ETag、语言以及各接口特有的查询参数，
// 以可变参数形式传给任意 API 方法。
type RequestOption func(*requestOptions)

func (o *requestOptions) setQuery(key, value string) {
	if o.query == nil {
		o.query = url.Values{}
	}
	o.query.Set(key, value)
}

func (o *requestOptions) addQuery(key, value string) {
	if o.query == nil {
		o.query = url.Values{}
	}
	o.query.Add(key, value)
}

func (o *requestOptions) setHeader(key, value string) {
	if o.headers == nil {
		o.headers = map[string]string{}
	}
	o.headers[key] = value
}

// newRequestOptions folds the given options into query values and headers.
// newRequestOptions 将选项折叠为查询参数与请求头。
func newRequestOptions(options ...RequestOption) (url.Values, map[string]string) {
	o := &requestOptions{query: url.Values{}, headers: map[string]string{}}
	for _, option := range options {
		option(o)
	}
	return o.query, o.headers
}

// WithPage selects which page of results to return (ESI default page size is 1000).
// WithPage 选择返回结果的页码（ESI 默认每页 1000 条）。
func WithPage(page int32) RequestOption {
	return func(o *requestOptions) { o.setQuery("page", strconv.FormatInt(int64(page), 10)) }
}

// WithIfNoneMatch sends an ETag from a previous request; the server returns 304
// and the method yields a nil result with a nil error when nothing changed.
//
// WithIfNoneMatch 携带上次请求的 ETag；数据未变化时服务器返回 304，
// 方法将返回 nil 结果与 nil error。
func WithIfNoneMatch(etag string) RequestOption {
	return func(o *requestOptions) { o.setHeader("If-None-Match", etag) }
}

// WithAcceptLanguage overrides the client's default Accept-Language for this request.
// WithAcceptLanguage 覆盖客户端默认的 Accept-Language（本次请求生效）。
func WithAcceptLanguage(language string) RequestOption {
	return func(o *requestOptions) { o.setHeader("Accept-Language", language) }
}

// WithAuthToken sends the given access token as Authorization: Bearer for this
// request only — use it on endpoints that require authentication. It overrides
// the client's TokenSource for this request; pass an empty token to make an
// unauthenticated request.
//
// WithAuthToken 仅本次请求携带指定的 access_token（Authorization: Bearer），
// 用于需要授权的接口。它会覆盖客户端的 TokenSource（传空字符串可强制匿名请求）。
func WithAuthToken(token string) RequestOption {
	return func(o *requestOptions) {
		if token != "" {
			o.setHeader("Authorization", "Bearer "+strings.TrimSpace(token))
		}
	}
}

// WithFromEvent filters calendar events starting after the given event ID.
// WithFromEvent 只返回晚于指定事件 ID 的日历事件。
func WithFromEvent(fromEvent int32) RequestOption {
	return func(o *requestOptions) { o.setQuery("from_event", strconv.FormatInt(int64(fromEvent), 10)) }
}

// WithLabelIDs filters contacts by the given label IDs (repeated query parameter).
// WithLabelIDs 按标签 ID 过滤联系人（重复查询参数）。
func WithLabelIDs(labelIDs []int32) RequestOption {
	return func(o *requestOptions) {
		for _, id := range labelIDs {
			o.addQuery("label_ids", strconv.FormatInt(int64(id), 10))
		}
	}
}

// WithWatched filters contacts or corporation structures by watched status.
// WithWatched 按关注状态过滤联系人或军团建筑。
func WithWatched(watched bool) RequestOption {
	return func(o *requestOptions) { o.setQuery("watched", strconv.FormatBool(watched)) }
}

// WithIncludeCompleted includes finished industry jobs.
// WithIncludeCompleted 包含已完成的工业作业。
func WithIncludeCompleted(include bool) RequestOption {
	return func(o *requestOptions) { o.setQuery("include_completed", strconv.FormatBool(include)) }
}

// WithLabels filters mails by label IDs (repeated query parameter).
// WithLabels 按标签 ID 过滤邮件（重复查询参数）。
func WithLabels(labels []int32) RequestOption {
	return func(o *requestOptions) {
		for _, id := range labels {
			o.addQuery("labels", strconv.FormatInt(int64(id), 10))
		}
	}
}

// WithLastMailID returns only mails older than the given mail ID.
// WithLastMailID 只返回早于指定邮件 ID 的邮件。
func WithLastMailID(lastMailID int32) RequestOption {
	return func(o *requestOptions) { o.setQuery("last_mail_id", strconv.FormatInt(int64(lastMailID), 10)) }
}

// WithStrict enables strict matching for name resolution.
// WithStrict 启用名称解析的严格匹配模式。
func WithStrict(strict bool) RequestOption {
	return func(o *requestOptions) { o.setQuery("strict", strconv.FormatBool(strict)) }
}

// WithFromID returns only entries older than the given ID (wallet/orders walk-back).
// WithFromID 只返回早于指定 ID 的记录（钱包/订单回溯翻页）。
func WithFromID(fromID int64) RequestOption {
	return func(o *requestOptions) { o.setQuery("from_id", strconv.FormatInt(fromID, 10)) }
}

// WithAvoid asks the route planner to avoid the given solar systems.
// WithAvoid 要求路线规划避开指定星系。
func WithAvoid(avoid []int32) RequestOption {
	return func(o *requestOptions) {
		for _, id := range avoid {
			o.addQuery("avoid", strconv.FormatInt(int64(id), 10))
		}
	}
}

// WithConnections adds custom system-to-system connections for route planning;
// each inner pair is joined with a pipe ("from|to").
//
// WithConnections 为路线规划添加自定义星系连接；每个内层二元组以管道符拼接（"from|to"）。
func WithConnections(connections [][]int32) RequestOption {
	return func(o *requestOptions) {
		for _, pair := range connections {
			if len(pair) == 2 {
				o.addQuery("connections", strconv.FormatInt(int64(pair[0]), 10)+"|"+strconv.FormatInt(int64(pair[1]), 10))
			}
		}
	}
}

// WithFlag selects the route security preference: "shortest", "secure" or "insecure".
// WithFlag 选择路线安全偏好："shortest"、"secure" 或 "insecure"。
func WithFlag(flag string) RequestOption {
	return func(o *requestOptions) { o.setQuery("flag", flag) }
}

// WithFilter filters structure market orders: "market" or "manufacturing_basic".
// WithFilter 过滤建筑市场订单："market" 或 "manufacturing_basic"。
func WithFilter(filter string) RequestOption {
	return func(o *requestOptions) { o.setQuery("filter", filter) }
}

// WithMaxWarID returns only wars with an ID below the given one.
// WithMaxWarID 只返回 ID 小于指定值的战争记录。
func WithMaxWarID(maxWarID int32) RequestOption {
	return func(o *requestOptions) { o.setQuery("max_war_id", strconv.FormatInt(int64(maxWarID), 10)) }
}
