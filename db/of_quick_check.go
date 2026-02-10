package db

type OFQuickCheck struct {
	UID             string `gorm:"primarykey;unique;index"`
	GateToken       string
	LastPackageName string
}

func OrCreateOFQuickCheck(uid string) (*OFQuickCheck, error) {
	q := &OFQuickCheck{
		UID: uid,
	}
	err := db.Where("uid = ?", uid).FirstOrCreate(q).Error
	if err != nil {
		return nil, err
	}
	return q, err
}

func UpOFQuickCheck(q *OFQuickCheck) error {
	return db.Save(q).Error
}
