package external

import (
	"context"
	"ewallet-ums/constans"
	"ewallet-ums/external/proto/notification"
	"ewallet-ums/helpers"
	"fmt"

	"google.golang.org/grpc"
)

func (*External) SendNotification(ctx context.Context, recipient, templateName string, placeHolder map[string]string) error {
	conn, err := grpc.Dial(helpers.GetEnv("NOTIFICATION_GRPC_HOST", ""), grpc.WithInsecure())
	if err != nil {
		return err
	}
	defer conn.Close()
	fmt.Println("Recipient:", recipient)
	fmt.Println("Template Name:", templateName)

	client := notification.NewNotificationServiceClient(conn)
	request := &notification.SendNotificationRequest{
		Recipient:    recipient,
		TemplateName: templateName,
		Placeholders: placeHolder,
	}

	response, err := client.SendNotification(ctx, request)
	if err != nil {
		return err
	}

	if response.Message != constans.SuccessMessage {
		return fmt.Errorf("failed to get response from notification: %s", response.Message)
	}

	return nil
}
