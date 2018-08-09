test_fund:
	go test -short -v fund_test.go fund.go fineract_client.go fineract_mock_client.go common.go error.go

test_fund_full:
	go test -v fund_test.go fund.go fineract_client.go fineract_mock_client.go common.go error.go

test_loan:
	go test -short -v loan_test.go loan.go fineract_client.go fineract_mock_client.go common.go error.go

test_loan_full:
	go test -v loan_test.go loan.go fineract_client.go fineract_mock_client.go common.go error.go

test_user_info:
	go test -v user_client_test.go user_client.go fineract_client.go fineract_mock_client.go common.go error.go fund.go

test_user_client:
	go test -v client_test.go client.go fineract_client.go fineract_mock_client.go common.go error.go fund.go

test_kyc:
	go test -short -v kycinfo_test.go kycinfo.go fineract_client.go fineract_mock_client.go error.go common.go fund.go fineract_client_test.go client.go

test_kyc_full:
	go test -v kycinfo_test.go kycinfo.go fineract_client.go fineract_mock_client.go error.go common.go fund.go fineract_client_test.go client.go