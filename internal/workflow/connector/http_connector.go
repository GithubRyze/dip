package connector

type HttpConnector struct {
	Endpoint  string
	Method    string
	Timeout   int
	Body      string
	Header    map[string]string
	TslConfig TslConfig
}

type TslConfig struct {
	InsecureSkipVerify bool
	CrtFilePath        string
	KeyFilePath        string
}

func (http *HttpConnector) Execute(ctx *DipContext) DipResult {

	return DipResult{}
}
