#!/bin/bash

PROGRAM_NAME="dg"

if [[ ":$PATH:" != *":$GOPATH/bin:"* ]]; then
    echo "Adding $GOPATH/bin to PATH"
    export PATH="$GOPATH/bin:$PATH"
    
    echo 'export PATH="$GOPATH/bin:$PATH"' >> ~/.zshrc  # For Zsh
    # echo 'export PATH="$GOPATH/bin:$PATH"' >> ~/.bashrc
    echo "Remember to run 'source ~/.zshrc' (or open a new terminal) to update your PATH"
fi

# Build the program
echo "Building $PROGRAM_NAME..."
go build -o $PROGRAM_NAME

# Check if the build was successful
if [ $? -eq 0 ]; then
    echo "Build successful"
    
    # Install the program
    echo "Installing $PROGRAM_NAME..."
    go install
    
    if [ $? -eq 0 ]; then
        echo "Installation successful."
        echo "Ensure you run '$PROGRAM_NAME' from the root of your dbt project."
    else
        echo "Installation failed"
    fi
else
    echo "Build failed"
fi
