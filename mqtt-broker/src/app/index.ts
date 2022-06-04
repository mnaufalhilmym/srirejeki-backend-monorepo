import aedes, { Client, AuthErrorCode, AuthenticateError } from "aedes";
import net from "net";
import axios from "axios";

const backendProtocol = process.env.BACKEND_PROTOCOL || "http";
const backendHost = process.env.BACKEND_HOST || "localhost";
const backendPort = process.env.BACKEND_PORT || "80";

const authenticate = async (
  client: Client,
  username: Readonly<string>,
  password: Readonly<Buffer>,
  callback: (error: AuthenticateError | null, success: boolean | null) => void
) => {
  if (!client) {
    const err = new Error("Client not provide it's identifier");
    (err as AuthenticateError).returnCode = AuthErrorCode.NOT_AUTHORIZED;
    console.error(err.message);
    callback(err as AuthenticateError, false);
    return;
  }
  if (client.id !== process.env.MQTT_SERVER_CLIENT_ID) {
    let authorized = false;
    try {
      const res = await axios.post(
        backendProtocol + "://" + backendHost + ":" + backendPort + "/mcu/auth",
        { deviceId: client.id }
      );
      if (res.status.toString()[0] === "2") {
        authorized = true;
      }
    } catch (error) {
      (error as AuthenticateError).returnCode =
        AuthErrorCode.SERVER_UNAVAILABLE;
      console.error(`Can't authorize client id: ${client.id}. Error: ${error}`);
      callback(error as AuthenticateError, false);
      return;
    }

    if (!authorized) {
      const error = new Error(`Client id ${client.id} is not registered`);
      (error as AuthenticateError).returnCode = AuthErrorCode.NOT_AUTHORIZED;
      console.error(`Can't authorize client id: ${client.id}. Error: ${error}`);
      callback(error as AuthenticateError, false);
      return;
    }
  }
  console.info(`Client id ${client.id} connection accepted`);
  callback(null, true);
};

const aedesInstance = aedes({ authenticate: authenticate });
const server = net.createServer(aedesInstance.handle);

// emitted when a client connects to the broker
aedesInstance.on("client", function (client) {
  const date = new Date();
  console.log(
    `[CLIENT_CONNECTED ${date.getHours()}:${date.getMinutes()}:${date.getSeconds()}] Client ${
      client ? client.id : client
    } connected to broker ${aedesInstance.id}`
  );
});

// emitted when a client disconnects from the broker
aedesInstance.on("clientDisconnect", function (client) {
  const date = new Date();
  console.log(
    `[CLIENT_DISCONNECTED ${date.getHours()}:${date.getMinutes()}:${date.getSeconds()}] Client ${
      client ? client.id : client
    } disconnected from the broker ${aedesInstance.id}`
  );
});

// emitted when a client subscribes to a message topic
aedesInstance.on("subscribe", function (subscriptions, client) {
  const date = new Date();
  console.log(
    `[TOPIC_SUBSCRIBED ${date.getHours()}:${date.getMinutes()}:${date.getSeconds()}] Client ${
      client ? client.id : client
    } subscribed to topics: ${subscriptions
      .map((s) => s.topic)
      .join(",")} on broker ${aedesInstance.id}`
  );
});

// emitted when a client unsubscribes from a message topic
aedesInstance.on("unsubscribe", function (subscriptions, client) {
  const date = new Date();
  console.log(
    `[TOPIC_UNSUBSCRIBED ${date.getHours()}:${date.getMinutes()}:${date.getSeconds()}] Client ${
      client ? client.id : client
    } unsubscribed to topics: ${subscriptions.join(",")} from broker ${
      aedesInstance.id
    }`
  );
});

// emitted when a client publishes a message packet on the topic
aedesInstance.on("publish", function (packet, client) {
  if (client) {
    const date = new Date();
    console.log(
      `[MESSAGE_PUBLISHED ${date.getHours()}:${date.getMinutes()}:${date.getSeconds()}] Client ${
        client ? client.id : "BROKER_" + aedesInstance.id
      } has published message ${packet.payload} on ${packet.topic} to broker ${
        aedesInstance.id
      }`
    );
    (async () => {
      const durations: string[] = [];
      if (date.getMinutes() === 0 && date.getSeconds() === 0) {
        durations.push("hour");
        if (date.getHours() === 0) {
          durations.push("day");
          if (date.getDate() === 1) {
            durations.push("month");
          }
        }
      }
      if (durations.length === 0) {
        return;
      }
      const topic = packet.topic.split("/");
      const topicFrom = topic.length >= 2 ? topic[1] : "";
      if (topicFrom === "client") {
        const deviceId = topic[2];
        const type = topic[3];
        try {
          const res = await axios.post(
            backendProtocol +
              "://" +
              backendHost +
              ":" +
              backendPort +
              "/data/snapshot",
            { type, data: packet.payload.toString(), deviceId, durations }
          );
          if (res.status.toString()[0] !== "2") {
            console.error(
              `Failed to post snapshot data.\nStatus: ${res.status}\nResponse: ${res.data}`
            );
          }
        } catch (error) {
          console.error(error);
        }
      }
    })();
  }
});

export default server;
