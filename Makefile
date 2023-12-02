dir:
	mkdir day_$(day)
	touch day_$(day)/main.go
	touch day_$(day)/input.txt
	touch day_$(day)/input_test.txt

run:
	cd '$(CURDIR)/day_$(day)' && go run main.go

