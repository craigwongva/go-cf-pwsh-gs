### Usage
```
cd ~/environment/go-cf-pwsh-gs
go mod init go-cf-pwsh-gs
go install ./...
#~/go/bin/go-cf-pwsh-gs <IP8080> <IP22> <region> <ami>
~/go/bin/go-cf-pwsh-gs "68.100.238.208" "100.101.102.103" "us-west-2" "ami-0a36eb8fadc976275"

tail -400 /var/log/cloud-init-output.log
```
