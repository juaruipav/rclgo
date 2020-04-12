package rclgo

import (
	"github.com/richardrigby/rclgo/err"
	cwrap "github.com/richardrigby/rclgo/internal"
	"github.com/richardrigby/rclgo/types"
)

type Publisher struct {
	rclPublisher *cwrap.RclPublisher
}

type PublisherOptions struct {
	rclPublisherOptions *cwrap.RclPublisherOptions
}

func NewZeroInitializedPublisher() Publisher {
	zeroPublisher := cwrap.RclGetZeroInitializedPublisher()
	return Publisher{&zeroPublisher}
}

func NewPublisherDefaultOptions() PublisherOptions {
	defOpts := cwrap.RclPublisherGetDefaultOptions()
	return PublisherOptions{&defOpts}
}

func (p *Publisher) GetTopicName() string {
	return cwrap.RclPublisherGetTopicName(p.rclPublisher)
}

func (p *Publisher) Init(
	publisherOptions PublisherOptions,
	node Node,
	topicName string,
	msg types.MessageTypeSupport,
) error {

	ret := cwrap.RclPublisherInit(
		p.rclPublisher,
		node.rclNode,
		msg.ROSIdlMessageTypeSupport,
		topicName,
		publisherOptions.rclPublisherOptions,
	)

	if ret != 0 {
		return err.NewErr("RclInitOptionsInit", int(ret))
	}

	return nil
}

func (p *Publisher) PublisherFini(node Node) error {
	ret := cwrap.RclPublisherFini(p.rclPublisher, node.rclNode)
	if ret != 0 {
		return err.NewErr("RclPublisherFini", ret)
	}

	return nil
}

func (p *Publisher) Publish(msg types.MessageTypeSupport, data types.MessageData) error {

	ret := cwrap.RclPublish(
		p.rclPublisher,
		msg.ROSIdlMessageTypeSupport,
		data.Data,
	)

	if ret != 0 {
		return err.NewErr("RclPublish", ret)
	}

	return nil
}

func (p *Publisher) IsValid() bool {
	return cwrap.RclPublisherIsValid(p.rclPublisher)
}
