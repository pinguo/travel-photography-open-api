package go_sdk

import (
	"bytes"
	"context"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"mime"
	"net/http"
	"net/url"
	"sort"
	"strconv"
	"strings"
	"time"
)

func NewSignatureBuilder(ak, sk string, expiredIn int) *SignatureBuilder {
	return &SignatureBuilder{
		accessKey: ak,
		secretKey: sk,
		expiredIn: expiredIn,
	}
}

type SignatureResult struct {
	FinalText string
	Timestamp string
	Sign      string
}

type SignatureBuilder struct {
	accessKey string
	secretKey string
	expiredIn int
}

func (s *SignatureBuilder) ValidateRequest(ctx context.Context, r *http.Request) error {
	ts := r.Header.Get("PG-Timestamp")
	timestamp, err := strconv.ParseInt(ts, 10, 64)
	if err != nil {
		return err
	}
	if s.expiredIn > 0 && int(time.Now().Unix()-timestamp) > s.expiredIn {
		return fmt.Errorf("timestamp expired")
	}

	rs, err := s.buildSignatureFromIncomeRequest(ctx, r)
	if err != nil {
		return err
	}
	if rs.Sign != r.Header.Get("PG-Sign") {
		return fmt.Errorf("signature validation failed")
	}
	return nil
}

func (s *SignatureBuilder) SignRequest(ctx context.Context, r *http.Request) error {
	rs, err := s.buildSignatureFromIncomeRequest(ctx, r)
	if err != nil {
		return err
	}

	r.Header.Set("PG-Timestamp", rs.Timestamp)
	r.Header.Set("PG-AccessKey", s.accessKey)
	r.Header.Set("PG-Sign", rs.Sign)
	return nil
}

func (s *SignatureBuilder) buildSignatureFromIncomeRequest(ctx context.Context, r *http.Request) (*SignatureResult, error) {
	queryParams := s.getGETParams(r)
	body, postParams, err := s.getPOSTParams(r)
	if err != nil {
		return nil, err
	}
	if len(postParams) > 0 {
		for k, v := range postParams {
			queryParams[k] = v
		}
	}
	ts := fmt.Sprintf("%d", time.Now().Unix())
	path := r.URL.Path
	finalText := fmt.Sprintf("%s%s%s%s", path, s.buildParamsSignatureText(queryParams, body), ts, s.secretKey)
	sign := s.md5Hash(finalText)
	return &SignatureResult{
		Sign:      sign,
		FinalText: finalText,
		Timestamp: ts,
	}, nil
}

// 计算 MD5 哈希值
func (s *SignatureBuilder) md5Hash(text string) string {
	hash := md5.New()
	hash.Write([]byte(text))
	return hex.EncodeToString(hash.Sum(nil))
}

func (s *SignatureBuilder) buildParamsSignatureText(params map[string]string, body string) string {
	// 获取所有键并排序
	keys := make([]string, 0, len(params))
	for k := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	// 拼接参数和值
	var paramStr string
	for _, k := range keys {
		paramStr += fmt.Sprintf("%s=%s", k, params[k])
	}
	return paramStr + body
}
func (s *SignatureBuilder) getGETParams(r *http.Request) map[string]string {
	params := make(map[string]string)
	for key, values := range r.URL.Query() {
		if len(values) > 0 {
			params[key] = values[0]
		}
	}
	return params
}

// 从 HTTP POST 请求中提取参数
func (s *SignatureBuilder) getPOSTParams(r *http.Request) (string, map[string]string, error) {
	const defaultContentType = "application/x-www-form-urlencoded"
	ct := r.Header.Get("Content-Type")
	if ct == "" {
		ct = defaultContentType
	}
	ct, _, err := mime.ParseMediaType(ct)
	if err != nil {
		return "", nil, fmt.Errorf("failed to parse content type: %v", err)
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		return "", nil, fmt.Errorf("failed to read request body: %v", err)
	}

	r.Body = io.NopCloser(bytes.NewBuffer(body))
	if strings.Contains(ct, "application/json") {
		// 把读取的内容重新放回 Body 中，以便后续处理不影响原请求
		return string(body), nil, nil
	}
	// 如果是其他类型，尝试解析为表单数据
	if strings.Contains(ct, defaultContentType) {
		values, err := url.ParseQuery(string(body))
		if err != nil {
			return "", nil, fmt.Errorf("failed to parse form data: %v", err)
		}
		params := make(map[string]string)
		for k, _ := range values {
			params[k] = values.Get(k)
		}
		return "", params, nil
	}
	return "", nil, fmt.Errorf("unsupported content type: " + ct)
}
