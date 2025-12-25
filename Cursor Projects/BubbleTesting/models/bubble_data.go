package models

// BubbleResponse represents a generic response from Bubble.io API
type BubbleResponse struct {
	Status  string      `json:"status"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

// BubbleRecord represents a record in Bubble.io
type BubbleRecord struct {
	ID       string                 `json:"_id"`
	Created  string                 `json:"Created Date"`
	Modified string                 `json:"Modified Date"`
	Fields   map[string]interface{} `json:"fields"`
}

// BubbleQuery represents a query to Bubble.io API
type BubbleQuery struct {
	Constraints []Constraint `json:"constraints"`
	Sort        []Sort       `json:"sort,omitempty"`
	Limit       int          `json:"limit,omitempty"`
	Skip        int          `json:"skip,omitempty"`
}

// Constraint represents a query constraint
type Constraint struct {
	Key            string      `json:"key"`
	ConstraintType string      `json:"constraint_type"`
	Value          interface{} `json:"value"`
}

// Sort represents a sort parameter
type Sort struct {
	Key       string `json:"key"`
	Direction string `json:"direction"`
}

// NewBubbleQuery creates a new query with the given constraints
func NewBubbleQuery(constraints ...Constraint) *BubbleQuery {
	return &BubbleQuery{
		Constraints: constraints,
	}
}

// AddConstraint adds a constraint to the query
func (q *BubbleQuery) AddConstraint(key, constraintType string, value interface{}) {
	q.Constraints = append(q.Constraints, Constraint{
		Key:            key,
		ConstraintType: constraintType,
		Value:          value,
	})
}

// AddSort adds a sort parameter to the query
func (q *BubbleQuery) AddSort(key, direction string) {
	q.Sort = append(q.Sort, Sort{
		Key:       key,
		Direction: direction,
	})
}
