package service


type GrpcService struct {
	ip string
	port uint16

}

func (service *GrpcService) Init(ip string , port uint16 ,  options ...interface{}) error {
	service.ip = ip
	service.port = port
	return nil
}

func (service *GrpcService) Start() error {

	return nil
}
