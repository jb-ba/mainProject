int ledPin = 12;

void setup()
{
    // put your setup code here, to run once:
    Serial.begin(115200);
    Serial.print("with power test");
    Serial.print("setupTest");
    pinMode(ledPin, OUTPUT);
}

void loop()
{
    digitalWrite(ledPin, HIGH);
    Serial.println("hallo");
    // put your main code here, to run repeatedly:
}