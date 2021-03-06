package sentry

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProjectKeyService_List(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/api/0/projects/the-interstellar-jurisdiction/pump-station/keys/", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "GET", r)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `[{
			"browserSdk": {
				"choices": [
					[
						"latest",
						"latest"
					],
					[
						"4.x",
						"4.x"
					]
				]
			},
			"browserSdkVersion": "4.x",
			"dateCreated": "2018-09-20T15:48:07.397Z",
			"dsn": {
				"cdn": "https://sentry.io/js-sdk-loader/cfc7b0341c6e4f6ea1a9d256a30dba00.min.js",
				"csp": "https://sentry.io/api/2/csp-report/?sentry_key=cfc7b0341c6e4f6ea1a9d256a30dba00",
				"minidump": "https://sentry.io/api/2/minidump/?sentry_key=cfc7b0341c6e4f6ea1a9d256a30dba00",
				"public": "https://cfc7b0341c6e4f6ea1a9d256a30dba00@sentry.io/2",
				"secret": "https://cfc7b0341c6e4f6ea1a9d256a30dba00:a07dcd97aa56481f82aeabaed43ca448@sentry.io/2",
				"security": "https://sentry.io/api/2/security/?sentry_key=cfc7b0341c6e4f6ea1a9d256a30dba00"
			},
			"id": "cfc7b0341c6e4f6ea1a9d256a30dba00",
			"isActive": true,
			"label": "Fabulous Key",
			"name": "Fabulous Key",
			"projectId": 2,
			"public": "cfc7b0341c6e4f6ea1a9d256a30dba00",
			"rateLimit": null,
			"secret": "a07dcd97aa56481f82aeabaed43ca448"
		}]`)
	})

	client := NewClient(httpClient, nil, "")
	projectKeys, _, err := client.ProjectKeys.List("the-interstellar-jurisdiction", "pump-station")
	assert.NoError(t, err)

	expected := []ProjectKey{
		{
			ID:        "cfc7b0341c6e4f6ea1a9d256a30dba00",
			Name:      "Fabulous Key",
			Label:     "Fabulous Key",
			Public:    "cfc7b0341c6e4f6ea1a9d256a30dba00",
			Secret:    "a07dcd97aa56481f82aeabaed43ca448",
			ProjectID: 2,
			IsActive:  true,
			DSN: ProjectKeyDSN{
				Secret:   "https://cfc7b0341c6e4f6ea1a9d256a30dba00:a07dcd97aa56481f82aeabaed43ca448@sentry.io/2",
				Public:   "https://cfc7b0341c6e4f6ea1a9d256a30dba00@sentry.io/2",
				CSP:      "https://sentry.io/api/2/csp-report/?sentry_key=cfc7b0341c6e4f6ea1a9d256a30dba00",
				Security: "https://sentry.io/api/2/security/?sentry_key=cfc7b0341c6e4f6ea1a9d256a30dba00",
				Minidump: "https://sentry.io/api/2/minidump/?sentry_key=cfc7b0341c6e4f6ea1a9d256a30dba00",
				CDN:      "https://sentry.io/js-sdk-loader/cfc7b0341c6e4f6ea1a9d256a30dba00.min.js",
			},
			DateCreated: mustParseTime("2018-09-20T15:48:07.397Z"),
		},
	}
	assert.Equal(t, expected, projectKeys)
}

func TestProjectKeyService_Create(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/api/0/projects/the-interstellar-jurisdiction/pump-station/keys/", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "POST", r)
		assertPostJSON(t, map[string]interface{}{
			"name": "Fabulous Key",
		}, r)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{
			"browserSdk": {
				"choices": [
					[
						"latest",
						"latest"
					],
					[
						"4.x",
						"4.x"
					]
				]
			},
			"browserSdkVersion": "4.x",
			"dateCreated": "2018-09-20T15:48:07.397Z",
			"dsn": {
				"cdn": "https://sentry.io/js-sdk-loader/cfc7b0341c6e4f6ea1a9d256a30dba00.min.js",
				"csp": "https://sentry.io/api/2/csp-report/?sentry_key=cfc7b0341c6e4f6ea1a9d256a30dba00",
				"minidump": "https://sentry.io/api/2/minidump/?sentry_key=cfc7b0341c6e4f6ea1a9d256a30dba00",
				"public": "https://cfc7b0341c6e4f6ea1a9d256a30dba00@sentry.io/2",
				"secret": "https://cfc7b0341c6e4f6ea1a9d256a30dba00:a07dcd97aa56481f82aeabaed43ca448@sentry.io/2",
				"security": "https://sentry.io/api/2/security/?sentry_key=cfc7b0341c6e4f6ea1a9d256a30dba00"
			},
			"id": "cfc7b0341c6e4f6ea1a9d256a30dba00",
			"isActive": true,
			"label": "Fabulous Key",
			"name": "Fabulous Key",
			"projectId": 2,
			"public": "cfc7b0341c6e4f6ea1a9d256a30dba00",
			"rateLimit": null,
			"secret": "a07dcd97aa56481f82aeabaed43ca448"
		}`)
	})

	client := NewClient(httpClient, nil, "")
	params := &CreateProjectKeyParams{
		Name: "Fabulous Key",
	}
	projectKey, _, err := client.ProjectKeys.Create("the-interstellar-jurisdiction", "pump-station", params)
	assert.NoError(t, err)
	expected := &ProjectKey{
		ID:        "cfc7b0341c6e4f6ea1a9d256a30dba00",
		Name:      "Fabulous Key",
		Label:     "Fabulous Key",
		Public:    "cfc7b0341c6e4f6ea1a9d256a30dba00",
		Secret:    "a07dcd97aa56481f82aeabaed43ca448",
		ProjectID: 2,
		IsActive:  true,
		DSN: ProjectKeyDSN{
			Secret:   "https://cfc7b0341c6e4f6ea1a9d256a30dba00:a07dcd97aa56481f82aeabaed43ca448@sentry.io/2",
			Public:   "https://cfc7b0341c6e4f6ea1a9d256a30dba00@sentry.io/2",
			CSP:      "https://sentry.io/api/2/csp-report/?sentry_key=cfc7b0341c6e4f6ea1a9d256a30dba00",
			Security: "https://sentry.io/api/2/security/?sentry_key=cfc7b0341c6e4f6ea1a9d256a30dba00",
			Minidump: "https://sentry.io/api/2/minidump/?sentry_key=cfc7b0341c6e4f6ea1a9d256a30dba00",
			CDN:      "https://sentry.io/js-sdk-loader/cfc7b0341c6e4f6ea1a9d256a30dba00.min.js",
		},
		DateCreated: mustParseTime("2018-09-20T15:48:07.397Z"),
	}
	assert.Equal(t, expected, projectKey)
}

func TestProjectKeyService_Update(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/api/0/projects/the-interstellar-jurisdiction/pump-station/keys/befdbf32724c4ae0a3d286717b1f8127/", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "PUT", r)
		assertPostJSON(t, map[string]interface{}{
			"name": "Fabulous Key",
		}, r)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{
			"browserSdk": {
				"choices": [
					[
						"latest",
						"latest"
					],
					[
						"4.x",
						"4.x"
					]
				]
			},
			"browserSdkVersion": "4.x",
			"dateCreated": "2018-09-20T15:48:07.397Z",
			"dsn": {
				"cdn": "https://sentry.io/js-sdk-loader/cfc7b0341c6e4f6ea1a9d256a30dba00.min.js",
				"csp": "https://sentry.io/api/2/csp-report/?sentry_key=cfc7b0341c6e4f6ea1a9d256a30dba00",
				"minidump": "https://sentry.io/api/2/minidump/?sentry_key=cfc7b0341c6e4f6ea1a9d256a30dba00",
				"public": "https://cfc7b0341c6e4f6ea1a9d256a30dba00@sentry.io/2",
				"secret": "https://cfc7b0341c6e4f6ea1a9d256a30dba00:a07dcd97aa56481f82aeabaed43ca448@sentry.io/2",
				"security": "https://sentry.io/api/2/security/?sentry_key=cfc7b0341c6e4f6ea1a9d256a30dba00"
			},
			"id": "cfc7b0341c6e4f6ea1a9d256a30dba00",
			"isActive": true,
			"label": "Fabulous Key",
			"name": "Fabulous Key",
			"projectId": 2,
			"public": "cfc7b0341c6e4f6ea1a9d256a30dba00",
			"rateLimit": null,
			"secret": "a07dcd97aa56481f82aeabaed43ca448"
		}`)
	})

	client := NewClient(httpClient, nil, "")
	params := &UpdateProjectKeyParams{
		Name: "Fabulous Key",
	}
	projectKey, _, err := client.ProjectKeys.Update("the-interstellar-jurisdiction", "pump-station", "befdbf32724c4ae0a3d286717b1f8127", params)
	assert.NoError(t, err)
	expected := &ProjectKey{
		ID:        "cfc7b0341c6e4f6ea1a9d256a30dba00",
		Name:      "Fabulous Key",
		Label:     "Fabulous Key",
		Public:    "cfc7b0341c6e4f6ea1a9d256a30dba00",
		Secret:    "a07dcd97aa56481f82aeabaed43ca448",
		ProjectID: 2,
		IsActive:  true,
		DSN: ProjectKeyDSN{
			Secret:   "https://cfc7b0341c6e4f6ea1a9d256a30dba00:a07dcd97aa56481f82aeabaed43ca448@sentry.io/2",
			Public:   "https://cfc7b0341c6e4f6ea1a9d256a30dba00@sentry.io/2",
			CSP:      "https://sentry.io/api/2/csp-report/?sentry_key=cfc7b0341c6e4f6ea1a9d256a30dba00",
			Security: "https://sentry.io/api/2/security/?sentry_key=cfc7b0341c6e4f6ea1a9d256a30dba00",
			Minidump: "https://sentry.io/api/2/minidump/?sentry_key=cfc7b0341c6e4f6ea1a9d256a30dba00",
			CDN:      "https://sentry.io/js-sdk-loader/cfc7b0341c6e4f6ea1a9d256a30dba00.min.js",
		},
		DateCreated: mustParseTime("2018-09-20T15:48:07.397Z"),
	}
	assert.Equal(t, expected, projectKey)
}

func TestProjectKeyService_Delete(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/api/0/projects/the-interstellar-jurisdiction/pump-station/keys/befdbf32724c4ae0a3d286717b1f8127/", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "DELETE", r)
	})

	client := NewClient(httpClient, nil, "")
	_, err := client.ProjectKeys.Delete("the-interstellar-jurisdiction", "pump-station", "befdbf32724c4ae0a3d286717b1f8127")
	assert.NoError(t, err)

}
