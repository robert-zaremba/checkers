package checkers

import (
	"testing"

	. "gopkg.in/check.v1"
)

func Test(t *testing.T) {
	Suite(&S{})
	Suite(&Numeric{})
	Suite(&Time{})
	Suite(&ContainerSuite{})
	Suite(&FileSuite{})
	Suite(&SamePathLinuxSuite{})
	Suite(&SamePathWindowsSuite{})
	TestingT(t)
}
