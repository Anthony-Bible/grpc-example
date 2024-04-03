From golang:1.22.1
Add . /src
RUN cd /src && go build -o myapp
CMD ["/src/myapp"]

