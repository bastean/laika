package sniffEmails_test

import (
	"testing"

	"github.com/bastean/laika/pkg/context/email/application/sniffEmails"
	aggregateMother "github.com/bastean/laika/test/pkg/context/shared/domain/aggregate"
	"github.com/stretchr/testify/suite"
)

type SniffEmailsTestSuite struct {
	suite.Suite
	sut *sniffEmails.SniffEmails
}

func (suite *SniffEmailsTestSuite) SetupTest() {
	suite.sut = sniffEmails.NewSniffEmails()
}

func (suite *SniffEmailsTestSuite) TestSniffEmails() {
	expected := aggregateMother.Create()

	suite.sut.Run(expected)

	suite.EqualValues(expected, expected)
}

func TestSniffEmailsSuite(t *testing.T) {
	suite.Run(t, new(SniffEmailsTestSuite))
}
