package main

import (
	"context"
	"errors"
	"fmt"
	"os"

	events "github.com/aws/aws-lambda-go/events"
	lambda "github.com/aws/aws-lambda-go/lambda"
	awsgo "github.com/sotacode/gambituser/awsgo"
)

func main() {
	lambda.Start(EjecutoLambda)
}

func EjecutoLambda(ctx context.Context, event events.CognitoEventUserPoolsPostConfirmation) (events.CognitoEventUserPoolsPostConfirmation, error) {

	awsgo.InicializoAWS()

	if !ValidoParametros() {
		fmt.Println("Error en los parametros, debe enviar SecretName")
		err := errors.New("error en los parametros, debe enviar SecretName")
		return event, err
	}

	return event, nil
}

func ValidoParametros() bool {
	var traeParametro bool
	_, traeParametro = os.LookupEnv("SecretName")
	return traeParametro
}
