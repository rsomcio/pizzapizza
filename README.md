# pizzapizza

docker build -t web:1.5 -t web:latest .

docker run --name webapp -p 5173:8080 --detach -it web:latest

docker start webapp

docker push rsomcio/web:1.2-b

kubectl set image deployments/webapp web=rsomcio/web:1.2-b
