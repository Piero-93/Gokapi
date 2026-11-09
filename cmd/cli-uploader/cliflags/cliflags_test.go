package cliflags

import (
	"os"
	"testing"

	"github.com/forceu/gokapi/internal/test"
)

func TestExpiryParametersSet(t *testing.T) {
	originalArgs := os.Args
	defer func() { os.Args = originalArgs }()

	os.Args = []string{"gokapi-cli", "upload", "-f", "file.txt"}
	result := GetUploadParameters(ModeUpload)
	test.IsEqualBool(t, result.IsExpiryDaysSet, false)
	test.IsEqualBool(t, result.IsExpiryDownloadsSet, false)

	os.Args = []string{"gokapi-cli", "upload", "-f", "file.txt", "--expiry-days", "0", "-d", "3"}
	result = GetUploadParameters(ModeUpload)
	test.IsEqualBool(t, result.IsExpiryDaysSet, true)
	test.IsEqualBool(t, result.IsExpiryDownloadsSet, true)
	test.IsEqualInt(t, result.ExpiryDays, 0)
	test.IsEqualInt(t, result.ExpiryDownloads, 3)
}
