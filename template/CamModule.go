package template

const (
	CamModuleTypeApplication = "application"
	CamModuleTypeLibrary     = "library"
	CamModuleTypeGrpc        = "grpc"
)

type CamModule struct {
	Type string `json:"type"`
}
