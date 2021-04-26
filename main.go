package main

import (
	"context"
	"fmt"
	"log"
	"mysql4prom/config"
	"time"

	_ "mysql4prom/bak"
	"mysql4prom/serivces"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)


func InitDBByGorm(dbUrl string)  (*gorm.DB, error){
	db, err := gorm.Open("mysql", dbUrl)
	if err!= nil{
		log.Fatalf("Open mysql database error: %s\n", err)
		return nil, err
	}
	return db, err
}


func RunService(ctx context.Context){

	db, err := InitDBByGorm(config.DBUrl)
	if err != nil || db == nil{
		fmt.Println(err)
		return
	}

	defer db.Close()
	var serivces *serivces.Rule4Prom

	for {
		select {
		case <-ctx.Done():
			fmt.Println("WriteRedis Done.")
			return
		default:
			fmt.Println("WriteRedis running")

			rules, errRule := serivces.GetAllRules(db, true)
			if errRule != nil {
				fmt.Println(errRule)
				ctx.Done()
				return
			}else{
				fmt.Println(rules)
			}
			time.Sleep(30 * time.Second)
		}
	}
}


func main(){
	fmt.Println("Make a mysql/maria for Promethues")

	ctx, cancel := context.WithCancel(context.Background())
	go RunService(ctx)

	ticker := time.NewTicker(60 * time.Second)
	defer ticker.Stop()
	for  {
		select {
		case <-ctx.Done():
			cancel()
			log.Printf("quit running after 3 seconds")
			time.Sleep(time.Second * 3)
			return
		case <- ticker.C:
			log.Println("心跳一下，继续运行")
		}
	}
}
