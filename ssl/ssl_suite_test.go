package ssl_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestSsl(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Ssl Suite")
}
