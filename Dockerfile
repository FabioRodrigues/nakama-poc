FROM heroiclabs/nakama-pluginbuilder:3.21.1 AS go-builder

ENV GO111MODULE on
ENV CGO_ENABLED 1

WORKDIR /backend

COPY go.mod .
COPY main.go .
COPY vendor/ vendor/
COPY constants/ constants/
COPY dtos/ dtos/
COPY entities/ entities/
COPY handlers/ handlers/
COPY services/ services/
COPY wrappers/ wrappers/


RUN go build --trimpath --mod=vendor --buildmode=plugin -o ./backend.so

FROM registry.heroiclabs.com/heroiclabs/nakama:3.21.1

COPY --from=go-builder /backend/backend.so /nakama/data/modules/
COPY local.yml /nakama/data/
COPY core/1.0.1.json /nakama/core/1.0.1.json