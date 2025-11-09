package model

type ArchiveModel struct {
	ArchiveMap map[string]string
}

func (s *Player) GetArchive() *ArchiveModel {
	if s.Archive == nil {
		s.Archive = new(ArchiveModel)
	}
	return s.Archive
}

func (a *ArchiveModel) GetArchiveMap() map[string]string {
	if a.ArchiveMap == nil {
		a.ArchiveMap = make(map[string]string)
	}
	return a.ArchiveMap
}

func (a *ArchiveModel) SetArchiveMap(k, v string) {
	archiveMap := a.GetArchiveMap()
	archiveMap[k] = v
}

func (a *ArchiveModel) GetArchiveValue(k string) string {
	archiveMap := a.GetArchiveMap()
	v, ok := archiveMap[k]
	if !ok {
		return ""
	}
	return v
}
