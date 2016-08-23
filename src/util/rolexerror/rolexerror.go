package rolexerror

const (
	//Default error code
	CodeUndefined = "503-10001"

	//Network error code

	//Node error code
	CodeErrorUpdateNodeMethod = "503-11302"
	CodeErrorNodeRole         = "503-11303"
	CodeErrorNodeAvailability = "503-11304"
	CodeGetNodeInfoError      = "503-11305"
)

type ContainerStatsStopError struct {
	ID  string
	Err error
}

func (e *ContainerStatsStopError) Error() string {
	if e.Err != nil {
		return e.Err.Error()
	}

	return "normal stop" + e.ID
}

type NodeConnError struct {
	ID       string
	Endpoint string
	Err      error
}

func (e *NodeConnError) Error() string {
	if e.Err != nil {
		return e.Err.Error()
	}

	return e.ID + " : " + e.Endpoint + " conn error"
}