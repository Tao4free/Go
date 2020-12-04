package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"reflect"
	"time"

	"cloud.google.com/go/datastore"
)

type MonthGcpUsage struct {
	BillingSubAccount string    `datastore:"billingSubAccount"`
	Cost              float64   `datastore:"cost"`
	DeleteFlag        bool      `datastore:"deleteFlag"`
	Detail            string    `datastore:"detail"`
	Id                string    `datastore:"id"`
	Project           string    `datastore:"project"`
	RegisterDate      time.Time `datastore:"registerDate"`
	Version           int64     `datastore:"version"`
	YearMonth         string    `datastore:"yearMonth"`
}

type OrderDetail struct {
	BillingCode             string    `datastore:"billingCode"`
	BillingDepartmentNumber string    `datastore:"billingDepartmentNumber"`
	BillingInfoCode         string    `datastore:"billingInfoCode"`
	CommissionRate          float64   `datastore:"comissionRate"`
	DeleteFlag              bool      `datastore:"deleteFlag"`
	ExcludeCommission       float64   `datastore:"excludeComission"`
	FreeWord                string    `datastore:"freeWord"`
	HideTargetMonthFlag     bool      `datastore:"hideTargetMonthFlag"`
	InvoiceDeliveryMethod   string    `datastore:"invoiceDeliveryMethod"`
	MaxComission            int64     `datastore:"maxComission"`
	MinComission            int64     `datastore:"minComission"`
	Notes                   string    `datastore:"notes"`
	OnlyUsageFeeFlag        bool      `datastore:"onlyUsageFeeFlag"`
	OrderNo                 int64     `datastore:"orderNo"`
	OutputByPaper           string    `datastore:"outputByPaper"`
	OverSeasSubsidiaryFlg   bool      `datastore:"overSeasSubsidiaryFlg"`
	PaymentDateChangedFlag  bool      `datastore:"paymentDateChangedFlag"`
	PaymentDayOfMonth       int64     `datastore:"paymentDayOfMonth"`
	PaymentMonthPeriod      int64     `datastore:"paymentMonthPeriod"`
	PaymentSupportAmount    int64     `datastore:"paymentSupportAmount"`
	PaymentSupportPlanName  string    `datastore:"paymentSupportPlanName"`
	PrepaidAccountFlag      bool      `datastore:"prepaidAccountFlag"`
	RegisterDate            time.Time `datastore:"registerDate"`
	SpecialDiscount         int64     `datastore:"specialDiscount"`
	SubAccount              string    `datastore:"subAccount"`
	TechSupportAmount       int64     `datastore:"techSupportAmount"`
	TechSupportPlanName     string    `datastore:"techSupportPlanName"`
	UniqueAccountFlag       bool      `datastore:"uniqueAccountFlag"`
	UniqueAccountNotes      string    `datastore:"uniqueAccountNotes"`
	UpdateDate              time.Time `datastore:"updateDate"`
	Version                 int64     `datastore:"version"`
}

type CustomizeProduct struct {
	DeleteFlag   bool           `datastore:"deleteFlag"`
	Id           int64          `datastore:"id"`
	Name         string         `datastore:"name"`
	OrderDetail  *datastore.Key `datastore:"orderDetail"`
	Price        float64        `datastore:"price"`
	Quantity     int64          `datastore:"quantity"`
	RegisterDate time.Time      `datastore:"registerDate"`
	UpdateDate   time.Time      `datastore:"updateDate"`
	Version      int64          `datastore:"version"`
}

type Exchange struct {
	DeleteFlag   bool      `datastore:"deleteFlag"`
	Month        string    `datastore:"month"`
	Rate         string    `datastore:"rate"`
	RegisterDate time.Time `datastore:"registerDate"`
	Registerer   string    `datastore:"registerer"`
	UpdateDate   time.Time `datastore:"updateDate"`
}

type InvoiceStock struct {
	OrderNo   int64  `datastore:"orderNo"`
	ChargeNo  int64  `datastore:"chargeNo"`
	ProjectId string `datastore:"projectId"`
	UnitPrice int64  `datastore:"unitPrice"`
	Quantity  int64  `datastore:"quantity"`
	YearMonth string `datastore:"yearMonth"`
}

type TestWrite struct {
	Nihaoma    bool   `datastore:"nihaoma"`
	TheWorld   string `datastore:"theWorld"`
	DoNotWorry int64  `datastore:"doNotWorry"`
}

const (
	MIN_COMISSION_FEE = int64(1500)
)

var monthGcpUsage []MonthGcpUsage
var monthGcpUsageDollar []MonthGcpUsage
var orderDetail []OrderDetail
var customizeProduct []CustomizeProduct
var exchange []Exchange
var invoiceStock []InvoiceStock
var invoicePerOrder []InvoiceStock

var yearMonth string = "default"

func main() {
	invoiceStock = []InvoiceStock{{OrderNo: 130034, ChargeNo: 0, ProjectId: "", UnitPrice: 0, Quantity: 1, YearMonth: "2019-11"}, {OrderNo: 130034, ChargeNo: 1, ProjectId: "ca-tsukubakokusai-univ", UnitPrice: 52841, Quantity: 1, YearMonth: "2019-11"}, {OrderNo: 130034, ChargeNo: 2, ProjectId: "", UnitPrice: 1500, Quantity: 1, YearMonth: "2019-11"}}
	fmt.Printf("%+v \n", invoiceStock)

	putInvoiceStock()
}

func mkInvoice(orderNo int64, chargeNo int64, projectId string, unitPrice int64, quantity int64) InvoiceStock {
	return InvoiceStock{
		OrderNo:   orderNo,
		ChargeNo:  chargeNo,
		ProjectId: projectId,
		UnitPrice: unitPrice,
		Quantity:  quantity,
		YearMonth: yearMonth}
}

func putInvoiceStock() {
	num := len(invoiceStock)

	invoiceStockInterface := make([]interface{}, num)
	for i := range invoiceStock {
		invoiceStockInterface[i] = &invoiceStock[i]
	}
	fmt.Printf("%+v \n", len(invoiceStock))

	write(invoiceStockInterface)
}

func write(invoiceStockInterface []interface{}) {

	file, err := os.OpenFile("aaaa.csv", os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	WriteCSV(file, invoiceStockInterface)
}

// Generate csv rows including header from interface{} slice or object
func genRows(sl []interface{}) [][]string {
	const COLTAG = "datastore"

	rows := make([][]string, 1)

	for n, d := range sl {
		rows = append(rows, []string{})
		v := reflect.ValueOf(d).Elem()
		// reflect.ValueOf(&orderDetailEntity).Elem()
		fmt.Printf("#############3%+v\n", v.Type())

		for i := 0; i < v.NumField(); i++ {
			if n == 0 {
				// Header
				colName := v.Type().Field(i).Tag.Get(COLTAG)
				rows[0] = append(rows[0], colName)
			}
			rows[len(rows)-1] = append(rows[len(rows)-1], fmt.Sprint(v.Field(i).Interface()))
		}
	}
	return rows
}

// Write csv file to path
func WriteCSV(file *os.File, data []interface{}) {
	rows := genRows(data)

	writer := csv.NewWriter(file)
	for _, row := range rows {
		writer.Write(row)
	}
	writer.Flush()
}
