package global

import (
	"regexp"
	"time"

	"code.cloudfoundry.org/cli/integration/helpers"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gexec"

	"testing"
)

const (
	CFEventuallyTimeout   = 300 * time.Second
	CFConsistentlyTimeout = 500 * time.Millisecond
)

var (
	// Per Test Level
	homeDir       string
	ReadOnlyOrg   string
	ReadOnlySpace string
)

func TestGlobal(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Global Suite")
}

var _ = SynchronizedBeforeSuite(func() []byte {
	// Ginkgo Globals
	SetDefaultEventuallyTimeout(CFEventuallyTimeout)
	SetDefaultConsistentlyDuration(CFConsistentlyTimeout)

	// Setup common environment variables
	helpers.TurnOffColors()

	helpers.SetupSynchronizedSuite(func() {
		helpers.EnableFeatureFlag("diego_docker")
		helpers.EnableFeatureFlag("service_instance_sharing")
	})

	ReadOnlyOrg, ReadOnlySpace = helpers.SetupReadOnlyOrgAndSpace()

	return nil
}, func(_ []byte) {
	if GinkgoParallelNode() != 1 {
		Fail("Test suite cannot run in parallel")
	}
})

var _ = BeforeEach(func() {
	homeDir = helpers.SetHomeDir()
	helpers.SetAPI()
})

var _ = AfterEach(func() {
	GinkgoWriter.Write([]byte("==============================Global After Each=============================="))
	helpers.DestroyHomeDir(homeDir)
})

var foundDefaultDomain string

func defaultSharedDomain() string {
	// TODO: Move this into helpers when other packages need it, figure out how
	// to cache cuz this is a wacky call otherwise
	if foundDefaultDomain == "" {
		session := helpers.CF("domains")
		Eventually(session).Should(Exit(0))

		regex, err := regexp.Compile(`(.+?)\s+shared`)
		Expect(err).ToNot(HaveOccurred())

		matches := regex.FindStringSubmatch(string(session.Out.Contents()))
		Expect(matches).To(HaveLen(2))

		foundDefaultDomain = matches[1]
	}
	return foundDefaultDomain
}

func setupCF(org string, space string) {
	helpers.LoginCF()
	if org != ReadOnlyOrg && space != ReadOnlySpace {
		helpers.CreateOrgAndSpace(org, space)
	}
	helpers.TargetOrgAndSpace(org, space)
}
