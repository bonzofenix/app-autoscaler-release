package quota_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"testing"
)

func TestQuota(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Quota Suite")
}
