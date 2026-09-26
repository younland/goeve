package goeve

import (
	"context"
	"errors"
	"fmt"
	"net/url"
	"strings"
	"sync"
	"time"

	"github.com/go-resty/resty/v2"
)

// EVE SSO (OAuth2) endpoints on the NetEase login server.
// 网易登录服务器上的 EVE SSO (OAuth2) 端点。
const (
	// SSOBaseURL is the base URL of the NetEase EVE SSO server.
	// SSOBaseURL 是网易 EVE SSO 服务器的基地址。
	SSOBaseURL = "https://login.evepc.163.com"

	// SSOAuthorizePath is the authorization endpoint.
	// SSOAuthorizePath 是授权端点路径。
	SSOAuthorizePath = "/v2/oauth/authorize"

	// SSOTokenPath is the token endpoint.
	// SSOTokenPath 是令牌端点路径。
	SSOTokenPath = "/v2/oauth/token"
)

// DefaultClientID is the client_id of the official ESI web application.
// NetEase does not provide a public place to register your own OAuth application;
// third-party tools commonly reuse this official client_id.
//
// DefaultClientID 是官方 ESI 网页应用使用的 client_id。
// 网易未开放自助申请 OAuth 应用的入口，第三方工具通常直接复用该官方 client_id。
const DefaultClientID = "bc90aa496a404724a93f41b4f4e97761"

// DefaultRedirectURI is the only officially usable redirect_uri.
// DefaultRedirectURI 是目前唯一可用的回调地址。
const DefaultRedirectURI = "https://esi.evepc.163.com/ui/oauth2-redirect.html"

// Response types supported by the EVE SSO authorize endpoint.
// EVE SSO 授权端点支持的授权模式。
const (
	// ResponseTypeToken is the implicit grant: the access token is returned
	// directly in the redirect URL fragment (expires in ~20 minutes, no refresh token).
	// ResponseTypeToken 为隐式模式：access_token 直接拼接在跳转地址的 fragment 中
	//（约 20 分钟过期，无 refresh_token）。
	ResponseTypeToken = "token"

	// ResponseTypeCode is the authorization code grant: exchange the code for an
	// access token plus a permanent refresh token.
	// ResponseTypeCode 为授权码模式：用授权码换取 access_token 及永久有效的 refresh_token。
	ResponseTypeCode = "code"
)

// Token is an OAuth2 token returned by the NetEase EVE SSO token endpoint.
// Token 是网易 EVE SSO 令牌端点返回的 OAuth2 令牌。
type Token struct {
	// AccessToken is the bearer token used to access protected ESI endpoints.
	// AccessToken 是访问 ESI 受保护端点所用的 Bearer 令牌。
	AccessToken string `json:"access_token"`

	// RefreshToken is the long-lived token used to obtain new access tokens.
	// Treat it like a password: anyone holding it can access the account.
	// RefreshToken 是长期有效的刷新令牌，可用于获取新的 access_token。
	// 请像对待密码一样保管它：任何持有它的人都能访问该账号。
	RefreshToken string `json:"refresh_token"`

	// TokenType is the token type, always "Bearer".
	// TokenType 为令牌类型，恒为 "Bearer"。
	TokenType string `json:"token_type"`

	// ExpiresIn is the lifetime of the access token in seconds (usually 1200 = 20 minutes).
	// ExpiresIn 为 access_token 的有效期（秒），通常为 1200（20 分钟）。
	ExpiresIn int `json:"expires_in"`

	// Expiry is the point in time when the access token expires.
	// Expiry 为 access_token 的过期时间点。
	Expiry time.Time `json:"-"`
}

// Valid reports whether the access token is present and not expired,
// leaving a 30-second safety margin.
//
// Valid 报告 access_token 是否存在且未过期（预留 30 秒安全余量）。
func (t *Token) Valid() bool {
	return t != nil && t.AccessToken != "" && time.Now().Before(t.Expiry.Add(-30*time.Second))
}

// AuthorizeConfig configures the EVE SSO authorization URL.
// AuthorizeConfig 用于配置 EVE SSO 授权链接。
type AuthorizeConfig struct {
	// ResponseType selects the grant type: ResponseTypeToken (implicit) or ResponseTypeCode.
	// ResponseType 选择授权模式：ResponseTypeToken（隐式）或 ResponseTypeCode（授权码）。
	ResponseType string

	// ClientID is the OAuth2 client id. Defaults to DefaultClientID when empty.
	// ClientID 为 OAuth2 客户端 ID，为空时使用 DefaultClientID。
	ClientID string

	// RedirectURI is the callback URI. Defaults to DefaultRedirectURI when empty.
	// RedirectURI 为回调地址，为空时使用 DefaultRedirectURI。
	RedirectURI string

	// State is an arbitrary value returned verbatim after authorization, used for CSRF checks.
	// State 为任意自定义值，授权后会原样返回，用于防止 CSRF 攻击。
	State string

	// Scope lists the requested ESI scopes, separated by spaces.
	// At most 4 scopes may be requested at once, e.g.
	// "esi-wallet.read_character_wallet.v1 esi-assets.read_assets.v1".
	// Scope 为申请的 ESI 权限范围，使用空格分隔。一次最多申请 4 项，例如
	// "esi-wallet.read_character_wallet.v1 esi-assets.read_assets.v1"。
	Scope string

	// Realm and DeviceID are required by the NetEase SSO but their values are not verified.
	// Realm 与 DeviceID 为网易 SSO 的必填参数，但取值不受校验，可随意填写。
	Realm    string
	DeviceID string
}

func (cfg AuthorizeConfig) clientID() string {
	if cfg.ClientID == "" {
		return DefaultClientID
	}
	return cfg.ClientID
}

func (cfg AuthorizeConfig) redirectURI() string {
	if cfg.RedirectURI == "" {
		return DefaultRedirectURI
	}
	return cfg.RedirectURI
}

// BuildAuthorizeURL builds the URL the end user must open in a browser to authorize
// the application. After logging in and selecting a character, the user is
// redirected to the redirect_uri:
//   - implicit mode (response_type=token): the fragment contains
//     access_token=...&expires_in=...&state=...
//   - code mode (response_type=code): the query contains code=...&state=...
//
// BuildAuthorizeURL 构造需要用户在浏览器中打开的授权链接。用户登录并选择角色后，
// 将跳转到 redirect_uri：
//   - 隐式模式（response_type=token）：fragment 中带 access_token=...&expires_in=...&state=...
//   - 授权码模式（response_type=code）：query 中带 code=...&state=...
func BuildAuthorizeURL(cfg AuthorizeConfig) string {
	responseType := cfg.ResponseType
	if responseType == "" {
		responseType = ResponseTypeCode
	}
	realm := cfg.Realm
	if realm == "" {
		realm = "goeve"
	}
	deviceID := cfg.DeviceID
	if deviceID == "" {
		deviceID = "goeve"
	}
	query := url.Values{}
	query.Set("response_type", responseType)
	query.Set("client_id", cfg.clientID())
	query.Set("redirect_uri", cfg.redirectURI())
	query.Set("state", cfg.State)
	if cfg.Scope != "" {
		query.Set("scope", cfg.Scope)
	}
	query.Set("realm", realm)
	query.Set("device_id", deviceID)
	return SSOBaseURL + SSOAuthorizePath + "?" + query.Encode()
}

// tokenResponse is the JSON payload returned by the SSO token endpoint.
// tokenResponse 为 SSO 令牌端点返回的 JSON 负载。
type tokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"`
	Error        string `json:"error"`
}

// exchangeToken calls the SSO token endpoint with the given form parameters.
// exchangeToken 使用指定的表单参数调用 SSO 令牌端点。
func exchangeToken(ctx context.Context, form url.Values) (*Token, error) {
	var tr tokenResponse
	response, err := resty.New().SetTimeout(30*time.Second).R().
		SetContext(ctx).
		SetHeader("Content-Type", "application/x-www-form-urlencoded").
		SetBody(form.Encode()).
		SetResult(&tr).
		Post(SSOBaseURL + SSOTokenPath)
	if err != nil {
		return nil, fmt.Errorf("esi sso: token request failed: %w", err)
	}
	if response.IsError() {
		if tr.Error != "" {
			return nil, fmt.Errorf("esi sso: token endpoint returned status %d: %s", response.StatusCode(), tr.Error)
		}
		return nil, fmt.Errorf("esi sso: token endpoint returned status %d: %s", response.StatusCode(), response.String())
	}
	if tr.AccessToken == "" {
		return nil, errors.New("esi sso: token endpoint returned an empty access_token")
	}
	return &Token{
		AccessToken:  tr.AccessToken,
		RefreshToken: tr.RefreshToken,
		TokenType:    tr.TokenType,
		ExpiresIn:    tr.ExpiresIn,
		Expiry:       time.Now().Add(time.Duration(tr.ExpiresIn) * time.Second),
	}, nil
}

// GetTokenFromCode exchanges an authorization code for an access token plus
// refresh token. The code is valid for about 10 minutes and can be used only once.
//
// GetTokenFromCode 使用授权码换取 access_token 与 refresh_token。
// 授权码有效期约 10 分钟，且只能使用一次。
func GetTokenFromCode(ctx context.Context, code string, cfg AuthorizeConfig) (*Token, error) {
	if code == "" {
		return nil, errors.New("esi sso: authorization code is empty")
	}
	form := url.Values{}
	form.Set("grant_type", "authorization_code")
	form.Set("client_id", cfg.clientID())
	form.Set("redirect_uri", cfg.redirectURI())
	form.Set("code", code)
	return exchangeToken(ctx, form)
}

// RefreshAccessToken exchanges a refresh token for a new access token (and
// usually a new refresh token). Refresh tokens are long-lived; store them safely.
//
// RefreshAccessToken 使用 refresh_token 换取新的 access_token
// （通常同时返回新的 refresh_token）。refresh_token 长期有效，请妥善保存。
func RefreshAccessToken(ctx context.Context, refreshToken string) (*Token, error) {
	if refreshToken == "" {
		return nil, errors.New("esi sso: refresh token is empty")
	}
	form := url.Values{}
	form.Set("grant_type", "refresh_token")
	form.Set("client_id", DefaultClientID)
	form.Set("refresh_token", refreshToken)
	return exchangeToken(ctx, form)
}

// TokenSource supplies access tokens to a Client and refreshes them automatically
// in the background using the refresh token. It is safe for concurrent use.
//
// TokenSource 向 Client 提供 access_token，并在后台使用 refresh_token 自动续期。
// 可并发安全使用。
type TokenSource struct {
	mu           sync.Mutex
	refreshToken string
	token        *Token
}

// NewTokenSource creates a TokenSource from an existing token (typically the
// result of GetTokenFromCode or RefreshAccessToken).
//
// NewTokenSource 基于已有令牌创建 TokenSource（通常为
// GetTokenFromCode 或 RefreshAccessToken 的返回结果）。
func NewTokenSource(token *Token) *TokenSource {
	return &TokenSource{token: token, refreshToken: token.RefreshToken}
}

// NewTokenSourceFromRefreshToken creates a TokenSource that fetches the initial
// access token lazily from the given refresh token on first use.
//
// NewTokenSourceFromRefreshToken 创建 TokenSource，首次使用时才用给定的
// refresh_token 惰性获取初始 access_token。
func NewTokenSourceFromRefreshToken(refreshToken string) *TokenSource {
	return &TokenSource{refreshToken: refreshToken}
}

// Token returns a valid access token, refreshing it when expired or missing.
// Token 返回有效的 access_token，过期或缺失时自动续期。
func (ts *TokenSource) Token(ctx context.Context) (string, error) {
	ts.mu.Lock()
	defer ts.mu.Unlock()
	if ts.token != nil && ts.token.Valid() {
		return ts.token.AccessToken, nil
	}
	if ts.refreshToken == "" {
		return "", errors.New("esi sso: no valid token and no refresh token available")
	}
	token, err := RefreshAccessToken(ctx, ts.refreshToken)
	if err != nil {
		return "", err
	}
	ts.token = token
	if token.RefreshToken != "" {
		ts.refreshToken = token.RefreshToken
	}
	return token.AccessToken, nil
}

// CurrentToken returns the current token held by the source (may be expired).
// CurrentToken 返回当前持有的令牌（可能已过期）。
func (ts *TokenSource) CurrentToken() *Token {
	ts.mu.Lock()
	defer ts.mu.Unlock()
	return ts.token
}

// ParseImplicitRedirect extracts the access token from a redirect URL fragment
// produced by the implicit grant flow, e.g. the URL copied from the browser
// address bar: https://esi.evepc.163.com/ui/oauth2-redirect.html#access_token=...&expires_in=...&state=...
//
// ParseImplicitRedirect 从隐式模式产生的跳转地址 fragment 中提取 access_token，
// 例如从浏览器地址栏复制的地址：
// https://esi.evepc.163.com/ui/oauth2-redirect.html#access_token=...&expires_in=...&state=...
func ParseImplicitRedirect(redirectURL string) (*Token, error) {
	if redirectURL == "" {
		return nil, errors.New("esi sso: redirect URL is empty")
	}
	if !strings.Contains(redirectURL, "#") {
		return nil, errors.New("esi sso: redirect URL contains no fragment")
	}
	fragment := strings.SplitN(redirectURL, "#", 2)[1]
	values, err := url.ParseQuery(fragment)
	if err != nil {
		return nil, fmt.Errorf("esi sso: cannot parse redirect fragment: %w", err)
	}
	accessToken := values.Get("access_token")
	if accessToken == "" {
		return nil, errors.New("esi sso: fragment contains no access_token")
	}
	expiresIn := 0
	if v := values.Get("expires_in"); v != "" {
		if _, err := fmt.Sscanf(v, "%d", &expiresIn); err != nil {
			expiresIn = 0
		}
	}
	return &Token{
		AccessToken: accessToken,
		TokenType:   "Bearer",
		ExpiresIn:   expiresIn,
		Expiry:      time.Now().Add(time.Duration(expiresIn) * time.Second),
	}, nil
}

// ParseAuthorizationCode extracts the authorization code from a redirect URL
// produced by the authorization code flow, e.g.
// https://esi.evepc.163.com/ui/oauth2-redirect.html?code=...&state=...
//
// ParseAuthorizationCode 从授权码模式产生的跳转地址中提取授权码，例如：
// https://esi.evepc.163.com/ui/oauth2-redirect.html?code=...&state=...
func ParseAuthorizationCode(redirectURL string) (code string, err error) {
	if redirectURL == "" {
		return "", errors.New("esi sso: redirect URL is empty")
	}
	parsed, err := url.Parse(redirectURL)
	if err != nil {
		return "", fmt.Errorf("esi sso: cannot parse redirect URL: %w", err)
	}
	code = parsed.Query().Get("code")
	if code == "" {
		return "", errors.New("esi sso: redirect URL contains no code")
	}
	return code, nil
}
