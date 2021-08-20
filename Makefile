build:
	go build -o build/LODeditor cmd/lodeditor/main.go

package: build
	fyne package -executable ./build/LODeditor -os darwin -icon ./assets/icon.png

clean:
	rm -Rf build/*

darwin:
	GOOS=darwin GOARCH=amd64 go build -o build/darwin/LODeditor cmd/lodeditor/main.go
	cd ./build/darwin/ && fyne package -executable LODeditor -os darwin -icon ../../assets/icon.png -name LODeditor

windows:
	GOOS=windows GOARCH=amd64 go build -o build/windows/LODeditor cmd/lodeditor/main.go
	cd ./build/darwin/ && fyne package -executable LODeditor -os windows -icon ../../assets/icon.png -name LODeditor