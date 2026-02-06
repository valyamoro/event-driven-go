package events

type UserData struct {
	UserID    string `json:"user_id"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Action    string `json:"action"`
	Timestamp string `json:"timestamp"`
}

type OrderData struct {
	OrderID     string  `json:"order_id"`
	UserID      string  `json:"user_id"`
	TotalAmount float64 `json:"total_amount"`
	Status      string  `json:"status"`
	Items       []OrderItem `json:"items"`
	ShippingAddress string `json:"shipping_address"`
	Reason      string `json:"reason,omitempty"`
}

type OrderItem struct {
	ProductID   string  `json:"product_id"`
	Name        string  `json:"name"`
	Quantity    int     `json:"quantity"`
	Price       float64 `json:"price"`
}

type PaymentData struct {
	PaymentID   string  `json:"payment_id"`
	OrderID     string  `json:"order_id"`
	Amount      float64 `json:"amount"`
	Currency    string  `json:"currency"`
	Status      string  `json:"status"`
	Method      string  `json:"method"`
	FailedReason string `json:"failed_reason,omitempty"`
	RefundAmount float64 `json:"refund_amount,omitempty"`
}

type InventoryData struct {
	ProductID   string `json:"product_id"`
	ProductName string `json:"product_name"`
	CurrentQty  int    `json:"current_quantity"`
	MinQty      int    `json:"min_required_quantity"`
	Warehouse   string `json:"warehouse"`
	Urgency     string `json:"urgency"`
}

type ReviewData struct {
	ReviewID    string  `json:"review_id"`
	ProductID   string  `json:"product_id"`
	UserID      string  `json:"user_id"`
	Rating      int     `json:"rating"`
	Title       string  `json:"title"`
	Comment     string  `json:"comment,omitempty"`
	VerifiedPurchase bool `json:"verified_purchase"`
}

type PromoCodeData struct {
	Code        string  `json:"code"`
	UserID      string  `json:"user_id"`
	OrderID     string  `json:"order_id"`
	Discount    float64 `json:"discount_amount"`
	DiscountPct float64 `json:"discount_percent,omitempty"`
	MinAmount   float64 `json:"min_order_amount,omitempty"`
}

type AlertData struct {
	Severity    string `json:"severity"`
	Service     string `json:"service"`
	Message     string `json:"message"`
	Code        string `json:"code,omitempty"`
	Action      string `json:"action"`
}
