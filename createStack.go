package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudformation"
	"io/ioutil"
	"os"
)

func main() {
	stackname := "s2"
	filename := "cf.yml"
	templateBody, _ := convertFileToString(filename)
	createStack(stackname, templateBody)
	waitStackCreateComplete(stackname)
	describeStacks(stackname)
}

func createStack(stackname, templateBody string) {
	svc := cloudformation.New(
		session.New(),
		&aws.Config{Region: aws.String(os.Args[3])})

	p1 := cloudformation.Parameter{
		ParameterKey:   aws.String("IP8080"),
		ParameterValue: aws.String(os.Args[1]),
	}
	p2 := cloudformation.Parameter{
		ParameterKey:   aws.String("IP22"),
		ParameterValue: aws.String(os.Args[2]),
	}
	p3 := cloudformation.Parameter{
		ParameterKey:   aws.String("region"),
		ParameterValue: aws.String(os.Args[3]),
	}
	p4 := cloudformation.Parameter{
		ParameterKey:   aws.String("gszipbucketobject"),
		ParameterValue: aws.String(os.Args[4]),
	}
	p5 := cloudformation.Parameter{
		ParameterKey:   aws.String("ami"),
		ParameterValue: aws.String(os.Args[5]),
	}

	input := cloudformation.CreateStackInput{
		Parameters: []*cloudformation.Parameter{
			&p1,
			&p2,
			&p3,
			&p4,
			&p5,
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

func waitStackCreateComplete(stackname string) {
	svc := cloudformation.New(
		session.New(),
		&aws.Config{Region: aws.String(os.Args[3])})

	input := cloudformation.DescribeStacksInput{
		StackName: aws.String(stackname),
	}

	err := svc.WaitUntilStackCreateComplete(&input)
	_ = err
}

func describeStacks(stackname string) {
	svc := cloudformation.New(
		session.New(),
		&aws.Config{Region: aws.String(os.Args[3])})

	input := cloudformation.DescribeStacksInput{
		StackName: aws.String(stackname),
	}

	output, err := svc.DescribeStacks(&input)
	if err != nil {
		fmt.Println(err)
	}

	for _, v := range output.Stacks {
		for _, w := range v.Outputs {
			//fmt.Printf("@88 %s, %s\n", *w.OutputKey, *w.OutputValue)
			if *w.OutputKey == "InstanceID" {
				fmt.Printf("%s:8080/geoserver/web", *w.OutputValue)
			}
		}
	}
}

func convertFileToString(filename string) (string, error) {

	dat, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Can't read dat file", err)
		return "", err
	}
	return string(dat), nil
}
