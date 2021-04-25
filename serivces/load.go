package serivces

import (
	"fmt"
	"github.com/prometheus/prometheus/pkg/rulefmt"
)


//未完成

type Updater struct{}


func (u *Updater) Parse(content []byte) error {

	rgs, errs := rulefmt.Parse(content)
	if errs != nil{
		for i := range errs {
			if errs[i] != nil{
				return errs[i]
			}
		}
	}

	for _, rg := range rgs.Groups {
		fmt.Println(rg.Name)

		//TODO:
	}

	return nil
}

