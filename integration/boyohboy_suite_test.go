package integration_test

import (
	"os/exec"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"
)

var boyOhBoyPath string

var _ = BeforeSuite(func() {
	var err error
	boyOhBoyPath, err = gexec.Build("github.com/williammartin/boyohboy")
	Expect(err).NotTo(HaveOccurred())
})

var _ = AfterSuite(func() {
	gexec.CleanupBuildArtifacts()
})

var execBOB = func(args ...string) *gexec.Session {
	cmd := exec.Command(boyOhBoyPath, args...)
	session, err := gexec.Start(cmd, GinkgoWriter, GinkgoWriter)
	Expect(err).NotTo(HaveOccurred())
	return session
}

func TestIntegration(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "BoyOhBoy Integration Suite")
}
