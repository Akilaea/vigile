//go:build testing

package hub

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsAppRoute(t *testing.T) {
	tests := []struct {
		path     string
		basePath string
		want     bool
	}{
		// known routes
		{"/", "/", true},
		{"/containers", "/", true},
		{"/containers/", "/", true},
		{"/Containers", "/", true},
		{"/smart", "/", true},
		{"/monitors", "/", true},
		{"/forgot-password", "/", true},
		{"/request-otp", "/", true},
		{"/system/abc123", "/", true},
		{"/system/abc123/", "/", true},
		{"/settings", "/", true},
		{"/settings/general", "/", true},

		// unknown paths
		{"/.env", "/", false},
		{"/phpinfo.php", "/", false},
		{"/wp-admin/", "/", false},
		{"/.git/config", "/", false},
		{"/system", "/", false},
		{"/system/", "/", false},
		{"/system/abc/def", "/", false},
		{"/settings/general/extra", "/", false},
		{"/containersx", "/", false},

		// base path, prefix not stripped by proxy
		{"/vigile", "/vigile/", true},
		{"/vigile/", "/vigile/", true},
		{"/vigile/containers", "/vigile/", true},
		{"/vigile/system/abc123", "/vigile/", true},
		{"/vigile/.env", "/vigile/", false},
		{"/vigilex", "/vigile/", false},

		// base path, prefix stripped by proxy
		{"/", "/vigile/", true},
		{"/containers", "/vigile/", true},
		{"/.env", "/vigile/", false},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.want, isAppRoute(tt.path, tt.basePath), "path=%q basePath=%q", tt.path, tt.basePath)
	}
}
