package fineract

import (
	"math/rand"
	"testing"
	"time"

	"log"

	"github.com/bmizerany/assert"
	"github.com/stretchr/testify/require"
)

func TestSuiteMockKYC(t *testing.T) {
	if !testing.Short() {
		t.Skip("Skipped mock tests in long mode")
	}
	client, err := makeClient(true)
	if err != nil {
		t.Fatal(err)
	}

	clientId := int64(3)
	KycSuite(t, client, clientId, "18733")
}

func TestSuiteKYC(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipped integrated tests in short mode")
	}
	client, err := makeClient(false)
	if err != nil {
		t.Fatal(err)
	}

	clientReq := &ClientInfo{
		FirstName:      "first name",
		LastName:       "last name",
		Active:         true,
		Locale:         "en",
		CountryCode:    "62",
		PhoneNumber:    toString(random(81100200000, 81100249999)),
		SubmitDate:     time.Now(),
		ActivationDate: time.Now(),
	}

	merchantName := "toko"
	merchantClientId := toString(random(11111111, 88888888))
	response, err := client.CreateClient(clientReq, merchantClientId, merchantName)
	require.Nil(t, err)
	require.NotNil(t, response)

	KycSuite(t, client, response.ID, toString(random(0, 99999)))
	KycISuite(t, client, response.ID)
}

func KycISuite(t *testing.T, client *Client, clientId int64) {
	fullName := "hung nguyen"
	updatedName := fullName + " updated"

	kycRes, err := client.GetKycInfosByClientID(&GetKycInfosByClientIDRequest{ClientID: clientId})
	assert.Equal(t, nil, err)
	if len(kycRes.KYCInfos) < 1 {
		t.Fatal("one kyc record should be avaiable to update")
	}
	kyc := kycRes.KYCInfos[0]
	kycId := kyc.ID

	t.Run("TestUpdateKYCInfo ", func(t *testing.T) {
		kyc := &KycInfoUpdateRequest{
			ID:       kycId,
			ClientID: clientId,
			BaseKycInfo: BaseKycInfo{
				FullName: updatedName,
			},
			DateFormat: "dd/MM/YYYY",
		}
		_, err := client.UpdateKYCInfo(kyc)
		assert.Equal(t, err, nil, "Error should be nil")

		kycRes, err := client.GetKycInfosByClientID(&GetKycInfosByClientIDRequest{ClientID: clientId})
		assert.Equal(t, nil, err)
		gKyc := kycRes.KYCInfos[0]
		assert.Equal(t, clientId, gKyc.ClientID, "Incorrect ClientID")
		assert.Equal(t, updatedName, gKyc.FullName, "Incorrect full name")
	})
}

func KycSuite(t *testing.T, client *Client, clientId int64, ktpNo string) {
	var kycId int64
	fullName := "hung nguyen"
	doB := "27/12/1988"
	formattedDoB := "1988-12-27"
	faceSimilarity := 10.0
	income := int64(100000)
	ktpURL := "http://google.co.in"
	selfieURL := "http://selfie-url.com"
	occupation := "student"
	postalCode := "411048"

	t.Run("TestCreateKYCInfo", func(t *testing.T) {
		kyc := &KycInfoCreateRequest{
			BaseKycInfo: BaseKycInfo{
				KtpUrl:           ktpURL,
				KtpNo:            ktpNo,
				SelfieUrl:        selfieURL,
				FullName:         fullName,
				Gender:           GenderMale,
				DayOfBirth:       doB,
				PlaceOfBirth:     "jakarta",
				HomeAddress:      "home address",
				MaritalStatus:    "kawin",
				Rt:               "rt",
				Rw:               "rw",
				Village:          "village",
				District:         "district",
				DomicileAddress:  "address",
				DomicileRt:       "rt",
				DomicileRw:       "rw",
				DomicileVillage:  "village",
				DomicileDistrict: "district",
				PostalCode:       postalCode,
				Income:           income,
				Occupation:       occupation,
				UserEmail:        "abc@def.com",
				UserMsisdn:       "81100200000",
				UserId:           "id",
				FaceSimilarity:   faceSimilarity,
				NationalID:       "123456789",
				Locale:           "en",
			},
			ClientID:   clientId,
			DateFormat: "dd/MM/YYYY",
		}
		res, err := client.CreateKYCInfo(kyc)
		if err != nil {
			log.Println(err)
			t.Fatal("kyc creation failed")
		}
		assert.Equal(t, err, nil, "Error should be nil")
		assert.Equal(t, res.ClientID, clientId, "Incorrect clientID")
	})

	t.Run("TestGetKycInfosByClientID", func(t *testing.T) {
		kycRes, err := client.GetKycInfosByClientID(&GetKycInfosByClientIDRequest{ClientID: clientId})
		log.Println(err)
		log.Println(kycRes)
		assert.Equal(t, nil, err)
		//assert.Equal(t, 1, len(kycRes.KYCInfos), "one kyc record should be fetched")
		kyc := kycRes.KYCInfos[0]
		assert.Equal(t, clientId, kyc.ClientID, "Incorrect ClientID")
		assert.Equal(t, fullName, kyc.FullName, "Incorrect full name")
		assert.Equal(t, ktpNo, kyc.KtpNo, "recorded ktp no doesnt match")
		assert.Equal(t, formattedDoB, kyc.DayOfBirth, "Date of Birth doesnt match")
		assert.Equal(t, faceSimilarity, kyc.FaceSimilarity)
		assert.Equal(t, income, kyc.Income)
		assert.Equal(t, Gender(GenderMale), kyc.Gender, "Incorrect gender")
		assert.Equal(t, ktpURL, kyc.KtpUrl)
		assert.Equal(t, selfieURL, kyc.SelfieUrl)
		assert.Equal(t, occupation, kyc.Occupation, "occupation doesnot match")
		assert.Equal(t, postalCode, kyc.PostalCode, "postalCode doesnot match")

		kycId = kyc.ID
	})

	t.Run("TestCreateKYCInfo:Add another KYC for this client", func(t *testing.T) {
		kyc := &KycInfoCreateRequest{
			BaseKycInfo: BaseKycInfo{
				KtpUrl:           ktpURL,
				KtpNo:            ktpNo,
				SelfieUrl:        selfieURL,
				FullName:         fullName,
				Gender:           GenderMale,
				DayOfBirth:       doB,
				PlaceOfBirth:     "jakarta",
				HomeAddress:      "home address",
				MaritalStatus:    "kawin",
				Rt:               "rt",
				Rw:               "rw",
				Village:          "village",
				District:         "district",
				DomicileAddress:  "address",
				DomicileRt:       "rt",
				DomicileRw:       "rw",
				DomicileVillage:  "village",
				DomicileDistrict: "district",
				PostalCode:       postalCode,
				Income:           income,
				Occupation:       occupation,
				UserEmail:        "abc@def.com",
				UserMsisdn:       "81100200000",
				UserId:           "id",
				FaceSimilarity:   faceSimilarity,
				NationalID:       "123456789",
				Locale:           "en",
			},
			ClientID:   clientId,
			DateFormat: "dd/MM/YYYY",
		}
		res, err := client.CreateKYCInfo(kyc)
		if err != nil {
			log.Println(err)
			t.Fatal("kyc creation failed")
		}
		assert.Equal(t, err, nil, "Error should be nil")
		assert.Equal(t, res.ClientID, clientId, "Incorrect clientID")
	})

	t.Run("TestGetKycInfosByClientID for multiple kycs", func(t *testing.T) {
		kycRes, err := client.GetKycInfosByClientID(&GetKycInfosByClientIDRequest{ClientID: clientId})
		assert.Equal(t, nil, err)
		assert.Equal(t, 2, len(kycRes.KYCInfos), "two kyc records should be fetched")
		kyc := kycRes.KYCInfos[1]
		assert.Equal(t, clientId, kyc.ClientID, "Incorrect ClientID")
		assert.Equal(t, fullName, kyc.FullName, "Incorrect full name")
		assert.Equal(t, ktpNo, kyc.KtpNo, "recorded ktp no doesnt match")
		assert.Equal(t, formattedDoB, kyc.DayOfBirth, "Date of Birth doesnt match")
		assert.Equal(t, faceSimilarity, kyc.FaceSimilarity)
		assert.Equal(t, income, kyc.Income)
		assert.Equal(t, Gender(GenderMale), kyc.Gender, "Incorrect gender")
		assert.Equal(t, ktpURL, kyc.KtpUrl)
		assert.Equal(t, selfieURL, kyc.SelfieUrl)
		assert.Equal(t, occupation, kyc.Occupation, "occupation doesnot match")
		assert.Equal(t, postalCode, kyc.PostalCode, "postalCode doesnot match")
	})
}

func random(min, max int) uint64 {
	rand.Seed(time.Now().Unix())
	return uint64(rand.Intn(max-min) + min)
}
