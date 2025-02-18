//go:build integration

// nolint:errcheck,stylecheck // because tests
package myhub

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/zikwall/myhub/internal/infrastructure/repositories/postgres/test_container"
)

// DatabaseTestSuite is the test suite.
type DatabaseTestSuite struct {
	suite.Suite

	container *test_container.TestDatabase
}

// SetupSuite is called once before the test suite runs.
func (s *DatabaseTestSuite) SetupSuite() {
	s.container = test_container.NewTestDatabase(s.T())
}

// TearDownSuite is called once after the test suite runs.
func (s *DatabaseTestSuite) TearDownSuite() {
	err := s.container.Drop()
	require.NoError(s.T(), err)
}

// TestSuite runs the test suite.
// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestDatabaseTestSuite(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test in short mode.")
	}

	suite.Run(t, new(DatabaseTestSuite))
}
