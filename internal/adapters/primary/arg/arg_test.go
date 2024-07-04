package argparser

import (
	"os"
	"testing"

	"github.com/stretchr/testify/suite"
)

type testSuite struct {
	suite.Suite
}

func TestRunner(t *testing.T) {
	suite.Run(t, new(testSuite))
}

func (t *testSuite) SetupTest() {
	// TODO
}

func (t *testSuite) TearDownTest() {
	// TODO
}

func (t *testSuite) TestFileExpected() {
	validFileName := "test.jpeg"
	os.Args = []string{"app", validFileName}
	arg := New()
	res, err := arg.FileName()
	t.Nil(err)
	t.Equal(validFileName, res.FileName())
}

func (t *testSuite) TestMissingArgumentReturnsError() {
	os.Args = []string{"app"}
	arg := New()
	_, err := arg.FileName()
	t.NotNil(err)
	t.ErrorIs(err, errInvalidArgumentCount)
}
