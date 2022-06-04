import app from "./app";

const port = process.env.MQTT_BROKER_PORT || 1883;

app.listen(port, function () {
  console.log("server started and listening on port ", port);
});

export default app;
