package netstorage_test

import (
	"testing"

	"github.com/takebox/rclone/backend/netstorage"
	"github.com/takebox/rclone/fstest/fstests"
)

// TestIntegration runs integration tests against the remote
func TestIntegration(t *testing.T) {
	fstests.Run(t, &fstests.Opt{
		RemoteName: "TestnStorage:",
		NilObject:  (*netstorage.Object)(nil),
	})
}
