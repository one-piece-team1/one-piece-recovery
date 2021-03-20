##################
# STEP 1 builder #
##################
FROM golang:alpine AS build

ADD . /images
WORKDIR /images
RUN go build -o /bin/starter

##################
# STEP 2 scratch #
##################
FROM scratch

COPY --from= build /bin/starter /bin/starter
ENTRYPOINT ["/bin/starter"]