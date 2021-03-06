package steps

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestInstallClusterPrerequisites(t *testing.T) {

	Convey("InstallClusterPrerequisites function", t, func() {

		Convey("In case no error occurs, should", func() {

			kymaTestSteps, testInst, _, mockCommandExecutor := getTestSetup()

			Convey("call UpdateInstallationStatus once and call RunCommand returning no error", func() {

				err := kymaTestSteps.InstallClusterPrerequisites(testInst)

				So(mockCommandExecutor.TimesMockCommandExecutorCalled, ShouldEqual, 1)
				So(err, ShouldBeNil)
			})
		})

		Convey("In case an error occurs, should", func() {

			kymaTestSteps, testInst, _, mockFailingCommandExecutor := getFailingTestSetup()

			Convey("call UpdateInstallationStatus once, call RunCommand, call UpdateInstallationStatus again and return the error", func() {

				err := kymaTestSteps.InstallClusterPrerequisites(testInst)

				So(mockFailingCommandExecutor.MockFailingCommandExecutorCalled, ShouldBeTrue)
				So(err, ShouldNotBeNil)
			})
		})
	})
}
