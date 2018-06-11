package integration_test

import (
	"os/exec"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"
	"github.com/onsi/gomega/gexec"
)

var _ = Describe("BoyOhBoy", func() {

	var (
		bobCmd *exec.Cmd

		session *gexec.Session
	)

	BeforeEach(func() {
		bobCmd = exec.Command(boyOhBoyPath)
		bobCmd.Args = append(bobCmd.Args, "PVTL")
	})

	JustBeforeEach(func() {
		session = execBin(bobCmd)
	})

	It("exits with a successful code", func() {
		Eventually(session).Should(gexec.Exit(0))
	})

	It("prints some price in dollars", func() {
		Eventually(session).Should(gbytes.Say(`[$]\d+\.\d+`))
	})

	Context("when no ticker is provided", func() {
		BeforeEach(func() {
			bobCmd.Args = []string{}
		})

		It("exits with a code of 1 to signal no retry", func() {
			Eventually(session).Should(gexec.Exit(1))
		})

		It("prints a useful error to stderr", func() {
			Eventually(session.Err).Should(gbytes.Say("please provide a ticker to fetch"))
		})
	})
})
