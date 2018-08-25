package imagesearch

import(
	"context"
	"log"
	proto "github.com/aeciovc/go-image-search/proto"
)

type Ping struct{}

func (p *Ping) Ping(ctx context.Context, req *proto.PingRequest, rsp *proto.PingResponse) error {
	log.Println("Responsing with pong...")
	rsp.Message = "Pong " + req.Name
	return nil
}

