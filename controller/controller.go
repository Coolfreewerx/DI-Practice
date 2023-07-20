package controller


type Controller struct {
	RequestId string
}

func New(requestId string) Controller {
	controllerObj := Controller{RequestId: requestId}
	return controllerObj
}