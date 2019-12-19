FROM alpine
ADD microtest-srv /microtest-srv
ENTRYPOINT [ "/microtest-srv" ]
