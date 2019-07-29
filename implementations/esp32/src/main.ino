#include "WiFi.h"        // ESP32 WiFi include
#include "coap_client.h" // for coap
#include "pb_common.h"   // for proto
#include "pb.h"          // for proto
// #include "pb_common.c"     // for proto
// #include "pb_encode.c"     // for proto
// #include "pb_decode.c"     // for proto
#include "pb_encode.h" // for proto
#include "notify.pb.h" // include protofile
// *pb_encode.c*, *pb_decode.c* and *pb_common.c*

const int PushButtonPin = 14;
const char *SSID = "joeWifi";                     // Wifi configurations
const char *WiFiPassword = "badPasswordsAreEasy"; // Wifi configurations

coapClient coap; //instance for coapclient
String DEVICE_SECRET_KEY = "your-device_secret_key";

IPAddress ip(0, 0, 0, 0);
int port = 5688;

uint8_t buffer[128]; // buffer for protobuf

// void callback_response(coapPacket &packet, IPAddress ip, int port);

// coap client response callback
void callback_response(coapPacket &packet, IPAddress ip, int port)
{
    char p[packet.payloadlen + 1];
    memcpy(p, packet.payload, packet.payloadlen);
    p[packet.payloadlen] = NULL;
    Serial.println("called");
    //response from coap server
    if (packet.type == 3 && packet.code == 0)
    {
        Serial.println("ping ok");
    }

    Serial.println(p);
}

// This Setup function is used to initialize everything
void setup()
{
    // Start serial connection
    Serial.begin(115200);
    Serial.print("setup");
    // This statement will declare pin 15 as digital input
    pinMode(PushButtonPin, INPUT);
    coap.response(callback_response);
    // client response callback.
    // this endpoint is single callback.
    // coap.start(9999);
}

void loop()

{
    if (WiFi.status() != WL_CONNECTED)
    {
        ConnectToWiFi();
    }
    else
    {
        checkInput();
    }
}

void ConnectToWiFi()
{

    WiFi.mode(WIFI_STA);
    WiFi.begin(SSID, WiFiPassword);
    Serial.print("Connecting to ");
    Serial.println(SSID);

    uint8_t i = 0;
    while (WiFi.status() != WL_CONNECTED)
    {
        Serial.print('.');
        delay(500);

        if ((++i % 16) == 0)
        {
            Serial.println(F(" still trying to connect"));
        }
    }
    ip = WiFi.gatewayIP();
    Serial.print(F("Connected. My IP address is: "));
    Serial.println(WiFi.localIP());
    Serial.println(WiFi.gatewayIP());
}

void checkInput()
{
    // digitalRead function stores the Push button state
    // in variable push_button_state
    int Push_button_state = digitalRead(PushButtonPin);
    // if condition checks if push button is pressed
    if (Push_button_state == HIGH)
    {
        Serial.println("high");
        sendStuff();
    }
}

void sendStuff()
{
    // root["name"] = "temperature";
    // root["data"] = 21.5;
    // root["accessToken"] = DEVICE_SECRET_KEY;

    String data;
    // root.printTo(data);
    char dataChar[data.length() + 1];
    data.toCharArray(dataChar, data.length() + 1);
    // bool state;

    //post request
    //arguments server ip address,default port,resource name, payload,payloadlength
    // int msgid = coap.post(ip, port, path, dataChar, data.length());
    coap.response(callback_response);
    int msgid = coap.get(ip, port, "a");
    msgid = coap.post(ip, port, "c", dataChar, data.length());
    Serial.println(String(msgid));

    coap.loop();

    delay(1000);
    genMessage();
}

// /*
void genMessage()
{
    ButtonMessage msg = ButtonMessage_init_zero;
    pb_ostream_t stream = pb_ostream_from_buffer(buffer, sizeof(buffer));

    msg.milli = 111;
    msg.sec = 13;
    char name[] = "urgent";
    strncpy(msg.label, name, strlen(name));
    bool status = pb_encode(&stream, ButtonMessage_fields, &msg);
    if (!status)
    {
        Serial.println("Failed to encode");
        return;
    }

    Serial.print("Message Length: ");
    Serial.println(stream.bytes_written);

    Serial.print("Message: ");

    for (int i = 0; i < stream.bytes_written; i++)
    {
        Serial.printf("%02X", buffer[i]);
    }
}
// */