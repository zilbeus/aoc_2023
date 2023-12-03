dir:
	mkdir day_$(day)
	touch day_$(day)/main.go
	touch day_$(day)/input_test.txt
	curl https://adventofcode.com/2023/day/$(day)/input -b $$(grep '^session=.*$$' $(CURDIR)/.env) > day_$(day)/input.txt

run:
	cd '$(CURDIR)/day_$(day)' && go run main.go

