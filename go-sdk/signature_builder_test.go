package go_sdk

import (
	"context"
	"net/http"
	"strings"
	"testing"
)

func TestSigRequest(t *testing.T) {
	sb := NewSignatureBuilder("ak", "sk", 3600)
	body := strings.NewReader("a=b&d=c")
	r, _ := http.NewRequest("POST", "https://api.open-platform.com/v1/photos/generate?data=a", body)

	ts, err := sb.SignRequest(context.Background(), r)
	if err != nil {
		t.Error(err)
	}
	t.Log(*ts)
	t.Log(r.Header)

	err = sb.ValidateRequest(context.Background(), r)
	if err != nil {
		t.Error(err)
	}
}
