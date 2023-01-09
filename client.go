package lotr_sdk

import (
	"fmt"
	"os"

	"github.com/imroc/req/v3"
)

type PaginationStats struct {
	Total  int `json:"total"`
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
	Page   int `json:"page"`
	Pages  int `json:"pages"`
}

type genericResponse[T any] struct {
	Docs []T `json:"docs"`
	PaginationStats
}

// LOTRClient is the go client for the Lord of the Rings api https://github.com/gitfrosh/lotr-api.
type LOTRClient struct {
	*req.Client
}

// NewLOTRClient creates and returns a new pointer to a new LOTRClient
func NewLOTRClient() *LOTRClient {
	authToken := os.Getenv("LOTR_ACCESS_TOKEN")
	if authToken == "" {
		fmt.Println("Warning: No access token found. Endpoints requiring authentication will not work.")
	}

	c := req.C().
		SetBaseURL("https://the-one-api.dev/v2").
		SetCommonBearerAuthToken(authToken).
		// EnableDump at the request level in request middleware which dump content into
		// memory (not print to stdout), we can record dump content only when unexpected
		// exception occurs, it is helpful to troubleshoot problems in production.
		OnBeforeRequest(func(c *req.Client, r *req.Request) error {
			if r.RetryAttempt > 0 { // Ignore on retry.
				return nil
			}
			r.EnableDump()
			return nil
		}).
		// Handle common exceptions in response middleware.
		OnAfterResponse(func(client *req.Client, resp *req.Response) error {
			if resp.Err != nil {
				return resp.Err
			}
			// Corner case: neither an error response nor a success response,
			// dump content to help troubleshoot.
			if !resp.IsSuccess() {
				return fmt.Errorf("bad response, raw dump:\n%s", resp.Dump())
			}
			return nil
		})

	return &LOTRClient{
		Client: c,
	}
}

// SetDebug enable debug logging if set to true, disable debug if set to false.
func (c *LOTRClient) SetDebug(enable bool) *LOTRClient {
	if enable {
		c.EnableDebugLog()
		c.EnableDumpAll()
	} else {
		c.DisableDebugLog()
		c.DisableDumpAll()
	}
	return c
}
