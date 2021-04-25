package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type RuleRecord struct {
	Id      int   	`gorm:"column:id"`
	Name	string		`gorm:"column:rule_name",`
	Fn      string		`gorm:"column:rule_fn",`
	Interval  int		`gorm:"column:rule_interval",`
	Alert     string	`gorm:"column:rule_alert",`
	Expr 	 string		`gorm:"column:rule_expr",`
	For      string		`gorm:"column:rule_for",`
	Note 	 string		`gorm:"column:note,omitempty",`
	State    int             `gorm:"column:state"`
	CreatedAt     *time.Time `gorm:"column:created_at"`
	UpdatedAt     *time.Time `gorm:"column:updated_at"`
}


func (*RuleRecord) TableName() string {
	return "t_rules"
}

func (r *RuleRecord) QueryAll(db *gorm.DB, onlyValid bool) ([]RuleRecord, error){
	records := []RuleRecord{}
	if onlyValid {
		db.Where("state = 1" ).Find(&records)
	}else{
		db.Find(&records)
	}

	return records, nil
}

func (r *RuleRecord) GetOneByGroupKey(db *gorm.DB, record RuleRecord) *RuleRecord {
	db.Where("rule_name = ? AND rule_fn = ?", record.Name, record.Fn).First(&record)
	return &record
}

func (r *RuleRecord) UpdatesByIds(db *gorm.DB, Ids []int, toUpdateMap map[string]interface{} ) {
	db.Where("id IN (?)}", Ids).Updates(toUpdateMap)
}

func (r *RuleRecord) Add(db *gorm.DB, record RuleRecord) int {
	db.Create(record)
	return record.Id
}

func (r *RuleRecord) Update(db *gorm.DB, record RuleRecord) int {
	db.Save(record)
	return record.Id
}

func (r *RuleRecord) Delete(db *gorm.DB, ruleId int) error {
	return db.Where(&RuleRecord{Id: ruleId}).Delete(RuleRecord{}).Error
}
