go mod init main
https://jmeubank.github.io/tdm-gcc/download/

npx create-react-app my-app
cd my-app
npm start

docker build -t cmstock:1.0 .       
docker run -it -p  8081:8081 cmstock:1.0
docker swarm init
docker service create --replicas 3 --publish 8081:8081 cmstock:1.0