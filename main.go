package main
 
import (
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-lambda-go/events"
	"github.com/google/uuid"
	"./structs"
	"context"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

func get_uuid() string {
	uuid, err := uuid.NewUUID()
	if err != nil {
		fmt.Println(err)
	}
	return uuid.String()
}

func Initiate_session() *dynamodb.DynamoDB {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
	SharedConfigState: session.SharedConfigEnable,
	}))
	svc := dynamodb.New(sess)

	return svc
}

func Upload_db(uuid string, patient structs.Patient) error {
	patient.Uuid = uuid

	svc := Initiate_session()

	av, err := dynamodbattribute.MarshalMap(patient)
	if err != nil {
		fmt.Println(err)
	}

	tableName := "patients"

	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(tableName),
	}

	_, err = svc.PutItem(input)
	if err != nil {
		fmt.Println(err)
	}

	return err
}

func HandleLambdaEvent(ctx context.Context, patient structs.Patient) (events.APIGatewayProxyResponse, error) {
	uuid := get_uuid()
	err := Upload_db(uuid, patient)

	resp := events.APIGatewayProxyResponse{Headers: make(map[string]string)}
	resp.Headers["Access-Control-Allow-Origin"] = "*"
	resp.Headers["Access-Control-Allow-Headers"] = "Content-Type"
	resp.Headers["content-type"] = "application/json"
	
	if err != nil {
		resp.StatusCode = 400
		resp.Body = "Something went wrong"
	} else {
		resp.StatusCode = 200
	}



	return resp, nil
}
 
func main() {
	lambda.Start(HandleLambdaEvent)
}