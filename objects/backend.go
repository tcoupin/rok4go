package objects

type Backend interface {
	Init(interface{}) error
	GetGlobalConfig(*GlobalConfig) error
	SetGlobalConfig(*GlobalConfig) error
	//	GetLayersId()
	//	GetLayerById()
	//	GetTile()
}
