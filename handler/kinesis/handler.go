package kinesis

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/kinesis"
	"github.com/phogolabs/log"
	"github.com/rogpeppe/fastuuid"
	k "github.com/tj/go-kinesis"
)

var _ log.Handler = &Handler{}

// Config of for this handler
type Config = k.Config

// Handler implementation.
type Handler struct {
	appName  string
	producer *k.Producer
	gen      *fastuuid.Generator
}

// New handler sending logs to Kinesis. To configure producer options or pass your
// own AWS Kinesis client use NewConfig instead.
func New(stream string) *Handler {
	return NewConfig(k.Config{
		StreamName: stream,
		Client:     kinesis.New(session.New(aws.NewConfig())),
	})
}

// NewConfig handler sending logs to Kinesis. The `config` given is passed to the batch
// Kinesis producer, and a random value is used as the partition key for even distribution.
func NewConfig(config Config) *Handler {
	producer := k.New(config)
	producer.Start()

	return &Handler{
		producer: producer,
		gen:      fastuuid.MustNewGenerator(),
	}
}

// Handle implements log.Handler.
func (h *Handler) Handle(e *log.Entry) {
	data, err := json.Marshal(e)
	if err != nil {
		fmt.Fprintln(os.Stdout, err)
		return
	}

	uuid := h.gen.Next()
	key := base64.StdEncoding.EncodeToString(uuid[:])

	if err := h.producer.Put(data, key); err != nil {
		fmt.Fprintln(os.Stdout, err)
	}
}
