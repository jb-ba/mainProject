# This is the main application for esp32

Example file used is from the official 
https://github.com/espressif/esp-idf/tree/e6afe28bafe5db5ab79fae213f2e8e1ccd9f937c/examples/protocols/coap_client/main
https://github.com/automote/ESP-CoAP

## Arduino libraries
Can be found under /snap/arduino-mhall119/current/libraries

## Protobufs
For linux with amd64 install x86 dependencies
https://code.google.com/archive/p/nanopb/issues/146
sudo apt-get install lib32stdc++6

Compile .proto File in older with dependencies
../generator-bin/protoc --proto_path=. --nanopb_out=. notify.proto
https://techtutorialsx.com/2018/10/19/esp32-esp8266-arduino-protocol-buffers/