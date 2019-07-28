#include "WiFi.h"        // ESP32 WiFi include
#include "coap_client.h" //
// #include "WiFiConfig.h" // My WiFi configuration.
// this will assign the name PushButton to pin numer 15
const int PushButton = 14;
// Wifi configurations
const char *SSID = "joeWifi";
const char *WiFiPassword = "badPasswordsAreEasy";
//instance for coapclient
coapClient coap;
String DEVICE_SECRET_KEY = "your-device_secret_key";

IPAddress ip(0, 0, 0, 0);
int port = 5688;

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
    pinMode(PushButton, INPUT);
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
    int Push_button_state = digitalRead(PushButton);
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

    // String data;
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
}