package common

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path"

	"github.com/pulumi/pulumi-go-provider/infer"
)

// buildURL constructs the full URL for a resource using the base admin URL and the resource's path template
func (b BaseResource[I, O]) buildURL(ctx context.Context, ID string) (string, error) {
	config := infer.GetConfig[Config](ctx)
	u, err := url.Parse(config.OtoroshiAdminUrl)
	if err != nil {
		return "", err
	}
	elements := []string{u.Path, b.Path}
	if ID != "" {
		elements = append(elements, ID)
	}
	u.Path = path.Join(elements...)

	return u.String(), nil
}

// doRequest sends an HTTP request to the Otoroshi API
func (b BaseResource[I, O]) doRequest(ctx context.Context, method, ID string, body any) (*http.Response, error) {
	config := infer.GetConfig[Config](ctx)
	u, err := b.buildURL(ctx, ID)
	if err != nil {
		return nil, err
	}
	var bodyReader io.Reader
	if body != nil {
		data, err := json.Marshal(body)
		if err != nil {
			return nil, err
		}
		bodyReader = bytes.NewReader(data)
	}
	req, err := http.NewRequestWithContext(ctx, method, u, bodyReader)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(config.OtoroshiAdminClientID, config.OtoroshiAdminClientSecret)
	return http.DefaultClient.Do(req)
}

func (b BaseResource[I, O]) doRequestAndGetResponse(ctx context.Context, method, id string, body any) (O, error) {
	var output O
	resp, err := b.doRequest(ctx, method, id, body)
	if err != nil {
		return output, err
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusNotFound && (method == http.MethodGet || method == http.MethodDelete) {
		return output, nil
	}
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		body, _ := io.ReadAll(resp.Body)
		return output, fmt.Errorf("unexpected status %d: %s", resp.StatusCode, body)
	}
	if err := json.NewDecoder(resp.Body).Decode(&output); err != nil {
		return output, err
	}
	return output, nil
}
