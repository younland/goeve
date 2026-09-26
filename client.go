package goeve

import (
	"context"
	"errors"
	"net/http"
	"net/url"
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

// WithAuthToken sets the access token used for Authorization: Bearer authentication.
// The token will be sent with every request.
// WithAuthToken 设置用于 Authorization: Bearer 认证的访问令牌，随每个请求发送。
func WithAuthToken(token string) Option {
	return func(c *Client) {
		c.restyClient.SetAuthToken(token)
	}
}

// WithTokenSource attaches a TokenSource to the client. Before every request the
// current access token is obtained from the source, refreshing it automatically
// when it has expired.
//
// WithTokenSource 为客户端挂载 TokenSource。每次请求前都会从该源获取当前
// 访问令牌，过期时自动使用 refresh_token 续期。
func WithTokenSource(ts *TokenSource) Option {
	return func(c *Client) {
		c.tokenSource = ts
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
	if c.tokenSource != nil {
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
