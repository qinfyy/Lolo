package config

type Resources struct {
	ResourcePath string `json:"ResourcePath"`
	DataPath     string `json:"DataPath"`
}

var defaultResources = &Resources{
	ResourcePath: "./Resource",
	DataPath:     "./data",
}

func GetResources() *Resources {
	return GetConfig().Resources
}

func (x *Resources) GetResourcePath() string {
	return x.ResourcePath
}

func (x *Resources) GetDataPath() string {
	return x.DataPath
}
