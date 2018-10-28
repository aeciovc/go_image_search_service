package imagesearch

import(
	"context"
	"log"
	proto "github.com/aeciovc/go-image-search/proto"
)

//GoMicro with Protobuf
type Ping struct{}

func (p *Ping) Ping(ctx context.Context, req *proto.PingRequest, rsp *proto.PingResponse) error {
	log.Println("[Ping] Responsing with pong...")
	rsp.Message = "Pong " + req.Name
	return nil
}

//RabbitMQ route
type RabbitMQService struct{}

func (rs *RabbitMQService) Ping() string {
	return "pong"
}

func (rs *RabbitMQService) Search(query string) []string {
	log.Println("[RabbitMQService] searching for ", query)
	return []string{"http://www.google.com.br/dkmfkng", "http://www.ghdjdkf.com"}
}