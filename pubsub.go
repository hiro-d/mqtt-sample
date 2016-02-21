package main

import (
    "fmt"
    MQTT "github.com/eclipse/paho.mqtt.golang"
    "os"
    "time"
)

var f MQTT.MessageHandler = func(client *MQTT.Client, msg MQTT.Message) {
    fmt.Printf("TOPIC: %s\n", msg.Topic())
    fmt.Printf("MSG: %s\n", msg.Payload())
}

func main() {
    user     := "username@github"
    password := "password"
    broker   := "broker"

    opts := MQTT.NewClientOptions().AddBroker(broker)
    opts.SetClientID("go-simple")
    opts.SetDefaultPublishHandler(f)
    opts.SetUsername(user)
    opts.SetPassword(password)

    c := MQTT.NewClient(opts)
    if token := c.Connect(); token.Wait() && token.Error() != nil {
        panic(token.Error())
    }

    topic := "path/to/topic"
    if token := c.Subscribe(topic, 0, nil); token.Wait() && token.Error() != nil {
        fmt.Println(token.Error())
        os.Exit(1)
    }

    for i := 0; i < 1; i++ {
        text := fmt.Sprintf("this is msg #%d!", i)
        token := c.Publish(topic, 0, false, text)
        token.Wait()
    }

    time.Sleep(3 * time.Second)

    c.Disconnect(250)
}
