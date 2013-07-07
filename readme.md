# WeatherLink CRC in Go

From http://www.davisnet.com/support/weather/download/VantageSerialProtocolDocs_v261.pdf

> The CRC checking used by the WeatherLink is based on the CRC-CCITT standard. The heart of 
> the method involves a CRC-accumulator that uses the following formula on each successive data 
> byte.
>
> crc = table[(crc >> 8) ^ data] ^ (crc << 8)
>
> After all the data bytes have been "accumulated", there will be a two byte CRC checksum 
> that will get processed in the same manner as the data bytes. If there has been no transmission 
> error, then the final CRC-accumulator value will be 0 (assuming it was set to zero before 
> accumulating data).


Weather_crc implements hash.Hash, but adds a Sum16() function to get the current value.


# Example

```go
package main

import (
  "github.com/derekpitt/weather_crc"
)

func main() {
  crc := weather_crc.New()
  crc.Write([]byte{0xc6})
  print(crc.Sum16())
}

```
