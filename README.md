### Usage
```
cd ~/environment/go-cf-pwsh-gs
go mod init go-cf-pwsh-gs
go install ./...
#~/go/bin/go-cf-pwsh-gs <IP8080> <IP22> <region> <gszipbucketobject> <ami> <keypair>
~/go/bin/go-cf-pwsh-gs "68.100.238.208" "100.101.102.103" "us-west-2" "deleteme1001/Deploy-Gs.zip" "ami-0a36eb8fadc976275" oregonkeypair "go_cf_pwsh_gs"

tail -400 /var/log/cloud-init-output.log
```
