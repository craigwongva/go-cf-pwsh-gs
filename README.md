### Usage

#### From the local host
```
cd ~/environment/go-cf-pwsh-gs
go mod init go-cf-pwsh-gs
go install ./...
#~/go/bin/go-cf-pwsh-gs <IP8080> <IP22> <region> <gszipbucketobject> <ami> <keypair> <stackname> <profile> <filename>
~/go/bin/go-cf-pwsh-gs "68.100.238.208" "100.101.102.103" "us-west-2" \
 "deleteme1001/Deploy-Gs.zip" "ami-0a36eb8fadc976275" oregonkeypair "go_cf_pwsh_gs" \
 "s2" "guitar" "cf.yml"

tail -400 /var/log/cloud-init-output.log
```

#### From a remote host
```
cd ~/environment/go-cf-pwsh-gs
aws cloudformation \
--endpoint-url https://cloudformation.us-gov-west-1.amazonaws.com \
--profile guitar create-stack --stack-name s1 --region us-gov-west-1 \
--template-body file://cf.yml \
--parameters ParameterKey=IP8080,ParameterValue=68.100.238.208 \
ParameterKey=IP22,ParameterValue=34.212.135.227 \
ParameterKey=region,ParameterValue=us-gov-west-1 \
ParameterKey=gszipbucketobject,ParameterValue=ogc-feeds/Deploy-Gs.zip \
ParameterKey=ami,ParameterValue=ami-2bad964a \
ParameterKey=keypair,ParameterValue=Stratus \
ParameterKey=instancerole,ParameterValue=go_cf_pwsh_gs
```

where ~/.aws/credentials looks like this:
```
[guitar]
aws_access_key_id = AKIAredacted 
aws_secret_access_key = redacted
```