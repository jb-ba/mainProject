#include "WiFi.h"        // ESP32 WiFi include
#include "coap_client.h" // for coap
#include "pb_common.h"   // for proto
#include "pb.h"          // for proto
#include "pb_encode.h"   // for proto
#include "notify.pb.h"   // include protofile
#include <ESPmDNS.h>     // for OTA update
#include <WiFiUdp.h>     // for OTA update
#include <ArduinoOTA.h>  // for OTA update

const int PushButtonPin = 14;
const char *SSID = "raspi-webgui";     // Wifi configurations
const char *WiFiPassword = "ChangeMe"; // Wifi configurations

coapClient coap; //instance for coapclient
String DEVICE_SECRET_KEY = "your-device_secret_key";

IPAddress ip(10, 3, 141, 1);
int port = 30002;
// int port = 5688;

// for led control
int ledPin = 27;
bool ledOn = false;

// coap client response callback
void callback_response(coapPacket &packet, IPAddress ip, int port)
{
    char p[packet.payloadlen + 1];
    memcpy(p, packet.payload, packet.payloadlen);
    p[packet.payloadlen] = NULL;
    Serial.println("response");
    //response from coap server
    if (packet.type == 3 && packet.code == 0)
    {
        Serial.println("ping ok");
    }
    Serial.println(p);
}
// observe(IPAddress ip, int port, char *url, uint8_t optionbuffer)
void observeLedStatus()
{
    int msgid = coap.observe(ip, port, "observe", 0);
    Serial.println("observing");
    Serial.println(msgid);
}

// This Setup function is used to initialize everything
void setup()
{
    Serial.begin(115200);
    Serial.print("setup");
    pinMode(ledPin, OUTPUT);
    pinMode(PushButtonPin, INPUT);
    coap.response(callback_response);
    // otaUpdate();
}

void loop()

{
    if (WiFi.status() != WL_CONNECTED)
    {
        ConnectToWiFi();
        if (WiFi.status() != WL_CONNECTED)
        {
            delay(2000);
        }
        else
        {
            observeLedStatus();
        }
    }
    else
    {
        coap.loop();
        ensureLed();
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
        delay(500);
    }
    ip = WiFi.gatewayIP();
    Serial.println(ip);
}

// just checks that bool and ledPin output are the same
void ensureLed()
{
    if (ledOn)
    {
        digitalWrite(ledPin, HIGH);
    }
    else
    {
        digitalWrite(ledPin, LOW);
    }
}

// checks if push button is pressed
void checkInput()
{
    if (digitalRead(PushButtonPin) == HIGH)
    {
        if (ledOn)
        {
            digitalWrite(ledPin, LOW);
            Serial.println("LED off");
            ledOn = false;
        }
        else
        {
            digitalWrite(ledPin, HIGH);
            Serial.println("LED on");
            ledOn = true;
        }
        sendClick();
    }
}

void sendClick()
{
    uint8_t buffer[128];

    ButtonMessage msg = ButtonMessage_init_zero;
    pb_ostream_t stream = pb_ostream_from_buffer(buffer, sizeof(buffer));

    msg.building = 111;
    msg.room = 13;
    msg.ledOn = ledOn;
    char name[] = "front";
    strncpy(msg.label, name, strlen(name));
    bool status = pb_encode(&stream, ButtonMessage_fields, &msg);
    if (!status)
    {
        Serial.println("Failed to encode");
        return;
    }

    int msgid = coap.post(ip, port, "clicked", (char *)buffer, stream.bytes_written);
    delay(1000);
}

// void otaUpdate()
// {
//     ArduinoOTA
//         .onStart([]() {
//             String type;
//             if (ArduinoOTA.getCommand() == U_FLASH)
//                 type = "sketch";
//             else // U_SPIFFS
//                 type = "filesystem";

//             // NOTE: if updating SPIFFS this would be the place to unmount SPIFFS using SPIFFS.end()
//             Serial.println("Start updating " + type);
//         })
//         .onEnd([]() {
//             Serial.println("\nEnd");
//         })
//         .onProgress([](unsigned int progress, unsigned int total) {
//             Serial.printf("Progress: %u%%\r", (progress / (total / 100)));
//         })
//         .onError([](ota_error_t error) {
//             Serial.printf("Error[%u]: ", error);
//             if (error == OTA_AUTH_ERROR)
//                 Serial.println("Auth Failed");
//             else if (error == OTA_BEGIN_ERROR)
//                 Serial.println("Begin Failed");
//             else if (error == OTA_CONNECT_ERROR)
//                 Serial.println("Connect Failed");
//             else if (error == OTA_RECEIVE_ERROR)
//                 Serial.println("Receive Failed");
//             else if (error == OTA_END_ERROR)
//                 Serial.println("End Failed");
//         });

//     ArduinoOTA.begin();

//     Serial.println("Ready");
//     Serial.print("IP address: ");
//     Serial.println(WiFi.localIP());
// }
