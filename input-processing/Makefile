clean:
	rm *.out
	rm -f input_5gb.txt input_5gb_nl.txt

test:
	./test.sh

package:
	docker build -t input-processor:1.0 .

profile:
	go tool pprof -http=:8080 heap_profile.out

dockerRun:
	docker run --rm -it --memory="1g" --entrypoint /bin/sh input-processor:1.0

largeTestFile:
	base64 /dev/urandom | head -c 5000000000 > input_5gb_nl.txt
	cat input_5gb_nl.txt | tr -d '\n' > input_5gb.txt

