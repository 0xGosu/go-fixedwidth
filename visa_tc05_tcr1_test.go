package fixedwidth

import (
	"os"
	"testing"
)

type VISA_TC05_TCR1 struct {
	TransactionCode                             int          `fixed:"1,2"`
	MessageHashTotal                            string       `fixed:"3,4"`
	TransactionCodeQualifier                    int          `fixed:"5,5"`
	TransactionComponentSequenceNumber          int          `fixed:"6,6"`
	BusinessFormatCode                          string       `fixed:"7,7,none,_"` // disable padding
	TokenAssuranceLevel                         string       `fixed:"8,8,none,_"` // disable padding
	Reserved1                                   string       `fixed:"10,18"`
	Reserved2                                   string       `fixed:"19,24"`
	DocumentationIndicator                      string       `fixed:"25,25,none,_"` // disable padding
	MemberMessageText                           EbcdicString `fixed:"26,75"`
	SpecialConditionIndicator                   string       `fixed:"76,77,none,_"` // disable padding
	FeeProgramIndicator                         string       `fixed:"78,80"`
	IssuerCharge                                string       `fixed:"81,81,none,_"` // disable padding
	Reserved3                                   string       `fixed:"82,82"`
	CardAcceptorID                              string       `fixed:"83,97"`
	TerminalID                                  string       `fixed:"98,105"`
	NationalReimbursementFee                    string       `fixed:"106,117"`
	MailPhoneElectronicCommercePaymentIndicator string       `fixed:"118,118,none,_"` // disable padding
	SpecialChargebackIndicator                  string       `fixed:"119,119,none,_"` // disable padding
	ConversionDate                              string       `fixed:"120,123"`
	Reserved4                                   string       `fixed:"124,125"`
	AcceptanceTerminalIndicator                 string       `fixed:"126,126,none,_"` // disable padding
	PrepaidCardIndicator                        string       `fixed:"127,127,none,_"` // disable padding
	ServiceDevelopmentField                     string       `fixed:"128,128,none,_"` // disable padding
	AVSResponseCode                             string       `fixed:"129,129,none,_"` // disable padding
	AuthorizationSourceCode                     string       `fixed:"130,130,none,_"` // disable padding
	PurchaseIdentifierFormat                    string       `fixed:"131,131,none,_"` // disable padding
	AccountSelection                            string       `fixed:"132,132,none,_"` // disable padding
	InstallmentPaymentCount                     string       `fixed:"133,134,none,_"` // disable padding
	PurchaseIdentifier                          string       `fixed:"135,159"`
	CardBlock                                   string       `fixed:"160,168"`
	ChipConditionCode                           string       `fixed:"169,169,none,_"` // disable padding
	POSEnvironment                              string       `fixed:"170,170,none,_"` // disable padding
}

func TestVISA_TC05_TCR1_Parse_Test(t *testing.T) {
	// Read test file
	testFile := "./visa_tc05_tcr1_test.hex"
	data, err := os.ReadFile(testFile)
	if err != nil {
		t.Fatalf("Failed to read test file: %v", err)
	}

	// Convert to string and parse
	record := string(data)
	var visa VISA_TC05_TCR1
	err = Unmarshal([]byte(record), &visa)
	if err != nil {
		t.Fatalf("Failed to unmarshal record: %v", err)
	}

	// Verify parsed data matches expected values
	expected := VISA_TC05_TCR1{
		TransactionCode:                    25,
		MessageHashTotal:                   string([]byte{227, 48}), // � in the file
		TransactionCodeQualifier:           0,
		TransactionComponentSequenceNumber: 1,
		BusinessFormatCode:                 " ",
		TokenAssuranceLevel:                " ",
		Reserved1:                          "",
		Reserved2:                          "000000",
		DocumentationIndicator:             " ",
		MemberMessageText:                  EbcdicString{S: "IR International_Pre_PS2000"},
		SpecialConditionIndicator:          "  ",
		FeeProgramIndicator:                "903",
		IssuerCharge:                       " ",
		Reserved3:                          "",
		CardAcceptorID:                     "888888888888888",
		TerminalID:                         "MVV00001",
		NationalReimbursementFee:           "000000000000",
		MailPhoneElectronicCommercePaymentIndicator: " ",
		SpecialChargebackIndicator:                  " ",
		ConversionDate:                              "9116",
		Reserved4:                                   "00",
		AcceptanceTerminalIndicator:                 " ",
		PrepaidCardIndicator:                        " ",
		ServiceDevelopmentField:                     " ",
		AVSResponseCode:                             " ",
		AuthorizationSourceCode:                     " ",
		PurchaseIdentifierFormat:                    " ",
		AccountSelection:                            "0",
		InstallmentPaymentCount:                     "  ",
		PurchaseIdentifier:                          "ITEM DESCRIPTOR",
		CardBlock:                                   "000000000",
		ChipConditionCode:                           " ",
		POSEnvironment:                              " ",
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

	if visa.BusinessFormatCode != expected.BusinessFormatCode {
		t.Errorf("BusinessFormatCode: got %s, want %s", visa.BusinessFormatCode, expected.BusinessFormatCode)
	}

	if visa.TokenAssuranceLevel != expected.TokenAssuranceLevel {
		t.Errorf("TokenAssuranceLevel: got %s, want %s", visa.TokenAssuranceLevel, expected.TokenAssuranceLevel)
	}

	if visa.Reserved1 != expected.Reserved1 {
		t.Errorf("Reserved1: got %s, want %s", visa.Reserved1, expected.Reserved1)
	}

	if visa.Reserved2 != expected.Reserved2 {
		t.Errorf("Reserved2: got %s, want %s", visa.Reserved2, expected.Reserved2)
	}

	if visa.DocumentationIndicator != expected.DocumentationIndicator {
		t.Errorf("DocumentationIndicator: got %s, want %s", visa.DocumentationIndicator, expected.DocumentationIndicator)
	}

	if visa.MemberMessageText != expected.MemberMessageText {
		t.Errorf("MemberMessageText: got %s, want %s", visa.MemberMessageText, expected.MemberMessageText)
	}

	if visa.SpecialConditionIndicator != expected.SpecialConditionIndicator {
		t.Errorf("SpecialConditionIndicator: got %s, want %s", visa.SpecialConditionIndicator, expected.SpecialConditionIndicator)
	}

	if visa.FeeProgramIndicator != expected.FeeProgramIndicator {
		t.Errorf("FeeProgramIndicator: got %s, want %s", visa.FeeProgramIndicator, expected.FeeProgramIndicator)
	}

	if visa.IssuerCharge != expected.IssuerCharge {
		t.Errorf("IssuerCharge: got %s, want %s", visa.IssuerCharge, expected.IssuerCharge)
	}

	if visa.Reserved3 != expected.Reserved3 {
		t.Errorf("Reserved3: got %s, want %s", visa.Reserved3, expected.Reserved3)
	}

	if visa.CardAcceptorID != expected.CardAcceptorID {
		t.Errorf("CardAcceptorID: got %s, want %s", visa.CardAcceptorID, expected.CardAcceptorID)
	}

	if visa.TerminalID != expected.TerminalID {
		t.Errorf("TerminalID: got %s, want %s", visa.TerminalID, expected.TerminalID)
	}

	if visa.NationalReimbursementFee != expected.NationalReimbursementFee {
		t.Errorf("NationalReimbursementFee: got %s, want %s", visa.NationalReimbursementFee, expected.NationalReimbursementFee)
	}

	if visa.MailPhoneElectronicCommercePaymentIndicator != expected.MailPhoneElectronicCommercePaymentIndicator {
		t.Errorf("MailPhoneElectronicCommercePaymentIndicator: got %s, want %s", visa.MailPhoneElectronicCommercePaymentIndicator, expected.MailPhoneElectronicCommercePaymentIndicator)
	}

	if visa.SpecialChargebackIndicator != expected.SpecialChargebackIndicator {
		t.Errorf("SpecialChargebackIndicator: got %s, want %s", visa.SpecialChargebackIndicator, expected.SpecialChargebackIndicator)
	}

	if visa.ConversionDate != expected.ConversionDate {
		t.Errorf("ConversionDate: got %s, want %s", visa.ConversionDate, expected.ConversionDate)
	}

	// Continue with remaining field comparisons
	if visa.Reserved4 != expected.Reserved4 {
		t.Errorf("Reserved4: got %s, want %s", visa.Reserved4, expected.Reserved4)
	}

	if visa.AcceptanceTerminalIndicator != expected.AcceptanceTerminalIndicator {
		t.Errorf("AcceptanceTerminalIndicator: got %s, want %s", visa.AcceptanceTerminalIndicator, expected.AcceptanceTerminalIndicator)
	}

	if visa.PrepaidCardIndicator != expected.PrepaidCardIndicator {
		t.Errorf("PrepaidCardIndicator: got %s, want %s", visa.PrepaidCardIndicator, expected.PrepaidCardIndicator)
	}

	if visa.ServiceDevelopmentField != expected.ServiceDevelopmentField {
		t.Errorf("ServiceDevelopmentField: got %s, want %s", visa.ServiceDevelopmentField, expected.ServiceDevelopmentField)
	}

	if visa.AVSResponseCode != expected.AVSResponseCode {
		t.Errorf("AVSResponseCode: got %s, want %s", visa.AVSResponseCode, expected.AVSResponseCode)
	}

	if visa.AuthorizationSourceCode != expected.AuthorizationSourceCode {
		t.Errorf("AuthorizationSourceCode: got %s, want %s", visa.AuthorizationSourceCode, expected.AuthorizationSourceCode)
	}

	if visa.PurchaseIdentifierFormat != expected.PurchaseIdentifierFormat {
		t.Errorf("PurchaseIdentifierFormat: got %s, want %s", visa.PurchaseIdentifierFormat, expected.PurchaseIdentifierFormat)
	}

	if visa.AccountSelection != expected.AccountSelection {
		t.Errorf("AccountSelection: got %s, want %s", visa.AccountSelection, expected.AccountSelection)
	}

	if visa.InstallmentPaymentCount != expected.InstallmentPaymentCount {
		t.Errorf("InstallmentPaymentCount: got %s, want %s", visa.InstallmentPaymentCount, expected.InstallmentPaymentCount)
	}

	if visa.PurchaseIdentifier != expected.PurchaseIdentifier {
		t.Errorf("PurchaseIdentifier: got %s, want %s", visa.PurchaseIdentifier, expected.PurchaseIdentifier)
	}

	if visa.CardBlock != expected.CardBlock {
		t.Errorf("CardBlock: got %s, want %s", visa.CardBlock, expected.CardBlock)
	}

	if visa.ChipConditionCode != expected.ChipConditionCode {
		t.Errorf("ChipConditionCode: got %s, want %s", visa.ChipConditionCode, expected.ChipConditionCode)
	}

	if visa.POSEnvironment != expected.POSEnvironment {
		t.Errorf("POSEnvironment: got %s, want %s", visa.POSEnvironment, expected.POSEnvironment)
	}
}
