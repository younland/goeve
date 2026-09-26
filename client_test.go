package goeve

import (
	"context"
	"github.com/younland/goeve/models"
	"net/url"
	"strings"
	"testing"
	"time"
)

// TestGetStatus performs a real network call against the public /status endpoint.
// TestGetStatus 对公开的 /status 接口发起真实网络调用。
func TestGetStatus(t *testing.T) {
	client := NewClient(WithTimeout(30 * time.Second))
	status, err := client.GetStatus(context.Background(), nil)
	if err != nil {
		t.Fatalf("GetStatus failed: %v", err)
	}
	if status.ServerVersion == "" {
		t.Fatal("server_version is empty")
	}
	t.Logf("server_version=%s players=%d vip=%v start_time=%s",
		status.ServerVersion, status.Players, status.Vip, status.StartTime)
}

// TestGetCharactersCharacterIdNotFound checks that querying a non-existent character
// returns a structured 404 APIError.
// TestGetCharactersCharacterIdNotFound 检查查询不存在的角色时返回结构化的 404 APIError。
func TestGetCharactersCharacterIdNotFound(t *testing.T) {
	client := NewClient(WithTimeout(30 * time.Second))
	_, err := client.GetCharacterId(context.Background(), 1, nil)
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

	categories, err := client.GetCategories(ctx, nil)
	if err != nil {
		t.Fatalf("GetCategories failed: %v", err)
	}
	if len(categories) == 0 {
		t.Fatal("no categories returned")
	}
	t.Logf("categories=%d", len(categories))

	// Conditional request: when nothing changed the server returns 304 and
	// the result stays nil — this must not be reported as an error.
	// 条件请求：数据未变化时服务器返回 304，结果保持为 nil，不应报错。
	etag := `W/"e224bfe767a9e9f3fa0e4615aac73f2c3a63fd0c792e553d5f203e03"`
	params := &models.GetCategoriesParams{IfNoneMatch: &etag}
	result, err := client.GetCategories(ctx, params)
	if err != nil {
		t.Fatalf("GetCategories with If-None-Match failed: %v", err)
	}
	if result != nil {
		t.Fatalf("expected nil result on 304, got %v", result)
	}
}

// TestAPIError checks that an invalid id is surfaced as a structured APIError.
// TestAPIError 检查无效 ID 是否以结构化的 APIError 返回。
func TestAPIError(t *testing.T) {
	client := NewClient(WithTimeout(30 * time.Second))
	_, err := client.GetCharacterId(context.Background(), 0, nil)
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
