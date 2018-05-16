test_fund:
	go test -short -v fund_test.go fund.go fineract_client.go fineract_mock_client.go common.go error.go

test_fund_full:
	go test -v fund_test.go fund.go fineract_client.go fineract_mock_client.go common.go error.go