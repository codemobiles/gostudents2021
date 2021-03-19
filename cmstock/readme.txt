go mod init main
https://jmeubank.github.io/tdm-gcc/download/


go get gorm.io/driver/mysql

# create and run react 
npx create-react-app my-app
cd my-app
npm start

# build or deploy react
npm run build

# Docker instruction
docker build -t cmstock:1.0 .       
docker run -it -p  8081:8081 cmstock:1.0


# docker swarm
docker swarm init
docker service create --name cmstock_service --replicas 3 --publish 8081:8081 cmstock:1.0
docker service ls 
docker service rm <service_name>
docker service scale <service_name>=5