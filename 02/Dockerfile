FROM golang:alpine as build
WORKDIR /src
COPY . /src
RUN cd /src && go build -o missing-numbers --tags netgo --ldflags '-extldflags "-lm -lstdc++ -static"'

FROM scratch
WORKDIR /app
COPY --from=build /src/missing-numbers /app/
ENTRYPOINT ["/app/missing-numbers"]

