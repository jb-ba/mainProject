# PiCoAP

This is an implementation of the CoAP server standard for the raspberry pi.

## Docker

To run it in docker you need to specify UDP as L4 protocol. 
I.e.: <pre>docker run --name picoap -p 5688:5688<b><i><u>/udp</u></i></b> jbba/picoap:v0.0.1arm</pre>

## Proto
Do `protoc --go_out=. *.proto` in proto folder