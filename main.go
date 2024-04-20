package main

import (
	"context"
	"fmt"
	"log"

	analyticsadmin "google.golang.org/api/analyticsadmin/v1beta"
	"google.golang.org/api/option"
)

func main() {
	ctx := context.Background()


	service,err := analyticsadmin.NewService(ctx,option.WithCredentialsFile("./eloquent-env-247113-a0d11e23fccc.json"))
	if err != nil {
		log.Fatal("new service err",err)
		return 
	}
	accountsService := analyticsadmin.NewAccountsService(service)
	accountsReply, err := accountsService.List().Do(); 
	if err != nil {
		log.Fatal("list account err",err)
		return
	}
	for _,acc := range accountsReply.Accounts {
		fmt.Println(acc.Name)
	}
	
	propertiesService := analyticsadmin.NewPropertiesService(service)
	propertiesReply,err := propertiesService.List().Filter(fmt.Sprintf("parent:%s",accountsReply.Accounts[0].Name)).Do()
	if err != nil {
		log.Fatal("list properties err",err)
		return
	}
	for _,pro := range propertiesReply.Properties {
		fmt.Println(pro)
	}
	propertie,err := propertiesService.Create(&analyticsadmin.GoogleAnalyticsAdminV1betaProperty{
		Account:         accountsReply.Accounts[0].Name,
		CurrencyCode:     "CNY",
		DisplayName:      "overstarrytest",
		// 行业示
		IndustryCategory: "ONLINE_COMMUNITIES",
		Parent:           accountsReply.Accounts[0].Name,
		PropertyType:     "PROPERTY_TYPE_ORDINARY",
		TimeZone:         "Asia/Shanghai",
	}).Do()
	if err != nil {
		log.Fatal("create  propertie err",err)
		return
	
	}
	propertiesDataStreamsService  := analyticsadmin.NewPropertiesDataStreamsService(service)
	res,err := propertiesDataStreamsService.Create(propertie.Name,&analyticsadmin.GoogleAnalyticsAdminV1betaDataStream{
		DisplayName:          "测试链接",
		Type:                 "WEB_DATA_STREAM",
		WebStreamData:        &analyticsadmin.GoogleAnalyticsAdminV1betaDataStreamWebStreamData{
			DefaultUri:      "https://www.overstarry.vip",
			MeasurementId:   "",
		},
	}).Do()
	fmt.Printf("%v",res.WebStreamData.MeasurementId)
}
