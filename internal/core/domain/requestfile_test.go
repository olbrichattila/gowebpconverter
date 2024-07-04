package request

import (
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

func (t *testSuite) TestFile() {
	validFileNames := []string{"test.jpeg", "test.jpg", "test.png", "test.gif", "test.svg", "test.webp", "test.tiff"}

	for _, validFileName := range validFileNames {
		res, err := NewFile(validFileName)
		t.Nil(err)
		t.Equal(validFileName, res.FileName())
	}
}
