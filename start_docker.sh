docker build -t adamastor:latest .
docker run -it -p 8000:8000  -v $HOME/projects/go_personal/adamastor:/adamastor adamastor:latest
