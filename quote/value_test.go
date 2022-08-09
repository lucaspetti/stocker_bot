package quote

import (
	"context"
	"testing"
)

func TestGetValueFunc(t *testing.T) {

}

type mockDefaultApiService struct {
	expectedSymbol string
}

func (s *mockDefaultApiService) CompanyBasicFinancials(ctx context.Context) {

}
