package integration_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"
)

var _ = Describe("BoyOhBoy", func() {

	It("exits with a successful code", func() {
		session := execBOB()
		Eventually(session).Should(gexec.Exit(0))
	})
})
