package objects

// Backend defines interface for all types of backend
type Backend interface {
	Init(interface{}) error
	GetGlobalConfig(*GlobalConfig) error
	SetGlobalConfig(*GlobalConfig) error
	//	GetLayersId()
	//	GetLayerById()
	//	GetTile()
}
