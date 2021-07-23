package database

type TrendDTO struct {
	BasicDTO
	Trend Trend
}

func (t *TrendDTO) Insert() {
	db := t.Get(DBTrend)
	if db.Model(&t.Trend).Where("title = ?", t.Trend.Title).Updates(&t.Trend).RowsAffected == 0 {
		db.Create(&t.Trend)
	}
}

func (t *TrendDTO) Update() error {
	db := t.Get(DBTrend)
	if err := db.Model(&t.Trend).Updates(t.Trend).Error; err != nil {
		return err
	}
	return nil
}

func (t *TrendDTO) FetchByTitle() (*Trend, error) {
	db := t.Get(DBTrend)
	row := Trend{}
	if err := db.Where("title = ?", t.Trend.Title).Find(&row).Error; err != nil {
		return nil, err
	}
	return &row, nil
}
