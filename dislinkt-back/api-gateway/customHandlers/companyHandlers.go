package customHandlers

import (
	"context"

	"github.com/blupulov/xwsDislinkt/api-gateway/grpcClients"
	cs "github.com/blupulov/xwsDislinkt/common/proto/services/company-service"
)

func (ch *CustomHandler) enableCompany(companyId string) error {
	companyClient := grpcClients.NewCompanyClient(ch.companyClientAddress)

	_, err := companyClient.EnableCompany(context.TODO(), &cs.EnableCompanyRequest{CompanyId: companyId})

	return err
}
