package publisher

import (
	cwrap "github.com/richardrigby/rclgo/internal"
	"github.com/richardrigby/rclgo/node"
	"github.com/richardrigby/rclgo/rcl"
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
	node node.Node,
	topicName string,
	msg types.MessageTypeSupport,
) error {

	ret := cwrap.RclPublisherInit(
		p.rclPublisher,
		node.RclNode,
		msg.ROSIdlMessageTypeSupport,
		topicName,
		publisherOptions.rclPublisherOptions,
	)

	if ret != 0 {
		return rcl.NewErr("cwrap.RclInitOptionsInit", int(ret))
	}

	return nil
}

func (p *Publisher) PublisherFini(node node.Node) error {
	ret := cwrap.RclPublisherFini(p.rclPublisher, node.RclNode)
	if ret != 0 {
		return rcl.NewErr("C.rcl_publisher_fini", int(ret))
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
		return rcl.NewErr("cwrap.Publish", ret)
	}

	return nil
}

func (p *Publisher) IsValid() bool {
	return cwrap.RclPublisherIsValid(p.rclPublisher)
}
