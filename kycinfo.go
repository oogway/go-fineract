package fineract

import (
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strconv"
)

const (
	KycTableName = "datatables/m_kyc"
)

type KycInfo struct {
	BaseKycInfo
	ID       int64 `json:"ID"`
	ClientID int64 `json:"clientID"`
}

type KycInfoCreateRequest struct {
	BaseKycInfo
	ClientID   int64  `json:"-"`
	DateFormat string `json:"dateFormat"`
}

type KycInfoUpdateRequest struct {
	BaseKycInfo
	ID         int64  `json:"-"`
	ClientID   int64  `json:"-"`
	DateFormat string `json:"dateFormat"`
}

type BaseKycInfo struct {
	KtpUrl        string `json:"ktp_url"`
	KtpNo         string `json:"ktp_no"`
	SelfieUrl     string `json:"selfie_url"`
	FullName      string `json:"full_name"`
	Gender        Gender `json:"-"`
	GenderCode    int64  `json:"Gender_cd_gender"`
	DayOfBirth    string `json:"date_of_birth"`
	PlaceOfBirth  string `json:"place_of_birth"`
	HomeAddress   string `json:"home_address"`
	MaritalStatus string `json:"marital_status"`
	Rt            string `json:"rt"`
	Rw            string `json:"rw"`
	Village       string `json:"kelurahan"`
	District      string `json:"kecamatan"`
	//TODO: Add domicile_same field
	DomicileAddress  string  `json:"domicile_address"`
	DomicileRt       string  `json:"domicile_rt"`
	DomicileRw       string  `json:"domicile_rw"`
	DomicileVillage  string  `json:"domicile_kelurahan"`
	DomicileDistrict string  `json:"domicile_kecamatan"`
	PostalCode       string  `json:"postal_code"`
	Income           int64   `json:"income"`
	Occupation       string  `json:"occupation"`
	UserEmail        string  `json:"user_email"`
	UserMsisdn       string  `json:"user_msisdn"`
	UserId           string  `json:"user_id"`
	FaceSimilarity   float64 `json:"face_similarity"`
	NationalID       string  `json:"national_id"`
	Locale           string  `json:"locale"`
	ExtraInfos       string  `json:"extra_infos"`
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
	ID       int64 `json:"id"`
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
	return path.Join(baseURL, KycTableName)
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

	err = client.MakeRequest(http.MethodGet, client.HostName.ResolveReference(tempPath).String(), nil, &response)
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

	err = client.MakeRequest(http.MethodGet, client.HostName.ResolveReference(tempPath).String(), nil, &response)
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

	err = client.MakeRequest(http.MethodPost, client.HostName.ResolveReference(tempPath).String(), r, &response)
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

	err = client.MakeRequest(http.MethodPut, client.HostName.ResolveReference(tempPath).String(), r, &response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

// convert raw data to kyc
func (client *Client) rowToKYC(row []interface{}) (*KycInfo, error) {
	if len(row) < 7 {
		return nil, fmt.Errorf("InvalID KYC info: %+v", row)
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
		// TODO: Bugfix for out of index and non-string fields
		BaseKycInfo: BaseKycInfo{
			KtpUrl:           row[2].(string),
			KtpNo:            row[3].(string),
			SelfieUrl:        row[4].(string),
			FullName:         row[5].(string),
			Gender:           fromCode(genderCode),
			DayOfBirth:       row[7].(string),
			PlaceOfBirth:     row[8].(string),
			HomeAddress:      row[9].(string),
			MaritalStatus:    row[10].(string),
			Rt:               row[11].(string),
			Rw:               row[12].(string),
			Village:          row[13].(string),
			District:         row[14].(string),
			DomicileAddress:  row[16].(string),
			DomicileRt:       row[17].(string),
			DomicileRw:       row[18].(string),
			DomicileVillage:  row[19].(string),
			DomicileDistrict: row[20].(string),
			PostalCode:       row[21].(string),
			Income:           0,
			Occupation:       row[23].(string),
			UserEmail:        row[24].(string),
			UserMsisdn:       row[25].(string),
			UserId:           row[26].(string),
			FaceSimilarity:   0,
			NationalID:       row[28].(string),
			ExtraInfos:       row[29].(string),
		},
		ID:       ID,
		ClientID: clientID,
	}
	return kycInfo, nil
}
