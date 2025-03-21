package fixedwidth

import (
	"os"
	"testing"
)

type VISA_TC_05_TCR_0 struct {
	TransactionCode int `fixed:"1,2"`
	// TODO : Handle MessageHashTotal as a byte array
	MessageHashTotal                                string `fixed:"3,4"`
	TransactionCodeQualifier                        int    `fixed:"5,5"`
	TransactionComponentSequenceNumber              int    `fixed:"6,6"`
	AccountNumber                                   int    `fixed:"7,22"`
	AccountNumberExtension                          int    `fixed:"23,25"`
	FloorLimitIndicator                             string `fixed:"26,26,none,_"` // disable padding
	CRBExceptionFileIndicator                       string `fixed:"27,27,none,_"` // disable padding
	PositiveCardholderAuthorizationServiceIndicator string `fixed:"28,28,none,_"` // disable padding
	AcquirerReferenceNumber                         string `fixed:"29,51"`
	AcquirerBusinessID                              string `fixed:"52,59"`
	PurchaseDate                                    string `fixed:"60,63"`
	DestinationAmount                               string `fixed:"64,75"`
	DestinationCurrencyCode                         string `fixed:"76,78"`
	SourceAmount                                    string `fixed:"79,90"`
	SourceCurrencyCode                              string `fixed:"91,93"`
	MerchantName                                    string `fixed:"94,118"`
	MerchantCity                                    string `fixed:"119,131"`
	MerchantCountryCode                             string `fixed:"132,134"`
	MerchantCategoryCode                            string `fixed:"135,138"`
	MerchantZIPCode                                 string `fixed:"139,143"`
	MerchantStateProvinceCode                       string `fixed:"144,146"`
	RequestedPaymentService                         string `fixed:"147,147,none,_"` // disable padding
	NumberOfPaymentForms                            string `fixed:"148,148,none,_"` // disable padding
	UsageCode                                       int    `fixed:"149,149"`
	ReasonCode                                      string `fixed:"150,151"`
	SettlementFlag                                  int    `fixed:"152,152"`
	AuthorizationCharacteristicsIndicator           string `fixed:"153,153"`
	AuthorizationCode                               string `fixed:"154,159"`
	POSTerminalCapability                           string `fixed:"160,160,none,_"` // disable padding
	ReservedField1                                  string `fixed:"161,161,none,_"` // disable padding
	CardholderIDMethod                              string `fixed:"162,162,none,_"` // disable padding
	CollectionOnlyFlag                              string `fixed:"163,163,none,_"` // disable padding
	POSEntryMode                                    string `fixed:"164,165,none,_"` // disable padding
	CentralProcessingDate                           string `fixed:"166,169"`
	ReimbursementAttribute                          string `fixed:"170,170,none,_"` // disable padding
}

func TestVISA_TC05_Parse_Test1(t *testing.T) {
	// Read test file
	testFile := "./visa_tc05_test1.data"
	data, err := os.ReadFile(testFile)
	if err != nil {
		t.Fatalf("Failed to read test file: %v", err)
	}

	// Convert to string and parse
	record := string(data)
	var visa VISA_TC_05_TCR_0
	err = Unmarshal([]byte(record), &visa)
	if err != nil {
		t.Fatalf("Failed to unmarshal record: %v", err)
	}

	// Verify parsed data matches expected values
	expected := VISA_TC_05_TCR_0{
		TransactionCode:                                 25,
		MessageHashTotal:                                string([]byte{28, 133}), // Represented as �� in the file
		TransactionCodeQualifier:                        0,
		TransactionComponentSequenceNumber:              0,
		AccountNumber:                                   4830970000162705,
		AccountNumberExtension:                          0,
		FloorLimitIndicator:                             " ",
		CRBExceptionFileIndicator:                       " ",
		PositiveCardholderAuthorizationServiceIndicator: " ",
		AcquirerReferenceNumber:                         "74064499116000000155872",
		AcquirerBusinessID:                              "10021249",
		PurchaseDate:                                    "0426",
		DestinationAmount:                               "000000003042",
		DestinationCurrencyCode:                         "840",
		SourceAmount:                                    "000000004001",
		SourceCurrencyCode:                              "124",
		MerchantName:                                    "MERCHANT NAME",
		MerchantCity:                                    "MERCHANT CITY",
		MerchantCountryCode:                             "CA",
		MerchantCategoryCode:                            "5611",
		MerchantZIPCode:                                 "00000",
		MerchantStateProvinceCode:                       "AB",
		RequestedPaymentService:                         " ",
		NumberOfPaymentForms:                            " ",
		UsageCode:                                       1,
		ReasonCode:                                      "00",
		SettlementFlag:                                  0,
		AuthorizationCharacteristicsIndicator:           "N",
		AuthorizationCode:                               "123456",
		POSTerminalCapability:                           "2",
		ReservedField1:                                  "2",
		CardholderIDMethod:                              "1",
		CollectionOnlyFlag:                              " ",
		POSEntryMode:                                    "05",
		CentralProcessingDate:                           "9116",
		ReimbursementAttribute:                          "B",
	}

	// Compare fields to check parsing accuracy
	if visa.TransactionCode != expected.TransactionCode {
		t.Errorf("TransactionCode: got %d, want %d", visa.TransactionCode, expected.TransactionCode)
	}

	if string(visa.MessageHashTotal) != string(expected.MessageHashTotal) {
		t.Errorf("MessageHashTotal: got %v, want %v", []byte(visa.MessageHashTotal), []byte(expected.MessageHashTotal))
	}

	if visa.TransactionCodeQualifier != expected.TransactionCodeQualifier {
		t.Errorf("TransactionCodeQualifier: got %d, want %d", visa.TransactionCodeQualifier, expected.TransactionCodeQualifier)
	}

	if visa.TransactionComponentSequenceNumber != expected.TransactionComponentSequenceNumber {
		t.Errorf("TransactionComponentSequenceNumber: got %d, want %d", visa.TransactionComponentSequenceNumber, expected.TransactionComponentSequenceNumber)
	}

	if visa.AccountNumber != expected.AccountNumber {
		t.Errorf("AccountNumber: got %d, want %d", visa.AccountNumber, expected.AccountNumber)
	}

	if visa.AccountNumberExtension != expected.AccountNumberExtension {
		t.Errorf("AccountNumberExtension: got %d, want %d", visa.AccountNumberExtension, expected.AccountNumberExtension)
	}

	if visa.FloorLimitIndicator != expected.FloorLimitIndicator {
		t.Errorf("FloorLimitIndicator: got %v, want %v", visa.FloorLimitIndicator, expected.FloorLimitIndicator)
	}

	if visa.CRBExceptionFileIndicator != expected.CRBExceptionFileIndicator {
		t.Errorf("CRBExceptionFileIndicator: got %s, want %s", visa.CRBExceptionFileIndicator, expected.CRBExceptionFileIndicator)
	}

	if visa.PositiveCardholderAuthorizationServiceIndicator != expected.PositiveCardholderAuthorizationServiceIndicator {
		t.Errorf("PositiveCardholderAuthorizationServiceIndicator: got %s, want %s", visa.PositiveCardholderAuthorizationServiceIndicator, expected.PositiveCardholderAuthorizationServiceIndicator)
	}

	if visa.AcquirerReferenceNumber != expected.AcquirerReferenceNumber {
		t.Errorf("AcquirerReferenceNumber: got %s, want %s", visa.AcquirerReferenceNumber, expected.AcquirerReferenceNumber)
	}

	if visa.AcquirerBusinessID != expected.AcquirerBusinessID {
		t.Errorf("AcquirerBusinessID: got %s, want %s", visa.AcquirerBusinessID, expected.AcquirerBusinessID)
	}

	if visa.PurchaseDate != expected.PurchaseDate {
		t.Errorf("PurchaseDate: got %s, want %s", visa.PurchaseDate, expected.PurchaseDate)
	}

	if visa.DestinationAmount != expected.DestinationAmount {
		t.Errorf("DestinationAmount: got %s, want %s", visa.DestinationAmount, expected.DestinationAmount)
	}

	if visa.DestinationCurrencyCode != expected.DestinationCurrencyCode {
		t.Errorf("DestinationCurrencyCode: got %s, want %s", visa.DestinationCurrencyCode, expected.DestinationCurrencyCode)
	}

	if visa.SourceAmount != expected.SourceAmount {
		t.Errorf("SourceAmount: got %s, want %s", visa.SourceAmount, expected.SourceAmount)
	}

	if visa.SourceCurrencyCode != expected.SourceCurrencyCode {
		t.Errorf("SourceCurrencyCode: got %s, want %s", visa.SourceCurrencyCode, expected.SourceCurrencyCode)
	}

	if visa.MerchantName != expected.MerchantName {
		t.Errorf("MerchantName: got %s, want %s", visa.MerchantName, expected.MerchantName)
	}

	if visa.MerchantCity != expected.MerchantCity {
		t.Errorf("MerchantCity: got %s, want %s", visa.MerchantCity, expected.MerchantCity)
	}

	if visa.MerchantCountryCode != expected.MerchantCountryCode {
		t.Errorf("MerchantCountryCode: got %s, want %s", visa.MerchantCountryCode, expected.MerchantCountryCode)
	}

	if visa.MerchantCategoryCode != expected.MerchantCategoryCode {
		t.Errorf("MerchantCategoryCode: got %s, want %s", visa.MerchantCategoryCode, expected.MerchantCategoryCode)
	}

	if visa.MerchantZIPCode != expected.MerchantZIPCode {
		t.Errorf("MerchantZIPCode: got %s, want %s", visa.MerchantZIPCode, expected.MerchantZIPCode)
	}

	if visa.MerchantStateProvinceCode != expected.MerchantStateProvinceCode {
		t.Errorf("MerchantStateProvinceCode: got %s, want %s", visa.MerchantStateProvinceCode, expected.MerchantStateProvinceCode)
	}

	if visa.RequestedPaymentService != expected.RequestedPaymentService {
		t.Errorf("RequestedPaymentService: got %s, want %s", visa.RequestedPaymentService, expected.RequestedPaymentService)
	}

	if visa.NumberOfPaymentForms != expected.NumberOfPaymentForms {
		t.Errorf("NumberOfPaymentForms: got %v, want %v", visa.NumberOfPaymentForms, expected.NumberOfPaymentForms)
	}

	if visa.UsageCode != expected.UsageCode {
		t.Errorf("UsageCode: got %v, want %v", visa.UsageCode, expected.UsageCode)
	}

	if visa.ReasonCode != expected.ReasonCode {
		t.Errorf("ReasonCode: got %v, want %v", visa.ReasonCode, expected.ReasonCode)
	}

	if visa.SettlementFlag != expected.SettlementFlag {
		t.Errorf("SettlementFlag: got %v, want %v", visa.SettlementFlag, expected.SettlementFlag)
	}

	if visa.AuthorizationCharacteristicsIndicator != expected.AuthorizationCharacteristicsIndicator {
		t.Errorf("AuthorizationCharacteristicsIndicator: got %s, want %s", visa.AuthorizationCharacteristicsIndicator, expected.AuthorizationCharacteristicsIndicator)
	}

	if visa.AuthorizationCode != expected.AuthorizationCode {
		t.Errorf("AuthorizationCode: got %s, want %s", visa.AuthorizationCode, expected.AuthorizationCode)
	}

	if visa.POSTerminalCapability != expected.POSTerminalCapability {
		t.Errorf("POSTerminalCapability: got %s, want %s", visa.POSTerminalCapability, expected.POSTerminalCapability)
	}

	if visa.ReservedField1 != expected.ReservedField1 {
		t.Errorf("ReservedField1: got %s, want %s", visa.ReservedField1, expected.ReservedField1)
	}

	if visa.CardholderIDMethod != expected.CardholderIDMethod {
		t.Errorf("CardholderIDMethod: got %s, want %s", visa.CardholderIDMethod, expected.CardholderIDMethod)
	}

	if visa.CollectionOnlyFlag != expected.CollectionOnlyFlag {
		t.Errorf("CollectionOnlyFlag: got %s, want %s", visa.CollectionOnlyFlag, expected.CollectionOnlyFlag)
	}

	if visa.POSEntryMode != expected.POSEntryMode {
		t.Errorf("POSEntryMode: got %s, want %s", visa.POSEntryMode, expected.POSEntryMode)
	}

	if visa.CentralProcessingDate != expected.CentralProcessingDate {
		t.Errorf("CentralProcessingDate: got %v, want %v", visa.CentralProcessingDate, expected.CentralProcessingDate)
	}

	if visa.ReimbursementAttribute != expected.ReimbursementAttribute {
		t.Errorf("ReimbursementAttribute: got %s, want %s", visa.ReimbursementAttribute, expected.ReimbursementAttribute)
	}

	// Additional field verification can be added for other important fields
}
