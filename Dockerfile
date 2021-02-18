FROM golang:1.15.3 AS build-env

WORKDIR /app/

ADD . .

RUN go build -o executable -mod vendor .


FROM gcr.io/distroless/base-debian10

WORKDIR /app/

COPY --from=build-env /app/executable .
COPY --from=build-env /app/database/mobile_applications.json ./database/mobile_applications.json
COPY --from=build-env /app/database/security_check_specs.json ./database/security_check_specs.json

EXPOSE 1323

CMD ["/app/executable"]