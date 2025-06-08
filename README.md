## How to Run the Microservice

Follow these steps to run and test the microservice:

### 1. Run the Service

Use the following command to start the microservice:

```
make run
```

### 2. Test the API
Once the service is running, you can access the API using the following GET request:
```
GET http://localhost:3000/fact
```
You can test this using a browser, curl, Postman, or any HTTP client of your choice.


Example using curl:
```
curl http://localhost:3000/fact
```

This should return a JSON response with the fact provided by the service.
```json
{
  "fact": "Tigers are excellent swimmers and do not avoid water."
}
```
