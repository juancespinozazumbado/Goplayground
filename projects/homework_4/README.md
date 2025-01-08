# Package sync System

### useage

-  run the command

```pws
  go mod tidy
```

run the server

```shell
  go run ./
```

- in other terminal run curl request

Add packages to the queue

```bash
  curl --location --request POST 'http://localhost:8080/add-package' \
       --header 'Content-Type: application/json' \
       --data'{
                 "Destination": "North",
                 "SenderName": "Webhook"
            }'
```

retrieve sorted packages 
```shell
  curl --location --request GET 'http://localhost:8080/get-packages?zone=Zone%201'
```