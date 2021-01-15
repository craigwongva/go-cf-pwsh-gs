package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudformation"
	"io/ioutil"
)

func main() {
	stackname := "s1"
	filename := "cf.yml"
	templateBody, _ := convertFileToString(filename)
	createStack(stackname, templateBody)
}

func createStack(stackname, templateBody string) {
	svc := cloudformation.New(
		session.New(),
		&aws.Config{Region: aws.String("us-west-2")})

	p1 := cloudformation.Parameter{
		ParameterKey:   aws.String("IP8080"),
		ParameterValue: aws.String(os.Args[1]),
	}
	p2 := cloudformation.Parameter{
		ParameterKey:   aws.String("IP22"),
		ParameterValue: aws.String(os.Args[2]),
	}
	p3 := cloudformation.Parameter{
		ParameterKey:   aws.String("githubpassword"),
		ParameterValue: aws.String(os.Args[3]),
	}
	input := cloudformation.CreateStackInput{
		Parameters: []*cloudformation.Parameter{
			&p1,
			&p2,
			&p3,
		},
		StackName:    aws.String(stackname),
		TemplateBody: aws.String(templateBody),
	}
	output, err := svc.CreateStack(&input)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%s\n", *output.StackId)
}

func convertFileToString(filename string) (string, error) {

	dat, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Can't read dat file", err)
		return "", err
	}
	return string(dat), nil
}
