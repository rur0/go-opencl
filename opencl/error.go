package opencl

// #cgo LDFLAGS: -lOpenCL
// #include <CL/cl.h>
import "C"
import (
	"errors"
	"fmt"
)

type clError int32

const (
	clSuccess             clError = clError(C.CL_SUCCESS)
	clDeviceNotFound              = clError(C.CL_DEVICE_NOT_FOUND)
	clInvalidValue                = clError(C.CL_INVALID_VALUE)
	clInvalidPlatform             = clError(C.CL_INVALID_PLATFORM)
	clBuildProgramFailure         = clError(C.CL_BUILD_PROGRAM_FAILURE)
	clOutOfResources              = clError(C.CL_OUT_OF_RESOURCES)
	clOutOfHostMemory             = clError(C.CL_OUT_OF_HOST_MEMORY)

	clPlatformNotFoundKHR = clError(C.CL_PLATFORM_NOT_FOUND_KHR)
)

var (
	DeviceNotFound      = errors.New("Device not found")
	InvalidValue        = errors.New("Invalid value")
	InvalidPlatform     = errors.New("Invalid platform")
	BuildProgramFailure = errors.New("Build program failure")
	OutOfResources      = errors.New("Out of resources")
	OutOfHostMemory     = errors.New("Out of host memory")

	PlatformNotFoundKHR = errors.New("No platforms found")

	UnexpectedType      = errors.New("Unexpected type")
	ErrorParsingVersion = errors.New("Error parsing OpenCL version")
	UnknownError        = errors.New("Unknown error")
)

var (
	errorMap = map[clError]error{
		clSuccess:             nil,
		clDeviceNotFound:      DeviceNotFound,
		clInvalidValue:        InvalidValue,
		clInvalidPlatform:     InvalidPlatform,
		clBuildProgramFailure: BuildProgramFailure,
		clOutOfResources:      OutOfResources,
		clOutOfHostMemory:     OutOfHostMemory,

		clPlatformNotFoundKHR: PlatformNotFoundKHR,
	}
)

func clErrorToError(clerr clError) error {
	err, ok := errorMap[clerr]
	if ok {
		return err
	}
	fmt.Println("Error", clerr)
	return UnknownError
}
