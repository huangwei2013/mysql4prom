package serivces

import (
	"github.com/jinzhu/gorm"
	"mysql4prom/models"
)

//   适配 Promethues 结构的规则数据
type Rule4Prom struct {
	Name	string
	Fn      string
	Interval  int
	Alert     string
	Expr 	 string
	For      string
	Labels   map[string]string
	Annotations map[string]string
}


func (r *Rule4Prom) GetAllRules(db *gorm.DB, onlyValid bool) ([]Rule4Prom, error){
	rules := []Rule4Prom{}

	var model *models.RuleRecord
	ruleRecords, errRule := model.QueryAll(db, onlyValid)
	if errRule != nil{
		return nil, errRule
	}

	var modelLabel *models.RuleLabelRecord
	var modelAnnotation *models.RuleAnnotationRecord

	for _, r := range ruleRecords {

		ruleLableRecords, errRuleLabel := modelLabel.QueryByRuleId(db, onlyValid, r.Id)
		if errRuleLabel != nil{
			continue
		}
		ruleAnnotationRecords, errRuleAnnotation := modelAnnotation.QueryByRuleId(db, onlyValid, r.Id)
		if errRuleAnnotation != nil{
			continue
		}

		rules = append(rules,  Rule4Prom{
			Name: r.Name,
			Fn: r.Fn,
			Interval: r.Interval,
			Alert: r.Alert,
			Expr: r.Expr,
			For: r.For,
			Labels: modelLabel.Records2Map(ruleLableRecords),
			Annotations: modelAnnotation.Records2Map(ruleAnnotationRecords),
		})

	}

	return rules, nil
}


func (r *Rule4Prom) AddRules(db *gorm.DB, rules []Rule4Prom) error {
	modelRule := models.RuleRecord{}
	modelRuleLabel := models.RuleLabelRecord{}
	modelRuleAnnotation := models.RuleAnnotationRecord{}

	for _, rule := range rules {
		ruleRecord := models.RuleRecord{
			Name: rule.Name,
			Fn: rule.Fn,
			Interval: rule.Interval,
			Alert: rule.Alert,
			Expr: rule.Expr,
			For: rule.For,
			State: 1,
		}
		ruleId := modelRule.Add(db, ruleRecord)

		if ruleId != 0 {
			for k, v := range rule.Labels{
				ruleLabelRecord := models.RuleLabelRecord{
					RuleId:    ruleId,
					Key:       k,
					Value:     v,
					State:     1,
				}
				modelRuleLabel.Add(db, ruleLabelRecord)
			}
			for k, v := range rule.Annotations{
				ruleAnnotationRecord := models.RuleAnnotationRecord{
					RuleId:    ruleId,
					Key:       k,
					Value:     v,
					State:     1,
				}
				modelRuleAnnotation.Add(db, ruleAnnotationRecord)
			}
		}
	}

	return nil
}




func (r *Rule4Prom) UpdateRules(db *gorm.DB, rules []Rule4Prom) error {
	modelRule := models.RuleRecord{}
	modelRuleLabel := models.RuleLabelRecord{}
	modelRuleAnnotation := models.RuleAnnotationRecord{}

	for _, rule := range rules {
		ruleRecord := models.RuleRecord{
			Name: rule.Name,
			Fn: rule.Fn,
			Interval: rule.Interval,
			Alert: rule.Alert,
			Expr: rule.Expr,
			For: rule.For,
			State: 1,
		}
		ruleId := modelRule.Update(db, ruleRecord)

		if ruleId != 0 {
			for k, v := range rule.Labels{
				ruleLabelRecord := models.RuleLabelRecord{
					RuleId:    ruleId,
					Key:       k,
					Value:     v,
					State:     1,
				}
				modelRuleLabel.Add(db, ruleLabelRecord)
			}
			for k, v := range rule.Annotations{
				ruleAnnotationRecord := models.RuleAnnotationRecord{
					RuleId:    ruleId,
					Key:       k,
					Value:     v,
					State:     1,
				}
				modelRuleAnnotation.Add(db, ruleAnnotationRecord)
			}
		}
	}

	return nil
}