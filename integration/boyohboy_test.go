package integration_test

import (
	"os/exec"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"
	"github.com/onsi/gomega/gexec"
)

var _ = Describe("BoyOhBoy fetch", func() {

	var (
		bobCmd *exec.Cmd

		session *gexec.Session
	)

	BeforeEach(func() {
		bobCmd = exec.Command(boyOhBoyPath)
	})

	JustBeforeEach(func() {
		session = execBin(bobCmd)
	})

	Context("when not all the args are provided", func() {
		It("exits with a code of 1 to signal no retry", func() {
			Eventually(session).Should(gexec.Exit(1))
		})

		It("prints the usage", func() {
			Eventually(session.Err).Should(gbytes.Say("please provide an action"))
		})
	})

	Context("when the action provided is not 'fetch'", func() {
		BeforeEach(func() {
			bobCmd.Args = append(bobCmd.Args, "not-fetch")
		})

		It("exits with a code of 1 to signal no retry", func() {
			Eventually(session).Should(gexec.Exit(1))
		})

		It("prints a useful error to stderr", func() {
			Eventually(session.Err).Should(gbytes.Say("please provide 'fetch' as the action"))
		})
	})

	Context("when no args are provided", func() {
		BeforeEach(func() {
			bobCmd.Args = append(bobCmd.Args, "fetch")
		})

		It("exits with a code of 1 to signal no retry", func() {
			Eventually(session).Should(gexec.Exit(1))
		})

		It("prints a useful error to stderr", func() {
			Eventually(session.Err).Should(gbytes.Say("please provide args to fetch"))
		})
	})

	Context("when the args json is malformed", func() {
		BeforeEach(func() {
			bobCmd.Args = append(bobCmd.Args, "fetch", "{malformed}}")
		})

		It("exits with a code of 1 to signal no retry", func() {
			Eventually(session).Should(gexec.Exit(1))
		})

		It("prints a useful error to stderr", func() {
			Eventually(session.Err).Should(gbytes.Say("please provide valid args json"))
		})
	})

	Context("when no ticker is provided in the args json", func() {
		BeforeEach(func() {
			bobCmd.Args = append(bobCmd.Args, "fetch", "{}")
		})

		It("exits with a code of 1 to signal no retry", func() {
			Eventually(session).Should(gexec.Exit(1))
		})

		It("prints a useful error to stderr", func() {
			Eventually(session.Err).Should(gbytes.Say("please provide a ticker in the args json"))
		})
	})

	Context("when given the fetch action and a ticker", func() {
		BeforeEach(func() {
			bobCmd.Args = append(bobCmd.Args, "fetch", "{ \"ticker\": \"PVTL\" }")
		})

		It("exits with a successful code", func() {
			Eventually(session).Should(gexec.Exit(0))
		})

		It("prints some fetch in dollars", func() {
			Eventually(session).Should(gbytes.Say(`\d+\.\d+`))
		})
	})

})
