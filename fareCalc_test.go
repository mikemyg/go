package main

import "testing"

func TestFareCalc(t *testing.T) {
	cFareCalc := make(chan exportData)

	//test day cost
	data := []userData{{
		id: 1,
		data: []gpsData{{
			lat:       37.966660,
			lng:       23.728308,
			timestamp: 1405594957},
			{
				lat:       37.866627,
				lng:       23.728263,
				timestamp: 1405594966},
			{
				lat:       37.766625,
				lng:       23.728263,
				timestamp: 1405594974},
		},
	}}

	for _, v := range data {
		go fareCalc(v, cFareCalc)
	}

	var expData []exportData
	for i := 0; i < len(data); i++ {
		expData = append(expData, <-cFareCalc)
	}

	if expData[0].fare_estimate != 17.75972961033891 {
		t.Errorf("expected fare estimate 17.75972961033891 but got %v", expData[0].fare_estimate)
	}

	//test night cost
	data1 := []userData{{
		id: 1,
		data: []gpsData{{
			lat:       37.966660,
			lng:       23.728308,
			timestamp: 1564952401},
			{
				lat:       37.866627,
				lng:       23.728263,
				timestamp: 1564952411},
			{
				lat:       37.766625,
				lng:       23.728263,
				timestamp: 1564952421},
		},
	}}

	for _, v := range data1 {
		go fareCalc(v, cFareCalc)
	}

	var expData1 []exportData
	for i := 0; i < len(data1); i++ {
		expData1 = append(expData1, <-cFareCalc)
	}

	if expData1[0].fare_estimate != 30.21574120735214 {
		t.Errorf("expected fare estimate 30.21574120735214 but got %v", expData1[0].fare_estimate)
	}

	//test idle cost
	data2 := []userData{{
		id: 1,
		data: []gpsData{{
			lat:       37.966660,
			lng:       23.728308,
			timestamp: 1564952401},
			{
				lat:       37.966660,
				lng:       23.728308,
				timestamp: 1564952511},
			{
				lat:       37.966660,
				lng:       23.728308,
				timestamp: 1564953921},
		},
	}}

	for _, v := range data2 {
		go fareCalc(v, cFareCalc)
	}

	var expData2 []exportData
	for i := 0; i < len(data2); i++ {
		expData2 = append(expData2, <-cFareCalc)
	}

	if expData2[0].fare_estimate != 6.3244444444444445 {
		t.Errorf("expected fare estimate 6.3244444444444445 but got %v", expData2[0].fare_estimate)
	}

}

func TestCheckTimestamp(t *testing.T) {

	//test if day cost at 00:00:00
	time := checkTimestamp(1564952400)
	if !time {
		t.Errorf("expected true for day cost at 00:00:00 but got %v", time)
	}

	//test if night cost at 00:00:01
	time1 := checkTimestamp(1564952401)
	if time1 {
		t.Errorf("expected false for night cost at 00:00:01 but got %v", time1)
	}
	//test if night cost at 05:00:00
	time2 := checkTimestamp(1564884000)
	if time2 {
		t.Errorf("expected false for night cost at 05:00:00 but got %v", time2)
	}
	//test if day cost at 05:00:01
	time3 := checkTimestamp(1564884001)
	if !time3 {
		t.Errorf("expected true for day cost at 05:00:01 but got %v", time3)
	}

}

func TestCalcMoving(t *testing.T) {
	//day calc
	calc := calcMoving(1405594957, 11.12316279749979)
	if calc != 8.231140470149844 {
		t.Errorf("expected 8.231140470149844 but got %v", calc)
	}

	// night calc
	calc1 := calcMoving(1564952421, 11.12316279749979)
	if calc1 != 14.460111636749726 {
		t.Errorf("expected 14.460111636749726 but got %v", calc1)
	}
}

func TestCalcIdle(t *testing.T) {
	amount := calcIdle(9)
	if amount != 0.029750000000000002 {
		t.Errorf("expected 0.029750000000000002 but got %v", amount)
	}
}

//0.3636111111111111  4.660833333333334 6,32444444444444
