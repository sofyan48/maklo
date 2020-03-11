# MAKLO 
Generate Or Insert AWS System Manager

## Installing
```
go build -o maklo
```

## Setting Environtment
```
cp .env.example .env
```

## Usage
Note: ***Run the app one place with environtment***
### GeneratePath Parameters
```
maklo generatePath -p /rll/dev/general/sdk_js -name sdk_js -s dev -d true
```
OPTIONS:
```
--path value, -p value       File Template Path
--name value, -n value       App Name
--stage value, -s value      Stage Parameters
--decrypt value, -d value    Decryption Option
```

### Insert Paramters
```
maklo insert -p templates/parameters.json
```
OPTIONS:
```
--path value, -p value       File Template Path
--overwrite value, -w value  Overwirte Option
```

### Generate By Templates
```
maklo generate -f yaml -p templates/parameter.yaml
```
OPTIONS:
```
--path value, -p value    Templates Path
--format value, -f value  Templates Formats | yaml or json
```