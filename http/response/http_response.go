package response

type HttpResponse struct {
	Code   int
	Status string
	Data   interface{}
}
