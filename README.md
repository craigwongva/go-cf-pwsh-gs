### Usage
```
cd ~/environment/go-cf-pwsh-gs
go mod init go-cf-pwsh-gs
go install ./...
#~/go/bin/go-cf-pwsh-gs <IP8080> <IP22> <githubpassword>
~/go/bin/go-cf-pwsh-gs "68.100.238.208/32" "0.0.0.0/0" <githubpassword>

tail -400 /var/log/cloud-init-output.log
```
