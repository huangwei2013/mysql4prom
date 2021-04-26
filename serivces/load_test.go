package serivces

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"io/ioutil"
	"log"
	"mysql4prom/config"
	"net/http"
	"testing"
)

func TestParseFromByte(t *testing.T){

	fileContent := `
groups:
   - name: kafka-consumergroup_lag
     rules:
     - alert: kafka 消费延迟
       expr: kafka_consumergroup_lag{consumergroup!="arch-hot-cold-split-sync-group",consumergroup!="nwd-rpc"}  > 1000
       for: 5m
       labels:
         status: warning
       annotations:
         summary: "{{$labels.job}}"
         description: "consumergroup: {{ $labels.consumergroup }} topic: {{$labels.topic}} 消费延迟 > 1000(当前值：{{ $value}})"

   - name: kafka-topic_partition_current_offset
     rules:
     - alert: kafka_topic监控
       expr: sum(changes(kafka_topic_partition_current_offset{topic!="order_sendOnlineCancelOrder_SaleOrder_w6r6",topic!="order_sendOnlinePayOrder_SaleOrder_w6r6",topic!="order_sendOnlinePreOrder_SaleOrder_w6r6",topic!="order_sendOnlineRefundOrder_SaleOrder_w6r6",topic!="order_sendPosRefundOrder_SaleOrder_w6r6",topic!="order_sendPosSaleOrder_SaleOrder_w6r6",topic!="__consumer_offsets",topic!="unUploadOrder",topic!="erp_purch_storeSku_w6r6",topic!="base_download",topic!="base_storesku_storageSync_w6r6",topic!="dh",topic!="order_dealStock",topic!="finance_supplier_w6r6",topic!="promotion.clearPromotionCache" }[30m]) )by (topic,job,instance) < 1
       #expr: sum(changes(kafka_topic_partition_current_offset{topic!="__consumer_offsets" }[30m]) )by (topic) < 1
       for: 5m
       labels:
         status: warning
       annotations:
         summary: "{{$labels.job}}"
         description: "topic: {{$labels.topic}} 30分钟未产生数据"
`
	updater := Updater{}
	db, err := gorm.Open("mysql", config.DBUrl)
	if err!= nil{
		log.Fatalf("Open mysql database error: %s\n", err)
		return
	}
	defer db.Close()

	updater.Parse(db, "testyaml", []byte(fileContent))

}



func TestParseFromHttpFile(t *testing.T){

	url := "https://raw.githubusercontent.com/xjjdog/prometheus-cnf-pro/master/prometheus/kafka_monit.yml"
	url = "https://raw.githubusercontent.com/xjjdog/prometheus-cnf-pro/master/prometheus/eureka_monit.yml"
	url = "https://raw.githubusercontent.com/xjjdog/prometheus-cnf-pro/master/prometheus/sys_monit.yml"

	resp,err := http.Get(url)
	if err != nil {
		log.Fatalf("Open http file error: %s\n", err)
		return
	}

	defer resp.Body.Close()
	fileContent,err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Read http file error: %s\n", err)
		return
	}

	updater := Updater{}
	db, err := gorm.Open("mysql", config.DBUrl)
	if err!= nil{
		log.Fatalf("Open mysql database error: %s\n", err)
		return
	}
	defer db.Close()
	updater.Parse(db, "testyaml", []byte(fileContent))
}