# Go Generate AWS System Manager


## Installing
```
go build main.go
```

## Setting Environtment
```
cp .env.example .env
```

## Usage
Note: ***Run the app one place with environtment***
### Generate Parameters
```
maklo generate -p /rll/dev/general/sdk_js -name sdk_js -s dev -d true
```
Option:
--path value, -p value       File Template Path
--name value, -n value       App Name
--stage value, -s value      Stage Parameters
--decrypt value, -d value    Decryption Option

### Insert Paramters
```
maklo insert -p templates/parameters.json
```
OPTIONS:
--path value, -p value       File Template Path
--overwrite value, -w value  Overwirte Option