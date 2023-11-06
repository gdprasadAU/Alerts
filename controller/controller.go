package controller

import (
	"Alerts/model"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

func ReadFromDB(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	// getting params from the url.
	paramServiceID := param["service_id"]
	paramStartTS := param["start_ts"]
	paramEndTS := param["end_ts"]
	//Read file Service.json which acts as a DB.
	serFileContents, err := os.ReadFile("Service.json")
	if err != nil {
		fmt.Println("error", err.Error())

	}
	services := []model.Service{}
	//check for empty file.
	if string(serFileContents) != "" {
		err := json.Unmarshal(serFileContents, &services)
		if err != nil {
			fmt.Println("error", err.Error())
		}
	}
	var serviceName string
	// To check if the service id exists in the db
	for _, service := range services {
		if service.ServiceID == paramServiceID {
			serviceName = service.ServiceName
		}
	}
	// Read Alert.json for DB
	contents, _ := os.ReadFile("Alert.json")
	alerts := []model.Alerts{}

	// check for empty alerts in DB
	if string(contents) != "" {
		err := json.Unmarshal(contents, &alerts)
		if err != nil {
			fmt.Println("error", err.Error())
		}
	}
	var serviceAlerts []model.Alerts
	alertIDs := []string{}
	// check if Service id  match the Service id in Alerts
	for _, alert := range alerts {
		if alert.AlertServiceID == paramServiceID {
			if FilterAlertsByTimeTS(alert, paramStartTS, paramEndTS) {
				serviceAlerts = append(serviceAlerts, alert)
				alertIDs = append(alertIDs, alert.AlertID)
			}
		}
	}

	resService := model.ResService{ServiceID: paramServiceID, ServiceName: serviceName, Alerts: serviceAlerts}

	fmt.Println("complete data :", resService)
	resp := model.Respo{AlertID: strings.Join(alertIDs, ","), Error: ""}
	byteRes, err := json.Marshal(resp)
	if err != nil {
		fmt.Println("error", err.Error())
	}

	_, err = w.Write(byteRes)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

// To check if the particular alert exists in the given ts
func FilterAlertsByTimeTS(alert model.Alerts, startTS string, endTS string) bool {

	startTSUnix, err := strconv.ParseInt(startTS, 10, 64)
	if err != nil {
		fmt.Println(err)
	}
	endTSUnix, err := strconv.ParseInt(endTS, 10, 64)
	if err != nil {
		fmt.Println(err)
	}

	if alert.AlertTs >= startTSUnix && alert.AlertTs <= endTSUnix {
		return true
	}

	return false
}
func WriteToDB(w http.ResponseWriter, r *http.Request) {
	// reading data from body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()
	var reqBody model.ReqBody
	json.Unmarshal(body, &reqBody)

	service := model.Service{ServiceID: reqBody.ServiceID, ServiceName: reqBody.ServiceName}
	// storing the data into Service.json file
	err = StoreService(service)
	if err != nil {
		resp := model.Respo{AlertID: reqBody.AlertID, Error: err.Error()}
		byteRes, err := json.Marshal(resp)
		if err != nil {
			fmt.Println("error", err.Error())
		}
		_, err = w.Write(byteRes)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	alertTS, err := strconv.ParseInt(reqBody.AlertTs, 10, 64)
	if err != nil {
		fmt.Println("error", err.Error())
	}

	alerts := model.Alerts{AlertID: reqBody.AlertID, Model: reqBody.Model, AlertType: reqBody.AlertType, AlertTs: alertTS, Severity: reqBody.Severity, AlertServiceID: reqBody.ServiceID, TeamSlack: reqBody.TeamSlack}
	//storing alerts
	err = StoreAlerts(alerts)
	if err != nil {
		resp := model.Respo{AlertID: reqBody.AlertID, Error: err.Error()}
		byteRes, err := json.Marshal(resp)
		if err != nil {
			fmt.Println("error", err.Error())
		}
		_, err = w.Write(byteRes)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	resp := model.Respo{AlertID: reqBody.AlertID, Error: ""}
	byteRes, err := json.Marshal(resp)
	if err != nil {
		fmt.Println("error", err.Error())
	}

	_, err = w.Write(byteRes)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
func StoreService(service model.Service) error {
	// check if srvice.json file exists or else create a file which acts as a db.
	serviceFile, err := os.OpenFile("Service.json", os.O_CREATE|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		fmt.Println("Error Creating a file")
	}
	// Reading file
	serviceContents, _ := os.ReadFile("Service.json")
	goServiceContents := []model.Service{}
	// check if file is empty
	if string(serviceContents) != "" {
		err = json.Unmarshal(serviceContents, &goServiceContents)

		if err != nil {
			fmt.Println("error", err.Error())
		}
		// check if service id already exists in the DB
		existsFlag := CheckServiceID(goServiceContents, service.ServiceID)
		if existsFlag {
			fmt.Println("Service exits")
		} else {
			goServiceContents = append(goServiceContents, service)
		}
	} else {
		goServiceContents = append(goServiceContents, service)
	}

	byteService, err := json.Marshal(goServiceContents)
	if err != nil {
		fmt.Println("Error encoding a file")
	}
	// write into DB
	_, err = serviceFile.Write(byteService)
	if err != nil {
		return err
	}
	serviceFile.Close()
	return nil
}

// to check if sservice id exits in the DB
// returns false if it does not exists.
func CheckServiceID(services []model.Service, serviceIDToCheck string) bool {
	for _, service := range services {
		if service.ServiceID == serviceIDToCheck {
			return true
		}
	}
	return false
}
func StoreAlerts(alerts model.Alerts) error {
	// check if the file exists or else create a file which acts as a db.
	alertFile, err := os.OpenFile("Alert.json", os.O_CREATE|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		fmt.Println("Error Creating a file")
	}
	// reading the file
	contents, _ := os.ReadFile("Alert.json")
	goContents := []model.Alerts{}
	//check if file is empty
	if string(contents) != "" {
		err = json.Unmarshal(contents, &goContents)
		if err != nil {
			fmt.Println("error", err.Error())
		}
	}
	// appending contents on to the file
	goContents = append(goContents, alerts)

	byteAlert, err := json.Marshal(goContents)
	if err != nil {
		fmt.Println("Error encoding a file")
	}
	// writing content tomthe file
	_, err = alertFile.Write(byteAlert)
	if err != nil {
		fmt.Println("Error writing a file", err.Error())
	}
	alertFile.Close()
	return nil
}
