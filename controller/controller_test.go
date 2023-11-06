package controller_test

import (
	"Alerts/controller"
	"Alerts/model"
	"testing"
)

func TestFilterAlertsByTimeTS(t *testing.T) {
	// Test case 1: Alert within the time range
	alertWithinRange := model.Alerts{AlertTs: 1695644000}
	startTS := "1695640000"
	endTS := "1695650000"
	result := controller.FilterAlertsByTimeTS(alertWithinRange, startTS, endTS)
	if !result {
		t.Errorf("Expected alert to be within the time range, but it was filtered out")
	}
	// Test case 2: Alert before the time range
	alertBeforeRange := model.Alerts{AlertTs: 1695630000}
	result = controller.FilterAlertsByTimeTS(alertBeforeRange, startTS, endTS)
	if result {
		t.Errorf("Expected alert to be before the time range, but it was not filtered out")
	}

	// Test case 3: Alert after the time range
	alertAfterRange := model.Alerts{AlertTs: 1695660000}
	result = controller.FilterAlertsByTimeTS(alertAfterRange, startTS, endTS)
	if result {
		t.Errorf("Expected alert to be after the time range, but it was not filtered out")
	}
}

func TestCheckServiceID(t *testing.T) {
	// Test case 1: Service ID exists in the slice
	services := []model.Service{
		{ServiceID: "service1", ServiceName: "Service 1"},
		{ServiceID: "service2", ServiceName: "Service 2"},
		{ServiceID: "service3", ServiceName: "Service 3"},
	}
	serviceIDToCheck := "service2"
	result := controller.CheckServiceID(services, serviceIDToCheck)
	if !result {
		t.Errorf("Expected service ID to be found, but it was not")
	}

	// Test case 2: Service ID does not exist in the slice
	serviceIDToCheck = "service4"
	result = controller.CheckServiceID(services, serviceIDToCheck)
	if result {
		t.Errorf("Expected service ID to not be found, but it was found")
	}

	// Test case 3: Empty slice
	emptySlice := []model.Service{}
	serviceIDToCheck = "service1"
	result = controller.CheckServiceID(emptySlice, serviceIDToCheck)
	if result {
		t.Errorf("Expected service ID to not be found in an empty slice")
	}
}
