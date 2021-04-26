package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type RuleRecord struct {
	Id      int   	    `gorm:"primary_key" json:"id"`
	Name	string		`gorm:"column:rule_name",`
	Type	string		`gorm:"column:rule_type",`
	Fn      string		`gorm:"column:rule_fn",`
	Gn      string		`gorm:"column:rule_gn",`
	Interval  int		`gorm:"column:rule_interval",`
	Alert     string	`gorm:"column:rule_alert",`
	Expr 	 string		`gorm:"column:rule_expr",`
	For      string		`gorm:"column:rule_for",`
	Note 	 string		`gorm:"column:note",`
	State    int             `gorm:"column:state"`
	CreatedAt     time.Time `gorm:"column:created_at"`
	UpdatedAt     time.Time `gorm:"column:updated_at"`
}


func (*RuleRecord) TableName() string {
	return "t_rules"
}

func (r *RuleRecord) QueryAll(db *gorm.DB, onlyValid bool) ([]RuleRecord, error){
	var records []RuleRecord
	if onlyValid {
		db.Where("state = 1" ).Find(&records)
	}else{
		db.Find(&records)
	}

	return records, nil
}



func (r *RuleRecord) GetsByGroupKey(db *gorm.DB, ruleName string, ruleFn string, records *[]RuleRecord) *[]RuleRecord {
	db.Where("rule_name = ? AND rule_fn = ?", ruleName, ruleFn).Find(records)
	return records
}

func (r *RuleRecord) UpdatesByIds(db *gorm.DB, Ids []int, toUpdateMap map[string]interface{} ) {
	db.Where("id IN (?)}", Ids).Updates(toUpdateMap)
}

func (r *RuleRecord) Add(db *gorm.DB, record RuleRecord) int {
	db.Create(&record)
	db.NewRecord(record)
	return record.Id
}

func (r *RuleRecord) Update(db *gorm.DB, record RuleRecord) int {
	db.Save(&record)
	return record.Id
}

func (r *RuleRecord) Delete(db *gorm.DB, ruleId int) error {
	return db.Where(&RuleRecord{Id: ruleId}).Delete(RuleRecord{}).Error
}
