package nats

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ROHITHSAKTHIVEL/Nats_sample/dto"

	im "github.com/nats-io/nats.go"
)

func Publish(ctx context.Context, request *dto.Request) {
	fmt.Println("Publishing.....")

	if err := validateVariable(request.Name); err != nil {
		fmt.Println("Validation error:", err)
		return
	}

	nc, err := im.Connect("localhost:4222")
	if err != nil {
		fmt.Println("Nats Connection Error:", err)
		return
	}
	defer nc.Close()

	data, encodeErr := Encode(request)
	if encodeErr != nil {
		fmt.Println("Encoding error:", encodeErr)
		return
	}

	if err := nc.Publish(request.Topic, data); err != nil {
		fmt.Println("Error publishing message:", err)
	}
	fmt.Println("publiesfed")
}

func Subscribe(ctx context.Context, req *dto.Request) {
	fmt.Println("Subscribing.....")

	nc, err := im.Connect("localhost:4222")
	if err != nil {
		fmt.Println("Nats Connection Error:", err)
		return
	}
	defer nc.Close()

	if _, err := nc.Subscribe(req.Topic, func(msg *im.Msg) {
		res, deerr := Decode(msg.Data)
		if deerr != nil {
			fmt.Println("Error decoding message:", deerr)
			return
		}
		fmt.Println("Received message:", res)
	}); err != nil {
		fmt.Println("Error subscribing to subject 'sample':", err)
	}

}

func Encode(i *dto.Request) ([]byte, error) {
	val, err := json.Marshal(i)
	if err != nil {
		return nil, err
	}
	return val, nil
}

func Decode(i []byte) (*dto.Request, error) {
	var data dto.Request
	err := json.Unmarshal(i, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func validateVariable(res interface{}) error {
	if res == nil {
		return errors.New("Invalid Input")
	}
	return nil
}
