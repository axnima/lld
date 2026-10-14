package main

import(
	"time"
	"errors"
	"fmt"
)

// Topic: money representation. Store minor units (12550 paise = INR 125.50) to avoid float rounding errors.
// Use NewMoney and keep amounts nonnegative; exported fields allow callers to bypass this contract.
type Money struct {
	MinorUnits int64
	Currency string
}

// Topic: named types and constants. These names document states but do not prevent other string values.
// The zero value is "", so explicitly initialize a new quote with Pending.
type QuoteStatus string

const (
	Pending QuoteStatus = "pending"
	Completed QuoteStatus = "completed"
)

// Topic: identity and lifecycle. ID stays stable as Status changes; this example assigns IDs manually.
type Quote struct {
	ID string
	Price Money
	ExpiresAt time.Time
	Status QuoteStatus
}

// Strategy: interchangeable pricing rules share Apply; use an interface when callers need different rules.
type PricePolicy interface {
	Apply(price Money) (Money , error)
}

// An empty struct suits a policy with no settings; matching Apply implicitly implements PricePolicy.
type RegularPrice struct{
}

// Policy configuration belongs to the policy, so the service need not understand discount percentages.
type PercentageDiscount struct{
	Discount int64
}

// Composition: the service has a policy; it depends on its behavior, not a concrete discount type.
type PricingService struct{
	Policy PricePolicy
}

// Topic: value semantics. Money contains value fields, so returning it does not expose shared mutable data.
func (r RegularPrice) Apply(price Money) (Money , error) {
	return price, nil
}

// Topic: integer rounding. Discounts include 0% and 100%; round the final price down to a minor unit.
func (p PercentageDiscount) Apply(price Money) (Money , error) {
	if p.Discount < 0 || p.Discount > 100 {
		return Money{} , errors.New("Validation failed : Invalid discount value")
	}
	remaining := 100 - p.Discount
	amount := price.MinorUnits
	// Split quotient and remainder to avoid amount*remaining overflow without losing the fractional contribution.
	newPrice := (amount/100)*remaining + (amount%100)*remaining/100
	return Money{
		MinorUnits : newPrice,
		Currency : price.Currency,
	}, nil
}

// Topic: delegation and errors. New policies implement Apply without requiring changes to Calculate.
// Reject a missing policy and propagate its result/error directly; callers must check errors before using results.
func (p PricingService) Calculate(price Money) (Money , error) {
	if p.Policy == nil {
		return Money{} , errors.New("Validation failed : Policy not set for this service")
	}
	return p.Policy.Apply(price)
}

// Topic: time boundaries. Valid only while now < ExpiresAt; equality counts as expired.
// Supplying now lets manual examples use exact timestamps without sleeping or depending on the wall clock.
func (q Quote) IsExpired(now time.Time) bool {
	if now.Before(q.ExpiresAt) {
		return false;
	}
	return true;
}

// Topic: money arithmetic. Add only matching, nonempty currencies and return a new value without mutating inputs.
// For nonnegative int64 operands, an overflowing sum wraps negative and is therefore less than a.MinorUnits.
func Add(a Money , b Money) (Money , error) {
	if a.Currency == b.Currency && a.MinorUnits + b.MinorUnits >= a.MinorUnits && len(a.Currency) > 0{
		return Money{
			MinorUnits : a.MinorUnits + b.MinorUnits,
			Currency : a.Currency,
		}, nil
	} else if a.Currency != b.Currency {
		return Money{}, errors.New("Validation failed : currency mismatch")
	} else if a.MinorUnits + b.MinorUnits < a.MinorUnits {
		return Money{}, errors.New("Validation failed : integer overflow")
	} else if len(a.Currency) == 0 || len(b.Currency) == 0 {
		return Money{}, errors.New("Validation failed : currency not specified")
	} else {
		return Money{}, errors.New("Operation failed")
	}

}

// Topic: constructors and error contracts. Zero is valid; negative amounts and empty currencies are rejected.
// Money{} accompanies failure, nil accompanies success; currency codes are not checked against a registry.
func NewMoney(amount int64 , currency string) (Money , error) {
	if amount < 0 {
		return Money{}, errors.New("Validation failed : amount less than 0")
	} else if len(currency) == 0 {
		return Money{}, errors.New("Validation failed : currency not specified")
	}
	return Money{
		MinorUnits : amount,
		Currency : currency,
	}, nil
}

// State machine: only Pending may become Completed, and failed transitions leave the quote unchanged.
// A pointer receiver mutates the original quote; the value receiver on IsExpired only reads it.
// The State pattern delegates behavior to state objects; these two states only need simple checks.
// Observer (not implemented): email and analytics listeners could react after successful acceptance.
// Command (not implemented): represent an acceptance request as data to queue for later execution.
func (q *Quote) Accept(now time.Time) error {
	if q.IsExpired(now) {
		return errors.New("Operation failed : Quote expired")
	} else if q.Status != Pending {
		return errors.New("Operation failed : Quote does not have Pending status")
	}
	q.Status = Completed
	return nil
}

func main() {
	// Manual demo 1: inject two strategies; expect 1000 INR and 800 INR, both with nil errors.
	a, _ := NewMoney(1000 , "INR")
	regularService := PricingService{
		Policy : RegularPrice{},
	}
	discountService := PricingService{
		Policy : PercentageDiscount{
			Discount : 20,
		},
	}
	fmt.Println(regularService.Calculate(a))
	fmt.Println(discountService.Calculate(a))

	// Manual demo 2: acceptance before expiry succeeds; accepting the same quote again fails.
	quote := Quote{
		ID:        "quote-1",
		Price:     a,
		ExpiresAt: time.Now().Add(5 * time.Minute),
		Status:    Pending,
	}

	if err := quote.Accept(time.Now()); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Quote Accepted")
	}

	if err := quote.Accept(time.Now()); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Quote Accepted")
	}

	// Manual demo 3: use a fresh Pending quote at exact expiry; expect an error and unchanged status.
	price, err := NewMoney(1000, "INR")
	if err != nil {
		fmt.Println(err)
		return
	}

	expiry := time.Date(2026, 10, 14, 12, 0, 0, 0, time.UTC)

	quote2 := Quote{
		ID:        "quote-2",
		Price:     price,
		ExpiresAt: expiry,
		Status:    Pending,
	}

	err = quote2.Accept(expiry)
	fmt.Println(err)          // quote expired
	fmt.Println(quote2.Status) // pending

}
