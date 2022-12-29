package types_test

import (
	"os"
	"testing"

	"github.com/furya-official/mage/app"
)

func TestMain(m *testing.M) {
	app.SetSDKConfig()
	os.Exit(m.Run())
}
