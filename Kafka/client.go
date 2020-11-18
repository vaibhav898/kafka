package main

import (
  "log"

  "golang.org/x/net/context"
  "google.golang.org/grpc"

  "github.com/tutorialedge/go-grpc-beginners-tutorial/chat"
)

func main() {

  consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
     "bootstrap.servers":    "host1:9092,host2:9092",
     "group.id":             "foo",
     "auto.offset.reset":    "smallest"})

  var conn *grpc.ClientConn
  conn, err := grpc.Dial(":9000", grpc.WithInsecure())
  if err != nil {
    log.Fatalf("did not connect: %s", err)
  }
  defer conn.Close()

  c := chat.NewChatServiceClient(conn)

  response, err := c.Username(context.Background(), &chat.Message{Body: "Username From Client!"})
  if err != nil {
    log.Fatalf("Error when calling Username: %s", err)
  }
  log.Printf("vaibhav from the server: %s", response.Body)

  response, err = c.Name(context.Background(), &chat.Message{Body: "Message to Name!"})
  if err != nil {
    log.Fatalf("Error when calling Name: %s", err)
  }
  log.Printf("vaibhav from the server: %s", response.Body)

}