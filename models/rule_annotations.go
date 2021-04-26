package models

import (
	"github.com/jinzhu/gorm"
	"time"
)


type RuleAnnotationRecord struct {
	Id      int   	`gorm:"column:id"`
	RuleId	int		`gorm:"column:rule_id",`
	Key      string		`gorm:"column:annotation_key",`
	Value      string		`gorm:"column:annotation_value",`
	State    int             `gorm:"column:state"`
	CreatedAt     time.Time `gorm:"column:created_at"`
	UpdatedAt     time.Time `gorm:"column:updated_at"`
}


func (*RuleAnnotationRecord) TableName() string {
	return "t_rule_annotations"
}

func (r *RuleAnnotationRecord) QueryByRuleId(db *gorm.DB, onlyValid bool, ruleId int) (*[]RuleAnnotationRecord, error){
	var records []RuleAnnotationRecord
	if onlyValid {
		db.Where(&RuleAnnotationRecord{State :1, RuleId: ruleId}).Find(&records)
	}else{
		db.Where(&RuleAnnotationRecord{RuleId: ruleId}).Find(&records)
	}

	return &records, nil
}


func (r *RuleAnnotationRecord) Records2Map(records *[]RuleAnnotationRecord)  map[string]string {
	maps := make(map[string]string)
	for _,v := range *records {
		maps[v.Key] = v.Value
	}
	return maps
}

func (r *RuleAnnotationRecord) UpdatesByIds(db *gorm.DB, Ids []int, toUpdateMap map[string]interface{} ) {
	db.Where("id IN (?)}", Ids).Updates(toUpdateMap)
}

func (r *RuleAnnotationRecord) Add(db *gorm.DB, record RuleAnnotationRecord) int {
	db.Create(&record)
	db.NewRecord(record)
	return record.Id
}

func (r *RuleAnnotationRecord) Update(db *gorm.DB, record RuleAnnotationRecord) {
	db.Save(record)
}

func (r *RuleAnnotationRecord) Delete(db *gorm.DB, ruleId int) error {
	return db.Where(&RuleAnnotationRecord{Id: ruleId}).Delete(RuleAnnotationRecord{}).Error
}