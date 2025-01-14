package main

import (
	"github.com/surgemq/message"
	"github.com/surgemq/surgemq/service"
)

func main() {

	// Instantiates a new Client
	c := &service.Client{}

	// Creates a new MQTT CONNECT message and sets the proper parameters
	msg := message.NewConnectMessage()
	msg.SetWillQos(1)
	msg.SetVersion(4)
	msg.SetCleanSession(true)
	msg.SetClientId([]byte("surgemq"))
	msg.SetKeepAlive(10)
	msg.SetWillTopic([]byte("will"))
	msg.SetWillMessage([]byte("send me home"))
	msg.SetUsername([]byte("surgemq"))
	msg.SetPassword([]byte("verysecret"))

	// Connects to the remote server at 127.0.0.1 port 1883
	c.Connect("tcp://127.0.0.1:1883", msg)

	// Creates a new SUBSCRIBE message to subscribe to topic "abc"
	submsg := message.NewSubscribeMessage()
	submsg.AddTopic([]byte("abc"), 0)

	// Subscribes to the topic by sending the message. The first nil in the function
	// call is a OnCompleteFunc that should handle the SUBACK message from the server.
	// Nil means we are ignoring the SUBACK messages. The second nil should be a
	// OnPublishFunc that handles any messages send to the client because of this
	// subscription. Nil means we are ignoring any PUBLISH messages for this topic.
	c.Subscribe(submsg, nil, nil)

	// Creates a new PUBLISH message with the appropriate contents for publishing
	pubmsg := message.NewPublishMessage()
	// pubmsg.SetPacketId(pktid)
	pubmsg.SetTopic([]byte("abc"))
	pubmsg.SetPayload(make([]byte, 1024))
	// pubmsg.SetQoS(qos)

	// Publishes to the server by sending the message
	c.Publish(pubmsg, nil)

	// Disconnects from the server
	c.Disconnect()

}
