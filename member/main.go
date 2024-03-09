package member

import (
	"log"
)

func main() {
	svc := NewMemberService("http://localhost")
	apiServer := NewApiServer(svc)

	log.Fatal(apiServer.Start(":3000"))

}
