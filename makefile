all:
	# test STO project
	cd STO && go test
	# test Go-Pattern project, using ./... to run all tests in subdirectories
	# see https://stackoverflow.com/a/21725603
	cd Go-pattern &&  go test ./...