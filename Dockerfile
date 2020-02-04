FROM scratch

LABEL maintainer="Stefano Pirrello <spirrello@gmail.com>"

WORKDIR /

COPY /url-test .

ENTRYPOINT ["/url-test"]
