package mmodels

// Hello contains the data for the GetHello routes
type Hello struct {
	Name string
}

// Message calculates the  message
func (model *Hello) Message() map[string]string {
	var message string
	if model.Name != "" {
		message = "Hello there " + model.Name + "!"
	} else {
		message = "Hello there!"
	}

	return map[string]string{
		"Message": message,
	}
}
