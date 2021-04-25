package models

import (
	"github.com/jinzhu/gorm"
	"time"
)


type RuleLabelRecord struct {
	Id      int   	`gorm:"column:id"`
	RuleId	int		`gorm:"column:rule_id",`
	Key      string		`gorm:"column:label_key",`
	Value      string		`gorm:"column:label_value",`
	State    int             `gorm:"column:state"`
	CreatedAt     *time.Time `gorm:"column:created_at"`
	UpdatedAt     *time.Time `gorm:"column:updated_at"`
}


func (*RuleLabelRecord) TableName() string {
	return "t_rule_labels"
}

func (r *RuleLabelRecord) QueryByRuleId(db *gorm.DB, onlyValid bool, ruleId int) (*[]RuleLabelRecord, error){
	records := []RuleLabelRecord{}
	if onlyValid {
		db.Where(&RuleLabelRecord{State :1, RuleId: ruleId}).Find(&records)
	}else{
		db.Where(&RuleLabelRecord{RuleId: ruleId}).Find(&records)
	}

	return &records, nil
}

func (r *RuleLabelRecord) Records2Map(records *[]RuleLabelRecord)  map[string]string {
	maps := make(map[string]string)
	for _,v := range *records {
		maps[v.Key] = v.Value
	}
	return maps
}

func (r *RuleLabelRecord) UpdatesByIds(db *gorm.DB, Ids []int, toUpdateMap map[string]interface{} ) {
	db.Where("id IN (?)}", Ids).Updates(toUpdateMap)
}

func (r *RuleLabelRecord) Add(db *gorm.DB, record RuleLabelRecord)  int {
	db.Create(record)
	return record.Id
}

func (r *RuleLabelRecord) Update(db *gorm.DB, record RuleLabelRecord) {
	db.Save(record)
}

func (r *RuleLabelRecord) Delete(db *gorm.DB, ruleId int) error {
	return db.Where(&RuleLabelRecord{Id: ruleId}).Delete(RuleLabelRecord{}).Error
}