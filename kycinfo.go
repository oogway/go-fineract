package fineract

import (
	"path"
	"fmt"
	"net/url"
	"strconv"
)

const (
	kycInfo = "datatables/m_kyc"
)

type KycInfo struct {
	BaseKycInfo
	ID       int64 `json:"ID"`
	ClientID int64 `json:"clientID"`
}

type KycInfoCreateRequest struct {
	BaseKycInfo
	ClientID   int64 `json:"-"`
	DateFormat string `json:"dateFormat"`
}

type KycInfoUpdateRequest struct {
	BaseKycInfo
	ID         int64 `json:"-"`
	ClientID   int64 `json:"-"`
	DateFormat string `json:"dateFormat"`
}

type BaseKycInfo struct {
	FullName    string `json:"full_name"`
	NationalID  string `json:"national_id"`
	HomeAddress string `json:"home_address"`
	DayOfBirth  string `json:"date_of_birth"`
	Gender      Gender `json:"-"`
	GenderCode  int64 `json:"Gender_cd_gender"`
	Locale      string `json:"locale"`
	ExtraInfos  string `json:"extra_infos"`
}
type Gender string

const (
	GenderMale   = "Male"
	GenderFemale = "Female"
)

func fromCode(code int64) Gender {
	if code == 14 {
		return GenderMale
	}
	return GenderFemale
}

func fromGender(gender Gender) int64 {
	if gender == GenderMale {
		return 14
	}
	return 15
}

type GetKycInfoByIDRequest struct {
	ClientID int64 `json:"clientID"`
	ID       int64 `json:"clientID"`
}

type GetKycInfoByIDResponse struct {
	KYCInfo *KycInfo `json:"kycInfos"`
}
type GetKycInfosByClientIDRequest struct {
	ClientID int64 `json:"clientID"`
}

type GetKycInfosByClientIDResponse struct {
	KYCInfos []KycInfo `json:"kycInfos"`
}

type CreateKycInfoResponse struct {
	OfficeID   int64 `json:"officeID"`
	ClientID   int64 `json:"clientID"`
	ResourceID int64 `json:"resourceID"`
}

func kycPath() string {
	return path.Join(baseURL, kycInfo)
}

func kycWithClientIDPath(clientID int64) string {
	return kycPath() + fmt.Sprintf("/%d?genericResultSet=true", clientID)
}

func kycWithIDPath(clientID int64, ID int64) string {
	return kycPath() + fmt.Sprintf("/%d/%d?genericResultSet=true", clientID, ID)
}

// GetCurrencyCode gets selected currency code from fineract
func (client *Client) GetKycInfosByClientID(r *GetKycInfosByClientIDRequest) (*GetKycInfosByClientIDResponse, error) {
	tempPath, err := url.Parse(kycWithClientIDPath(r.ClientID))
	if err != nil {
		return nil, err
	}
	response := make(map[string]interface{})

	err = client.MakeRequest("GET", client.HostName.ResolveReference(tempPath).String(), nil, &response)
	if err != nil {
		return nil, err
	}
	var kycInfos []KycInfo
	rowsData := response["data"]
	if nil != rowsData {
		rows := rowsData.([]interface{})
		for _, row := range rows {
			kycRawData := row.(map[string]interface{})["row"].([]interface{})
			kycInfo, err := client.rowToKYC(kycRawData)
			if nil == err {
				kycInfos = append(kycInfos, *kycInfo)
			}
		}
	}
	return &GetKycInfosByClientIDResponse{
		KYCInfos: kycInfos,
	}, nil
}

// GetCurrencyCode gets selected currency code from fineract
func (client *Client) GetKycInfoByID(r *GetKycInfoByIDRequest) (*GetKycInfoByIDResponse, error) {
	tempPath, err := url.Parse(kycWithIDPath(r.ClientID, r.ID))
	if err != nil {
		return nil, err
	}
	response := make(map[string]interface{})

	err = client.MakeRequest("GET", client.HostName.ResolveReference(tempPath).String(), nil, &response)
	if err != nil {
		return nil, err
	}
	var kycInfo *KycInfo
	rowsData := response["data"]
	if nil != rowsData {
		rows := rowsData.([]interface{})
		if len(rows) > 0 {
			kycRawData := rows[0].(map[string]interface{})["row"].([]interface{})
			kycInfo, _ = client.rowToKYC(kycRawData)
		}
	}

	return &GetKycInfoByIDResponse{
		KYCInfo: kycInfo,
	}, nil
}

// Create KYC info
func (client *Client) CreateKYCInfo(r *KycInfoCreateRequest) (*CreateKycInfoResponse, error) {
	tempPath, err := url.Parse(kycWithClientIDPath(r.ClientID))
	r.GenderCode = fromGender(r.Gender)
	if err != nil {
		return nil, err
	}
	var response *CreateKycInfoResponse

	err = client.MakeRequest("POST", client.HostName.ResolveReference(tempPath).String(), r, &response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

// Update KYC info
func (client *Client) UpdateKYCInfo(r *KycInfoUpdateRequest) (*CreateKycInfoResponse, error) {
	tempPath, err := url.Parse(kycWithIDPath(r.ClientID, r.ID))
	r.GenderCode = fromGender(r.Gender)
	if err != nil {
		return nil, err
	}
	var response *CreateKycInfoResponse

	err = client.MakeRequest("PUT", client.HostName.ResolveReference(tempPath).String(), r, &response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

// convert raw data to kyc
func (client *Client) rowToKYC(row []interface{}) (*KycInfo, error) {
	if len(row) < 7 {
		return nil, fmt.Errorf("InvalID KYC info", row)
	}
	ID, err := strconv.ParseInt(row[0].(string), 10, 64)
	if nil != err {
		return nil, err
	}
	clientID, err := strconv.ParseInt(row[1].(string), 10, 64)
	if nil != err {
		return nil, err
	}
	genderCode, err := strconv.ParseInt(row[6].(string), 10, 8)
	kycInfo := &KycInfo{
		BaseKycInfo: BaseKycInfo{
			FullName:    row[2].(string),
			NationalID:  row[3].(string),
			HomeAddress: row[4].(string),
			DayOfBirth:  row[5].(string),
			Gender:      fromCode(genderCode),
			ExtraInfos:  row[7].(string),
		},
		ID:       ID,
		ClientID: clientID,
	}
	return kycInfo, nil
}
