package purchase

import (
	"errors"
	"github.com/Rhymond/go-money"
	"github.com/google/uuid"
	"github.com/mamalmaleki/coffeego"
	"github.com/mamalmaleki/coffeego/internal/payment"
	"github.com/mamalmaleki/coffeego/internal/store"
	"time"
)

type Purchase struct {
	id                 uuid.UUID
	Store              store.Store
	ProductsToPurchase []coffeego.Product
	total              money.Money
	PaymentMeans       payment.Means
	timeOfPurchase     time.Time
	CardToken          *string
}

func (p *Purchase) validateAndEnrich() error {
	if len(p.ProductsToPurchase) == 0 {
		return errors.New("purchase must consist of at least one product")
	}
	p.total = *money.New(0, "USD")
	for _, v := range p.ProductsToPurchase {
		newTotal, _ := p.total.Add(&v.BasePrice)
		p.total = *newTotal
	}
	if p.total.IsZero() {
		return errors.New("likely mistake; purchase should never be 0")
	}

	p.id = uuid.New()
	p.timeOfPurchase = time.Now()

	return nil
}