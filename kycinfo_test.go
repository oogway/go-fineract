package fineract

import (
	"testing"

	"github.com/bmizerany/assert"
)

func TestClient_GetKycInfosByClientID(t *testing.T) {
	client, err := makeClient(true)
	if err != nil {
		t.Fatal(err)
	}
	kycsRes, err := client.GetKycInfosByClientID(&GetKycInfosByClientIDRequest{ClientID: 1})
	kycs := kycsRes.KYCInfos
	assert.Equal(t, nil, err)
	assert.Equal(t, 2, len(kycs))
	for _, kyc := range kycs {
		assert.NotEqual(t, nil, kyc, "kyc should not be null")
	}
	assert.Equal(t, kycs[0].ID, uint64(8), "Incorrect ID")
	assert.Equal(t, kycs[0].ClientID, uint64(2), "Incorrect ClientID")
	assert.Equal(t, kycs[0].FullName, "kyc full name 1", "Incorrect full name")
	assert.Equal(t, kycs[0].NationalID, "123", "Incorrect national ID")
	assert.Equal(t, kycs[0].HomeAddress, "kyc address", "Incorrect address")
	assert.Equal(t, kycs[0].Gender, Gender(GenderMale), "Incorrect gender")
	assert.Equal(t, kycs[0].DayOfBirth, "2018-07-02", "Incorrect dob")
	assert.Equal(t, kycs[0].ExtraInfos, "{\"extraInfo\":\"extraInfo\"}", "Incorrect extra infos")

	assert.Equal(t, kycs[1].ID, uint64(9), "Incorrect ID")
	assert.Equal(t, kycs[1].ClientID, uint64(2), "Incorrect ClientID")
	assert.Equal(t, kycs[1].FullName, "test name 2", "Incorrect full name")
	assert.Equal(t, kycs[1].NationalID, "123456", "Incorrect national ID")
	assert.Equal(t, kycs[1].HomeAddress, "address", "Incorrect address")
	assert.Equal(t, kycs[1].Gender, Gender(GenderFemale), "Incorrect gender")
	assert.Equal(t, kycs[1].DayOfBirth, "2018-07-01", "Incorrect dob")
	assert.Equal(t, kycs[1].ExtraInfos, "{}", "Incorrect extra infos")
}

func TestClient_GetKycInfosByID(t *testing.T) {
	client, err := makeClient(true)
	if err != nil {
		t.Fatal(err)
	}
	kycsRes, err := client.GetKycInfoByID(&GetKycInfoByIDRequest{ClientID: 3, ID: 17})
	kyc := kycsRes.KYCInfo
	assert.Equal(t, nil, err)
	assert.Equal(t, kyc != nil, true, "kyc should not be nil")
	assert.Equal(t, kyc.ID, uint64(17), "Incorrect ID")
	assert.Equal(t, kyc.ClientID, uint64(3), "Incorrect ClientID")
	assert.Equal(t, kyc.FullName, "test name", "Incorrect full name")
	assert.Equal(t, kyc.NationalID, "1234567", "Incorrect national ID")
	assert.Equal(t, kyc.HomeAddress, "132 ham nghi", "Incorrect address")
	assert.Equal(t, kyc.Gender, Gender(GenderMale), "Incorrect gender")
	assert.Equal(t, kyc.DayOfBirth, "1988-07-01", "Incorrect dob")
	assert.Equal(t, kyc.ExtraInfos, "{\"extraInfo\":\"extraInfo\"}", "Incorrect extra infos")
}

func TestClient_GetKycInfosByIDWithNotExistedID(t *testing.T) {
	client, err := makeClient(true)
	if err != nil {
		t.Fatal(err)
	}
	kycsRes, err := client.GetKycInfoByID(&GetKycInfoByIDRequest{ClientID: 3, ID: 18})
	kyc := kycsRes.KYCInfo
	assert.Equal(t, nil, err)
	assert.Equal(t, kyc == nil, true, "kyc should not be nil")
}

func TestClient_CreateKYCInfo(t *testing.T) {
	client, err := makeClient(true)
	if err != nil {
		t.Fatal(err)
	}
	kyc := &KycInfoCreateRequest{
		BaseKycInfo: BaseKycInfo{
			FullName:    "hung nguyen",
			ExtraInfos:  "{}",
			Gender:      GenderMale,
			NationalID:  "123456789",
			HomeAddress: "home address",
			DayOfBirth:  "27/12/1988",
			Locale:      "en",
		},
		ClientID:   3,
		DateFormat: "dd/MM/YYYY",
	}
	res, err := client.CreateKYCInfo(kyc)
	assert.Equal(t, err, nil, "Error should be nil")
	assert.Equal(t, res.OfficeID, uint64(1), "Incorrect officeID")
	assert.Equal(t, res.ClientID, uint64(3), "Incorrect clientID")
	assert.Equal(t, res.ResourceID, uint64(3), "Incorrect resourceID")

}

func TestClient_UpdateKYCInfo(t *testing.T) {
	client, err := makeClient(true)
	if err != nil {
		t.Fatal(err)
	}
	kyc := &KycInfoUpdateRequest{
		ID:       16,
		ClientID: 3,
		BaseKycInfo: BaseKycInfo{
			FullName:    "hung nguyen new name 123",
			ExtraInfos:  "{}",
			Gender:      GenderMale,
			NationalID:  "123456789",
			HomeAddress: "home address",
			DayOfBirth:  "27/12/1988",
			Locale:      "en",
		},
		DateFormat: "dd/MM/YYYY",
	}
	res, err := client.UpdateKYCInfo(kyc)
	assert.Equal(t, err, nil, "Error should be nil")
	assert.Equal(t, res.OfficeID, uint64(1), "Incorrect officeID")
	assert.Equal(t, res.ClientID, uint64(3), "Incorrect clientID")
	assert.Equal(t, res.ResourceID, uint64(3), "Incorrect resourceID")

}
