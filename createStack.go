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

	IP8080 := os.Args[1]
	IP22 := os.Args[2]
	region := os.Args[3]
	gszipbucketobject := os.Args[4]
	ami := os.Args[5]
	keypair := os.Args[6]
	instancerole := os.Args[7]
	stackname := os.Args[8]
	profile := os.Args[9]
	filename := os.Args[10]

	templateBody, _ := convertFileToString(filename)

	createStack(
		IP8080, IP22, region, gszipbucketobject,
		ami, keypair, instancerole, stackname, templateBody, profile)
	waitStackCreateComplete(region, stackname, profile)
	describeStacks(region, stackname, profile)
}

func createStack(
	IP8080, IP22, region, gszipbucketobject,
	ami, keypair, instancerole, stackname, templateBody, profile string) {

	e := fmt.Sprintf("https://cloudformation.%s.amazonaws.com", region)
	svc := cloudformation.New(
		session.New(),
		&aws.Config{
			Region:      aws.String(region),
			Endpoint:    aws.String(e),
			Credentials: credentials.NewSharedCredentials("", profile),
		})

	p1 := cloudformation.Parameter{
		ParameterKey:   aws.String("IP8080"),
		ParameterValue: aws.String(IP8080),
	}
	p2 := cloudformation.Parameter{
		ParameterKey:   aws.String("IP22"),
		ParameterValue: aws.String(IP22),
	}
	p3 := cloudformation.Parameter{
		ParameterKey:   aws.String("region"),
		ParameterValue: aws.String(region),
	}
	p4 := cloudformation.Parameter{
		ParameterKey:   aws.String("gszipbucketobject"),
		ParameterValue: aws.String(gszipbucketobject),
	}
	p5 := cloudformation.Parameter{
		ParameterKey:   aws.String("ami"),
		ParameterValue: aws.String(ami),
	}
	p6 := cloudformation.Parameter{
		ParameterKey:   aws.String("keypair"),
		ParameterValue: aws.String(keypair),
	}
	p7 := cloudformation.Parameter{
		ParameterKey:   aws.String("instancerole"),
		ParameterValue: aws.String(instancerole),
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

	fmt.Printf("The stack is being created: %s\n", *output.StackId)
}

func waitStackCreateComplete(region, stackname, profile string) {

	e := fmt.Sprintf("https://cloudformation.%s.amazonaws.com", region)
	svc := cloudformation.New(
		session.New(),
		&aws.Config{
			Region:      aws.String(region),
			Endpoint:    aws.String(e),
			Credentials: credentials.NewSharedCredentials("", profile),
		})

	input := cloudformation.DescribeStacksInput{
		StackName: aws.String(stackname),
	}

	err := svc.WaitUntilStackCreateComplete(&input)
	_ = err
}

func describeStacks(region, stackname, profile string) {

	e := fmt.Sprintf("https://cloudformation.%s.amazonaws.com", region)
	svc := cloudformation.New(
		session.New(),
		&aws.Config{
			Region:      aws.String(region),
			Endpoint:    aws.String(e),
			Credentials: credentials.NewSharedCredentials("", profile),
		})

	input := cloudformation.DescribeStacksInput{
		StackName: aws.String(stackname),
	}

	output, err := svc.DescribeStacks(&input)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Stack %s has been created. I am sleeping 5 minutes while GeoServer installs...", stackname)
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
