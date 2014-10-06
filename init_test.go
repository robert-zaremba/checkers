package checkers

import (
	. "gopkg.in/check.v1"
	"testing"
)

func Test(t *testing.T) {
	Suite(&S{})
	Suite(&ContainerSuite{})
	Suite(&FileSuite{})
	Suite(&SamePathLinuxSuite{})
	Suite(&SamePathWindowsSuite{})
	TestingT(t)
}
