FROM golang:onbuild


WORKDIR /app


ENV SRC_DIR=github.com/dictybase-playground/async-job-server/worker
ADD . $SRC_DIR
RUN cd $SRC_DIR; go build -o myapp; cp myapp /app/


ENTRYPOINT ["./myapp"]
