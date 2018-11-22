# goproject

This application is a modified version of the Microclimate Go template that's been configured to support package
management with the Dep package manager. New packages can be added to the Gopkg.toml file, and they will be pulled in during the Docker build

## External Editor Support
Most Go editors and package managers require the project to be under the system's Gopath, which requires some additional steps to work with Microclimate. To use this project in Microclimate with an external editor, do the following:
#### Prerequisites
- Go installed on your local machine
- `GOPATH` environment set (usually to `~/go`)
- Editor installed on your machine with Go support (VS Code with Go extension recommended)

#### Steps
1. Import the application into Microclimate
2. Update all references to `goproject` to the name of your Microclimate project (Should be just Dockerfile and main.go)
3. Once the application has been imported, create a symlink from the application (e.g. `/microclimate-workspace/goproject`) to a folder under your GOPATH/src:
    ```
    ln -s /microclimate-workspace/<project> $GOPATH/src/<project>
    ```
4. Open the project in your editor

**Note**: As you add dependencies to your application, you should call `dep ensure -v` on the local copy of your project so that the packages are visible to the editor while you're making changes