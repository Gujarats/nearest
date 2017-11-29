go test -v -race ./...
RESULT=$?
[ $RESULT -ne 0 ] && exit 1
exit 0
