all: win_clean win_build

win_build:
	@if not exist target mkdir target
	@if not exist target\config mkdir target\config
	@echo Copy of the configuration file...
	@copy /V src\application.yaml target\config
	@echo Compling GO sources in progress...
	@cd src && go build -v -o ../target
	
unix_build:
	@mkdir target/config
	@echo Copy of the configuration file...
	@cp -f src/application.yaml target/config
	@echo Compling GO sources in progress...
	@cd src && go build -v -o ../target

win_clean:
	@if exist target rd /s /q target

unix_clean:
	@rm -rf target

godoc:
	@cd src && godoc -http=:6060

run:
	@cd target && sample-rest-api
	
test:
	@cd src && go test -v ./...