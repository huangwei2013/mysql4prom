package serivces

import (
	"github.com/jinzhu/gorm"
	"github.com/prometheus/prometheus/pkg/rulefmt"
	"log"
	"mysql4prom/config"
	"mysql4prom/models"
)


//未完成

type Updater struct{}


func (u *Updater) Parse(db *gorm.DB, filename string, content []byte) error {

	rgs, errs := rulefmt.Parse(content)
	if errs != nil{
		for i := range errs {
			if errs[i] != nil{
				return errs[i]
			}
		}
	}

	var modelRule *models.RuleRecord
	for _, rg := range rgs.Groups {
		log.Println("rg: ",rg.Name)
		var ruleRecord4Query  *[]models.RuleRecord
		ruleRecord4Query = modelRule.GetsByGroupKey(db, rg.Name, filename, ruleRecord4Query)
		if ruleRecord4Query != nil {
			//update rules & labels & annotations
			u.UpdateRuleAndAll(db, ruleRecord4Query, &rg)
		}else{
			//insert rules & labels & annotations
			u.AddRuleAndAll(db, filename, &rg)
		}
	}

	return nil
}

func (u *Updater) UpdateRuleAndAll(db *gorm.DB, ruleRecord *[]models.RuleRecord, rg *rulefmt.RuleGroup) error {


	return nil
}

func (u *Updater) AddRuleAndAll(db *gorm.DB, filename string, rg *rulefmt.RuleGroup) error {
	interval := int(rg.Interval)
	if interval == 0 {
		interval = int(config.GlobalInterval.Seconds())
	}

	rulesRecord4Insert := make([]Rule4Prom, len(rg.Rules))
	for i, r := range rg.Rules{

		nr := Rule4Prom{
			Name:       rg.Name,
			Fn:       	filename,
			Alert:      r.Alert,
			Interval: 	interval,
			Expr:     	r.Expr,
			For:      	r.For.String(),
			Labels:		r.Labels,
			Annotations:	r.Annotations,
		}

		rulesRecord4Insert[i] = nr
	}

	if len(rulesRecord4Insert) > 0 {
		rule4Prom := Rule4Prom{}
		rule4Prom.AddRules(db, rulesRecord4Insert)
	}

	return nil
}
