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
```
./main -n $APP_NAME -s $STAGE -t $TYPE
```
- APP_NAME e.g core-api
- STAGE e.g dev
- $TYPE e.g secret or general

