# Gatorloop Webapp

### Pre-installation
Before you can run the project, you will need to have golang installed and configured (https://golang.org/doc/code.html). You will also need to install mysql, and create a running instance of the gatorloop database on localhost:3306.

### Installation Instructions
1. Clone the project into $GOPATH/src/github.com
2. Install Godep
    ```
     go get github.com/tools/godep
     ```
3. Install godep dependencies. First cd into the project, then run:
    ```
    godep restore
    ```
4. Install the project. Run from the project directory:
    ```
    go install
    ```
5. Run the executable in $GOPATH/bin.
    ```
    gatorloopwebapp
    ```
