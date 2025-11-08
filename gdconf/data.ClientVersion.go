package gdconf

type ClientVersion struct {
	VersionMap map[string]string `json:"VersionMap"`
}

func (g *GameConfig) loadClientVersion() {
	g.Data.ClientVersion = new(ClientVersion)
	ReadJson(g.dataPath, "ClientVersion.json", &g.Data.ClientVersion)
}

func GetClientVersion(v string) string {
	return cc.Data.ClientVersion.VersionMap[v]
}
