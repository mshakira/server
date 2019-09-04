FROM scratch
ADD server.crt /
ADD server.key /
ADD bin/server /
CMD ["/server"]
