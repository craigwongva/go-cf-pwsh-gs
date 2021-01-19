package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudformation"
	"io/ioutil"
	"os"
	"time"
)

func main() {
	stackname := "s2"
	filename := "cf.yml"
	templateBody, _ := convertFileToString(filename)
	fmt.Println("createStack.go@18 I am calling createStack")
	createStack(stackname, templateBody)
	fmt.Println("createStack.go@18 I am calling waitStackCreateComplete")
	waitStackCreateComplete(stackname)
	fmt.Println("createStack.go@18 I am calling describeStacks")
	describeStacks(stackname)
}

func createStack(stackname, templateBody string) {

	e := fmt.Sprintf("https://cloudformation.%s.amazonaws.com", os.Args[3])
	svc := cloudformation.New(
		session.New(),
		&aws.Config{
			Region:      aws.String(os.Args[3]),
			Endpoint:    aws.String(e),
			Credentials: credentials.NewSharedCredentials("", "guitar"),
		})

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
	p6 := cloudformation.Parameter{
		ParameterKey:   aws.String("keypair"),
		ParameterValue: aws.String(os.Args[6]),
	}
	p7 := cloudformation.Parameter{
		ParameterKey:   aws.String("instancerole"),
		ParameterValue: aws.String(os.Args[7]),
	}
	input := cloudformation.CreateStackInput{
		Parameters: []*cloudformation.Parameter{
			&p1,
			&p2,
			&p3,
			&p4,
			&p5,
			&p6,
			&p7,
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
	fmt.Println("createStack.go@100 starting")
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

	fmt.Println("createStack.go@113 The stack has been created. I am sleeping 5 minutes while GeoServer installs...")
	time.Sleep(300 * time.Second)

	for _, v := range output.Stacks {
		for _, w := range v.Outputs {
			if *w.OutputKey == "InstanceID" {
				fmt.Printf("GeoServer is ready here: %s:8080/geoserver/web\n", *w.OutputValue)
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
