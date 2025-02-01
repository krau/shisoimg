package dao

func CreateRule(prefix string, path string) error {
	return db.Create(&UrlRule{
		Prefix: prefix,
		Path:   path,
	}).Error
}

func DeleteRule(path string) error {
	return db.Where("path = ?", path).Delete(&UrlRule{}).Error
}

func GetRules() ([]UrlRule, error) {
	var rules []UrlRule
	err := db.Find(&rules).Error
	return rules, err
}
