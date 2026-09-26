package goeve

import (
	"context"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
	"time"
)

// TestGetServerStatus performs a real network call against the public /status endpoint.
// TestGetServerStatus 对公开的 /status 接口发起真实网络调用。
func TestGetServerStatus(t *testing.T) {
	client := NewClient(WithTimeout(30 * time.Second))
	status, err := client.GetServerStatus(context.Background())
	if err != nil {
		t.Fatalf("GetServerStatus failed: %v", err)
	}
	if status.ServerVersion == "" {
		t.Fatal("server_version is empty")
	}
	t.Logf("server_version=%s players=%d vip=%v start_time=%s",
		status.ServerVersion, status.Players, status.Vip, status.StartTime)
}

// TestGetCharacterNotFound checks that querying a non-existent character
// returns a structured 404 APIError.
// TestGetCharacterNotFound 检查查询不存在的角色时返回结构化的 404 APIError。
func TestGetCharacterNotFound(t *testing.T) {
	client := NewClient(WithTimeout(30 * time.Second))
	_, err := client.GetCharacter(context.Background(), 1)
	if err == nil {
		t.Fatal("expected an error")
	}
	apiErr, ok := err.(*APIError)
	if !ok {
		t.Fatalf("expected *APIError, got %T: %v", err, err)
	}
	if apiErr.StatusCode != 404 {
		t.Fatalf("expected status 404, got %d", apiErr.StatusCode)
	}
	t.Logf("api error: %s", apiErr)
}

// TestGetUniverseCategories fetches the public universe category list.
// TestGetUniverseCategories 获取公开的宇宙分类列表。
func TestGetUniverseCategories(t *testing.T) {
	client := NewClient(WithTimeout(30 * time.Second))
	ctx := context.Background()

	categories, err := client.GetUniverseCategories(ctx)
	if err != nil {
		t.Fatalf("GetUniverseCategories failed: %v", err)
	}
	if len(categories) == 0 {
		t.Fatal("no categories returned")
	}
	t.Logf("categories=%d", len(categories))

	// Conditional request: when nothing changed the server returns 304 and
	// the result stays nil — this must not be reported as an error.
	// 条件请求：数据未变化时服务器返回 304，结果保持为 nil，不应报错。
	etag := `W/"e224bfe767a9e9f3fa0e4615aac73f2c3a63fd0c792e553d5f203e03"`
	result, err := client.GetUniverseCategories(ctx, WithIfNoneMatch(etag))
	if err != nil {
		t.Fatalf("GetUniverseCategories with If-None-Match failed: %v", err)
	}
	if result != nil {
		t.Fatalf("expected nil result on 304, got %v", result)
	}
}

// TestAPIError checks that an invalid id is surfaced as a structured APIError.
// TestAPIError 检查无效 ID 是否以结构化的 APIError 返回。
func TestAPIError(t *testing.T) {
	client := NewClient(WithTimeout(30 * time.Second))
	_, err := client.GetCharacter(context.Background(), 0)
	if err == nil {
		t.Fatal("expected an error")
	}
	apiErr, ok := err.(*APIError)
	if !ok {
		t.Fatalf("expected *APIError, got %T: %v", err, err)
	}
	if apiErr.StatusCode != 400 {
		t.Fatalf("expected status 400, got %d", apiErr.StatusCode)
	}
	t.Logf("api error: %s", apiErr)
}

// TestBuildAuthorizeURL verifies the SSO authorization URL construction.
// TestBuildAuthorizeURL 验证 SSO 授权链接的构造。
func TestBuildAuthorizeURL(t *testing.T) {
	got := BuildAuthorizeURL(AuthorizeConfig{
		ResponseType: ResponseTypeCode,
		State:        "test-state",
		Scope:        "esi-wallet.read_character_wallet.v1",
	})
	if !strings.HasPrefix(got, SSOBaseURL+SSOAuthorizePath+"?") {
		t.Fatalf("unexpected authorize URL prefix: %s", got)
	}
	values, err := url.ParseQuery(strings.SplitN(got, "?", 2)[1])
	if err != nil {
		t.Fatalf("cannot parse query: %v", err)
	}
	checks := map[string]string{
		"response_type": "code",
		"client_id":     DefaultClientID,
		"redirect_uri":  DefaultRedirectURI,
		"state":         "test-state",
		"scope":         "esi-wallet.read_character_wallet.v1",
	}
	for key, want := range checks {
		if values.Get(key) != want {
			t.Errorf("query %q = %q, want %q", key, values.Get(key), want)
		}
	}
	if values.Get("realm") == "" || values.Get("device_id") == "" {
		t.Error("realm and device_id must be set")
	}
}

// TestParseImplicitRedirect verifies parsing of the implicit grant redirect URL.
// TestParseImplicitRedirect 验证隐式模式跳转地址的解析。
func TestParseImplicitRedirect(t *testing.T) {
	redirect := "https://esi.evepc.163.com/ui/oauth2-redirect.html#access_token=abc123&expires_in=1199&state=xyz"
	token, err := ParseImplicitRedirect(redirect)
	if err != nil {
		t.Fatalf("ParseImplicitRedirect failed: %v", err)
	}
	if token.AccessToken != "abc123" || token.ExpiresIn != 1199 {
		t.Fatalf("unexpected token: %+v", token)
	}
	if !token.Valid() {
		t.Fatal("token should be valid")
	}
}

// TestParseAuthorizationCode verifies parsing of the authorization code redirect URL.
// TestParseAuthorizationCode 验证授权码模式跳转地址的解析。
func TestParseAuthorizationCode(t *testing.T) {
	redirect := "https://esi.evepc.163.com/ui/oauth2-redirect.html?code=the-code&state=xyz"
	code, err := ParseAuthorizationCode(redirect)
	if err != nil {
		t.Fatalf("ParseAuthorizationCode failed: %v", err)
	}
	if code != "the-code" {
		t.Fatalf("unexpected code: %s", code)
	}
}

// TestTokenValid checks token expiry logic.
// TestTokenValid 检查令牌过期逻辑。
func TestTokenValid(t *testing.T) {
	if (&Token{AccessToken: "x", Expiry: time.Now().Add(time.Minute)}).Valid() != true {
		t.Fatal("fresh token should be valid")
	}
	if (&Token{AccessToken: "x", Expiry: time.Now().Add(-time.Hour)}).Valid() != false {
		t.Fatal("expired token should be invalid")
	}
	if (&Token{Expiry: time.Now().Add(time.Hour)}).Valid() != false {
		t.Fatal("token without access token should be invalid")
	}
}

// TestWithAuthTokenPerRequest verifies that WithAuthToken sends the token only
// for the request that carries the option.
// TestWithAuthTokenPerRequest 验证 WithAuthToken 只为携带该选项的请求发送令牌。
func TestWithAuthTokenPerRequest(t *testing.T) {
	var gotAuth []string
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotAuth = append(gotAuth, r.Header.Get("Authorization"))
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"players":1,"server_version":"test","start_time":"2026-09-27T00:00:00Z","vip":false}`))
	}))
	defer server.Close()

	client := NewClient(WithBaseURL(server.URL))

	// without the option: no Authorization header / 不带选项：无 Authorization 头
	if _, err := client.GetServerStatus(context.Background()); err != nil {
		t.Fatalf("unauthenticated request failed: %v", err)
	}
	// with the option: Bearer token sent / 带选项：发送 Bearer 令牌
	if _, err := client.GetServerStatus(context.Background(), WithAuthToken("abc123")); err != nil {
		t.Fatalf("authenticated request failed: %v", err)
	}

	if len(gotAuth) != 2 {
		t.Fatalf("expected 2 requests, got %d", len(gotAuth))
	}
	if gotAuth[0] != "" {
		t.Errorf("request without option should not carry Authorization, got %q", gotAuth[0])
	}
	if gotAuth[1] != "Bearer abc123" {
		t.Errorf("Authorization = %q, want %q", gotAuth[1], "Bearer abc123")
	}
}
