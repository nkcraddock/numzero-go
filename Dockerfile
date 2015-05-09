FROM scratch
ADD build/numzero /numzero
EXPOSE 3001
ENTRYPOINT ["/numzero"]

