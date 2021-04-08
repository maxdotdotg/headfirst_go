package magazine

type Subscriber struct {
	Name    string
	Rate    float64
	Active  bool
	Address // this field is defined by it's type ONLY
	// the Address struct is now _embedded_ in the Subscriber struct
	//HomeAddress Address
}

type Employee struct {
	Name    string
	Salary  float64
	Address // defined by it's type ONLY
	//HomeAddress Address
	// with this removed, the underlying Address struct can be
	// accessed via anonymous fields
	// Employee.Address.Street, for example
	// or more likely, Employee.Street
}

type Address struct {
	Street     string
	City       string
	State      string
	PostalCode string
}
