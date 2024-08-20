package gomeos

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	// Initialize MEOS before running any tests or examples
	MeosInitialize("UTC")

	// Run the tests (including examples)
	code := m.Run()

	// Finalize MEOS after all tests and examples have run
	MeosFinalize()

	// Exit with the code returned by m.Run()
	os.Exit(code)
}
