dir = day_$(shell seq -f "%02g" $(day) $(day))
dir:
	mkdir $(CURDIR)/$(dir)
	touch $(CURDIR)/$(dir)/main.go
	touch $(CURDIR)/$(dir)/input_test.txt
	curl https://adventofcode.com/2023/day/$(day)/input -b $$(grep '^session=.*$$' $(CURDIR)/.env) > $(CURDIR)/$(dir)/input.txt

run:
	cd $(CURDIR)/$(dir) && go run main.go

input:
	touch $(CURDIR)/$(dir)/input_test.txt
	curl https://adventofcode.com/2023/day/$(day)/input -b $$(grep '^session=.*$$' $(CURDIR)/.env) > $(CURDIR)/$(dir)/input.txt
