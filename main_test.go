package main

import(
	"testing"
	"./structs"

	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/aws"
	"fmt"
)

func TestHandleLambdaEvent(t *testing.T) {
	patient := structs.Patient{Uuid: "test",
		Illness: "testIllness",
		Pain_level: 0,
		Hospital: "testHospital",
		Name: "testName",
	}

	test, _ := HandleLambdaEvent(nil, patient)
	if test.StatusCode != 200 {
		t.Error("Wrong status code (not 200)")
	}
}

func TestUpload_db(t *testing.T) {
	uuid := "2"
	patient := structs.Patient{Uuid: "test",
		Illness: "testIllness",
		Pain_level: 0,
		Hospital: "testHospital",
		Name: "testName",
	}
	test := Upload_db(uuid, patient)

	if test != nil {
		t.Error(test)
	}

	svc := Initiate_session()
	tableName := "patients"

	result, err := svc.GetItem(&dynamodb.GetItemInput{
	    TableName: aws.String(tableName),
	    Key: map[string]*dynamodb.AttributeValue{
	        "uuid": {
	            S: aws.String(uuid),
	        },
	    },
	})
	if *result.Item["name"].S!= "testName" {
		t.Error("Wrong name")
	}
	if *result.Item["illness"].S!= "testIllness" {
		t.Error("Wrong ill")
	}
	if *result.Item["hospital"].S!= "testHospital" {
		t.Error("Wrong hos")
	}
	if *result.Item["painLevel"].N!= "0" {
		t.Error("Wrong pain")
	}
	if err != nil {
	    fmt.Println(err.Error())
		return
	}
}