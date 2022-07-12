# loadBalancingGrpcApi
Build a simple system balancing Nginx - Grpc 
<p>Test Load banlancing API use Vegeta package Golang</p>

echo "GET http://192.168.0.157" | ./vegeta.exe attack -duration=30s -rate=10 -output=results-veg-httpbin-get.bin && cat results-veg-httpbin-get.bin | ./vegeta.exe plot --title="HTTP Bin GET 10 rps for 30 seconds" > http-bin-get-10rps-30seconds.html
