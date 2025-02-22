package main

import (
	"time"

	"github.com/cloudentity/openbanking-quickstart/openbanking/obbr/consents/models"
	paymentModels "github.com/cloudentity/openbanking-quickstart/openbanking/obbr/payments/models"
	"github.com/gin-gonic/gin"
	"github.com/go-openapi/strfmt"

	acpClient "github.com/cloudentity/acp-client-go/models"
)

func OBBRMapError(c *gin.Context, err *Error) (code int, resp interface{}) {
	code, resp = err.Code, models.OpenbankingBrasilResponseError{
		Errors: []*models.OpenbankingBrasilError{
			{
				Detail: err.Message,
			},
		},
	}
	return
}

func NewOBBRAccountsResponse(accounts []AccountData) ResponseAccountList {
	return ResponseAccountList{
		Data: accounts,
	}
}

func OBBRPermsToStringSlice(perms []acpClient.OpenbankingBrasilPermission) []string {
	var slice []string
	for _, perm := range perms {
		slice = append(slice, string(perm))
	}
	return slice
}

func NewOBBRPayment(introspectionResponse *acpClient.IntrospectOBBRPaymentConsentResponse, self strfmt.URI, id string) paymentModels.OpenbankingBrasilResponsePixPayment {
	now := strfmt.DateTime(time.Now())
	status := paymentModels.OpenbankingBrasilStatus1PDNG
	localInstrument := paymentModels.OpenbankingBrasilEnumLocalInstrumentMANU
	return paymentModels.OpenbankingBrasilResponsePixPayment{
		Data: &paymentModels.OpenbankingBrasilResponsePixPaymentData{
			PaymentID:            id,
			ConsentID:            introspectionResponse.ConsentID,
			CreationDateTime:     now,
			StatusUpdateDateTime: now,
			Status:               &status,
			LocalInstrument:      &localInstrument,
			Payment: &paymentModels.OpenbankingBrasilPaymentPix{
				Amount:   introspectionResponse.OBBRCustomerPaymentConsent.Payment.Amount,
				Currency: introspectionResponse.OBBRCustomerPaymentConsent.Payment.Currency,
			},
			CreditorAccount: &paymentModels.OpenbankingBrasilCreditorAccount{},
		},
		Links: &paymentModels.OpenbankingBrasilLinks{
			Self: string(self),
		},
		Meta: &paymentModels.OpenbankingBrasilMeta{},
	}
}
