# Alerts
Alerts Recording System.
The API Alerts are used to store alerts of many services and retrieve Alerts of a particular service in a specific period of time.

## Storage:
The API can be more effectively designed using DB but will make the execution of the code a bit challenging, So I have used a local system as DB by store the data in the form of json files.
I have designed the json files in a way I would have designed a SQL DB where I have added AlertService ID which is the foreign key for the table Alerts ie. the primary key in the Service table.
In the same way, I have designed two json files (Service.json and Alerts.json) to store the data.

## Rest Service
The API is designed mux for routes.

## Design
The API will store the alerts in JSON files, where the service file will only add data if the service ID does not exist in the file. However, the alerts will be added to the file every time because the alert might have the same service ID and alert ID but a different timestamp.
The API Routes are as follows.

### Write Alerts:
Users should be able to send requests to this API to write alert data to the chosen data storage.
Write Request
HTTP Method: POST
Endpoint: /alerts
Request Body: 
```
{
"alert_id": "b950482e9911ec7e41f7ca5e5d9a424f",
"service_id": "my_test_service_id",
"service_name": "my_test_service",
"model": "my_test_model",
"alert_type": "anomaly",
"alert_ts": "1695644160",
"severity": "warning",
"team_slack": "slack_ch"
}
```
HTTP Status Code: 200 OK
Response Body:
```
{
"alert_id": "b950482e9911ec7e41f7ca5e5d9a424f",
"error": ""
}
```

HTTP Status Code: 500 Internal Server Error
Response Body:

```
{
"alert_id": "b950482e9911ec7e41f7ca5e5d9a424f",
"error": "<error details>"
}
```

### Read Request:
HTTP Method: GET
Endpoint: /alerts
Query Parameters:
service_id: The identifier of the service for which alerts are requested.
start_ts: The starting timestamp epoch of the time period.
end_ts: The ending timestamp epoch of the time period.
Example: /alerts?
service_id=my_test_service_id&start_ts=1695643160&end_ts=1695644360
Read Response
Success
HTTP Status Code: 200 OK
Response Body:
```
{
"service_id" : "my_test_service_id"
"service_name": "my_test_service",
"alerts" : [
{
"alert_id": "b950482e9911ec7e41f7ca5e5d9a424f",
"model": "my_test_model",
"alert_type": "anomaly",
"alert_ts": "1695644060",
"severity": "warning",
"team_slack": "slack_ch"
},
{
"alert_id": "b950482e9911ecsdfs41f75e5d9az23cv",
"model": "my_test_model",
"alert_type": "anomaly",
"alert_ts": "1695644160",
"severity": "warning",
"team_slack": "slack_ch"
},
]
}
```

HTTP Status Code: Appropriate HTTP error status (e.g., 400 Bad Request, 404 Not Found, 500 Internal
Server Error)
Response Body:
```
{
"alert_id": "b950482e9911ec7e41f7ca5e5d9a424f",
"error": "<error details>"
}
```


## Steps to RUN the code.
- clone the Repo.
- move inside the Alerts directory.
- ```go mod tidy```
- ```go run main.go```
  
