all: unix_clean unix_build

win_build:
	@echo Compling GO sources in progress...
	@cd src && go build -v -o ../target/
	@echo Copy the configuration file...
	@if not exist target\config mkdir target\config
	@copy /V src\config\application.yaml target\config
	
unix_build:
	@echo Compling GO sources in progress...
	@cd src && go build -v -o ../target/
	@echo Copy the configuration file...
	@mkdir -p target/config
	@cp -f src/config/application.yaml target/config/

win_clean:
	@if exist target rd /s /q target

unix_clean:
	@rm -rf target

godoc:
	@cd src && godoc -http=:6060

run:
	@cd target && ./sample-rest-api
	
test:
	@cd src && go test -v ./...
