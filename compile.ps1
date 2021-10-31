# Runs the compile.sh script in ubuntu container without switch between VSCode dev container and windows env

# the VSCode container name. VSCode mounts the project under /workspaces/${projectName}
# Container name varies and if it gets rebuild we need to change the name.
$CONT_NAME = "infallible_kowalevski"

# The non-root user is missing `go/bin` in its PATH env variable, causing protoc to fail
$CMD = "cd /workspaces/hello-again && export PATH=$('$')PATH:/root/go/bin && ./compile.sh"

# Get the current container status - https://docs.docker.com/engine/reference/commandline/inspect/
$STATUS=$(docker inspect --format='{{.State.Status}}' $CONT_NAME)


Write-Output "Container status -" $STATUS

# Makes sure the container is running
if($STATUS -notcontains "running") {
    Write-Output "Trying to start the container"
    docker start $CONT_NAME
}

# Runs the command as non-root user,
docker exec $CONT_NAME sh -c $CMD
